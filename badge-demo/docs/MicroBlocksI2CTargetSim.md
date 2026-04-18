# MicroBlocks I2C Target Simulation

This document defines the first MicroBlocks I2C target simulation path while
live Zepto flashing is still unstable.

## Scope

- Linux MicroBlocks VM with a filesystem spooler backend
- Boardie web build with an HTTP bridge backend
- Simulation models the Zepto as an I2C target that receives a controller write
  payload and may later see a controller read request
- The simulation now follows Zephyr's target-style split between write-receive
  and read-request handling rather than advertising a controller read length up front

## Primitive Surface

The `components/microblocks-smallvm` Linux VM now exposes a new primitive set:

- `[i2ctarget:start]` address
- `[i2ctarget:stop]`
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
- [I2C Target.ubl](/root/interactor/badge-demo/examples/microblocks/I2C%20Target.ubl): importable MicroBlocks library for these primitives
- [web_i2c_transaction.py](/root/interactor/badge-demo/scripts/web_i2c_transaction.py): send controller transactions into the hosted Boardie bridge
- [test_web_i2c_bridge.sh](/root/interactor/badge-demo/scripts/test_web_i2c_bridge.sh): smoke test for the hosted bridge endpoints

## Intended Usage

1. Build and run the Linux MicroBlocks VM.
2. Import `I2C Target.ubl`.
3. Start the simulated target at the intended Zepto address.
4. Poll `i2c write pending?` and `i2c read requested?` from student code.
5. Use `receive i2c write` to consume any controller write payload.
6. When a read is requested, publish bytes with `reply with _`.
7. Drive transactions from the host with `scripts/microblocks_i2c_sim.py transaction ...`.

## Web And IDE Hosting

The hosted web path now works through the existing Python server:

- the browser polls `/api/i2c/next` for pending controller transactions
- Boardie injects those requests into the `i2ctarget` queue
- the browser posts replies back to `/api/i2c/respond`
- Linux-side tests can use `scripts/web_i2c_transaction.py`

## What This Is Not

The hosted Boardie path is not a kernel-visible I2C target device.

That means:

- you cannot point `i2ctransfer` or another Linux I2C controller tool directly at
  the browser session and expect a real `/dev/i2c-*` target transaction
- the current bridge preserves target semantics at the application layer, not at
  the Linux I2C device-model layer

If we want true Linux-visible controller transactions for the hosted simulation,
the next step is a lower-level adapter such as:

- a kernel-backed target or emulation device that appears on a real I2C bus
- an `i2c-stub`-style or virtual adapter path with a userspace bridge
- a custom host-side process that owns a real or emulated adapter and forwards
  I2C transactions into Boardie/MicroBlocks

For the current BeagleBadge image, see
[I2CKernelSimulation.md](/root/interactor/badge-demo/docs/I2CKernelSimulation.md#L1)
for why `i2c-stub` is not a drop-in replacement for the hosted Boardie bridge.
