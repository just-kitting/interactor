# BeagleBadge I2C Device-Tree Status

## Live mapping on this board

The currently running device tree exposes:

- `/dev/i2c-0` -> `i2c@20000000` (`main_i2c0`)
- `/dev/i2c-2` -> `i2c@2b200000` (`wkup_i2c0`)

The Badge QWIIC buses are not currently exposed:

- `main_i2c1` at `0x20010000` is present but `status = "disabled"`
- `main_i2c2` at `0x20020000` is present but `status = "disabled"`

This explains why probing J6 through `/dev/i2c-0` was incorrect.

## Board connector mapping

Per the BeagleBadge documentation:

- J6 -> `I2C1`
- J7 -> `I2C2`

On this SoC and DT:

- J6 should map to `main_i2c1` / `i2c@20010000`
- J7 should map to `main_i2c2` / `i2c@20020000`

## Current DTB status

The shipped `/boot/dtb/ti/k3-am62l3-badge.dtb` already contains:

- `main_i2c1_default_pins`
- `main_i2c1` node
- `main_i2c2` node

What is missing:

- `main_i2c1` is not enabled
- `main_i2c2` is not enabled
- there is no `main_i2c2_default_pins` block in the shipped badge DTB

## Candidate fix

`overlays/beaglebadge-qwiic-i2c.dtso` enables:

- `main_i2c1`
- `main_i2c2`
- a new pinctrl block for `I2C2_SCL` and `I2C2_SDA`

Use `scripts/validate_qwiic_i2c_overlay.sh` to compile and merge the overlay against the live badge DTB before attempting boot-time activation.

Use `scripts/install_qwiic_i2c_overlay.sh` to:

- compile the overlay into `/boot/dtb/ti/k3-am62l3-badge-qwiic-i2c.dtbo`
- back up `/boot/uEnv.txt`
- append the overlay to `name_overlays`

After installation, a reboot is required to verify that Linux exposes the QWIIC buses.

## Current staged state on this board

This installation step has already been performed locally:

- overlay installed at `/boot/dtb/ti/k3-am62l3-badge-qwiic-i2c.dtbo`
- `/boot/uEnv.txt` updated to append `ti/k3-am62l3-badge-qwiic-i2c.dtbo`
- backup written to `/boot/uEnv.txt.bak.20260415T105230Z`

The remaining work is reboot-time validation.
