# Zepto Control Wiring

This document records the current temporary hardware control wiring between the BeagleBadge Grove connector and the attached Zepto.

## Current Wiring

User-corrected live wiring on 2026-04-30:

- BeagleBadge Grove pin 1
  - wire color: yellow
  - connected to Zepto `BSL`
  - user-noted package ball: `G23`
- BeagleBadge Grove pin 2
  - wire color: white
  - connected to Zepto `RST`
  - user-noted package ball: `G22`

The earlier pin-3/pin-4 note was backwards and should not be reused.

This wiring is intended to support automated Zepto bootloader entry from the BeagleBadge host instead of requiring only manual button timing.

## SoC And Linux Mapping

Validated from the AM62L datasheet / TRM and the live BeagleBadge pinctrl state:

- Grove pin 1 / Zepto `BSL`
  - package ball: `G23`
  - SoC pad: `GPMC0_AD12`
  - pad config register: `PADCONFIG42`
  - pad config address: `0x040840A8`
  - GPIO function: `GPIO0_27`
  - Linux line: `gpiochip1` line `27`
- Grove pin 2 / Zepto `RST`
  - package ball: `G22`
  - SoC pad: `GPMC0_AD13`
  - pad config register: `PADCONFIG43`
  - pad config address: `0x040840AC`
  - GPIO function: `GPIO0_28`
  - Linux line: `gpiochip1` line `28`

The board DTS does not currently assign badge-specific GPIO line names to these two lines, which is why they appear as unnamed entries in `gpioinfo`.

## Live Pinmux State

Validated on the live board after reboot with the Zepto attached:

- `gpiochip1` line `27` idle read: `0`
- `gpiochip1` line `28` idle read: `1`
- `/sys/kernel/debug/pinctrl/4084000.pinctrl-pinctrl-single/pinmux-pins`
  - `pin 42`: `MUX UNCLAIMED`, `GPIO UNCLAIMED`
  - `pin 43`: `MUX UNCLAIMED`, `GPIO UNCLAIMED`
- `/sys/kernel/debug/pinctrl/4084000.pinctrl-pinctrl-single/pins`
  - `pin 42` / `0x40840a8`: `08254007`
  - `pin 43` / `0x40840ac`: `08254007`

Those live pad values are consistent with GPIO-mode pads and match the expected `GPIO0_27` / `GPIO0_28` control candidates.

## Current State

The electrical wiring, SoC pad mapping, Linux GPIO mapping, and live pinmux state are now known and recorded here.

Automated Zepto BSL entry is still not working yet. A host-side probe was attempted using:

- `gpiochip1` line `27` as `BSL` and line `28` as `RST`
- `gpiochip1` line `28` as `BSL` and line `27` as `RST`
- both candidate `BSL` hold levels: `0` and `1`
- both candidate reset pulse directions: `0 -> 1` and `1 -> 0`

Neither ordering produced an MSPM0 BSL ACK at `0x48` on `/dev/i2c-1`, so one of the following is still true:

- the logical signal assignment is inverted relative to the wire note
- the reset / boot timing needs adjustment
- the Zepto side is not currently entering BSL from this host-driven sequence

## Next Step

Use the confirmed Linux lines `gpiochip1` / `27` and `28` for future BSL automation attempts.

Remaining work:

- tune the host-side reset / boot timing and polarity
- confirm whether the Zepto expects `BSL` sampled only during a narrower reset window
- once the correct sequence is known, add a host-side toggle script and integrate it into the flashing wrappers
