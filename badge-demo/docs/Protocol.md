# BadgeSnake Protocol

## Goal

Keep `components/battlesnake-rules` as close to upstream as possible by translating Battlesnake HTTP semantics into a compact I2C message protocol.

## Endpoint Mapping

The Zepto player should behave as if it implements these logical Battlesnake endpoints:

- `GET /`
- `POST /start`
- `POST /move`
- `POST /end`

To reduce traffic, the host should replace long HTTP paths with compact path tokens.

Initial token table:

- `0x01`: `GET /`
- `0x02`: `POST /start`
- `0x03`: `POST /move`
- `0x04`: `POST /end`

## Transaction Model

### Metadata request

Equivalent to `GET /`.

1. Host writes a single-frame request containing token `0x01`
2. Player returns the JSON body normally returned by the Battlesnake info endpoint

### Start request

Equivalent to `POST /start`.

1. Host writes a frame containing token `0x02` and a JSON request body
2. Player performs any setup and may return an optional JSON response body

### Move request

Equivalent to `POST /move`.

1. Host writes a frame containing token `0x03` and a JSON request body
2. Host reads back the JSON response body containing the move decision

### End request

Equivalent to `POST /end`.

1. Host writes a frame containing token `0x04` and a JSON request body
2. Player performs cleanup and may return an optional JSON response body

## Frame Format

Initial host-to-player frame:

```text
byte 0  : protocol version
byte 1  : path token
byte 2  : flags
byte 3  : reserved
byte 4-5: payload length, little-endian
byte N  : payload bytes
```

Initial player-to-host frame:

```text
byte 0  : protocol version
byte 1  : status code
byte 2  : flags
byte 3  : reserved
byte 4-5: payload length, little-endian
byte N  : payload bytes
```

## Status Codes

- `0x00`: success
- `0x01`: bad request
- `0x02`: unsupported token
- `0x03`: internal error
- `0x04`: timeout or busy

## Notes

- JSON bodies should remain structurally compatible with Battlesnake expectations
- HTTP headers are intentionally removed
- The protocol should favor simple blocking transactions first
- If response sizes exceed practical I2C transfer limits, chunking can be added later without changing the logical endpoint model

## Fixtures

Machine-readable starting fixtures live under `fixtures/protocol/`.

These should be treated as the initial contract for:

- token assignments
- example request and response bodies
- early simulation and regression tests
