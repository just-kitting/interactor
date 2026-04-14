# BadgeSnake TODO

## Repository Setup

- [ ] Create `components/` and establish the first submodules
- [ ] Add `docs/` structure for hardware, protocol, build, and recovery notes
- [ ] Decide initial local code layout (`src/`, `app/`, or similar)
- [ ] Add a `scripts/` directory for setup, flashing, and demo launch tooling

## Product Definition

- [ ] Write a short BadgeSnake product spec
- [ ] Define player count and supported connection methods
- [ ] Define turn timing, board size, and failure behavior
- [ ] Decide whether the game is fully Battlesnake-compatible or a local variant
- [ ] Define what must be visible on the ePaper, RGB LED, buzzer, and 7-segment displays

## Hardware Integration

- [ ] Document BeagleBadge connectors used for Zepto players
- [ ] Choose the Zepto transport layer: I2C, UART, SPI, or mixed
- [ ] Define hot-plug and disconnect behavior
- [ ] Define power budget and connector limits for multiple attached Zeptos
- [ ] Capture display, LED, buzzer, and segment-display interface notes

## Host Runtime

- [ ] Choose implementation language for the BadgeSnake host runtime
- [ ] Build a local simulation mode that runs without attached hardware
- [ ] Implement player registration and health monitoring
- [ ] Implement the game loop and deterministic move resolution
- [ ] Implement ePaper rendering pipeline
- [ ] Implement status indicators and audible feedback
- [ ] Add structured logging for gameplay and device events

## Zepto Firmware Track

- [ ] Select the first Zepto SDK/examples repo to import as a submodule
- [ ] Define the protocol between BeagleBadge and Zepto boards
- [ ] Provide a minimal reference snake implementation
- [ ] Provide a failure-test firmware image that behaves badly on purpose
- [ ] Document classroom-friendly firmware update steps

## Build And Test

- [ ] Add a host-side smoke test for the game engine
- [ ] Add protocol tests with recorded fixtures
- [ ] Add a hardware bring-up checklist for a fresh board
- [ ] Add a demo-start script that validates prerequisites before launch
- [ ] Add a release checklist for live classroom/demo use

## Bootloader And Reliability

- [ ] Record current boot flow and storage map in a dedicated recovery doc
- [ ] Configure `/etc/fw_env.config` for OSPI environment access
- [ ] Verify `fw_printenv` and `fw_setenv` against both env copies
- [ ] Identify the correct bootloader source and artifact provenance
- [ ] Create a safe OSPI flashing procedure from Linux
- [ ] Test booting from OSPI without relying on microSD loader stages
- [ ] Define a good-boot marker and boot-attempt counter
- [ ] Add automatic rollback for failed system updates
- [ ] Add a systemd-based post-update reboot validation flow
- [ ] Test failure cases: bad DTB, bad kernel, bad rootfs, interrupted update
- [ ] Document full recovery from a non-booting OSPI image

## Immediate Next Actions

- [ ] Decide the first set of upstream repos to add under `components/`
- [ ] Add a recovery/boot documentation file with real commands and offsets
- [ ] Determine whether current `/boot/tiboot3.bin`, `tispl.bin`, and `u-boot.img` are the intended production boot artifacts
- [ ] Confirm how BeagleBadge currently reaches serial console during recovery work
- [ ] Pick the host runtime language and the initial simulation target
