# System State Preservation

Do not replace the current microSD starting image until the current board state has been captured and copied off-board.

This repo now includes a capture script for that purpose:

```sh
./scripts/capture_beaglebadge_state.sh
```

## What It Preserves

The capture includes:

- the current `badge-demo` git state and recent history
- a git bundle for the superproject
- a tarball of the current working tree, including the checked-out submodule contents, excluding `.git`, `.cache`, prior `artifacts/`, and Armbian output directories
- current board boot configuration from `/boot`
- package inventory, block-device layout, mounted filesystems, and kernel command line
- current `/etc/fstab` and `/etc/fw_env.config` if present

The capture is written under:

- `artifacts/state-capture/<timestamp>/`
- `artifacts/state-capture/beaglebadge-state-<timestamp>.tar.gz`

## Required Process Before Replacing The microSD Image

1. Ensure the repo is committed:
   - `git status --short --branch`
   - `git submodule status`
2. Run the capture:
   - `./scripts/capture_beaglebadge_state.sh`
3. Copy the resulting archive off the BeagleBadge to another machine or storage device.
4. Record where the archive was copied in `PROGRESS.md`.
5. Only then discuss or attempt writing a new starting microSD image.

This capture can take several minutes because the working-tree archive includes the checked-out source trees under `components/`.

## Restoring From The Capture

The capture is intended to preserve both the repository state and the board-specific boot/runtime context.

For repo recovery:

1. Clone or recreate the destination working directory.
2. Restore the git history with the saved superproject bundle if needed.
3. Restore any untracked but important local files from `repo-working-tree.tar.gz`.

For board-state review:

1. Review `boot/` for the previous `uEnv.txt`, DTB, and overlay state.
2. Review `system/` for package inventory, block devices, and filesystem layout.
3. Use that record before attempting to recreate prior behavior on a new card.

## Current Project Policy

- Kernel bring-up can proceed with package builds from `components/armbian-build`.
- A fresh microSD image is not yet an approved next step.
- Any future image replacement must start with this capture flow.
