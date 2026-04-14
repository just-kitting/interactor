# Display And Status Plan

This document defines what BadgeSnake should show on the BeagleBadge outputs in the first pass.

## ePaper

The ePaper display is the primary output.

It should show:

- board state
- snake positions
- food positions
- current turn number
- active player names
- final game result

During pre-game or fault states it should show:

- detected player slots
- missing player slots
- transport failure summary if a game cannot start

## RGB LED

The RGB LED should provide coarse state at a glance:

- idle / waiting for players
- ready to start
- active game
- player failure or transport error
- completed game

Exact colors can be chosen later once the LED control path is identified cleanly.

## Buzzer

The buzzer should be sparse and purposeful:

- one short cue on game start
- one short cue on player failure or disconnect
- one short cue on game end

Avoid continuous or distracting sounds in classroom use.

## 7-Segment Displays

First-pass target use:

- show player count before game start
- show current turn modulo a small range during gameplay
- show a simple end-state indicator when the game completes

Implementation direction should follow the kernel-module style demonstrated in:

- `components/vsx-examples/PocketBeagle-2/seven_segment`

Avoid spending more time on generic userspace exploration unless that example path proves unsuitable.

## Buttons

Front-panel buttons are already visible to Linux and can support:

- start or confirm action
- menu or back action
- optional local navigation between idle/status screens

Buttons should not become the primary gameplay control path; they are for host operation only.

## First-Pass Priorities

1. ePaper board rendering
2. failure visibility on the ePaper
3. a minimal RGB LED state machine
4. sparse buzzer notifications
5. optional 7-segment integration
