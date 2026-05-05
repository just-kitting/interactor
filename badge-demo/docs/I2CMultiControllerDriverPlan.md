# Multi-Controller I2C Driver Plan

This document plans a Linux kernel transport for BadgeSnake that uses the AM62L
host and MSPM0L targets as multi-controller I2C participants while creating
serial-style endpoints for software on top of the I2C bus.

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

## Current Status

That first implementation phase is now underway in the TI kernel submodule:

- a first `i2c-omap.c` patch is staged locally
- it adds:
  - `reg_slave()` / `unreg_slave()`
  - slave IRQ routing into `i2c_slave_event()`
  - restoration of slave-listen mode after host master transfers

That first implementation phase is now far enough along that J6 target mode has
been validated with a true second controller:

- J6 `/dev/i2c-1` hosts `slave-testunit`
- J7 `/dev/i2c-3` acts as the initiator via a short between J6 and J7
- both write and read transactions succeed in that topology

The remaining failure is specifically the same-adapter self-test path, not
generic AM62L target support.

## Recommended Architecture

The end goal is **not** a target-only driver.

The end goal is a serial-to-I2C endpoint bridge where:

- each peer on the transport is represented as an endpoint
- software can open a serial-style endpoint for a peer
- the bridge maps serial payload flow onto framed multi-controller I2C traffic
- either side may initiate a transaction when it has queued data

The kernel transport still needs an internal framed model because I2C is
transactional and address-based, but the exposed interface should be able to
support serial semantics cleanly.

### Layering

1. OMAP adapter slave support
   - extend `i2c-omap.c` so the AM62L adapter can be both controller and slave
2. Bridge core
   - new driver owns endpoint state, frame queues, arbitration/retry handling,
     and serial-facing device nodes
3. Inbound slave path
   - slave events are decoded into frames and delivered to endpoint RX queues
4. Outbound controller path
   - queued endpoint TX bytes are packetized and sent with `i2c_transfer()`
5. Endpoint presentation
   - create serial-style character devices, one per logical peer endpoint

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

Recommended driver family:

- core file: `drivers/i2c/i2c-badgesnake-bridge.c`
- config: `CONFIG_I2C_BADGESNAKE_BRIDGE`
- model:
  - uses standard I2C slave registration for inbound transactions
  - uses controller-mode `i2c_transfer()` for outbound transactions
  - exposes serial-style endpoint devices to software

### Core Responsibilities

- register one or more own slave addresses via `i2c_slave_register()`
- collect inbound bytes using:
  - `I2C_SLAVE_WRITE_REQUESTED`
  - `I2C_SLAVE_WRITE_RECEIVED`
  - `I2C_SLAVE_STOP`
- decode inbound transactions into bridge frames
- maintain per-endpoint RX/TX queues
- packetize outbound queued data into bridge frames
- send outbound frames with `i2c_transfer()`
- handle arbitration loss and backoff as normal bridge behavior
- expose serial-style endpoints to userspace

### Suggested Userspace ABI

Expose one serial-style character device per logical peer endpoint.

Example shape:

- `/dev/badgesnake/zepto0`
- `/dev/badgesnake/zepto1`
- `/dev/badgesnake/sim0`

Semantics:

- `read()` returns ordered byte-stream data for that endpoint
- `write()` queues byte-stream data for that endpoint
- the driver is responsible for framing, destination addressing, and retries
- `poll()`/`epoll()` reports stream readiness
- `ioctl()` configures:
  - remote I2C address
  - local endpoint ID
  - MTU / frame sizing policy
  - statistics
  - trace and fault-injection controls

If a tty-compatible layer is needed for existing tools, the bridge can either:

- implement tty-port glue directly, or
- expose character devices and pair them with a userspace PTY helper

The key point is that endpoint-oriented serial semantics are the target ABI.

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

Before writing the full bridge, prove OMAP slave support with the existing
kernel backend:

- enable `CONFIG_I2C_SLAVE_TESTUNIT`
- instantiate `slave-testunit`
- run multi-controller and repeated-start tests from Linux tools

If that step fails, do not proceed to the bridge driver yet.

## Serial Endpoint Model

The bridge should treat each peer as a duplex logical endpoint even though the
wire is simplex per transaction.

That means:

- each endpoint has independent RX and TX queues
- either side may initiate a transaction to drain its TX queue
- framing is invisible to the endpoint consumer
- ordering is per endpoint, not per raw bus transaction

This is the model that should align with MicroBlocks.

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
- follow [I2CSlaveBringup.md](/root/interactor/badge-demo/docs/I2CSlaveBringup.md#L1)
  for the first `slave-testunit` validation path

### Phase 1: OMAP Slave Support

- add slave registration and IRQ event plumbing to `i2c-omap.c`
- instantiate `slave-testunit` via sysfs on J6/J7 bus
- verify multi-controller tests from `Documentation/i2c/slave-testunit-backend.rst`

### Phase 2: Bridge Driver

- add `i2c-badgesnake-bridge.c`
- support inbound frame receive and outbound controller writes
- create serial-style endpoint devices

### Phase 3: Host Integration

- add a small host-side library or daemon in this repo
- give Go runtime and MicroBlocks a common serial-endpoint transport entrypoint

### Phase 4: Compatibility Layer

- if needed, add tty-port or PTY compatibility for tools that insist on tty nodes
- keep the internal bridge logic endpoint-oriented and framed

## Risks

- AM62L OMAP adapter may need nontrivial slave-mode work
- multi-controller corner cases will surface:
  - arbitration loss
  - repeated starts
  - clock stretching expectations
  - runtime PM interactions
- hiding endpoint identity too aggressively may make transport debugging harder
- tty-style compatibility may still be needed for some user tools

## Recommendation

Yes, this kernel is a workable starting point for the plan.

But the sequence matters:

1. enable slave mode on the OMAP/AM62L adapter
2. validate with `slave-testunit`
3. add the BadgeSnake serial-to-I2C bridge driver
4. then decide whether explicit tty glue is still needed beyond the endpoint devices
