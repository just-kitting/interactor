# MicroBlocks SmallVM

This submodule contains the source for the MicroBlocks interpreter and development environment that should be updated to support BeagleConnect Zepto and allow communication with the development environment over I2C, rather than UART.

Upstream: https://bitbucket.org/john_maloney/smallvm
Our fork: https://github.com/just-kitting/microblocks-smallvm

Modifications to add hardware support for I2C transport of development environment communications should be performed directly to this repository suitable for upstreaming.

Commit locally and the maintainer will provide feedback.

BadgeSnake will use this submodule repo to provide support for MicroBlocks graphical programming to the user on BeagleBadge and pushing data to BeagleConnect Zepto for immediate commands and executable opcodes.

## Local patches in use

- Added a generic `i2ctarget` primitive set for the Linux simulator so MicroBlocks
  code can act like an I2C target during host-side BadgeSnake simulation.
- Added the same primitive surface in Boardie with a browser-side queue backend.
- Added a hosted HTTP bridge so Linux-side controller requests can be sent into
  the Boardie web simulation without browser-console interaction.

## Runtime notes on BeagleBadge

- The bundled `gp-raspberryPi` executable is a 32-bit `armhf` binary, but it is
  not a generic ARM Linux build.
- It depends on Raspberry Pi userspace libraries including `libbcm_host.so`, so
  installing only `libc6:armhf` on BeagleBadge would not be sufficient to run it.
- The bundled `gp-linux64bit` binary is `x86-64` and cannot run on BeagleBadge.
- For BeagleBadge, the viable paths are:
  - build or obtain a native `gp` runtime for `aarch64`, or
  - use the web-hosted MicroBlocks app with the browser-side `i2ctarget`
    backend and optional HTTP bridge rather than relying only on the Linux spooler

## Update procedure

- Commit generic simulator/runtime changes inside the `components/microblocks-smallvm`
  submodule.
- Keep BadgeSnake-specific MicroBlocks libraries and helper scripts in the
  superproject under `examples/microblocks/`, `scripts/`, and `docs/`.
