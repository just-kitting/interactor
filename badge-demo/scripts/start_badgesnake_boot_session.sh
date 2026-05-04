#!/usr/bin/env bash

set -euo pipefail

REPO_ROOT="/root/interactor/badge-demo"
SESSION_NAME="badgesnake"
BOOT_LOG_DIR="${REPO_ROOT}/artifacts/boot-status"
BOOT_LOG="${BOOT_LOG_DIR}/latest.txt"

install -d -m 0755 "${BOOT_LOG_DIR}"

{
	echo "timestamp=$(date -u +%Y-%m-%dT%H:%M:%SZ)"
	echo "hostname=$(hostname)"
	echo "kernel=$(uname -a)"
	echo "uptime=$(uptime -p 2>/dev/null || true)"
	echo "cwd=${REPO_ROOT}"
	echo "git_head=$(git -C "${REPO_ROOT}" rev-parse --short HEAD 2>/dev/null || true)"
	echo "git_branch=$(git -C "${REPO_ROOT}" rev-parse --abbrev-ref HEAD 2>/dev/null || true)"
	echo "git_status=$(git -C "${REPO_ROOT}" status --short --branch --ignore-submodules=dirty 2>/dev/null | tr '\n' ';' || true)"
} > "${BOOT_LOG}"

if ! command -v tmux >/dev/null 2>&1; then
	exit 0
fi

if tmux has-session -t "${SESSION_NAME}" 2>/dev/null; then
	exit 0
fi

tmux new-session -d -s "${SESSION_NAME}" -c "${REPO_ROOT}" \
	"printf 'BadgeSnake boot session started at %s\n\n' \"\$(date -u +%Y-%m-%dT%H:%M:%SZ)\"; \
	 uname -a; printf '\n'; \
	 git status --short --branch --ignore-submodules=dirty || true; \
	 printf '\nBoot log: %s\n' '${BOOT_LOG}'; \
	 exec bash -l"
