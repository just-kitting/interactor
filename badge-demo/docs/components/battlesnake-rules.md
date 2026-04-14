# BattleSnake Rules

This submodule contains the source for Battlesnake game rules and logic. It provides a command-line interface for playing the game.

Upstream: https://github.com/BattlesnakeOfficial/rules
Our fork: https://github.com/just-kitting/battlesnake-rules

Modifications to add hardware support for I2C transport should be performed directly to this repository suitable for upstreaming.

Commit locally and the maintainer will provide feedback.

BadgeSnake will use this submodule repo to test game execution from the command-line. Integration with the `arena` web server or a local rendering will be determined later.

## Expected modifications

The first BadgeSnake transport modifications are now present in the local submodule history:

* `sim://` URLs create deterministic in-process players for CLI testing
* `i2c://` URLs preserve the same Battlesnake request lifecycle while standing in for BadgeSnake I2C transport

These are currently CLI-layer changes only. The core rules engine remains close to upstream.

The next transport step is to connect the `i2c://` shape to a real adapter path or an emulated bus path once the kernel-side test environment is available.

Notes:

* A controller over I2C will have a bus and an address.
* The current transport-token draft maps:
  * `GET /` -> `0x01`
  * `POST /start` -> `0x02`
  * `POST /move` -> `0x03`
  * `POST /end` -> `0x04`
* This board currently lacks the `i2c-stub` kernel module, so `sim://` is the primary executable transport path today.
