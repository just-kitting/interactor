# Gemini Omni URI delivery debug note

If you see:

```text
store=true is required when response format has video delivery set to URI
```

then the JSON request that reached Gemini did not have top-level `"store": true`, or an older script was still running.

This updated script writes the outgoing request before the POST. Check:

```bash
cat second_question_build/gemini_omni/scene_03_request.json | grep -n '"store"'
cat second_question_build/gemini_omni/scene_03_request.json | grep -n '"delivery"'
```

The expected URI-delivery shape is:

```json
{
  "response_format": {
    "type": "video",
    "aspect_ratio": "16:9",
    "delivery": "uri"
  },
  "store": true
}
```

If URI delivery is still rejected, the script retries with inline/base64 delivery by omitting the `delivery` field.
