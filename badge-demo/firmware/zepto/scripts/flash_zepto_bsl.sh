#!/usr/bin/env bash
set -euo pipefail

if [ "$#" -ne 1 ]; then
  echo "usage: $0 <image.bin>" >&2
  exit 2
fi

image="$1"
repo_root="$(cd "$(dirname "${BASH_SOURCE[0]}")/../../.." && pwd)"
bb_imager_cli="${repo_root}/components/bb-imager-rs/target/debug/bb-imager-cli"
i2c_bus="${ZEPTO_I2C_BUS:-/dev/i2c-1}"
probe_addr="${ZEPTO_BSL_ADDR:-0x48}"
probe_attempts="${ZEPTO_BSL_WAIT_ATTEMPTS:-50}"
probe_sleep_s="${ZEPTO_BSL_WAIT_INTERVAL:-0.2}"
flash_attempts="${ZEPTO_FLASH_ATTEMPTS:-3}"

bus_num="${i2c_bus##*/i2c-}"

if [ "${bus_num}" = "${i2c_bus}" ]; then
  echo "expected ZEPTO_I2C_BUS like /dev/i2c-1, got ${i2c_bus}" >&2
  exit 1
fi

if [ ! -f "${image}" ]; then
  echo "missing image: ${image}" >&2
  exit 1
fi

if [ ! -x "${bb_imager_cli}" ]; then
  echo "missing ${bb_imager_cli}" >&2
  echo "build it first with scripts/build_bb_imager_cli_zepto_i2c.sh" >&2
  exit 1
fi

status=1

for attempt in $(seq 1 "${flash_attempts}"); do
  echo
  echo "Flash attempt ${attempt}/${flash_attempts}"
  echo "Waiting for Zepto MSPM0 BSL on ${i2c_bus}."
  echo "Perform the Zepto BOOT/RST sequence now if the board is not already in BSL."
  "${repo_root}/scripts/probe_zepto_bsl_active.sh" "${bus_num}" "${probe_addr}" "${probe_attempts}" "${probe_sleep_s}"

  echo "Flashing ${image} to Zepto via ${i2c_bus} (attempt ${attempt}/${flash_attempts})"
  set +e
  "${bb_imager_cli}" flash zepto "${image}" "${i2c_bus}"
  status=$?
  set -e
  if [ "${status}" -eq 0 ]; then
    exit 0
  fi

  echo "Flash attempt ${attempt}/${flash_attempts} failed with exit code ${status}" >&2
  if [ "${attempt}" -lt "${flash_attempts}" ]; then
    echo "Retrying after a short delay. Re-enter BSL now if needed." >&2
    sleep 0.5
  fi
done

exit "${status}"
