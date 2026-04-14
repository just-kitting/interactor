#!/usr/bin/env bash
set -euo pipefail

kernel_release="$(uname -r)"

echo "Kernel: ${kernel_release}"
echo
echo "Detected I2C adapters:"
if compgen -G "/sys/bus/i2c/devices/i2c-*" > /dev/null; then
  ls /sys/bus/i2c/devices | grep '^i2c-' | sed 's/^/  - /'
else
  echo "  - none"
fi

echo
echo "Checking i2c-stub module support..."
if modinfo i2c-stub >/dev/null 2>&1; then
  echo "  - i2c-stub module is available"
else
  echo "  - i2c-stub module is not available in the current kernel package"
  echo "  - sim:// remains the primary local transport test path"
  exit 1
fi
