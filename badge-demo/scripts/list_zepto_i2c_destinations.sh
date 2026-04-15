#!/usr/bin/env bash
set -euo pipefail

repo_root="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
bb_imager_cli="${repo_root}/components/bb-imager-rs/target/debug/bb-imager-cli"

if [ ! -x "${bb_imager_cli}" ]; then
  echo "Missing ${bb_imager_cli}"
  echo "Build it first with scripts/build_bb_imager_cli_zepto_i2c.sh"
  exit 1
fi

exec "${bb_imager_cli}" list-destinations zepto --no-frills
