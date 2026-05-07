#!/bin/bash
set -euo pipefail

SCRIPTS_DIR=$(dirname $(realpath $0))
echo "Running $SCRIPTS_DIR/setup-repo.sh"

cd $SCRIPTS_DIR/../workdir/badge-snake/seed-rootfs/root
git clone --recurse-submodules https://github.com/just-kitting/interactor

