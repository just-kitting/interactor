#!/usr/bin/env python3
"""Serve the MicroBlocks web app over HTTPS from the checked-out webapp directory."""

from __future__ import annotations

import argparse
import http.server
import json
import os
import ssl
import subprocess
import threading
import time
from collections import deque
from pathlib import Path
from urllib.parse import parse_qs, urlparse


ROOT_DIR = Path(__file__).resolve().parents[1]
WEBAPP_DIR = ROOT_DIR / "components" / "microblocks-smallvm" / "chromeApp" / "webapp"
CERT_DIR = ROOT_DIR / ".cache" / "microblocks-web"
CERT_FILE = CERT_DIR / "cert.pem"
KEY_FILE = CERT_DIR / "key.pem"


class I2CBridge:
    def __init__(self) -> None:
        self._cond = threading.Condition()
        self._next_id = 1
        self._pending: deque[dict[str, object]] = deque()
        self._responses: dict[int, list[int]] = {}

    def queue_transaction(
        self,
        *,
        address: int,
        write: list[int],
        timeout_ms: int,
    ) -> tuple[int, list[int] | None]:
        deadline = time.monotonic() + (timeout_ms / 1000.0)
        with self._cond:
            request_id = self._next_id
            self._next_id += 1
            self._pending.append(
                {
                    "id": request_id,
                    "address": address,
                    "write": write,
                    "timeout_ms": timeout_ms,
                }
            )
            self._cond.notify_all()

            while True:
                if request_id in self._responses:
                    return request_id, self._responses.pop(request_id)
                remaining = deadline - time.monotonic()
                if remaining <= 0:
                    self._pending = deque(
                        request for request in self._pending if request["id"] != request_id
                    )
                    return request_id, None
                self._cond.wait(timeout=remaining)

    def next_request(self, timeout_ms: int) -> dict[str, object] | None:
        deadline = time.monotonic() + (timeout_ms / 1000.0)
        with self._cond:
            while not self._pending:
                remaining = deadline - time.monotonic()
                if remaining <= 0:
                    return None
                self._cond.wait(timeout=remaining)
            return self._pending.popleft()

    def store_response(self, request_id: int, response: list[int]) -> bool:
        with self._cond:
            self._responses[request_id] = response
            self._cond.notify_all()
            return True


BRIDGE = I2CBridge()


def ensure_certificate() -> None:
    CERT_DIR.mkdir(parents=True, exist_ok=True)
    if CERT_FILE.exists() and KEY_FILE.exists():
        return
    subprocess.run(
        [
            "openssl",
            "req",
            "-x509",
            "-newkey",
            "rsa:2048",
            "-keyout",
            str(KEY_FILE),
            "-out",
            str(CERT_FILE),
            "-days",
            "3650",
            "-nodes",
            "-subj",
            "/CN=beaglebadge.local",
        ],
        check=True,
        stdout=subprocess.DEVNULL,
        stderr=subprocess.DEVNULL,
    )


class MicroBlocksHandler(http.server.SimpleHTTPRequestHandler):
    def __init__(self, *args, directory: str | None = None, **kwargs) -> None:
        super().__init__(*args, directory=str(WEBAPP_DIR if directory is None else directory), **kwargs)

    def _json_body(self) -> dict[str, object]:
        length = int(self.headers.get("Content-Length", "0"))
        raw = self.rfile.read(length) if length else b"{}"
        return json.loads(raw.decode("utf-8"))

    def _send_json(self, status: int, payload: dict[str, object]) -> None:
        body = json.dumps(payload).encode("utf-8")
        self.send_response(status)
        self.send_header("Content-Type", "application/json")
        self.send_header("Content-Length", str(len(body)))
        self.send_header("Cache-Control", "no-store")
        self.end_headers()
        self.wfile.write(body)

    def do_GET(self) -> None:
        parsed = urlparse(self.path)
        if parsed.path == "/api/i2c/next":
            params = parse_qs(parsed.query)
            timeout_ms = int(params.get("timeout_ms", ["1000"])[0])
            request = BRIDGE.next_request(timeout_ms)
            self._send_json(200, {"request": request})
            return
        if parsed.path == "/api/i2c/status":
            self._send_json(200, {"ok": True})
            return
        self.path = parsed.path
        super().do_GET()

    def do_POST(self) -> None:
        parsed = urlparse(self.path)
        if parsed.path == "/api/i2c/transaction":
            body = self._json_body()
            address = int(body.get("address", 0))
            timeout_ms = int(body.get("timeout_ms", 5000))
            write = [int(value) & 0xFF for value in body.get("write", [])]
            request_id, response = BRIDGE.queue_transaction(
                address=address,
                write=write,
                timeout_ms=timeout_ms,
            )
            if response is None:
                self._send_json(
                    504,
                    {
                        "ok": False,
                        "id": request_id,
                        "error": "timed out waiting for i2c target response",
                    },
                )
                return
            self._send_json(200, {"ok": True, "id": request_id, "response": response})
            return

        if parsed.path == "/api/i2c/respond":
            body = self._json_body()
            request_id = int(body["id"])
            response = [int(value) & 0xFF for value in body.get("response", [])]
            BRIDGE.store_response(request_id, response)
            self._send_json(200, {"ok": True})
            return

        self._send_json(404, {"ok": False, "error": "unknown endpoint"})


def main() -> int:
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument("--host", default="0.0.0.0")
    parser.add_argument(
        "--public-host",
        default=None,
        help="hostname or IP to show in the startup URL; defaults to --host",
    )
    parser.add_argument("--port", type=int, default=8443)
    parser.add_argument("--http", action="store_true", help="serve without TLS")
    args = parser.parse_args()

    os.chdir(WEBAPP_DIR)
    server = http.server.ThreadingHTTPServer((args.host, args.port), MicroBlocksHandler)

    scheme = "http"
    if not args.http:
        ensure_certificate()
        context = ssl.SSLContext(ssl.PROTOCOL_TLS_SERVER)
        context.load_cert_chain(certfile=str(CERT_FILE), keyfile=str(KEY_FILE))
        server.socket = context.wrap_socket(server.socket, server_side=True)
        scheme = "https"

    public_host = args.public_host or args.host
    if public_host == "0.0.0.0":
        public_host = "localhost"

    print(f"Serving MicroBlocks web app from {WEBAPP_DIR}")
    print(f"Listening on {scheme}://{args.host}:{args.port}/")
    print(f"Open {scheme}://{public_host}:{args.port}/microblocks.html")
    print(f"I2C bridge endpoint: {scheme}://{public_host}:{args.port}/api/i2c/transaction")
    server.serve_forever()
    return 0


if __name__ == "__main__":
    raise SystemExit(main())
