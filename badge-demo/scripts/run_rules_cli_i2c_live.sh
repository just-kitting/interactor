#!/usr/bin/env bash
set -euo pipefail

repo_root="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
rules_dir="${repo_root}/components/battlesnake-rules"
cache_dir="${repo_root}/.cache"

mkdir -p "${cache_dir}/tmp" "${cache_dir}/go-tmp" "${cache_dir}/go-build"

export TMPDIR="${cache_dir}/tmp"
export GOTMPDIR="${cache_dir}/go-tmp"
export GOCACHE="${cache_dir}/go-build"
export CGO_ENABLED=0

bus="${BADGESNAKE_I2C_BUS:-1}"
addr="${BADGESNAKE_I2C_ADDR:-0x42}"
name="${BADGESNAKE_NAME:-Zepto}"
max_response_len="${BADGESNAKE_MAX_RESPONSE_LEN:-512}"
opponent_name="${BADGESNAKE_OPPONENT_NAME:-Clocky}"
opponent_url="${BADGESNAKE_OPPONENT_URL:-sim://clockwise?name=Clocky&color=%23a85f20}"

cd "${rules_dir}"

exec go run ./cli/battlesnake play \
  --width "${BADGESNAKE_WIDTH:-11}" \
  --height "${BADGESNAKE_HEIGHT:-11}" \
  --seed "${BADGESNAKE_SEED:-20260611}" \
  --timeout "${BADGESNAKE_TIMEOUT_MS:-500}" \
  --name "${name}" \
  --url "i2c://${bus}?addr=${addr}&max_response_len=${max_response_len}" \
  --name "${opponent_name}" \
  --url "${opponent_url}" \
  "$@"
