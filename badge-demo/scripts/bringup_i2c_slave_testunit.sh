#!/bin/sh
set -eu

usage() {
  cat <<'EOF'
Usage:
  bringup_i2c_slave_testunit.sh start [bus] [addr]
  bringup_i2c_slave_testunit.sh stop [bus] [addr]
  bringup_i2c_slave_testunit.sh status [bus] [addr]

Defaults:
  bus  = 1
  addr = 0x30

Notes:
  - The slave backend address is written to sysfs as 0x1000 + addr.
  - This requires a kernel with CONFIG_I2C_SLAVE_TESTUNIT available.
EOF
}

subcmd="${1:-}"
bus="${2:-1}"
addr="${3:-0x30}"

if [ -z "$subcmd" ]; then
  usage >&2
  exit 1
fi

bus_path="/sys/bus/i2c/devices/i2c-$bus"
new_device="$bus_path/new_device"
delete_device="$bus_path/delete_device"
addr_dec=$((addr))
slave_addr=$(printf '0x%x' $((0x1000 + addr_dec)))
dev_node="$bus-$(printf '%04x' "$addr_dec")"

modprobe_fallback() {
  modprobe i2c-slave-testunit 2>/dev/null || modprobe i2c_slave_testunit 2>/dev/null || return 1
}

case "$subcmd" in
  start)
    [ -e "$new_device" ] || { echo "missing $new_device" >&2; exit 1; }
    if [ -e "/sys/bus/i2c/devices/$dev_node" ]; then
      echo "slave-testunit already present at $dev_node"
      exit 0
    fi
    if ! modprobe_fallback; then
      echo "i2c-slave-testunit module not available on this kernel" >&2
      exit 1
    fi
    printf 'slave-testunit %s\n' "$slave_addr" >"$new_device"
    echo "started slave-testunit on bus $bus at addr $(printf '0x%x' "$addr_dec")"
    ;;
  stop)
    [ -e "$delete_device" ] || { echo "missing $delete_device" >&2; exit 1; }
    if [ ! -e "/sys/bus/i2c/devices/$dev_node" ]; then
      echo "no slave-testunit device at $dev_node"
      exit 0
    fi
    printf '%s\n' "$slave_addr" >"$delete_device"
    echo "stopped slave-testunit on bus $bus at addr $(printf '0x%x' "$addr_dec")"
    ;;
  status)
    if [ -e "/sys/bus/i2c/devices/$dev_node" ]; then
      echo "present: $dev_node"
      if [ -r "/sys/bus/i2c/devices/$dev_node/name" ]; then
        cat "/sys/bus/i2c/devices/$dev_node/name"
      fi
    else
      echo "absent: $dev_node"
      exit 1
    fi
    ;;
  *)
    usage >&2
    exit 1
    ;;
esac
