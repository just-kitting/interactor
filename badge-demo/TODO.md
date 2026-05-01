# BadgeSnake TODO

## Repository Setup

- [X] Create `components/` and establish the first submodules
- [X] Add `docs/` structure for hardware, protocol, build, and recovery notes
- [X] Decide initial local code layout (`cmd/badgesnake`, `internal/`, `examples/microblocks`)
- [X] Add a `scripts/` directory for setup, flashing, and demo launch tooling

## Product Definition

- [X] Write a short BadgeSnake product spec based on Battlesnake `rules` defaults
- [X] Define player count and supported connection methods
- [X] Define turn timing, board size, and failure behavior
- [X] Decide whether the game is fully Battlesnake-compatible or a local variant
- [X] Define what must be visible on the ePaper, RGB LED, buzzer, and 7-segment displays

## Hardware Integration

- [X] Document BeagleBadge connectors used for Zepto players
- [X] Choose the Zepto transport layer: I2C, UART, SPI, or mixed
- [X] Define hot-plug and disconnect behavior
- [ ] Define power budget and connector limits for multiple attached Zeptos
- [ ] Capture display, LED, buzzer, and segment-display interface notes with control paths
- [X] Enable BeagleBadge J6/J7 I2C controllers in the device tree and verify they appear as Linux adapters
- [ ] Document Grove connector Linux line names and add a pinmux control path for switching Grove signals between UART, I2C, and GPIO roles

## Host Runtime

- [X] Choose implementation language for the BadgeSnake host runtime
- [~] Build a local simulation mode that runs without attached hardware
- [~] Make the MicroBlocks IDE usable on BeagleBadge through a web-hosted path
- [X] Complete one successful end-to-end MicroBlocks web build and launch smoke test on BeagleBadge
- [ ] Add a kernel-visible I2C simulation path so Linux controller tools can issue real bus transactions against the simulator
- [ ] Choose and implement the kernel-visible simulation mechanism; `i2c-stub` is not sufficient for dynamic Boardie target behavior
- [ ] Add AM62L `i2c-omap` slave-mode support and validate it with `slave-testunit`
- [ ] Design and implement a serial-to-I2C endpoint bridge kernel driver for BadgeSnake
- [ ] Define the userspace endpoint ABI exposed by the bridge driver
- [X] Stage Armbian kernel config changes for `I2C_SLAVE_TESTUNIT` and `I2C_SLAVE_EEPROM`
- [X] Add a concrete `slave-testunit` bring-up guide and helper script
- [X] Add an exact x86 Docker host wrapper for the BeagleBadge `vendor-edge` kernel build
- [X] Add initial host-runtime code skeleton for protocol and in-memory transport
- [ ] Implement player registration and health monitoring
- [ ] Implement the game loop and deterministic move resolution
- [ ] Implement ePaper rendering pipeline
- [ ] Implement status indicators and audible feedback
- [ ] Add structured logging for gameplay and device events

## Zepto Firmware Track

- [X] Select the first Zepto SDK/examples repo to import as a submodule
- [X] Define the protocol between BeagleBadge and Zepto boards
- [ ] Provide a minimal reference snake implementation
- [~] Add a MicroBlocks simulation path for I2C target handling
- [ ] Provide a failure-test firmware image that behaves badly on purpose
- [ ] Document classroom-friendly firmware update steps
- [X] Build and verify `bb-imager-cli` with `zepto_i2c` support on-device
- [X] Confirm which BeagleBadge I2C bus exposes a Zepto in MSPM0 BSL mode
- [X] Add a first native Zepto PlatformIO project and build `blink`
- [ ] Flash the first PlatformIO `blink` image successfully through MSPM0 I2C BSL
- [ ] Diagnose the current `bb-imager-cli` MSPM0 flash failure after BSL handshake succeeds
- [ ] Identify the Linux GPIO mapping and current pinmux state for Grove pin 1 (`BSL`) and pin 2 (`RST`) so Zepto bootloader entry can be automated

## Build And Test

- [X] Add a host-side smoke test for the game engine
- [X] Add protocol tests with recorded fixtures
- [X] Add a hardware bring-up checklist for a fresh board
- [X] Add a demo-start script that validates prerequisites before launch
- [X] Add a release checklist for live classroom/demo use
- [X] Add a fixture validator for protocol inputs

## Bootloader And Reliability

- [X] Record current boot flow and storage map in a dedicated recovery doc
- [ ] Configure `/etc/fw_env.config` for OSPI environment access using the derived candidate layout
- [ ] Verify `fw_printenv` and `fw_setenv` against both env copies after environment format is validated
- [X] Identify the correct bootloader source and artifact provenance
- [ ] Create a safe OSPI flashing procedure from Linux
- [ ] Test booting from OSPI without relying on microSD loader stages
- [ ] Define a good-boot marker and boot-attempt counter
- [ ] Add automatic rollback for failed system updates
- [ ] Add a systemd-based post-update reboot validation flow
- [ ] Test failure cases: bad DTB, bad kernel, bad rootfs, interrupted update
- [ ] Document full recovery from a non-booting OSPI image

## Immediate Next Actions

- [X] Decide the first set of upstream repos to add under `components/`
- [X] Add a recovery/boot documentation file with real commands, offsets, and serial-console expectations
- [X] Determine whether current `/boot/tiboot3.bin`, `tispl.bin`, and `u-boot.img` are the intended production boot artifacts
- [X] Confirm how BeagleBadge currently reaches serial console during recovery work
- [X] Pick the host runtime language and the initial simulation target
- [X] Run and capture one successful `components/battlesnake-rules` simulated CLI game on-device
- [ ] Decide whether the current Armbian kernel packaging should grow `i2c-stub` support or whether BadgeSnake testing should stay on `sim://` until real Zepto devices are attached
- [X] Configure a default Rust toolchain for local `bb-imager-rs` builds
- [X] Reboot with the QWIIC I2C overlay enabled and verify new Linux adapters appear for J6/J7
- [X] Verify the Zepto in BSL mode ACKs at `0x48` on the correct QWIIC bus
- [~] Build `bb-imager-cli` with `zepto_i2c` and enumerate/flash the live Zepto over `/dev/i2c-1`
- [X] Review `docs/ZeptoPlatformIO.md`
- [X] Add a board-state preservation capture process before any fresh microSD image work
- [X] Run `scripts/capture_beaglebadge_state.sh` on the live board and copy the resulting archive off-board
- [X] Run `scripts/build_beaglebadge_vendor_edge_kernel_x86_docker.sh` on an x86 Docker host
- [X] Reinstall the returned `vendor-edge-k3` kernel artifacts on BeagleBadge and reboot into the rebuilt kernel
- [ ] Rebuild the BeagleBadge `vendor-edge-k3` kernel artifacts after restoring the missing K3 I2C slave config lines
- [~] Reinstall the corrected `vendor-edge-k3` kernel artifacts on BeagleBadge and reboot again
- [X] Reinstall the corrected `vendor-edge-k3` kernel artifacts on BeagleBadge
- [X] Restore the local QWIIC overlay after the DTB package refresh
- [X] Reboot into the corrected kernel+overlay state
- [X] Rebuild and install a kernel whose host build log shows the `archive/k3-6.12` AM62L slave patch series applied
- [X] Reboot into the newly installed `P5507` AM62L slave-test kernel and rerun `slave-testunit` binding checks
- [X] Verify `i2c-slave-testunit` is present after reboot
- [X] Prove that `new_device` can instantiate a slave-testunit node on J6 / `i2c-1`
- [X] Add enough AM62L `i2c-omap` slave registration support for `slave-testunit` to bind on J6 / `i2c-1`
- [ ] Diagnose why forced same-adapter transactions to bound `slave-testunit` on J6 currently time out
- [X] Rebuild and install the AM62L `i2c-omap` slave TX underflow/FIFO follow-up kernel
- [X] Reboot into the newly installed `P024c` AM62L slave-test kernel and rerun forced-read validation
- [X] Rebuild and install the AM62L 1-byte slave FIFO threshold follow-up kernel
- [ ] Reboot into the pinned `P9d8b` 4-patch kernel
- [ ] Re-test `i2ctransfer -f -y 1 r1@0x30` after the 1-byte slave FIFO threshold follow-up kernel boots
- [X] Mirror the staged AM62L `i2c-omap` slave-support change into the Armbian `archive/k3-6.12` kernel patchset
