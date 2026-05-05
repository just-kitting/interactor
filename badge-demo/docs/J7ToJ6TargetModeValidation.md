# J7 To J6 Target-Mode Validation

This document records the now-working validation path where J7 acts as a true
second controller against J6 in target mode.

## Purpose

The forced same-adapter self-test on J6 (`/dev/i2c-1` talking to its own target
address) still has edge cases and timeouts.

That is no longer the best proof point for AM62L target support.

The stronger proof is:

- J6 hosts the slave backend
- J7 acts as a separate initiator
- a short connects J6 and J7

In this setup:

- `/dev/i2c-1` is the target/slave side
- `/dev/i2c-3` is the controller side

## Expected Wiring

- remove the Zepto from J6 while doing this validation
- install a short between J6 and J7

This makes J7 a real second controller on the same physical bus as J6.

## One-Command Validation

Use:

```sh
./scripts/validate_j7_to_j6_target_mode.sh
```

The script will:

1. instantiate `slave-testunit` on J6 / `i2c-1`
2. write one byte from J7 / `i2c-3` to `0x30`
3. read one byte back from J7 / `i2c-3`
4. print the slave-side `dmesg` diagnostics for both operations

## Expected Success Shape

The current known-good pattern is:

- write exits successfully
- read exits successfully and returns `0x00`
- slave-side diagnostics show:
  - `AAS`
  - `RRDY` for write-side receive flow
  - `XRDY` and `XUDF` for read-side transmit flow

## Why This Matters

This is the current best proof that:

- AM62L target mode on J6 works with a real second controller
- the remaining issue is specifically the same-adapter self-test path
- future transport development should use J7 or another true initiator model as
  the validation baseline

## Next-Level Validation

Once the byte-level smoke test is passing, use:

```sh
./scripts/validate_j7_to_j6_testunit_features.sh
```

That script validates richer `slave-testunit` features over the same J7 -> J6
topology:

- repeated-start version query
- SMBus block-proc-call style multi-byte reply
