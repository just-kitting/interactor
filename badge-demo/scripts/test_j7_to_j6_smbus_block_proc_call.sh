#!/usr/bin/env bash

set -euo pipefail

SCRIPT_DIR="$(cd -- "$(dirname -- "${BASH_SOURCE[0]}")" && pwd)"
REPO_ROOT="$(cd -- "${SCRIPT_DIR}/.." && pwd)"
BRINGUP="${REPO_ROOT}/scripts/bringup_i2c_slave_testunit.sh"
SRC="${REPO_ROOT}/scripts/test_smbus_block_proc_call.c"
CACHE_DIR="${REPO_ROOT}/.cache"
BIN="${CACHE_DIR}/test_smbus_block_proc_call"

SLAVE_BUS="${BADGESNAKE_SLAVE_BUS:-1}"
MASTER_BUS="${BADGESNAKE_MASTER_BUS:-3}"
ADDR="${BADGESNAKE_TARGET_ADDR:-0x30}"
CMD="${BADGESNAKE_PROC_CMD:-0x03}"
COUNT="${BADGESNAKE_PROC_COUNT:-1}"
VALUE="${BADGESNAKE_PROC_VALUE:-4}"

command -v gcc >/dev/null 2>&1 || {
	echo "missing required command: gcc" >&2
	exit 1
}

[[ -x "${BRINGUP}" ]] || {
	echo "missing helper: ${BRINGUP}" >&2
	exit 1
}

mkdir -p "${CACHE_DIR}"

gcc -O2 -Wall -Wextra -o "${BIN}" "${SRC}"

if "${BRINGUP}" status "${SLAVE_BUS}" "${ADDR}" >/dev/null 2>&1; then
	"${BRINGUP}" stop "${SLAVE_BUS}" "${ADDR}" >/dev/null
	sleep 1
fi
"${BRINGUP}" start "${SLAVE_BUS}" "${ADDR}" >/dev/null
sleep 1
"${BRINGUP}" status "${SLAVE_BUS}" "${ADDR}" >/dev/null

echo "SMBus block-proc-call via /dev/i2c-${MASTER_BUS} -> ${ADDR}"
dmesg -C
"${BIN}" "${MASTER_BUS}" "${ADDR}" "${CMD}" "${COUNT}" "${VALUE}"

echo
echo "Recent slave diagnostics:"
dmesg | tail -n 80
