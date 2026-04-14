#!/usr/bin/env bash
set -euo pipefail

repo_root="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
rules_dir="${repo_root}/components/battlesnake-rules"

cd "${rules_dir}"

exec go run ./cli/battlesnake play \
  --width 7 \
  --height 7 \
  --seed 20260414 \
  --name ZeptoA \
  --url 'i2c://stub?addr=0x10&move=clockwise&name=ZeptoA' \
  --name ZeptoB \
  --url 'i2c://stub?addr=0x11&move=counterclockwise&name=ZeptoB'
