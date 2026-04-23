# Armbian Kernel Build

For the current I2C target-mode bring-up work, the requested host-side build is a BeagleBadge `vendor-edge` kernel package build, not a fresh microSD image build.

Use:

```sh
./scripts/build_beaglebadge_vendor_edge_kernel_x86_docker.sh
```

## Host Requirements

- Linux `x86_64`
- Docker installed and usable by the invoking user
- this repo checked out with submodules

## What The Script Does

The wrapper runs the Armbian build framework in `components/armbian-build` with this exact target:

```sh
./compile.sh \
  kernel \
  BOARD=beaglebadge \
  BRANCH=vendor-edge \
  RELEASE=trixie \
  BUILD_MINIMAL=yes \
  BUILD_DESKTOP=no \
  KERNEL_BTF=no
```

Armbian handles the Docker relaunch internally.

## Expected Output

The resulting kernel packages are written under:

- `components/armbian-build/output/debs/`

For the current target-mode work, the expected useful artifacts are the rebuilt `linux-image-*`, `linux-dtb-*`, and any corresponding modules package for the BeagleBadge `vendor-edge` family.

## Why This Is The Current Build Target

- it is enough to validate the staged `CONFIG_I2C_SLAVE_TESTUNIT=m` and `CONFIG_I2C_SLAVE_EEPROM=m` changes
- it avoids jumping ahead to a full image replacement before state-preservation work is in place
- it keeps the next test focused on `i2c-omap` slave bring-up
