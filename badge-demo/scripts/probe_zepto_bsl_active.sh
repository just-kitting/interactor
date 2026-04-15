#!/usr/bin/env bash
set -euo pipefail

bus="${1:-0}"
addr="${2:-0x48}"
attempts="${3:-50}"
sleep_s="${4:-0.2}"

# MSPM0 BSL connection request from bb-imager-rs test vectors.
request=(
  0x80 0x01 0x00 0x12 0x3A 0x61 0x44 0xDE
)

echo "Active probing J6/J7-style MSPM0 BSL on /dev/i2c-${bus} at ${addr}"
echo "Attempts: ${attempts}, interval: ${sleep_s}s"
echo

for attempt in $(seq 1 "${attempts}"); do
  printf '[%02d/%02d] ' "${attempt}" "${attempts}"

  if out="$(i2ctransfer -f -y "${bus}" "w8@${addr}" "${request[@]}" "r1" 2>&1)"; then
    echo "ACK received: ${out}"
    exit 0
  fi

  echo "no response"
  sleep "${sleep_s}"
done

echo
echo "No ACK from ${addr} on /dev/i2c-${bus}."
echo "If the Zepto BSL timed out, re-enter BSL and rerun this script immediately."
exit 1
