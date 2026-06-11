#!/usr/bin/env bash
set -euo pipefail

repo_root="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
"${repo_root}/scripts/build_zepto_zephyr_native_sim_i2c_transport.sh" "$@"
"${repo_root}/scripts/run_zepto_zephyr_native_sim_i2c_transport.sh"
