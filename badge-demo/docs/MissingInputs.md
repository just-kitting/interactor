# Missing Inputs

Use this file only for immediate questions that need answers before the next meaningful implementation step.

Reply by writing plain text below each block-quoted question.

If the bottom of this file is a block quote, there are still unanswered questions.
If the bottom of this file is not a block quote, all current questions have been answered.

> Which BeagleBadge connector is the Zepto currently attached to: J6 on `/dev/i2c-0` or J7 on `/dev/i2c-2`?

I have it attached to the one closest to the edge, J6.

> When the Zepto is held in MSPM0 BSL mode on this hardware, should it ACK at I2C address `0x48`, or is there a board-specific address or extra reset/enable sequence I should use before probing?

I don't think it was reconfigured in any way, but it is possible. There isn't a defined board-specific address. It is possible that the BSL may have timed-out or otherwise failed/hung. The only verification I did to know the BSL is running was to note that the firmware I programmed in previously did not execute (it was an example of fading an RGB LED).

I added an MSPM0BSLUsersGuide.md to try to answer BSL questions.

> Please re-enter the Zepto on J6 into MSPM0 BSL mode, then leave it connected and let me know when I should immediately run `scripts/probe_zepto_bsl_active.sh 0`.
