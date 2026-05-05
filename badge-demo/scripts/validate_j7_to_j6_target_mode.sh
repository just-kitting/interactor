#!/usr/bin/env bash

set -euo pipefail

SCRIPT_DIR="$(cd -- "$(dirname -- "${BASH_SOURCE[0]}")" && pwd)"
REPO_ROOT="$(cd -- "${SCRIPT_DIR}/.." && pwd)"
BRINGUP="${REPO_ROOT}/scripts/bringup_i2c_slave_testunit.sh"

SLAVE_BUS="${BADGESNAKE_SLAVE_BUS:-1}"
MASTER_BUS="${BADGESNAKE_MASTER_BUS:-3}"
ADDR="${BADGESNAKE_TARGET_ADDR:-0x30}"
WRITE_BYTE="${BADGESNAKE_WRITE_BYTE:-0x00}"
READ_LEN="${BADGESNAKE_READ_LEN:-1}"

usage() {
	cat <<'EOF'
Usage:
  validate_j7_to_j6_target_mode.sh

Environment:
  BADGESNAKE_SLAVE_BUS    Target/slave bus (default: 1)
  BADGESNAKE_MASTER_BUS   Initiator/controller bus (default: 3)
  BADGESNAKE_TARGET_ADDR  7-bit slave address (default: 0x30)
  BADGESNAKE_WRITE_BYTE   Byte used for the write smoke test (default: 0x00)
  BADGESNAKE_READ_LEN     Bytes requested for the read smoke test (default: 1)

Expected wiring:
  - J6 shorted to J7 so /dev/i2c-${MASTER_BUS} can act as a second controller
    against the slave backend instantiated on /dev/i2c-${SLAVE_BUS}

What it does:
  1. Ensures slave-testunit is bound on the slave bus
  2. Runs one write from master bus to target address
  3. Runs one read from master bus to target address
  4. Prints the captured slave-side dmesg diagnostics
EOF
}

if [[ "${1:-}" == "-h" || "${1:-}" == "--help" ]]; then
	usage
	exit 0
fi

for cmd in i2ctransfer dmesg; do
	command -v "${cmd}" >/dev/null 2>&1 || {
		echo "missing required command: ${cmd}" >&2
		exit 1
	}
done

[[ -x "${BRINGUP}" ]] || {
	echo "missing helper: ${BRINGUP}" >&2
	exit 1
}

echo "Ensuring slave-testunit is present on i2c-${SLAVE_BUS} at ${ADDR}..."
"${BRINGUP}" start "${SLAVE_BUS}" "${ADDR}"
sleep 1
"${BRINGUP}" status "${SLAVE_BUS}" "${ADDR}"

echo
echo "Write test: /dev/i2c-${MASTER_BUS} -> ${ADDR} byte ${WRITE_BYTE}"
dmesg -C
i2ctransfer -f -y "${MASTER_BUS}" "w1@${ADDR}" "${WRITE_BYTE}"
write_dmesg="$(dmesg)"

echo
echo "Read test: /dev/i2c-${MASTER_BUS} -> ${ADDR} length ${READ_LEN}"
dmesg -C
read_output="$(i2ctransfer -f -y "${MASTER_BUS}" "r${READ_LEN}@${ADDR}")"
read_dmesg="$(dmesg)"

echo
echo "Read returned: ${read_output}"

echo
echo "Write-side slave diagnostics:"
if [[ -n "${write_dmesg}" ]]; then
	printf '%s\n' "${write_dmesg}"
else
	echo "(none)"
fi

echo
echo "Read-side slave diagnostics:"
if [[ -n "${read_dmesg}" ]]; then
	printf '%s\n' "${read_dmesg}"
else
	echo "(none)"
fi
