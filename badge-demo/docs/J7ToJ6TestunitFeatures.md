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

## Next `P665e` Follow-up

The next staged kernel change is:

- `0013-Trace-slave-TX-byte-sequence.patch`

Purpose:

- log each slave TX callback as either:
  - `slave tx-requested`
  - `slave tx-processed`
- log the byte value returned for that slot

This should tell us whether the proc-call byte shift is already present on the
J6 target side, or whether the J7 initiator path is transforming an otherwise
correct target response.

## Current `P0e03` Result

Booting the thirteen-patch `P0e03` kernel answers the TX-side question.

- repeated-start version query still works
- proc-call response is still:
  - `0x00 0x04 0x03 0x02 0x01`

But the new J6 trace now shows the target-side TX callback/value sequence:

- `slave tx-requested stat=0x400 value=0x4`
- `slave tx-processed stat=0x10 value=0x3`
- `slave tx-processed stat=0x400 value=0x2`
- `slave tx-processed stat=0x400 value=0x1`

So the target is already producing the leading four meaningful bytes in the
expected order. The remaining proc-call bug is narrower:

- the repeated-start read phase still begins with an empty slot on the bus
- the final `0x00` never gets a chance to leave the target

That points to read-start handling on the repeated-start address phase, not to
the `slave-testunit` state machine or byte ordering within the target callback
sequence.

## Next `P0e03` Follow-up

The next staged kernel change is:

- `0014-Prime-first-slave-TX-byte-on-read-address-match.patch`

Purpose:

- preload the first proc-call response byte on the read-address match (`AAS`)
- avoid losing the first bus slot to an initial empty TX condition before
  `XUDF|XRDY` is raised

Expected effect:

- proc-call response should become:
  - `0x04 0x03 0x02 0x01 0x00`

## Current `P21f5` Result

Booting the fourteen-patch `P21f5` kernel proves the new `AAS`-time priming
path is active, but the visible proc-call bytes are still unchanged.

- repeated-start version query still works
- proc-call response still is:
  - `0x00 0x04 0x03 0x02 0x01`

The new slave trace shows:

- `slave tx-requested stat=0x200 value=0x4`
- `slave tx-processed stat=0x400 value=0x3`
- `slave tx-processed stat=0x10 value=0x2`
- `slave tx-processed stat=0x400 value=0x1`
- `slave tx-processed stat=0x10 value=0x0`
- `slave tx-processed stat=0x400 value=0x0`

So the target is now definitely trying to preload the read phase earlier, but
the controller still exposes the same visible five-byte proc-call response to
J7. The remaining problem is therefore more specific than “prime the first byte
earlier”; it is about how the controller treats the combined startup state at
the beginning of the repeated-start read phase.

## Next `P21f5` Follow-up

The next debugging target is the combined read-start state seen on J6, such as:

- `stat=0x614` (`AAS|ARDY|XUDF|XRDY`)

Purpose:

- determine whether data written on plain `AAS` is ignored until the following
  `XUDF|XRDY` cycle
- decide whether the combined startup interrupt needs a different servicing
  order to make the primed byte visible on the bus

## True SMBus Proc-Call Note

There is now an important distinction between:

- the raw `i2ctransfer` surrogate used in this repo
- a true userspace `I2C_SMBUS_BLOCK_PROC_CALL`

The AM62L `i2c-omap` initiator on `/dev/i2c-3` currently does **not**
advertise `I2C_FUNC_SMBUS_BLOCK_PROC_CALL`, so a direct userspace block-proc-
call ioctl is not available yet on this board.

The probe for that is now:

```sh
./scripts/test_j7_to_j6_smbus_block_proc_call.sh
```

Current result:

```text
adapter /dev/i2c-3 does not advertise I2C_FUNC_SMBUS_BLOCK_PROC_CALL
```

So the current raw proc-call mismatch:

- `0x00 0x04 0x03 0x02 0x01`

is still useful as a surrogate signal, but it is not yet the same thing as a
successful userspace SMBus block-proc-call validation on J7.

## Next Master-Side Follow-up

The next staged kernel change shifts to J7 master capability:

- `0015-Add-OMAP-SMBus-recv-len-support.patch`

Purpose:

- add `I2C_M_RECV_LEN` handling to the OMAP master receive path
- advertise `(I2C_FUNC_SMBUS_EMUL_ALL & ~I2C_FUNC_SMBUS_QUICK)`
- enable Linux's generic userspace SMBus block-proc-call emulation on `/dev/i2c-3`

If this works, the next live check is no longer only the raw surrogate:

```sh
./scripts/test_j7_to_j6_smbus_block_proc_call.sh
```

## Current `P6926` Result

Booting the fifteen-patch `P6926` kernel changes the J7 initiator capability
surface, but does not yet make the true userspace SMBus path reliable.

- J7 adapter functionality changed from:
  - `0xefe002d`
  - to `0xffe802d`
- so the old capability-gate failure is gone

But a clean direct run of the real userspace SMBus path can still fail:

```text
ioctl(I2C_SMBUS BLOCK_PROC_CALL): Connection timed out
```

and `dmesg` shows repeated:

```text
omap_i2c 20010000.i2c: Too much work in one IRQ
```

So the new master-side patch did change the adapter-level story, but the
receive-length handling behind that path is still not correct enough yet for a
stable SMBus block-proc-call on J7.

## Current `P6926` Raw `I2C_RDWR|I2C_M_RECV_LEN` Result

A direct userspace `I2C_RDWR` probe with `I2C_M_RECV_LEN` shows the new failure
mode more clearly than the higher-level SMBus wrapper:

```text
data=0x04 0x04 0x04 0x04 0x04
```

That is distinct from the raw `i2ctransfer` surrogate:

```text
0x00 0x04 0x03 0x02 0x01
```

So the current master-side recv-len support is not merely exposing the old
target-side proc-call misalignment. It is also capable of repeatedly consuming
the first length byte inside the same IRQ work pass.

The staged 16th patch therefore changes the OMAP master receive path so that
the first recv-len byte only reprograms `CNT`/FIFO state and then exits that
IRQ pass, leaving payload-byte consumption for the next real interrupt.

## Current `P641a` Result

Booting the sixteen-patch `P641a` kernel changed the failure mode again, but it
did not fix the true userspace recv-len path.

- repeated-start version query still works
- the raw proc-call surrogate still returns:
  - `0x00 0x04 0x03 0x02 0x01`
- true SMBus block-proc-call still times out
- direct `I2C_RDWR` + `I2C_M_RECV_LEN` now fails with:

```text
ioctl(I2C_RDWR): Protocol error
```

and `dmesg` shows:

```text
omap_i2c 20010000.i2c: Too much work in one IRQ
```

followed by a kernel oops in the IRQ thread:

```text
pc : omap_i2c_transmit_data.isra.0+0x70/0x1ac
lr : omap_i2c_xfer_data+0x264/0x338
x0 : 0000000000000000
```

So `P641a` should be treated as a regression diagnostic step, not a working
recv-len fix.

## Current `P3659` Result

Booting the distinct `P3659` build fixes the `P641a` crash, but the recv-len
data path is still not correct.

- true SMBus block-proc-call now returns:

```text
count=4
data=0x00 0x00 0x00 0x00
```

- direct `I2C_RDWR` + `I2C_M_RECV_LEN` now returns:

```text
data=0x04 0x00 0x00 0x00 0x00
```

- the older raw `i2ctransfer` surrogate still returns:

```text
0x00 0x04 0x03 0x02 0x01
```

Most importantly, the IRQ-thread crash from `P641a` was not reproduced on
`P3659`, so the stale opposite-direction TX status guard appears to have fixed
that regression.

The remaining problem is now the master-visible payload path after the count
byte, not the old crash path.
