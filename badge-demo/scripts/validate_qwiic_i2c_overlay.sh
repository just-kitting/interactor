#!/usr/bin/env bash
set -euo pipefail

repo_root="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
overlay_src="${repo_root}/overlays/beaglebadge-qwiic-i2c.dtso"
overlay_dtb="${repo_root}/.cache/beaglebadge-qwiic-i2c.dtbo"
merged_dtb="${repo_root}/.cache/k3-am62l3-badge-qwiic-i2c.dtb"
base_dtb="/boot/dtb/ti/k3-am62l3-badge.dtb"

mkdir -p "${repo_root}/.cache"

dtc -@ -I dts -O dtb -o "${overlay_dtb}" "${overlay_src}"
fdtoverlay -i "${base_dtb}" -o "${merged_dtb}" "${overlay_dtb}"

echo "Validated overlay against ${base_dtb}"
echo
dtc -I dtb -O dts "${merged_dtb}" | grep -nE 'i2c@20010000|i2c@20020000|badge-main-i2c2-default-pins|status = \"okay\"' | sed -n '1,120p'
