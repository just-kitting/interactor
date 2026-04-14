# Simulation Plan

## Purpose

Build a host-only BadgeSnake simulation path before hardware integration so protocol, game-loop, and failure behavior can be tested repeatably on a development machine.

## First Simulation Target

A CLI-hosted simulation should:

- use the upstream `components/battlesnake-rules` engine
- emulate Zepto players in-process
- preserve the BadgeSnake path-token and JSON-body protocol model
- support deterministic seeds for reproducible games
- capture turn-by-turn logs suitable for protocol and regression tests

## Planned Runtime Shape

- `cmd/badgesnake`: CLI entry point
- `internal/protocol`: path tokens, frame structures, encode/decode helpers
- `internal/simtransport`: in-memory transport used in place of I2C
- `internal/player`: player abstractions for mock or fixture-driven snakes
- `internal/game`: host-side orchestration around the rules engine

## First Milestones

1. Implement a protocol encoder/decoder that turns logical endpoint calls into BadgeSnake frames.
2. Implement an in-memory transport that can send those frames to mock players.
3. Create a fixture-driven mock player that returns canned metadata, start, move, and end responses.
4. Run one deterministic two-player game through the Battlesnake rules engine.
5. Save logs/fixtures for regression tests.

## Minimum Success Criteria

- A development host can run a complete game without BeagleBadge hardware.
- The host can exercise the same logical player API that the I2C path will use later.
- A failing or slow mock player can be simulated.
- The resulting logs can be checked into the repo as fixtures.

## Deferred Until After First Simulation

- Real I2C bus access
- ePaper rendering
- RGB LED/buzzer integration
- Zepto flashing
