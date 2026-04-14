# Progress Log

## 2026-04-14

Initial planning pass completed.

### Repository state

- `badge-demo` currently contains only a project `README.md`
- The repo is appropriate as an integration root rather than a code-heavy project today
- No submodules exist yet

### Board state confirmed on this system

- Hostname: `beaglebadge`
- OS: Armbian-unofficial `26.02.0-trunk trixie`
- Kernel: `6.12.57-vendor-edge-k3`
- Running on BeagleBadge / AM62L3
- Root and `/boot` are on microSD (`mmcblk1p2` and `mmcblk1p1`)
- eMMC exists as `mmcblk0`
- OSPI flash is visible via MTD partitions from Linux
- `/boot` contains `tiboot3.bin`, `tispl.bin`, `u-boot.img`, `Image`, and `uInitrd`
- `fw_printenv` currently returns nothing because environment access is not configured yet

### Boot-related observations

- `dmesg` shows Linux sees OSPI partitions named for `tiboot3`, `tispl`, `u-boot`, primary/backup env, and an OSPI rootfs area
- `/usr/lib/u-boot/platform_install.sh` on this image only copies bootloader artifacts into `/boot`; it does not provide a complete OSPI flashing workflow by itself
- The current repository should therefore treat OSPI boot enablement as a planned deliverable, not as already solved

### Documentation added this session

- `AGENTS.md` for strategy and execution rules
- `TODO.md` for staged delivery tracking
- `PROGRESS.md` for verified findings and session notes

### Highest-value information still needed

- Exact upstream repo list to use as submodules
- BeagleBadge recovery instructions for bad OSPI flashes
- Serial-console wiring/access details
- Preferred source of truth for U-Boot and kernel artifacts
- BadgeSnake gameplay and protocol requirements

## 2026-04-14 (later)

Repository inputs expanded by the user and reviewed.

### New repository assets confirmed

- Added submodules:
  - `components/armbian-build`
  - `components/battlesnake-rules`
  - `components/beaglebadge`
  - `components/beagleconnect-zepto`
- Added architecture note: `docs/Architecture.md`
- Added AM62L reference PDFs under `docs/`

### Architecture findings now grounded in repo content

- BadgeSnake is currently targeting I2C over the two BeagleBadge QWIIC connectors
- `components/beaglebadge/docs/FAQ.md` confirms:
  - J6 is on I2C1
  - J7 is on I2C2
  - the OSPI flash device is ISSI `IS25WX256-JHLE` at 256Mbit
- `components/beagleconnect-zepto` hardware files expose explicit BSL-related nets for:
  - I2C
  - UART
  - bootloader invocation

### Important remaining gaps after review

- No board recovery guide was found in the imported BeagleBadge docs
- No serial-console access instructions were found in the imported BeagleBadge docs
- No exact `fw_env.config` offsets were found for the OSPI environment partitions
- No Zepto firmware SDK/examples repo was added yet beyond the hardware-design repo
- The BadgeSnake gameplay contract and host-side protocol remain undefined

### Documentation added in this follow-up

- `docs/MissingInputs.md` to keep the missing component/document list explicit

## 2026-04-14 (workflow)

Collaboration workflow clarified by the user.

- Open questions should be asked through committed documents in the repo
- Chat is acceptable for live monitoring, but not as the primary place to park pending decisions
- Every wrap-up must include a git commit
- `docs/MissingInputs.md` is now reserved for immediate blocking questions in a Q&A format
- Non-immediate items were moved out of `docs/MissingInputs.md`

## 2026-04-14 (user answers)

The user answered the immediate blocking questions in `docs/MissingInputs.md`.

### Decisions captured

- Recovery is operator-assisted, not agent-autonomous, because the serial console is only available through an external USB-to-serial host
- Holding `SEL` at power-on is the standard recovery path to boot from microSD instead of OSPI
- Bootloader source of truth is the Armbian build tree, currently targeting branch `2025.12-beaglebadge`
- Current `/boot/tiboot3.bin`, `/boot/tispl.bin`, and `/boot/u-boot.img` are intermediate artifacts and should not be treated as proven OSPI install images yet
- Gameplay should keep Battlesnake HTTP semantics and translate them over I2C
- Student-facing firmware should target MicroBlocks
- Example BadgeSnake player firmware should live in this repository
- Zepto flashing should use I2C BSL
- Preferred flashing tool today is `bb-imager-rs`

### Follow-up documentation added

- `docs/BootAndRecovery.md` for recovery, boot order, artifact provenance, and candidate U-Boot environment layout

### U-Boot environment status

- The live MTD layout clearly identifies `ospi.env` and `ospi.env.backup`
- A candidate `fw_env.config` can be derived from the partition table
- `fw_printenv` still does not read a valid environment with that candidate config, so environment handling remains unvalidated and should be treated carefully

## 2026-04-14 (repo scaffolding)

Further implementation planning and scaffolding completed without requiring new user input.

### Decisions captured

- Host runtime language selected: Go
- Initial repo layout selected:
  - `cmd/badgesnake`
  - `internal/protocol`
  - `examples/microblocks`
- Product defaults should track Battlesnake `rules` defaults for the first pass

### New documentation

- `docs/ProductSpec.md`
- `docs/Protocol.md`
- expanded `docs/BootAndRecovery.md` with packaged U-Boot metadata and environment findings

### New helper scripts

- `scripts/check_ospi_env.sh`
- `scripts/check_uboot_artifacts.sh`

### Important bootloader finding

- The packaged U-Boot config on this system uses `CONFIG_ENV_IS_NOWHERE=y`
- This explains why `fw_printenv` cannot read a persistent OSPI environment today
- Reliable rollback based on U-Boot environment variables therefore requires a future bootloader change, not just a userspace config file
- `/boot/tiboot3.bin`, `/boot/tispl.bin`, and `/boot/u-boot.img` are byte-for-byte identical to the packaged files under `/usr/lib/linux-u-boot-vendor-edge-beaglebadge/`
