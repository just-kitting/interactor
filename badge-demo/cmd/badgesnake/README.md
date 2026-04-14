# `cmd/badgesnake`

Planned home of the BadgeSnake host binary.

Why Go:

- `components/battlesnake-rules` is already a Go module
- Keeping the host runtime in Go minimizes impedance with the upstream rules engine
- The host can eventually wrap or embed rules logic without introducing a second systems language

The live BeagleBadge image does not currently have the Go toolchain installed, so local development on the device should not be assumed. Prefer CI or a larger external build host for compilation until the image strategy changes.
