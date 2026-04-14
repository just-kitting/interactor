# Rules CLI Transport Notes

## Current Scope

`components/battlesnake-rules` now carries the first BadgeSnake-oriented transport extensions in the CLI layer:

- `http://` and `https://` continue to use the upstream HTTP client path
- `sim://` creates deterministic in-process players without a network server
- `i2c://` preserves the same request lifecycle while standing in for BadgeSnake bus transport

These changes are intentionally limited to the CLI request layer so the core rules engine remains close to upstream.

## Simulated Player URLs

Example:

```sh
go run ./cli/battlesnake play \
  --name Clocky --url 'sim://clockwise?name=Clocky&color=%23112233' \
  --name Righty --url 'sim://right?name=Righty&color=%2300ff00'
```

Supported move modes:

- `up`
- `down`
- `left`
- `right`
- `clockwise`
- `counterclockwise`

Optional query parameters:

- `name`
- `color`
- `move`
- `latency_ms`

## I2C-style URLs

Example:

```sh
go run ./cli/battlesnake play \
  --name ZeptoA --url 'i2c://stub?addr=0x10&move=clockwise&name=ZeptoA' \
  --name ZeptoB --url 'i2c://stub?addr=0x11&move=counterclockwise&name=ZeptoB'
```

The current `i2c://` implementation is still host-side simulation. It maps Battlesnake operations onto the intended BadgeSnake transport tokens:

- `GET /` -> `0x01`
- `POST /start` -> `0x02`
- `POST /move` -> `0x03`
- `POST /end` -> `0x04`

This provides a stable bridge between the upstream CLI and the BadgeSnake protocol work in this repo before real I2C bus plumbing is attached.

## Current Board Status

This Armbian image has `i2c-tools` installed, but `modinfo i2c-stub` currently fails because the `i2c-stub` kernel module is not present in the running kernel package.

That means:

- `sim://` is the primary executable test path on this board today
- `i2c://` is useful for protocol-shape testing inside the `rules` CLI
- true `i2c-stub` adapter emulation still depends on kernel/module availability

Use `scripts/check_i2c_stub_support.sh` to confirm local readiness before depending on `i2c-stub` in test plans.
