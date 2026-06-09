#!/bin/sh
set -eu

SCRIPT_DIR="$(CDPATH= cd -- "$(dirname "$0")" && pwd)"
ROOT_DIR="$(CDPATH= cd -- "$SCRIPT_DIR/.." && pwd)"
BIN_PATH="${BADGESNAKE_BIN:-$ROOT_DIR/.cache/badgesnake-ui-sim}"
FORCE_BUILD="${FORCE_BUILD:-0}"

needs_build=0
if [ ! -x "$BIN_PATH" ]; then
    needs_build=1
fi

if [ "$FORCE_BUILD" = "1" ]; then
    needs_build=1
fi

if [ "$needs_build" = "0" ]; then
    if [ "$ROOT_DIR/go.mod" -nt "$BIN_PATH" ]; then
        needs_build=1
    fi
fi

if [ "$needs_build" = "0" ]; then
    if find "$ROOT_DIR/cmd" "$ROOT_DIR/internal" -name '*.go' -newer "$BIN_PATH" -print -quit | grep -q .; then
        needs_build=1
    fi
fi

if [ "$needs_build" = "1" ]; then
    mkdir -p "$(dirname "$BIN_PATH")"
    cd "$ROOT_DIR"
    go build -o "$BIN_PATH" ./cmd/badgesnake
fi

echo "$BIN_PATH"
