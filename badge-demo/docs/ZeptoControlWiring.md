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

Those live pad values match the expected `GPIO0_27` / `GPIO0_28` control candidates, but they also encode an important limitation:

- `MUXMODE = 7` (GPIO)
- `TX_DIS = 1` (pad output driver disabled)

That means the pads are readable as GPIO inputs but are not directly drivable as outputs until a pinctrl state reprograms the pad configuration.

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

Additional live debugging on 2026-06-18 found that:

- `gpioset` can claim `gpiochip1` lines `27/28`
- the external `RST` signal still does not move low
- the likely reason is that `PADCONFIG42/43` are left at `0x08254007`, which keeps `TX_DIS=1`
- `/sys/kernel/debug/pinctrl/4084000.pinctrl-pinctrl-single/gpio-ranges` is also empty, so the pinctrl driver is not handling GPIO range transitions for these lines

This is now treated as a software configuration issue rather than a wiring-only issue.

## Next Step

Use the confirmed Linux lines `gpiochip1` / `27` and `28` for future BSL automation attempts.

Remaining work:

- install a DT overlay that programs PADCONFIG42/43 into a non-`TX_DIS` GPIO-capable state
- reboot and re-test whether `RST` can be driven low from Linux
- once `RST` is proven, tune the host-side reset / boot timing and polarity
- confirm whether the Zepto expects `BSL` sampled only during a narrower reset window
- once the correct sequence is known, add a host-side toggle script and integrate it into the flashing wrappers

## Overlay Candidate

This repo now includes a first overlay candidate intended to clear the boot-time `TX_DIS` state on these pads:

- `overlays/beaglebadge-zepto-control.dtso`
- `scripts/validate_zepto_control_overlay.sh`
- `scripts/install_zepto_control_overlay.sh`

## Zepto Qwiic Validation

Additional bare-metal validation on 2026-06-19 confirms that the Zepto firmware can actively take ownership of the Qwiic bus pins from the MSPM0 side:

- diagnostic builds:
  - `firmware/zepto/examples/baremetal/qwiic_diag/`
  - `scripts/build_zepto_qwiic_diag.sh`
  - `scripts/flash_zepto_qwiic_diag.sh`
- `zepto_qwiic_scl_low` was flashed successfully through the known-good raw BSL script:
  - `scripts/flash_zepto_bsl_direct_i2c.py`
- after booting that image, BeagleBadge host probes on `/dev/i2c-1` no longer behaved like a normal idle empty bus:
  - `i2cdetect -y -r 1` hit the timeout wrapper
  - `i2ctransfer -f -y 1 ...` failed with `Device or resource busy`

That is strong evidence that Zepto `PA1` really is the Qwiic `SCL` line and that bare-metal firmware can drive it low from the attached Zepto.

The corresponding `zepto_qwiic_sda_low` image did not produce an equally clear host-side symptom, so `PA0` ownership is not yet as strongly host-validated as `PA1`.

This means the remaining blocker for Zepto-hosted I2C target mode is no longer basic Qwiic pin reachability. The remaining issue is target-mode peripheral configuration / protocol behavior after boot.
