# Windows Encoding Fix for parse_benchmarks.py

## Problem

Windows CI/CD benchmark job was failing with encoding error:

```
Traceback (most recent call last):
  File "D:\a\beve-go\beve-go\scripts\parse_benchmarks.py", line 276, in <module>
    main()
  File "D:\a\beve-go\beve-go\scripts\parse_benchmarks.py", line 261, in main
    print(f"\u2705 Parsed {len(results)} benchmark results")
  File "C:\hostedtoolcache\windows\Python\3.11.9\x64\Lib\encodings\cp1252.py", line 19, in encode
    return codecs.charmap_encode(input,self.errors,encoding_table)[0]
UnicodeEncodeError: 'charmap' codec can't encode character '\u2705' in position 0: character maps to <undefined>
```

## Root Cause

Windows terminal uses CP1252 encoding by default, which cannot display Unicode emoji characters like ✅, ❌, ⚠️, 🎉, 📁.

## Solution

### 1. Force UTF-8 Encoding on Windows

Added UTF-8 encoding wrapper at the top of the script:

```python
# Force UTF-8 encoding on Windows
if sys.platform == 'win32':
    import io
    sys.stdout = io.TextIOWrapper(sys.stdout.buffer, encoding='utf-8')
    sys.stderr = io.TextIOWrapper(sys.stderr.buffer, encoding='utf-8')
```

### 2. Replace Emoji with ASCII Tags

Changed all emoji characters to ASCII equivalents:

| Before | After |
|--------|-------|
| ✅ | `[SUCCESS]` |
| ❌ | `[ERROR]` |
| ⚠️ | `[WARNING]` |
| 🎉 | `[DONE]` |
| 📁 | `[INFO]` |

## Changes Made

### File: `scripts/parse_benchmarks.py`

**Line 1-2**: Added encoding declaration
```python
#!/usr/bin/env python3
# -*- coding: utf-8 -*-
```

**Lines 13-17**: Added Windows UTF-8 wrapper
```python
# Force UTF-8 encoding on Windows
if sys.platform == 'win32':
    import io
    sys.stdout = io.TextIOWrapper(sys.stdout.buffer, encoding='utf-8')
    sys.stderr = io.TextIOWrapper(sys.stderr.buffer, encoding='utf-8')
```

**All print statements**: Replaced emoji with ASCII tags
- Line 127: `✅ Generated Markdown` → `[SUCCESS] Generated Markdown`
- Line 144: `✅ Generated JSON` → `[SUCCESS] Generated JSON`
- Line 155: `⚠️  matplotlib not available` → `[WARNING] matplotlib not available`
- Line 166: `⚠️  No scenarios to chart` → `[WARNING] No scenarios to chart`
- Line 239: `✅ Generated Chart` → `[SUCCESS] Generated Chart`
- Line 251: `❌ Error: Raw file not found` → `[ERROR] Raw file not found`
- Line 261: `❌ Error: No benchmark results` → `[ERROR] No benchmark results`
- Line 268: `✅ Parsed N benchmark results` → `[SUCCESS] Parsed N benchmark results`
- Line 278: `🎉 Benchmark reports generated` → `[DONE] Benchmark reports generated`
- Line 279: `📁 Output directory` → `[INFO] Output directory`

## Validation

### Local Test (macOS)
```bash
$ python3 scripts/parse_benchmarks.py benchmarks/latest_raw.txt benchmarks/
[WARNING] matplotlib not available, skipping chart generation
[SUCCESS] Parsed 33 benchmark results
[SUCCESS] Generated JSON: benchmarks/latest.json
[SUCCESS] Generated Markdown: benchmarks/latest.md

[DONE] Benchmark reports generated successfully!
[INFO] Output directory: benchmarks
```

### Expected Windows Output
Script should now work on Windows CI/CD without UnicodeEncodeError.

## Benefits

1. **Cross-Platform Compatibility**: Works on Windows, macOS, and Linux
2. **No External Dependencies**: Pure Python solution
3. **Backward Compatible**: ASCII tags are readable on all platforms
4. **Fail-Safe**: UTF-8 wrapper ensures proper encoding handling

## Testing

To test on Windows:
```powershell
python scripts\parse_benchmarks.py benchmarks\latest_raw.txt benchmarks\
```

To test CI/CD:
```bash
git add scripts/parse_benchmarks.py
git commit -m "fix: Windows encoding error in parse_benchmarks.py"
git push origin main
```

Watch GitHub Actions benchmark job on Windows runner.

## Related Files

- `scripts/parse_benchmarks.py` - Main script with fix
- `scripts/bench.sh` - Calls parse_benchmarks.py
- `.github/workflows/benchmarks.yml` - CI/CD workflow

## References

- [Python Windows UTF-8 Mode](https://peps.python.org/pep-0540/)
- [Unicode on Windows Terminal](https://docs.microsoft.com/en-us/windows/console/console-virtual-terminal-sequences)
- [CP1252 Encoding](https://en.wikipedia.org/wiki/Windows-1252)
