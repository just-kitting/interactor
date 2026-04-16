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
* upstream I2C destinations were discovered by enumerating `/dev/i2c-*`
* the CLI `Zepto` target accepts an image path and an I2C device path such as `/dev/i2c-0`

Local patch direction:

* destination discovery should actively probe each `/dev/i2c-*` bus with the MSPM0 BSL connection packet
* only buses that ACK as a live BSL target should be reported by `list-destinations zepto`
* the Zepto CLI path must pass its `no_filter` flag through consistently so filtered discovery is actually used by default

## Current board findings

On this BeagleBadge:

* `rustup` is installed
* default Rust toolchain is now configured as `stable-aarch64-unknown-linux-gnu`
* read-based scans on `/dev/i2c-0` and `/dev/i2c-2` did not show an ACK at `0x48` during this pass
* the next best probe is an active BSL connection request, not a passive bus scan, because the ROM BSL can sleep or time out if no interface is detected quickly enough

## Current local patch state

The local tree now includes:

* active I2C probing in `bb-flasher-mspm0/src/i2c.rs` using the MSPM0 BSL connection request and ACK byte
* a Zepto CLI fix in `bb-imager-cli/src/main.rs` so `list-destinations zepto` no longer inverts the `no_filter` flag
* a local diagnostic patch in `bb-flasher-mspm0/src/lib.rs` and `bb-flasher-mspm0/src/bsl.rs` so unknown MSPM0 BSL status bytes can be surfaced as real values instead of collapsing to a generic unknown error

The remaining validation step is to rebuild `bb-imager-cli`, rerun the live Zepto flash path, and capture the actual BSL status/error returned by the target.
