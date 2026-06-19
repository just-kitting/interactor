# MicroBlocks Examples

This directory will hold student-facing BadgeSnake examples for BeagleConnect Zepto.

Current contents:

- `I2C Target.ubl`: library wrappers around the MicroBlocks `i2ctarget`
  primitive set
- `I2C Target Echo Example.md`: first hosted Boardie test program and matching
  Linux-side transaction
- `BadgeSnake Live Host Workflow.md`: current 1-live-Zepto plus 1-simulated-opponent path for the launcher demo

Planned contents:

- minimal info endpoint example
- move-only example
- full reference snake
- intentionally broken example for failure testing

## Simulation Workflow

The first simulator path is Linux-host only. It is intended for rapid protocol and
student-firmware development before live Zepto flashing is stable.

1. Build/run the Linux MicroBlocks VM from `components/microblocks-smallvm/linux+pi`.
2. Import `I2C Target.ubl` into the MicroBlocks IDE.
3. Start the simulated target with the Zepto address you want to emulate.
4. Use `scripts/microblocks_i2c_sim.py transaction ...` from this repo to enqueue
   an optional controller write followed by a controller read and wait for the
   student program's reply.

For the hosted web IDE, use `scripts/web_i2c_transaction.py` to drive controller
transactions through the Boardie bridge while the browser UI is open.

## Current Live Demo Boundary

The current badge demo path is:

1. author and test the protocol logic with the hosted MicroBlocks workflow
2. manually load the program onto a Zepto
3. run the badge-launcher Battlesnake app with the default `ui-live` backend
4. let the live Zepto snake play against one local simulated opponent

That gives us a usable end-to-end demo without depending on target-role I2C on
BeagleBadge or on a launcher-integrated flashing flow.
