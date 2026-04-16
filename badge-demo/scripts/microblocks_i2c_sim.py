#!/usr/bin/env python3
"""Interact with the Linux MicroBlocks I2C target simulator spool directory."""

from __future__ import annotations

import argparse
import os
import sys
import tempfile
import time
from pathlib import Path

DEFAULT_SIM_DIR = Path(os.environ.get("MICROBLOCKS_I2C_SIM_DIR", "/tmp/microblocks_i2c_target_sim"))


def parse_address(value: str) -> int:
    address = int(value, 0)
    if not 0 <= address <= 0x7F:
        raise argparse.ArgumentTypeError("I2C address must be between 0x00 and 0x7f")
    return address


def parse_hex_bytes(text: str) -> bytes:
    cleaned = text.replace(" ", "").replace(",", "").replace(":", "")
    if cleaned.startswith("0x"):
        cleaned = cleaned[2:]
    if len(cleaned) % 2:
        raise argparse.ArgumentTypeError("hex payload must contain an even number of nybbles")
    try:
        return bytes.fromhex(cleaned)
    except ValueError as exc:
        raise argparse.ArgumentTypeError(str(exc)) from exc


def ensure_dir(path: Path) -> None:
    path.mkdir(parents=True, exist_ok=True)


def next_request_id(sim_dir: Path, address: int) -> str:
    prefix = f"request-{address:02d}-"
    last = 0
    for entry in sim_dir.glob(f"{prefix}*.bin"):
        stem = entry.stem
        try:
            request_id = int(stem.split("-")[-1], 10)
        except ValueError:
            continue
        last = max(last, request_id)
    return f"{last + 1:020d}"


def request_path(sim_dir: Path, address: int, request_id: str) -> Path:
    return sim_dir / f"request-{address:02d}-{request_id}.bin"


def response_path(sim_dir: Path, address: int, request_id: str) -> Path:
    return sim_dir / f"response-{address:02d}-{request_id}.bin"


def write_atomic(path: Path, payload: bytes) -> None:
    ensure_dir(path.parent)
    with tempfile.NamedTemporaryFile(dir=path.parent, delete=False) as tmp:
        tmp.write(payload)
        tmp.flush()
        os.fsync(tmp.fileno())
        tmp_path = Path(tmp.name)
    tmp_path.replace(path)


def enqueue(sim_dir: Path, address: int, write_data: bytes, read_length: int) -> str:
    if read_length < 0:
        raise ValueError("read length must be non-negative")
    request_id = next_request_id(sim_dir, address)
    header = read_length.to_bytes(4, "little", signed=False)
    write_atomic(request_path(sim_dir, address, request_id), header + write_data)
    return request_id


def wait_for_response(sim_dir: Path, address: int, request_id: str, timeout: float, poll: float) -> bytes:
    rsp_path = response_path(sim_dir, address, request_id)
    deadline = time.monotonic() + timeout
    while time.monotonic() < deadline:
        if rsp_path.exists():
            payload = rsp_path.read_bytes()
            rsp_path.unlink()
            return payload
        time.sleep(poll)
    raise TimeoutError(f"timed out waiting for response {request_id} on address 0x{address:02x}")


def format_bytes(payload: bytes, mode: str) -> str:
    if mode == "hex":
        return payload.hex()
    if mode == "text":
        return payload.decode("utf-8", errors="replace")
    raise ValueError(f"unsupported format {mode}")


def cmd_enqueue(args: argparse.Namespace) -> int:
    request_id = enqueue(args.sim_dir, args.address, args.write_data, args.read_length)
    print(request_id)
    return 0


def cmd_wait_response(args: argparse.Namespace) -> int:
    payload = wait_for_response(args.sim_dir, args.address, args.request_id, args.timeout, args.poll_interval)
    print(format_bytes(payload, args.format))
    return 0


def cmd_transaction(args: argparse.Namespace) -> int:
    request_id = enqueue(args.sim_dir, args.address, args.write_data, args.read_length)
    print(
        f"queued request {request_id} for address 0x{args.address:02x}; "
        f"waiting for MicroBlocks reply in {args.sim_dir}",
        file=sys.stderr,
    )
    payload = wait_for_response(args.sim_dir, args.address, request_id, args.timeout, args.poll_interval)
    print(format_bytes(payload, args.format))
    return 0


def cmd_respond(args: argparse.Namespace) -> int:
    write_atomic(response_path(args.sim_dir, args.address, args.request_id), args.write_data)
    return 0


def build_parser() -> argparse.ArgumentParser:
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument("--sim-dir", type=Path, default=DEFAULT_SIM_DIR, help="spool directory path")
    subparsers = parser.add_subparsers(dest="command", required=True)

    enqueue_parser = subparsers.add_parser("enqueue", help="enqueue a controller request")
    enqueue_parser.add_argument("--address", type=parse_address, required=True)
    enqueue_parser.add_argument("--write", dest="write_data", type=parse_hex_bytes, default=b"", help="controller write payload as hex")
    enqueue_parser.add_argument("--read-length", type=int, default=0, help="controller read length after the write phase")
    enqueue_parser.set_defaults(func=cmd_enqueue)

    wait_parser = subparsers.add_parser("wait-response", help="wait for a controller response")
    wait_parser.add_argument("--address", type=parse_address, required=True)
    wait_parser.add_argument("--request-id", required=True)
    wait_parser.add_argument("--timeout", type=float, default=10.0)
    wait_parser.add_argument("--poll-interval", type=float, default=0.1)
    wait_parser.add_argument("--format", choices=("hex", "text"), default="hex")
    wait_parser.set_defaults(func=cmd_wait_response)

    tx_parser = subparsers.add_parser("transaction", help="enqueue a request and wait for a response")
    tx_parser.add_argument("--address", type=parse_address, required=True)
    tx_parser.add_argument("--write", dest="write_data", type=parse_hex_bytes, default=b"", help="controller write payload as hex")
    tx_parser.add_argument("--read-length", type=int, default=0, help="controller read length after the write phase")
    tx_parser.add_argument("--timeout", type=float, default=10.0)
    tx_parser.add_argument("--poll-interval", type=float, default=0.1)
    tx_parser.add_argument("--format", choices=("hex", "text"), default="hex")
    tx_parser.set_defaults(func=cmd_transaction)

    respond_parser = subparsers.add_parser("respond", help="write a response file directly")
    respond_parser.add_argument("--address", type=parse_address, required=True)
    respond_parser.add_argument("--request-id", required=True)
    respond_parser.add_argument("--write", dest="write_data", type=parse_hex_bytes, required=True, help="response payload as hex")
    respond_parser.set_defaults(func=cmd_respond)

    return parser


def main() -> int:
    parser = build_parser()
    args = parser.parse_args()
    try:
        return args.func(args)
    except TimeoutError as exc:
        print(str(exc), file=sys.stderr)
        return 1
    except ValueError as exc:
        print(str(exc), file=sys.stderr)
        return 1


if __name__ == "__main__":
    raise SystemExit(main())
