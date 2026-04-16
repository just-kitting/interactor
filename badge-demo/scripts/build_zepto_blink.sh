#!/usr/bin/env bash
set -euo pipefail

repo_root="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
exec /root/.platformio/penv/bin/pio run -d "${repo_root}/firmware/zepto" -e zepto_blink
