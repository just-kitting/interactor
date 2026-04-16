# MicroBlocks SmallVM

This submodule contains the source for the MicroBlocks interpreter and development environment that should be updated to support BeagleConnect Zepto and aallow communication with the development environment over I2C, rather than UART.

Upstream: https://bitbucket.org/john_maloney/smallvm
Our fork: https://github.com/just-kitting/microblocks-smallvm

Modifications to add hardware support for I2C transport of development environment communications should be performed directly to this repository suitable for upstreaming.

Commit locally and the maintainer will provide feedback.

BadgeSnake will use this submodule repo to provide support for MicroBlocks graphical programming to the user on BeagleBadge and pushing data to BeagleConnect Zepto for immediate commands and executable opcodes.
