#!/usr/bin/env bash

set -euo pipefail

SCRIPT_DIR="$(cd -- "$(dirname -- "${BASH_SOURCE[0]}")" && pwd)"
REPO_ROOT="$(cd -- "${SCRIPT_DIR}/.." && pwd)"
INSTALL_HELPER="${REPO_ROOT}/scripts/install_beaglebadge_vendor_edge_kernel_artifacts.sh"
REBOOT_DELAY="${BADGESNAKE_REBOOT_DELAY:-5}"

usage() {
	cat <<'EOF'
Usage:
  install_latest_kernel_and_reboot.sh [--delay SECONDS]

Environment:
  BADGESNAKE_BUILD_SUFFIX   Optional exact build suffix to install
  BADGESNAKE_REBOOT_DELAY   Reboot delay in seconds (default: 5)

Notes:
  - Uses scripts/install_beaglebadge_vendor_edge_kernel_artifacts.sh
  - Reboots only after the install helper exits successfully
  - Runs sync before reboot
EOF
}

while [[ $# -gt 0 ]]; do
	case "$1" in
		--delay)
			REBOOT_DELAY="$2"
			shift 2
			;;
		-h|--help)
			usage
			exit 0
			;;
		*)
			echo "Unknown argument: $1" >&2
			usage >&2
			exit 1
			;;
	esac
done

if [[ ! -x "${INSTALL_HELPER}" ]]; then
	echo "Missing install helper: ${INSTALL_HELPER}" >&2
	exit 1
fi

if ! [[ "${REBOOT_DELAY}" =~ ^[0-9]+$ ]]; then
	echo "BADGESNAKE_REBOOT_DELAY / --delay must be an integer number of seconds" >&2
	exit 1
fi

echo "Installing BeagleBadge vendor-edge kernel artifacts..."
if [[ -n "${BADGESNAKE_BUILD_SUFFIX:-}" ]]; then
	echo "Pinned build suffix: ${BADGESNAKE_BUILD_SUFFIX}"
fi
"${INSTALL_HELPER}"

echo
echo "Install completed successfully."
echo "Current running kernel before reboot:"
uname -a
echo
echo "Rebooting in ${REBOOT_DELAY}s..."
sync
sleep "${REBOOT_DELAY}"
exec systemctl reboot
