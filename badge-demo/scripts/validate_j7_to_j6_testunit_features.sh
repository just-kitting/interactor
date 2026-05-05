#!/usr/bin/env bash

set -euo pipefail

SCRIPT_DIR="$(cd -- "$(dirname -- "${BASH_SOURCE[0]}")" && pwd)"
REPO_ROOT="$(cd -- "${SCRIPT_DIR}/.." && pwd)"
BRINGUP="${REPO_ROOT}/scripts/bringup_i2c_slave_testunit.sh"

SLAVE_BUS="${BADGESNAKE_SLAVE_BUS:-1}"
MASTER_BUS="${BADGESNAKE_MASTER_BUS:-3}"
ADDR="${BADGESNAKE_TARGET_ADDR:-0x30}"
VERSION_READ_LEN="${BADGESNAKE_VERSION_READ_LEN:-32}"
PROC_LEN="${BADGESNAKE_PROC_LEN:-4}"

usage() {
	cat <<'EOF'
Usage:
  validate_j7_to_j6_testunit_features.sh

Environment:
  BADGESNAKE_SLAVE_BUS        Target/slave bus (default: 1)
  BADGESNAKE_MASTER_BUS       Initiator/controller bus (default: 3)
  BADGESNAKE_TARGET_ADDR      7-bit slave address (default: 0x30)
  BADGESNAKE_VERSION_READ_LEN Bytes to read for repeated-start version test (default: 32)
  BADGESNAKE_PROC_LEN         DATAH for SMBus block-proc-call test (default: 4)

Expected wiring:
  - J6 shorted to J7 so /dev/i2c-${MASTER_BUS} can act as a second controller
    against the slave backend instantiated on /dev/i2c-${SLAVE_BUS}

Tests:
  1. GET_VERSION_WITH_REP_START-like sequence over J7->J6
  2. SMBUS_BLOCK_PROC_CALL-like sequence over J7->J6
EOF
}

if [[ "${1:-}" == "-h" || "${1:-}" == "--help" ]]; then
	usage
	exit 0
fi

for cmd in i2ctransfer dmesg tr; do
	command -v "${cmd}" >/dev/null 2>&1 || {
		echo "missing required command: ${cmd}" >&2
		exit 1
	}
done

[[ -x "${BRINGUP}" ]] || {
	echo "missing helper: ${BRINGUP}" >&2
	exit 1
}

if "${BRINGUP}" status "${SLAVE_BUS}" "${ADDR}" >/dev/null 2>&1; then
	"${BRINGUP}" stop "${SLAVE_BUS}" "${ADDR}" >/dev/null
	sleep 1
fi
"${BRINGUP}" start "${SLAVE_BUS}" "${ADDR}" >/dev/null
sleep 1
"${BRINGUP}" status "${SLAVE_BUS}" "${ADDR}" >/dev/null

echo "Repeated-start version test via /dev/i2c-${MASTER_BUS} -> ${ADDR}"
dmesg -C
version_tmp="$(mktemp)"
trap 'rm -f "${version_tmp}"' EXIT
i2ctransfer -f -y -b "${MASTER_BUS}" "w3@${ADDR}" 4 0 0 "r${VERSION_READ_LEN}" > "${version_tmp}"
version_output="$(tr -d '\000' < "${version_tmp}")"
version_dmesg="$(dmesg)"

if [[ "${version_output}" != v* ]]; then
	echo "unexpected version response: ${version_output}" >&2
	exit 1
fi

echo "Version response: ${version_output}"

echo
echo "SMBus block-proc-call style test via /dev/i2c-${MASTER_BUS} -> ${ADDR}"
dmesg -C
proc_read_len=$((PROC_LEN + 1))
proc_output="$(i2ctransfer -f -y "${MASTER_BUS}" "w3@${ADDR}" 3 1 "${PROC_LEN}" "r${proc_read_len}@${ADDR}")"
proc_dmesg="$(dmesg)"

expected_proc=""
for ((i=PROC_LEN; i>=0; i--)); do
	part="$(printf '0x%02x' "${i}")"
	if [[ -z "${expected_proc}" ]]; then
		expected_proc="${part}"
	else
		expected_proc="${expected_proc} ${part}"
	fi
done

if [[ "${proc_output}" != "${expected_proc}" ]]; then
	echo "unexpected block-proc-call response: ${proc_output}" >&2
	echo "expected: ${expected_proc}" >&2
	exit 1
fi

echo "Block-proc-call response: ${proc_output}"

echo
echo "Repeated-start slave diagnostics:"
if [[ -n "${version_dmesg}" ]]; then
	printf '%s\n' "${version_dmesg}"
else
	echo "(none)"
fi

echo
echo "Block-proc-call slave diagnostics:"
if [[ -n "${proc_dmesg}" ]]; then
	printf '%s\n' "${proc_dmesg}"
else
	echo "(none)"
fi
