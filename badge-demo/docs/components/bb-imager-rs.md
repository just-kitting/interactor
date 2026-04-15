# BeagleBoard Imaging Utility

This submodule contains the code for flashing BeagleConnect Zepto and downloading/flashing the latest BeagleBadge images to a microSD card. This component is the current source tree for the preferred host-side flashing utility.

Upstream: https://github.com/beagleboard/bb-imager-rs
Our fork: https://github.com/embedconsult/bb-imager-rs
Local path: `components/bb-imager-rs`

Modifications to improve `bb-imager-rs` should be performed directly to this repository suitable for all `bb-imager-rs` users.

Commit locally and the maintainer will provide feedback.

BadgeSnake will use this submodule repo to understand and perform flashing of BeagleConnect Zepto.

## Relevant findings

The tree already contains MSPM0 support for Zepto-style targets:

* CLI feature `zepto_i2c`
* CLI feature `zepto_uart`
* `bb-flasher-mspm0` library with BSL helpers
* Linux I2C implementation under `bb-flasher-mspm0/src/i2c.rs`

## Important implementation details

Current I2C flashing assumptions in the source:

* BSL target address is hard-coded to `0x48`
* I2C destinations are discovered by enumerating `/dev/i2c-*`
* the CLI `Zepto` target accepts an image path and an I2C device path such as `/dev/i2c-0`

## Current board findings

On this BeagleBadge:

* `rustup` is installed
* default Rust toolchain is now configured as `stable-aarch64-unknown-linux-gnu`
* read-based scans on `/dev/i2c-0` and `/dev/i2c-2` did not show an ACK at `0x48` during this pass

## Next actions

* configure a Rust toolchain for local `bb-imager-rs` builds
* build `bb-imager-cli` with `zepto_i2c`
* verify whether the connected Zepto BSL device should appear at `0x48` on `i2c-0` or `i2c-2`
* once the device is visible, use `bb-imager-cli` to enumerate or flash it
