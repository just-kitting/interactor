# Missing Inputs

This document tracks the components, documents, and decisions that are still needed to build BadgeSnake with confidence.

## Already Supplied

- AM62L datasheet: `docs/am62l.pdf`
- AM62L technical reference manual: `docs/sprujb4a.pdf`
- BadgeSnake architecture note: `docs/Architecture.md`
- Imported submodules:
  - `components/armbian-build`
  - `components/battlesnake-rules`
  - `components/beaglebadge`
  - `components/beagleconnect-zepto`
- BeagleBadge board FAQ in submodule:
  - QWIIC connector mapping
  - OSPI device identification
- Zepto hardware files showing BSL/I2C/UART-related nets

## Still Needed: High Priority

- BeagleBadge recovery guide
  - how to get serial console access on this board
  - how to force ROM/USB/recovery boot modes
  - how to recover from a bad OSPI flash
  - what a successful recovery boot looks like on the console
- U-Boot source-of-truth
  - authoritative upstream repo
  - branch
  - expected commit or release family
  - whether the `/boot` artifacts on this image are the intended production ones
- OSPI environment details
  - exact `fw_env.config` offsets and sizes for `ospi.env` and `ospi.env.backup`
  - expected redundancy strategy for U-Boot environment updates
- BadgeSnake product spec
  - player count
  - turn timing
  - board size
  - elimination/failure behavior
  - whether gameplay remains wire-compatible with Battlesnake HTTP semantics or becomes a local I2C-native variant

## Still Needed: Hardware And Electrical

- BeagleBadge serial-console wiring details
  - connector name
  - voltage level
  - pinout
  - known-good USB adapter wiring
- QWIIC power budget
  - current available per connector
  - whether both attached Zeptos can be powered directly from BeagleBadge during gameplay and flashing
- Display/status hardware notes
  - ePaper update constraints and preferred software stack
  - RGB LED control path
  - 7-segment display control path
  - buzzer control path

## Still Needed: Zepto Software Track

- Zepto firmware/software repo for actual BadgeSnake player code
  - SDK or examples repo if different from the hardware-design submodule
  - preferred language/runtime for student-facing examples
- Host-driven flashing method
  - whether the AM62L host should use I2C BSL, UART BSL, JTAG, or a helper MCU workflow
  - command-line tooling expected on the BeagleBadge image
- Initial I2C protocol spec
  - address assignment
  - registration handshake
  - move request/response encoding
  - timeout and error handling

## Still Needed: Build And Deployment

- Preferred image build flow
  - whether `components/armbian-build` is the canonical path for producing deployable BadgeSnake images
  - how local patches should be carried and tested
- Update/rollback design constraints
  - target storage for rootfs A/B or fallback assets
  - acceptable reboot downtime
  - whether fallback should prefer microSD, eMMC, or alternate OSPI environment entries

## Components To Consider Adding Later

Only add these once the design needs them:

- U-Boot source tree as a submodule if local bootloader changes become necessary
- Linux kernel or DTS source tree as a submodule if BadgeSnake needs board-specific patches
- Zepto firmware SDK/examples repo as a submodule for classroom/player firmware
- BadgeSnake host application repo if it gets split from this integration repo
