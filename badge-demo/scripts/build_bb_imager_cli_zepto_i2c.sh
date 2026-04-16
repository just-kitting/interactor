#!/usr/bin/env bash
set -euo pipefail

repo_root="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
cd "${repo_root}/components/bb-imager-rs"

exec cargo build -p bb-imager-cli --features zepto_i2c
