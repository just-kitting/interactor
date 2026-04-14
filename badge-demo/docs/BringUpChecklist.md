# Bring-Up Checklist

Use this checklist when validating a fresh BeagleBadge system for BadgeSnake work.

## Base System

- [ ] Board boots successfully
- [ ] `uname -a` reports the expected kernel
- [ ] `/boot` is mounted and readable
- [ ] `/proc/mtd` shows the expected OSPI partition layout
- [ ] `components/` submodules are initialized

## Bootloader Inspection

- [ ] Run `scripts/check_uboot_artifacts.sh`
- [ ] Confirm `/boot` loader artifacts match the packaged U-Boot payload
- [ ] Run `scripts/check_ospi_env.sh`
- [ ] Confirm whether the active U-Boot build still uses `CONFIG_ENV_IS_NOWHERE=y`

## Hardware Visibility

- [ ] Run `scripts/probe_badge_outputs.sh`
- [ ] Confirm `/dev/fb0` is present
- [ ] Confirm the ePaper path enumerates as `spi:gdey042t81`
- [ ] Confirm `gpio-keys` is present
- [ ] Confirm the expected QWIIC I2C buses are present

## BadgeSnake Repo State

- [ ] `docs/MissingInputs.md` does not end with an unanswered block quote
- [ ] `TODO.md` reflects current work state
- [ ] `PROGRESS.md` records any new findings
- [ ] The working tree is clean before risky boot changes
