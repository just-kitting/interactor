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

cd "${rules_dir}"

exec go run ./cli/battlesnake play \
  --width 3 \
  --height 3 \
  --seed 20260414 \
  --foodSpawnChance 0 \
  --minimumFood 0 \
  --name ZeptoA \
  --url 'i2c://stub?addr=0x10&move=clockwise&name=ZeptoA' \
  --name ZeptoB \
  --url 'i2c://stub?addr=0x11&move=counterclockwise&name=ZeptoB'
