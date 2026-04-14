# Boot And Recovery

This document captures the current working understanding of BeagleBadge boot, recovery, and U-Boot environment handling for BadgeSnake.

## Recovery Model

- The board normally attempts to boot from OSPI first.
- If OSPI boot fails, the board attempts UART boot using the attached USB-to-UART path.
- Holding the `SEL` button during power-on changes boot mode to microSD first.
- If microSD boot fails in that mode, the board attempts USB DFU boot.

These boot modes are documented in `components/beaglebadge/docs/FAQ.md`.

## Practical Recovery Constraints

- The agent cannot use the serial console directly because the agent session runs on the board itself.
- Recovery over serial therefore depends on an external host and operator.
- The minimum operator-assisted recovery path for a bad OSPI flash is:
  - power off the board
  - hold `SEL`
  - power on
  - boot from microSD
  - restore or replace OSPI contents from the recovered Linux system

Because the session cannot survive a failed boot on its own, every meaningful step must be committed before risky boot changes.

## Target Reliability Behavior

OSPI bootloader work should not stop at "board boots once".

The intended behavior is:

- a known-good previous boot path remains available
- a new boot is only marked good after userland reaches a healthy state
- failure to mark a boot good causes rollback to the previous known-good kernel, device tree, and agent execution image

That likely implies:

- explicit U-Boot boot-attempt state
- a userspace good-boot marker
- initrd or equivalent early recovery logic if userspace needs to restore state before normal boot proceeds

## U-Boot Artifact Provenance

- Source of truth today: `components/armbian-build`
- Target branch: `2025.12-beaglebadge`
- The current running image and the bootloader artifacts on this board were built from the HEAD of that branch, per user guidance

Current `/boot/tiboot3.bin`, `/boot/tispl.bin`, and `/boot/u-boot.img` should be treated as intermediate artifacts.
They may not be suitable for direct OSPI installation without validation.

## Candidate U-Boot Environment Layout

The live MTD partition map on this board is:

```text
mtd0  ospi.tiboot3      0x00080000
mtd1  ospi.tispl        0x00200000
mtd2  ospi.u-boot       0x00400000
mtd3  ospi.env          0x00040000
mtd4  ospi.env.backup   0x00040000
mtd5  ospi.rootfs       0x018c0000
mtd6  ospi.phypattern   0x00040000
```

Each partition reports an erase size of `0x20000`.

The current best candidate `fw_env.config` derived from that layout is:

```text
/dev/mtd3 0x0 0x40000 0x20000
/dev/mtd4 0x0 0x40000 0x20000
```

This is a derived candidate, not yet a validated production setting.

Testing `fw_printenv -c /tmp/fw_env_ospi.config` with that layout does not yield a readable environment on the current board state, which suggests at least one of:

- the OSPI environment has not been initialized by the installed bootloader
- the true environment size is smaller than the full partition size
- the stored format differs from what `fw_printenv` expects with this config alone

Do not treat environment writes as safe until this is validated against the U-Boot build configuration or a known-good deployed bootloader.
