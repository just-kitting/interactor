# BattleSnake Rules

This submodule contains the source for Battlesnake game rules and logic. It provides a command-line interface for playing the game.

Upstream: https://github.com/BattlesnakeOfficial/rules
Our fork: https://github.com/just-kitting/battlesnake-rules

Modifications to add hardware support for I2C transport should be performed directly to this repository suitable for upstreaming.

Commit locally and the maintainer will provide feedback.

BadgeSnake will use this submodule repo to test game execution from the command-line. Integration with the `arena` web server or a local rendering will be determined later.

## Expected modifications

Currently, `rules` expects URLs where a Battlesnake controller will respond. We need to replace this with I2C requests where a Battlesnake controller will respond.

Notes:

* A controller over I2C will have a bus and an address.
* We'll need to define a simple translation of URL path (the part after the domain name) such that an I2C-based controller can satisfy the API.
