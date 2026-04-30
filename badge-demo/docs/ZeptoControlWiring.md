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

## Current State

The electrical wiring is now known, but the Linux GPIO mapping is not yet captured in the repo.

Current `gpioinfo` output on the BeagleBadge does not expose obvious Grove-specific line names, so additional board or DT mapping work is still needed before a reliable control script can toggle these lines.

## Next Step

Identify the exact Linux GPIO chip/line mapping and pinmux state for:

- Grove pin 1 -> Zepto `BSL`
- Grove pin 2 -> Zepto `RST`

Once those are known, add:

- a host-side toggle script for reset/BSL sequencing
- integration into the Zepto flashing wrappers
