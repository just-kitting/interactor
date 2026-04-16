# MicroBlocks I2C Target Simulation

This document defines the first BadgeSnake simulation path for MicroBlocks while
live Zepto flashing is still unstable.

## Scope

- Linux MicroBlocks VM only for the first functional backend
- Boardie keeps the primitive names reserved but does not implement the spooler yet
- Simulation models the Zepto as an I2C target that receives a controller write
  payload and optionally returns a reply payload

## Primitive Surface

The `components/microblocks-smallvm` Linux VM now exposes a new primitive set:

- `[i2ctarget:start]` address
- `[i2ctarget:stop]`
- `[i2ctarget:isStarted]`
- `[i2ctarget:address]`
- `[i2ctarget:hasRequest]`
- `[i2ctarget:receive]`
- `[i2ctarget:requestedBytes]`
- `[i2ctarget:reply]`

`[i2ctarget:receive]` returns a byte array containing the controller's write
payload. `[i2ctarget:requestedBytes]` reports how many bytes the controller
expects in the reply phase for the currently active transaction. `[i2ctarget:reply]`
publishes the response bytes and completes that transaction.

## Host Spool Format

The Linux VM uses a spool directory, defaulting to
`/tmp/microblocks_i2c_target_sim`. Override it with `MICROBLOCKS_I2C_SIM_DIR`.

Request files:

- `request-<address>-<request_id>.bin`
- first 4 bytes: little-endian reply length requested by the controller
- remaining bytes: controller write payload

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
4. Poll `badge i2c request pending?` from student code.
5. Use `badge i2c receive request` and `badge i2c requested bytes` to parse the transaction.
6. Reply with `badge i2c reply _`.
7. Drive transactions from the host with `scripts/microblocks_i2c_sim.py transaction ...`.
