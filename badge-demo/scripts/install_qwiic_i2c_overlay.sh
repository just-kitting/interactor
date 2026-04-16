#!/usr/bin/env bash
set -euo pipefail

repo_root="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
overlay_src="${repo_root}/overlays/beaglebadge-qwiic-i2c.dtso"
overlay_dst="/boot/dtb/ti/k3-am62l3-badge-qwiic-i2c.dtbo"
uenv="/boot/uEnv.txt"
backup_suffix="$(date -u +%Y%m%dT%H%M%SZ)"

tmp_dtbo="$(mktemp)"
cleanup() {
	rm -f "${tmp_dtbo}"
}
trap cleanup EXIT

dtc -@ -I dts -O dtb -o "${tmp_dtbo}" "${overlay_src}"

cp -a "${uenv}" "${uenv}.bak.${backup_suffix}"
cp -a "${tmp_dtbo}" "${overlay_dst}"

python3 - <<'PY'
from pathlib import Path

uenv = Path("/boot/uEnv.txt")
overlay = "ti/k3-am62l3-badge-qwiic-i2c.dtbo"
lines = uenv.read_text().splitlines()
updated = []
changed = False

for line in lines:
    if line.startswith("    name_overlays=") or line.startswith("name_overlays="):
        key, value = line.split("=", 1)
        parts = value.split()
        if overlay not in parts:
            parts.append(overlay)
            changed = True
        line = f"{key}={' '.join(parts)}"
    updated.append(line)

if not changed and not any(overlay in line for line in updated):
    updated.append(f"name_overlays={overlay}")

uenv.write_text("\n".join(updated) + "\n")
PY

echo "Installed ${overlay_dst}"
echo "Backed up ${uenv} to ${uenv}.bak.${backup_suffix}"
echo
grep -n "name_overlays" "${uenv}" || true
echo
echo "Reboot is required before J6/J7 I2C adapters can be validated."
