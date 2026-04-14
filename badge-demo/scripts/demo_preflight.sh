#!/bin/sh
set -eu

fail=0

check_file() {
    path="$1"
    if [ -e "$path" ]; then
        printf '[ok] %s\n' "$path"
    else
        printf '[missing] %s\n' "$path"
        fail=1
    fi
}

check_cmd() {
    cmd="$1"
    if command -v "$cmd" >/dev/null 2>&1; then
        printf '[ok] command %s\n' "$cmd"
    else
        printf '[missing] command %s\n' "$cmd"
        fail=1
    fi
}

echo '== BadgeSnake demo preflight =='
echo

echo '== repo structure =='
check_file components/armbian-build
check_file components/battlesnake-rules
check_file components/beaglebadge
check_file components/beagleconnect-zepto
check_file docs/ProductSpec.md
check_file docs/Protocol.md
check_file docs/BootAndRecovery.md
check_file docs/HardwareOutputs.md
check_file fixtures/protocol/path_tokens.json
echo

echo '== helper tooling =='
check_cmd fw_printenv
check_cmd sha256sum
echo

echo '== device exposure =='
check_file /dev/fb0
check_file /dev/i2c-0
check_file /dev/i2c-2
check_file /proc/mtd
echo

echo '== missing-input state =='
if tail -n 1 docs/MissingInputs.md | grep -q '^>'; then
    echo '[missing] docs/MissingInputs.md still ends with an unanswered block quote'
    fail=1
else
    echo '[ok] docs/MissingInputs.md reports no unanswered blocking questions'
fi
echo

echo '== u-boot artifact probe =='
if ./scripts/check_uboot_artifacts.sh >/dev/null 2>&1; then
    echo '[ok] scripts/check_uboot_artifacts.sh ran successfully'
else
    echo '[failed] scripts/check_uboot_artifacts.sh'
    fail=1
fi

echo
echo '== output probe =='
if ./scripts/probe_badge_outputs.sh >/dev/null 2>&1; then
    echo '[ok] scripts/probe_badge_outputs.sh ran successfully'
else
    echo '[failed] scripts/probe_badge_outputs.sh'
    fail=1
fi

echo
if [ "$fail" -eq 0 ]; then
    echo 'Preflight passed.'
else
    echo 'Preflight failed.'
    exit 1
fi
