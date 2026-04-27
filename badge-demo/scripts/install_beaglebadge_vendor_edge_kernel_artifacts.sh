#!/usr/bin/env bash

set -euo pipefail

SCRIPT_DIR="$(cd -- "$(dirname -- "${BASH_SOURCE[0]}")" && pwd)"
REPO_ROOT="$(cd -- "${SCRIPT_DIR}/.." && pwd)"
DEB_DIR="${REPO_ROOT}/components/armbian-build/output/debs"

required=(
	"${DEB_DIR}/linux-image-vendor-edge-k3_26.02.0-trunk_arm64__"*.deb
	"${DEB_DIR}/linux-dtb-vendor-edge-k3_26.02.0-trunk_arm64__"*.deb
	"${DEB_DIR}/linux-headers-vendor-edge-k3_26.02.0-trunk_arm64__"*.deb
	"${DEB_DIR}/linux-libc-dev-vendor-edge-k3_26.02.0-trunk_arm64__"*.deb
)

for pattern in "${required[@]}"; do
	matches=( $pattern )
	if [[ ${#matches[@]} -ne 1 ]]; then
		echo "Expected exactly one package matching: $pattern" >&2
		exit 1
	fi
done

image_deb=( "${DEB_DIR}/linux-image-vendor-edge-k3_26.02.0-trunk_arm64__"*.deb )
dtb_deb=( "${DEB_DIR}/linux-dtb-vendor-edge-k3_26.02.0-trunk_arm64__"*.deb )
headers_deb=( "${DEB_DIR}/linux-headers-vendor-edge-k3_26.02.0-trunk_arm64__"*.deb )
libc_deb=( "${DEB_DIR}/linux-libc-dev-vendor-edge-k3_26.02.0-trunk_arm64__"*.deb )

echo "Reinstalling locally built BeagleBadge vendor-edge kernel artifacts:"
printf '  %s\n' "${image_deb[0]}" "${dtb_deb[0]}" "${headers_deb[0]}" "${libc_deb[0]}"

export DEBIAN_FRONTEND=noninteractive
apt-get update
apt-get install --reinstall -y \
	"${image_deb[0]}" \
	"${dtb_deb[0]}" \
	"${headers_deb[0]}" \
	"${libc_deb[0]}"

echo
echo "Installed local kernel artifacts."
echo "Next recommended steps:"
echo "  1. reboot"
echo "  2. uname -a"
echo "  3. modinfo i2c-slave-testunit"
echo "  4. ./scripts/bringup_i2c_slave_testunit.sh start 1 0x30"
