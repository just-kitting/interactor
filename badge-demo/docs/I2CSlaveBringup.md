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

## Suggested Bring-Up Sequence

1. Boot a kernel that includes the above config options.
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
