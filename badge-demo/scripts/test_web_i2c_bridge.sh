#!/bin/sh
set -eu

ROOT_DIR=$(CDPATH= cd -- "$(dirname -- "$0")/.." && pwd)
PORT="${PORT:-18443}"
HOST="${HOST:-127.0.0.1}"
BASE_URL="https://${HOST}:${PORT}"

cleanup() {
  if [ -n "${SERVER_PID:-}" ]; then
    kill "$SERVER_PID" >/dev/null 2>&1 || true
    wait "$SERVER_PID" >/dev/null 2>&1 || true
  fi
  if [ -n "${TX_PID:-}" ]; then
    wait "$TX_PID" >/dev/null 2>&1 || true
  fi
}
trap cleanup EXIT INT TERM

python3 "$ROOT_DIR/scripts/serve_microblocks_web.py" \
  --host "$HOST" \
  --public-host "$HOST" \
  --port "$PORT" >/dev/null 2>&1 &
SERVER_PID=$!

sleep 1

python3 "$ROOT_DIR/scripts/web_i2c_transaction.py" \
  --url "$BASE_URL" \
  --address 0x42 \
  --write 0x10 0x20 \
  --timeout-ms 5000 \
  --insecure >"$ROOT_DIR/.cache/web_i2c_bridge_result.json" &
TX_PID=$!

sleep 1

REQUEST_JSON=$(curl -sk "${BASE_URL}/api/i2c/next?timeout_ms=100")
printf '%s' "$REQUEST_JSON" | python3 -c '
import json, sys
payload = json.load(sys.stdin)
request = payload.get("request")
assert request is not None, "expected queued request"
assert request["address"] == 0x42, request
assert request["write"] == [0x10, 0x20], request
print(request["id"])
' >"$ROOT_DIR/.cache/web_i2c_bridge_id.txt"

REQUEST_ID=$(cat "$ROOT_DIR/.cache/web_i2c_bridge_id.txt")

curl -sk \
  -H 'Content-Type: application/json' \
  -d "{\"id\":${REQUEST_ID},\"response\":[170,85]}" \
  "${BASE_URL}/api/i2c/respond" >/dev/null

wait "$TX_PID"
unset TX_PID

python3 - <<'PY' "$ROOT_DIR/.cache/web_i2c_bridge_result.json"
import json, sys
with open(sys.argv[1], "r", encoding="utf-8") as handle:
    payload = json.load(handle)
assert payload["ok"] is True, payload
assert payload["response"] == [170, 85], payload
print("web i2c bridge smoke ok")
PY
