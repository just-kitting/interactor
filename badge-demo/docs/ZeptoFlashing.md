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

Because of that, the flash wrapper now waits for the BSL to become reachable before starting `bb-imager-cli`, and it can retry the flash command several times in the same invocation.

The probe now treats only `0x00` as a valid BSL ACK.
Other returned bytes such as `0x02` or `0x06` are reported as unexpected responses rather than successful readiness.

## Current flashing stance

The wrapper-level wait/retry behavior is kept in this repo.

The deeper MSPM0 transport changes inside `bb-imager-rs` were backed out so the actual flash sequence again follows the previously debugged upstream path.

Flash a Zepto image:

```sh
components/bb-imager-rs/target/debug/bb-imager-cli flash zepto <IMAGE> /dev/i2c-1
```

The flasher accepts `bin`, `hex`, `txt`, and `xz` inputs through the MSPM0 flasher path.
