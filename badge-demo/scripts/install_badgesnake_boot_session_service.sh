#!/usr/bin/env bash

set -euo pipefail

SCRIPT_DIR="$(cd -- "$(dirname -- "${BASH_SOURCE[0]}")" && pwd)"
REPO_ROOT="$(cd -- "${SCRIPT_DIR}/.." && pwd)"
UNIT_SRC="${REPO_ROOT}/systemd/badgesnake-boot-session.service"
UNIT_DST="/etc/systemd/system/badgesnake-boot-session.service"

if [[ ! -f "${UNIT_SRC}" ]]; then
	echo "Missing unit file: ${UNIT_SRC}" >&2
	exit 1
fi

install -m 0644 "${UNIT_SRC}" "${UNIT_DST}"
systemctl daemon-reload
systemctl enable badgesnake-boot-session.service
systemctl restart badgesnake-boot-session.service

echo "Installed and enabled badgesnake-boot-session.service"
echo "Check status with:"
echo "  systemctl status badgesnake-boot-session.service"
echo "Attach to the workspace tmux session with:"
echo "  tmux attach -t badgesnake"
