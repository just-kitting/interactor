#!/bin/sh
set -u

SCRIPT_DIR=$(CDPATH= cd -- "$(dirname -- "$0")" && pwd)
REPO_ROOT=$(CDPATH= cd -- "$SCRIPT_DIR/.." && pwd)
BRINGUP="$REPO_ROOT/scripts/bringup_i2c_slave_testunit.sh"

J6_BUS="${BADGESNAKE_J6_BUS:-1}"
J7_BUS="${BADGESNAKE_J7_BUS:-3}"
J6_ADDR="${BADGESNAKE_J6_ADDR:-0x30}"
J7_ADDR="${BADGESNAKE_J7_ADDR:-0x31}"
OUT_DIR="${1:-/tmp/badgesnake-ti-e2e-i2c-$(date -u +%Y%m%dT%H%M%SZ)}"
LOG="$OUT_DIR/run.log"
DMESG_LOG="$OUT_DIR/dmesg.log"

mkdir -p "$OUT_DIR"
: >"$LOG"

log()
{
	printf '%s\n' "$*" | tee -a "$LOG"
}

run()
{
	log ""
	log "$ $*"
	"$@" >>"$LOG" 2>&1
	rc=$?
	log "exit=$rc"
	return 0
}

capture_dmesg()
{
	dmesg >"$DMESG_LOG" 2>&1 || true
	log ""
	log "Filtered dmesg:"
	grep -E 'omap_i2c|Transmit underflow|Arbitration lost|timed out|timeout|slave irq|isr-master|master-enter' "$DMESG_LOG" >>"$LOG" 2>&1 || true
}

log "BadgeSnake AM62L OMAP I2C dual-listener reproduction"
log "output_dir=$OUT_DIR"
run uname -a
if [ -r /proc/device-tree/model ]; then
	run cat /proc/device-tree/model
fi

for cmd in i2ctransfer dmesg grep tee; do
	command -v "$cmd" >/dev/null 2>&1 || {
		log "missing required command: $cmd"
		exit 1
	}
done

[ -x "$BRINGUP" ] || {
	log "missing helper: $BRINGUP"
	exit 1
}

log ""
log "Reset target backend state"
run "$BRINGUP" stop "$J6_BUS" "$J6_ADDR"
run "$BRINGUP" stop "$J7_BUS" "$J7_ADDR"
run "$BRINGUP" stop "$J7_BUS" "$J6_ADDR"

log ""
log "Start both target backends"
run "$BRINGUP" start "$J6_BUS" "$J6_ADDR"
run "$BRINGUP" start "$J7_BUS" "$J7_ADDR"
run "$BRINGUP" status "$J6_BUS" "$J6_ADDR"
run "$BRINGUP" status "$J7_BUS" "$J7_ADDR"

run dmesg -C

log ""
log "Dual-listener transfers expected to show current failure signature"
run i2ctransfer -f -y "$J7_BUS" "w1@$J6_ADDR" 0x00
run i2ctransfer -f -y "$J7_BUS" "r1@$J6_ADDR"
run i2ctransfer -f -y "$J6_BUS" "w1@$J7_ADDR" 0x00
run i2ctransfer -f -y "$J6_BUS" "r1@$J7_ADDR"

log ""
log "Target status after transfers"
run "$BRINGUP" status "$J6_BUS" "$J6_ADDR"
run "$BRINGUP" status "$J7_BUS" "$J7_ADDR"

capture_dmesg

log ""
log "Wrote:"
log "  $LOG"
log "  $DMESG_LOG"
