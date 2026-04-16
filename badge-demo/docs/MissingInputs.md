# Missing Inputs

Use this file as the primary method for immediate questions that need answers before the next meaningful implementation step. This provides a logged history of interactions. When communication has been consolidated into other documentation, it should be removed from this file.

User will reply by writing plain text below each block-quoted question.

If the bottom of this file is a block quote, there are still unanswered questions.
If the bottom of this file is not a block quote, all current questions have been answered.

> Which BeagleBadge connector is the Zepto currently attached to: J6 on `/dev/i2c-0` or J7 on `/dev/i2c-2`?

I have it attached to the one closest to the edge, J6.

> When the Zepto is held in MSPM0 BSL mode on this hardware, should it ACK at I2C address `0x48`, or is there a board-specific address or extra reset/enable sequence I should use before probing?

I don't think it was reconfigured in any way, but it is possible. There isn't a defined board-specific address. It is possible that the BSL may have timed-out or otherwise failed/hung. The only verification I did to know the BSL is running was to note that the firmware I programmed in previously did not execute (it was an example of fading an RGB LED).

I added an MSPM0BSLUsersGuide.md to try to answer BSL questions.

> Please re-enter the Zepto on J6 into MSPM0 BSL mode, then leave it connected and let me know when I should immediately run `scripts/probe_zepto_bsl_active.sh 0`.

The active probe was already run on J6 with the exact `bb-imager-rs` MSPM0 connection packet and did not get an ACK at `0x48`.

The exact physical sequence in use is:

I hold BOOT, press RST, then release BOOT fairly quickly. The firmware stops execution. USB power is not related (Zepto is powered from BeagleBadge over the QWIIC connector, not any separate USB power).

The user also pointed out the likely root cause correctly:

`/dev/i2c-0` is main I2C0, `/dev/i2c-2` is likely WKUP_I2C0, and the QWIIC buses were not enabled in the device tree.

All current blocking questions have been answered or reduced to local build and flashing work.

> What additional guidance do you have for the agent?

I have guidance on using PlatformIO as a build tool for Zepto firmware based on its usage in MicroBlocks.

These are not open questions unless they become false later.

- The current Zepto firmware track is **native MSPM0L1117 firmware**, not Arduino-core on Zephyr.
- The current flashing path is **MSPM0 BSL via `bb-imager-rs` / `bb-imager-cli`**.
- **OpenOCD/SWD is out of scope** for the current critical path.
- **UART is out of scope** for bring-up and early development.
- The first milestones are **`blink`** and **`I2C target echo`**.
- A Zepto is currently attached to **J6**, the connector closest to the edge.
- The physical BSL entry sequence currently in use is: **hold BOOT, press RST, then release BOOT fairly quickly**.
- The Zepto is powered from the **BeagleBadge over the QWIIC connector**, not from separate USB power.

> What is the preferred repo/path/branch for the native Zepto firmware work?
>
> Example answer format:
> `components/beagleconnect-zepto @ <commit-or-branch>`
> `components/microblocks-smallvm @ <commit-or-branch>`
> `components/bb-imager-rs @ <commit-or-branch>`

`examples/baremetal/blink` and `examples/baremetal/i2c-target-echo` on the main branch in the primary repository should be use.

> Which file(s) in `beagleconnect-zepto` should be treated as the best source of truth for Zepto hardware pin information?
>
> Paste exact relative paths if known.

I just created `components/beagleconnect-zepto/FAQ.md` sto hold the best source of truth for Zepto hardware pin information. 

> What guidance do you have for Zepto bring-up to unblock a correct `blink` target and the first PlatformIO board definition?

See questions and answers below.

> What is the exact user-visible LED pin and polarity on Zepto?
>
> If there is an enable GPIO, transistor gate, or RGB LED detail that matters for the first `blink`, describe it here.

Read the above mentioned FAQ.md. `blink` should use `TIMG0_C1`/`PA12` with active-high polarity.

> Is there a preferred board name / board ID string to use consistently across PlatformIO, docs, and future MicroBlocks board metadata?
>
> Example: `beagleconnect_zepto`, `beagleconnect-zepto`, `MSPM0L1117-Zepto`

Use `zeptomspm0l1117` as the board name.

> Are there any known flash or RAM regions that early firmware should avoid?
>
> Examples: bootloader reservation, config page, BSL-related reservation, calibration page, or “safe to use full app region for now”.

I'm absolutely certain that there are, but you'll need to read `MSPM0BSLUsersGuide.md`, `MSPM0L111xDataSheet.md` and `MSPM0LTechnicalReferenceManual.md` to find out.

> Which MSPM0 I2C peripheral and pins should the first Zepto application firmware use for the QWIIC connector?

I placed an `FAQ.md` into the `components/beagleconnect-zepto` repository. The I2C used on the QWIIC connectors is on PA0 (SDA) and PA1 (SCL).

> What 7-bit I2C address should the first non-BSL Zepto application firmware use for the target echo?

Agent should decide.

> What is the preferred maximum transaction size for the first echo milestone?
>
> Suggested answer format: `8`, `16`, `32`, or `64` bytes.
>
> Note: this is only for the first echo milestone, not the eventual MicroBlocks logical message size.

Let's see if we can't transfer at least 64 bytes at first. We'll want to be able to do a 1024 byte transfer for MicroBlocks.

> What should count as “echo milestone complete” from the host side?

A script writes `01 02 03 04 ...` and reads back the same bytes ten times in a row without reset.

> What guidance do you have on PlatformIO?

These unblock the first custom MSPM0 PlatformIO platform scaffold.

> Should the early PlatformIO support vendor a minimal MSPM0 SDK subset into the repo, fetch it during setup, or assume a separately installed SDK path?

The MSPM0 SDK can be read as a guideline, but it is bloated and confusing as a source base and should not be included. We want a very minimal implementation to start.

> What host environments matter for early PlatformIO bring-up?

BeagleBadge only for now

> Is there an existing Zepto or MSPM0 example in any of your repos that the agent should copy from first?

No

> MicroBlocks-Over-I2C Planning Inputs
>
> These are not required for `blink`, but they help the agent avoid painting itself into a corner during the echo design.
>
> What I2C address should be reserved for future MicroBlocks VM transport messages?

The agent should choose and it should be easy to change.

> Do you want the first echo transport format to already resemble a future MicroBlocks framing layer, or should it stay deliberately tiny and throwaway?

The `blink` and `i2c-target-echo` examples should stay as minimal as possible without any concept of reuse outside of understanding.

> Are there any application-side I2C addresses you already want to reserve separately from the future VM transport address?

https://github.com/adafruit/I2C_Addresses is an excellent list of "used" I2C addresses. MicroBlocks support several of the devices listed there. I think 0x48 as the bootloader address is the only one to seriously avoid.
