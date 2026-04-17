#!/bin/sh
set -eu

ROOT_DIR=$(CDPATH= cd -- "$(dirname -- "$0")/.." && pwd)
MB_DIR="$ROOT_DIR/components/microblocks-smallvm"
EM_CACHE_DIR="$ROOT_DIR/.cache/emscripten"

if ! command -v emcc >/dev/null 2>&1; then
  echo "emcc not found; install emscripten first" >&2
  exit 1
fi

mkdir -p "$EM_CACHE_DIR"
export EM_FROZEN_CACHE=0
export EM_CACHE="$EM_CACHE_DIR"

echo "Building MicroBlocks web IDE..."
(cd "$MB_DIR/chromeApp/emscripten" && sh ./buildEmcc.sh)

echo "Building Boardie web simulator..."
(cd "$MB_DIR/boardie" && sh ./buildBoardie.sh)

echo "Built web assets in $MB_DIR/chromeApp/webapp"
