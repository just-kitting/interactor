# Missing Inputs

Use this file only for immediate questions that need answers before the next meaningful implementation step.

Reply by writing plain text below each block-quoted question.

If the bottom of this file is a block quote, there are still unanswered questions.
If the bottom of this file is not a block quote, all current questions have been answered.

> What is the BeagleBadge recovery procedure after a bad OSPI flash, including how to reach serial console, how to force recovery/ROM boot, and what successful recovery output should look like?

> What is the source of truth for BeagleBadge U-Boot artifacts: repository, branch, and preferred commit or release family?

> Are the current `/boot/tiboot3.bin`, `/boot/tispl.bin`, and `/boot/u-boot.img` files on this image the intended production boot artifacts, or only temporary/intermediate ones?

> What are the exact `fw_env.config` values we should use for `ospi.env` and `ospi.env.backup`, including device path, offset, environment size, erase block size, and sector count if relevant?

> What is the intended BadgeSnake gameplay contract: player count, board size, turn timing, elimination/failure behavior, and whether the host should preserve Battlesnake HTTP semantics internally or define a new I2C-native protocol?

> Which repository should hold the actual Zepto player firmware or SDK/examples for BadgeSnake, and what student-facing language/runtime should be the default starting point?

> What flashing path should the BeagleBadge host implement for Zepto boards: I2C BSL, UART BSL, JTAG, or another method?

> What command-line tools, scripts, or upstream utilities should be treated as the preferred flashing toolchain on the BeagleBadge image?
