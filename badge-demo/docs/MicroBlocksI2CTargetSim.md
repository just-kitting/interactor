# MicroBlocks I2C Target Simulation

This document defines the first BadgeSnake simulation path for MicroBlocks while
live Zepto flashing is still unstable.

## Scope

- Linux MicroBlocks VM only for the first functional backend
- Boardie keeps the primitive names reserved but does not implement the spooler yet
- Simulation models the Zepto as an I2C target that receives a controller write
  payload and may later see a controller read request
- The simulation now follows Zephyr's target-style split between write-receive
  and read-request handling rather than advertising a controller read length up front

## Primitive Surface

The `components/microblocks-smallvm` Linux VM now exposes a new primitive set:

- `[i2ctarget:start]` address
- `[i2ctarget:stop]`
- `[i2ctarget:isStarted]`
- `[i2ctarget:address]`
- `[i2ctarget:writeAvailable]`
- `[i2ctarget:receiveWrite]`
- `[i2ctarget:readRequested]`
- `[i2ctarget:reply]`

`[i2ctarget:receiveWrite]` returns a byte array containing the controller's
write payload. `[i2ctarget:readRequested]` becomes true when the controller
initiates a read phase. `[i2ctarget:reply]` publishes the response bytes for
that read request and completes that phase.

## Host Spool Format

The Linux VM uses a spool directory, defaulting to
`/tmp/microblocks_i2c_target_sim`. Override it with `MICROBLOCKS_I2C_SIM_DIR`.

Write request files:

- `write-<address>-<request_id>.bin`
- body: controller write payload

Read request files:

- `read-<address>-<request_id>.bin`
- body: empty marker file indicating that the controller has started a read phase

Response files:

- `response-<address>-<request_id>.bin`
- body: target reply payload

This format is intentionally simple so repo-local tools can drive it without
depending on the MicroBlocks IDE protocol.

## Repo Helpers

- [microblocks_i2c_sim.py](/root/interactor/badge-demo/scripts/microblocks_i2c_sim.py): enqueue requests and wait for replies
- [test_microblocks_i2c_sim.sh](/root/interactor/badge-demo/scripts/test_microblocks_i2c_sim.sh): smoke test for the helper and spool format
- [BadgeSnake I2C Target Sim.ubl](/root/interactor/badge-demo/examples/microblocks/BadgeSnake%20I2C%20Target%20Sim.ubl): importable MicroBlocks library for these primitives

## Intended Usage

1. Build and run the Linux MicroBlocks VM.
2. Import `BadgeSnake I2C Target Sim.ubl`.
3. Start the simulated target at the intended Zepto address.
4. Poll `badge i2c write pending?` and `badge i2c read requested?` from student code.
5. Use `badge i2c receive write` to consume any controller write payload.
6. When a read is requested, publish bytes with `badge i2c reply _`.
7. Drive transactions from the host with `scripts/microblocks_i2c_sim.py transaction ...`.

## Web And IDE Hosting

The current functional backend is Linux VM only, because the spooler uses the
local filesystem. Boardie already builds to WebAssembly/JavaScript, but its
`i2ctarget` backend is stubbed today. That means:

- the MicroBlocks web app can still be hosted from a web server
- the new I2C target simulation path does not yet work in that web build
- a browser-capable backend should use JavaScript messaging or WebSockets rather
  than the filesystem spooler
