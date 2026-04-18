# I2C Kernel Simulation

This document records the difference between the current hosted Boardie bridge
and a real Linux-visible I2C simulation path.

## Current State

The hosted Boardie path is an application-level bridge:

- the browser receives queued controller-style transactions from the local web server
- Boardie/MicroBlocks processes those requests using the `i2ctarget` primitive surface
- the browser posts the reply back to the server

This preserves request semantics, but it does **not** create a Linux kernel I2C
target device. As a result, normal controller-side tools such as `i2ctransfer`
cannot talk to the browser session directly.

## Why `i2c-stub` Is Not The Right Replacement

Linux documents `i2c-stub` as a simple fake I2C/SMBus device backed by in-memory
arrays. Its typical use is:

1. load `i2c-stub`
2. preload register data
3. load a kernel client driver
4. observe that driver’s behavior

That is useful for static chip emulation, but it does not match what Boardie or
MicroBlocks needs:

- student code must observe controller writes as events
- student code must decide replies dynamically
- the simulation should match target-style request sequencing, not static register reads

Even on a kernel where `i2c-stub` is available, it would be the wrong primitive
for replacing the current Boardie simulation.

## Better Kernel-Level Directions

If the goal is to let Linux controller tools make real bus transactions against
the simulation, we need a lower-level path such as:

- an I2C slave/target backend that is visible from userspace and supports dynamic data flow
- a virtual adapter or stub bus with a userspace bridge process behind it
- a custom kernel/helper pair that forwards real I2C transactions into the Boardie runtime

The closest standard kernel pieces are the target/slave interface and backends
such as `i2c-slave-eeprom`, but those are still EEPROM-style backends, not a
complete dynamic Boardie target runtime.

## Live Probe

Use:

```sh
./scripts/check_i2c_kernel_sim_support.sh
```

to check whether the running kernel even ships likely pieces such as `i2c-stub`
or `i2c-slave-eeprom`.
