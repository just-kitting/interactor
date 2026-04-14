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
- The packaged U-Boot payload installed on this system reports:
  - source repo: `https://github.com/TexasInstruments/ti-u-boot`
  - source branch: `ti-u-boot-2025.01-next`
  - git revision: `7b9dedb046eb6a720997f61582c7b13da1b5b9f0`
  - artifact version: `2025.01-S7b9d-P0000-H8a1d-V5719-Bbf55-R448a`

Current `/boot/tiboot3.bin`, `/boot/tispl.bin`, and `/boot/u-boot.img` should be treated as intermediate artifacts.
They may not be suitable for direct OSPI installation without validation.

On this board, those `/boot` artifacts are byte-for-byte identical to the files packaged under `/usr/lib/linux-u-boot-vendor-edge-beaglebadge/`.
That improves provenance confidence for the running image, but it still does not prove that writing them into OSPI is the complete safe deployment procedure.

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

However, the packaged U-Boot config on this system contains:

```text
CONFIG_ENV_SIZE=0x1f000
CONFIG_ENV_IS_DEFAULT=y
CONFIG_ENV_IS_NOWHERE=y
```

That means the current packaged bootloader is not configured to persist its environment in OSPI at all.

Testing `fw_printenv -c /tmp/fw_env_ospi.config` with the candidate layout therefore does not yield a readable environment on the current board state, which is consistent with the packaged configuration.

To make rollback and boot-attempt tracking work through U-Boot environment variables, a future bootloader build will need to move away from `CONFIG_ENV_IS_NOWHERE` and into an explicit persistent storage location.

Until then, any `fw_env.config` work is preparation for a later bootloader change, not a feature that the currently packaged U-Boot can use.

Do not treat environment writes as safe until this is validated against the U-Boot build configuration or a known-good deployed bootloader.
