# Multi-Instance Workflow

This project now has more than one Codex-capable environment:

- BeagleBadge
- `bq2`

The best way to connect them is **git-first**, not through a fragile live
Codex-to-Codex session link.

## Recommendation

Use the repository as the canonical coordination surface:

- commits
- tracked docs
- scripts
- branch state

Treat any live shell/session connection as secondary and optional.

## Why Git-First Is Better Here

- The board is memory-constrained, so large live coordination layers add risk.
- The work already depends on exact, replayable kernel/build/install state.
- The project agreement is that important decisions belong in git anyway.
- A git-based handoff survives reboots, session exits, and machine changes.

## Practical Split

Use `bq2` for:

- large codebase inspection
- kernel/source edits
- Armbian patch updates
- build-log analysis
- x86-host build orchestration

Current limitation:

- `bq2` does not have Docker access, so it should not be treated as the current
  Armbian kernel build host
- for now, actual Armbian kernel package builds still run on the separate x86
  host, while `bq2` prepares source/patch state and reviews logs

Use BeagleBadge for:

- runtime validation
- install/reboot loops
- J6/J7 bus tests
- Zepto flashing/probing
- verifying board state after boot

## Handoff Rules

Before handing work across instances:

1. Commit current work.
2. Update `PROGRESS.md` if the board state, diagnosis, or active next step changed.
3. Update `TODO.md` if the task state changed.
4. If the handoff is subtle, add a short committed note in the relevant doc.

## What Not To Do

Avoid relying on:

- an uncommitted shared shell transcript
- a live Codex session bridge as the only source of truth
- concurrent edits to the same file from both instances

Those can still be used temporarily, but only as convenience layers above the
repo, not instead of it.

## Current Recommendation

For the current AM62L `i2c-omap` work:

- do source edits and Armbian patch prep on `bq2`
- run the actual Armbian kernel build on the x86 host
- copy returned build artifacts into this repo
- run `install_latest_kernel_and_reboot.sh` and all J7 -> J6 validation on the
  BeagleBadge instance
