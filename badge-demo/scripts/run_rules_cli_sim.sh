#!/usr/bin/env bash
set -euo pipefail

repo_root="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
rules_dir="${repo_root}/components/battlesnake-rules"

cd "${rules_dir}"

exec go run ./cli/battlesnake play \
  --width 7 \
  --height 7 \
  --seed 20260414 \
  --name Clocky \
  --url 'sim://clockwise?name=Clocky&color=%23112233' \
  --name Righty \
  --url 'sim://right?name=Righty&color=%2300ff00'
