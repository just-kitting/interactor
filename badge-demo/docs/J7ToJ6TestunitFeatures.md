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
