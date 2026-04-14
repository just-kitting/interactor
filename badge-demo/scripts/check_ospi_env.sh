#!/bin/sh
set -eu

config_file="${1:-/tmp/badgesnake_fw_env.config}"

echo "== U-Boot packaged environment mode =="
grep -E 'CONFIG_ENV_(IS_NOWHERE|IS_IN_SPI_FLASH|SIZE=)' /usr/lib/linux-u-boot-vendor-edge-beaglebadge/u-boot-config-target-1 || true
echo

echo "== Live MTD layout =="
cat /proc/mtd
echo

cat >"$config_file" <<'EOF'
/dev/mtd3 0x0 0x40000 0x20000
/dev/mtd4 0x0 0x40000 0x20000
EOF

echo "== Candidate fw_env.config =="
cat "$config_file"
echo

echo "== fw_printenv probe =="
if fw_printenv -c "$config_file" 2>&1; then
    echo
    echo "fw_printenv succeeded with candidate config."
else
    rc=$?
    echo
    echo "fw_printenv failed with rc=$rc."
    echo "This is expected on the current packaged U-Boot because CONFIG_ENV_IS_NOWHERE=y."
    exit 0
fi
