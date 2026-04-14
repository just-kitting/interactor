#!/bin/sh
set -eu

pkg_dir="/usr/lib/linux-u-boot-vendor-edge-beaglebadge"

echo "== Packaged U-Boot metadata =="
sed -n '1,200p' "$pkg_dir/u-boot-metadata.sh"
echo

echo "== Artifact checksums: packaged vs /boot =="
for f in tiboot3.bin tispl.bin u-boot.img; do
    echo "-- $f --"
    sha256sum "$pkg_dir/$f" "/boot/$f"
    echo
done
