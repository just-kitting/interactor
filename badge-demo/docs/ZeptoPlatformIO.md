# Zepto PlatformIO Guidance

## Purpose

This note is for the Zepto firmware track only.

The goal is to create the smallest maintainable path to build native MSPM0L1117 firmware for BeagleConnect Zepto inside the MicroBlocks work tree, using the existing MSPM0 BSL flashing flow via `bb-imager-rs` / `bb-imager-cli`.

The first milestone is deliberately small:

1. `blink`
2. `I2C target echo`

Do not let future MicroBlocks integration expand this first milestone.

## First-Milestone Constraints

- Native MSPM0L1117 firmware path only.
- No Arduino-core on Zephyr for this bring-up track.
- No UART bring-up path.
- No OpenOCD work.
- No BLE, USB, or IDE GUI work in this first pass.
- Do not block on the later MicroBlocks VM transport.
- Keep memory use visible and conservative from day one.

## Definition Of Done For Milestone 1

Milestone 1 is done when all of the following are true:

- A custom MSPM0 PlatformIO path exists and is checked into the chosen repo/location.
- A BeagleConnect Zepto board manifest and `platformio.ini` environment exist.
- `pio run` builds a native Zepto `blink` firmware image.
- `pio run -t upload` or an equivalent documented command flashes that image through the existing BSL flow.
- The Zepto LED blink is verified on real hardware.
- `pio run` builds a native Zepto `I2C target echo` firmware image.
- A host-side smoke test can perform write/read echo against the physical Zepto on the chosen J6 Linux adapter.
- The implementation notes record code size, RAM usage, and any unresolved blockers.

## Recommended Implementation Order

1. Freeze the missing inputs.
2. Add the custom PlatformIO development platform scaffold.
3. Add the Zepto board manifest and environment.
4. Add minimal startup/linker/clock/GPIO support.
5. Build `blink`.
6. Wire upload to the existing `bb-imager-rs` / `bb-imager-cli` path.
7. Flash and verify `blink`.
8. Add `I2C target echo`.
9. Add a host smoke test for the echo protocol.
10. Record size/memory results.
11. Stop and reassess before any MicroBlocks VM transport work.

## Suggested Repository Layout

Use the smallest layout that keeps the work discoverable. A reasonable first pass looks like this:

```text
platforms/
  mspm0/
    platform.json
    main.py
boards/
  BeagleConnectZepto.json
src/
  zepto/
    board_init.c
    board_init.h
    led.c
    led.h
    i2c_target_echo.c
    i2c_target_echo.h
examples/
  zepto-blink/
  zepto-i2c-echo/
scripts/
  flash_zepto_bsl.sh
  test_zepto_i2c_echo.sh
docs/
  zepto/
    platformio.md
    firmware-bringup.md
```

If the project prefers a different layout, keep the same separation of concerns:

- PlatformIO platform scaffold
- board manifest
- minimal native board support
- flashing wrapper
- host smoke test
- notes/docs

## PlatformIO Strategy

Treat this as two layers:

### 1. Custom MSPM0 development platform

This layer teaches PlatformIO how to build and upload the Zepto firmware.

It should contain:

- package/tool declarations
- toolchain selection
- framework selection for a native bare-metal build
- builder logic
- upload integration

Keep the first version extremely small. The goal is not to create a polished general MSPM0 ecosystem yet. The goal is to support Zepto firmware bring-up with as little friction as possible.

### 2. Zepto board manifest

This layer should describe the Zepto board as a board target.

At minimum, capture:

- board name
- vendor
- MCU family/name
- CPU architecture
- clock frequency
- flash size
- RAM size
- upload protocol as `custom` if needed
- any board-level build flags that are truly fixed

Avoid stuffing board-specific policy into the global platform layer if it only applies to Zepto.

## PlatformIO Environment Guidance

The first environment should be boring and explicit.

Use a dedicated environment for Zepto bring-up, for example:

```ini
[env:zepto_blink]
platform = <local-or-custom-mspm0-platform>
board = BeagleConnectZepto
framework = baremetal
upload_protocol = custom
extra_scripts =
  pre:scripts/pio_zepto_upload.py
build_flags =
  -D ZEPTO_BLINK
```

Add a second environment for echo, for example:

```ini
[env:zepto_i2c_echo]
platform = <local-or-custom-mspm0-platform>
board = BeagleConnectZepto
framework = baremetal
upload_protocol = custom
extra_scripts =
  pre:scripts/pio_zepto_upload.py
build_flags =
  -D ZEPTO_I2C_ECHO
```

The exact `platform` syntax can be adjusted to match however the custom platform is stored locally, but the important point is that the Zepto environments must be isolated and explicit.

## Upload Integration Guidance

Do not make OpenOCD or UART a prerequisite for the first upload path.

Wrap the known-good BSL flashing flow in one stable place and have PlatformIO call that wrapper.

Recommended pattern:

- `scripts/flash_zepto_bsl.sh` contains the authoritative shell wrapper.
- PlatformIO upload logic calls that wrapper.
- The wrapper takes the built image path as an argument.
- The wrapper also takes or derives the final Linux adapter path for J6.

The upload wrapper should:

- fail loudly on missing inputs
- echo the exact command it is about to run
- support a dry-run mode if easy
- avoid hidden environment assumptions

Keep the wrapper simple enough that it can also be run by hand outside PlatformIO.

## Native Firmware Bring-Up Rules

For the first pass, keep the native support layer thin.

Target the minimum pieces needed for `blink` and `I2C target echo`:

- reset/startup
- vector table
- clock init
- GPIO output for LED
- basic delay or timer tick
- I2C target init
- I2C target RX/TX handling

Prefer a thin HAL boundary such as:

```c
void board_init(void);
void board_led_set(int on);
void board_delay_ms(unsigned ms);
void board_i2c_target_init(void);
void board_i2c_target_poll(void);
```

The point is not abstraction for abstraction's sake. The point is to keep future MicroBlocks integration from having to untangle bring-up code from app logic.

## Memory Discipline

Keep memory pressure obvious from the start.

Rules for milestone 1:

- avoid dynamic allocation unless there is a compelling reason
- avoid large debug buffers
- avoid `printf` in steady-state paths if a smaller alternative exists
- keep echo buffers intentionally small at first
- record code size and RAM usage for every bring-up milestone
- keep map files or size summaries with the build outputs or in docs

Do not introduce the larger future MicroBlocks buffer expectations into the first echo implementation. The echo milestone exists to prove the board support and target-mode I2C plumbing, not to solve full VM transport sizing yet.

## I2C Echo Guidance

The first echo protocol should be deliberately simple.

Goals:

- prove target-mode I2C works on the real board
- prove host write/read sequences are reliable
- prove the chosen address is reachable on the J6 path
- prove recovery from repeated test cycles

The first implementation should document:

- chosen 7-bit target address
- chosen transaction model
- maximum supported payload for the first milestone
- timeout behavior
- buffer ownership rules
- what happens if the host reads before a reply is ready

Start conservative. A stable 16- or 32-byte echo is more valuable than a speculative large transfer that is hard to debug.

## Host Smoke Test Guidance

Add one script dedicated to the echo milestone.

The script should:

- take the Linux adapter path as an argument or environment variable
- take the target address as an argument or environment variable
- perform one known write/read transaction
- print pass/fail in a way that is easy to paste into logs
- exit nonzero on failure

A minimal script is enough. The important thing is that it becomes the repeatable proof that the target-mode I2C path is alive.

## What Not To Do Yet

Do not spend time on these in the first pass:

- serial/UART transport support
- OpenOCD setup
- IDE-side I2C connection GUI
- full MicroBlocks message transport over I2C
- large-message segmentation
- file transfer support
- polishing multi-board support

These are follow-on tasks after `blink` and `I2C target echo` are stable.

## Future MicroBlocks Transport Note

Once `I2C target echo` is stable, the next design step should treat the current serial framing as transport-specific rather than universal.

Working direction for the later phase:

- preserve the logical MicroBlocks message semantics
- do not assume serial start/terminator bytes are needed on I2C
- use explicit I2C transaction framing and segmentation outside the logical message body
- keep room for a dedicated VM/control address and possibly future app/message address separation

This note is here only to prevent architectural drift during the bring-up work. Do not implement it in milestone 1 unless explicitly requested.

## Recommended Deliverables From The Agent

For the first pass, ask the agent to produce all of the following:

1. the PlatformIO scaffold
2. the Zepto board manifest
3. the upload wrapper
4. the `blink` firmware
5. the `I2C target echo` firmware
6. the host echo smoke test
7. a short bring-up note describing how to build, flash, and test
8. a short file-by-file change summary
9. a short list of blockers or assumptions that still need user input

## Stop Condition

If the agent completes `blink` and `I2C target echo` and the remaining work would require guessing about addresses, adapter mapping, or upload commands, it should stop and update `MissingInputs.md` instead of improvising.