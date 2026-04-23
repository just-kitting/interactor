#!/usr/bin/env bash

set -euo pipefail

SCRIPT_DIR="$(cd -- "$(dirname -- "${BASH_SOURCE[0]}")" && pwd)"
REPO_ROOT="$(cd -- "${SCRIPT_DIR}/.." && pwd)"
OUT_BASE="${REPO_ROOT}/artifacts/state-capture"
TIMESTAMP="$(date -u +%Y%m%dT%H%M%SZ)"
OUT_DIR="${OUT_BASE}/${TIMESTAMP}"
WORKTREE_ARCHIVE="${OUT_DIR}/repo-working-tree.tar.gz"
STATE_ARCHIVE="${OUT_BASE}/beaglebadge-state-${TIMESTAMP}.tar.gz"

mkdir -p "${OUT_DIR}/system" "${OUT_DIR}/repo" "${OUT_DIR}/bundles" "${OUT_DIR}/boot"

capture_cmd() {
	local name="$1"
	shift
	"$@" > "${OUT_DIR}/system/${name}.txt" 2>&1 || true
}

copy_if_exists() {
	local src="$1"
	local dest="$2"
	if [[ -e "${src}" ]]; then
		cp -a "${src}" "${dest}"
	fi
}

echo "Capturing BeagleBadge state into ${OUT_DIR}"

capture_cmd uname uname -a
capture_cmd os-release cat /etc/os-release
capture_cmd armbian-release cat /etc/armbian-release
capture_cmd cmdline cat /proc/cmdline
capture_cmd mounts mount
capture_cmd findmnt findmnt -A
capture_cmd lsblk lsblk -o NAME,PATH,SIZE,FSTYPE,LABEL,PARTLABEL,MOUNTPOINTS
capture_cmd blkid blkid
capture_cmd df df -h
capture_cmd dpkg-query dpkg-query -W
capture_cmd dpkg-selections dpkg --get-selections
capture_cmd modules lsmod
capture_cmd dmesg-tail sh -c "dmesg | tail -n 400"

copy_if_exists /etc/fstab "${OUT_DIR}/system/fstab"
copy_if_exists /etc/fw_env.config "${OUT_DIR}/system/fw_env.config"

copy_if_exists /boot/uEnv.txt "${OUT_DIR}/boot/uEnv.txt"
copy_if_exists /boot/boot.scr "${OUT_DIR}/boot/boot.scr"
copy_if_exists /boot/armbianEnv.txt "${OUT_DIR}/boot/armbianEnv.txt"

if compgen -G "/boot/dtb/ti/k3-am62l3-badge*" > /dev/null; then
	cp -a /boot/dtb/ti/k3-am62l3-badge* "${OUT_DIR}/boot/"
fi

if compgen -G "/boot/dtb/ti/*.dtbo" > /dev/null; then
	mkdir -p "${OUT_DIR}/boot/dtbo"
	cp -a /boot/dtb/ti/*.dtbo "${OUT_DIR}/boot/dtbo/"
fi

{
	echo "repo_root=${REPO_ROOT}"
	echo "capture_time_utc=${TIMESTAMP}"
	echo "host=$(hostname -f 2>/dev/null || hostname)"
} > "${OUT_DIR}/repo/metadata.txt"

git -C "${REPO_ROOT}" status --short --branch > "${OUT_DIR}/repo/git-status.txt" 2>&1 || true
git -C "${REPO_ROOT}" submodule status > "${OUT_DIR}/repo/git-submodule-status.txt" 2>&1 || true
git -C "${REPO_ROOT}" log --oneline -n 20 > "${OUT_DIR}/repo/git-log.txt" 2>&1 || true
git -C "${REPO_ROOT}" diff --stat > "${OUT_DIR}/repo/git-diff-stat.txt" 2>&1 || true

git -C "${REPO_ROOT}" bundle create "${OUT_DIR}/bundles/badge-demo.bundle" HEAD

git -C "${REPO_ROOT}" submodule foreach --recursive '
	name="$(printf "%s" "$displaypath" | tr "/" "_")"
	if git rev-parse --is-inside-work-tree >/dev/null 2>&1; then
		git status --short --branch > "'"${OUT_DIR}"'/repo/${name}-status.txt" 2>&1 || true
		git log --oneline -n 20 > "'"${OUT_DIR}"'/repo/${name}-log.txt" 2>&1 || true
	fi
' >/dev/null

tar \
	--exclude=.git \
	--exclude=.cache \
	--exclude=artifacts \
	--exclude='components/*/output' \
	-czf "${WORKTREE_ARCHIVE}" \
	-C "${REPO_ROOT}" .

tar -czf "${STATE_ARCHIVE}" -C "${OUT_BASE}" "${TIMESTAMP}"

cat <<EOF
State capture complete.

Directory:
  ${OUT_DIR}

Archives:
  ${WORKTREE_ARCHIVE}
  ${STATE_ARCHIVE}

Next step:
  Copy ${STATE_ARCHIVE} off the BeagleBadge before replacing the microSD image.
EOF
