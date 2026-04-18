#!/usr/bin/env python3
"""Send an I2C controller transaction to the hosted Boardie bridge."""

from __future__ import annotations

import argparse
import json
import ssl
import sys
import urllib.error
import urllib.request


def parse_byte(value: str) -> int:
    return int(value, 0) & 0xFF


def main() -> int:
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument("--url", default="https://127.0.0.1:8443")
    parser.add_argument("--address", type=parse_byte, required=True)
    parser.add_argument("--write", nargs="*", default=[])
    parser.add_argument("--timeout-ms", type=int, default=5000)
    parser.add_argument("--insecure", action="store_true", help="skip TLS certificate verification")
    args = parser.parse_args()

    payload = json.dumps(
        {
            "address": args.address,
            "write": [parse_byte(value) for value in args.write],
            "timeout_ms": args.timeout_ms,
        }
    ).encode("utf-8")
    request = urllib.request.Request(
        args.url.rstrip("/") + "/api/i2c/transaction",
        data=payload,
        headers={"Content-Type": "application/json", "Accept": "application/json"},
        method="POST",
    )

    context = None
    if args.insecure:
        context = ssl._create_unverified_context()

    try:
        with urllib.request.urlopen(request, context=context) as response:
            body = json.loads(response.read().decode("utf-8"))
    except urllib.error.HTTPError as exc:
        body = json.loads(exc.read().decode("utf-8"))
        print(json.dumps(body, indent=2, sort_keys=True))
        return exc.code

    print(json.dumps(body, indent=2, sort_keys=True))
    return 0


if __name__ == "__main__":
    raise SystemExit(main())
