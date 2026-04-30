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

## Current Returned Artifacts

The current x86-host build artifacts have been copied back into `components/armbian-build/output/` in this repo. The relevant `.deb` files are:

- `linux-image-vendor-edge-k3`
- `linux-dtb-vendor-edge-k3`
- `linux-headers-vendor-edge-k3`
- `linux-libc-dev-vendor-edge-k3`

These packages currently use the same Debian package version string as the installed kernel packages on the board: `26.02.0-trunk`.

Because of that, the board-side install step should be a reinstall from the local `.deb` files, not a normal upgrade.

Use:

```sh
./scripts/install_beaglebadge_vendor_edge_kernel_artifacts.sh
```

Then reboot and continue with the `slave-testunit` bring-up steps in `docs/I2CSlaveBringup.md`.

## Current Correction

After validating the live kernel, it turned out the corrected `components/armbian-build` branch no longer carried the intended K3 config deltas for:

- `CONFIG_I2C_SLAVE_EEPROM`
- `CONFIG_I2C_SLAVE_TESTUNIT`

That means the previously returned `vendor-edge-k3` artifacts were built from the wrong effective config and do not contain `i2c-slave-testunit`.

Before continuing AM62L slave-mode bring-up, rebuild the kernel artifacts again from the restored config in `components/armbian-build/config/kernel/`.

## Mixed Artifact Directory

`components/armbian-build/output/debs/` can contain more than one `vendor-edge-k3` artifact set at the same time.

The current reinstall wrapper handles that by:

- scanning the available `linux-image-vendor-edge-k3` packages
- preferring the first one whose contents include `i2c-slave-testunit.ko`
- selecting the matching `dtb`, `headers`, and `libc-dev` packages from the same build suffix
- reinstalling the local QWIIC overlay after the DTB package refresh

That avoids accidentally reinstalling an older stale artifact set just because it was copied in earlier.

## Current Source-Of-Truth Caveat

The current Armbian kernel build still uses the source tuple from the build framework:

- `KERNELSOURCE=https://github.com/TexasInstruments/ti-linux-kernel`
- `KERNELBRANCH=branch:ti-linux-6.12.y-cicd`

So copying back a freshly built artifact set does **not** by itself prove that it included local changes staged in `components/ti-linux-kernel/`.

For the staged local `i2c-omap` slave-support patch, a successful validation build must either:

- be redirected to the local kernel source tree, or
- carry the patch through the Armbian kernel patch path, or
- otherwise prove in the build log that the patched kernel source was used

If the copied artifacts still only show the existing `C2876...` suffix, they are not a new post-`i2c-omap` test build.

On 2026-04-30, a newer copied artifact set with suffix `S3b4a...` did appear in `output/debs/`.
However, the corresponding build log still shows:

- `KERNELSOURCE='https://github.com/TexasInstruments/ti-linux-kernel'`
- `KERNELBRANCH='branch:ti-linux-6.12.y-cicd'`
- `kernel patching: 0 total patches; 0 applied; 0 with problems`

So that newer copied build is not yet evidence that the staged local `components/ti-linux-kernel` `i2c-omap` patch was compiled.

## Expected Install Notes

Two install-time messages are expected on the current board image and do not indicate failure:

- `update-initramfs: Symlink failed ... moving ...`:
  - `/boot` is FAT32 on this image, so Armbian cannot keep normal Linux symlinks there and falls back to `rename()`
- `_apt ... couldn't be accessed by user '_apt'`:
  - this warning appears when local `.deb` files are installed from a root-only path
  - the current reinstall wrapper stages the packages into `/var/tmp/badgesnake-kernel-debs/` first so future runs avoid that warning

## Why This Is The Current Build Target

- it is enough to validate the staged `CONFIG_I2C_SLAVE_TESTUNIT=m` and `CONFIG_I2C_SLAVE_EEPROM=m` changes
- it avoids jumping ahead to a full image replacement before state-preservation work is in place
- it keeps the next test focused on `i2c-omap` slave bring-up

## Local Module Iteration

There is a good reason to distinguish between the short-term and long-term work here:

- `i2c-slave-testunit` and `i2c-slave-eeprom` are built as modules when enabled
- `i2c-omap` is built into the K3 kernel image with `CONFIG_I2C_OMAP=y`

So yes, local module-only builds on the BeagleBadge are a valid faster iteration path for the stock slave test backends once the config is correct.

What module-only builds cannot avoid is the eventual `i2c-omap.c` work:

- changes to `i2c-omap` require rebuilding and booting a new kernel image
- they cannot be delivered as a standalone `.ko`

For the immediate missing-`slave-testunit` problem, the current rebuild is only correcting the package contents. After that, module-only iteration is the better path until bus-driver changes begin.
