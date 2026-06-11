# Zephyr `native_sim` BadgeSnake I2C Transport Tests

This app provides a host-only validation lane for the current BadgeSnake I2C
transport contract.

## Scope

- runs as a Zephyr `native_sim` executable on a Linux host
- uses Zephyr's `zephyr,i2c-emul-controller` bus
- registers a custom I2C target that behaves like the current Zepto-side
  BadgeSnake transport endpoint
- executes `ztest` cases that perform controller write/read transactions against
  that target

## What It Validates

- request and response frame header encoding assumptions
- current token mapping for `info`, `start`, `move`, and `end`
- current fixed-role transaction model:
  - controller write for request
  - later controller read for response
- zero padding after the framed response prefix when the controller reads a
  larger fixed buffer

## What It Does Not Validate

- BeagleBadge Linux `/dev/i2c-*` behavior
- AM62L `i2c-omap` behavior
- MSPM0 hardware target-mode behavior
- electrical timing, pull-ups, or bus integrity

## Usage

Use the separate build and run wrappers:

```sh
./scripts/build_zepto_zephyr_native_sim_i2c_transport.sh
./scripts/run_zepto_zephyr_native_sim_i2c_transport.sh
```

Or use the combined helper:

```sh
./scripts/test_zepto_zephyr_native_sim_i2c_transport.sh
```

Requirements:

- `west`
- a Zephyr workspace (`ZEPHYR_BASE` set, or runnable through `west`)
- host build tools required by `native_sim`
