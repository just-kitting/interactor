# Hot-Plug Policy

This document defines first-pass behavior when Zepto players are connected, missing, or disconnected during a BadgeSnake session.

## Scope

Applies to the direct QWIIC/I2C player connections planned for the first demo.

## Before Game Start

- The host should probe expected player slots before starting a game.
- A slot counts as available only if the player responds successfully to the logical metadata request.
- Missing slots should be reported clearly in logs and in the local UI.
- A game should not auto-start until the required number of players are present.

## During Game Start

- If a player fails to answer `/start`, the host should abort the game start and report the failure.
- Partial game start should not be treated as success.

## During Gameplay

- If a player fails to answer `/move` within the configured timeout, that player should be treated as failed for the turn.
- First-pass failure behavior:
  - log the timeout or transport failure
  - mark the player as eliminated
  - continue the game with remaining players if the rules engine permits
- A disconnect mid-game should not crash the host runtime.

## Reconnect Behavior

- A player that disconnects during an active game should not be reinserted into that same game automatically.
- Reconnected players should only be eligible for the next game session.

## UI Expectations

- The local display should indicate:
  - detected players before game start
  - missing player slots
  - mid-game disconnections or timeouts
- Status outputs should make transport failure visible even if the ePaper refresh is delayed.

## Follow-Up

- refine player-slot discovery once the I2C addressing scheme is implemented
- decide whether repeated move failures should produce a distinct elimination cause in logs or UI
