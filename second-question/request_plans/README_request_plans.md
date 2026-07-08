# Dry-run provider request plans

These JSONL files are examples of the requests the script will prepare before calling the video providers.

- `sora/sora_requests.jsonl`: all scenes using the Sora context-visual profile. Locked local scenes are marked as local renders.
- `omni/omni_requests.jsonl`: all scenes using the Gemini Omni context-visual profile. Locked local scenes are marked as local renders.
- `compare/sora/sora_requests.jsonl`: configured comparison scenes (3, 5) for Sora.
- `compare/omni/omni_requests.jsonl`: configured comparison scenes (3, 5) for Gemini Omni.

Model names are intentionally edited in `CONTEXT_GENERATION_PROFILES` inside `second_question_production.py`, not passed on the CLI.
