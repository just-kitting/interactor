# Progress Log

## 2026-04-14

Initial planning pass completed.

### Repository state

- `badge-demo` currently contains only a project `README.md`
- The repo is appropriate as an integration root rather than a code-heavy project today
- No submodules exist yet

### Board state confirmed on this system

- Hostname: `beaglebadge`
- OS: Armbian-unofficial `26.02.0-trunk trixie`
- Kernel: `6.12.57-vendor-edge-k3`
- Running on BeagleBadge / AM62L3
- Root and `/boot` are on microSD (`mmcblk1p2` and `mmcblk1p1`)
- eMMC exists as `mmcblk0`
- OSPI flash is visible via MTD partitions from Linux
- `/boot` contains `tiboot3.bin`, `tispl.bin`, `u-boot.img`, `Image`, and `uInitrd`
- `fw_printenv` currently returns nothing because environment access is not configured yet

### Boot-related observations

- `dmesg` shows Linux sees OSPI partitions named for `tiboot3`, `tispl`, `u-boot`, primary/backup env, and an OSPI rootfs area
- `/usr/lib/u-boot/platform_install.sh` on this image only copies bootloader artifacts into `/boot`; it does not provide a complete OSPI flashing workflow by itself
- The current repository should therefore treat OSPI boot enablement as a planned deliverable, not as already solved

### Documentation added this session

- `AGENTS.md` for strategy and execution rules
- `TODO.md` for staged delivery tracking
- `PROGRESS.md` for verified findings and session notes

### Highest-value information still needed

- Exact upstream repo list to use as submodules
- BeagleBadge recovery instructions for bad OSPI flashes
- Serial-console wiring/access details
- Preferred source of truth for U-Boot and kernel artifacts
- BadgeSnake gameplay and protocol requirements
