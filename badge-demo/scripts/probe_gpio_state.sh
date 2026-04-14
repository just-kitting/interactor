#!/bin/sh
set -eu

echo "== gpio chips =="
find /dev -maxdepth 1 -type c -name 'gpiochip*' | sort
echo

if command -v gpioinfo >/dev/null 2>&1; then
    echo "== gpioinfo =="
    gpioinfo
else
    echo "gpioinfo not installed"
fi
