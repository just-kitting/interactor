# MicroBlocks Examples

This directory will hold student-facing BadgeSnake examples for BeagleConnect Zepto.

Current contents:

- `BadgeSnake I2C Target Sim.ubl`: library wrappers around the Linux MicroBlocks
  simulator's `i2ctarget` primitive set

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
2. Import `BadgeSnake I2C Target Sim.ubl` into the MicroBlocks IDE.
3. Start the simulated target with the Zepto address you want to emulate.
4. Use `scripts/microblocks_i2c_sim.py transaction ...` from this repo to enqueue
   a controller request and wait for the student program's reply.
