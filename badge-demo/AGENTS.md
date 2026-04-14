# BadgeSnake Agent Strategy

## Mission

Build a reproducible BeagleBadge demonstration where multiple BeagleConnect Zepto boards participate in a Battlesnake-style game controlled by student firmware and rendered locally on the badge.

This repository is the integration point. Prefer keeping upstream technology components in git submodules under this directory and keeping BadgeSnake-specific glue, build scripts, docs, and test assets in this repo.

## Current Baseline

Observed on 2026-04-14 from this board:

- Board: BeagleBadge (`Machine model: Texas Instruments AM62L3 Badge`)
- OS: Armbian-unofficial `26.02.0-trunk` on Debian `13 (trixie)`
- Kernel: `6.12.57-vendor-edge-k3`
- Root filesystem: microSD (`mmcblk1p2`)
- Boot partition: microSD (`mmcblk1p1`, mounted at `/boot`)
- eMMC present: `mmcblk0` (`3.6G`)
- OSPI MTD partitions visible in Linux:
  - `ospi.tiboot3`
  - `ospi.tispl`
  - `ospi.u-boot`
  - `ospi.env`
  - `ospi.env.backup`
  - `ospi.rootfs`
  - `ospi.phypattern`
- `/boot` already contains `tiboot3.bin`, `tispl.bin`, and `u-boot.img`
- `fw_printenv` is not yet configured on this image, so boot-environment management is not ready

Treat this as a starting point, not a guarantee. Re-check after any system or bootloader work.

## Working Agreements

- Commit at the end of each meaningful work session.
- Prefer small commits with descriptive messages over large undocumented batches.
- Update `PROGRESS.md` whenever repository direction, board state, or confirmed findings change.
- Put open questions and missing-user-input requests into tracked repo documents rather than relying on chat.
- Use `docs/MissingInputs.md` only for immediate blocking questions.
- Format `docs/MissingInputs.md` as question/answer prompts with block quotes for questions and plain text answers below them.
- Keep the last line of `docs/MissingInputs.md` as a block quote whenever unanswered questions remain.
- Treat git history and committed docs as the primary collaboration channel for decisions and follow-ups.
- Do not consider a wrap-up complete if local changes are still uncommitted.
- Do not rewrite or discard user changes unless explicitly asked.
- When touching system boot or update logic, document the exact verification and recovery steps before marking a task done.

## Repository Shape

Target structure:

- `README.md`: user-facing overview
- `AGENTS.md`: execution strategy for future agents
- `TODO.md`: live delivery checklist
- `PROGRESS.md`: dated notes with confirmed board/repo state
- `docs/`: hardware notes, protocols, rollout procedures, recovery guides
- `scripts/`: host-side setup, flashing, validation, packaging, and demo launch helpers
- `components/`: git submodules for upstream technology pieces
- `src/` or `app/`: BadgeSnake integration code specific to this demo
- `tests/`: host-side simulation, protocol validation, and smoke tests

## Submodule Policy

Use git submodules for upstream components that have their own lifecycle and should be updated independently.

Expected submodule classes:

- BeagleBadge platform software
  - kernel/device-tree sources if local patches become necessary
  - U-Boot / bootloader sources or board support trees
  - Armbian build or packaging logic if image generation becomes part of the workflow
- BadgeSnake application dependencies
  - game engine / rules service
  - display stack
  - hardware abstraction libraries
- Zepto-side technology
  - firmware SDKs
  - protocol examples
  - classroom starter firmware

Rules:

- Place submodules under `components/`.
- Always commit integration patches into submodule repos in a way that does not break upstream.
- Upstream breaking patches should not be made and project-specific integration code should be kept in this repo.
- For each submodule, add a short `docs/components/<name>.md` note describing:
  - why it exists
  - upstream URL
  - local patches or overlays
  - update procedure
  - how it is used by BadgeSnake

## Delivery Streams

Work should advance in these streams in parallel where practical:

1. Demo definition
   - settle game rules, player count, update rate, UX, and success criteria
2. Hardware interfaces
   - Zepto connection topology, discovery, hot-plug behavior, fault handling
3. Badge runtime
   - game loop, rendering, audio/status outputs, local logging
4. Developer workflow
   - submodule layout, build scripts, simulation, packaging
5. Device reliability
   - update process, bootloader in OSPI, reboot safety, rollback

## Reliability Strategy

The board should be able to restart after software changes without depending on a fragile single-path boot setup.

Target end state:

- Primary boot chain stored in OSPI
- Boot environment managed explicitly with both primary and fallback entries
- Rootfs and boot artifacts versioned and deployable in a controlled way
- Failed updates detectable and reversible
- Recovery path documented and testable from a cold boot

Recommended phases:

1. Baseline capture
   - document boot order, current storage usage, environment offsets, and recovery access
2. Environment control
   - configure `fw_printenv` / `fw_setenv` against the OSPI env partitions
   - confirm read/write of both primary and backup environment locations
3. Bootloader deployment
   - install a known-good bootloader into OSPI from the running Linux system
   - verify that the board still boots without depending on microSD-resident loader stages
4. Rollback design
   - define a good-boot marker and boot-attempt counter
   - keep a fallback boot target that remains untouched until the new system proves healthy
5. Update automation
   - create scripts/services that stage updates, reboot, validate health, and roll back on failure
6. Failure drills
   - test bad kernel, bad DTB, bad rootfs, and interrupted deployment scenarios

Never treat OSPI flashing as routine until recovery has been rehearsed and documented.

## Helpful Documentation To Add

The most valuable inputs from the user or upstream repos would be:

- BeagleBadge hardware references:
  - schematics
  - connector pin maps
  - OSPI part details
  - boot mode strap information
  - serial console access details
- Known-good software references:
  - U-Boot source repo / branch / commit
  - kernel source repo / branch / commit
  - any BeagleBadge-specific DTS or overlay repos
- Recovery instructions:
  - ROM/USB boot procedure
  - SD rescue image procedure
  - how to recover from a broken OSPI bootloader
- BadgeSnake product definition:
  - game rules
  - expected player count
  - Zepto protocol choice
  - latency targets
  - display and sound requirements
- Upstream component list:
  - which repos should become submodules first
  - where local changes are expected to be pushed

If only one document gets added next, make it a board recovery guide with exact steps and expected serial-console output.

## Definition Of Done

A work item is not done unless:

- the code or doc change is committed
- `TODO.md` reflects the new state
- `PROGRESS.md` records important findings
- verification is documented
- rollback or recovery impact is stated for any boot/system change
