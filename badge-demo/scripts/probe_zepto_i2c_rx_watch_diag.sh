#!/usr/bin/env bash
set -euo pipefail

bus="${1:-1}"
addr="${2:-0x42}"

echo "probe: rx watch on /dev/i2c-${bus} addr=${addr}"
echo
echo "[prime]"
timeout 5 i2ctransfer -f -y "${bus}" "w1@${addr}" 0x00 || true
echo

if timeout 5 i2cdetect -y -r "${bus}" >/tmp/zepto-rx-watch-i2cdetect.txt 2>&1; then
    cat /tmp/zepto-rx-watch-i2cdetect.txt
    echo
    echo "result: target did not latch on received write data"
    exit 1
fi

cat /tmp/zepto-rx-watch-i2cdetect.txt
echo
echo "result: host probe timed out after write to ${addr}; target latched on received write data"
