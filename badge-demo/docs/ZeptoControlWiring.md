# Zepto Control Wiring

This document records the current temporary hardware control wiring between the BeagleBadge Grove connector and the attached Zepto.

## Current Wiring

User-provided live wiring on 2026-04-30:

- BeagleBadge Grove pin 4
  - wire color: yellow
  - connected to Zepto `BSL`
- BeagleBadge Grove pin 3
  - wire color: white
  - connected to Zepto `RST`

This wiring is intended to support automated Zepto bootloader entry from the BeagleBadge host instead of requiring only manual button timing.

## Current State

The electrical wiring is now known, but the Linux GPIO mapping is not yet captured in the repo.

Current `gpioinfo` output on the BeagleBadge does not expose obvious Grove-specific line names, so additional board or DT mapping work is still needed before a reliable control script can toggle these lines.

## Next Step

Identify the exact Linux GPIO chip/line mapping for:

- Grove pin 3 -> Zepto `RST`
- Grove pin 4 -> Zepto `BSL`

Once those are known, add:

- a host-side toggle script for reset/BSL sequencing
- integration into the Zepto flashing wrappers
