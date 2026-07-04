# The Second Question of Technology - Production Script

This kit contains a runnable Python production helper for assembling the source-grounded explainer video **The Second Question of Technology**.

## What it does

- Downloads source pages, public/reference assets, and official logo assets where direct downloads are available.
- Creates either Sora-generated clips through the OpenAI Video API or local placeholder timing clips.
- Builds exact on-screen labels in post, avoiding unreliable generated text.
- Creates optional OpenAI TTS voiceover.
- Assembles clips, labels, audio, and an official-logo end card with ffmpeg.
- Exports metadata as XLSX, CSVs, source manifests, and a draft DDEX ERN-style XML crosswalk.

## Requirements

- Python 3.10+
- ffmpeg and ffprobe on PATH
- Python packages in `requirements_second_question.txt`
- `OPENAI_API_KEY` set for Sora and/or TTS generation

Install Python dependencies:

```bash
python -m pip install -r requirements_second_question.txt
```

## Typical commands

Download source/audit assets and logo files:

```bash
python second_question_production.py download-assets
```

Create metadata workbook, CSVs, and draft DDEX XML:

```bash
python second_question_production.py metadata
```

Generate Sora clips:

```bash
python second_question_production.py generate-clips --sora
```

Generate TTS voiceover:

```bash
python second_question_production.py voiceover
```

Assemble final video:

```bash
python second_question_production.py assemble --tts
```

Fast local mock cut without API calls:

```bash
python second_question_production.py all --mock --no-tts
```

## Important notes

- The DDEX XML is intentionally labeled as a draft crosswalk. It is not recipient-profile-validated and still requires real identifiers such as DPID, ISRC, UPC/GRid, contributor identifiers, and recipient-specific business rules.
- Sora should generate mechanism visuals only. Exact labels, citations, credits, and official logos are composited in post.
- Use only official logo files. Do not ask the video model to create or redraw logos.
