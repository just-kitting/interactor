#!/usr/bin/env bash
set -euo pipefail

repo_root="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
overlay_src="${repo_root}/overlays/beaglebadge-zepto-control.dtso"
overlay_dtb="${repo_root}/.cache/beaglebadge-zepto-control.dtbo"
merged_dtb="${repo_root}/.cache/k3-am62l3-badge-zepto-control.dtb"
base_dtb="/boot/dtb/ti/k3-am62l3-badge.dtb"

mkdir -p "${repo_root}/.cache"

dtc -@ -I dts -O dtb -o "${overlay_dtb}" "${overlay_src}"
fdtoverlay -i "${base_dtb}" -o "${merged_dtb}" "${overlay_dtb}"

echo "Validated overlay against ${base_dtb}"
echo
dtc -I dtb -O dts "${merged_dtb}" | grep -nE 'zepto-control-default-pins|gpio@600000|pinctrl-0 = <.*zepto_control_default_pins' | sed -n '1,120p'
