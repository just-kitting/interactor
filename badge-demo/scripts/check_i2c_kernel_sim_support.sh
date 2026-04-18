#!/bin/sh
set -eu

KVER=$(uname -r)
MODULE_DIR="/lib/modules/$KVER"

echo "Kernel: $KVER"
echo

check_module() {
  name="$1"
  label="$2"
  if modinfo "$name" >/dev/null 2>&1; then
    echo "[yes] $label ($name)"
  else
    echo "[no]  $label ($name)"
  fi
}

echo "Kernel module availability:"
check_module i2c-stub "I2C stub emulator"
check_module i2c_stub "I2C stub emulator"
check_module i2c-slave-eeprom "I2C slave EEPROM backend"
check_module slave-24c02 "I2C slave EEPROM backend alias"
echo

echo "I2C bus sysfs endpoints:"
for bus in /sys/bus/i2c/devices/i2c-*; do
  [ -e "$bus" ] || continue
  printf '%s:' "$bus"
  for entry in new_device delete_device name; do
    if [ -e "$bus/$entry" ]; then
      printf ' %s' "$entry"
    fi
  done
  printf '\n'
done

echo
echo "Assessment:"
cat <<'EOF'
- `i2c-stub` is a memory-backed fake SMBus device model. It is useful for testing
  kernel client drivers, but it is not an interactive target path for Boardie or
  MicroBlocks logic.
- A proper Linux-visible Boardie simulation needs a kernel-visible target or
  emulated adapter path that controller-side tools such as `i2ctransfer` can use.
- If `i2c-slave-eeprom` or another usable target backend is missing from the
  running kernel, the hosted browser bridge remains the only working path on this
  image.
EOF
