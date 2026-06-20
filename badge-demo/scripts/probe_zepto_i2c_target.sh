#!/usr/bin/env bash
set -euo pipefail

bus="${1:-1}"
addr="${2:-0x42}"

echo "probe: bus=/dev/i2c-${bus} addr=${addr}"
echo
echo "[i2cdetect]"
timeout 5 i2cdetect -y "${bus}" || true
echo
echo "[i2cdetect -r]"
timeout 5 i2cdetect -y -r "${bus}" || true
echo
echo "[write-1]"
timeout 5 i2ctransfer -f -y "${bus}" "w1@${addr}" 0x00 || true
echo
echo "[write-6]"
timeout 5 i2ctransfer -f -y "${bus}" "w6@${addr}" 0x01 0x01 0x00 0x00 0x00 0x00 || true
