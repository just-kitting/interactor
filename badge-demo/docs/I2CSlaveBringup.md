# I2C Slave Bring-Up

This document captures the first concrete bring-up path for AM62L target/slave
mode on BeagleBadge using the kernel's existing `slave-testunit` backend.

## Goal

Before building the BadgeSnake bridge driver, prove three things:

1. the AM62L host adapter can register a slave address
2. the Linux I2C slave event path works on J6/J7
3. multi-controller behavior can be exercised with the stock kernel test backend

## Required Kernel Config

The Armbian config should include:

- `CONFIG_I2C_SLAVE=y`
- `CONFIG_I2C_SLAVE_EEPROM=m`
- `CONFIG_I2C_SLAVE_TESTUNIT=m`

These options are now staged in:

- `components/armbian-build/config/kernel/linux-k3-vendor-edge.config`
- `components/armbian-build/config/kernel/linux-k3-beagle-vendor.config`
- `components/armbian-build/config/kernel/linux-k3-beagle-vendor-rt.config`

## Important Limitation

Even with those configs enabled, target-mode bring-up still depends on
`drivers/i2c/busses/i2c-omap.c` learning how to:

- register a slave address
- unregister a slave address
- translate slave IRQs into `i2c_slave_event()` callbacks

Until that exists, `slave-testunit` cannot actually bind on AM62L.

The first BeagleBadge test patch for that gap is now carried through the Armbian patchset at:

- `components/armbian-build/patch/kernel/archive/k3-6.12/0001-Stage-OMAP-I2C-slave-registration-support.patch`
- `components/armbian-build/patch/kernel/archive/k3-6.12/0002-Fix-OMAP-slave-helper-declaration-order.patch`

## Build Scope

There are two different iteration paths here:

- `i2c-slave-testunit` and `i2c-slave-eeprom` are loadable modules
- `i2c-omap` is built into the K3 kernel image

That means:

- for validating the stock slave backends, a local module-only build is a reasonable faster path
- for the eventual AM62L slave-mode enablement in `i2c-omap.c`, a full kernel image rebuild and reboot are still required

So the current corrective rebuild is only needed because the returned kernel packages were missing the module config entirely. After that, module-only iteration should be preferred until work moves into `i2c-omap.c`.

## First Validation Target

Use the kernel's own backend:

- source: `drivers/i2c/i2c-slave-testunit.c`
- docs: `Documentation/i2c/slave-testunit-backend.rst`

This is the right first target because it already exercises:

- multi-controller reads
- repeated-start handling
- SMBus host notify
- SMBus alert behavior

## Repo Helper

Use:

```sh
./scripts/bringup_i2c_slave_testunit.sh start 1 0x30
```

That helper:

- loads `i2c-slave-testunit` if present
- instantiates `slave-testunit` at `0x30` on the selected bus
- uses the Linux slave-address sysfs convention of `0x1000 + addr`

To remove it:

```sh
./scripts/bringup_i2c_slave_testunit.sh stop 1 0x30
```

To check status:

```sh
./scripts/bringup_i2c_slave_testunit.sh status 1 0x30
```

## Current Build Path

To build the updated BeagleBadge `vendor-edge` kernel packages on an x86 host with Docker:

```sh
./scripts/build_beaglebadge_vendor_edge_kernel_x86_docker.sh
```

This is intentionally a kernel-package build only. It is the next step for target-mode validation and does not imply that a replacement microSD image should be produced yet.

## Current Install Path

After copying the returned `.deb` artifacts back into `components/armbian-build/output/debs/` on the BeagleBadge, reinstall them with:

```sh
./scripts/install_beaglebadge_vendor_edge_kernel_artifacts.sh
```

This uses `apt-get install --reinstall` because the rebuilt packages currently carry the same Debian version string as the installed `vendor-edge-k3` packages.

## Suggested Bring-Up Sequence

1. Install the returned kernel artifacts and reboot into that kernel.
2. Confirm the module exists:

```sh
modinfo i2c-slave-testunit
```

3. Instantiate on J6 / `/dev/i2c-1`:

```sh
./scripts/bringup_i2c_slave_testunit.sh start 1 0x30
```

4. Run a simple read test:

```sh
i2ctransfer -y 1 r1@0x30
```

Expected initial value:

- `0x00` when the testunit is idle

5. Run a repeated-start version query test:

```sh
i2ctransfer -y -b 1 w3@0x30 4 0 0 r128
```

6. Run a multi-controller read test:

```sh
i2cset -y 1 0x30 1 0x50 0x10 1 i
```

This should force the slave testunit to act as a controller and read from
another address on the bus after a short delay.

## Success Criteria

Target-mode bring-up is considered real when:

- `slave-testunit` can be instantiated on AM62L
- simple read traffic works
- repeated-start test works
- at least one multi-controller command works

Only after that should BadgeSnake-specific bridge code be added.

## Current AM62L Result

The current BeagleBadge state has now moved past the old adapter-registration failure:

- `modinfo i2c-slave-testunit` works
- the QWIIC overlay restores `i2c-1` and `i2c-3`
- `new_device` can instantiate the slave address node on `i2c-1`
- `slave-testunit` now binds successfully on `i2c-1`

The old blocker:

```text
i2c_slave_register: not supported by adapter
```

is no longer present after booting the `P5507` kernel.

## Current Rebuild Input

The next x86-host rebuild should be treated as valid only if the build log shows kernel patching from `archive/k3-6.12`.

If the build log still shows zero applied kernel patches, the wrong Armbian checkout was used and the returned artifacts should not be trusted for AM62L slave-mode testing.

The first valid rebuilt artifact set is now:

- build UUID: `1b53c454-a989-4c66-a5e6-e8b706e4bc05`
- artifact suffix: `P5507`
- patching result: `2 total patches; 2 applied; 0 with problems`

That kernel has been installed on the BeagleBadge. The immediate next step is to reboot into it and repeat the `slave-testunit` binding test on `i2c-1`.

## Current Runtime Result

After rebooting into the `P5507` kernel:

- `uname -a` shows build `#4` from `Thu Apr 30 08:25:00 UTC 2026`
- `./scripts/bringup_i2c_slave_testunit.sh status 1 0x30` reports:
  - `present: 1-1030`
  - `bound: yes`
  - `slave-testunit`
- `dmesg` shows only:
  - `i2c i2c-1: new_device: Instantiated device slave-testunit at 0x30`

The next failure has moved further forward:

- forced same-adapter controller traffic such as `i2ctransfer -f -y 1 r1@0x30` now times out
- there is no longer an immediate `-95` adapter rejection at slave registration time

So the remaining work is no longer “make the slave backend bind.”
It is “make bound slave-mode traffic actually complete.”
