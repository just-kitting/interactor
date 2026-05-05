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
  3. add a BadgeSnake serial-to-I2C endpoint bridge
  4. only add explicit tty glue if endpoint devices are still insufficient

## 2026-04-23 (bridge direction clarified)

The user clarified that the desired end state is not a target-only driver.

### Direction change

- the kernel deliverable should be a serial-to-I2C endpoint bridge
- it should expose serial-style endpoints for initiators and targets
- internal framing still matters, but endpoint-style serial compatibility is now
  part of the intended driver shape rather than an optional afterthought

### What did not change

- AM62L `i2c-omap` slave support is still the first prerequisite
- `slave-testunit` is still the right first validation target
- the bridge still has to tolerate multi-controller arbitration loss and
  transaction-bounded wire semantics

## 2026-04-23 (target-mode bring-up staged)

Continued from planning into concrete AM62L target-mode bring-up prep.

### Armbian config changes staged

In `components/armbian-build`:

- `config/kernel/linux-k3-vendor-edge.config`
- `config/kernel/linux-k3-beagle-vendor.config`
- `config/kernel/linux-k3-beagle-vendor-rt.config`

all now include:

- `CONFIG_I2C_SLAVE_EEPROM=m`
- `CONFIG_I2C_SLAVE_TESTUNIT=m`

### Repo additions

- `docs/I2CSlaveBringup.md`
- `scripts/bringup_i2c_slave_testunit.sh`

### Current state

- the running kernel currently only exposes `CONFIG_I2C_SLAVE=y`
- it does not yet ship the `slave-testunit` backend
- `i2c-omap.c` still lacks `reg_slave` / `unreg_slave`

### Practical implication

- the next bootable validation target is now explicit:
  - rebuild/boot kernel with `I2C_SLAVE_TESTUNIT`
  - try `slave-testunit` instantiation on J6/J7
  - only then start implementing OMAP slave hooks or the BadgeSnake bridge,
    depending on what fails first

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

## 2026-04-23 (tighten Armbian kernel build handoff)

The user corrected the `components/armbian-build` branch and asked for a concrete host-side build path plus a preservation process before any new starting microSD image is discussed.

### Findings

- the immediate need for I2C target-mode bring-up is a rebuilt BeagleBadge `vendor-edge` kernel package set, not a fresh image
- `components/armbian-build` exposes the expected `kernel` artifact shortcut through `./compile.sh`
- the exact current target tuple is:
  - `BOARD=beaglebadge`
  - `BRANCH=vendor-edge`
  - `RELEASE=trixie`
- we had not yet committed a preservation workflow for the live board state before talking about replacing the current starting microSD image

### Changes

- added `scripts/build_beaglebadge_vendor_edge_kernel_x86_docker.sh` as the exact x86 host wrapper for the Armbian kernel build
- added `scripts/capture_beaglebadge_state.sh` to preserve current board and repo state before any image replacement
- added `docs/ArmbianKernelBuild.md` and `docs/SystemStatePreservation.md`
- updated `docs/components/armbian-build.md` and `docs/I2CSlaveBringup.md` to point at the new workflow

### Current policy

- build kernel packages first
- do not treat a fresh starting microSD image as the next default step
- before any image replacement, run the capture script on the live board and copy the resulting archive off-board

## 2026-04-27 (state capture and x86 kernel build completed)

The preservation and host-build steps are now complete.

### Findings

- the saved board-state archive now exists at `artifacts/state-capture/beaglebadge-state-20260427T115544Z.tar.gz`
- the x86 Docker build output has been copied back into `components/armbian-build/output/`
- the returned build includes:
  - `linux-image-vendor-edge-k3`
  - `linux-dtb-vendor-edge-k3`
  - `linux-headers-vendor-edge-k3`
  - `linux-libc-dev-vendor-edge-k3`
- the rebuilt packages currently use the same Debian package version string as the installed packages on the board: `26.02.0-trunk`

### Changes

- added `scripts/install_beaglebadge_vendor_edge_kernel_artifacts.sh` to reinstall the returned local `.deb` files on-device
- updated the kernel-build and slave-bring-up docs to use the reinstall path instead of a generic upgrade assumption

### Next step

- reinstall the returned kernel packages on the BeagleBadge
- reboot into the rebuilt kernel
- verify `i2c-slave-testunit` exists and continue AM62L slave-mode bring-up

## 2026-04-27 (kernel reinstall notes on live board)

The local reinstall of the returned `vendor-edge-k3` kernel packages completed, and the reported messages were reviewed.

### Findings

- `update-initramfs: Symlink failed ... moving ...` is expected on this image because `/boot` is FAT32 and cannot preserve normal Linux symlinks
- Armbian handled that correctly by renaming:
  - `/boot/uInitrd-6.12.57-vendor-edge-k3` to `/boot/uInitrd`
  - `/boot/vmlinuz-6.12.57-vendor-edge-k3` to `/boot/Image`
- the `_apt` unsandboxed note was caused by installing local `.deb` files directly from a root-only repo path, not by a package failure

### Changes

- updated `scripts/install_beaglebadge_vendor_edge_kernel_artifacts.sh` to stage the `.deb` files into `/var/tmp/badgesnake-kernel-debs/` before calling `apt-get install --reinstall`
- updated `docs/ArmbianKernelBuild.md` to document the expected FAT32 `/boot` behavior and the prior `_apt` warning

## 2026-04-27 (rebuilt kernel missing slave-testunit)

The first rebuilt `vendor-edge-k3` kernel install did not provide the expected slave test module.

### Findings

- after reboot, the live kernel still reported only:
  - `CONFIG_I2C_SLAVE=y`
- `modinfo i2c-slave-testunit` failed because the module was not present in the rebuilt package
- the returned x86-host build log shows it used `config/kernel/linux-k3-vendor-edge.config`
- on the corrected `components/armbian-build` branch, that config no longer contained the intended:
  - `CONFIG_I2C_SLAVE_EEPROM=m`
  - `CONFIG_I2C_SLAVE_TESTUNIT=m`
- therefore the first returned `vendor-edge-k3` artifacts were built from the wrong effective config and are not the right validation kernel

### Changes

- restored `CONFIG_I2C_SLAVE_EEPROM=m` and `CONFIG_I2C_SLAVE_TESTUNIT=m` in:
  - `components/armbian-build/config/kernel/linux-k3-vendor-edge.config`
  - `components/armbian-build/config/kernel/linux-k3-beagle-vendor.config`
  - `components/armbian-build/config/kernel/linux-k3-beagle-vendor-rt.config`

### Next step

- rebuild the kernel artifacts again on the x86 Docker host
- copy the corrected artifacts back into `components/armbian-build/output/`
- reinstall and reboot again before retrying `modinfo i2c-slave-testunit`

## 2026-04-27 (corrected artifact set restored)

The rebuilt x86-host artifacts with the restored K3 config have now been copied back into `components/armbian-build/output/debs/`.

### Findings

- `output/debs/` now contains both:
  - the older stale `Cd5e6...` artifact set
  - the corrected `C2876...` artifact set
- the corrected `linux-image-vendor-edge-k3_...C2876...deb` contains:
  - `i2c-slave-eeprom.ko`
  - `i2c-slave-testunit.ko`
- after reinstalling the corrected image package, those module files are present on disk under:
  - `/lib/modules/6.12.57-vendor-edge-k3/kernel/drivers/i2c/`

### Changes

- updated `scripts/install_beaglebadge_vendor_edge_kernel_artifacts.sh` so it no longer assumes a single artifact set in `output/debs/`
- the reinstall wrapper now prefers the image package whose contents include `i2c-slave-testunit.ko`, then picks the matching `dtb`, `headers`, and `libc-dev` packages from the same suffix

### Next step

- let the corrected reinstall finish
- reboot
- verify `modinfo i2c-slave-testunit`
- continue with `./scripts/bringup_i2c_slave_testunit.sh start 1 0x30`

## 2026-04-27 (corrected reinstall verified pre-reboot)

The corrected artifact set is now installed on disk and the missing QWIIC overlay was restored into `/boot`.

### Findings

- `modinfo i2c-slave-testunit` now succeeds by module name
- the running kernel package reports:
  - `Linux beaglebadge 6.12.57-vendor-edge-k3 #2 ...`
- the DTB package refresh removed `/boot/dtb/ti/k3-am62l3-badge-qwiic-i2c.dtbo`, which explains why only `i2c-0` and `i2c-2` were present and why the bring-up helper could not find `/sys/bus/i2c/devices/i2c-1/new_device`
- reinstalling the overlay restored:
  - `/boot/dtb/ti/k3-am62l3-badge-qwiic-i2c.dtbo`
  - `name_overlays=... ti/k3-am62l3-badge-qwiic-i2c.dtbo` in `/boot/uEnv.txt`

### Changes

- updated `scripts/install_beaglebadge_vendor_edge_kernel_artifacts.sh` to reapply the local QWIIC overlay after reinstalling the kernel DTB package

### Next step

- reboot so the restored QWIIC overlay takes effect again
- verify `i2c-1` and `i2c-3` reappear
- rerun `./scripts/bringup_i2c_slave_testunit.sh start 1 0x30`

## 2026-04-27 (adapter-side slave registration failure confirmed)

The BeagleBadge is now past the packaging, module, and overlay issues. The remaining failure is the expected adapter support gap.

### Findings

- after reboot, `i2c-1` and `i2c-3` reappeared
- `modinfo i2c-slave-testunit` succeeded normally
- instantiating `slave-testunit` on `i2c-1` created the slave-address node `1-1030`
- `dmesg` then reported:
  - `i2c_slave_register: not supported by adapter`
  - `probe with driver i2c-slave-testunit failed with error -95`
- this is the expected `i2c-omap` adapter limitation surfacing for real

### Changes

- updated `scripts/bringup_i2c_slave_testunit.sh` to use the correct slave-address sysfs node (`1-1030` for `0x30`) and to fail when the device node exists but the driver did not bind

### Meaning

- `new_device` alone is not enough proof of working slave mode on AM62L
- the project is now correctly narrowed to adding `reg_slave` / `unreg_slave` support and slave-event handling in `drivers/i2c/busses/i2c-omap.c`

## 2026-04-27 (first `i2c-omap` slave patch staged)

The project has now moved from reproducing the adapter limitation to staging the
first bus-driver patch in the TI kernel tree.

### Changes

- updated `components/ti-linux-kernel/drivers/i2c/busses/i2c-omap.c` to add:
  - `reg_slave()` / `unreg_slave()`
  - `I2C_FUNC_SLAVE` advertisement
  - a slave IRQ path that translates OMAP events into `i2c_slave_event()`
  - restoration of slave-listen mode after host master transfers

### Scope

- this patch is intended to get past the current `i2c_slave_register: not supported by adapter` failure
- it has not yet been compiled into a BeagleBadge test kernel in this session
- the next required step is another x86-host kernel rebuild and on-device reinstall

## 2026-04-28 (copied build check after staged `i2c-omap` patch)

The copied Armbian output was checked to see whether a new kernel build had arrived for the staged local `i2c-omap` patch.

### Findings

- there is a newer copied build session in `components/armbian-build/output/logs/`:
  - `85c46fcc-f113-4c36-881e-c0f1535e7e23`
- that build still reports:
  - `KERNELSOURCE='https://github.com/TexasInstruments/ti-linux-kernel'`
  - `KERNELBRANCH='branch:ti-linux-6.12.y-cicd'`
- the copied artifacts in `output/debs/` are still only the two already-known sets:
  - stale `Cd5e6...`
  - corrected module-enabled `C2876...`
- there is no third artifact suffix corresponding to a new post-`i2c-omap` build
- therefore no copied build is yet visible that proves the staged local `components/ti-linux-kernel` patch was compiled

### Meaning

- a newer build was copied back
- but it was not a new test build for the staged local `i2c-omap` patch

## 2026-04-30 (newer copied artifacts and Zepto control wiring)

Two new pieces of state were confirmed.

### Findings

- `components/armbian-build/output/debs/` now includes a newer artifact set with suffix `S3b4a...`
- the corresponding copied build log still reports:
  - `KERNELSOURCE='https://github.com/TexasInstruments/ti-linux-kernel'`
  - `KERNELBRANCH='branch:ti-linux-6.12.y-cicd'`
  - `kernel patching: 0 total patches; 0 applied; 0 with problems`
- so the `S3b4a...` build is newer than the earlier copied artifacts, but still does not prove the staged local `i2c-omap` patch was compiled
- the BeagleBadge Grove connector is now wired to the attached Zepto control pins:
  - this original pin-number note was later corrected
  - current corrected mapping is recorded below and in `docs/ZeptoControlWiring.md`

### Changes

- added `docs/ZeptoControlWiring.md`
- updated flashing docs to point at the new control-wiring record

### Next step

- identify the Linux GPIO mapping and mux state for the corrected Grove pins
- add host-side BSL/RST automation once that mapping is known

## 2026-04-30 (Armbian patch integration for AM62L slave-mode test kernel)

The next BeagleBadge `vendor-edge` rebuild has been wired to carry the staged AM62L `i2c-omap` slave-mode work through Armbian's kernel patch path instead of depending on a separate local TI-kernel source tree.

### Changes

- added:
  - `components/armbian-build/patch/kernel/archive/k3-6.12/0001-Stage-OMAP-I2C-slave-registration-support.patch`
  - `components/armbian-build/patch/kernel/archive/k3-6.12/0002-Fix-OMAP-slave-helper-declaration-order.patch`
- updated:
  - `scripts/build_beaglebadge_vendor_edge_kernel_x86_docker.sh`
  - `docs/ArmbianKernelBuild.md`
  - `docs/I2CSlaveBringup.md`
  - `docs/components/armbian-build.md`
  - `docs/components/ti-linux-kernel.md`

### Meaning

- the x86 host build now has one concrete patch input for AM62L target-mode bring-up
- the build wrapper fails fast if that patch file is missing
- the next returned artifact set should only be trusted if the host build log shows kernel patching from `archive/k3-6.12`

### Next step

- rerun `./scripts/build_beaglebadge_vendor_edge_kernel_x86_docker.sh` on the x86 Docker host
- copy the returned artifacts back into `components/armbian-build/output/`
- reinstall on BeagleBadge and rerun `slave-testunit` binding tests

## 2026-04-30 (first AM62L slave-patch build failure and Grove pin correction)

The first x86-host rebuild that carried the AM62L `i2c-omap` slave patch failed in compile, but it usefully proved that the Armbian patch path is now active.

### Findings

- copied build log:
  - `components/armbian-build/output/logs/log-kernel-7afe2c77-a105-4db4-954a-6d5cef614223.log`
- the log now shows:
  - `Summary: kernel patching: 1 total patches; 1 applied; 0 with problems`
- the actual compile failure is narrow and local to the staged patch:
  - `drivers/i2c/busses/i2c-omap.c:812:17: error: implicit declaration of function 'omap_i2c_set_master_mode'`
  - `drivers/i2c/busses/i2c-omap.c:838:17: error: implicit declaration of function 'omap_i2c_restore_slave_listen'`
  - later duplicate-definition diagnostics confirm the problem is missing forward declarations, not a broader design failure

### Wiring correction

- the temporary Grove-to-Zepto control wiring was corrected by the user:
  - Grove pin 1, yellow -> `BSL`
  - Grove pin 2, white -> `RST`
  - earlier pin-3/pin-4 notes were backwards
- user also noted package balls:
  - Grove pin 1 / `BSL` -> `G23`
  - Grove pin 2 / `RST` -> `G22`
- Linux GPIO line names and final mux state are still not captured

### Next step

- add the missing helper forward declarations in `i2c-omap.c`
- refresh the Armbian patch copy from the corrected source
- rerun the x86 host kernel build

## 2026-04-30 (clean two-patch AM62L build artifacts returned)

The rebuilt Armbian artifacts copied back after the source and patch-series fixes are the first clean AM62L slave-mode test set.

### Findings

- copied build log:
  - `components/armbian-build/output/logs/log-kernel-1b53c454-a989-4c66-a5e6-e8b706e4bc05.log`
- the corresponding summary confirms:
  - `kernel patching: 2 total patches; 2 applied; 0 with problems`
- the new returned artifact suffix is:
  - `P5507`
- older `P0000` and earlier stale artifact sets are still present in `output/debs/`

### Changes

- updated `scripts/install_beaglebadge_vendor_edge_kernel_artifacts.sh` to choose the newest artifact set that contains `i2c-slave-testunit.ko`
- recorded a new generic hardware TODO to document Grove line names and pinmux control for UART/I2C/GPIO roles, not only Zepto control

### Next step

- reinstall the `P5507` BeagleBadge `vendor-edge-k3` kernel artifacts on the board
- reboot and rerun the `slave-testunit` binding test on `i2c-1`

## 2026-04-30 (installed `P5507` AM62L slave-test kernel on BeagleBadge)

The first clean post-fix AM62L slave-test kernel artifact set has now been installed on the live board.

### Findings

- selected artifact suffix:
  - `P5507`
- selected packages:
  - `linux-image-vendor-edge-k3_..._P5507_...`
  - `linux-dtb-vendor-edge-k3_..._P5507_...`
  - `linux-headers-vendor-edge-k3_..._P5507_...`
  - `linux-libc-dev-vendor-edge-k3_..._P5507_...`
- install completed successfully
- expected install-time notes were observed:
  - FAT32 `/boot` symlink fallback for `uInitrd`
  - DTB refresh followed by local QWIIC overlay reinstall
  - non-fatal firmware warnings during `update-initramfs`
- `/boot/uEnv.txt` was backed up to:
  - `/boot/uEnv.txt.bak.20260430T144428Z`

### Changes

- `scripts/install_beaglebadge_vendor_edge_kernel_artifacts.sh` now prefers the newest returned artifact set that contains `i2c-slave-testunit.ko`
- the live system was reinstalled with the `P5507` artifact set
- the QWIIC overlay was reapplied automatically after the DTB package refresh

### Next step

- reboot the BeagleBadge into the installed `P5507` kernel
- run:
  - `uname -a`
  - `modinfo i2c-slave-testunit`
  - `ls /sys/bus/i2c/devices | grep '^i2c-[13]$'`
  - `./scripts/bringup_i2c_slave_testunit.sh start 1 0x30`

## 2026-04-30 (runtime validation after booting `P5507`)

The board has now been rebooted into the installed `P5507` kernel and the first live target-mode validation has been rerun.

### Findings

- running kernel:
  - `Linux beaglebadge 6.12.57-vendor-edge-k3 #4 SMP PREEMPT Thu Apr 30 08:25:00 UTC 2026 aarch64 GNU/Linux`
- `modinfo i2c-slave-testunit` still works
- QWIIC adapters are present after reboot:
  - `i2c-1`
  - `i2c-3`
- `./scripts/bringup_i2c_slave_testunit.sh start 1 0x30` now succeeds
- `./scripts/bringup_i2c_slave_testunit.sh status 1 0x30` reports:
  - `present: 1-1030`
  - `bound: yes`
  - `slave-testunit`
- `dmesg` no longer shows:
  - `i2c_slave_register: not supported by adapter`
- forced same-adapter traffic still fails:
  - `i2ctransfer -f -y 1 r1@0x30`
  - `i2ctransfer -f -y -b 1 w3@0x30 4 0 0 r32`
  - both currently time out

### Meaning

- the AM62L `i2c-omap` work has crossed the first important threshold:
  - slave registration and driver binding now work on live hardware
- the remaining failure is in transaction completion after bind, not in adapter capability advertisement or slave registration

### Next step

- inspect why the bound slave target is not completing controller-originated transactions on `i2c-1`
- add the next round of `i2c-omap` instrumentation or fixes based on that timeout behavior

## 2026-04-30 (next `i2c-omap` follow-up for slave TX timeout)

The first live `P5507` validation narrowed the remaining AM62L problem to transaction completion after slave bind. The next staged driver change is targeted specifically at slave-transmit startup.

### Findings

- the local AM62L TRM notes say target-transmit mode can hold SCL low while software intervention is required on `XUDF`
- TI forum guidance and examples show `XUDF` appearing together with slave-transmit startup events
- the previous patch series did not enable `XUDF` interrupts in slave mode and did not service `XUDF` as part of the slave TX path
- the previous slave-listen restore path also did not explicitly reset FIFO state for slave operation

### Changes

- updated `components/ti-linux-kernel/drivers/i2c/busses/i2c-omap.c` to:
  - enable `XUDF` interrupts in slave mode
  - service `XUDF` and `XRDY` through a shared slave-TX helper
  - reset FIFO state and keep 1-byte thresholds when returning to slave-listen mode
- refreshed the Armbian `archive/k3-6.12` patch series to include:
  - `0003-Handle-slave-TX-underflow-on-OMAP-I2C.patch`

### Next step

- rebuild the BeagleBadge `vendor-edge` kernel packages again on the x86 Docker host
- copy the returned artifacts back into `components/armbian-build/output/`
- reinstall, reboot, and repeat the `i2ctransfer -f -y 1 r1@0x30` validation

## 2026-04-30 (installed `P024c` slave-TX follow-up kernel)

The next AM62L slave-mode follow-up build has now been copied back, selected, and installed on the live board.

### Findings

- copied build log:
  - `components/armbian-build/output/logs/log-kernel-a7a0cecc-b133-4f5e-86d2-3dc2f1235eea.log`
- build summary:
  - `kernel patching: 3 total patches; 3 applied; 0 with problems`
- selected artifact suffix:
  - `P024c`
- selected packages:
  - `linux-image-vendor-edge-k3_..._P024c_...`
  - `linux-dtb-vendor-edge-k3_..._P024c_...`
  - `linux-headers-vendor-edge-k3_..._P024c_...`
  - `linux-libc-dev-vendor-edge-k3_..._P024c_...`

### Changes

- installed the `P024c` kernel package set on the BeagleBadge
- reapplied the local QWIIC overlay after the DTB package refresh
- backed up `/boot/uEnv.txt` to:
  - `/boot/uEnv.txt.bak.20260430T210620Z`

### Meaning

- the next runtime test will be the first one with all three staged `i2c-omap` patches:
  - slave registration support
  - declaration-order fix
  - slave TX underflow/FIFO follow-up

### Next step

- reboot the board
- rerun:
  - `uname -a`
  - `./scripts/bringup_i2c_slave_testunit.sh start 1 0x30`
  - `./scripts/bringup_i2c_slave_testunit.sh status 1 0x30`
  - `i2ctransfer -f -y 1 r1@0x30`

## 2026-05-01 (validated `P024c` runtime and staged 1-byte FIFO threshold follow-up)

The `P024c` kernel is now booted and validated on the live BeagleBadge. Binding works, but the first forced same-adapter read still times out, so the next follow-up is narrowed to slave FIFO threshold programming.

### Findings

- running kernel after reboot:
  - `Linux beaglebadge 6.12.57-vendor-edge-k3 #5 SMP PREEMPT Thu Apr 30 08:25:00 UTC 2026 aarch64 GNU/Linux`
- `modinfo i2c-slave-testunit` still works
- J6 / J7 adapters are present after reboot:
  - `i2c-1`
  - `i2c-3`
- `./scripts/bringup_i2c_slave_testunit.sh start 1 0x30` reports:
  - `slave-testunit already present at 1-1030`
- `./scripts/bringup_i2c_slave_testunit.sh status 1 0x30` reports:
  - `present: 1-1030`
  - `bound: yes`
  - `slave-testunit`
- forced same-adapter read still fails:
  - `i2ctransfer -f -y 1 r1@0x30`
  - `Error: Sending messages failed: Connection timed out`
- adapter IRQ activity is present during the timed-out read:
  - IRQ line: `288`
  - source: `20010000.i2c`
  - observed count after the timed-out read: `4`
- `dmesg` after the timeout still shows no new `i2c-omap` slave transaction diagnostics beyond device instantiation:
  - `i2c i2c-1: new_device: Instantiated device slave-testunit at 0x30`

### Meaning

- the `P024c` patch series is sufficient for slave registration and bind on AM62L
- the remaining runtime gap is in active slave transaction service, not adapter registration
- the most concrete remaining mismatch with the TRM is slave FIFO threshold programming:
  - AM62L target-transmit requires TX threshold `1`
  - the previous slave-listen helper only cleared FIFOs and did not actually program the TX/RX thresholds to `1`

### Changes

- committed a follow-up TI-kernel change in `components/ti-linux-kernel`:
  - `38e256a46` `Program 1-byte FIFO thresholds in slave listen mode`
- mirrored that into the Armbian patch series as:
  - `components/armbian-build/patch/kernel/archive/k3-6.12/0004-Program-1-byte-FIFO-thresholds-in-slave-listen-mode.patch`

### Next step

- rebuild the BeagleBadge `vendor-edge` kernel package set with the four-patch `k3-6.12` series
- reinstall the returned artifacts
- reboot and rerun:
  - `./scripts/bringup_i2c_slave_testunit.sh start 1 0x30`
  - `i2ctransfer -f -y 1 r1@0x30`

## 2026-05-01 (returned `P9d8b` build and pinned reinstall path)

The next four-patch host build was returned successfully, but the on-device reinstall helper initially selected the older `P024c` package set because multiple artifact sets coexist under `components/armbian-build/output/debs`.

### Findings

- returned build UUID:
  - `f13b89a7-6f9f-4da6-b75f-1856bd7715dd`
- returned build summary:
  - `kernel patching: 4 total patches; 4 applied; 0 with problems`
- returned artifact suffix:
  - `P9d8b`
- the default reinstall helper selected:
  - `P024c`
- root cause:
  - the helper ranked matching `.deb` files by filesystem modification time, which is not a reliable proxy for the newest copied build set

### Changes

- updated `scripts/install_beaglebadge_vendor_edge_kernel_artifacts.sh` to accept:
  - `BADGESNAKE_BUILD_SUFFIX='<full package suffix>.deb'`
- documented the explicit-suffix path in `docs/ArmbianKernelBuild.md`
- launched the pinned reinstall using:
  - `BADGESNAKE_BUILD_SUFFIX='6.12.57-S3b4a-D0000-P9d8b-C2876Hb496-HK01ba-Vc222-Be8e3-R448a.deb' ./scripts/install_beaglebadge_vendor_edge_kernel_artifacts.sh`

### Changes

- the pinned `P9d8b` reinstall completed on the live BeagleBadge
- `dpkg.log` shows the image package finished at:
  - `2026-05-01 14:21:55 UTC`
- the local QWIIC overlay files were reapplied afterward:
  - `/boot/dtb/ti/k3-am62l3-badge-qwiic-i2c.dtbo`
  - `/boot/uEnv.txt`

### Next step

- reboot into the pinned `P9d8b` kernel
- rerun:
  - `uname -a`
  - `./scripts/bringup_i2c_slave_testunit.sh start 1 0x30`
  - `i2ctransfer -f -y 1 r1@0x30`

## 2026-05-01 (validated `P9d8b` runtime)

The four-patch `P9d8b` kernel is now booted on the live board. The 1-byte FIFO threshold follow-up did not resolve the read timeout.

### Findings

- running kernel after reboot:
  - `Linux beaglebadge 6.12.57-vendor-edge-k3 #6 SMP PREEMPT Thu Apr 30 08:25:00 UTC 2026 aarch64 GNU/Linux`
- J6 / J7 adapters are present after reboot:
  - `i2c-1`
  - `i2c-3`
- `modinfo i2c-slave-testunit` still works
- `./scripts/bringup_i2c_slave_testunit.sh start 1 0x30` succeeds
- `./scripts/bringup_i2c_slave_testunit.sh status 1 0x30` reports:
  - `present: 1-1030`
  - `bound: yes`
  - `slave-testunit`
- forced same-adapter read still fails:
  - `i2ctransfer -f -y 1 r1@0x30`
  - `Error: Sending messages failed: Connection timed out`
- adapter IRQ activity is still present during the timed-out read:
  - before: `1`
  - after: `2`
  - source: `20010000.i2c`
- `dmesg` still shows no new `i2c-omap` transaction diagnostics beyond:
  - `i2c i2c-1: new_device: Instantiated device slave-testunit at 0x30`

### Meaning

- the 1-byte slave FIFO threshold follow-up did not change the externally visible failure mode
- AM62L slave registration and binding remain good
- the remaining gap is still in the first active slave transaction path after address match

### Next step

- instrument or trace the first live AM62L slave transaction status path on J6 / `i2c-1`
- determine whether the controller is stalling at:
  - `AAS`
  - `XRDY`
  - `XUDF`
  - another status transition not currently surfaced in `dmesg`

## 2026-05-01 (staged 5-patch diagnostic kernel)

The next follow-up is now a diagnostic kernel build that logs the live AM62L slave IRQ/status path during J6 target traffic.

### Changes

- committed a TI-kernel instrumentation patch in `components/ti-linux-kernel`:
  - `d1fdd7a4f` `Instrument OMAP slave transaction state`
- mirrored that into the Armbian patch series as:
  - `components/armbian-build/patch/kernel/archive/k3-6.12/0005-Instrument-OMAP-slave-transaction-state.patch`
- updated the x86 build wrapper so it now expects the full five-patch series

### Meaning

- the next host build is not trying to "fix" the timeout yet
- it is specifically intended to surface:
  - enabled slave IRQ status bits
  - `I2C_CON`
  - `I2C_BUFSTAT`
  - `slave_read`
  - current threshold
- those logs should make it clear whether the first live target transaction stalls at:
  - `AAS`
  - `XRDY`
  - `XUDF`
  - or another unexpected status combination

### Next step

- rebuild the BeagleBadge `vendor-edge` kernel package set with the five-patch `k3-6.12` series
- copy the returned artifacts back
- reinstall and reboot into that diagnostic kernel
- capture the new `i2c-omap` log lines while reproducing:
  - `./scripts/bringup_i2c_slave_testunit.sh start 1 0x30`
  - `i2ctransfer -f -y 1 r1@0x30`

## 2026-05-04 (installed `P0380` diagnostic kernel)

The five-patch diagnostic build has been returned and installed on the live BeagleBadge.

### Findings

- returned build UUID:
  - `aeec0475-c697-43f7-b4c7-bee8e8e4539c`
- returned build summary:
  - `kernel patching: 5 total patches; 5 applied; 0 with problems`
- selected artifact suffix:
  - `P0380`
- the diagnostic reinstall was pinned explicitly with:
  - `BADGESNAKE_BUILD_SUFFIX='6.12.57-S3b4a-D0000-P0380-C2876Hb496-HK01ba-Vc222-Be8e3-R448a.deb'`
- `update-initramfs` completed normally with the usual FAT32 `/boot` rename fallback
- the local QWIIC overlay was reinstalled automatically afterward:
  - `/boot/dtb/ti/k3-am62l3-badge-qwiic-i2c.dtbo`
  - `/boot/uEnv.txt`
- `/boot/uEnv.txt` now again includes:
  - `name_overlays=ti/k3-am62l3-badge-eink-gdey042t81.dtbo ti/k3-am62l3-badge-qwiic-i2c.dtbo`

### Meaning

- the diagnostic kernel is ready to boot
- the next runtime pass should finally surface `i2c-omap` slave-path status in `dmesg` during the timed-out read

### Next step

- reboot into `P0380`
- rerun:
  - `uname -a`
  - `./scripts/bringup_i2c_slave_testunit.sh start 1 0x30`
  - `i2ctransfer -f -y 1 r1@0x30`
  - `dmesg | tail -n 80`

## 2026-05-04 (validated `P0380` diagnostic logs and staged ISR-branch follow-up)

The first diagnostic kernel produced usable log data. It still does not show any slave IRQ processing beyond listen-mode entry.

### Findings

- running kernel after reboot:
  - `Linux beaglebadge 6.12.57-vendor-edge-k3 #7 SMP PREEMPT Thu Apr 30 08:25:00 UTC 2026 aarch64 GNU/Linux`
- J6 / J7 adapters are present after reboot:
  - `i2c-1`
  - `i2c-3`
- `slave-testunit` still instantiates and binds on `i2c-1`
- first reproduced read after fresh bind failed with:
  - `Error: Sending messages failed: Remote I/O error`
- the adapter IRQ count still increments during the failing read:
  - before: `0`
  - after: `1`
- the only new diagnostic log line from `i2c-omap` was:
  - `omap_i2c 20010000.i2c: slave listen stat=0x00 con=0x8002 bufstat=0x8001 read=0 threshold=1`
- there was still no `slave irq` log line during the transaction
- the client node remains present afterward:
  - `present: 1-1030`
  - `bound: yes`
  - `slave-testunit`

### Meaning

- the current diagnostic point is too low in the stack
- we know the adapter entered listen mode, but we still do not know whether the failing self-read is being serviced:
  - in the slave IRQ path
  - or in the master IRQ path after `omap_i2c_set_master_mode()`

### Changes

- committed a second TI-kernel diagnostic patch in `components/ti-linux-kernel`:
  - `c5a94159b` `Trace OMAP master/slave ISR branch selection`
- mirrored that into the Armbian patch series as:
  - `components/armbian-build/patch/kernel/archive/k3-6.12/0006-Trace-OMAP-master-slave-ISR-branch-selection.patch`
- updated the x86 build wrapper to expect the six-patch series

### Next step

- rebuild the BeagleBadge `vendor-edge` kernel package set with the six-patch `k3-6.12` series
- copy the returned artifacts back
- reinstall and reboot into that branch-selection diagnostic kernel
- reproduce:
  - `./scripts/bringup_i2c_slave_testunit.sh start 1 0x30`
  - `i2ctransfer -f -y 1 r1@0x30`
  - `dmesg | tail -n 80`

## 2026-05-04 (fixed six-patch compile break)

The first host build of the six-patch diagnostic kernel failed at compile time, but the failure was narrow and local.

### Findings

- failed build UUID:
  - `7051ba70-863c-4ecf-98ef-d796f6382523`
- patching still succeeded:
  - `kernel patching: 6 total patches; 6 applied; 0 with problems`
- compile failure in `drivers/i2c/busses/i2c-omap.c`:
  - implicit declaration of `omap_i2c_slave_log_state`
  - followed by conflicting/static declaration errors
- root cause:
  - the ISR-branch diagnostic patch called `omap_i2c_slave_log_state()` before its definition

### Changes

- committed a TI-kernel fix in `components/ti-linux-kernel`:
  - `2e86d4d5a` `Fix OMAP ISR diagnostics declaration order`
- regenerated the Armbian sixth patch as:
  - `components/armbian-build/patch/kernel/archive/k3-6.12/0006-Fix-OMAP-ISR-diagnostics-declaration-order.patch`
- updated the x86 build wrapper to expect that regenerated `0006`

### Next step

- rerun the same x86 host build:
  - `./scripts/build_beaglebadge_vendor_edge_kernel_x86_docker.sh`
- copy the returned artifacts back
- continue with the six-patch diagnostic reinstall on the BeagleBadge

## 2026-05-04 (corrected stale regenerated `0006` patch payload)

The next host build still failed with the same compile error even after the TI-kernel source fix. The cause was not new code; it was a stale Armbian patch payload.

### Findings

- failed build UUID:
  - `308cf8d8-5664-47a2-bd0b-e35424d40f92`
- patching still succeeded:
  - `kernel patching: 6 total patches; 6 applied; 0 with problems`
- compile failure was unchanged:
  - implicit declaration of `omap_i2c_slave_log_state`
- root cause:
  - `components/armbian-build/patch/kernel/archive/k3-6.12/0006-Fix-OMAP-ISR-diagnostics-declaration-order.patch`
    still contained the old branch-selection patch payload instead of the 2-line declaration-order fix

### Changes

- regenerated `0006-Fix-OMAP-ISR-diagnostics-declaration-order.patch` from the actual TI-kernel commit:
  - `2e86d4d5a`
- committed the corrected Armbian patch payload in `components/armbian-build`:
  - `1be756694` `Correct OMAP ISR declaration fix patch payload`
- locally preflighted the corrected source on the BeagleBadge by building:
  - `make -C components/ti-linux-kernel O=/usr/src/linux-headers-$(uname -r) CC=gcc HOSTCC=gcc drivers/i2c/busses/i2c-omap.o`
  - result: `drivers/i2c/busses/i2c-omap.o` compiled successfully

### Next step

- rerun the same x86 host build again
- the expected `0006` diffstat is now:
  - `(+2/-0)[1M]`

## 2026-05-04 (installed `P803a` six-patch diagnostic kernel)

The corrected six-patch diagnostic build has now been returned and installed on the live BeagleBadge.

### Findings

- returned build UUID:
  - `d29eab67-4837-464d-8f22-14ebbc1062c6`
- returned build summary:
  - `kernel patching: 6 total patches; 6 applied; 0 with problems`
- selected artifact suffix:
  - `P803a`
- the diagnostic reinstall was pinned explicitly with:
  - `BADGESNAKE_BUILD_SUFFIX='6.12.57-S3b4a-D0000-P803a-C2876Hb496-HK01ba-Vc222-Be8e3-R448a.deb'`
- `update-initramfs` completed normally with the usual FAT32 `/boot` rename fallback
- the local QWIIC overlay was reinstalled automatically afterward:
  - `/boot/dtb/ti/k3-am62l3-badge-qwiic-i2c.dtbo`
  - `/boot/uEnv.txt`
- `/boot/uEnv.txt` still includes:
  - `name_overlays=ti/k3-am62l3-badge-eink-gdey042t81.dtbo ti/k3-am62l3-badge-qwiic-i2c.dtbo`

### Meaning

- the six-patch ISR-branch diagnostic kernel is ready to boot
- the next runtime pass should tell us whether the self-read stays on the master ISR path instead of entering the slave IRQ path

### Next step

- reboot into `P803a`
- rerun:
  - `uname -a`
  - `./scripts/bringup_i2c_slave_testunit.sh start 1 0x30`
  - `i2ctransfer -f -y 1 r1@0x30`
  - `dmesg | tail -n 80`

## 2026-05-04 (validated `P803a` ISR-branch diagnostics and staged slave-lifetime follow-up)

The six-patch ISR-branch diagnostic kernel still did not produce `master-enter` or `isr-master` logs during the failing self-read.

### Findings

- running kernel after reboot:
  - `Linux beaglebadge 6.12.57-vendor-edge-k3 #8 SMP PREEMPT Thu Apr 30 08:25:00 UTC 2026 aarch64 GNU/Linux`
- J6 / J7 adapters are present after reboot:
  - `i2c-1`
  - `i2c-3`
- fresh `slave-testunit` instantiation still reports success
- first reproduced self-read after fresh bind failed with:
  - `Error: Sending messages failed: Remote I/O error`
- the adapter IRQ count still increments during the failing read:
  - before: `0`
  - after: `1`
- the only new diagnostic line remains:
  - `omap_i2c 20010000.i2c: slave listen stat=0x00 con=0x8002 bufstat=0x8001 read=0 threshold=1`
- there was still no:
  - `master-enter`
  - `isr-master`
  - `slave irq`
- after a short delay, the client node is still present and bound:
  - `present: 1-1030`
  - `bound: yes`
  - `slave-testunit`

### Meaning

- the failing self-read is still not reaching any of the new xfer/ISR branch diagnostics
- the next most plausible gap is the slave registration lifetime itself:
  - whether `reg_slave` is really sticking
  - whether `unreg_slave` is being called unexpectedly
  - whether `xfer_common()` is starting with no slave pointer at all

### Changes

- committed a TI-kernel follow-up in `components/ti-linux-kernel`:
  - `a0413f647` `Trace OMAP slave registration lifetime`
- mirrored that into the Armbian patch series as:
  - `components/armbian-build/patch/kernel/archive/k3-6.12/0007-Trace-OMAP-slave-registration-lifetime.patch`
- updated the x86 build wrapper to expect the seven-patch series

### Next step

- rebuild the BeagleBadge `vendor-edge` kernel package set with the seven-patch `k3-6.12` series
- copy the returned artifacts back
- install and boot that slave-lifetime diagnostic kernel
- reproduce:
  - `./scripts/bringup_i2c_slave_testunit.sh start 1 0x30`
  - `i2ctransfer -f -y 1 r1@0x30`
  - `dmesg | tail -n 80`

## 2026-05-04 (installed `Pb92b` slave-lifetime diagnostic kernel)

The seven-patch diagnostic build completed and has now been installed on the live BeagleBadge.

### Findings

- returned build UUID:
  - `5e4b0975-693d-46fa-a009-9756e65b7d75`
- build summary:
  - `kernel patching: 7 total patches; 7 applied; 1 with problems; 1 needs_rebase`
- important detail:
  - Armbian marked `0007-Trace-OMAP-slave-registration-lifetime` as `needs_rebase`, but the build still completed and produced installable packages
- selected artifact suffix:
  - `Pb92b`
- the diagnostic reinstall was pinned explicitly with:
  - `BADGESNAKE_BUILD_SUFFIX='6.12.57-S3b4a-D0000-Pb92b-C2876Hb496-HK01ba-Vc222-Be8e3-R448a.deb'`
- `update-initramfs` completed normally with the usual FAT32 `/boot` rename fallback
- the local QWIIC overlay was reinstalled automatically afterward:
  - `/boot/dtb/ti/k3-am62l3-badge-qwiic-i2c.dtbo`
  - `/boot/uEnv.txt`
- `/boot/uEnv.txt` still includes:
  - `name_overlays=ti/k3-am62l3-badge-eink-gdey042t81.dtbo ti/k3-am62l3-badge-qwiic-i2c.dtbo`

### Meaning

- the seven-patch slave-lifetime diagnostic kernel is ready to boot
- the next runtime pass should tell us whether:
  - `reg_slave()` logs
  - `unreg_slave()` runs unexpectedly
  - `xfer_common()` starts a transfer to `0x30` with no live slave pointer

### Next step

- reboot into `Pb92b`
- rerun:
  - `uname -a`
  - `./scripts/bringup_i2c_slave_testunit.sh start 1 0x30`
  - `i2ctransfer -f -y 1 r1@0x30`
  - `dmesg | tail -n 80`

## 2026-05-04 (validated `Pb92b` slave-lifetime diagnostics and narrowed the failure mode)

The seven-patch `Pb92b` runtime result changed the picture again. The failing same-adapter self-read no longer looks like a missing slave registration at transfer start.

### Findings

- running kernel after reboot:
  - `Linux beaglebadge 6.12.57-vendor-edge-k3 #9 SMP PREEMPT Thu Apr 30 08:25:00 UTC 2026 aarch64 GNU/Linux`
- J6 / J7 adapters are present after reboot:
  - `i2c-1`
  - `i2c-3`
- `slave-testunit` remains present and bound on J6 / `i2c-1`:
  - `present: 1-1030`
  - `bound: yes`
  - `slave-testunit`
- a clean same-adapter self-read still fails with:
  - `Error: Sending messages failed: Connection timed out`
- the current diagnostic `dmesg` output for that failing read is now:
  - `omap_i2c 20010000.i2c: slave irq stat=0x06 con=0x8000 bufstat=0x8001 read=0 threshold=1`
  - `omap_i2c 20010000.i2c: slave irq stat=0x02 con=0x8000 bufstat=0x8001 read=0 threshold=1`
- decoded against the OMAP status bits, those IRQs are:
  - `0x06` = `ARDY | NACK`
  - `0x02` = `NACK`
- notably absent during the failing read:
  - `AAS`
  - `RRDY`
  - `XRDY`
  - `XUDF`
  - `master-xfer addr=0x30 no registered slave`
- an additional same-adapter self-write probe:
  - `i2ctransfer -f -y 1 w1@0x30 0x00`
  - also times out
  - produces no slave IRQ log lines

### Meaning

- the current `Pb92b` kernel is no longer failing at the old "no registered slave" point
- same-adapter self-read is now reaching `omap_i2c_slave_irq()`
- but it is only reaching the `ARDY/NACK` tail of a transaction, not an address-match or data-ready event
- same-adapter self-write is even earlier-failing: timeout with no slave IRQ activity at all
- the next useful discriminator is own-address and master/slave register state around listen mode and transaction start, not more slave lifetime logging alone

### Next step

- stage a follow-up diagnostic patch that logs:
  - `OA`
  - `SA`
  - `CON`
  - `STAT`
  - `IE`
- capture those values for both:
  - `i2ctransfer -f -y 1 r1@0x30`
  - `i2ctransfer -f -y 1 w1@0x30 0x00`

## 2026-05-04 (staged own-address register diagnostics)

The next AM62L follow-up is now staged in both the TI kernel tree and the Armbian patch series.

### Changes

- committed in `components/ti-linux-kernel`:
  - `79146a484` `Trace OMAP own-address registers`
- mirrored into the Armbian `k3-6.12` patch stack as:
  - `components/armbian-build/patch/kernel/archive/k3-6.12/0008-Trace-OMAP-own-address-registers.patch`
- committed in `components/armbian-build`:
  - `0dcb10d36` `Carry OMAP own-address diagnostics patch`
- updated the x86 build wrapper to expect the eight-patch series

### Purpose

- include `OA`, `SA`, and `IE` in the existing slave diagnostic log lines
- emit an `xfer-msg` trace when the adapter starts a same-adapter master transfer to its own registered slave address

### Next step

- rebuild the BeagleBadge `vendor-edge` kernel package set with the eight-patch `k3-6.12` series
- copy the returned artifacts back
- install and boot that diagnostic kernel
- reproduce both:
  - `i2ctransfer -f -y 1 r1@0x30`
  - `i2ctransfer -f -y 1 w1@0x30 0x00`

## 2026-05-04 (added install-and-reboot wrapper and boot workspace service)

To reduce repetitive manual steps around kernel iteration, the repo now includes:

- `scripts/install_latest_kernel_and_reboot.sh`
- `systemd/badgesnake-boot-session.service`
- `scripts/install_badgesnake_boot_session_service.sh`
- `scripts/uninstall_badgesnake_boot_session_service.sh`
- `scripts/start_badgesnake_boot_session.sh`

### Purpose

- one command to install the newest returned kernel artifacts and reboot only on success
- a reversible boot-time workspace helper so post-reboot recovery is predictable

### Verification and recovery

- the install wrapper delegates to the existing kernel reinstall helper and only calls `systemctl reboot` after that helper exits successfully
- the boot-session service is repo-managed and reversible with:
  - `./scripts/uninstall_badgesnake_boot_session_service.sh`
- the boot-session helper writes:
  - `artifacts/boot-status/latest.txt`
- if `tmux` is available, it also creates:
  - session `badgesnake`

### Next step

- enable the boot-session service on the live BeagleBadge
- use `./scripts/install_latest_kernel_and_reboot.sh` for the next returned kernel install

## 2026-05-04 (enabled boot workspace service on the live board)

The repo-managed reboot helpers are now installed and verified on this BeagleBadge.

### Findings

- `./scripts/install_badgesnake_boot_session_service.sh` installed and enabled:
  - `badgesnake-boot-session.service`
- live service state:
  - `Loaded: loaded (/etc/systemd/system/badgesnake-boot-session.service; enabled; preset: enabled)`
  - `Active: active (exited)`
- `./scripts/start_badgesnake_boot_session.sh` created:
  - `tmux` session `badgesnake`
- boot status file written and verified:
  - `artifacts/boot-status/latest.txt`
- `./scripts/install_latest_kernel_and_reboot.sh --help` prints the expected usage and environment controls

### Meaning

- future kernel install iterations can use one command to install and reboot
- after reboot, reconnecting to the board can immediately recover the prepared workspace via:
  - `tmux attach -t badgesnake`

### Recovery

- remove the boot-session service with:
  - `./scripts/uninstall_badgesnake_boot_session_service.sh`

## 2026-05-04 (updated boot workspace service to resume Codex)

The boot-session helper now attempts to resume the latest local Codex session inside the `tmux` workspace instead of dropping directly to a shell.

### Changes

- updated:
  - `scripts/start_badgesnake_boot_session.sh`
- the `tmux` session now prefers:
  - `codex resume --last --dangerously-bypass-approvals-and-sandbox --no-alt-screen -C /root/interactor/badge-demo`
- if that exits, it falls back to `bash -l`

### Meaning

- the post-reboot `tmux` session can resume the latest Codex session automatically
- the literal `--yolo` spelling is not used here because the installed `codex` build exposes `resume --last` and the current approval/sandbox flags instead

## 2026-05-04 (verified rebooted kernel and fixed tmux attach target)

Two live follow-ups are now confirmed on this board:

### Findings

- the install-and-reboot wrapper did boot the newly installed kernel:
  - `Linux beaglebadge 6.12.57-vendor-edge-k3 #10 SMP PREEMPT Thu Apr 30 08:25:00 UTC 2026 aarch64 GNU/Linux`
- the boot-session helper did start Codex, but the tmux session name ended up as `0` instead of `badgesnake`
- the live pane showed Codex running inside tmux and resuming the latest session

### Fix

- updated `scripts/start_badgesnake_boot_session.sh` to force stable names after session creation:
  - rename session to `badgesnake`
  - rename window `0` to `workspace`
  - disable automatic window rename on that target
- applied the same rename fix to the live session immediately

### Current live state

- `tmux has-session -t badgesnake` succeeds
- `tmux display-message -p -t badgesnake:0 '#{session_name} #{window_name} #{pane_current_command}'` reports:
  - `badgesnake workspace codex`

## 2026-05-04 (validated `Pdc8d` own-address diagnostics)

The eight-patch own-address diagnostic kernel is now running and produced the first useful split between same-adapter read and write behavior.

### Findings

- the board booted the newly installed kernel and still reports:
  - `Linux beaglebadge 6.12.57-vendor-edge-k3 #10 SMP PREEMPT Thu Apr 30 08:25:00 UTC 2026 aarch64 GNU/Linux`
- the stable boot tmux target now works as intended:
  - `tmux ls` shows `badgesnake`
  - `tmux attach -t badgesnake` succeeds
- after re-instantiating `slave-testunit`:
  - `present: 1-1030`
  - `bound: yes`
  - `slave-testunit`

Same-adapter self-read:

- command:
  - `i2ctransfer -f -y 1 r1@0x30`
- result:
  - `Error: Sending messages failed: Connection timed out`
- diagnostics:
  - `slave xfer-msg stat=0x00 con=0x8400 ie=0x61f oa=0x30 sa=0x30 bufstat=0x8000 read=0 threshold=1`
  - `slave irq stat=0x06 con=0x8000 ie=0x61f oa=0x30 sa=0x30 bufstat=0x8001 read=0 threshold=1`
  - `slave irq stat=0x02 con=0x8000 ie=0x61f oa=0x30 sa=0x30 bufstat=0x8001 read=0 threshold=1`

Same-adapter self-write:

- command:
  - `i2ctransfer -f -y 1 w1@0x30 0x00`
- result:
  - `Error: Sending messages failed: Connection timed out`
- diagnostics:
  - `slave xfer-msg stat=0x00 con=0x8400 ie=0x61f oa=0x30 sa=0x30 bufstat=0x8000 read=0 threshold=1`
  - `slave irq stat=0x16 con=0x8200 ie=0x61f oa=0x30 sa=0x30 bufstat=0x8001 read=0 threshold=1`
  - `slave irq stat=0x06 con=0x8200 ie=0x61f oa=0x30 sa=0x30 bufstat=0x8000 read=1 threshold=1`
  - `slave irq stat=0x02 con=0x8200 ie=0x61f oa=0x30 sa=0x30 bufstat=0x8000 read=0 threshold=1`

Decoded status bits:

- `0x16` = `XRDY | ARDY | NACK`
- `0x06` = `ARDY | NACK`
- `0x02` = `NACK`

### Meaning

- the own-address state is sane during both transactions:
  - `OA=0x30`
  - `SA=0x30`
  - `IE=0x61f`
- same-adapter self-write does reach address-match / transmit-ready territory:
  - `XRDY` is present
  - `con=0x8200` indicates slave transmit path activity
- same-adapter self-read still does not show `AAS`, `XRDY`, or `RRDY`
- so the next problem is more specific than "slave mode is not working":
  - write-side address-match exists but does not complete cleanly
  - read-side path never reaches target-transmit startup

## 2026-05-04 (staged slave-listen control cleanup follow-up)

The next kernel change is now staged around `I2C_CON` cleanup when returning from master mode to slave listen mode.

### Changes

- committed in `components/ti-linux-kernel`:
  - `5e394583e` `Clear stale master state before slave listen`
- mirrored into the Armbian `k3-6.12` patch stack as:
  - `components/armbian-build/patch/kernel/archive/k3-6.12/0009-Clear-stale-master-state-before-slave-listen.patch`
- committed in `components/armbian-build`:
  - `391059c2e` `Carry slave-listen master-state cleanup patch`
- updated the x86 build wrapper to expect the nine-patch series

### Hypothesis

- same-adapter self-transfers are leaking the previous master transfer's `TRX` state into the slave listen phase
- that matches the current split:
  - self-write reaches `XRDY`
  - self-read does not
- clearing master-only `I2C_CON` bits before slave listen should remove that stale state

### Next step

- rebuild the BeagleBadge `vendor-edge` kernel package set with the nine-patch `k3-6.12` series
- copy the returned artifacts back
- install and boot that follow-up kernel
- reproduce:
  - `i2ctransfer -f -y 1 r1@0x30`
  - `i2ctransfer -f -y 1 w1@0x30 0x00`

## 2026-05-05 (validated `P5123` slave-listen cleanup follow-up)

The nine-patch cleanup kernel booted, but it did not change the same-adapter read/write behavior.

### Findings

- running kernel after reboot:
  - `Linux beaglebadge 6.12.57-vendor-edge-k3 #11 SMP PREEMPT Thu Apr 30 08:25:00 UTC 2026 aarch64 GNU/Linux`
- boot-session behavior across reboot remains acceptable:
  - `tmux ls` shows `badgesnake`
  - `tmux attach -t badgesnake` succeeds
- `slave-testunit` still binds on J6 / `i2c-1`

Same-adapter self-read on `P5123`:

- command:
  - `i2ctransfer -f -y 1 r1@0x30`
- result:
  - `Error: Sending messages failed: Connection timed out`
- diagnostics:
  - `slave xfer-msg stat=0x00 con=0x8400 ie=0x61f oa=0x30 sa=0x30 bufstat=0x8000 read=0 threshold=1`
  - `slave irq stat=0x06 con=0x8000 ie=0x61f oa=0x30 sa=0x30 bufstat=0x8001 read=0 threshold=1`
  - `slave irq stat=0x02 con=0x8000 ie=0x61f oa=0x30 sa=0x30 bufstat=0x8001 read=0 threshold=1`

Same-adapter self-write on `P5123`:

- command:
  - `i2ctransfer -f -y 1 w1@0x30 0x00`
- result:
  - `Error: Sending messages failed: Connection timed out`
- diagnostics:
  - `slave xfer-msg stat=0x00 con=0x8400 ie=0x61f oa=0x30 sa=0x30 bufstat=0x8000 read=0 threshold=1`
  - `slave irq stat=0x16 con=0x8200 ie=0x61f oa=0x30 sa=0x30 bufstat=0x8001 read=0 threshold=1`
  - `slave irq stat=0x06 con=0x8200 ie=0x61f oa=0x30 sa=0x30 bufstat=0x8000 read=1 threshold=1`
  - `slave irq stat=0x02 con=0x8200 ie=0x61f oa=0x30 sa=0x30 bufstat=0x8000 read=0 threshold=1`

### Meaning

- clearing master-only `I2C_CON` bits on the return-to-listen path did not change the same-adapter behavior
- the current same-adapter divergence remains:
  - self-write reaches `XRDY`
  - self-read does not
- that strongly suggests the next useful validation is no longer another same-adapter kernel tweak
- the next useful validation is a true second controller or external initiator talking to J6 / `i2c-1`

### Next step

- obtain a real second-controller test path for J6 target mode
- use that path to validate whether AM62L slave mode works correctly when the initiator is not the same adapter

## 2026-05-05 (validated J6 target mode with J7 as a true second controller)

The QWIIC loopback between J7 (`/dev/i2c-3`) and J6 (`/dev/i2c-1`) proves the AM62L slave path is working with a true second controller.

### Findings

- user-provided hardware setup:
  - Zepto temporarily removed from J6
  - short installed between J6 and J7
- `slave-testunit` remained instantiated and bound on J6 / `i2c-1` at `0x30`

Write from J7 to J6:

- command:
  - `i2ctransfer -f -y 3 w1@0x30 0x00`
- result:
  - `EXIT=0`
- J6 slave diagnostics:
  - `slave irq stat=0x200 con=0x8000 ie=0x61f oa=0x30 sa=0x3ff bufstat=0x8000 read=0 threshold=1`
  - `slave irq stat=0x08 con=0x8000 ie=0x61f oa=0x30 sa=0x3ff bufstat=0x8100 read=0 threshold=1`
  - `slave irq stat=0x04 con=0x8000 ie=0x61f oa=0x30 sa=0x3ff bufstat=0x8000 read=0 threshold=1`

Read from J7 to J6:

- command:
  - `i2ctransfer -f -y 3 r1@0x30`
- result:
  - `0x00`
  - `EXIT=0`
- J6 slave diagnostics:
  - `slave irq stat=0x610 con=0x8000 ie=0x61f oa=0x30 sa=0x3ff bufstat=0x8000 read=0 threshold=1`
  - `slave irq stat=0x400 con=0x8000 ie=0x61f oa=0x30 sa=0x3ff bufstat=0x803f read=1 threshold=1`
  - `slave irq stat=0x06 con=0x8000 ie=0x61f oa=0x30 sa=0x3ff bufstat=0x8000 read=1 threshold=1`
  - `slave irq stat=0x02 con=0x8000 ie=0x61f oa=0x30 sa=0x3ff bufstat=0x8000 read=0 threshold=1`

Decoded status bits:

- `0x200` = `AAS`
- `0x08` = `RRDY`
- `0x04` = `ARDY`
- `0x610` = `AAS | XUDF | XRDY`
- `0x400` = `XUDF`
- `0x06` = `ARDY | NACK`
- `0x02` = `NACK`

### Meaning

- AM62L target mode on J6 works when another controller drives the bus
- the major blocker is no longer generic slave support in `i2c-omap`
- the remaining problem is specifically the forced same-adapter self-test path
- that means the next software layer can move forward using a true second initiator model
  - J7 as a host-side initiator for validation
  - or an external controller / Zepto in the final topology

## 2026-05-05 (added repeatable J7-to-J6 validation harness)

The second-controller proof is now captured as a repo-owned validation script instead of a manual terminal recipe.

### Changes

- added:
  - `scripts/validate_j7_to_j6_target_mode.sh`
  - `docs/J7ToJ6TargetModeValidation.md`

### Verification

- `bash -n scripts/validate_j7_to_j6_target_mode.sh`
- live run on this board:
  - `./scripts/validate_j7_to_j6_target_mode.sh`

Observed result from the live script run:

- write succeeds from J7 to J6
- read succeeds from J7 to J6 and returns `0x00`
- representative slave-side diagnostics:
  - write:
    - `slave irq stat=0x208 ...`
    - `slave irq stat=0x04 ...`
  - read:
    - `slave irq stat=0x606 ...`
    - `slave irq stat=0x06 ...`
    - `slave irq stat=0x02 ...`

### Meaning

- the working second-controller path is now easy to re-run after kernel or driver changes
- future transport work can use this script as the current baseline validation path instead of the same-adapter self-test

## 2026-05-05 (staged repeated-start follow-up for J7 -> J6 feature tests)

The next kernel change is now aimed at the repeated-start / partial-command path exposed by the richer `slave-testunit` feature checks.

### Findings behind this change

- `./scripts/validate_j7_to_j6_testunit_features.sh` currently shows:
  - combined write+read transactions complete
  - but both the version query and block-proc-call style read return only zero bytes
- the write-side portion of those transactions clearly reaches J6 and raises slave IRQs
- the most plausible failure is that the slave backend state is being reset before the read phase starts

### Changes

- committed in `components/ti-linux-kernel`:
  - `049fcadf0` `Defer slave STOP until bus is idle`
- mirrored into the Armbian `k3-6.12` patch stack as:
  - `components/armbian-build/patch/kernel/archive/k3-6.12/0010-Defer-slave-STOP-until-bus-is-idle.patch`
- committed in `components/armbian-build`:
  - `2707d0e1c` `Carry deferred slave STOP patch`
- updated the x86 build wrapper to expect the ten-patch series

### Hypothesis

- the OMAP slave path currently maps `ARDY` directly to `I2C_SLAVE_STOP`
- on repeated-start style write->read traffic, that likely resets the
  `slave-testunit` state machine too early
- if `STOP` is deferred until `STAT_BB` clears, partial-command state should
  survive into the read phase

### Next step

- rebuild the BeagleBadge `vendor-edge` kernel package set with the ten-patch `k3-6.12` series
- copy the returned artifacts back
- install and boot that follow-up kernel
- rerun:
  - `./scripts/validate_j7_to_j6_testunit_features.sh`

## 2026-04-27 (module-only iteration boundary)

The reason for using a full rebuild versus a local module build is now explicit.

### Findings

- `i2c-slave-testunit` and `i2c-slave-eeprom` are normal loadable modules under `drivers/i2c/`
- `i2c-omap` is built into the K3 kernel image with `CONFIG_I2C_OMAP=y`
- therefore:
  - module-only builds are a valid faster path for stock slave backend validation
  - module-only builds are not enough for the eventual AM62L slave-mode work in `i2c-omap.c`

### Current recommendation

- finish one corrective rebuild so the returned packages actually contain the slave backend modules
- after that, prefer local module-only iteration until development moves into `i2c-omap.c`
