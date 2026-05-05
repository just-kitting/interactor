# J7 To J6 Testunit Feature Validation

This document captures higher-level `slave-testunit` backend checks on the
working J7 -> J6 second-controller topology.

## Purpose

The basic J7 -> J6 smoke test proves:

- a byte write succeeds
- a byte read succeeds

That is enough to prove target mode basically works, but it is still a very
small slice of transport behavior.

The next useful checks are backend features which exercise:

- repeated-start behavior
- multi-byte reply generation

## One-Command Validation

Use:

```sh
./scripts/validate_j7_to_j6_testunit_features.sh
```

The script assumes:

- J6 is shorted to J7
- J6 (`/dev/i2c-1`) hosts `slave-testunit`
- J7 (`/dev/i2c-3`) is the initiator

## Checks Performed

### GET_VERSION_WITH_REP_START

This sends:

- command `0x04`
- followed by a repeated-start read

Success criteria:

- the returned byte stream starts with `v`
- which proves the repeated-start sequence was preserved

### SMBUS_BLOCK_PROC_CALL

This sends:

- command `0x03`
- `DATAL=1`
- `DATAH=<N>`
- followed by a read of `N + 1` bytes

Success criteria:

- the returned stream is:
  - length byte `N`
  - then descending bytes `N-1` to `0`

For the default script setting `N=4`, the expected response is:

```text
0x04 0x03 0x02 0x01 0x00
```

## Why This Matters

These checks are a better baseline for the future transport than a single
one-byte read/write smoke test, because they validate:

- multi-message transaction behavior
- multi-byte target responses
- nontrivial target-side state handling

## Current Known Result

On the current `#11` kernel:

- both combined write+read transactions complete electrically
- but both feature tests still return only zero bytes

That is a meaningful result, because it suggests the slave backend is losing
partial-command state across the write->read boundary.

The current kernel hypothesis is:

- `ARDY` is being treated as `STOP` too early in the OMAP slave path
- on repeated-start style traffic, that likely resets the `slave-testunit`
  state machine before the read phase begins

The next staged kernel follow-up therefore defers `I2C_SLAVE_STOP` on `ARDY`
until the bus is no longer busy.

## Current `P92b0` Result

On the current `#12` kernel:

- the repeated-start version query completes
- the block-proc-call style query also completes
- both still return only zero bytes

The remaining bug is now narrower than the original `STOP` hypothesis. In the
current OMAP slave IRQ path, `I2C_SLAVE_WRITE_REQUESTED` is emitted on every
`RRDY` byte. For `i2c-slave-testunit`, that callback clears the staged command
registers, so a command such as `4 0 0` is repeatedly reset before the read
phase starts.

The next staged kernel follow-up is:

- `0011-Track-slave-write-transaction-state.patch`

Its purpose is to send `WRITE_REQUESTED` only once per write transaction and
keep the command payload intact across the repeated-start boundary.

## Current `P1d46` Result

Booting the eleven-patch `P1d46` kernel fixes the repeated-start version path:

- `i2ctransfer -f -y -b 3 w3@0x30 4 0 0 r32`
- now returns `v6.12.57-vendor-edge-k3`

The block-proc-call path is improved but not fully correct yet:

- `i2ctransfer -f -y 3 w3@0x30 3 1 4 r5@0x30`
- returns:
  - `0x00 0x04 0x03 0x02 0x01`
- documented `slave-testunit` expectation remains:
  - `0x04 0x03 0x02 0x01 0x00`

So the current remaining bug is no longer loss of the full repeated-start
command state. It is a narrower transmit-side issue where the proc-call
response is still shifted by one byte.

## Next `P1d46` Follow-up

The next staged kernel change is:

- `0012-Handle-combined-slave-TX-slots.patch`

Reason:

- the proc-call read phase is the case where the first slave TX interrupt shows
  up with combined `XUDF|XRDY`
- the version path does not need this
- the current hypothesis is that the controller wants two TX bytes queued at the
  start of the proc-call read phase, but the current slave TX path only services
  that combined condition once

The follow-up therefore handles `XUDF` and `XRDY` as two explicit TX slots
instead of collapsing them into one callback.

## Current `P665e` Result

Booting the twelve-patch `P665e` kernel does not change the remaining proc-call
misalignment.

- repeated-start version query still works
- fixed-length proc-call read still returns:
  - `0x00 0x04 0x03 0x02 0x01`
- recv-len proc-call read still returns:
  - `0x04 0x00 0x00 0x00 0x00`

So the combined `XUDF|XRDY` two-slot hypothesis was not sufficient. The next
iteration needs more detailed proc-call transmit tracing or a different model
for how the controller consumes the first returned byte.
