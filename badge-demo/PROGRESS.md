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

## 2026-04-14 (output probing)

Live hardware-output probing completed.

### Findings

- The current image exposes:
  - `/dev/fb0`
  - `/dev/dri`
  - `/dev/i2c-0`
  - `/dev/i2c-2`
- The current image does not expose an ALSA sound card
- SPI devices identify:
  - OSPI NOR flash
  - the `gdey042t81` ePaper panel path
  - an `mcp23s18` GPIO expander
  - a LoRa-related SPI device
- The active device tree has named symbols for:
  - RGB LED pins
  - PWM beeper pins
  - `gpio-keys`
  - `mcp23s18`
- The front-panel buttons are discoverable from the DT as `UP`, `DOWN`, `LEFT`, `RIGHT`, `SELECT`, and `BACK`
- Linux input devices currently enumerate:
  - `gpio-keys` as `event0`
  - PMIC power button as `event1`
  - `pwm-beeper` as `event2`

### Documentation added

- `docs/HardwareOutputs.md`
- `scripts/probe_badge_outputs.sh`

## 2026-04-14 (execution checklists)

Operational checklists and simulation planning were added.

### Documentation added

- `docs/SimulationPlan.md`
- `docs/BringUpChecklist.md`
- `docs/ReleaseChecklist.md`

### TODO updates

- Hardware bring-up checklist is now documented
- Release checklist is now documented

## 2026-04-14 (protocol fixtures)

Initial protocol fixture set added for simulation and future tests.

### Added fixtures

- `fixtures/protocol/path_tokens.json`
- `fixtures/protocol/info-response.json`
- `fixtures/protocol/start-request.json`
- `fixtures/protocol/move-request.json`
- `fixtures/protocol/move-response.json`
- `fixtures/protocol/end-request.json`

### Purpose

- provide a stable initial contract for endpoint-token mapping
- support future simulation work
- seed protocol and regression tests without waiting on hardware

## 2026-04-14 (preflight)

A demo preflight script was added.

## 2026-04-14 (rules CLI transport)

BadgeSnake transport work started inside `components/battlesnake-rules`.

### Changes in progress

- added CLI request-path support for:
  - `sim://` deterministic local players
  - `i2c://` BadgeSnake transport-shaped simulated players
- added focused tests for simulated metadata and move handling
- added local wrapper scripts and docs in the superproject to exercise the upstream CLI path

### Current board findings

- `go` is installed locally as `go1.24.4`
- first-run `go build` and `go test` on this board remain slow because the runtime and standard library compile locally
- `i2c-tools` are installed
- `modinfo i2c-stub` currently fails because this kernel package does not provide the `i2c-stub` module

### Implication

- the fastest executable path for CLI integration on this board today is `sim://`
- `i2c://` is useful now for transport-shape and token-mapping tests
- true `i2c-stub` bus emulation will require a kernel/module packaging change or a different test host
- initial `go run` attempts can also fail if they rely on the board's 104MB `/tmp` zram mount for compiler scratch space
- the local wrapper scripts now redirect Go temporary build output into a repo-local `.cache/` directory on the root filesystem
- the local wrapper scripts now also force `CGO_ENABLED=0` for faster and more reliable on-device CLI smoke runs
- the local smoke wrappers were reduced to a deterministic 3x3 zero-food game so they terminate quickly instead of running a long standard match

### Verified smoke runs

- `scripts/run_rules_cli_sim.sh`
  - completed on-device
  - 3x3 deterministic zero-food game
  - completed after 2 turns as a draw
- `scripts/run_rules_cli_i2c_sim.sh`
  - completed on-device
  - 3x3 deterministic zero-food game
  - completed after 2 turns with `ZeptoB` as winner

### Remaining gap

- `i2c://` is still transport-shaped simulation inside the CLI layer
- the next implementation step is to connect that scheme to a real I2C adapter path or a kernel-backed emulation path when available

## 2026-04-15 (bb-imager-rs and Zepto BSL)

Repository state changed under the user:

- large AM62L PDF documents were replaced by Markdown notes
- `components/bb-imager-rs` was added
- `rustup` was installed
- `components/battlesnake-rules` history was pushed to branch `badge-snake`

### Verified findings

- `components/battlesnake-rules` local HEAD `29c9fdf` is on branch `badge-snake` and matches `origin/badge-snake`
- `bb-imager-rs` already contains an MSPM0 I2C flashing path intended for Zepto-style targets
- `bb-flasher-mspm0/src/i2c.rs` hard-codes the BSL target address as `0x48`
- default Rust toolchain is now configured as `stable-aarch64-unknown-linux-gnu`
- read-based `i2cdetect` scans on `/dev/i2c-0` and `/dev/i2c-2` did not show an ACK at `0x48` during this pass

### New repo support

- `docs/components/bb-imager-rs.md`
- `scripts/probe_zepto_bsl.sh`

### Added

- `scripts/demo_preflight.sh`

### Purpose

- provide a single entry point to validate repo state, helper availability, core device exposure, and unresolved-input status before a demo run

## 2026-04-14 (fixture validation)

A protocol fixture validator was added and can run on the current image.

### Added

- `scripts/validate_protocol_fixtures.py`

### Purpose

- catch malformed or inconsistent protocol fixtures before simulation code is written

## 2026-04-14 (go skeleton)

Initial host-runtime code skeleton added.

### Added

- `go.mod`
- `cmd/badgesnake/main.go`
- `internal/protocol/tokens.go`
- `internal/protocol/frame.go`
- `internal/player/player.go`
- `internal/player/fixture_player.go`
- `internal/simtransport/memory.go`

### Notes

- The current image does not provide the Go toolchain, so this code was not compiled on-device
- The skeleton is intended to anchor future simulation work around:
  - path-token handling
  - frame encode/decode
  - a fixture-backed player abstraction
  - an in-memory transport adapter using upstream Battlesnake client models

## 2026-04-14 (runtime policies)

First-pass runtime behavior was further specified.

### Documentation added

- `docs/HotplugPolicy.md`
- `docs/DisplayStatusPlan.md`

### Decisions captured

- player presence must be validated before game start
- move timeout or disconnect should eliminate the affected player rather than crash the host
- disconnected players should not be reinserted into an active game
- the first-pass output priority is ePaper first, then minimal RGB LED and buzzer cues, with 7-segment support remaining provisional

## 2026-04-14 (gpio probing)

GPIO probing added more concrete output-path information.

### Findings

- `/dev/gpiochip0` through `/dev/gpiochip3` are present
- `gpioinfo` shows named front-panel button lines on `gpiochip1`
- `gpioinfo` shows a 16-line `gpiochip3` block consumed as `segment`
- the most likely interpretation is that the `mcp23s18` SPI expander backs the 7-segment display path

### Documentation added

- `scripts/probe_gpio_state.sh`

## 2026-04-14 (ui wrap-up)

UI/output exploration was intentionally narrowed.

### Direction recorded

- 7-segment implementation should follow the kernel-module style used by `components/vsx-examples/PocketBeagle-2/seven_segment`
- further generic UI exploration should stop until that path is actually exercised

## 2026-04-14 (go tests added)

Focused Go unit tests were added for the new runtime skeleton.

### Added

- `internal/protocol/frame_test.go`
- `internal/simtransport/memory_test.go`

### Notes

- A full `go test ./...` run on-device is currently slow because the board is compiling the Go runtime locally
- The target was narrowed from generic tree validation to focused package tests for protocol and in-memory transport
