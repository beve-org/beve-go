# Bench.sh Fix: Empty Results Issue

## 🐛 Problem

Bench.sh script was creating empty `results` arrays in `latest.json`:

```json
{
  "generated_at": "2025-10-14T14:36:19Z",
  "environment": { ... },
  "results": []  ← EMPTY!
}
```

**Error Message:**
```
ls: /var/folders/.../tmp.xxx/result_*.json: No such file or directory
[DEBUG] Total result files: 0
[DEBUG] Processed 0 result files into JSON
```

## 🔍 Root Cause

**`./...` glob pattern was including `examples/` directory with broken dependencies:**

```bash
# Original command:
go test -bench='^Benchmark...$' -benchmem -benchtime=30000x -run=^$ ./...

# Output:
# github.com/beve-org/beve-go/examples/axios-interceptor
examples/axios-interceptor/server-fiber.go:9:2: no required module provides package github.com/gofiber/fiber/v2
FAIL    github.com/beve-org/beve-go/examples/axios-interceptor [setup failed]
                                                                ^^^^^^^^^^^^^^
                                                                Exit code != 0!
```

**Consequence:**
1. Test command fails (non-zero exit code)
2. `run_bench` function hits `return 1`
3. Background job terminates without writing result file
4. `results_dir/` stays empty
5. `latest.json` has empty `results` array

## ✅ Solution

### 1. **Fixed Test Command Pattern**

Changed from `./...` (all packages recursively) to `.` (root package only):

```bash
# Before:
run_bench "Small Struct" "BEVE" "Marshal" go test ... -run=^$ ./...

# After:
run_bench "Small Struct" "BEVE" "Marshal" go test ... -run=^$ .
```

### 2. **Improved Error Handling**

```bash
# Before:
if [[ ${cmd_status} -ne 0 ]]; then
  echo "[FAILED] ${label}" >&2
  rm -f "${tmp_out}"
  return 1  ← Terminated job without writing result!
fi

# After:
if [[ ${cmd_status} -ne 0 ]]; then
  echo "[FAILED] ${label} (exit code: ${cmd_status})" >&2
  echo "[FAILED] Last output:" >&2
  tail -5 "${tmp_out}" >&2
  rm -f "${tmp_out}"
  # Don't return - let the function continue to write a null result
  local bench_line=""
else
  local bench_line
  bench_line="$(grep '^Benchmark' "${tmp_out}" | head -n1 || true)"
fi
```

Now failed benchmarks write `null` results instead of terminating silently.

### 3. **Enhanced Debug Output**

```bash
# Before:
echo "[FAILED] ${label}" >&2

# After:
echo "[FAILED] ${label} (exit code: ${cmd_status})" >&2
echo "[FAILED] Last output:" >&2
tail -5 "${tmp_out}" >&2
```

Shows **why** benchmark failed for easier debugging.

### 4. **Added Result File Validation**

```bash
[DEBUG] Writing to: /tmp/result_abc123.json
[DEBUG] File created successfully: 245 bytes
[DEBUG] Results directory: /tmp/xyz
[DEBUG] Total result files: 33
[DEBUG] Processed 33 result files into JSON
```

## 🧪 Testing

### Before Fix:

```bash
$ ./scripts/bench.sh
[START] Small Struct · Marshal · BEVE
[START] Small Struct · Marshal · JSON
...
[DEBUG] Total result files: 0  ← EMPTY!
[DEBUG] Processed 0 result files into JSON

$ cat benchmarks/latest.json
{
  "results": []
}
```

### After Fix:

```bash
$ ./scripts/bench.sh
[START] Small Struct · Marshal · BEVE
[START] Small Struct · Marshal · JSON
...
[DONE] Small Struct · Marshal · BEVE
[DONE] Small Struct · Marshal · JSON
...
[DEBUG] Total result files: 33  ← SUCCESS!
[DEBUG] Processed 33 result files into JSON

$ cat benchmarks/latest.json
{
  "results": [
    {"scenario":"Small Struct","codec":"BEVE","operation":"Marshal",...},
    ...
  ]
}
```

## 📊 Impact

| Metric | Before | After |
|--------|--------|-------|
| **Result files created** | 0 | 33 |
| **JSON results array** | Empty | 33 entries |
| **Chart generation** | ❌ Failed | ✅ Success |
| **Error visibility** | ❌ Silent failures | ✅ Detailed logs |

## 🎯 Files Modified

- `scripts/bench.sh`
  - Changed all `./...` → `.` (33 occurrences)
  - Improved error handling (no `return 1` on failure)
  - Added debug output for failures
  - Enhanced result validation

## 🚀 Verification

1. ✅ **Local test passes** - 33 benchmarks complete successfully
2. ✅ **Result files created** - All benchmarks write to `results_dir/`
3. ✅ **JSON populated** - `latest.json` contains all results
4. ✅ **Chart renders** - `plot_benchmarks.py` works without errors
5. ✅ **Failed benchmarks logged** - Shows exit code + last output

## 📝 Lessons Learned

1. **`./...` is dangerous** - Can include broken example code
2. **Silent failures are evil** - Always log failures before returning
3. **Background jobs need care** - `return` terminates job without cleanup
4. **Debug early, debug often** - Validation checks catch issues immediately

---

**Author:** GitHub Copilot  
**Date:** 2025-10-14  
**Issue:** Empty benchmark results in CI/CD workflow  
**Status:** ✅ FIXED
