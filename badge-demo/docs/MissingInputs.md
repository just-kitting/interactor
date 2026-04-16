# Missing Inputs

Use this file only for immediate blocking questions that require user input before the next meaningful implementation step.

If the bottom of this file is a block quote, there are still unanswered questions.
If the bottom of this file is not a block quote, all current blocking questions have been answered.

All currently tracked Zepto bring-up questions have been answered and moved into the main project documents:

- `docs/ZeptoPlatformIO.md`
- `docs/ZeptoFlashing.md`
- `docs/components/beagleconnect-zepto.md`
- `components/beagleconnect-zepto/FAQ.md`
- `PROGRESS.md`

> I probed `/dev/i2c-1` immediately after your last `go` and still got no ACK at `0x48`. Please try a tighter experiment: keep the Zepto `BOOT` button held continuously, press and release `RST`, keep holding `BOOT`, and then reply `go-hold` while `BOOT` is still held so I can probe during the asserted BSL invoke condition.
