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

if [ ! -f "${image}" ]; then
  echo "missing image: ${image}" >&2
  exit 1
fi

if [ ! -x "${bb_imager_cli}" ]; then
  echo "missing ${bb_imager_cli}" >&2
  echo "build it first with scripts/build_bb_imager_cli_zepto_i2c.sh" >&2
  exit 1
fi

echo "Flashing ${image} to Zepto via ${i2c_bus}"
exec "${bb_imager_cli}" flash zepto "${image}" "${i2c_bus}"
