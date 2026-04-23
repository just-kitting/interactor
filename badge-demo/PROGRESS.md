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

## 2026-04-16 (MicroBlocks I2C simulation start)

Started the MicroBlocks-side simulation path needed while real Zepto flashing
remains unreliable.

### Changes in `components/microblocks-smallvm`

- Added a new `i2ctarget` primitive set for the Linux VM:
  - `start`
  - `stop`
  - `isStarted`
  - `address`
  - `writeAvailable`
  - `receiveWrite`
  - `readRequested`
  - `reply`
- The Linux backend uses a spool directory at `/tmp/microblocks_i2c_target_sim`
  by default, overridable with `MICROBLOCKS_I2C_SIM_DIR`
- Boardie now reserves the same primitive names with stubs so the primitive
  surface does not have to change later
- Removed the earlier fake `requestedBytes` concept after review, because it does
  not match Zephyr-style target semantics where the target sees read initiation
  callbacks but not a controller-advertised total read count

### Superproject support added

- `examples/microblocks/I2C Target.ubl`
- `scripts/microblocks_i2c_sim.py`
- `scripts/test_microblocks_i2c_sim.sh`
- `docs/MicroBlocksI2CTargetSim.md`

### Intent

- let student MicroBlocks code poll and reply to simulated controller I2C requests
- keep the simulation path decoupled from the current MSPM0 BSL flashing blocker
- make it possible to exercise BadgeSnake request/response behavior before
  hardware firmware delivery is stable

## 2026-04-16 (MicroBlocks IDE runtime viability)

Checked whether the repository's bundled GP/MicroBlocks runtimes can be used
directly on BeagleBadge.

### Findings

- `components/microblocks-smallvm/gp/gp-raspberryPi` is a 32-bit `armhf` ELF
  expecting `/lib/ld-linux-armhf.so.3`
- that binary also depends on Raspberry Pi-specific userspace, including
  `libbcm_host.so`
- BeagleBadge does not provide `libbcm_host.so`, so installing only 32-bit glibc
  would not make the bundled Raspberry Pi executable usable here
- `components/microblocks-smallvm/gp/gp-linux64bit` is an `x86-64` binary and is
  not usable on BeagleBadge either
- the checked-in web app shell exists, but the generated browser runtime payloads
  are not currently present in the repo

### Implication

- the shortest promising path on BeagleBadge is a web-hosted MicroBlocks IDE
  backed by generated browser artifacts, not trying to run the bundled
  Raspberry Pi desktop binary

## 2026-04-17 (web-hosted MicroBlocks direction)

Moved the active MicroBlocks path to browser hosting.

### Changes in `components/microblocks-smallvm`

- Boardie now has a browser-side BadgeSnake `i2ctarget` queue model instead of
  only stubbed primitives
- the web app now includes `badgesnake-boardie.js`, exposing a simple helper API
  for driving Boardie I2C target transactions from the browser console
- the webapp service worker and manifest were adjusted to work from a normal
  served directory instead of assuming `/run/`

### Superproject support added

- `scripts/build_microblocks_web.sh`
- `scripts/serve_microblocks_web.py`
- `docs/WebMicroBlocks.md`

### Goal

- let the user open a hosted MicroBlocks IDE on BeagleBadge
- use Boardie as the first browser-executable target for BadgeSnake student code
- defer native GP runtime work until later

### Verification status

- `emscripten` and `nodejs` were installed successfully on BeagleBadge
- the first full web build now gets past the earlier hard failures:
  - removed unsupported `--memory-init-file`
  - added a writable repo-local Emscripten cache
  - restored `zlib` port use against that writable cache
  - made the Emscripten build script idempotent across retries
- Debian's `emcc` package still tried to invoke a missing `html-minifier-terser`
  helper when the build emitted `gp_wasm.html`; the build was corrected to emit
  `gp_wasm.js` directly so the hosted webapp no longer depends on that tool
- the Boardie web build also depended on an external `closure-compiler` binary;
  the local build was corrected to skip Closure so the hosted path builds with
  the stock Armbian Emscripten install
- the tracked web server now binds to `0.0.0.0` by default so the hosted IDE can
  be reached remotely without a local script hack
- a full end-to-end web build now completes on BeagleBadge:
  - `./scripts/build_microblocks_web.sh`
  - outputs verified in `chromeApp/webapp`:
    - `gp_wasm.js`
    - `gp_wasm.wasm`
    - `gp_wasm.data`
    - `boardie/run_boardie.js`
    - `boardie/run_boardie.wasm`
- hosted smoke test passed with:
  - `python3 scripts/serve_microblocks_web.py --host 127.0.0.1 --public-host 127.0.0.1 --port 18443`
  - `curl -kI https://127.0.0.1:18443/microblocks.html`
  - `curl -kI https://127.0.0.1:18443/gp_wasm.js`
  - `curl -kI https://127.0.0.1:18443/boardie/run_boardie.js`

### Remaining gap

- `i2c://` is still transport-shaped simulation inside the CLI layer

## 2026-04-18 (MicroBlocks I2C target cleanup and CLI bridge)

Cleaned up the student-facing I2C target block labels and added a Linux-side
request path for the hosted Boardie simulation.

### Student-facing changes

- renamed the importable MicroBlocks wrapper to `examples/microblocks/I2C Target.ubl`
- removed BadgeSnake/simulation wording from the block text
- dropped the wrapper for `isStarted`, which was not a useful student-facing
  target concept

### Hosted bridge changes

- `scripts/serve_microblocks_web.py` now exposes:
  - `POST /api/i2c/transaction`
  - `GET /api/i2c/next`
  - `POST /api/i2c/respond`
- the hosted `badgesnake-boardie.js` helper now polls that bridge and forwards
  controller transactions into the running Boardie I2C target queue
- added `scripts/web_i2c_transaction.py` as the Linux-side controller client
- added `scripts/test_web_i2c_bridge.sh` as a smoke test for the bridge

### Verification status

- `python3 -m py_compile scripts/serve_microblocks_web.py scripts/web_i2c_transaction.py`
- `bash -n scripts/test_web_i2c_bridge.sh`
- `node --check components/microblocks-smallvm/chromeApp/webapp/badgesnake-boardie.js`
- `node --check components/microblocks-smallvm/chromeApp/MicroBlocks/badgesnake-boardie.js`
- `./scripts/test_web_i2c_bridge.sh`

### Architectural note

- the hosted Boardie bridge is intentionally above the kernel I2C layer
- it preserves controller/target request semantics, but it does not create a
  Linux-visible `/dev/i2c-*` target endpoint
- using `i2ctransfer` directly against the browser simulation will require a new
  lower-level adapter or kernel-backed emulation path

## 2026-04-18 (i2c-stub suitability check)

Checked whether the hosted Boardie simulation should be replaced with `i2c-stub`.

### Findings

- kernel documentation describes `i2c-stub` as a simple memory-backed fake SMBus
  device for testing kernel client drivers
- that model does not provide the dynamic event/reply behavior needed for
  MicroBlocks target logic
- this BeagleBadge image does not currently expose `i2c-stub` or
  `i2c-slave-eeprom` as loadable modules
- `i2ctransfer` is installed and available, but there is no kernel-visible
  target backend here for it to talk to on behalf of Boardie

### Repo additions

- `docs/I2CKernelSimulation.md`
- `scripts/check_i2c_kernel_sim_support.sh`

### Conclusion

- the current hosted bridge remains the only working Boardie path on this image
- replacing it with `i2c-stub` would be incorrect even on a kernel that shipped
  that module

## 2026-04-18 (running Armbian kernel source mapping)

Mapped the running BeagleBadge kernel package back to the Armbian source tree.

### Verified live package metadata

- running kernel: `6.12.57-vendor-edge-k3`
- package: `linux-image-vendor-edge-k3` version `26.02.0-trunk`
- package source field: `linux-6.12.57`
- package build git revision:
  `da3c0f0a33ac00f7138c695a16d90301cf7ec02b`

### Matching Armbian build config

- board config: `config/boards/beaglebadge.conf`
- family config: `config/sources/families/k3-beagle.conf`
- BeagleBadge `vendor` / `vendor-rt` kernel source:
  - `https://github.com/beagleboard/linux`
  - `branch:v6.12.49-ti-arm64-r56`
- generic `k3` `vendor-edge` source is configured separately in
  `config/sources/families/k3.conf`:
  - `https://github.com/TexasInstruments/ti-linux-kernel`
  - `branch:ti-linux-6.12.y-cicd`

### Important conclusion

- the running image is named `vendor-edge-k3`, but the BeagleBadge family source
  config in the checked-in Armbian tree currently only overrides `vendor` and
  `vendor-rt`
- so the most defensible source-tree answer for the live image is:
  - family/package lane: Armbian `vendor-edge-k3`
  - source tree configured for generic `k3 vendor-edge`:
    `TexasInstruments/ti-linux-kernel` on `ti-linux-6.12.y-cicd`
- the exact packaged content on this installed image is additionally identified
  by package git revision `da3c0f0a33ac00f7138c695a16d90301cf7ec02b`

## 2026-04-23 (multi-controller I2C transport planning)

The user added the TI kernel source as a submodule and clarified that the real
transport problem is asynchronous messaging on a multi-controller I2C bus, not
a simple request/response controller-target link.

### Kernel-tree findings

- `components/ti-linux-kernel` contains:
  - `Documentation/i2c/slave-interface.rst`
  - `drivers/i2c/i2c-core-slave.c`
  - `drivers/i2c/i2c-slave-testunit.c`
  - `drivers/i2c/i2c-slave-eeprom.c`
  - `drivers/i2c/i2c-stub.c`
- the generic slave framework exists in this tree
- `config/kernel/linux-k3-vendor-edge.config` enables `CONFIG_I2C_SLAVE=y`
- `drivers/i2c/busses/i2c-omap.c` does not currently implement:
  - `reg_slave`
  - `unreg_slave`

### Implication

- AM62L host slave-mode support is the first kernel task
- `slave-testunit` is the right initial validation target
- `i2c-stub` remains unsuitable because it is for static SMBus client-driver
  testing, not a dynamic multi-controller message transport

### Repo additions

- `docs/I2CMultiControllerDriverPlan.md`
- `docs/components/ti-linux-kernel.md`

### Conclusion

- yes, the added TI kernel tree is a workable starting point for planning the
  driver
- the recommended sequence is:
  1. extend `i2c-omap.c` with slave support
  2. validate with `slave-testunit`
  3. add a BadgeSnake-specific framed backend
  4. only add a stream-style shim above that if MicroBlocks still needs it

## 2026-04-18 (hosted Boardie test program)

Added a concrete first program for the hosted I2C target bridge.

### Additions

- `examples/microblocks/I2C Target Echo Example.md`

### Intent

- give the user a specific Boardie-side MicroBlocks program to run
- document that `web_i2c_transaction.py` should talk to the same port as the
  running hosted IDE rather than starting a second server
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
- user confirmed the connected Zepto is on J6, which maps to `/dev/i2c-0`
- the pulled MSPM0 BSL guide confirms:
  - default I2C target address is `0x48`
  - interface autodetection times out after 10 seconds
  - command reception also times out and can require wakeup or re-entry
- an active BSL connection probe was run on J6 (`/dev/i2c-0`) with the exact `bb-imager-rs` connection packet for hundreds of attempts and still saw no ACK at `0x48`

### New repo support

- `docs/components/bb-imager-rs.md`
- `scripts/probe_zepto_bsl.sh`
- `scripts/probe_zepto_bsl_active.sh`

## 2026-04-15 (QWIIC DT mapping)

The user's device-tree suspicion was confirmed on the live system.

### Verified live mapping

- `/dev/i2c-0` maps to `main_i2c0` at `i2c@20000000`
- `/dev/i2c-2` maps to `wkup_i2c0` at `i2c@2b200000`
- `main_i2c1` at `i2c@20010000` exists in the badge DT but is `status = "disabled"`
- `main_i2c2` at `i2c@20020000` exists in the badge DT but is `status = "disabled"`

### Implication

- J6 and J7 are not currently exposed to Linux userspace on this image
- probing the Zepto on J6 through `/dev/i2c-0` was targeting the wrong controller

### Repo additions

- `docs/DeviceTreeI2C.md`
- `overlays/beaglebadge-qwiic-i2c.dtso`
- `scripts/validate_qwiic_i2c_overlay.sh`
- `scripts/install_qwiic_i2c_overlay.sh`

### Boot staging performed

- installed `/boot/dtb/ti/k3-am62l3-badge-qwiic-i2c.dtbo`
- backed up `/boot/uEnv.txt` to `/boot/uEnv.txt.bak.20260415T105230Z`
- updated `name_overlays` in `/boot/uEnv.txt` to include `ti/k3-am62l3-badge-qwiic-i2c.dtbo`
- reboot was then completed and Linux now exposes:
  - `/dev/i2c-1` -> `main_i2c1` / J6
  - `/dev/i2c-3` -> `main_i2c2` / J7
- the Zepto in MSPM0 BSL mode ACKs at `0x48` on `/dev/i2c-1`

## 2026-04-15 (Zepto flashing workflow)

The live Zepto flashing path is now concrete enough to script even though the Rust CLI build has not finished yet.

### Verified path

- `bb-imager-rs` should be built with the `zepto_i2c` feature
- the live destination for the connected Zepto is `/dev/i2c-1`
- the active BSL probe confirms the device ACKs at `0x48` on that bus

### Repo additions

- `docs/ZeptoFlashing.md`
- `scripts/build_bb_imager_cli_zepto_i2c.sh`
- `scripts/list_zepto_i2c_destinations.sh`

### Current limitation

- `cargo build -p bb-imager-cli --features zepto_i2c` was started on-device and progressed through dependency compilation
- the build was stopped before completion to avoid burning the entire turn on compile time alone

## 2026-04-15 (bb-imager-rs Zepto enumeration)

The built `bb-imager-cli` currently lists all `/dev/i2c-*` controllers for the Zepto target instead of only live MSPM0 BSL buses.

### Local patch direction

- `components/bb-imager-rs/bb-flasher-mspm0/src/i2c.rs` is being updated so `ports()` actively sends the MSPM0 BSL connection request
- only buses that return the `0x00` ACK should be exposed through `list-destinations zepto`

### Current status

- the source patch is in progress in the submodule
- narrow `cargo check` validation on-device is still dominated by first-pass dependency compilation

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

## 2026-04-15 (bb-imager-rs Zepto enumeration)

The `bb-imager-rs` Zepto I2C destination filtering patch was corrected after a live build failure.

### Findings

- `bb-flasher-mspm0/src/i2c.rs` cannot use `write_all` or `read_exact` on `LinuxI2CDevice`; the correct API is `I2CDevice::write` and `I2CDevice::read`
- `bb-imager-cli` had a second logic bug for `list-destinations zepto`: the `no_filter` flag was inverted only for the Zepto target, which forced unfiltered output even when active probing was patched in
- the live board still confirms the real MSPM0 BSL target only on J6 `/dev/i2c-1` at address `0x48`

### Actions

- fixed the active-probe implementation in `components/bb-imager-rs/bb-flasher-mspm0/src/i2c.rs`
- fixed the Zepto `list-destinations` flag handling in `components/bb-imager-rs/bb-imager-cli/src/main.rs`
- kicked off another on-device rebuild after the source fix

### Verification status

- source-level fix is in place
- full `bb-imager-cli` relink on this board is still the slow step, so the final filtered `list-destinations zepto` output should be rechecked against the rebuilt binary after link completion

## 2026-04-16 (Zepto PlatformIO blink scaffold)

The first native Zepto firmware scaffold now exists and builds locally with PlatformIO on the BeagleBadge.

### Added

- `firmware/zepto/platformio.ini`
- `firmware/zepto/boards/zeptomspm0l1117.json`
- `firmware/zepto/platforms/zepto-mspm0/`
- `firmware/zepto/linker/zeptomspm0l1117.ld`
- `firmware/zepto/src/startup_mspm0l1117.c`
- `firmware/zepto/src/zepto_board.c`
- `firmware/zepto/examples/baremetal/blink/src/main.c`
- `firmware/zepto/scripts/flash_zepto_bsl.sh`
- `scripts/build_zepto_blink.sh`
- `scripts/flash_zepto_blink.sh`

### Verified

- PlatformIO Core is available under `/root/.platformio/penv/bin/pio`
- PlatformIO successfully installed `toolchain-gccarmnoneeabi`
- `./scripts/build_zepto_blink.sh` succeeds on-device
- the current `blink` output is `364` bytes at `firmware/zepto/.pio/build/zepto_blink/firmware.bin`
- the Zepto still ACKs the MSPM0 BSL connection packet on `/dev/i2c-1`

### Current implementation details

- board ID is `zeptomspm0l1117`
- the first `blink` uses PA12 as the LED GPIO path with active-high behavior
- the build uses a minimal Cortex-M0+ startup file, linker script, and direct IOMUX/GPIO register writes
- the current upload path is routed through `bb-imager-cli flash zepto ... /dev/i2c-1`

### Current blocker

- `pio run -t upload` now reaches the Rust MSPM0 flasher but the flash still fails after BSL handshake with `Unknown error occured`
- this means PlatformIO project setup, image generation, and bus selection are working; the remaining problem is inside the current `bb-imager-rs` MSPM0 flash transaction path

## 2026-04-16 (Zepto flash retry)

The user retried Zepto flashing after re-entering the bootloader.

### Findings

- the expected repo-root helper path was missing; only `firmware/zepto/scripts/flash_zepto_bsl.sh` existed
- `./scripts/flash_zepto_blink.sh` still reaches `bb-imager-cli`
- the current live failure is now narrowed to `Remote I/O error (os error 121)` from the MSPM0 flashing path
- this failure still occurs after the user explicitly restarted the Zepto into BSL mode, so the current blocker is not explained solely by a stale timed-out bootloader session

### Follow-up

- added a repo-root `scripts/flash_zepto_bsl.sh` wrapper
- reset `docs/MissingInputs.md` back to the agreed no-open-questions state

## 2026-04-16 (live BSL state re-check)

After the user retried the flash path, the live Zepto was re-checked directly.

### Findings

- a direct `i2ctransfer` attempt of the MSPM0 BSL connection packet on `/dev/i2c-1` failed with `Remote I/O error`
- a fresh run of `scripts/probe_zepto_bsl_active.sh 1` then showed no ACK at `0x48` across the full retry window
- the current blocker is therefore the live Zepto no longer being in an active BSL session, not just the higher-level PlatformIO wrapper

### Next action

- re-enter the Zepto into BSL on J6 and respond through `docs/MissingInputs.md`, then continue immediately with live probing while the bootloader is still active

## 2026-04-16 (immediate re-probe after user `go`)

The Zepto was reprobed immediately after the user indicated that BSL had just been re-entered.

### Findings

- `i2cdetect -r -y 1` showed no device at all on the J6 bus
- `scripts/probe_zepto_bsl_active.sh 1` still got no ACK at `0x48` across the full retry window
- this means the current manual BOOT/RST sequence is not producing a Linux-visible I2C BSL target on demand, at least not reproducibly enough for live debugging

### Next action

- try probing while `BOOT` is still physically held high during and after reset, instead of releasing it before the probe window

## 2026-04-16 (correction after user-run probe)

The previous assumption that the Zepto was not reachable in BSL was too strong.

### Findings

- the user directly ran `scripts/probe_zepto_bsl_active.sh 1`
- that live user-run probe succeeded immediately with `ACK received: 0x00`
- the correct current fact is that the Zepto BSL is reachable on `/dev/i2c-1`
- the real problem is therefore reproducibility and sequencing around who runs the probe or flash and when, not a confirmed absence of the BSL target

### Follow-up

- cleared the stale blocking question from `docs/MissingInputs.md`
- corrected the repo record so future work treats J6 `/dev/i2c-1` as the active BSL path unless disproven by a fresh direct probe

## 2026-04-16 (combined-transfer hypothesis)

The current best hypothesis for the MSPM0 I2C flashing failure is that the backend is splitting BSL transactions too aggressively.

### Findings

- the working `scripts/probe_zepto_bsl_active.sh` path uses `i2ctransfer` with a combined write/read transaction
- the prior Rust `bb-flasher-mspm0` I2C backend buffered nothing and issued separate `write()` and `read()` syscalls
- the BSL request flow includes multi-part requests such as header, payload, CRC, then ACK, which are good candidates for breakage if the target expects them as one transaction boundary

### Changes

- updated `components/bb-imager-rs/bb-flasher-mspm0/src/i2c.rs` to buffer pending writes and execute them with `I2C_RDWR` combined transfer on the next read
- added `scripts/trace_zepto_bsl.py` to test `connect` and `get_device_info` without going through the full `bb-imager-cli` stack

## 2026-04-16 (BSL timing hypothesis)

The user provided a critical timing detail from live testing.

### Findings

- a user-run `scripts/probe_zepto_bsl_active.sh 1` did not get an ACK until attempt `13/50`
- that delay was caused by the user manually invoking BSL while the probe loop was already running, not by the BSL naturally taking `2.6s` to become ready
- the useful conclusion is only that a long-running waiter can catch manual BSL entry in-flight

### Changes

- updated `firmware/zepto/scripts/flash_zepto_bsl.sh` to:
  - wait for the BSL to ACK before calling `bb-imager-cli`
  - retry the flash command up to a configurable number of times
  - expose timing knobs through environment variables

## 2026-04-16 (flash wrapper bugfixes)

The first timing-aware flash wrapper exposed two shell-level bugs during live use.

### Findings

- `scripts/probe_zepto_bsl_active.sh` treated any successful `i2ctransfer` read byte as success, even when the returned byte was not the valid BSL ACK `0x00`
- `firmware/zepto/scripts/flash_zepto_bsl.sh` captured the wrong status code after a failed `bb-imager-cli` call because it relied on `if cmd; then ...; fi` and then read `$?`

### Changes

- updated `scripts/probe_zepto_bsl_active.sh` so only `0x00` counts as a valid ACK
- updated `firmware/zepto/scripts/flash_zepto_bsl.sh` to:
  - print a clearer prompt that it is waiting for the user to perform the BOOT/RST sequence
  - preserve the real `bb-imager-cli` exit status

## 2026-04-16 (restore known-good flasher path)

The earlier MSPM0 transport debugging inside `bb-imager-rs` was backed out.

### Findings

- `bb-imager-rs` commit `a0fb8954c60f9ef04c71e8fd9451912a2b22cf45` had already been debugged before this session
- the additional MSPM0 I2C transport rewrite and custom error-reporting changes were therefore higher-risk than the surrounding wrapper work
- the user-reported odd bytes during probing were better evidence that we were drifting away from the known-good path than evidence of a validated protocol insight

### Changes

- reverted the `bb-imager-rs` MSPM0 combined-transfer rewrite
- reverted the `bb-imager-rs` custom MSPM0 error-reporting patch
- kept the useful outer wrapper behavior in this repo: wait for manual BSL entry, then launch the known-good flasher path immediately

## 2026-04-16 (MSPM0 reserved flash regions)

The local MSPM0 documentation is now explicit enough to pin down which Zepto flash regions early firmware must avoid.

### Reserved regions

- `NONMAIN` configuration NVM at `0x41C0.0000` to `0x41C0.07FF`
- `FACTORY` constants at `0x41C4.0000` to `0x41C4.01FF`

### Why they must be avoided

- `NONMAIN` holds the BCR/BSL configuration that controls bootloader enablement, invoke behavior, integrity checks, passwords, and static write protection
- `FACTORY` holds factory-programmed calibration and identity data

### First-pass rule

- for `blink` and early I2C echo work, use MAIN flash only and leave `NONMAIN` and `FACTORY` untouched
