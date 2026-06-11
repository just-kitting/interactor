#!/bin/sh
set -eu

SCRIPT_DIR="$(CDPATH= cd -- "$(dirname "$0")" && pwd)"
REPO_ROOT="$(CDPATH= cd -- "$SCRIPT_DIR/.." && pwd)"
MODULE_PATH="$("$SCRIPT_DIR/build_opt3001_oot_module.sh")"
KERNEL_RELEASE="$(uname -r)"
DEST_DIR="/lib/modules/$KERNEL_RELEASE/extra"
MODULES_LOAD_CONF="/etc/modules-load.d/opt3001.conf"

if [ ! -f "$MODULE_PATH" ]; then
    echo "built module not found at $MODULE_PATH" >&2
    exit 1
fi

mkdir -p "$DEST_DIR"
install -m 0644 "$MODULE_PATH" "$DEST_DIR/opt3001.ko"
depmod -a "$KERNEL_RELEASE"
modprobe opt3001
printf 'opt3001\n' > "$MODULES_LOAD_CONF"

echo "Installed opt3001.ko to $DEST_DIR"
echo "Configured $MODULES_LOAD_CONF"
