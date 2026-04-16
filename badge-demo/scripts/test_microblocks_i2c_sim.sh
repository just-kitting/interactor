#!/bin/sh
set -eu

ROOT_DIR=$(CDPATH= cd -- "$(dirname -- "$0")/.." && pwd)
SIM_DIR=$(mktemp -d)
trap 'rm -rf "$SIM_DIR"' EXIT

REQ_ID=$("$ROOT_DIR/scripts/microblocks_i2c_sim.py" --sim-dir "$SIM_DIR" enqueue --address 0x42 --write 010203 --read-length 2)
"$ROOT_DIR/scripts/microblocks_i2c_sim.py" --sim-dir "$SIM_DIR" respond --address 0x42 --request-id "$REQ_ID" --write aabb
RESPONSE=$("$ROOT_DIR/scripts/microblocks_i2c_sim.py" --sim-dir "$SIM_DIR" wait-response --address 0x42 --request-id "$REQ_ID" --timeout 1)

test "$RESPONSE" = "aabb"
printf '%s\n' "microblocks i2c sim helper smoke passed"
