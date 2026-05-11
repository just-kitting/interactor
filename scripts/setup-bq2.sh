#!/bin/bash
set -euo pipefail

SCRIPTS_DIR=$(dirname $(realpath $0))
echo "Running $SCRIPTS_DIR/setup-bq2.sh"

cd $SCRIPTS_DIR/../utilities/bq2
shards install
shards build
cd $SCRIPTS_DIR/..
mkdir -p bin
cp ./utilities/bq2/bin/bq2 bin/
bin/bq2 --install
bin/sysroot-builder -s Alpine --workdir=workdir/badge-snake
ls scripts/
cp scripts/sysroot-build-plan.json workdir/badge-snake/seed-rootfs/bq2-rootfs/var/lib/sysroot-build-plan.json
bin/sysroot-runner --workdir=workdir/badge-snake
#cd workdir/badge-snake/seed-rootfs/root
#git clone https://github.com/just-kitting/interactor --recurse-submodules

