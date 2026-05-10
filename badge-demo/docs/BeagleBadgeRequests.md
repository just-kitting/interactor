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
