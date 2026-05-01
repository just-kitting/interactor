#!/usr/bin/env bash

set -euo pipefail

SCRIPT_DIR="$(cd -- "$(dirname -- "${BASH_SOURCE[0]}")" && pwd)"
REPO_ROOT="$(cd -- "${SCRIPT_DIR}/.." && pwd)"
DEB_DIR="${REPO_ROOT}/components/armbian-build/output/debs"
STAGE_DIR="/var/tmp/badgesnake-kernel-debs"
QWIIC_OVERLAY_HELPER="${REPO_ROOT}/scripts/install_qwiic_i2c_overlay.sh"
SELECTED_SUFFIX="${BADGESNAKE_BUILD_SUFFIX:-}"
shopt -s nullglob

image_matches=( "${DEB_DIR}/linux-image-vendor-edge-k3_26.02.0-trunk_arm64__"*.deb )
if [[ ${#image_matches[@]} -lt 1 ]]; then
	echo "No linux-image-vendor-edge-k3 artifact found in ${DEB_DIR}" >&2
	exit 1
fi

image_deb=""
filtered_matches=()
for candidate in "${image_matches[@]}"; do
	listing="$(dpkg-deb -c "${candidate}" 2>/dev/null || true)"
	if grep -q 'i2c-slave-testunit\.ko' <<< "${listing}"; then
		filtered_matches+=( "${candidate}" )
	fi
done

if [[ -n "${SELECTED_SUFFIX}" ]]; then
	image_deb="${DEB_DIR}/linux-image-vendor-edge-k3_26.02.0-trunk_arm64__${SELECTED_SUFFIX}"
	if [[ ! -f "${image_deb}" ]]; then
		echo "Requested BADGESNAKE_BUILD_SUFFIX does not exist: ${SELECTED_SUFFIX}" >&2
		exit 1
	fi
elif [[ ${#filtered_matches[@]} -gt 0 ]]; then
	IFS=$'\n' read -r -d '' -a sorted_images < <(printf '%s\n' "${filtered_matches[@]}" | while IFS= read -r f; do printf '%s\t%s\n' "$(stat -c '%Z' "${f}")" "${f}"; done | sort -rn | cut -f2- && printf '\0')
	image_deb="${sorted_images[0]}"
else
	IFS=$'\n' read -r -d '' -a sorted_images < <(printf '%s\n' "${image_matches[@]}" | while IFS= read -r f; do printf '%s\t%s\n' "$(stat -c '%Z' "${f}")" "${f}"; done | sort -rn | cut -f2- && printf '\0')
	image_deb="${sorted_images[0]}"
fi

build_suffix="${image_deb##*__}"

dtb_deb="${DEB_DIR}/linux-dtb-vendor-edge-k3_26.02.0-trunk_arm64__${build_suffix}"
headers_deb="${DEB_DIR}/linux-headers-vendor-edge-k3_26.02.0-trunk_arm64__${build_suffix}"
libc_deb="${DEB_DIR}/linux-libc-dev-vendor-edge-k3_26.02.0-trunk_arm64__${build_suffix}"

for f in "${image_deb}" "${dtb_deb}" "${headers_deb}" "${libc_deb}"; do
	if [[ ! -f "${f}" ]]; then
		echo "Missing matching artifact for selected build suffix ${build_suffix}: ${f}" >&2
		exit 1
	fi
done

echo "Reinstalling locally built BeagleBadge vendor-edge kernel artifacts:"
printf '  %s\n' "${image_deb}" "${dtb_deb}" "${headers_deb}" "${libc_deb}"
if [[ -n "${SELECTED_SUFFIX}" ]]; then
	echo "Using requested BADGESNAKE_BUILD_SUFFIX=${SELECTED_SUFFIX}"
fi

rm -rf "${STAGE_DIR}"
install -d -m 0755 "${STAGE_DIR}"
install -m 0644 "${image_deb}" "${STAGE_DIR}/"
install -m 0644 "${dtb_deb}" "${STAGE_DIR}/"
install -m 0644 "${headers_deb}" "${STAGE_DIR}/"
install -m 0644 "${libc_deb}" "${STAGE_DIR}/"

export DEBIAN_FRONTEND=noninteractive
apt-get update
apt-get install --reinstall -y \
	"${STAGE_DIR}/$(basename "${image_deb}")" \
	"${STAGE_DIR}/$(basename "${dtb_deb}")" \
	"${STAGE_DIR}/$(basename "${headers_deb}")" \
	"${STAGE_DIR}/$(basename "${libc_deb}")"

if [[ -x "${QWIIC_OVERLAY_HELPER}" ]]; then
	echo
	echo "Reinstalling local QWIIC overlay after DTB package refresh..."
	"${QWIIC_OVERLAY_HELPER}"
fi

echo
echo "Installed local kernel artifacts."
echo "Expected install notes on this board:"
echo "  - FAT32 /boot cannot keep Linux symlinks, so Armbian falls back to rename()"
echo "  - the QWIIC overlay is reinstalled after the DTB package refresh"
echo "Next recommended steps:"
echo "  1. reboot"
echo "  2. uname -a"
echo "  3. modinfo i2c-slave-testunit"
echo "  4. ls /sys/bus/i2c/devices | grep '^i2c-[13]$'"
echo "  5. ./scripts/bringup_i2c_slave_testunit.sh start 1 0x30"
