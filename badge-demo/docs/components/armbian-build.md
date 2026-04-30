# Armbian Build

This submodule contains the code for building Armbian for BeagleBadge.

Upstream: https://github.com/armbian/build
TI fork: https://github.com/TexasInstruments/armbian-build
Our fork: https://github.com/embedconsult/armbian-build

Modifications to improve Armbian should be performed directly to this repository suitable for all Armbian users.

Commit locally and the maintainer will provide feedback.

BadgeSnake will use this submodule repo to understand the running distro image, including bootloader and kernel.

## Current BadgeSnake Use

The current immediate use is rebuilding the BeagleBadge `vendor-edge` kernel packages so:

- the staged `I2C_SLAVE_TESTUNIT` and `I2C_SLAVE_EEPROM` config changes can be boot-tested
- the staged AM62L `i2c-omap` slave-registration patch can be compiled into a test kernel

Use the top-level wrapper:

```sh
./scripts/build_beaglebadge_vendor_edge_kernel_x86_docker.sh
```

This deliberately builds kernel packages only. It does not yet authorize replacing the current starting microSD image.

The current BeagleBadge-specific kernel patch for target-mode bring-up lives here:

- `components/armbian-build/patch/kernel/archive/k3-6.12/0001-Stage-OMAP-I2C-slave-registration-support.patch`
- `components/armbian-build/patch/kernel/archive/k3-6.12/0002-Fix-OMAP-slave-helper-declaration-order.patch`

## Update Procedure

1. Update this submodule to the intended branch and commit.
2. Commit any BadgeSnake-required Armbian changes inside the submodule first.
3. For K3 kernel driver work, carry the live test patch in the Armbian kernel patch directory used by the build.
4. Commit the superproject pointer update.
5. Run the top-level x86 Docker wrapper to produce test kernel packages.
6. Preserve current BeagleBadge state with `./scripts/capture_beaglebadge_state.sh` before discussing any fresh image replacement.
