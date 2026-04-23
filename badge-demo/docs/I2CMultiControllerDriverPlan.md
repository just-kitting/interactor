# Multi-Controller I2C Driver Plan

This document plans a Linux kernel transport for BadgeSnake that uses the AM62L
host and MSPM0L targets as multi-controller I2C participants while preserving a
message-oriented software model.

## Why This Driver Exists

MicroBlocks does not want I2C to behave like a simple request/response
controller/target link, because both sides can generate messages asynchronously.

The important hardware fact is:

- transfers are still simplex at the wire level
- but both the AM62L host and MSPM0L target can arbitrate as I2C controllers

So the transport should model:

- asynchronous messages in both directions
- arbitration loss as a normal retry condition
- transaction boundaries as message boundaries
- no assumption that one side is permanently the sole initiator

## Kernel Starting Point In `components/ti-linux-kernel`

The added TI kernel tree already contains the main building blocks we need:

- generic slave core:
  - `drivers/i2c/i2c-core-slave.c`
  - `Documentation/i2c/slave-interface.rst`
- example slave backends:
  - `drivers/i2c/i2c-slave-eeprom.c`
  - `drivers/i2c/i2c-slave-testunit.c`
- `i2c-stub` for static SMBus client-driver testing:
  - `drivers/i2c/i2c-stub.c`
  - `Documentation/i2c/i2c-stub.rst`

The critical missing piece for BeagleBadge is:

- `drivers/i2c/busses/i2c-omap.c` does **not** currently wire up:
  - `reg_slave`
  - `unreg_slave`

That means the first implementation phase is enabling slave mode on the AM62L
host adapter itself.

## Recommended Architecture

Do **not** make the kernel ABI look like a raw tty first.

The kernel transport should be:

- message-oriented
- address-aware
- transaction-bounded
- usable from kernel or userspace

If MicroBlocks wants a byte-stream abstraction, add that as a userspace shim on
top of the kernel message transport.

### Layering

1. OMAP adapter slave support
   - extend `i2c-omap.c` so the AM62L adapter can be both controller and slave
2. BadgeSnake slave backend
   - new backend driver that receives slave events and exposes inbound messages
3. Outbound controller path
   - same driver uses `i2c_transfer()` to send framed writes to peer addresses
4. Userspace bridge
   - miscdevice or character-device API for host applications
5. Optional serial shim
   - PTY or stream adapter in userspace if MicroBlocks insists on serial semantics

## Proposed On-Wire Model

Use one I2C write transaction as one transport frame.

Do not depend on multi-message combined transactions for the base protocol.

### Frame Header

Each frame should begin with a fixed header, for example:

- `version`
- `flags`
- `source_endpoint`
- `destination_endpoint`
- `sequence`
- `payload_length`
- `header_crc`

Followed by:

- payload bytes
- optional payload CRC

### Why Include Source And Destination

When Linux acts as an I2C slave receiving a write, the hardware generally knows
the local destination address but not a useful higher-level source identity.

So the frame must self-identify:

- who sent it
- who it is for
- how it should be ordered/retried

## Host Kernel Driver Shape

Recommended first driver:

- file: `drivers/i2c/i2c-slave-badgesnake.c`
- config: `CONFIG_I2C_SLAVE_BADGESNAKE`
- model: standard I2C slave backend registered at `0x1000 + local_address`

### Core Responsibilities

- register an own slave address via `i2c_slave_register()`
- collect inbound bytes using:
  - `I2C_SLAVE_WRITE_REQUESTED`
  - `I2C_SLAVE_WRITE_RECEIVED`
  - `I2C_SLAVE_STOP`
- service outbound reads only if we decide the protocol needs them
  - likely not needed in the first push-only frame model
- expose inbound frames to userspace through a miscdevice or character device
- accept outbound frames from userspace and transmit them with `i2c_transfer()`

### Suggested Userspace ABI

Prefer a datagram-style character device over a tty.

Example shape:

- `/dev/badgesnake-i2c0`
- `read()` returns one full frame at a time
- `write()` accepts one full frame including destination endpoint
- `poll()`/`epoll()` for readiness
- `ioctl()` for:
  - local endpoint configuration
  - MTU query
  - statistics
  - fault injection or trace toggles

If we later need per-peer devices, add a small userspace mux daemon rather than
baking address-specific minors into the first kernel ABI.

## OMAP Adapter Work

Files likely touched:

- `components/ti-linux-kernel/drivers/i2c/busses/i2c-omap.c`
- `components/ti-linux-kernel/drivers/i2c/Kconfig`
- `components/ti-linux-kernel/drivers/i2c/Makefile`

### Expected OMAP Additions

- implement `reg_slave()` and `unreg_slave()`
- program own-address registers and enable slave interrupts
- translate OMAP slave IRQ conditions into `i2c_slave_event()` calls
- preserve master-mode support while slave mode is enabled
- handle arbitration loss as a retryable controller-side condition
- validate runtime-PM behavior when the adapter must remain responsive as a slave

### Validation Target

Before writing BadgeSnake-specific logic, prove OMAP slave support with the
existing kernel backend:

- enable `CONFIG_I2C_SLAVE_TESTUNIT`
- instantiate `slave-testunit`
- run multi-controller and repeated-start tests from Linux tools

If that step fails, do not proceed to the BadgeSnake-specific backend yet.

## Zepto Firmware Side

The MSPM0L side should mirror the same framing model:

- one I2C write transaction equals one transport frame
- inbound ISR or driver code reconstructs frames
- controller-originated sends are retried after arbitration loss or bus-busy
- endpoint IDs live in the frame header, not in implicit bus state

This keeps the host and MCU logic symmetric.

## Simulation Strategy

There should be two simulation levels:

1. Application-level simulation
   - current hosted Boardie bridge
   - good for IDE work and quick iteration
2. Kernel-visible transport simulation
   - future path for `i2ctransfer` and real controller-tool testing

The second path should be based on the same frame format and userspace ABI as
the real host driver, not on the current HTTP bridge semantics.

## Phased Delivery

### Phase 0: Kernel Capability Bring-Up

- enable in config:
  - `CONFIG_I2C_SLAVE`
  - `CONFIG_I2C_SLAVE_TESTUNIT`
  - optionally `CONFIG_I2C_SLAVE_EEPROM`
- prove whether `i2c-omap.c` can be extended cleanly on AM62L

### Phase 1: OMAP Slave Support

- add slave registration and IRQ event plumbing to `i2c-omap.c`
- instantiate `slave-testunit` via sysfs on J6/J7 bus
- verify multi-controller tests from `Documentation/i2c/slave-testunit-backend.rst`

### Phase 2: BadgeSnake Backend

- add `i2c-slave-badgesnake.c`
- expose miscdevice/chardev ABI
- support inbound frame receive and outbound controller writes

### Phase 3: Userspace Host Integration

- add a small host-side library or daemon in this repo
- give Go runtime and MicroBlocks bridge a common transport entrypoint

### Phase 4: Optional Stream Shim

- only if needed, add a PTY adapter for tools expecting a serial stream
- keep this outside the kernel ABI

## Risks

- AM62L OMAP adapter may need nontrivial slave-mode work
- multi-controller corner cases will surface:
  - arbitration loss
  - repeated starts
  - clock stretching expectations
  - runtime PM interactions
- trying to force a tty abstraction too early may hide addressing and framing bugs

## Recommendation

Yes, this kernel is a workable starting point for the plan.

But the sequence matters:

1. enable slave mode on the OMAP/AM62L adapter
2. validate with `slave-testunit`
3. add a BadgeSnake-specific message backend
4. only then decide whether a serial-style shim is still needed
