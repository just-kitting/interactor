# `cmd/badgesnake`

Planned home of the BadgeSnake host binary.

Why Go:

- `components/battlesnake-rules` is already a Go module
- Keeping the host runtime in Go minimizes impedance with the upstream rules engine
- The host can eventually wrap or embed rules logic without introducing a second systems language

Go 1.24.4 is now installed on the live BeagleBadge image. Local development is possible, but first-run builds on the board are still slow enough that longer compiles should be treated as background work.
