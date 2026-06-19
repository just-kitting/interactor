# `cmd/badgesnake`

BadgeSnake host binary.

Why Go:

- `components/battlesnake-rules` is already a Go module
- Keeping the host runtime in Go minimizes impedance with the upstream rules engine
- The host can eventually wrap or embed rules logic without introducing a second systems language

Go 1.24.4 is now installed on the live BeagleBadge image. Local development is possible, but first-run builds on the board are still slow enough that longer compiles should be treated as background work.

## Commands

- `tokens`: print the logical Battlesnake-to-I2C token table
- `frame-example`: emit one encoded sample move frame
- `ui-sim`: file-backed two-snake simulator for the launcher UI
- `ui-live`: file-backed live host that runs one Zepto over I2C against one local simulated opponent

`ui-live` keeps the launcher JSON contract used by
`components/badge-launcher/applications/games/battlesnake/` and is the current
path for the 1-live-player badge demo.
