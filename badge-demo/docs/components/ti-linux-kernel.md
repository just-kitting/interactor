# TI Linux Kernel

This submodule contains the Texas Instruments Linux kernel source tree used as
the starting point for BeagleBadge kernel work.

Upstream: https://github.com/TexasInstruments/ti-linux-kernel

## Why It Exists

BadgeSnake needs this tree for:

- AM62L I2C controller and possible slave-mode work
- BeagleBadge DTS and driver inspection
- evaluating whether the host can support a multi-controller I2C transport
- eventual kernel patches for transport, simulation, or flashing support

## Current Relevance To BadgeSnake

Important files already identified:

- `drivers/i2c/busses/i2c-omap.c`
- `drivers/i2c/i2c-core-slave.c`
- `drivers/i2c/i2c-slave-testunit.c`
- `drivers/i2c/i2c-slave-eeprom.c`
- `Documentation/i2c/slave-interface.rst`
- `Documentation/i2c/slave-testunit-backend.rst`
- `Documentation/i2c/i2c-stub.rst`

Current conclusion:

- the tree has the generic I2C slave framework and useful example backends
- `i2c-omap.c` does not currently wire up `reg_slave` / `unreg_slave`
- so AM62L host slave support is a prerequisite for a kernel-visible BadgeSnake
  transport on BeagleBadge

## Planned BadgeSnake Use

- phase 1: validate AM62L slave-mode support using `slave-testunit`
- phase 2: add a BadgeSnake-specific slave backend for framed messages
- phase 3: connect the kernel ABI to the Go runtime and simulation tools

## Update Procedure

- keep generic kernel work in this submodule
- keep BadgeSnake planning, docs, and userspace integration in the superproject
- when changing kernel behavior, record:
  - config deltas
  - DTS changes
  - exact verification commands
  - rollback impact
