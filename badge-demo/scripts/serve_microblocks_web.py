#!/usr/bin/env python3
"""Serve the MicroBlocks web app over HTTPS from the checked-out webapp directory."""

from __future__ import annotations

import argparse
import http.server
import os
import ssl
import subprocess
from pathlib import Path


ROOT_DIR = Path(__file__).resolve().parents[1]
WEBAPP_DIR = ROOT_DIR / "components" / "microblocks-smallvm" / "chromeApp" / "webapp"
CERT_DIR = ROOT_DIR / ".cache" / "microblocks-web"
CERT_FILE = CERT_DIR / "cert.pem"
KEY_FILE = CERT_DIR / "key.pem"


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


def main() -> int:
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument("--host", default="beaglebadge.local")
    parser.add_argument("--port", type=int, default=8443)
    parser.add_argument("--http", action="store_true", help="serve without TLS")
    args = parser.parse_args()

    os.chdir(WEBAPP_DIR)
    handler = http.server.SimpleHTTPRequestHandler
    server = http.server.ThreadingHTTPServer((args.host, args.port), handler)

    scheme = "http"
    if not args.http:
        ensure_certificate()
        context = ssl.SSLContext(ssl.PROTOCOL_TLS_SERVER)
        context.load_cert_chain(certfile=str(CERT_FILE), keyfile=str(KEY_FILE))
        server.socket = context.wrap_socket(server.socket, server_side=True)
        scheme = "https"

    print(f"Serving MicroBlocks web app from {WEBAPP_DIR}")
    print(f"Open {scheme}://{args.host if args.host != '0.0.0.0' else 'localhost'}:{args.port}/microblocks.html")
    server.serve_forever()
    return 0


if __name__ == "__main__":
    raise SystemExit(main())
