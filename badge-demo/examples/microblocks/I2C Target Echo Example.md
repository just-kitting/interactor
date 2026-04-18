# I2C Target Echo Example

Use this first on Boardie to confirm that the hosted I2C target bridge is
working end to end.

## Goal

- target address: `0x42`
- controller write: any byte array
- controller read reply: the same bytes, prefixed with `0xE0`

Example:

- controller writes: `0x10 0x20`
- MicroBlocks replies: `0xE0 0x10 0x20`

## Setup

1. Start the hosted IDE with:

```sh
./scripts/serve_microblocks_web.py --public-host <badge-ip-or-hostname>
```

2. Open `https://<badge-ip-or-hostname>:8443/microblocks.html`.
3. Open Boardie from the MicroBlocks runtime menu.
4. Import [I2C Target.ubl](/root/interactor/badge-demo/examples/microblocks/I2C%20Target.ubl#L1).

## Blocks To Build

Create one global variable:

- `lastWrite`

Initialize it to an empty byte array at startup.

### `when started`

Build this script:

1. `set lastWrite to (new byte array 0)`
2. `start i2c target at address 66`
3. `forever`
4. `if (i2c write pending?)`
5. `set lastWrite to (receive i2c write)`
6. `if (i2c read requested?)`
7. `set response to (join byte arrays (new byte array 1) lastWrite)`
8. `set byte 1 of response to 224`
9. `reply with response`
10. `wait 10 ms`

If your MicroBlocks build does not expose a direct `join byte arrays` block,
use this equivalent pattern:

1. `set response to (new byte array ((byte array length of lastWrite) + 1))`
2. `set byte 1 of response to 224`
3. copy each byte from `lastWrite` into `response` starting at index `2`
4. `reply with response`

## Linux-Side Test

Run this against the same server that is already hosting the IDE:

```sh
python3 ./scripts/web_i2c_transaction.py \
  --url https://127.0.0.1:8443 \
  --address 0x42 \
  --write 0x10 0x20 \
  --insecure
```

Expected JSON result:

```json
{
  "id": 1,
  "ok": true,
  "response": [224, 16, 32]
}
```

## Port Note

`8443` is the port of the running hosted IDE server. The CLI test should talk to
that same server. It does not need a second web server on another port.

If your IDE is running on a different port, pass the matching base URL:

```sh
python3 ./scripts/web_i2c_transaction.py \
  --url https://127.0.0.1:18443 \
  --address 0x42 \
  --write 0x10 0x20 \
  --insecure
```
