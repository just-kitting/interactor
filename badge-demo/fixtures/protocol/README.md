# Protocol Fixtures

These fixtures support early BadgeSnake protocol and simulation work before the host runtime is implemented.

Contents:

- `path_tokens.json`: token assignments for logical Battlesnake endpoints
- `info-response.json`: example body for logical `GET /`
- `start-request.json`: example body for logical `POST /start`
- `move-request.json`: example body for logical `POST /move`
- `move-response.json`: example body for logical `POST /move` response
- `end-request.json`: example body for logical `POST /end`

These payloads are intentionally compact and deterministic so they can be reused in regression tests.
