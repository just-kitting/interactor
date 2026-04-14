# Release Checklist

Use this checklist before a live BadgeSnake demo or classroom session.

## Repo And Build Inputs

- [ ] All current work is committed
- [ ] Required submodules are initialized at the intended revisions
- [ ] The host runtime build artifacts come from a known-good commit
- [ ] Any Zepto example firmware used in the session is committed

## Board Readiness

- [ ] Board boots reliably from the intended medium
- [ ] Recovery path has been re-verified recently
- [ ] `scripts/check_uboot_artifacts.sh` output has been reviewed
- [ ] Display and button probes show expected devices

## Demo Validation

- [ ] A host-only simulation game has completed successfully
- [ ] At least one hardware-attached test game has completed successfully
- [ ] Failure behavior has been spot-checked with a missing or non-responsive player
- [ ] Logs from the latest validation run were preserved

## Operational Readiness

- [ ] Required cables and Zepto boards are labeled and available
- [ ] microSD recovery media is available
- [ ] The operator knows how to force `SEL` boot-to-microSD
- [ ] The operator knows which commit to return to if rollback is needed
