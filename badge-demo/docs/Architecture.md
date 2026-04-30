# Architecture

This document provides a quick reference around the technical architecture of BadgeSnake.

## Transport

I2C is still the physical transport between BeagleBadge QWIIC connectors and up
to 2 BeagleConnect Zepto boards, but the architectural direction has changed.

The key constraint is that both the development host and the target can
generate asynchronous messages. Because AM62L and MSPM0L both support
multi-controller I2C, the preferred direction is now:

* simplex transactions on the wire
* either side may arbitrate as controller when it has data to send
* one I2C transaction maps to one framed transport message

So the longer-term plan is a kernel-visible serial-to-I2C endpoint bridge rather
than a strict host-controller / target-responder model.

Battlesnake HTTP semantics can still be preserved at the host-runtime layer, but
they should ride on top of endpoint-oriented framed transport instead of directly
mapping `GET`/`PUT` to fixed I2C read/write choreography.

## BeagleBadge

There are 2 QWIIC connectors on BeagleBadge and each is connected to a different I2C bus on the AM62L32 processor:

* I2C1: badge connector J6, SDA on AM62L pin A6, SCL on AM62L pin D7
* I2C2: badge connector J7, SDA on AM62L pin D8, SCL on AM62L pin B8
* OSPI device: ISSI IS25WX256-JHLE, 256Mbit, powered from PMIC B2 1.8V rail

There is also active work to document the Grove connector as a general-purpose host interface:

* capture Linux GPIO line names for Grove-accessible signals
* capture current pad muxing for those pins
* add a controlled way to switch Grove pins between UART, I2C, and GPIO modes

That work is broader than Zepto control and should be treated as generic board capability, not a Zepto-only special case.

Local documents:

* `docs/AM62LTechnicalReferenceManual.md`: AM62L Technical Reference Manual notes
* `docs/AM62LDataSheet.md`: AM62L Datasheet notes

Submodules:

* ../components/beaglebadge: board design repository

## BeagleConnect Zepto

BeagleConnect Zepto exposes BSL-oriented signals on the JST/QWIIC path in the hardware design:

* `BSL_I2C_SDA`
* `BSL_I2C_SCL`
* `BSL_UART_TX`
* `BSL_UART_RX`
* `BSL_INVOKE`

This is enough to justify keeping I2C as the first transport choice for both gameplay and firmware loading experiments, but the exact host-side flashing flow still needs to be documented.

## Firmware Direction

* Student-facing programming target: MicroBlocks
* Example BadgeSnake player firmware should live in this repository
* Preferred host-side flashing tool: `bb-imager-rs`
* Longer term preference: Linux kernel support for flashing where practical

## Host Runtime

* Implementation language: Go
* Planned binary location: `cmd/badgesnake/`
* Planned shared packages: `internal/protocol/` and additional `internal/` packages as needed
* Current board state: `go1.24.4` is installed locally, but first-run builds on-device are slow enough that they should be treated as background jobs

Submodules:

* ../components/beagleconnect-zepto: board design repository
* ../components/bb-imager-rs: flashing utility source repository
* ../components/ti-linux-kernel: TI kernel source starting point for AM62L I2C driver work
