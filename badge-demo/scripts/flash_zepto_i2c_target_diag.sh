#!/usr/bin/env bash
set -euo pipefail

repo_root="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
env_name="${1:-zepto_i2c_target_diag_stage6}"

exec /root/.platformio/penv/bin/pio run -d "${repo_root}/firmware/zepto" -e "${env_name}" -t upload
