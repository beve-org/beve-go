# BEVE Extension Proposal: Essential Data Types

**Version**: 1.0  
**Date**: October 14, 2025  
**Status**: Draft  
**Authors**: BEVE Go Contributors

## Abstract

This proposal extends the BEVE v1.0 specification with **high-performance, widely-used data types** that deserve native binary representations:
- **Temporal Types**: Timestamps, durations, intervals (with optional timezone)
- **Identifiers**: UUID/ULID (128-bit, 55% smaller than string)
- **Patterns**: Regular expressions (validation, search, config)

These extensions are **performance-focused** and **commonly used** in modern distributed systems, APIs, and databases.

## Motivation

### Why These Types?

✅ **Performance Impact**: These types appear in 90%+ of modern APIs  
✅ **Space Efficiency**: 30-55% smaller than JSON string representations  
✅ **Semantic Meaning**: Binary format preserves type information (UUID ≠ random string)  
✅ **Real-World Usage**: UUID, timestamps, and regex are in MessagePack/CBOR for a reason

### Current Problems

**Temporal Data**:
- `time.Time` as `int64` loses timezone → breaks user-facing apps
- No standard for durations/intervals → each app reinvents the wheel

**UUIDs**:
- String `"550e8400-e29b-41d4-a716-446655440000"` = 36 bytes
- Binary `0x550e8400e29b41d4a716446655440000` = 16 bytes (**55% savings**)
- Databases use binary UUIDs internally anyway (PostgreSQL, MongoDB)

**Regular Expressions**:
- Validation schemas sent as strings → no semantic meaning
- Pattern rules in config files → verbose and repetitive

### Use Cases
- **APIs**: ISO 8601 timestamps with timezone, UUID entity IDs
- **Databases**: Binary UUID primary keys, temporal indexing
- **IoT/Telemetry**: High-precision timestamps, compact identifiers
- **Distributed Systems**: Trace IDs (OpenTelemetry), correlation tokens
- **Validation**: Email/phone regex patterns, input sanitization

## Specification

### Extension Header (0x06)

Following BEVE v1.0 spec section 6 (Extensions), we use the reserved extension space:

```c++
6 -> extensions                            0b00000'110
```

The next 5 bits denote extension types:

```c++
4 -> timestamp                             0b00100'110
5 -> duration                              0b00101'110
6 -> interval                              0b00110'110
7 -> recurring event (cron-like)           0b00111'110
8 -> UUID/ULID (128-bit identifier)        0b01000'110
9 -> regular expression                    0b01001'110
```

### 4 - Timestamp

High-precision timestamp with **optional timezone offset**. UTC assumed if timezone not specified.

**Design Philosophy**:
- ✅ **UTC by default**: If no timezone specified, UTC assumed (like MessagePack)
- ✅ **Optional timezone**: Can store offset when known (like CBOR Tag 0)
- ✅ **Compact**: Timezone only 2 bytes when needed
- ✅ **Interoperable**: Supports both PostgreSQL `TIMESTAMP` and `TIMESTAMPTZ` models
- ✅ **Best of both**: Flexibility + simplicity

**Layout**: `HEADER | PRECISION | EPOCH_SECONDS | SUB_SECOND [| TZ_OFFSET]`

**HEADER**: `0b00100'110` (1 byte)

**PRECISION** (1 byte):
```c++
Bit 0:   Has timezone offset (0=no/UTC, 1=yes)
Bits 1-3: Precision
  0 -> seconds only
  1 -> milliseconds (3 decimal places)
  2 -> microseconds (6 decimal places)
  3 -> nanoseconds  (9 decimal places)
```

**EPOCH_SECONDS**: Signed 64-bit integer (8 bytes, little-endian)
- Unix epoch (1970-01-01T00:00:00Z)
- Range: ~292 billion years before/after epoch

**SUB_SECOND**: Unsigned integer (precision-dependent)
```c++
milliseconds: uint16_t (2 bytes) -> 0-999
microseconds: uint32_t (4 bytes) -> 0-999,999
nanoseconds:  uint32_t (4 bytes) -> 0-999,999,999
```

**TZ_OFFSET**: Signed 16-bit integer (2 bytes, optional)
- Minutes offset from UTC (-1439 to +1439)
- Only present if precision bit 0 is set
- Examples: UTC+5:30 = +330, UTC-8:00 = -480, UTC = 0

**Total Size**:
- Seconds (UTC):      9 bytes (header + precision + epoch)
- Seconds (with TZ):  11 bytes (+ 2 bytes offset)
- Nanoseconds (UTC):  14 bytes
- Nanoseconds (TZ):   16 bytes (+ 2 bytes offset)

**JSON Representation** (RFC 3339):
```json
"2025-10-14T15:30:45.123456789Z"          // UTC (no timezone offset)
"2025-10-14T10:30:45.123456789-05:00"     // With timezone offset
```

**Timezone Handling**:
- **No timezone** (bit 0 = 0): Assumed UTC, like MessagePack
- **With timezone** (bit 0 = 1): Explicit offset stored, like CBOR Tag 0
- **Storage flexibility**: Application chooses based on data source
- **Display**: Application converts to local timezone when rendering
- **Example**: 
  - Server without timezone info → stores UTC (9-14 bytes)
  - User input with timezone → stores offset (11-16 bytes)
  - IoT device → always UTC (smaller payload)

### 5 - Duration

Time span without specific start/end points.

**Layout**: `HEADER | SIGN_PRECISION | SECONDS | SUB_SECOND`

**HEADER**: `0b00110'110` (1 byte)

**SIGN_PRECISION** (1 byte):
```c++
Bit 0:   Sign (0=positive, 1=negative)
Bits 1-7: Precision (same as timestamp)
```

**SECONDS**: Unsigned 64-bit integer (8 bytes)

**SUB_SECOND**: Same as timestamp (0, 2, or 4 bytes)

**Total Size**: 10-14 bytes

**JSON Representation**:
```json
{
  "duration": "PT2H30M45.123S"
}
```
(ISO 8601 duration format)

### 6 - Interval

Time range between two timestamps.

**Layout**: `HEADER | START_TIMESTAMP | END_TIMESTAMP`

**HEADER**: `0b00110'110` (1 byte)

**Timestamps**: Two UTC timestamps (compact encoding)

**Total Size**: 19-29 bytes (2× timestamp size + 1)

**JSON Representation**:
```json
{
  "start": "2025-10-14T00:00:00Z",
  "end": "2025-10-14T23:59:59Z"
}
```

### 7 - Recurring Event (Experimental)

Compact cron-like expression for recurring events.

**Layout**: `HEADER | CRON_TYPE | CRON_DATA`

**HEADER**: `0b00111'110` (1 byte)

**CRON_TYPE** (1 byte):
```c++
0 -> daily (at specific time)
1 -> weekly (day + time)
2 -> monthly (date + time)
3 -> custom cron expression
```

**CRON_DATA**: Variable length based on type

**JSON Representation**:
```json
{
  "recurrence": "0 9 * * 1-5",
  "description": "Every weekday at 9:00 AM"
}
```

## Implementation Guide

### Go Implementation

```go
// Extension types
const (
    ExtTimestamp         = 0x26 // 0b00100'110
    ExtDuration          = 0x2E // 0b00101'110
    ExtInterval          = 0x36 // 0b00110'110
    ExtRecurringEvent    = 0x3E // 0b00111'110
)

// Precision flags
const (
    PrecisionSeconds      = 0 << 1
    PrecisionMilliseconds = 1 << 1
    PrecisionMicroseconds = 2 << 1
    PrecisionNanoseconds  = 3 << 1
    
    FlagHasTimezone = 0x01 // Bit 0: timezone offset present
)

// Precision constants
const (
    PrecisionSeconds      = 0
    PrecisionMilliseconds = 1
    PrecisionMicroseconds = 2
    PrecisionNanoseconds  = 3
)

// Timestamp encodes a timestamp with optional timezone
type Timestamp struct {
    Seconds         int64
    Nanoseconds     uint32
    TimezoneOffset  *int16  // nil = UTC, otherwise minutes from UTC
}

func (e *Encoder) EncodeTimestamp(ts Timestamp) error {
    // Header
    if err := e.WriteByte(ExtTimestamp); err != nil {
        return err
    }
    
    // Precision + timezone flag
    precision := PrecisionNanoseconds
    if ts.TimezoneOffset != nil {
        precision |= FlagHasTimezone
    }
    if err := e.WriteByte(precision); err != nil {
        return err
    }
    
    // Epoch seconds (little-endian)
    epochBuf := make([]byte, 8)
    binary.LittleEndian.PutUint64(epochBuf, uint64(ts.Seconds))
    if err := e.WriteBytes(epochBuf); err != nil {
        return err
    }
    
    // Nanoseconds (little-endian)
    nanoBuf := make([]byte, 4)
    binary.LittleEndian.PutUint32(nanoBuf, ts.Nanoseconds)
    if err := e.WriteBytes(nanoBuf); err != nil {
        return err
    }
    
    // Optional timezone offset (little-endian)
    if ts.TimezoneOffset != nil {
        tzBuf := make([]byte, 2)
        binary.LittleEndian.PutUint16(tzBuf, uint16(*ts.TimezoneOffset))
        return e.WriteBytes(tzBuf)
    }
    
    return nil
}

// Helper: Create UTC timestamp (no timezone)
func NewTimestampUTC(seconds int64, nanos uint32) Timestamp {
    return Timestamp{Seconds: seconds, Nanoseconds: nanos, TimezoneOffset: nil}
}

// Helper: Create timestamp with timezone
func NewTimestampWithTZ(seconds int64, nanos uint32, offsetMinutes int16) Timestamp {
    return Timestamp{Seconds: seconds, Nanoseconds: nanos, TimezoneOffset: &offsetMinutes}
}
```

### JavaScript/TypeScript Implementation

```typescript
interface Timestamp {
  seconds: bigint;
  nanoseconds: number;
  timezoneOffset?: number; // minutes from UTC (optional)
}

function encodeTimestamp(ts: Timestamp): Uint8Array {
  const hasTimezone = ts.timezoneOffset !== undefined;
  const size = hasTimezone ? 16 : 14;
  const buffer = new Uint8Array(size);
  const view = new DataView(buffer.buffer);
  
  // Header + precision
  buffer[0] = 0x26; // ExtTimestamp
  buffer[1] = (3 << 1) | (hasTimezone ? 1 : 0); // Nanoseconds + timezone flag
  
  // Epoch seconds (little-endian)
  view.setBigInt64(2, ts.seconds, true);
  
  // Nanoseconds (little-endian)
  view.setUint32(10, ts.nanoseconds, true);
  
  // Optional timezone offset (little-endian)
  if (hasTimezone) {
    view.setInt16(14, ts.timezoneOffset!, true);
  }
  
  return buffer;
}

// Helper: Convert Date to BEVE Timestamp (UTC)
function dateToTimestampUTC(date: Date): Timestamp {
  const millis = date.getTime();
  const seconds = BigInt(Math.floor(millis / 1000));
  const nanoseconds = (millis % 1000) * 1_000_000;
  return { seconds, nanoseconds };
}

// Helper: Convert Date with timezone to BEVE Timestamp
function dateToTimestampWithTZ(date: Date): Timestamp {
  const millis = date.getTime();
  const seconds = BigInt(Math.floor(millis / 1000));
  const nanoseconds = (millis % 1000) * 1_000_000;
  const timezoneOffset = -date.getTimezoneOffset(); // JavaScript returns negative of UTC offset
  return { seconds, nanoseconds, timezoneOffset };
}
```

### Python Implementation

```python
import struct
from datetime import datetime, timezone, timedelta

EXT_TIMESTAMP = 0x26
PRECISION_NANOSECONDS = 3 << 1
FLAG_HAS_TIMEZONE = 0x01

def encode_timestamp(dt: datetime) -> bytes:
    """Encode datetime to BEVE timestamp extension with optional timezone."""
    # Convert to Unix epoch
    if dt.tzinfo is None:
        # No timezone: assume UTC
        epoch = dt.replace(tzinfo=timezone.utc).timestamp()
        tz_offset = None
    else:
        # Has timezone: extract offset
        epoch = dt.timestamp()
        tz_offset = int(dt.utcoffset().total_seconds() / 60)  # minutes
    
    seconds = int(epoch)
    nanoseconds = int((epoch - seconds) * 1_000_000_000)
    
    # Precision + timezone flag
    precision = PRECISION_NANOSECONDS | (FLAG_HAS_TIMEZONE if tz_offset is not None else 0)
    
    # Pack: header + precision + seconds + nanoseconds [+ timezone]
    if tz_offset is not None:
        return struct.pack('<BBqIh', 
            EXT_TIMESTAMP,
            precision,
            seconds,
            nanoseconds,
            tz_offset
        )
    else:
        return struct.pack('<BBqI', 
            EXT_TIMESTAMP,
            precision,
            seconds,
            nanoseconds
        )

def decode_timestamp(data: bytes) -> datetime:
    """Decode BEVE timestamp to datetime."""
    header, precision = struct.unpack_from('<BB', data, 0)
    has_tz = bool(precision & FLAG_HAS_TIMEZONE)
    
    if has_tz:
        seconds, nanos, tz_offset = struct.unpack('<qIh', data[2:])
        tz = timezone(timedelta(minutes=tz_offset))
    else:
        seconds, nanos = struct.unpack('<qI', data[2:])
        tz = timezone.utc
    
    epoch = seconds + (nanos / 1_000_000_000)
    return datetime.fromtimestamp(epoch, tz=tz)
```

---

## Extension 8: UUID/ULID (128-bit Identifier)

### Motivation

**Problem**: 
- UUID strings waste 36 bytes (with dashes) or 32 bytes (hex only)
- Binary representation is 16 bytes = **55% space savings**
- UUIDs are ubiquitous in distributed systems, databases, APIs

**Use Cases**:
- Database primary keys (PostgreSQL UUID type)
- Distributed tracing (OpenTelemetry trace IDs)
- Session tokens and API keys
- Message queue correlation IDs
- Microservice entity IDs

### Binary Layout

```
HEADER (1 byte) | VERSION_FLAGS (1 byte) | UUID_BYTES (16 bytes)
```

**Total Size**: 18 bytes (vs 36 bytes string)

#### VERSION_FLAGS Byte

```
Bits 0-3: UUID version (4 = random, 6 = sortable, 7 = Unix timestamp)
Bits 4-7: Reserved (must be 0)
```

Common values:
```c++
0x04 -> UUID v4 (random)          // Most common
0x01 -> UUID v1 (timestamp+MAC)
0x06 -> UUID v6 (reordered v1)    // Sortable
0x07 -> UUID v7 (Unix timestamp)  // ULID-like
0x08 -> ULID (Universally Unique Lexicographically Sortable ID)
```

### Example

UUID `550e8400-e29b-41d4-a716-446655440000` becomes:

```
Header:      0x48 (0b01000'110)
Version:     0x04 (UUID v4)
Bytes[0-15]: 55 0e 84 00 e2 9b 41 d4 a7 16 44 66 55 44 00 00
```

### Implementation Examples

#### Go

```go
import "github.com/google/uuid"

func encodeUUID(u uuid.UUID) []byte {
    buf := make([]byte, 18)
    buf[0] = 0x48  // HEADER
    buf[1] = 0x04  // UUID v4
    copy(buf[2:], u[:])
    return buf
}

func decodeUUID(data []byte) uuid.UUID {
    var u uuid.UUID
    copy(u[:], data[2:18])
    return u
}
```

#### TypeScript

```typescript
import { v4 as uuidv4, parse as uuidParse } from 'uuid';

function encodeUUID(uuidString: string): Uint8Array {
    const buf = new Uint8Array(18);
    buf[0] = 0x48;  // HEADER
    buf[1] = 0x04;  // UUID v4
    
    const bytes = uuidParse(uuidString);
    buf.set(bytes, 2);
    return buf;
}

function decodeUUID(data: Uint8Array): string {
    const bytes = data.slice(2, 18);
    return Array.from(bytes)
        .map(b => b.toString(16).padStart(2, '0'))
        .join('')
        .replace(/(.{8})(.{4})(.{4})(.{4})(.{12})/, '$1-$2-$3-$4-$5');
}
```

#### Python

```python
import uuid
import struct

def encode_uuid(u: uuid.UUID) -> bytes:
    """Encode UUID to BEVE binary."""
    return struct.pack('<BB16s', 0x48, 0x04, u.bytes)

def decode_uuid(data: bytes) -> uuid.UUID:
    """Decode BEVE binary to UUID."""
    _, version, uuid_bytes = struct.unpack('<BB16s', data[:18])
    return uuid.UUID(bytes=uuid_bytes)
```

### JSON Mapping

```json
{
  "id": "550e8400-e29b-41d4-a716-446655440000",
  "type": "uuid-v4"
}
```

Or simplified:
```json
"550e8400-e29b-41d4-a716-446655440000"
```

### Size Comparison

| Format | Size | Example |
|--------|------|---------|
| **BEVE Binary** | **18 bytes** | `48 04 55 0e 84 00 ...` |
| JSON (with dashes) | 38 bytes | `"550e8400-e29b-41d4-a716-446655440000"` |
| JSON (hex only) | 34 bytes | `"550e8400e29b41d4a716446655440000"` |
| MessagePack (fixext 16) | 18 bytes | Same as BEVE |
| CBOR (Tag 37) | 19 bytes | 1 byte tag + 1 byte size + 16 bytes |

**Result**: 47% smaller than JSON string representation

---

## Extension 9: Regular Expression

### Motivation

**Problem**: 
- Validation rules sent repeatedly as strings
- Pattern matching rules in config files
- API request/response validation schemas
- Search patterns in query languages

**Use Cases**:
- Input validation (email, phone, URL patterns)
- API schema definitions (OpenAPI, JSON Schema)
- Rule engines and policy enforcement
- Log parsing and filtering
- Search query patterns

### Binary Layout

```
HEADER (1 byte) | SYNTAX_FLAGS (1 byte) | PATTERN_SIZE | PATTERN_UTF8
```

**Size**: Variable (typically 3 + pattern length)

#### SYNTAX_FLAGS Byte

```
Bit 0: Case insensitive (i)
Bit 1: Multiline (m)
Bit 2: Dot matches newline (s)
Bit 3: Extended syntax (x)
Bit 4: Unicode-aware (u)
Bit 5-7: Reserved (must be 0)
```

Common flag combinations:
```c++
0x01 -> Case insensitive only (/pattern/i)
0x03 -> Case insensitive + multiline (/pattern/im)
0x11 -> Case insensitive + Unicode (/pattern/iu)
```

#### PATTERN_SIZE

Uses BEVE's compressed unsigned integer format (same as string SIZE).

### Examples

#### Email Validation Pattern

Pattern: `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

```
Header:       0x49 (0b01001'110)
Flags:        0x00 (no flags)
Size:         0x96 (compressed: 54 < 64, fits in 1 byte with 2-bit size indicator)
Pattern:      "^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\\.[a-zA-Z]{2,}$" (UTF-8)
```

**Total Size**: 3 + 54 = 57 bytes (vs 56 bytes as plain string, but with semantic meaning)

#### Case-Insensitive Search

Pattern: `hello world` (case insensitive)

```
Header:       0x49
Flags:        0x01 (case insensitive)
Size:         0x2C (11 bytes)
Pattern:      "hello world"
```

**Total Size**: 3 + 11 = 14 bytes

### Implementation Examples

#### Go

```go
import "regexp"

func encodeRegex(pattern string, flags int) []byte {
    patternBytes := []byte(pattern)
    size := len(patternBytes)
    
    // Simplified: assuming size < 64 for example
    buf := make([]byte, 3+size)
    buf[0] = 0x49  // HEADER
    buf[1] = byte(flags)
    buf[2] = byte(size << 2)  // SIZE with 2-bit indicator
    copy(buf[3:], patternBytes)
    return buf
}

func decodeRegex(data []byte) (*regexp.Regexp, error) {
    flags := data[1]
    size := data[2] >> 2
    pattern := string(data[3 : 3+size])
    
    // Apply flags
    if flags&0x01 != 0 {
        pattern = "(?i)" + pattern  // Case insensitive
    }
    if flags&0x02 != 0 {
        pattern = "(?m)" + pattern  // Multiline
    }
    if flags&0x04 != 0 {
        pattern = "(?s)" + pattern  // Dot matches newline
    }
    
    return regexp.Compile(pattern)
}
```

#### TypeScript

```typescript
function encodeRegex(regex: RegExp): Uint8Array {
    const pattern = regex.source;
    const flags = 
        (regex.ignoreCase ? 0x01 : 0) |
        (regex.multiline ? 0x02 : 0) |
        (regex.dotAll ? 0x04 : 0);
    
    const patternBytes = new TextEncoder().encode(pattern);
    const size = patternBytes.length;
    
    const buf = new Uint8Array(3 + size);
    buf[0] = 0x49;  // HEADER
    buf[1] = flags;
    buf[2] = size << 2;  // SIZE
    buf.set(patternBytes, 3);
    return buf;
}

function decodeRegex(data: Uint8Array): RegExp {
    const flags = data[1];
    const size = data[2] >> 2;
    const pattern = new TextDecoder().decode(data.slice(3, 3 + size));
    
    let flagStr = '';
    if (flags & 0x01) flagStr += 'i';
    if (flags & 0x02) flagStr += 'm';
    if (flags & 0x04) flagStr += 's';
    if (flags & 0x10) flagStr += 'u';
    
    return new RegExp(pattern, flagStr);
}
```

#### Python

```python
import re
import struct

def encode_regex(pattern: str, flags: int = 0) -> bytes:
    """Encode regex pattern to BEVE binary."""
    pattern_bytes = pattern.encode('utf-8')
    size = len(pattern_bytes)
    
    # Simplified: assuming size < 64
    return struct.pack('<BBB', 0x49, flags, size << 2) + pattern_bytes

def decode_regex(data: bytes) -> re.Pattern:
    """Decode BEVE binary to compiled regex."""
    _, flags_byte, size_byte = struct.unpack('<BBB', data[:3])
    size = size_byte >> 2
    pattern = data[3:3+size].decode('utf-8')
    
    # Convert BEVE flags to Python re flags
    py_flags = 0
    if flags_byte & 0x01: py_flags |= re.IGNORECASE
    if flags_byte & 0x02: py_flags |= re.MULTILINE
    if flags_byte & 0x04: py_flags |= re.DOTALL
    if flags_byte & 0x08: py_flags |= re.VERBOSE
    
    return re.compile(pattern, py_flags)
```

### JSON Mapping

```json
{
  "pattern": "^[a-z]+$",
  "flags": ["i", "m"]
}
```

Or simplified (JavaScript-like):
```json
"/^[a-z]+$/im"
```

### Use Case Examples

**API Validation Schema**:
```json
{
  "email": {
    "type": "string",
    "pattern": "^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\\.[a-zA-Z]{2,}$"
  }
}
```

Becomes more compact in BEVE binary (57 bytes vs ~100+ bytes JSON overhead).

**Log Filter Config**:
```json
{
  "filters": [
    {"pattern": "ERROR|FATAL", "flags": ["i"]},
    {"pattern": "^\\d{4}-\\d{2}-\\d{2}", "flags": []}
  ]
}
```

Each pattern stored efficiently with semantic meaning preserved.

---

## Compatibility

### Backward Compatibility
✅ **Fully backward compatible** with BEVE v1.0
- Uses reserved extension space (0x06)
- Decoders without extension support can skip unknown types
- Fallback: Encode as int64 Unix nanos (current workaround)

### Forward Compatibility
✅ **Extensible design**
- 5-bit extension type space (32 possible types)
- 28 types still available for future use
- Version negotiation via metadata (optional)

## Performance Analysis

### Size Comparison

| Type | BEVE Extension | JSON | MessagePack | CBOR |
|------|----------------|------|-------------|------|
| UTC Timestamp (ns) | 14 bytes | ~30 bytes | 12 bytes | 13 bytes |
| Timestamp + TZ | 16 bytes | ~36 bytes | 14 bytes | 15 bytes |
| Duration | 14 bytes | ~20 bytes | 12 bytes | 13 bytes |
| UUID | 18 bytes | 38 bytes | 18 bytes | 19 bytes |
| RegExp (avg) | ~50 bytes | ~100 bytes | ~50 bytes | ~50 bytes |

**Result**: BEVE extensions provide **30-50% space savings** vs JSON while maintaining semantic meaning.

### Speed Benchmarks (Estimated)

| Operation | BEVE Extension | time.Time (int64) | JSON (RFC3339) |
|-----------|----------------|-------------------|----------------|
| Encode | ~10 ns | ~8 ns | ~150 ns |
| Decode | ~12 ns | ~10 ns | ~200 ns |

**Trade-off**: ~2-4 ns overhead for timezone/precision support vs int64, but 15-20× faster than JSON.

## Security Considerations

1. **Integer Overflow**: Use 64-bit signed integers for epoch (safe until year ~292 billion)
2. **Timezone Validation**: Validate offset range (-1439 to +1439 minutes)
3. **Precision Loss**: Document precision limits (nanoseconds = 9 decimal places)
4. **Leap Seconds**: Not explicitly handled (follow Unix epoch semantics)

## Migration Path

### Phase 1: Go Implementation (v1.4.0) - **HIGH PRIORITY**
- [ ] Timestamp (UTC + optional timezone)
- [ ] Duration
- [ ] UUID/ULID (128-bit identifier)
- [ ] time.Time auto-detection and encoding
- [ ] Benchmark vs current int64 approach
- [ ] Documentation and examples

### Phase 2: Extended Support (v1.5.0)
- [ ] Regular Expression type
- [ ] Interval type
- [ ] JavaScript/TypeScript library
- [ ] Python library

### Phase 3: Advanced Features (v2.0.0) - **LOW PRIORITY**
- [ ] Recurring events (cron-like)
- [ ] Calendar-aware operations
- [ ] Multi-language support (Rust, Java, etc.)

## Design Decisions

### ✅ Optional Timezone Offset

**Rationale**:
- **Hybrid approach**: Best of both MessagePack (UTC only) and CBOR (timezone support)
- **Storage efficiency**: UTC timestamps save 2 bytes when timezone unknown
- **Context preservation**: Can store user's original timezone when known
- **Real-world flexibility**: Matches how applications actually handle timestamps

**Comparison**:

| Format | Timezone Support | Size (ns) | Trade-off |
|--------|------------------|-----------|-----------|
| MessagePack | UTC only | 12 bytes | Simple, but loses timezone context |
| CBOR Tag 0 | RFC 3339 text | ~30 bytes | Full ISO 8601, but large/string-based |
| CBOR Tag 1 | UTC only | ~10 bytes | Compact, but loses timezone context |
| **BEVE** | **Optional** | **14-16 bytes** | **Compact + flexible** |

**Use Case Examples**:
- **IoT sensor data** → UTC only (14 bytes, no timezone needed)
- **User calendar event** → with timezone (16 bytes, preserves "+05:30")
- **Server logs** → UTC only (14 bytes, simpler aggregation)
- **E-commerce order** → with timezone (16 bytes, legal compliance)

### ✅ Precision Field (Variable Sub-Second Resolution)

**Rationale**:
- IoT sensors: milliseconds sufficient (saves 2 bytes)
- Financial systems: microseconds needed
- Scientific computing: nanoseconds required
- Flexibility without waste

### ✅ UUID Binary Format (Not String)

**Rationale**:
- **Performance**: Binary UUIDs are standard in databases (PostgreSQL, Cassandra, MongoDB)
- **Space efficiency**: 18 bytes vs 36 bytes (50% savings)
- **Semantic meaning**: Type system knows it's an ID, not a random string
- **Compatibility**: MessagePack and CBOR also use binary UUID

**Why NOT add everything from CBOR/MessagePack?**:
- ❌ Decimal fractions → Float64 sufficient, niche use case
- ❌ Rational numbers → Scientific computing only, rare
- ❌ URI/URL type → String is fine, app can validate
- ❌ Set type → Array works, app can deduplicate
- ❌ Indefinite-length encoding → Hurts performance (no size prefix)

**BEVE Philosophy**: Only add types that are **ubiquitous + performance-critical**.

## Open Questions

1. **Should we support leap seconds explicitly?**
   - Proposal: Follow Unix epoch semantics (no explicit support)
   
2. **Should recurring events use cron syntax or custom format?**
   - Proposal: Custom binary format for efficiency, document cron conversion

3. **Should we support calendar systems beyond Gregorian?**
   - Proposal: Not in v1.0, defer to v2.0 if demand exists

4. **Should UUID version be validated on decode?**
   - Proposal: Store version but don't validate (allow future UUID formats)

## Summary

This proposal adds **6 high-value extension types** to BEVE:

| Extension | Type | Why? | Space Savings |
|-----------|------|------|---------------|
| **4** | Timestamp | Ubiquitous in APIs/DBs | 14-16 bytes vs ~30 (47%) |
| **5** | Duration | Time spans everywhere | 14 bytes vs ~20 (30%) |
| **6** | Interval | Date ranges, schedules | 30 bytes vs ~50 (40%) |
| **7** | Recurring Event | Cron jobs, calendars | Variable (compact) |
| **8** | UUID/ULID | Database IDs, tracing | 18 bytes vs 36 (50%) |
| **9** | RegExp | Validation, search | Semantic + compact |

**Philosophy**: Only extensions that are **performance-critical** and **widely used**.

**Not included** (intentionally):
- Decimal fractions, rationals, bigfloats → Niche use cases
- URI/URL types → String is sufficient
- Set type → Array + app logic works
- Indefinite-length encoding → Hurts performance

## References

- BEVE Specification v1.0: [SPECIFICATION.md](../SPECIFICATION.md)
- ISO 8601: Date and time format standard
- RFC 3339: Date/Time on the Internet
- RFC 4122: UUID specification
- MessagePack Timestamp Extension: https://github.com/msgpack/msgpack/blob/master/spec.md#timestamp-extension-type
- CBOR Tags: https://www.rfc-editor.org/rfc/rfc8949.html#name-standard-date-time-string
- PCRE2: Regular Expression Syntax

## Changelog

- **2025-10-14**: Initial proposal with temporal types
- **2025-10-14**: Added UUID/ULID and RegExp extensions
- **TBD**: Community feedback period
- **TBD**: Implementation in beve-go v1.4.0

---

**Proposal Status**: 📝 **DRAFT** - Ready for community review

**Next Steps**:
1. Community discussion on GitHub Discussions
2. Prototype implementation in beve-go (Phase 1: Timestamp, Duration, UUID)
3. Benchmark validation vs JSON/MessagePack/CBOR
4. Specification update PR

**Priority Implementation Order**:
1. ✅ **Timestamp + UUID** (most impactful, 90% of use cases)
2. Duration + Interval
3. RegExp (validation use cases)
4. Recurring events (lower priority)

**Contributors welcome!** Join the discussion at: https://github.com/stephenberry/eve/discussions
