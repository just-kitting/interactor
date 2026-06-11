#!/usr/bin/env bash
set -euo pipefail

repo_root="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
build_dir="${BADGESNAKE_ZEPHYR_BUILD_DIR:-${repo_root}/.cache/zepto-native-sim-i2c-transport}"
exe_path="${BADGESNAKE_ZEPHYR_EXE:-${build_dir}/zephyr/zephyr.exe}"

if [ ! -x "${exe_path}" ]; then
  echo "expected test executable not found at ${exe_path}" >&2
  exit 1
fi

echo "Running ${exe_path}..."
"${exe_path}" "$@"
