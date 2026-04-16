#!/bin/sh
set -eu

ROOT_DIR=$(CDPATH= cd -- "$(dirname -- "$0")/.." && pwd)
SIM_DIR=$(mktemp -d)
trap 'rm -rf "$SIM_DIR"' EXIT

WRITE_ID=$("$ROOT_DIR/scripts/microblocks_i2c_sim.py" --sim-dir "$SIM_DIR" enqueue --address 0x42 --write 010203)
test -f "$SIM_DIR/write-66-$WRITE_ID.bin"

REQ_ID=$("$ROOT_DIR/scripts/microblocks_i2c_sim.py" --sim-dir "$SIM_DIR" enqueue-read --address 0x42)
test -f "$SIM_DIR/read-66-$REQ_ID.bin"
"$ROOT_DIR/scripts/microblocks_i2c_sim.py" --sim-dir "$SIM_DIR" respond --address 0x42 --request-id "$REQ_ID" --write aabb
RESPONSE=$("$ROOT_DIR/scripts/microblocks_i2c_sim.py" --sim-dir "$SIM_DIR" wait-response --address 0x42 --request-id "$REQ_ID" --timeout 1)

test "$RESPONSE" = "aabb"
printf '%s\n' "microblocks i2c sim helper smoke passed"
