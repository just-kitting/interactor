#!/usr/bin/env bash
set -euo pipefail

target_addr="${1:-0x48}"
bus="${2:-}"

printf 'Probing Zepto MSPM0 BSL target address %s\n' "${target_addr}"
echo

probe_bus() {
  local busdev="$1"
  local busnum="${busdev##*-}"
  echo "Bus ${busnum} (${busdev})"
  i2cdetect -r -y "${busnum}" | sed 's/^/  /'
  echo
}

if [ -n "${bus}" ]; then
  probe_bus "/dev/i2c-${bus}"
else
  for busdev in /dev/i2c-*; do
    [ -e "${busdev}" ] || continue
    probe_bus "${busdev}"
  done
fi

echo "Expected by bb-imager-rs: ${target_addr}"
echo "If the Zepto is in BSL over I2C, an ACK at ${target_addr} should appear on the connected bus."
echo "J6 is expected on /dev/i2c-0."
