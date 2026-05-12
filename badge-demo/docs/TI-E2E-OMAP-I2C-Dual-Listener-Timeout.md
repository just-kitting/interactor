# TI E2E Draft: AM62L OMAP I2C Target-Capable Dual-Listener Timeout

## Post Title

AM62L I2C OMAP: two target-capable controllers on same bus time out when either initiates

## What To Post

Copy the text below into a TI E2E forum post. Attach or paste the output from:

```sh
./scripts/reproduce_omap_i2c_dual_listener_timeout.sh
```

Run the script on the BeagleBadge while J6 and J7 are still shorted together
and no BeagleConnect Zepto boards are connected.

Latest captured artifact set in this repo:

- [run.log](/root/interactor/badge-demo/artifacts/ti-e2e/20260512T095152Z/run.log)
- [dmesg.log](/root/interactor/badge-demo/artifacts/ti-e2e/20260512T095152Z/dmesg.log)

## Post Body

I am working on BeagleBadge, based on TI AM62L, using the K3 vendor kernel
`6.12.57-vendor-edge-k3`. I am extending `drivers/i2c/busses/i2c-omap.c` so an
AM62L I2C adapter can register a Linux I2C slave backend and still temporarily
act as an I2C master/initiator.

The immediate test setup uses two AM62L OMAP I2C instances on the same board:

- J6 is `/dev/i2c-1`, controller instance `20010000.i2c`
- J7 is `/dev/i2c-3`, controller instance `20020000.i2c`
- J6 and J7 are physically shorted together
- Linux `i2c-slave-testunit` is used as the target/slave backend
- No external I2C devices are connected for this test

Validated working cases:

- J6 target at `0x30`, J7 initiator: true SMBus block-proc-call works
- J7 target at `0x30`, J6 initiator: true SMBus block-proc-call works
- repeated-start version reads work in the single-target reverse-topology case

The remaining failing case is when both adapters are target-capable at the same
time:

- J6 target at `0x30`
- J7 target at `0x31`

After both targets are registered and bound, either initiation direction times
out:

```sh
i2ctransfer -f -y 3 w1@0x30 0x00
i2ctransfer -f -y 3 r1@0x30
i2ctransfer -f -y 1 w1@0x31 0x00
i2ctransfer -f -y 1 r1@0x31
```

All four commands time out, even though each single-target topology works.

Driver changes already tested:

- added OMAP `reg_slave` / `unreg_slave` plumbing
- clear stale master state before returning to slave listen mode
- switch from the slave IRQ mask to the normal master IRQ mask when an adapter
  with a registered slave backend initiates a master transfer
- on AM62L/IP-v2, clear `I2C_IRQENABLE_CLR` before writing the next role's
  desired interrupt mask, because `I2C_IRQENABLE_SET` only sets bits

The last change did remove the previous stale IRQ mask problem:

- before: failing dual-listener logs showed accumulated `ie=0x661f`
- now: logs show cleaned masks such as `ie=0x61f` and `ie=0x601f`
- `Transmit underflow` is no longer observed

But the dual-listener case still times out in both directions.

The relevant current kernel state is:

- `components/ti-linux-kernel`: `eb09330dc065`
  `Clear OMAP IRQ enables before role mask writes`
- `components/armbian-build`: `b17e72e32`
  `Add OMAP IRQENABLE clear patch`
- tested artifact: `6.12.57-S22fb-D0000-Pb163-C2876Hb496-HK01ba-Vc222-Be8e3-R448a`
- booted kernel:
  `Linux beaglebadge 6.12.57-vendor-edge-k3 #27 SMP PREEMPT Tue May 5 16:45:44 UTC 2026 aarch64 GNU/Linux`

Question:

For AM62L OMAP I2C IP-v2, is there an additional required sequence when two
controllers on the same physical I2C bus are both left in target/slave listen
mode, but either one may temporarily become master?

Specifically:

1. When switching a target-capable adapter from slave-listen mode into master
   mode, besides clearing `I2C_IRQENABLE_CLR`, clearing status, setting `MST`,
   and restoring slave-listen mode afterward, is there another register or FIFO
   state that must be reset?
2. Is it expected that an OMAP I2C controller with own address enabled can
   interfere with another controller's transaction to a different target address
   on the same bus?
3. Should `OA` / target addressing be disabled while an adapter is initiating a
   master transfer, then restored afterward?
4. Are there AM62L-specific constraints for using two OMAP I2C instances as
   multi-controller participants on the same physical bus?

The attached/pasted reproduction log includes:

- exact commands
- `uname -a`
- target backend status
- exit status for each `i2ctransfer`
- `dmesg` lines containing `omap_i2c`, `Transmit underflow`, `Arbitration lost`,
  timeout, `slave irq`, `isr-master`, and `master-enter`

The goal is not to debug the Linux slave-testunit backend itself. The single
target cases prove that target mode and true SMBus block-proc-call work in both
directions. The problem only appears when both OMAP adapters are simultaneously
registered as target-capable listeners on the shorted J6/J7 bus.

## Local Context For BadgeSnake Agents

Do not move to BeagleConnect Zepto testing until this badge-only dual-listener
boundary is understood.

If TI confirms that `OA` should be disabled during master transfers, the next
local patch should test clearing/restoring own-address state around
`omap_i2c_set_master_mode()` / `omap_i2c_restore_slave_listen()`.
