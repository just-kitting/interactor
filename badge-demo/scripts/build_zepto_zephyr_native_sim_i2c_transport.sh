#!/usr/bin/env bash
set -euo pipefail

repo_root="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
app_dir="${repo_root}/firmware/zepto/zephyr/native_sim_i2c_transport"
build_dir="${BADGESNAKE_ZEPHYR_BUILD_DIR:-${repo_root}/.cache/zepto-native-sim-i2c-transport}"
board="${BADGESNAKE_ZEPHYR_BOARD:-native_sim}"

if ! command -v west >/dev/null 2>&1; then
  echo "west is required to build the Zephyr native_sim transport tests" >&2
  exit 1
fi

mkdir -p "${build_dir}"

echo "Building ${app_dir} for board ${board}..."
west build -b "${board}" "${app_dir}" -d "${build_dir}" "$@"

echo "${build_dir}/zephyr/zephyr.exe"
