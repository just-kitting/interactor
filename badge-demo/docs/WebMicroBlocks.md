# Web-Hosted MicroBlocks

This is the current preferred path for trying BadgeSnake student firmware on
BeagleBadge while native GP runtime support remains unresolved.

## What Works

- MicroBlocks web IDE build path through Emscripten
- Boardie browser simulator build
- browser-side `i2ctarget` simulation for Boardie
- local HTTPS hosting from the checked-out repository

## Build

```sh
./scripts/build_microblocks_web.sh
```

The build script emits `gp_wasm.js` directly instead of `gp_wasm.html`, so it
does not depend on Debian's missing `html-minifier-terser` packaging.
The Boardie build also avoids Closure Compiler so the hosted build does not
require a separate `closure-compiler` binary on Armbian.

That builds:

- `components/microblocks-smallvm/chromeApp/webapp/gp_wasm.js`
- `components/microblocks-smallvm/chromeApp/webapp/gp_wasm.wasm`
- `components/microblocks-smallvm/chromeApp/webapp/gp_wasm.data`
- `components/microblocks-smallvm/chromeApp/webapp/boardie/run_boardie.js`
- `components/microblocks-smallvm/chromeApp/webapp/boardie/run_boardie.wasm`

## Serve

```sh
./scripts/serve_microblocks_web.py
```

For remote access, the server now listens on `0.0.0.0` by default. If you want
the printed URL to use a specific hostname or IP, pass `--public-host`.

Default URL:

- `https://localhost:8443/microblocks.html`

The script generates a self-signed certificate under `.cache/microblocks-web/`
the first time it runs.

## Boardie I2C Target Workflow

Open the IDE, then open Boardie from the MicroBlocks runtime connection menu.
Import [I2C Target.ubl](/root/interactor/badge-demo/examples/microblocks/I2C%20Target.ubl#L1)
and start the simulated target at the Zepto address you want to emulate.

The hosted page exposes a browser-console helper:

```js
await BadgeSnakeBoardie.open()
await BadgeSnakeBoardie.enqueueWrite(0x42, [0x01, 0x02, 0x03])
await BadgeSnakeBoardie.transaction(0x42, [0x10, 0x20])
```

`transaction()` performs an optional controller write phase, then a controller
read phase, and waits for the Boardie-side reply.

The hosted server also exposes an HTTP bridge so Linux-side tests can behave
like an I2C controller:

```sh
python3 ./scripts/web_i2c_transaction.py \
  --url https://127.0.0.1:8443 \
  --address 0x42 \
  --write 0x10 0x20 \
  --insecure
```

That `8443` endpoint is the same running web server that is already serving the
IDE. The CLI transaction does not require a second listener; it should point at
the port where `microblocks.html` is already being served.

That request will block until the running Boardie/MicroBlocks program answers
the corresponding target read. The bridge smoke test is:

```sh
./scripts/test_web_i2c_bridge.sh
```

A concrete first program to run on Boardie is documented in
[I2C Target Echo Example.md](/root/interactor/badge-demo/examples/microblocks/I2C%20Target%20Echo%20Example.md#L1).

## Notes

- The browser path is the preferred short-term route on BeagleBadge.
- The bundled `gp-raspberryPi` binary is Raspberry Pi-specific and is not a
  viable runtime on BeagleBadge.
- Boardie is the first browser backend for BadgeSnake simulation; it is not a
  substitute for live Zepto firmware, but it is good enough to start writing
  MicroBlocks snake logic.
