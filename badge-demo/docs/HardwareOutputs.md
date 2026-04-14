# Hardware Outputs

This note captures what the current BeagleBadge Linux image exposes for display and status hardware.

## Confirmed From The Live System

- ePaper framebuffer is available at `/dev/fb0`
- DRM is present via `/dev/dri`
- ALSA sound cards are not currently exposed
- SPI devices currently enumerate as:
  - `spi0.0`: OSPI NOR flash
  - `spi1.0`: `gdey042t81` ePaper panel driver path
  - `spi2.0`: `mcp23s18` GPIO expander
  - `spi3.0`: `dh2228fv` on the `lora` DT node
- I2C buses visible to Linux:
  - `/dev/i2c-0`
  - `/dev/i2c-2`
- Buttons are exposed through the `gpio-keys` DT node:
  - `UP`
  - `DOWN`
  - `LEFT`
  - `RIGHT`
  - `SELECT`
  - `BACK`
- Kernel input handlers currently show:
  - `gpio-keys` on `event0`
  - PMIC power button on `event1`
  - `pwm-beeper` on `event2`

## Device-Tree Clues

The device tree contains named symbols for:

- `rgb_led_default_pins`
- `pwm_beeper_default_pins`
- `mcp23s18`
- `gpio_keys`

This indicates the board description knows about RGB LED wiring, a PWM-driven beeper path, and an SPI GPIO expander, even though generic sysfs classes for LEDs or PWM are not currently surfaced in this image.

## What This Means For BadgeSnake

- ePaper rendering can be approached through framebuffer/DRM first
- Button input is already discoverable through normal Linux input paths
- RGB LED and buzzer support likely require either:
  - enabling additional kernel drivers or DT nodes
  - direct control through GPIO/PWM once the responsible devices are identified cleanly
- The beeper already appears as a kernel `pwm-beeper` input/sound device, but not as an ALSA sound card
- 7-segment display control is not yet identified from the current live-system probe and remains an open implementation task

## Follow-Up Work

- identify the Linux input event nodes corresponding to the front-panel buttons
- determine whether the `mcp23s18` expander owns LEDs, 7-segment lines, or other status outputs
- determine whether PWM beeper support is disabled in the current DT or only missing a userspace control path
- choose whether the ePaper path should be framebuffer-based or DRM-native for the first pass
