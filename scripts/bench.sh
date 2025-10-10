#!/usr/bin/env bash
set -euo pipefail

ROOT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
OUT_DIR="${ROOT_DIR}/benchmarks"
OUT_FILE="${OUT_DIR}/latest.md"
JSON_OUT_FILE="${OUT_DIR}/latest.json"

mkdir -p "${OUT_DIR}"

timestamp="$(date -u '+%Y-%m-%dT%H:%M:%SZ')"
host_name="$(hostname)"
os_name="$(uname -s)"
kernel="$(uname -srv)"
arch="$(uname -m)"
cpu_brand="$( (sysctl -n machdep.cpu.brand_string 2>/dev/null || lscpu | awk -F: '/Model name/ {gsub(/^ +| +$/, "", $2); print $2; exit}' 2>/dev/null || uname -p || echo 'unknown') | tr -s ' ' )"
goversion="$(go version)"
git_rev="$(git -C "${ROOT_DIR}" rev-parse --short HEAD 2>/dev/null || echo 'unknown')"

header_tmp=$(mktemp)
tmp_files=("${header_tmp}")

json_tmp=$(mktemp)
tmp_files+=("${json_tmp}")
exec 3>"${json_tmp}"

json_escape() {
  local str="${1-}"
  str="${str//\\/\\\\}"
  str="${str//\"/\\\"}"
  str="${str//$'\n'/\\n}"
  str="${str//$'\r'/\\r}"
  str="${str//$'\t'/\\t}"
  printf '%s' "${str}"
}

cleanup() {
  # Close JSON descriptor if still open
  if { true >&3; } 2>/dev/null; then
    exec 3>&-
  fi
  rm -f "${tmp_files[@]}"
}
trap cleanup EXIT

host_json="$(json_escape "${host_name}")"
os_json="$(json_escape "${os_name}")"
kernel_json="$(json_escape "${kernel}")"
arch_json="$(json_escape "${arch}")"
cpu_json="$(json_escape "${cpu_brand}")"
go_json="$(json_escape "${goversion}")"
git_json="$(json_escape "${git_rev}")"

cat <<EOF >"${header_tmp}"
# BEVE Benchmark Snapshot

> Generated: ${timestamp}
> Hostname: ${host_name}
> OS: ${os_name}
> Kernel: ${kernel}
> Architecture: ${arch}
> CPU: ${cpu_brand}
> Go: ${goversion}
> Git: ${git_rev}

Metrics below cover BEVE alongside CBOR, Sonic, MessagePack, and Go's encoding/json implementations.

EOF

cat <<EOF >&3
{
  "generated_at": "${timestamp}",
  "environment": {
    "hostname": "${host_json}",
    "os": "${os_json}",
    "kernel": "${kernel_json}",
    "architecture": "${arch_json}",
    "cpu": "${cpu_json}",
    "go_version": "${go_json}",
    "git_revision": "${git_json}"
  },
  "results": [
EOF

first_json_entry=1


append_tmp_file() {
  tmp_files+=("$1")
}

format_command() {
  local formatted=""
  local part
  for part in "$@"; do
    if [[ -z "${formatted}" ]]; then
      formatted="$(printf '%q' "${part}")"
    else
      formatted+=" $(printf '%q' "${part}")"
    fi
  done
  printf '%s' "${formatted}"
}

run_bench() {
  local scenario="$1"
  local codec="$2"
  local operation="$3"
  shift 3
  local command=("$@")
  local label="${scenario} · ${operation} · ${codec}"

  echo "Running ${label}..." >&2

  local tmp_out
  tmp_out=$(mktemp)
  append_tmp_file "${tmp_out}"

  (cd "${ROOT_DIR}" && "${command[@]}") | tee "${tmp_out}"
  local cmd_status=${PIPESTATUS[0]}
  if [[ ${cmd_status} -ne 0 ]]; then
    echo "Benchmark failed: ${label}" >&2
    exit 1
  fi

  local bench_line
  bench_line="$(grep '^Benchmark' "${tmp_out}" | head -n1 || true)"

  local ns="n/a" bytes="n/a" allocs="n/a"
  if [[ -n "${bench_line}" ]]; then
    local metrics
    metrics="$(awk '{
      for (i = 1; i <= NF; ++i) {
        if ($(i+1) == "ns/op") ns = $(i);
        if ($(i+1) == "B/op") bytes = $(i);
        if ($(i+1) == "allocs/op") allocs = $(i);
      }
    }
    END {
      if (ns == "") ns = "n/a";
      if (bytes == "") bytes = "n/a";
      if (allocs == "") allocs = "n/a";
      printf "%s %s %s", ns, bytes, allocs;
    }' <<<"${bench_line}" || true)"
    if [[ -n "${metrics}" ]]; then
      set -- ${metrics}
      ns="$1"
      bytes="$2"
      allocs="$3"
    fi
  fi

  local ns_json bytes_json allocs_json
  if [[ "${ns}" == "n/a" ]]; then
    ns_json="null"
  else
    ns_json="${ns}"
  fi
  if [[ "${bytes}" == "n/a" ]]; then
    bytes_json="null"
  else
    bytes_json="${bytes}"
  fi
  if [[ "${allocs}" == "n/a" ]]; then
    allocs_json="null"
  else
    allocs_json="${allocs}"
  fi

  local scenario_json codec_json operation_json command_json
  scenario_json="$(json_escape "${scenario}")"
  codec_json="$(json_escape "${codec}")"
  operation_json="$(json_escape "${operation}")"
  command_json="$(json_escape "$(format_command "${command[@]}")")"

  if [[ ${first_json_entry} -eq 0 ]]; then
    printf ',\n' >&3
  else
    first_json_entry=0
  fi
  printf '    {"scenario":"%s","codec":"%s","operation":"%s","ns_per_op":%s,"bytes_per_op":%s,"allocs_per_op":%s,"command":"%s"}' \
    "${scenario_json}" "${codec_json}" "${operation_json}" \
    "${ns_json}" "${bytes_json}" "${allocs_json}" "${command_json}" >&3
}

# Small struct benchmarks (marshal + unmarshal)
run_bench "Small Struct" "BEVE" "Marshal" go test '-bench=^BenchmarkSmallStruct_BEVE_Marshal$' -benchmem -benchtime=10000x -run=^$ ./...
run_bench "Small Struct" "BEVE ZeroCopy" "Marshal" go test '-bench=^BenchmarkSmallStruct_BEVE_MarshalZeroCopy$' -benchmem -benchtime=10000x -run=^$ ./...
run_bench "Small Struct" "CBOR" "Marshal" go test '-bench=^BenchmarkSmallStruct_CBOR_Marshal$' -benchmem -benchtime=10000x -run=^$ ./...
run_bench "Small Struct" "JSON" "Marshal" go test '-bench=^BenchmarkSmallStruct_JSON_Marshal$' -benchmem -benchtime=10000x -run=^$ ./...
run_bench "Small Struct" "Sonic" "Marshal" go test '-bench=^BenchmarkSmallStruct_Sonic_Marshal$' -benchmem -benchtime=10000x -run=^$ ./...
run_bench "Small Struct" "MessagePack" "Marshal" go test '-bench=^BenchmarkSmallStruct_MessagePack_Marshal$' -benchmem -benchtime=10000x -run=^$ ./...

run_bench "Small Struct" "BEVE" "Unmarshal" go test '-bench=^BenchmarkSmallStruct_BEVE_Unmarshal$' -benchmem -benchtime=10000x -run=^$ ./...
run_bench "Small Struct" "CBOR" "Unmarshal" go test '-bench=^BenchmarkSmallStruct_CBOR_Unmarshal$' -benchmem -benchtime=10000x -run=^$ ./...
run_bench "Small Struct" "JSON" "Unmarshal" go test '-bench=^BenchmarkSmallStruct_JSON_Unmarshal$' -benchmem -benchtime=10000x -run=^$ ./...
run_bench "Small Struct" "Sonic" "Unmarshal" go test '-bench=^BenchmarkSmallStruct_Sonic_Unmarshal$' -benchmem -benchtime=10000x -run=^$ ./...
run_bench "Small Struct" "MessagePack" "Unmarshal" go test '-bench=^BenchmarkSmallStruct_MessagePack_Unmarshal$' -benchmem -benchtime=10000x -run=^$ ./...

# Medium payload marshal
run_bench "Medium Payload" "BEVE" "Marshal" go test '-bench=^BenchmarkMedium_BEVE_Marshal$' -benchmem -benchtime=5000x -run=^$ ./...
run_bench "Medium Payload" "BEVE ZeroCopy" "Marshal" go test '-bench=^BenchmarkMedium_BEVE_MarshalZeroCopy$' -benchmem -benchtime=5000x -run=^$ ./...
run_bench "Medium Payload" "CBOR" "Marshal" go test '-bench=^BenchmarkMedium_CBOR_Marshal$' -benchmem -benchtime=5000x -run=^$ ./...
run_bench "Medium Payload" "JSON" "Marshal" go test '-bench=^BenchmarkMedium_JSON_Marshal$' -benchmem -benchtime=5000x -run=^$ ./...
run_bench "Medium Payload" "Sonic" "Marshal" go test '-bench=^BenchmarkMedium_Sonic_Marshal$' -benchmem -benchtime=5000x -run=^$ ./...
run_bench "Medium Payload" "MessagePack" "Marshal" go test '-bench=^BenchmarkMedium_MessagePack_Marshal$' -benchmem -benchtime=5000x -run=^$ ./...

# Medium payload unmarshal
run_bench "Medium Payload" "BEVE" "Unmarshal" go test '-bench=^BenchmarkMedium_BEVE_Unmarshal$' -benchmem -benchtime=5000x -run=^$ ./...
run_bench "Medium Payload" "JSON" "Unmarshal" go test '-bench=^BenchmarkMedium_JSON_Unmarshal$' -benchmem -benchtime=5000x -run=^$ ./...
run_bench "Medium Payload" "Sonic" "Unmarshal" go test '-bench=^BenchmarkMedium_Sonic_Unmarshal$' -benchmem -benchtime=5000x -run=^$ ./...
run_bench "Medium Payload" "CBOR" "Unmarshal" go test '-bench=^BenchmarkMedium_CBOR_Unmarshal$' -benchmem -benchtime=5000x -run=^$ ./...
run_bench "Medium Payload" "MessagePack" "Unmarshal" go test '-bench=^BenchmarkMedium_MessagePack_Unmarshal$' -benchmem -benchtime=5000x -run=^$ ./...

# Large payload marshal
run_bench "Large Payload" "BEVE" "Marshal" go test '-bench=^BenchmarkLarge_BEVE_Marshal$' -benchmem -benchtime=3000x -run=^$ ./...
run_bench "Large Payload" "BEVE ZeroCopy" "Marshal" go test '-bench=^BenchmarkLarge_BEVE_MarshalZeroCopy$' -benchmem -benchtime=3000x -run=^$ ./...
run_bench "Large Payload" "CBOR" "Marshal" go test '-bench=^BenchmarkLarge_CBOR_Marshal$' -benchmem -benchtime=3000x -run=^$ ./...
run_bench "Large Payload" "JSON" "Marshal" go test '-bench=^BenchmarkLarge_JSON_Marshal$' -benchmem -benchtime=3000x -run=^$ ./...
run_bench "Large Payload" "Sonic" "Marshal" go test '-bench=^BenchmarkLarge_Sonic_Marshal$' -benchmem -benchtime=3000x -run=^$ ./...
run_bench "Large Payload" "MessagePack" "Marshal" go test '-bench=^BenchmarkLarge_MessagePack_Marshal$' -benchmem -benchtime=3000x -run=^$ ./...

# Large payload unmarshal
run_bench "Large Payload" "BEVE" "Unmarshal" go test '-bench=^BenchmarkLarge_BEVE_Unmarshal$' -benchmem -benchtime=3000x -run=^$ ./...
run_bench "Large Payload" "JSON" "Unmarshal" go test '-bench=^BenchmarkLarge_JSON_Unmarshal$' -benchmem -benchtime=3000x -run=^$ ./...
run_bench "Large Payload" "Sonic" "Unmarshal" go test '-bench=^BenchmarkLarge_Sonic_Unmarshal$' -benchmem -benchtime=3000x -run=^$ ./...
run_bench "Large Payload" "CBOR" "Unmarshal" go test '-bench=^BenchmarkLarge_CBOR_Unmarshal$' -benchmem -benchtime=3000x -run=^$ ./...
run_bench "Large Payload" "MessagePack" "Unmarshal" go test '-bench=^BenchmarkLarge_MessagePack_Unmarshal$' -benchmem -benchtime=3000x -run=^$ ./...

printf '\n  ]\n}\n' >&3
exec 3>&-
mv "${json_tmp}" "${JSON_OUT_FILE}"

python - <<'PY' "${header_tmp}" "${JSON_OUT_FILE}" "${OUT_FILE}"
import json
import math
import sys
from pathlib import Path

header_path = Path(sys.argv[1])
json_path = Path(sys.argv[2])
out_path = Path(sys.argv[3])

header = header_path.read_text()
data = json.loads(json_path.read_text())
results = data.get("results", [])

order = []
groups = {}

for entry in results:
  scenario = entry.get("scenario", "")
  operation = entry.get("operation", "")
  key = (scenario, operation)
  if key not in groups:
    groups[key] = []
    order.append(key)
  groups[key].append(entry)

def fmt(value):
  if value is None:
    return "n/a"
  if isinstance(value, (int, float)) and math.isfinite(value):
    if isinstance(value, int) or value.is_integer():
      return f"{int(value)}"
    return f"{value:.3f}".rstrip("0").rstrip(".")
  return str(value)

lines = [header.rstrip(), "", "## Summary", "", "| Scenario | Codec | Operation | ns/op | B/op | allocs/op |", "|----------|-------|-----------|-------|------|-----------|"]

for key in order:
  scenario, operation = key
  entries = groups[key]
  sorted_entries = sorted(entries, key=lambda e: (float("inf") if e.get("ns_per_op") in (None, "n/a") else e["ns_per_op"]))
  for entry in sorted_entries:
    lines.append(f"| {scenario} | {entry.get('codec', '')} | {operation} | {fmt(entry.get('ns_per_op'))} | {fmt(entry.get('bytes_per_op'))} | {fmt(entry.get('allocs_per_op'))} |")

lines.extend(["", "## Commands", "", "| Scenario | Codec | Operation | Command |", "|----------|-------|-----------|---------|"])

for key in order:
  scenario, operation = key
  entries = groups[key]
  sorted_entries = sorted(entries, key=lambda e: (float("inf") if e.get("ns_per_op") in (None, "n/a") else e["ns_per_op"]))
  for entry in sorted_entries:
    cmd = entry.get("command", "")
    lines.append(f"| {scenario} | {entry.get('codec', '')} | {operation} | `{cmd}` |")

out_path.write_text("\n".join(lines) + "\n")
PY

printf '\nBenchmark report written to %s\n' "${OUT_FILE}" >&2
printf 'Benchmark JSON written to %s\n' "${JSON_OUT_FILE}" >&2
