#!/usr/bin/env bash
#
# previous session: 019e0a2a-bf01-74e2-9f4b-9d456d842498
#
set -euo pipefail

MAX_ITERS="${MAX_ITERS:-10}"
OUT_DIR="${OUT_DIR:-.codex-loop}"
SANDBOX="${SANDBOX:-workspace-write}"
SESSION_ID=""

usage() {
  cat <<'EOF'
Usage:
  codex-until-stuck.sh [options] [task prompt]

Options:
  -s, --session SESSION_ID   Resume this exact Codex session every iteration.
  -n, --max-iters N          Maximum loop iterations. Default: MAX_ITERS env or 10.
  -o, --out-dir DIR          Output/log directory. Default: OUT_DIR env or .codex-loop.
  --sandbox MODE             Codex sandbox mode. Default: SANDBOX env or workspace-write.
                             Common values: read-only, workspace-write, danger-full-access.
  -h, --help                 Show this help.

Examples:
  codex-until-stuck.sh \
    --session 019e0a2a-bf01-74e2-9f4b-9d456d842498

  codex-until-stuck.sh \
    --session 019e0a2a-bf01-74e2-9f4b-9d456d842498 \
    "Continue the existing task. Keep fixing issues until done or genuinely stuck."

  MAX_ITERS=25 codex-until-stuck.sh \
    -s 019e0a2a-bf01-74e2-9f4b-9d456d842498
EOF
}

while [[ $# -gt 0 ]]; do
  case "$1" in
    -s|--session)
      if [[ $# -lt 2 ]]; then
        echo "Missing value for $1" >&2
        exit 64
      fi
      SESSION_ID="$2"
      shift 2
      ;;
    -n|--max-iters)
      if [[ $# -lt 2 ]]; then
        echo "Missing value for $1" >&2
        exit 64
      fi
      MAX_ITERS="$2"
      shift 2
      ;;
    -o|--out-dir)
      if [[ $# -lt 2 ]]; then
        echo "Missing value for $1" >&2
        exit 64
      fi
      OUT_DIR="$2"
      shift 2
      ;;
    --sandbox)
      if [[ $# -lt 2 ]]; then
        echo "Missing value for $1" >&2
        exit 64
      fi
      SANDBOX="$2"
      shift 2
      ;;
    -h|--help)
      usage
      exit 0
      ;;
    --)
      shift
      break
      ;;
    -*)
      echo "Unknown option: $1" >&2
      usage >&2
      exit 64
      ;;
    *)
      break
      ;;
  esac
done

if ! [[ "$MAX_ITERS" =~ ^[0-9]+$ ]] || [[ "$MAX_ITERS" -lt 1 ]]; then
  echo "MAX_ITERS must be a positive integer; got: $MAX_ITERS" >&2
  exit 64
fi

LAST="$OUT_DIR/last-message.md"
LOG="$OUT_DIR/run.log"

TASK_PROMPT="${*:-"Continue the current task until it is complete, verified, or genuinely blocked."}"

mkdir -p "$OUT_DIR"

CONTROL_PROMPT="$(cat <<'EOF'
Work autonomously. Keep making concrete progress.

Rules:
- Inspect the repo and choose the next useful action.
- Make code changes when appropriate.
- Run relevant tests/checks when possible.
- Fix failures you caused.
- Do not stop just because there is more work.
- Only say STUCK when you cannot make further safe progress because of missing credentials, missing product decisions, unavailable external services, destructive permissions, or repeated failure after trying reasonable alternatives.
- Do not claim DONE unless the relevant implementation and verification steps are complete.
- End your final message with exactly one of these lines:
  CONTINUE: <next concrete action>
  DONE: <what is complete and how it was verified>
  STUCK: <blocking reason and what human input is needed>
EOF
)"

continue_prompt() {
  cat <<EOF
Continue from the previous turn in this exact session. Take the next concrete action now.

Do not merely summarize unless summarizing is necessary to proceed.
Do not stop because there is more work.
End with exactly one CONTINUE, DONE, or STUCK line.

$CONTROL_PROMPT
EOF
}

initial_prompt() {
  cat <<EOF
$TASK_PROMPT

$CONTROL_PROMPT
EOF
}

run_codex_initial() {
  if [[ -n "$SESSION_ID" ]]; then
    echo "=== iteration 0: resume explicit session $SESSION_ID ===" | tee -a "$LOG"

    codex exec \
      --sandbox "$SANDBOX" \
      --output-last-message "$LAST" \
      resume "$SESSION_ID" \
      "$(initial_prompt)" | tee -a "$LOG"
  else
    echo "=== iteration 0: start new exec session ===" | tee -a "$LOG"

    codex exec \
      --sandbox "$SANDBOX" \
      --output-last-message "$LAST" \
      "$(initial_prompt)" | tee -a "$LOG"
  fi
}

run_codex_continue() {
  local i="$1"

  if [[ -n "$SESSION_ID" ]]; then
    echo "=== iteration $i: resume explicit session $SESSION_ID ===" | tee -a "$LOG"

    codex exec \
      --sandbox "$SANDBOX" \
      --output-last-message "$LAST" \
      resume "$SESSION_ID" \
      "$(continue_prompt)" | tee -a "$LOG"
  else
    echo "=== iteration $i: resume most recent session for this working directory ===" | tee -a "$LOG"

    codex exec \
      --sandbox "$SANDBOX" \
      --output-last-message "$LAST" \
      resume --last \
      "$(continue_prompt)" | tee -a "$LOG"
  fi
}

check_last_message() {
  if [[ ! -f "$LAST" ]]; then
    echo "Codex did not write expected last-message file: $LAST" >&2
    exit 1
  fi

  if grep -Eq '^DONE:' "$LAST"; then
    echo "Codex reported DONE."
    grep -E '^DONE:' "$LAST" || true
    exit 0
  fi

  if grep -Eq '^STUCK:' "$LAST"; then
    echo "Codex reported STUCK:"
    grep -E '^STUCK:' "$LAST" || true
    exit 2
  fi
}

run_codex_initial
check_last_message

for i in $(seq 1 "$MAX_ITERS"); do
  run_codex_continue "$i"
  check_last_message
done

echo "Reached MAX_ITERS=$MAX_ITERS without DONE or STUCK."
echo "Last assistant message is in: $LAST"
exit 3
