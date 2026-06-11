#!/bin/sh
set -eu

SCRIPT_DIR="$(CDPATH= cd -- "$(dirname "$0")" && pwd)"
REPO_ROOT="$(CDPATH= cd -- "$SCRIPT_DIR/.." && pwd)"
KERNEL_SOURCE="${OPT3001_KERNEL_SOURCE:-$REPO_ROOT/components/ti-linux-kernel}"
KERNEL_HEADERS="${OPT3001_KERNEL_HEADERS:-/lib/modules/$(uname -r)/build}"
BUILD_ROOT="${OPT3001_BUILD_ROOT:-$REPO_ROOT/.cache/opt3001-oot/$(uname -r)}"
SOURCE_FILE="$KERNEL_SOURCE/drivers/iio/light/opt3001.c"

if [ ! -f "$SOURCE_FILE" ]; then
    echo "opt3001 source not found at $SOURCE_FILE" >&2
    exit 1
fi

if [ ! -d "$KERNEL_HEADERS" ]; then
    echo "kernel headers not found at $KERNEL_HEADERS" >&2
    exit 1
fi

mkdir -p "$BUILD_ROOT"
cp "$SOURCE_FILE" "$BUILD_ROOT/opt3001.c"
cat > "$BUILD_ROOT/Makefile" <<'EOF'
obj-m += opt3001.o
KDIR ?= /lib/modules/$(shell uname -r)/build
all:
	$(MAKE) -C $(KDIR) M=$(PWD) modules
clean:
	$(MAKE) -C $(KDIR) M=$(PWD) clean
EOF

make -C "$KERNEL_HEADERS" M="$BUILD_ROOT" modules >&2
printf '%s\n' "$BUILD_ROOT/opt3001.ko"
