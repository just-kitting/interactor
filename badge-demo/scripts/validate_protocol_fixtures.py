#!/usr/bin/env python3
import json
from pathlib import Path
import sys


ROOT = Path(__file__).resolve().parent.parent
FIXTURES = ROOT / "fixtures" / "protocol"


def load_json(name: str):
    path = FIXTURES / name
    with path.open("r", encoding="utf-8") as handle:
        return json.load(handle)


def require(condition: bool, message: str):
    if not condition:
        raise SystemExit(message)


def validate_tokens():
    payload = load_json("path_tokens.json")
    require(payload.get("version") == 1, "path_tokens.json version must be 1")
    tokens = payload.get("tokens")
    require(isinstance(tokens, list) and tokens, "path_tokens.json tokens must be a non-empty list")

    values = set()
    names = set()
    endpoints = set()
    for item in tokens:
        token = item.get("token")
        method = item.get("method")
        path = item.get("path")
        name = item.get("name")
        require(isinstance(token, int) and 0 < token < 256, f"invalid token value: {token!r}")
        require(method in {"GET", "POST"}, f"invalid method for token {token}: {method!r}")
        require(isinstance(path, str) and path.startswith("/"), f"invalid path for token {token}: {path!r}")
        require(isinstance(name, str) and name, f"invalid name for token {token}: {name!r}")
        require(token not in values, f"duplicate token value: {token}")
        require(name not in names, f"duplicate token name: {name}")
        require((method, path) not in endpoints, f"duplicate endpoint mapping: {method} {path}")
        values.add(token)
        names.add(name)
        endpoints.add((method, path))


def validate_move_response():
    payload = load_json("move-response.json")
    require(payload.get("move") in {"up", "down", "left", "right"}, "move-response.json move must be one of up/down/left/right")
    require(isinstance(payload.get("shout"), str), "move-response.json shout must be a string")


def validate_snake_request(name: str):
    payload = load_json(name)
    require("game" in payload, f"{name} missing game")
    require("board" in payload, f"{name} missing board")
    require("you" in payload, f"{name} missing you")
    require(isinstance(payload.get("turn"), int), f"{name} turn must be an int")

    game = payload["game"]
    board = payload["board"]
    you = payload["you"]

    require(isinstance(game.get("id"), str) and game["id"], f"{name} game.id must be non-empty")
    require(isinstance(game.get("timeout"), int) and game["timeout"] > 0, f"{name} game.timeout must be > 0")
    require(isinstance(board.get("height"), int) and board["height"] > 0, f"{name} board.height must be > 0")
    require(isinstance(board.get("width"), int) and board["width"] > 0, f"{name} board.width must be > 0")
    require(isinstance(board.get("snakes"), list), f"{name} board.snakes must be a list")
    require(isinstance(board.get("food"), list), f"{name} board.food must be a list")
    require(isinstance(board.get("hazards"), list), f"{name} board.hazards must be a list")
    require(isinstance(you.get("id"), str) and you["id"], f"{name} you.id must be non-empty")


def validate_info_response():
    payload = load_json("info-response.json")
    for key in ("apiversion", "author", "color", "head", "tail", "version"):
        require(isinstance(payload.get(key), str) and payload[key], f"info-response.json {key} must be a non-empty string")


def main():
    required = [
        "path_tokens.json",
        "info-response.json",
        "start-request.json",
        "move-request.json",
        "move-response.json",
        "end-request.json",
    ]
    for name in required:
        require((FIXTURES / name).is_file(), f"missing required fixture: {name}")

    validate_tokens()
    validate_info_response()
    validate_snake_request("start-request.json")
    validate_snake_request("move-request.json")
    validate_snake_request("end-request.json")
    validate_move_response()

    print("Protocol fixtures validated successfully.")


if __name__ == "__main__":
    try:
        main()
    except SystemExit as exc:
        if isinstance(exc.code, str):
            print(exc.code, file=sys.stderr)
            raise SystemExit(1)
        raise
