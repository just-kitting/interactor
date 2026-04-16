#!/usr/bin/env python3
import argparse
import binascii
import subprocess
import sys


def jamcrc(data: bytes) -> int:
    return binascii.crc32(data, 0xFFFFFFFF) ^ 0xFFFFFFFF


def packet(cmd: int, payload: bytes = b"") -> bytes:
    length = len(payload) + 1
    body = bytes([cmd]) + payload
    crc = jamcrc(body).to_bytes(4, "little")
    return bytes([0x80, length & 0xFF, (length >> 8) & 0xFF, cmd]) + payload + crc


def i2ctransfer(bus: int, addr: int, write_bytes: bytes, read_len: int) -> bytes:
    cmd = [
        "i2ctransfer",
        "-f",
        "-y",
        str(bus),
        f"w{len(write_bytes)}@0x{addr:02x}",
    ]
    cmd.extend(f"0x{b:02x}" for b in write_bytes)
    cmd.append(f"r{read_len}")

    proc = subprocess.run(cmd, capture_output=True, text=True)
    if proc.returncode != 0:
        raise RuntimeError(proc.stderr.strip() or proc.stdout.strip() or "i2ctransfer failed")

    text = proc.stdout.strip()
    if not text:
        return b""

    return bytes(int(part, 16) for part in text.split())


def print_hex(label: str, data: bytes) -> None:
    print(f"{label}: {data.hex()}")


def main() -> int:
    parser = argparse.ArgumentParser(description="Trace MSPM0 Zepto BSL over I2C with combined transfers")
    parser.add_argument("--bus", type=int, default=1)
    parser.add_argument("--addr", type=lambda x: int(x, 0), default=0x48)
    args = parser.parse_args()

    connect = packet(0x12)
    get_info = packet(0x19)

    print_hex("connect-req", connect)
    ack = i2ctransfer(args.bus, args.addr, connect, 1)
    print_hex("connect-ack", ack)

    print_hex("get-info-req", get_info)
    ack = i2ctransfer(args.bus, args.addr, get_info, 1)
    print_hex("get-info-ack", ack)

    response = i2ctransfer(args.bus, args.addr, b"", 32)
    print_hex("get-info-resp", response)
    return 0


if __name__ == "__main__":
    try:
        raise SystemExit(main())
    except Exception as exc:
        print(f"error: {exc}", file=sys.stderr)
        raise SystemExit(1)
