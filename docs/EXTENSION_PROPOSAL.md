# BEVE Extension Proposal: Time & Temporal Types

**Version**: 1.0  
**Date**: October 14, 202### 5 - Duration

Time span without specific start/end points.

**Layout**: `HEADER | SIGN_PRECISION | SECONDS | SUB_SECOND`

**HEADER**: `0b00101'110` (1 byte)tatus**: Draft  
**Authors**: BEVE Go Contributors

## Abstract

This proposal extends the BEVE v1.0 specification to support temporal data types (timestamps, durations, time zones) as formal extensions. These types are ubiquitous in modern applications and deserve efficient, standardized binary representations.

## Motivation

Current workarounds for temporal data in BEVE:
- **Problem 1**: `time.Time` encoded as `int64` Unix nanoseconds loses timezone information
- **Problem 2**: No standard representation for durations, intervals, or recurring events
- **Problem 3**: Applications must implement custom solutions, breaking interoperability

### Use Cases
- **APIs**: ISO 8601 timestamps with timezone awareness
- **Databases**: Efficient temporal indexing and queries
- **IoT/Telemetry**: High-precision event timestamps
- **Scheduling**: Recurring events, cron-like expressions
- **Distributed Systems**: Clock synchronization, causality tracking

## Specification

### Extension Header (0x06)

Following BEVE v1.0 spec section 6 (Extensions), we use the reserved extension space:

```c++
6 -> extensions                            0b00000'110
```

The next 5 bits denote temporal extension types:

```c++
4 -> timestamp                             0b00100'110
5 -> duration                              0b00101'110
6 -> interval                              0b00110'110
7 -> recurring event (cron-like)           0b00111'110
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

| Type | BEVE Extension | JSON (ISO 8601) | MessagePack | CBOR |
|------|----------------|-----------------|-------------|------|
| UTC Timestamp (ns) | 14 bytes | ~30 bytes | 12 bytes | 13 bytes |
| Timestamp + TZ | 16 bytes | ~36 bytes | 14 bytes | 15 bytes |
| Duration | 14 bytes | ~20 bytes | 12 bytes | 13 bytes |

**Result**: BEVE extensions are comparable to MessagePack/CBOR while maintaining human-readable JSON conversion.

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

### Phase 1: Go Implementation (v1.4.0)
- [ ] Core extension types (timestamp UTC, duration)
- [ ] time.Time auto-detection and encoding
- [ ] Benchmark vs current int64 approach
- [ ] Documentation and examples

### Phase 2: Extended Support (v1.5.0)
- [ ] Timezone-aware timestamps
- [ ] Interval type
- [ ] JavaScript/TypeScript library
- [ ] Python library

### Phase 3: Advanced Features (v2.0.0)
- [ ] Recurring events
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

## Open Questions

1. **Should we support leap seconds explicitly?**
   - Proposal: Follow Unix epoch semantics (no explicit support)
   
2. **Should recurring events use cron syntax or custom format?**
   - Proposal: Custom binary format for efficiency, document cron conversion

3. **Should we support calendar systems beyond Gregorian?**
   - Proposal: Not in v1.0, defer to v2.0 if demand exists

## References

- BEVE Specification v1.0: [SPECIFICATION.md](../SPECIFICATION.md)
- ISO 8601: Date and time format standard
- RFC 3339: Date/Time on the Internet
- MessagePack Timestamp Extension: https://github.com/msgpack/msgpack/blob/master/spec.md#timestamp-extension-type
- CBOR Tags: https://www.rfc-editor.org/rfc/rfc8949.html#name-standard-date-time-string

## Changelog

- **2025-10-14**: Initial proposal draft
- **TBD**: Community feedback period
- **TBD**: Implementation in beve-go v1.4.0

---

**Proposal Status**: 📝 **DRAFT** - Ready for community review

**Next Steps**:
1. Community discussion on GitHub Discussions
2. Prototype implementation in beve-go
3. Benchmark validation
4. Specification update PR

**Contributors welcome!** Join the discussion at: https://github.com/stephenberry/eve/discussions
