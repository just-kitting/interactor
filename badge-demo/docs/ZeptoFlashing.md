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

## Expected CLI flow

Build:

```sh
scripts/build_bb_imager_cli_zepto_i2c.sh
```

List destinations:

```sh
scripts/list_zepto_i2c_destinations.sh
```

Current limitation:

- upstream `bb-imager-cli` lists every `/dev/i2c-*` controller for the Zepto target
- a local `bb-imager-rs` patch is being applied so enumeration filters by a real MSPM0 BSL ACK instead

Flash a Zepto image:

```sh
components/bb-imager-rs/target/debug/bb-imager-cli flash zepto <IMAGE> /dev/i2c-1
```

The flasher accepts `bin`, `hex`, `txt`, and `xz` inputs through the MSPM0 flasher path.
