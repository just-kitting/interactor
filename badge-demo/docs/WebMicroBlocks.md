# Web-Hosted MicroBlocks

This is the current preferred path for trying BadgeSnake student firmware on
BeagleBadge while native GP runtime support remains unresolved.

## What Works

- MicroBlocks web IDE build path through Emscripten
- Boardie browser simulator build
- browser-side BadgeSnake `i2ctarget` simulation for Boardie
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
Import [BadgeSnake I2C Target Sim.ubl](/root/interactor/badge-demo/examples/microblocks/BadgeSnake%20I2C%20Target%20Sim.ubl#L1)
and start the simulated target at the Zepto address you want to emulate.

The hosted page exposes a browser-console helper:

```js
await BadgeSnakeBoardie.open()
await BadgeSnakeBoardie.enqueueWrite(0x42, [0x01, 0x02, 0x03])
await BadgeSnakeBoardie.transaction(0x42, [0x10, 0x20])
```

`transaction()` performs an optional controller write phase, then a controller
read phase, and waits for the Boardie-side reply.

## Notes

- The browser path is the preferred short-term route on BeagleBadge.
- The bundled `gp-raspberryPi` binary is Raspberry Pi-specific and is not a
  viable runtime on BeagleBadge.
- Boardie is the first browser backend for BadgeSnake simulation; it is not a
  substitute for live Zepto firmware, but it is good enough to start writing
  MicroBlocks snake logic.
