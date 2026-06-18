# Zepto Flashing

## Current live path

The current intended host-side path is:

- `bb-imager-rs`
- `bb-imager-cli`
- feature set: `zepto_i2c`
- destination device: `/dev/i2c-1` for the Zepto currently attached on J6

## Board state

Validated on this board:

- J6 is exposed as `/dev/i2c-1`
- J7 is exposed as `/dev/i2c-3`
- the Zepto in MSPM0 BSL mode ACKs at `0x48` on `/dev/i2c-1`

## Repo wrappers

- `scripts/build_bb_imager_cli_zepto_i2c.sh`
- `scripts/list_zepto_i2c_destinations.sh`
- `scripts/probe_zepto_bsl_active.sh`
- `scripts/flash_zepto_bsl.sh`
- `scripts/flash_zepto_blink.sh`

## Automated Control Direction

The BeagleBadge Grove connector is now wired into the attached Zepto control signals:

- Grove pin 1 (yellow) -> Zepto `BSL`
- Grove pin 2 (white) -> Zepto `RST`

See [ZeptoControlWiring.md](/root/interactor/badge-demo/docs/ZeptoControlWiring.md#L1).

This should allow the host to automate Zepto BSL entry once the Linux GPIO line mapping and pinmux state for those Grove pins are identified.

## Expected CLI flow

Build:

```sh
scripts/build_bb_imager_cli_zepto_i2c.sh
```

List destinations:

```sh
scripts/list_zepto_i2c_destinations.sh
```

Current local patch direction:

- `bb-flasher-mspm0` now probes each `/dev/i2c-*` bus with the MSPM0 BSL connection request and keeps only ACKing buses
- `bb-imager-cli` Zepto listing was also fixed so it does not invert `--no-filter` and accidentally bypass that probe logic
- the remaining verification step is to rerun destination listing after the rebuilt CLI binary finishes linking on-device

## Timing note

Live probing on this board shows the Zepto BSL is not always reachable immediately after the manual BOOT/RST sequence.

- one successful user-run probe only received the first ACK on attempt `13/50`
- that delay was caused by the user invoking BSL while the probe loop was already running, not by a demonstrated intrinsic BSL startup delay

The probe now treats only `0x00` as a valid BSL ACK.
Other returned bytes such as `0x02` or `0x06` are reported as unexpected responses rather than successful readiness.

## Current flashing stance

The wrapper-level wait/retry behavior is kept in this repo.

The deeper MSPM0 transport changes inside `bb-imager-rs` were backed out so the actual flash sequence again follows the previously debugged upstream path.

## Current live limitation

Live testing on 2026-06-18 found:

- host-driven BSL entry now works when Grove signal pads are put in a non-`TX_DIS` GPIO-capable state
- the MSPM0 I2C flasher path in `bb-flasher-mspm0` works once `standalone_verification` is skipped on I2C
- a separate active BSL probe immediately before flashing can consume the first response packet and make the next `get_device_info` request fail
- the wrapper therefore must not actively probe the Zepto right before invoking the Rust flasher
- linking the full `bb-imager-cli` binary on-device can still hit the BeagleBadge memory limit, so lower-level validation was done with `bb-flasher-mspm0/examples/flash_i2c.rs`

Until the full `bb-imager-cli` path is rebuilt and revalidated on-device, this repo still includes a direct raw-I2C fallback flasher:

```sh
scripts/flash_zepto_bsl_direct_i2c.py <IMAGE.bin>
```

That helper:

- patches `PADCONFIG42/43` live with `busybox devmem`
- drives Grove `BSL` / `RST` through `gpioset`
- enters MSPM0 BSL over `/dev/i2c-1`
- performs `connect`, `get_device_info`, `unlock`, `mass_erase`, `program_data`, and `start_application`
- restores the original pad configuration on exit

Flash a Zepto image:

```sh
components/bb-imager-rs/target/debug/bb-imager-cli flash zepto <IMAGE> /dev/i2c-1
```

The flasher accepts `bin`, `hex`, `txt`, and `xz` inputs through the MSPM0 flasher path.
