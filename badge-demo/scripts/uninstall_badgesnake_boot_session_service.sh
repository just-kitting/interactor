#!/usr/bin/env bash

set -euo pipefail

UNIT_DST="/etc/systemd/system/badgesnake-boot-session.service"

systemctl disable --now badgesnake-boot-session.service || true
rm -f "${UNIT_DST}"
systemctl daemon-reload

echo "Removed badgesnake-boot-session.service"
