# MicroBlocks Examples

This directory will hold student-facing BadgeSnake examples for BeagleConnect Zepto.

Current contents:

- `I2C Target.ubl`: library wrappers around the MicroBlocks `i2ctarget`
  primitive set
- `I2C Target Echo Example.md`: first hosted Boardie test program and matching
  Linux-side transaction

Planned contents:

- minimal info endpoint example
- move-only example
- full reference snake
- intentionally broken example for failure testing
- notes on how host-side flashing should load the program onto MSPM0L1117-based Zepto boards

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
