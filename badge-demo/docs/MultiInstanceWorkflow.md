# Multi-Instance Workflow

This project now has more than one Codex-capable environment:

- BeagleBadge
- `bq2`
- local Ollama analysis helper

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

Current bootstrap path outside this repo:

- `/root/interactor/scripts/setup-bq2.sh`
- `/root/interactor/scripts/setup-repo.sh`
- `/root/interactor/scripts/sysroot-build-plan.json`

Those belong to the enclosing `interactor` checkout, not `badge-demo` itself.
They are useful for bringing up the `bq2` environment, but they do not change
the git-first coordination model for this repo.

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

## How To Task `bq2`

`bq2` appears to work better with narrow tasks than with open-ended “continue”
requests on the full repo.

Recommended task shape:

- give one concrete goal
- name the exact files it may edit
- include the current relevant commit(s)
- include the exact validation command to run or update
- require a commit at the end

Good examples:

- “Update `components/ti-linux-kernel/drivers/i2c/busses/i2c-omap.c` to avoid
  the `P641a` recv-len regression, then mirror it into
  `components/armbian-build/patch/kernel/archive/k3-6.12/0016-...`, and commit
  both repos.”
- “Read `log-kernel-...` and summarize only the new build failure in
  `PROGRESS.md`.”

Avoid:

- “continue”
- “work on whatever is next”
- tasks that span unrelated files or subsystems without a concrete boundary

Recommended first `bq2` task for the current state:

- inspect `components/ti-linux-kernel/drivers/i2c/busses/i2c-omap.c`
- explain why the `P641a` recv-len follow-up can fall through into
  `omap_i2c_transmit_data()` during a receive transaction
- update only:
  - `components/ti-linux-kernel/drivers/i2c/busses/i2c-omap.c`
  - `components/armbian-build/patch/kernel/archive/k3-6.12/0016-Pause-OMAP-recv-len-after-count-byte.patch`
- commit both repos

## What Not To Do

Avoid relying on:

- an uncommitted shared shell transcript
- a live Codex session bridge as the only source of truth
- concurrent edits to the same file from both instances

Those can still be used temporarily, but only as convenience layers above the
repo, not instead of it.

## How To Use The Ollama Helper

The Ollama helper should be treated as a read-only analysis engine, not as a
driver of workflow.

Use it for:

- patch review
- log interpretation
- Linux I2C / OMAP behavior analysis
- narrowing hypotheses before source edits

Do not use it for:

- making coordination decisions
- deciding what to edit next across the whole repo
- editing files directly

Recommended task shape:

- provide one narrow question
- include the exact file excerpt or `dmesg` excerpt
- ask for concise findings with file/line references
- ask it to rank 1-3 likely causes or review one specific patch

Best return path:

- paste the Ollama findings back into chat as plain text
- if the findings are especially important, also save them in a tracked note or
  commit them into the repo through the main Codex instance

Current best Ollama task:

- review the `P641a` recv-len regression
- focus on why a receive transaction can later reach
  `omap_i2c_transmit_data()` and crash in the IRQ thread
- inspect:
  - `components/ti-linux-kernel/drivers/i2c/busses/i2c-omap.c`
  - the `P641a` dmesg excerpt with `Too much work in one IRQ`
  - the oops showing `pc : omap_i2c_transmit_data.isra.0`

Suggested prompt:

```text
Review this AM62L/OMAP I2C regression in read-only mode.

Context:
- true SMBus block-proc-call on J7 -> J6 still fails
- direct I2C_RDWR + I2C_M_RECV_LEN now fails with Protocol error
- dmesg shows many "omap_i2c 20010000.i2c: Too much work in one IRQ"
- then an oops with:
  - pc : omap_i2c_transmit_data.isra.0+0x70/0x1ac
  - lr : omap_i2c_xfer_data+0x264/0x338
  - x0 : 0000000000000000

Please analyze only:
- components/ti-linux-kernel/drivers/i2c/busses/i2c-omap.c

Question:
- how can the recv-len receive path fall through into transmit handling with invalid state?
- identify the most likely control-flow bug(s)
- cite exact function names and the relevant branches
- keep the answer concise
```

## Current Recommendation

For the current AM62L `i2c-omap` work:

- do source edits and Armbian patch prep on `bq2`
- run the actual Armbian kernel build on the x86 host
- copy returned build artifacts into this repo
- run `install_latest_kernel_and_reboot.sh` and all J7 -> J6 validation on the
  BeagleBadge instance
