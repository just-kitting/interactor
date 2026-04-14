# Missing Inputs

Use this file only for immediate questions that need answers before the next meaningful implementation step.

Reply by writing plain text below each block-quoted question.

If the bottom of this file is a block quote, there are still unanswered questions.
If the bottom of this file is not a block quote, all current questions have been answered.

> What is the BeagleBadge recovery procedure after a bad OSPI flash, including how to reach serial console, how to force recovery/ROM boot, and what successful recovery output should look like?

An agent won't be able to reach the serial console. It is provided over a USB-to-serial connection by external hosts, but agent sessions are run locally and cannot resume unless there is a successful boot that invokes the session to continue.

In the case of complete disaster, I can recreate the image based on commits that have been pushed before, which is why performing a commit at every task is so critical. I will press the Select button at boot to force booting from the microSD card, rather than OSPI.

See `badge-demo/components/beaglebadge/docs/FAQ.md` for details about the bootmode selection.

Please make sure the bootloader in OSPI sets up some kind of watchdog where if the agent does not mark the boot as good, the previous known-good kernel, device tree and agent-execution-image are executed. You probably need to make an initrd image or some such that is known suitable to restore the agent session for recovery.

> What is the source of truth for BeagleBadge U-Boot artifacts: repository, branch, and preferred commit or release family?

The Armbian build repository was used to produce the U-Boot artifacts living in this image. Ideally, there would be a CI workflow that captured the artifacts and they could be labeled as known-good against a well-known build environment. I suggest we enable the appropriate workflows and fetch those assets to avoid long build times on BeagleBadge itself. The `2025.12-beaglebadge` is a reasonable target on https://github.com/embedconsult/armbian-build for this work.

At some point, I want to move image creation over to https://github.com/embedconsult/bootstrap-qcow2, which is far more minimal than Armbian, but the build process currently rebuilds LLVM from source and that takes a long time.

If you want to make a fork of the various components required to build the bootloader directly, I won't object. You'd have to extract the information out of the Armbian project.

For the images on-board, the current HEAD of the 2025.12-beaglebadge branch was used to create the bootloader assets and all of the base Armbian image being run here now. I couldn't tell you exactly where Armbian puts the configuration, but here's what it would not let me commit to the userpatches subdirectory:

```console
ubuntu@ip-172-31-1-184:~/workspace/armbian-build$ cat userpatches/config-default.conf
PREFER_DOCKER=no
SHARE_LOG=yes                                                   
BOARD=beaglebadge                                                    
BRANCH=vendor-edge                                                    
BUILD_DESKTOP=no                                                    
BUILD_MINIMAL=no                                                    
KERNEL_CONFIGURE=no                                                    
RELEASE=trixie
ubuntu@ip-172-31-1-184:~/workspace/armbian-build$ fgrep -A 13 trixie userpatches/customize-image.sh 
                trixie)
                        apt-get update
                        apt-get install -y gpg
                        curl -fsSL https://packagecloud.io/84codes/crystal/gpgkey | gpg --dearmor > /etc/apt/trusted.gpg.d/84codes_crystal.gpg
                        apt-key adv --keyserver hkp://keyserver.ubuntu.com:80 --recv-keys 7AA630A1EDEE7D73
                        . /etc/os-release
                        cat > /etc/apt/sources.list.d/packagecloud.list <<- EOF
                        deb https://packagecloud.io/84codes/crystal/$ID $VERSION_CODENAME main
                        EOF
                        apt-get update
                        apt-get install -y crystal clang lld libsdl2-dev rustup
                        rustup install stable
                        ;;
        esac
```

> Are the current `/boot/tiboot3.bin`, `/boot/tispl.bin`, and `/boot/u-boot.img` files on this image the intended production boot artifacts, or only temporary/intermediate ones?

Intermediate. They may or may not be suitable for installing into the OSPI.

> What are the exact `fw_env.config` values we should use for `ospi.env` and `ospi.env.backup`, including device path, offset, environment size, erase block size, and sector count if relevant?

I do not know, nor do I know what information is required to figure this out.

> What is the intended BadgeSnake gameplay contract: player count, board size, turn timing, elimination/failure behavior, and whether the host should preserve Battlesnake HTTP semantics internally or define a new I2C-native protocol?

Just use the Battlesnake `rules` engine defaults and existing provided options.

We want minimal changes to `rules`, so the existing HTTP semantics should be maintained with a simple HTTP<->I2C semantic translation. If paths are long in the HTTP semantics, the translation should shorten them to minimize the byte transfer. An HTTP GET should be an I2C write of a path, followed by an I2C read. An HTTP put should be an I2C write of a path, followed by an I2C write of data. Headers should be eliminated where possible.

> Which repository should hold the actual Zepto player firmware or SDK/examples for BadgeSnake, and what student-facing language/runtime should be the default starting point?

For examples, let's keep those in this repository.

Our target is MicroBlocks for the student programming language.

Reminders to self:
* provide bb-imager-rs repo with flashing instructions for MSPM0L1117
* provide information on Zephyr for MSPM0L1117 and Zepto board

> What flashing path should the BeagleBadge host implement for Zepto boards: I2C BSL, UART BSL, JTAG, or another method?

I2C BSL should be used to flash the Zepto boards.

> What command-line tools, scripts, or upstream utilities should be treated as the preferred flashing toolchain on the BeagleBadge image?

`bb-imager-rs` is the prefered flashing tool to run to program BeagleConnect Zepto and to write microSD cards for BeagleBadge. Linux kernel drivers are even more preferred when available, such as for the BeagleBadge OSPI. Eventually, BeagleConnect Zepto should be programmed using a Linux kernel module.
