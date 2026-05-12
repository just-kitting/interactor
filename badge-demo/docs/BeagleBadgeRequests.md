# BeagleBadge Requests

This file holds explicit requests from `bq2` to the live `beaglebadge`
validation agent. Treat it as the repo-backed handoff channel for hardware
work.

## 2026-05-09: Validate recv-len payload tracing

### Source State To Build

Build and copy back a new BeagleBadge `vendor-edge-k3` kernel artifact from
the repo state containing:

- top-level repo:
  - `2a18b76` `Advance OMAP recv-len payload tracing`
- `components/ti-linux-kernel`:
  - `a140e1de8f36` `Trace OMAP recv-len payload handling`
- `components/armbian-build`:
  - `4d98c2f38` `Carry OMAP recv-len payload tracing`

The new build should be distinct from the already-tested `P5248` artifact.
Record the new package suffix and build summary path before installing.

### Install And Test

After installing and rebooting into the new distinct artifact, run the recv-len
tests sequentially to keep memory pressure low:

```sh
./scripts/test_j7_to_j6_smbus_block_proc_call.sh
./scripts/validate_j7_to_j6_testunit_features.sh
```

Also run the direct `I2C_RDWR | I2C_M_RECV_LEN` probe used for the `P3659`,
`Pa309`, and `P5248` comparisons.

### Required Log Capture

Capture the new master-side recv-len diagnostics immediately after the direct
recv-len probe:

```sh
dmesg | grep 'recv-len'
```

Record enough surrounding `dmesg` to correlate the direct probe with J6
slave-side TX traces.

### What `bq2` Needs From The Result

Classify the J7 master-side failure into one of these cases:

- only the count byte is logged, which means payload RX interrupt/drain is not
  firing after count-byte reprogramming
- payload byte logs show `0x00`, which means J7 is reading zeros from
  `DATA_REG`
- payload byte logs show `0x03 0x02 0x01`, which means the driver copies good
  bytes and the data is lost later

For each run, record:

- running kernel identity from `uname -a`
- package suffix installed
- direct recv-len userspace output
- SMBus block-proc-call userspace output
- raw surrogate output
- `recv-len` dmesg lines
- relevant J6 slave TX trace lines

### 2026-05-10 live validation result

Completed on BeagleBadge with distinct kernel artifact:

- `6.12.57-S22fb-D0000-P7f58-C2876Hb496-HK01ba-Vc222-Be8e3-R448a`

Classification:

- J7 reads zeros from `DATA_REG` after the count-byte reprogramming

Observed outputs:

- `uname -a`:
  - `Linux beaglebadge 6.12.57-vendor-edge-k3 #23 SMP PREEMPT Tue May  5 16:45:44 UTC 2026 aarch64 GNU/Linux`
- true SMBus block-proc-call:
  - `count=4`
  - `data=0x00 0x00 0x00 0x00`
- raw surrogate:
  - `0x00 0x04 0x03 0x02 0x01`
- direct `I2C_RDWR | I2C_M_RECV_LEN`:
  - `0x04 0x00 0x00 0x00 0x00`

Relevant master-side `recv-len` lines:

```text
recv-len byte value=0x4 offset=1
recv-len count=4
recv-len byte value=0x0 offset=2
recv-len byte value=0x0 offset=3
recv-len byte value=0x0 offset=4
recv-len byte value=0x0 offset=5
```

Relevant J6 slave TX trace:

```text
slave tx-requested stat=0x200 value=0x4
slave tx-processed stat=0x400 value=0x3
slave tx-processed stat=0x10 value=0x2
slave tx-processed stat=0x400 value=0x1
slave tx-processed stat=0x10 value=0x0
```

## 2026-05-10: Validate recv-len without CNT rewrite

### Source State To Build

Build and copy back a new BeagleBadge `vendor-edge-k3` kernel artifact from
the repo state containing:

- top-level repo:
  - this request plus the submodule pointers to the commits below
- `components/ti-linux-kernel`:
  - `e2cf5df7bd9f` `Keep OMAP recv-len transfer count after length byte`
- `components/armbian-build`:
  - `3d54c59e8` `Carry OMAP recv-len no-CNT-rewrite experiment`

The new build should be distinct from the already-tested `P7f58` payload
tracing artifact. Record the new package suffix and build summary path before
installing.

### Purpose

`P7f58` proved that J7 reaches the payload receive path but reads zeros from
`DATA_REG` after count-byte reprogramming. This experiment avoids rewriting
`CNT` and `buf_len` after the first recv-len byte. It keeps the original
full-length controller transfer running while trimming `msg->len` so the
caller should still see only `1 + count + extra` bytes.

This tests whether the mid-read `CNT` rewrite is what makes OMAP lose the
target's early non-zero payload bytes.

### Install And Test

After installing and rebooting into the new distinct artifact, run:

```sh
./scripts/test_j7_to_j6_smbus_block_proc_call.sh
./scripts/validate_j7_to_j6_testunit_features.sh
```

Also run the direct `I2C_RDWR | I2C_M_RECV_LEN` probe used in the previous
comparisons.

### Required Log Capture

Capture the recv-len diagnostics immediately after the direct recv-len probe:

```sh
dmesg | grep 'recv-len'
```

### What `bq2` Needs From The Result

Record whether the direct recv-len userspace output changes from:

```text
0x04 0x00 0x00 0x00 0x00
```

to any form containing the non-zero payload bytes:

```text
0x04 0x03 0x02 0x01 ...
```

Also record whether the new `recv-len count=... keep_buf_len=...` log appears
and what subsequent `recv-len byte value=...` lines show.

### 2026-05-10 live validation result

Completed on BeagleBadge with distinct kernel artifact:

- `6.12.57-S22fb-D0000-Pac0a-C2876Hb496-HK01ba-Vc222-Be8e3-R448a`

Observed outputs:

- `uname -a`:
  - `Linux beaglebadge 6.12.57-vendor-edge-k3 #24 SMP PREEMPT Tue May  5 16:45:44 UTC 2026 aarch64 GNU/Linux`
- true SMBus block-proc-call:
  - `count=4`
  - `data=0x03 0x02 0x01 0x00`
- raw surrogate:
  - `0x00 0x04 0x03 0x02 0x01`
- direct `I2C_RDWR | I2C_M_RECV_LEN`:
  - `0x04 0x03 0x02 0x01 0x00`

Required new log evidence:

```text
recv-len count=4 extra=0 remaining=4 msg_len=5 keep_buf_len=32 cnt=0x0f threshold=16
recv-len byte value=0x3 offset=2
recv-len byte value=0x2 offset=3
recv-len byte value=0x1 offset=4
recv-len byte value=0x0 offset=5
```

Relevant J6 slave TX trace:

```text
slave tx-requested stat=0x200 value=0x4
slave tx-processed stat=0x400 value=0x3
slave tx-processed stat=0x10 value=0x2
slave tx-processed stat=0x400 value=0x1
slave tx-processed stat=0x10 value=0x0
```

Conclusion:

- keeping the original controller transfer active after the recv-len count
  byte restores the non-zero payload bytes on true SMBus block-proc-call and
  direct `I2C_RDWR | I2C_M_RECV_LEN`
- the remaining mismatch is now limited to the raw `i2ctransfer` surrogate path

## 2026-05-10: Validate cleaned recv-len patch

### Source State To Build

Build and copy back a new BeagleBadge `vendor-edge-k3` kernel artifact from
the repo state containing:

- top-level repo:
  - this request plus the submodule pointers to the commits below
- `components/ti-linux-kernel`:
  - `3c92960e2833` `Clean OMAP recv-len count-byte handling`
- `components/armbian-build`:
  - `e49bac3e9` `Clean OMAP recv-len Armbian patch`

The new build should be distinct from `Pac0a`.

### Purpose

`Pac0a` proved that keeping the original controller transfer active after the
recv-len count byte restores payload bytes on true SMBus block-proc-call and
direct `I2C_RDWR | I2C_M_RECV_LEN`.

This cleaned state preserves that behavior but removes the temporary
successful-path `recv-len` `dev_info_ratelimited()` diagnostics from
`i2c-omap.c`. The expected result is behavioral parity with `Pac0a`, not new
raw-surrogate behavior.

### Install And Test

After installing and rebooting into the new distinct artifact, run:

```sh
./scripts/test_j7_to_j6_smbus_block_proc_call.sh
```

Also run the direct `I2C_RDWR | I2C_M_RECV_LEN` probe used in the previous
comparisons.

### What `bq2` Needs From The Result

Record:

- the new package suffix and build summary path
- true SMBus block-proc-call output
- direct `I2C_RDWR | I2C_M_RECV_LEN` output
- whether `dmesg | grep 'recv-len'` is empty after successful runs

Expected functional outputs remain:

```text
count=4
data=0x03 0x02 0x01 0x00
```

and:

```text
0x04 0x03 0x02 0x01 0x00
```

### 2026-05-11 live validation result

Completed on BeagleBadge with distinct kernel artifact:

- `6.12.57-S22fb-D0000-P5910-C2876Hb496-HK01ba-Vc222-Be8e3-R448a`

Observed outputs:

- `uname -a`:
  - `Linux beaglebadge 6.12.57-vendor-edge-k3 #25 SMP PREEMPT Tue May  5 16:45:44 UTC 2026 aarch64 GNU/Linux`
- true SMBus block-proc-call:
  - `count=4`
  - `data=0x03 0x02 0x01 0x00`
- direct `I2C_RDWR | I2C_M_RECV_LEN`:
  - `0x04 0x03 0x02 0x01 0x00`

Successful-run `recv-len` log check:

```text
dmesg | grep 'recv-len'
```

returned no lines after the direct recv-len probe.

Conclusion:

- the cleaned patch preserves the working `Pac0a` recv-len behavior
- the temporary successful-path `recv-len` diagnostics are gone as intended

## 2026-05-11: Validate J6/J7 role reversal before Zepto

### Purpose

Before connecting BeagleConnect Zepto boards, validate the remaining
multi-controller boundary with only the BeagleBadge J6/J7 shorted-bus setup.

The current `P5910` result proves a controlled J7-initiator to J6-target path.
It does not yet prove a general multi-master arbitration transport.

### Keep Zepto Disconnected

For this validation pass, keep Zepto boards disconnected. The goal is to reduce
the problem to the two AM62L OMAP adapters first.

### Test 1: Reverse The Current Topology

Run the existing feature checks with J7 hosting `slave-testunit` and J6 acting
as initiator:

```sh
BADGESNAKE_SLAVE_BUS=3 BADGESNAKE_MASTER_BUS=1 ./scripts/test_j7_to_j6_smbus_block_proc_call.sh
BADGESNAKE_SLAVE_BUS=3 BADGESNAKE_MASTER_BUS=1 ./scripts/validate_j7_to_j6_testunit_features.sh
```

Record whether:

- true SMBus block-proc-call returns `count=4` and `data=0x03 0x02 0x01 0x00`
- direct or script-level feature checks return `0x04 0x03 0x02 0x01 0x00`
- any `Arbitration lost`, timeout, NACK, or reset messages appear in `dmesg`

### Test 2: Two-Address Dual-Listener Setup

If Test 1 passes, try keeping both adapters target-capable on the same shorted
bus at different target addresses:

- J6 target address: `0x30`
- J7 target address: `0x31`

Then run alternating transactions:

- J7 initiates to J6 at `0x30`
- J6 initiates to J7 at `0x31`

Record the exact commands used, outputs, and `dmesg` after each direction.

### What `bq2` Needs From The Result

Record:

- whether reverse topology works
- whether dual-listener setup works
- whether adapters can return to target-listen mode after acting as initiators
- whether any arbitration-lost or bus-recovery behavior is observed
- whether the next step should stay J6/J7 or move to one Zepto

### 2026-05-11 live validation result

Completed on BeagleBadge on the already-booted `P5910` kernel:

- `Linux beaglebadge 6.12.57-vendor-edge-k3 #25 SMP PREEMPT Tue May  5 16:45:44 UTC 2026 aarch64 GNU/Linux`

#### Test 1: Reverse topology

Commands:

```sh
BADGESNAKE_SLAVE_BUS=3 BADGESNAKE_MASTER_BUS=1 ./scripts/test_j7_to_j6_smbus_block_proc_call.sh
BADGESNAKE_SLAVE_BUS=3 BADGESNAKE_MASTER_BUS=1 ./scripts/validate_j7_to_j6_testunit_features.sh
```

Observed result:

- reverse topology does **not** work
- SMBus block-proc-call on `/dev/i2c-1` to J7 target `0x30` fails with:
  - `ioctl(I2C_SMBUS BLOCK_PROC_CALL): Input/output error`
- repeated-start version read on `/dev/i2c-1` to J7 target `0x30` fails with:
  - `Error: Sending messages failed: Input/output error`

Relevant `dmesg`:

```text
omap_i2c 20010000.i2c: slave xfer-msg stat=0x00 con=0x8400 ie=0x61f oa=0x30 sa=0x30 bufstat=0x8000 read=0 threshold=3
omap_i2c 20010000.i2c: Transmit underflow
omap_i2c 20020000.i2c: slave irq stat=0x208 ...
```

J7 target status after failure:

- `present: 3-1030`
- `bound: yes`

So J7 can still bind as a target, but J6 acting as initiator is not yet a
working mirror of the validated J7 -> J6 direction.

#### Test 2: Dual-listener setup

Target setup:

- J6 target `0x30`
- J7 target `0x31`

Commands used:

```sh
./scripts/bringup_i2c_slave_testunit.sh start 1 0x30
./scripts/bringup_i2c_slave_testunit.sh start 3 0x31
i2ctransfer -f -y 3 w1@0x30 0x00
i2ctransfer -f -y 3 r1@0x30
i2ctransfer -f -y 1 w1@0x31 0x00
i2ctransfer -f -y 1 r1@0x31
```

Observed result:

- both target backends instantiate and bind successfully
- once both listeners are active, **neither** initiation direction works

J7 -> J6 at `0x30`:

- write: `Input/output error`
- read: `Connection timed out`

Relevant `dmesg`:

```text
omap_i2c 20020000.i2c: Transmit underflow
omap_i2c 20020000.i2c: slave irq stat=0x04 con=0x8200 ie=0x661f oa=0x31 sa=0x30 ...
omap_i2c 20010000.i2c: slave irq stat=0x208 ...
omap_i2c 20010000.i2c: slave irq stat=0x610 ...
```

J6 -> J7 at `0x31`:

- write: `Input/output error`
- read: `Connection timed out`

Relevant `dmesg`:

```text
omap_i2c 20010000.i2c: Transmit underflow
omap_i2c 20010000.i2c: slave irq stat=0x04 con=0x8200 ie=0x661f oa=0x30 sa=0x31 ...
omap_i2c 20020000.i2c: slave tx-requested stat=0x400 value=0x0 read=1 write=0
```

Target status after both failed directions:

- `present: 1-1030`, `bound: yes`
- `present: 3-1031`, `bound: yes`

So both adapters do return to a bound target state after acting as initiators,
but the presence of the second active listener appears to poison the bus-level
transaction path.

#### Recommendation

The next step should stay J6/J7, not move to Zepto yet.

Reason:

- the controlled J7 -> J6 single-target path is validated
- reverse topology is still broken
- dual-listener setup breaks both directions before Zepto enters the picture
- this boundary should be understood on the two OMAP adapters first

## 2026-05-11: Validate role IRQ-mask follow-up

### Source State To Build

Build and copy back a new BeagleBadge `vendor-edge-k3` kernel artifact from
the repo state containing:

- top-level repo:
  - this request plus the submodule pointers to the commits below
- `components/ti-linux-kernel`:
  - `ef2ef02cdcd0` `Switch OMAP IRQ masks across master-slave roles`
- `components/armbian-build`:
  - `4fb843095` `Add OMAP role IRQ mask patch`

The new build should be distinct from `P5910`.

### Purpose

The latest badge-only failures have a common shape: the adapter acting as
initiator also has a slave backend registered, and the live interrupt mask still
contains slave-only `XUDF`.

The new patch switches to the normal master interrupt mask while an adapter is
temporarily acting as a master, then restores the slave interrupt mask when it
returns to target-listen mode.

### Keep Zepto Disconnected

Keep BeagleConnect Zepto disconnected for this pass. The goal is still to clear
the J6/J7 badge-only boundary first.

### Test 1: Clean Reverse Topology

Before running the reverse test, explicitly remove any stale J6 target at
`0x30`:

```sh
./scripts/bringup_i2c_slave_testunit.sh stop 1 0x30
./scripts/bringup_i2c_slave_testunit.sh stop 3 0x30
BADGESNAKE_SLAVE_BUS=3 BADGESNAKE_MASTER_BUS=1 ./scripts/test_j7_to_j6_smbus_block_proc_call.sh
BADGESNAKE_SLAVE_BUS=3 BADGESNAKE_MASTER_BUS=1 ./scripts/validate_j7_to_j6_testunit_features.sh
```

Record:

- true SMBus block-proc-call output
- repeated-start feature output
- whether any `Transmit underflow`, `Arbitration lost`, timeout, or NACK
  appears in `dmesg`

### Test 2: Dual-Listener Setup

If Test 1 passes, retry the two-listener setup:

```sh
./scripts/bringup_i2c_slave_testunit.sh stop 1 0x30
./scripts/bringup_i2c_slave_testunit.sh stop 3 0x31
./scripts/bringup_i2c_slave_testunit.sh start 1 0x30
./scripts/bringup_i2c_slave_testunit.sh start 3 0x31
i2ctransfer -f -y 3 w1@0x30 0x00
i2ctransfer -f -y 3 r1@0x30
i2ctransfer -f -y 1 w1@0x31 0x00
i2ctransfer -f -y 1 r1@0x31
```

Record:

- output and exit status for each direction
- target status after each direction
- `dmesg` lines mentioning `Transmit underflow`, `slave irq`, `isr-master`,
  `Arbitration lost`, timeout, NACK, or reset

### What `bq2` Needs From The Result

Record whether `Transmit underflow` disappears when a target-capable adapter
initiates after the role IRQ-mask patch.

If reverse topology and dual-listener tests both pass, the next validation can
move to one Zepto. If either still fails, keep the work on J6/J7 and capture the
new `dmesg` signature.

### 2026-05-11 live validation result

Completed on BeagleBadge with distinct kernel artifact:

- `6.12.57-S22fb-D0000-P957d-C2876Hb496-HK01ba-Vc222-Be8e3-R448a`
- `Linux beaglebadge 6.12.57-vendor-edge-k3 #26 SMP PREEMPT Tue May  5 16:45:44 UTC 2026 aarch64 GNU/Linux`

#### Test 1: Clean reverse topology

Commands:

```sh
./scripts/bringup_i2c_slave_testunit.sh stop 1 0x30
./scripts/bringup_i2c_slave_testunit.sh stop 3 0x30
BADGESNAKE_SLAVE_BUS=3 BADGESNAKE_MASTER_BUS=1 ./scripts/test_j7_to_j6_smbus_block_proc_call.sh
BADGESNAKE_SLAVE_BUS=3 BADGESNAKE_MASTER_BUS=1 ./scripts/validate_j7_to_j6_testunit_features.sh
```

Observed result:

- true SMBus block-proc-call on `/dev/i2c-1` -> J7 target now works:
  - `count=4`
  - `data=0x03 0x02 0x01 0x00`
- repeated-start version read also works
- the raw surrogate still returns an unexpected shifted form:
  - `0x00 0x00 0x04 0x03 0x02`

Important `dmesg` difference from the old failure:

- the previous reverse-topology `Transmit underflow` on the J6 initiator path
  is gone
- `dmesg` for the successful SMBus reverse test ends with:

```text
omap_i2c 20010000.i2c: master-xfer addr=0x30 no registered slave
```

So the role IRQ-mask patch clears the old hard reverse-topology failure, but
the raw surrogate is still malformed.

#### Test 2: Dual-listener setup

Target setup:

- J6 target `0x30`
- J7 target `0x31`

Commands:

```sh
./scripts/bringup_i2c_slave_testunit.sh start 1 0x30
./scripts/bringup_i2c_slave_testunit.sh start 3 0x31
i2ctransfer -f -y 3 w1@0x30 0x00
i2ctransfer -f -y 3 r1@0x30
i2ctransfer -f -y 1 w1@0x31 0x00
i2ctransfer -f -y 1 r1@0x31
```

Observed result:

- both targets instantiate and bind
- but both initiation directions still fail

J7 -> J6 at `0x30`:

- write: `Input/output error`
- read: `Connection timed out`

Relevant `dmesg`:

```text
omap_i2c 20020000.i2c: Transmit underflow
omap_i2c 20020000.i2c: slave irq stat=0x04 con=0x8200 ie=0x661f oa=0x31 sa=0x30 ...
```

J6 -> J7 at `0x31`:

- write: `Input/output error`
- read: `Connection timed out`

Relevant `dmesg`:

```text
omap_i2c 20010000.i2c: Transmit underflow
omap_i2c 20010000.i2c: slave irq stat=0x04 con=0x8200 ie=0x661f oa=0x30 sa=0x31 ...
```

After both failed directions:

- `present: 1-1030`, `bound: yes`
- `present: 3-1031`, `bound: yes`

So the dual-listener boundary is still not clear even after the role IRQ-mask
patch.

#### Recommendation

Do **not** move to Zepto yet.

The role IRQ-mask patch fixes the old reverse-topology hard failure, but the
dual-listener setup still breaks both directions with `Transmit underflow`.
Continue on J6/J7 until that boundary is understood.

## 2026-05-11: Validate IRQENABLE clear follow-up

### Source State To Build

Build and copy back a new BeagleBadge `vendor-edge-k3` kernel artifact from
the repo state containing:

- top-level repo:
  - this request plus the submodule pointers to the commits below
- `components/ti-linux-kernel`:
  - `eb09330dc065` `Clear OMAP IRQ enables before role mask writes`
- `components/armbian-build`:
  - `b17e72e32` `Add OMAP IRQENABLE clear patch`

The new build should be distinct from `P957d`.

### Purpose

`P957d` fixed the reverse-topology hard failure, but the dual-listener setup
still shows `ie=0x661f` and target-side `Transmit underflow`.

The likely remaining bug is that AM62L/IP-v2 maps the normal interrupt-enable
register to `I2C_IRQENABLE_SET`. Writing a new role mask there can leave old
role bits enabled. The new patch clears `I2C_IRQENABLE_CLR` before writing the
desired mask, so role changes should replace the live interrupt mask rather than
accumulate it.

### Keep Zepto Disconnected

Keep BeagleConnect Zepto disconnected for this pass.

### Install And Test

After installing and rebooting into the distinct artifact, rerun the
dual-listener setup:

```sh
./scripts/bringup_i2c_slave_testunit.sh stop 1 0x30
./scripts/bringup_i2c_slave_testunit.sh stop 3 0x31
./scripts/bringup_i2c_slave_testunit.sh start 1 0x30
./scripts/bringup_i2c_slave_testunit.sh start 3 0x31
i2ctransfer -f -y 3 w1@0x30 0x00
i2ctransfer -f -y 3 r1@0x30
i2ctransfer -f -y 1 w1@0x31 0x00
i2ctransfer -f -y 1 r1@0x31
```

Also rerun the clean reverse-topology true SMBus path as a regression check:

```sh
./scripts/bringup_i2c_slave_testunit.sh stop 1 0x30
./scripts/bringup_i2c_slave_testunit.sh stop 3 0x30
BADGESNAKE_SLAVE_BUS=3 BADGESNAKE_MASTER_BUS=1 ./scripts/test_j7_to_j6_smbus_block_proc_call.sh
```

### What `bq2` Needs From The Result

Record:

- package suffix and build summary path
- whether dual-listener J7 -> J6 and J6 -> J7 write/read commands pass
- `dmesg` lines showing `ie=...` for `slave irq`, `isr-master`, and
  `master-enter`
- whether `Transmit underflow` remains
- whether reverse-topology true SMBus still works

If the live `ie=` values no longer contain stale role bits and the dual-listener
case still fails, that is the point where a concise TI E2E question becomes
worth preparing.

### 2026-05-12 live validation result

Completed on BeagleBadge with distinct kernel artifact:

- `6.12.57-S22fb-D0000-Pb163-C2876Hb496-HK01ba-Vc222-Be8e3-R448a`
- `Linux beaglebadge 6.12.57-vendor-edge-k3 #27 SMP PREEMPT Tue May  5 16:45:44 UTC 2026 aarch64 GNU/Linux`

#### Dual-listener setup

Setup:

```sh
./scripts/bringup_i2c_slave_testunit.sh stop 1 0x30
./scripts/bringup_i2c_slave_testunit.sh stop 3 0x31
./scripts/bringup_i2c_slave_testunit.sh start 1 0x30
./scripts/bringup_i2c_slave_testunit.sh start 3 0x31
```

Observed:

- both targets bind

J7 -> J6 at `0x30`:

- write: `Connection timed out`
- read: `Connection timed out`

Relevant `dmesg`:

```text
omap_i2c 20010000.i2c: slave irq ... ie=0x61f ...
omap_i2c 20020000.i2c: slave irq ... ie=0x601f ...
```

J6 -> J7 at `0x31`:

- write: `Connection timed out`
- read: `Connection timed out`

No `Transmit underflow` was observed in this pass.

#### Reverse-topology regression check

Commands:

```sh
./scripts/bringup_i2c_slave_testunit.sh stop 1 0x30
./scripts/bringup_i2c_slave_testunit.sh stop 3 0x31
./scripts/bringup_i2c_slave_testunit.sh stop 3 0x30
BADGESNAKE_SLAVE_BUS=3 BADGESNAKE_MASTER_BUS=1 ./scripts/test_j7_to_j6_smbus_block_proc_call.sh
```

Observed:

- reverse-topology true SMBus still works:
  - `count=4`
  - `data=0x03 0x02 0x01 0x00`
- relevant `dmesg` shows only slave-side `ie=0x61f`
- no `Transmit underflow` was observed in this regression check

#### Conclusion

- the stale role-bit hypothesis is now validated:
  - old `ie=0x661f` accumulation is gone
  - dual-listener failures remain, but the signature changed from underflow to timeout
- reverse-topology true SMBus remains good
- this is the point where preparing a concise TI E2E question is worth doing
- do **not** move to Zepto yet

### 2026-05-11 copied-build check

The first copied build after this request did **not** satisfy the requested
source state.

Observed copied summary:

- `components/armbian-build/output/logs/summary-kernel-a3757ada-e16e-4d64-9688-5b518ebf372d.md`

Observed package version:

- `6.12.57-S22fb-D0000-P5910-C2876Hb496-HK01ba-Vc222-Be8e3-R448a`

Why it is stale:

- the summary still shows only `16 total patches`
- the resulting package suffix is the already-tested `P5910`
- therefore it does **not** include the requested `0017` role IRQ-mask patch

Current superproject pointers are already correct:

- `components/ti-linux-kernel`:
  - `ef2ef02cdcd0` `Switch OMAP IRQ masks across master-slave roles`
- `components/armbian-build`:
  - `4fb843095` `Add OMAP role IRQ mask patch`

So the next required step is to rebuild from the updated submodule state and
copy back a distinct post-`P5910` artifact before attempting live validation.
