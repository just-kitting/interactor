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

## Host Runtime

- [X] Choose implementation language for the BadgeSnake host runtime
- [~] Build a local simulation mode that runs without attached hardware
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
- [ ] Provide a failure-test firmware image that behaves badly on purpose
- [ ] Document classroom-friendly firmware update steps

## Build And Test

- [~] Add a host-side smoke test for the game engine
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
- [ ] Run and capture one successful `components/battlesnake-rules` simulated CLI game on-device
- [ ] Decide whether the current Armbian kernel packaging should grow `i2c-stub` support or whether BadgeSnake testing should stay on `sim://` until real Zepto devices are attached
