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

The BeagleBadge `vendor-edge` build still pulls the upstream TI source tuple from the build framework:

- `KERNELSOURCE=https://github.com/TexasInstruments/ti-linux-kernel`
- `KERNELBRANCH=branch:ti-linux-6.12.y-cicd`

The AM62L slave-mode test kernel is now carried through the Armbian patch path instead of relying on the local `components/ti-linux-kernel/` worktree:

- `components/armbian-build/patch/kernel/archive/k3-6.12/0001-Stage-OMAP-I2C-slave-registration-support.patch`
- `components/armbian-build/patch/kernel/archive/k3-6.12/0002-Fix-OMAP-slave-helper-declaration-order.patch`
- `components/armbian-build/patch/kernel/archive/k3-6.12/0003-Handle-slave-TX-underflow-on-OMAP-I2C.patch`

Those patches were generated from the staged `components/ti-linux-kernel` `i2c-omap` change set and are now the build input that matters for x86-host kernel-package rebuilds.

For the next validation build, the important proof point in the host build log is no longer “local kernel source was used.”
It is that kernel patching reports at least this patch as applied from `archive/k3-6.12`.

If the host build log still reports:

- `kernel patching: 0 total patches; 0 applied; 0 with problems`

then the wrong Armbian worktree or branch was used for the build.

On 2026-04-30, a newer copied artifact set with suffix `S3b4a...` did appear in `output/debs/`.
However, the corresponding build log still shows:

- `KERNELSOURCE='https://github.com/TexasInstruments/ti-linux-kernel'`
- `KERNELBRANCH='branch:ti-linux-6.12.y-cicd'`
- `kernel patching: 0 total patches; 0 applied; 0 with problems`

So that newer copied build is not yet evidence that the staged AM62L slave-mode patch was compiled.

Later on 2026-04-30, a corrected build log and artifact set were copied back:

- build UUID: `1b53c454-a989-4c66-a5e6-e8b706e4bc05`
- artifact suffix: `P5507`
- kernel patching summary: `2 total patches; 2 applied; 0 with problems`

That `P5507` set is now the first valid AM62L slave-test kernel artifact set.
It has also been reinstalled on the BeagleBadge, but the board still needs to be rebooted into it before `slave-testunit` can be re-tested.

After runtime validation on `P5507`, binding succeeded but same-adapter traffic to the bound target still timed out.
The next host build should therefore include the additional slave-TX follow-up patch:

- `0003-Handle-slave-TX-underflow-on-OMAP-I2C.patch`

That follow-up build has now been returned and installed with:

- build UUID: `a7a0cecc-b133-4f5e-86d2-3dc2f1235eea`
- artifact suffix: `P024c`
- kernel patching summary: `3 total patches; 3 applied; 0 with problems`

The board has now been rebooted into `P024c` and validated:

- `slave-testunit` still binds on J6 / `i2c-1`
- forced same-adapter reads still time out
- adapter IRQ activity is present during the timeout

So the next follow-up is a fourth `k3-6.12` patch:

- `0004-Program-1-byte-FIFO-thresholds-in-slave-listen-mode.patch`

The next host build should therefore carry all four staged `k3-6.12` patches, then return to the same runtime validation sequence on the BeagleBadge.

If multiple returned artifact sets coexist in `components/armbian-build/output/debs`, the reinstall helper can be pinned explicitly:

```sh
BADGESNAKE_BUILD_SUFFIX='<full package suffix>.deb' ./scripts/install_beaglebadge_vendor_edge_kernel_artifacts.sh
```

Example:

```sh
BADGESNAKE_BUILD_SUFFIX='6.12.57-S3b4a-D0000-P9d8b-C2876Hb496-HK01ba-Vc222-Be8e3-R448a.deb' ./scripts/install_beaglebadge_vendor_edge_kernel_artifacts.sh
```

That pinned reinstall has now completed on the BeagleBadge. The next step is only the runtime reboot and validation into `P9d8b`.

The current diagnostic reinstall uses the same explicit-suffix path for the five-patch kernel:

```sh
BADGESNAKE_BUILD_SUFFIX='6.12.57-S3b4a-D0000-P0380-C2876Hb496-HK01ba-Vc222-Be8e3-R448a.deb' ./scripts/install_beaglebadge_vendor_edge_kernel_artifacts.sh
```

That `P0380` reinstall has completed on the BeagleBadge. The next step is to reboot into it and capture the new `i2c-omap` slave diagnostic log lines during the failing J6 read.

The current ISR-branch diagnostic reinstall uses:

```sh
BADGESNAKE_BUILD_SUFFIX='6.12.57-S3b4a-D0000-P803a-C2876Hb496-HK01ba-Vc222-Be8e3-R448a.deb' ./scripts/install_beaglebadge_vendor_edge_kernel_artifacts.sh
```

That `P803a` reinstall has completed on the BeagleBadge. The next step is to reboot into it and capture whether the J6 self-read is handled on the master ISR path instead of the slave ISR path.

The current slave-lifetime diagnostic reinstall uses:

```sh
BADGESNAKE_BUILD_SUFFIX='6.12.57-S3b4a-D0000-Pb92b-C2876Hb496-HK01ba-Vc222-Be8e3-R448a.deb' ./scripts/install_beaglebadge_vendor_edge_kernel_artifacts.sh
```

That `Pb92b` reinstall has completed on the BeagleBadge. Armbian marked `0007` as `needs_rebase`, but the build still completed and produced installable packages.

The current next build should carry an eighth `k3-6.12` diagnostic patch:

- `0008-Trace-OMAP-own-address-registers.patch`

That follow-up adds `OA`, `SA`, and `IE` register context to the existing slave diagnostics and adds a targeted `xfer-msg` trace when J6 starts a same-adapter master transfer to its own registered slave address.

The current next build after `Pdc8d` should also carry:

- `0009-Clear-stale-master-state-before-slave-listen.patch`

That follow-up clears master-only `I2C_CON` bits before returning the controller to slave listen mode, targeting the observed same-adapter divergence where self-write reaches `XRDY` but self-read does not.

For on-device install + reboot, use:

```sh
./scripts/install_latest_kernel_and_reboot.sh
```

Optional exact-build pinning still works through:

```sh
BADGESNAKE_BUILD_SUFFIX='<full package suffix>.deb' ./scripts/install_latest_kernel_and_reboot.sh
```

The wrapper:

- reuses `scripts/install_beaglebadge_vendor_edge_kernel_artifacts.sh`
- stops on install failure
- shows the currently running kernel before reboot
- runs `sync`
- triggers `systemctl reboot`

For repeated recovery after reboot, the repo also carries a boot-session service:

```sh
./scripts/install_badgesnake_boot_session_service.sh
```

That installs a oneshot systemd unit which:

- writes a boot summary to `artifacts/boot-status/latest.txt`
- creates a `tmux` session named `badgesnake` in `/root/interactor/badge-demo`
- on this board, attempts:

```sh
codex resume --last --dangerously-bypass-approvals-and-sandbox --no-alt-screen -C /root/interactor/badge-demo
```

- if Codex exits, falls back to a login shell inside the same `tmux` session
- after session creation, the helper forces stable tmux names:
  - session `badgesnake`
  - window `workspace`
- observed on the latest reboot:
  - `tmux attach -t badgesnake` works reliably
  - the pane can return to `bash` if `codex resume --last` exits

Control knob:

- set `BADGESNAKE_CODEX_BOOT_MODE=shell` before running the installer if you want the service to skip Codex auto-resume and open only a shell

Recovery / rollback:

- remove the service with:

```sh
./scripts/uninstall_badgesnake_boot_session_service.sh
```

## Current Installed-But-Not-Yet-Booted Kernel

The most recent completed reinstall on the live BeagleBadge is the
slave-TX-trace follow-up kernel built from the thirteen-patch `k3-6.12` stack.

- build summary:
  - `components/armbian-build/output/logs/summary-kernel-f2f0b16c-b1bc-4684-ac57-ea0ba75d1cd2.md`
- selected suffix:
  - `6.12.57-S22fb-D0000-P0e03-C2876Hb496-HK01ba-Vc222-Be8e3-R448a.deb`
- pinned reinstall command used:

```sh
BADGESNAKE_BUILD_SUFFIX='6.12.57-S22fb-D0000-P0e03-C2876Hb496-HK01ba-Vc222-Be8e3-R448a.deb' ./scripts/install_beaglebadge_vendor_edge_kernel_artifacts.sh
```

- `/var/log/dpkg.log` confirms completion through:
  - `2026-05-06 00:22:12 status installed linux-headers-vendor-edge-k3:arm64 26.02.0-trunk`
  - `2026-05-06 00:23:00 status installed linux-image-vendor-edge-k3:arm64 26.02.0-trunk`
- post-install overlay check still passes:
  - `/boot/dtb/ti/k3-am62l3-badge-qwiic-i2c.dtbo` exists
  - `/boot/uEnv.txt` still includes `ti/k3-am62l3-badge-qwiic-i2c.dtbo`

Next validation after reboot:

```sh
./scripts/validate_j7_to_j6_testunit_features.sh
i2ctransfer -f -y 3 w3@0x30 3 1 4 r5@0x30
i2ctransfer -f -y -b 3 w3@0x30 3 1 4 r5@0x30
dmesg | tail -n 120
```

That reboot has now happened. The `P21f5` fourteen-patch kernel boots as `#16`,
the repeated-start version query still works, and the proc-call response is
still `0x00 0x04 0x03 0x02 0x01` even though the new `AAS` priming code fires.

The current next debugging target is no longer “prime the first byte on `AAS`”.
It is the combined read-start condition seen on J6 when the proc-call read
phase begins.

The next useful follow-up should focus on combined slave startup state such as:

- `stat=0x614` (`AAS|ARDY|XUDF|XRDY`)

The build wrapper still enforces that `0014` is present before starting the
x86-host kernel build.

## Expected Install Notes

Two install-time messages are expected on the current board image and do not indicate failure:

- `update-initramfs: Symlink failed ... moving ...`:
  - `/boot` is FAT32 on this image, so Armbian cannot keep normal Linux symlinks there and falls back to `rename()`
- `_apt ... couldn't be accessed by user '_apt'`:
  - this warning appears when local `.deb` files are installed from a root-only path
  - the current reinstall wrapper stages the packages into `/var/tmp/badgesnake-kernel-debs/` first so future runs avoid that warning
- `W: Possible missing firmware ...` during `update-initramfs`:
  - these warnings were observed during the `P5507` reinstall and are not the blocker for current I2C slave-mode work

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
