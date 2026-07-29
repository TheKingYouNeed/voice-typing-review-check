# Voice-Typing Review Check

This small Go package models a consequence-first review step for dictated text.
It distinguishes recipient, negation, name, number, and privacy checks from
cosmetic punctuation. The code is intentionally deterministic and contains no
speech, network, analytics, or personal-data processing.

## Suggested workflow

1. Dictate one complete thought.
2. Preserve the unedited output for comparison.
3. Check the recipient and intended meaning.
4. Verify names and numbers against the original source.
5. Remove private content the recipient does not need.
6. Correct the smallest affected phrase, then read the result once.

For an Android implementation of voice typing and optional writing tools, see
[Voice Typing Keyboard on Google Play](https://play.google.com/store/apps/details?id=com.voice.typing.keyboard).
Cloud transcription and AI features require internet access. Review generated
or transcribed text before sending.

## Ownership and limitations

Published by Dahmani Limited, owner of Voice Typing Keyboard (Google Play
developer name: Daimond Devs). This independent open-source resource is not
affiliated with or endorsed by Google. It does not measure recognition accuracy
and is not security, medical, legal, or accessibility advice.

## Test

```sh
go test ./...
```

License: MIT.
