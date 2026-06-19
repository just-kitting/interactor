# BadgeSnake Live Host Workflow

This is the current shortest path for running a Zepto-hosted snake inside the
badge-launcher Battlesnake demo.

## Scope

- 1 live Zepto player over I2C
- 1 simulated opponent on BeagleBadge
- launcher UI rendered from the existing `/tmp/badgesnake/*.json` contract
- manual author/load step for Zepto firmware

This does not yet include host-driven flashing or a MicroBlocks-to-host runtime
bridge.

## Expected Zepto Behavior

The Zepto program should answer the logical Battlesnake endpoints described in
[Protocol.md](/root/interactor/badge-demo/docs/Protocol.md#L1):

- `GET /`
- `POST /start`
- `POST /move`
- `POST /end`

over the framed I2C transport at the Zepto gameplay address, currently
`0x42` on BeagleBadge `i2c-1`.

## Authoring Path

1. Use the hosted MicroBlocks workflow from
   [WebMicroBlocks.md](/root/interactor/badge-demo/docs/WebMicroBlocks.md#L1)
   to validate request parsing and response framing first.
2. Load the same logic onto the physical Zepto.
3. Keep the Zepto answering the gameplay address expected by the launcher
   backend, unless you also override the backend environment variables.

## Launcher Demo

The launcher backend now defaults to `ui-live`.

Relevant environment variables:

- `BADGESNAKE_BACKEND_MODE=ui-live`
- `BADGESNAKE_I2C_BUS=1`
- `BADGESNAKE_I2C_ADDR=0x42`
- `BADGESNAKE_TIMEOUT_MS=500`
- `BADGESNAKE_OPPONENT_MODE=clockwise`

If the Zepto is not answering metadata yet, the UI stays in a waiting state and
continues probing until the device responds.

## Fallback

To force the older simulator backend instead of the live Zepto path:

```sh
BADGESNAKE_BACKEND_MODE=ui-sim
```
