# Architecture

This document provides a quick reference around the technical architecture of BadgeSnake.

## Transport

I2C will be used to transport data between BeagleBadge's QWIIC connectors and up to 2 BeagleConnect Zepto's QWIIC connectors for both flashing and game play.

Gameplay should preserve Battlesnake HTTP semantics and translate them onto I2C:

* HTTP `GET`: I2C write of a shortened path token, then I2C read of the response body
* HTTP `PUT`: I2C write of a shortened path token, then I2C write of the request body
* HTTP headers should be removed wherever possible

This keeps `components/battlesnake-rules` changes minimal while allowing Zepto players to behave like Battlesnake controllers.

## BeagleBadge

There are 2 QWIIC connectors on BeagleBadge and each is connected to a different I2C bus on the AM62L32 processor:

* I2C1: badge connector J6, SDA on AM62L pin A6, SCL on AM62L pin D7
* I2C2: badge connector J7, SDA on AM62L pin D8, SCL on AM62L pin B8
* OSPI device: ISSI IS25WX256-JHLE, 256Mbit, powered from PMIC B2 1.8V rail

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
