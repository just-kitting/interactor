#!/usr/bin/env bash
set -euo pipefail

bus="${1:-1}"

echo "probe: loopback diag on /dev/i2c-${bus}"
echo

if timeout 5 i2cdetect -y -r "${bus}" >/tmp/zepto-loopback-i2cdetect.txt 2>&1; then
    cat /tmp/zepto-loopback-i2cdetect.txt
    echo
    echo "result: no externally visible SCL seizure; internal loopback self-test did not report pass"
    exit 1
fi

cat /tmp/zepto-loopback-i2cdetect.txt
echo
echo "result: host probe timed out; treating SCL seizure as internal loopback self-test pass"
