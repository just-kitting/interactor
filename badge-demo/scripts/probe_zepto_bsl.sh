#!/usr/bin/env bash
set -euo pipefail

target_addr="${1:-0x48}"

printf 'Probing Zepto MSPM0 BSL target address %s\n' "${target_addr}"
echo

for busdev in /dev/i2c-*; do
  [ -e "${busdev}" ] || continue
  bus="${busdev##*-}"
  echo "Bus ${bus} (${busdev})"
  i2cdetect -r -y "${bus}" | sed 's/^/  /'
  echo
done

echo "Expected by bb-imager-rs: ${target_addr}"
echo "If the Zepto is in BSL over I2C, an ACK at ${target_addr} should appear on the connected bus."
