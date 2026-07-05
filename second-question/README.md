# The Second Question of Technology production kit

This version treats AI video generation as context-visual generation only.

The deterministic edit owns:

- voiceover and final timing
- all on-screen labels and subtitles
- official logos and end-card layout
- specific images such as Angelica Kauffmann's `Invention`
- final clip duration, resolution, fps, and audio muxing
- spreadsheet/DDEX-style metadata outputs

AI-generated clips are deliberately replaceable. Assembly strips generated audio, loops or trims short clips to the scripted scene duration, burns labels in post, and composites the local end card.

## Provider/model configuration

Model selection is intentionally not a CLI option. Edit the constants near the top of `second_question_production.py`:

```python
CONTEXT_GENERATION_PROFILES = {
    "sora": {
        "provider": "sora",
        "model": DEFAULT_SORA_MODEL,
        "request_seconds": 8,
        "size": "1280x720",
    },
    "gemini_omni": {
        "provider": "gemini_omni",
        "model": DEFAULT_GEMINI_OMNI_MODEL,
        "aspect_ratio": "16:9",
        "delivery": "uri",
    },
}

DEFAULT_FULL_AI_PROFILE = "sora"
COMPARE_CONTEXT_PROFILES = ("sora", "gemini_omni")
COMPARE_CONTEXT_SCENE_IDS = (3, 5)
LOCKED_LOCAL_SCENE_IDS = {1, 8, 9}
```

Scenes 1, 8, and 9 are locked local renders by default. Scene 1 preserves the cold open. Scene 8 preserves the Kauffmann image beat. Scene 9 preserves the final alternatives-network beat.

## Install

```bash
python -m pip install -r requirements_second_question.txt
```

You also need `ffmpeg` and `ffprobe` on your PATH.

## Environment variables

For Sora context clips and OpenAI TTS:

```bash
export OPENAI_API_KEY="your_openai_key"
```

For Gemini Omni Flash context clips:

```bash
export GEMINI_API_KEY="your_gemini_key"
```

`GOOGLE_API_KEY` is accepted as a fallback for Gemini.

## Dry-run the provider requests

```bash
python second_question_production.py generate-clips --sora --dry-run
python second_question_production.py generate-clips --omni --dry-run
python second_question_production.py compare-context-generators --dry-run
```

Dry runs write JSONL request plans without calling provider APIs.

## Generate the same test scenes with both Omni and Sora

The comparison command uses `COMPARE_CONTEXT_SCENE_IDS` and `COMPARE_CONTEXT_PROFILES` from the code. It does not expose model selection as a CLI option.

```bash
python second_question_production.py download-assets
python second_question_production.py compare-context-generators
```

Outputs:

```text
second_question_build/clips/compare/sora/
second_question_build/clips/compare/omni/
second_question_build/clips/compare/README_compare_context_generators.md
```

## Generate a full Sora-backed cut

```bash
python second_question_production.py download-assets
python second_question_production.py metadata
python second_question_production.py generate-clips --sora
python second_question_production.py assemble --sora --tts --force
```

## Generate a full Gemini Omni-backed cut

```bash
python second_question_production.py download-assets
python second_question_production.py metadata
python second_question_production.py generate-clips --omni
python second_question_production.py assemble --omni --tts --force
```

## Fast local mock cut

```bash
python second_question_production.py all --mock --no-tts
```

## Important architecture notes

- Generated clip audio is stripped during normalization.
- Generated clips can be shorter than the scripted scene. The assembler loops or trims them to each scene's target duration.
- Exact labels live in `assembly/labels.ass` and are burned after the picture edit.
- Official logos are never generated. Add uploaded official logo overrides here if downloads fail or you prefer specific official files:

```text
second_question_build/assets/logos/osi_uploaded.png
second_question_build/assets/logos/oshw_uploaded.png
second_question_build/assets/logos/beagleboard_uploaded.png
```

- The DDEX XML remains a draft metadata crosswalk, not a recipient-profile-validated delivery package.

## What is included in this completed kit

```text
second_question_production.py          Main production script
requirements_second_question.txt       Python dependencies
metadata/second_question_metadata.xlsx Workbook: release, scenes, sources, providers, rights, DDEX crosswalk
metadata/*.csv                         Flat metadata exports
metadata/source_manifest.json          Machine-readable source and scene manifest
metadata/ern43_draft_not_validated.xml Draft ERN-style XML crosswalk
request_plans/                         Dry-run Sora and Gemini Omni request plans
docs/LICENSE_SELECTION_NOTE.md         Reminder to choose final Creative Commons license
```

For macOS Python 3.9 environments that use LibreSSL, the requirements file pins `urllib3<2` to avoid the urllib3 v2 `NotOpenSSLWarning`.
