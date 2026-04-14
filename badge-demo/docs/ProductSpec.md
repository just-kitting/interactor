# BadgeSnake Product Spec

## Objective

BadgeSnake is a local BeagleBadge demonstration that lets attached BeagleConnect Zepto boards participate as Battlesnake controllers without requiring each player to host a networked web service.

The host system on BeagleBadge should preserve the Battlesnake request/response model closely enough that upstream `rules` remains the game authority and only transport adaptation is added.

## Initial Scope

- Host board: BeagleBadge
- Player boards: up to 2 BeagleConnect Zepto boards over QWIIC/I2C
- Student-facing programming environment: MicroBlocks
- Game engine: `components/battlesnake-rules`
- Transport model: HTTP semantics translated onto I2C transactions

## Initial Simulation Target

The first simulation target should be a host-only CLI that:

- runs on a development host without BeagleBadge hardware
- uses the upstream Battlesnake rules engine
- replaces real I2C Zepto transactions with an in-process transport shim
- allows canned or mock snake responses for repeatable testing

This should validate game-loop integration and protocol framing before real hardware is involved.

## Gameplay Defaults

Use Battlesnake `rules` defaults unless a later demo need forces divergence.

Initial assumptions:

- Ruleset: `standard`
- Map: `standard`
- Width: `11`
- Height: `11`
- Timeout: `500ms`
- Food spawn chance: `15`
- Minimum food: `1`

## Host Responsibilities

- Detect and register connected Zepto players
- Expose Battlesnake-compatible request semantics to those players over I2C
- Run the game through the upstream rules engine
- Render game state locally on BeagleBadge
- Surface basic status through non-ePaper outputs when useful
- Log enough state to reproduce failures and field issues

## Player Responsibilities

- Implement Battlesnake endpoints logically equivalent to `GET /`, `POST /start`, `POST /move`, and `POST /end`
- Respond within the host-enforced timeout
- Tolerate compact path-token encodings in place of full HTTP paths

## Non-Goals For First Pass

- Full online Battlesnake compatibility as a public web service
- More than 2 directly attached Zepto players
- Dynamic protocol negotiation
- Broad firmware-language support beyond MicroBlocks examples

## Immediate Follow-On

- Define how gameplay state maps onto the ePaper, RGB LED, buzzer, and 7-segment displays
- Build a local host-side simulation path that runs without hardware
- Create a minimal reference snake example under `examples/microblocks/`
