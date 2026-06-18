#!/usr/bin/env python3
import argparse
import fcntl
import os
import struct
import subprocess
import sys
import time
import zlib
from pathlib import Path


I2C_SLAVE = 0x0703
REQ = 0x80
RESP = 0x08
ACK = 0x00

CMD_CONNECTION = 0x12
CMD_GET_DEVICE_INFO = 0x19
CMD_UNLOCK = 0x21
CMD_MASS_ERASE = 0x15
CMD_PROGRAM_DATA = 0x20
CMD_START_APPLICATION = 0x40
CMD_CORE_MESSAGE = 0x3B
CMD_GET_DEVICE_INFO_RESP = 0x31

DEFAULT_PAD_GPIO = 0x00054007


def jamcrc(data: bytes) -> int:
    return zlib.crc32(data) ^ 0xFFFFFFFF


def pkt_no_data(cmd: int) -> bytes:
    return bytes([REQ, 0x01, 0x00, cmd]) + struct.pack("<I", jamcrc(bytes([cmd])))


def pkt_unlock() -> bytes:
    pwd = bytes([0xFF]) * 32
    body = bytes([CMD_UNLOCK]) + pwd
    return bytes([REQ, 0x21, 0x00, CMD_UNLOCK]) + pwd + struct.pack("<I", jamcrc(body))


def pkt_program(address: int, data: bytes) -> bytes:
    length = 1 + 4 + len(data)
    body = bytes([CMD_PROGRAM_DATA]) + struct.pack("<I", address) + data
    head = bytes([REQ]) + struct.pack("<H", length) + bytes([CMD_PROGRAM_DATA]) + struct.pack("<I", address)
    return head + data + struct.pack("<I", jamcrc(body))


def read_exact(fd: int, count: int, timeout_s: float) -> bytes:
    deadline = time.time() + timeout_s
    buf = b""
    while len(buf) < count:
        chunk = os.read(fd, count - len(buf))
        if chunk:
            buf += chunk
            continue
        if time.time() > deadline:
            raise TimeoutError(f"timeout reading {count} bytes, got {len(buf)}")
        time.sleep(0.01)
    return buf


def write_then_read(fd: int, tx: bytes, rx_len: int, delay_s: float = 0.02, timeout_s: float = 3.0) -> bytes:
    os.write(fd, tx)
    ack = read_exact(fd, 1, timeout_s)
    if ack[0] != ACK:
        raise RuntimeError(f"unexpected ACK byte 0x{ack[0]:02x}")
    if rx_len == 0:
        return b""
    time.sleep(delay_s)
    return read_exact(fd, rx_len, timeout_s)


def devmem_read(addr: int) -> int:
    out = subprocess.check_output(["busybox", "devmem", hex(addr), "32"], text=True).strip()
    return int(out, 16)


def devmem_write(addr: int, value: int) -> None:
    subprocess.check_call(
        ["busybox", "devmem", hex(addr), "32", hex(value)],
        stdout=subprocess.DEVNULL,
        stderr=subprocess.DEVNULL,
    )


def parse_device_info(payload: bytes) -> dict[str, int]:
    if len(payload) != 32:
        raise RuntimeError(f"unexpected device-info length {len(payload)}")
    header, length, cmd = payload[0], int.from_bytes(payload[1:3], "little"), payload[3]
    if header != RESP or cmd != CMD_GET_DEVICE_INFO_RESP:
        raise RuntimeError(f"unexpected device-info header/cmd: {payload[:4].hex()}")
    return {
        "cmd_interpreter_version": int.from_bytes(payload[4:6], "little"),
        "build_id": int.from_bytes(payload[6:8], "little"),
        "application_version": int.from_bytes(payload[8:12], "little"),
        "active_plugin_version": int.from_bytes(payload[12:14], "little"),
        "bsl_max_buffer_size": int.from_bytes(payload[14:16], "little"),
        "bsl_buffer_start_address": int.from_bytes(payload[16:20], "little"),
        "bcr_configuration_id": int.from_bytes(payload[20:24], "little"),
        "bsl_configuration_id": int.from_bytes(payload[24:28], "little"),
        "packet_len": length,
    }


def parse_core_response(payload: bytes) -> None:
    if len(payload) != 9:
        raise RuntimeError(f"unexpected core-response length {len(payload)}")
    header, length, cmd, msg = payload[0], int.from_bytes(payload[1:3], "little"), payload[3], payload[4]
    if header != RESP or length != 2 or cmd != CMD_CORE_MESSAGE or msg != 0x00:
        raise RuntimeError(f"unexpected core response: {payload.hex()}")


def chunk_image(image: bytes, max_data_len: int):
    chunk_len = max_data_len - (max_data_len % 8)
    if chunk_len <= 0:
        raise RuntimeError(f"invalid max_data_len {max_data_len}")
    addr = 0
    while addr < len(image):
        chunk = image[addr : addr + chunk_len]
        rem = len(chunk) % 8
        if rem:
            chunk += bytes([0xFF]) * (8 - rem)
        yield addr, chunk
        addr += chunk_len


def main() -> int:
    parser = argparse.ArgumentParser()
    parser.add_argument("image", type=Path)
    parser.add_argument("--bus", default="/dev/i2c-1")
    parser.add_argument("--address", default="0x48")
    parser.add_argument("--bsl-line", type=int, default=27)
    parser.add_argument("--rst-line", type=int, default=28)
    parser.add_argument("--gpiochip", default="gpiochip1")
    parser.add_argument("--pad-bsl", type=lambda s: int(s, 0), default=0x040840A8)
    parser.add_argument("--pad-rst", type=lambda s: int(s, 0), default=0x040840AC)
    parser.add_argument("--pad-value", type=lambda s: int(s, 0), default=DEFAULT_PAD_GPIO)
    args = parser.parse_args()

    image = args.image.read_bytes()
    addr = int(args.address, 0)

    orig_bsl = devmem_read(args.pad_bsl)
    orig_rst = devmem_read(args.pad_rst)
    hold = None
    try:
        devmem_write(args.pad_bsl, args.pad_value)
        devmem_write(args.pad_rst, args.pad_value)

        pulse = subprocess.Popen(
            ["gpioset", "-c", args.gpiochip, f"{args.bsl_line}=1", f"{args.rst_line}=0"]
        )
        time.sleep(0.25)
        pulse.terminate()
        pulse.wait(timeout=1)

        hold = subprocess.Popen(
            ["gpioset", "-c", args.gpiochip, f"{args.bsl_line}=1", f"{args.rst_line}=1"]
        )
        time.sleep(0.12)

        fd = os.open(args.bus, os.O_RDWR)
        fcntl.ioctl(fd, I2C_SLAVE, addr)
        try:
            write_then_read(fd, pkt_no_data(CMD_CONNECTION), 0)
            print("connect: ok")

            info = parse_device_info(write_then_read(fd, pkt_no_data(CMD_GET_DEVICE_INFO), 32, delay_s=0.05))
            print(
                "device_info:"
                f" build=0x{info['build_id']:04x}"
                f" plugin=0x{info['active_plugin_version']:04x}"
                f" max_buffer={info['bsl_max_buffer_size']}"
            )

            parse_core_response(write_then_read(fd, pkt_unlock(), 9, delay_s=0.05))
            print("unlock: ok")

            parse_core_response(write_then_read(fd, pkt_no_data(CMD_MASS_ERASE), 9, delay_s=0.2))
            print("mass_erase: ok")

            max_data_len = info["bsl_max_buffer_size"] - (8 + 4)
            total_chunks = sum(1 for _ in chunk_image(image, max_data_len))
            for i, (offset, chunk) in enumerate(chunk_image(image, max_data_len), start=1):
                parse_core_response(write_then_read(fd, pkt_program(offset, chunk), 9, delay_s=0.05))
                print(f"program_data: chunk {i}/{total_chunks} addr=0x{offset:08x} len={len(chunk)}")

            write_then_read(fd, pkt_no_data(CMD_START_APPLICATION), 0)
            print("start_application: ok")
        finally:
            os.close(fd)
    finally:
        if hold is not None:
            hold.terminate()
            try:
                hold.wait(timeout=1)
            except subprocess.TimeoutExpired:
                hold.kill()
        devmem_write(args.pad_bsl, orig_bsl)
        devmem_write(args.pad_rst, orig_rst)

    return 0


if __name__ == "__main__":
    sys.exit(main())
