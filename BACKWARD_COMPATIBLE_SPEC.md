# 🔧 BEVE Typed Schema - Backward Compatible Implementation Strategy

**Date:** 16 Ekim 2025  
**Goal:** Add Typed Schema as OPTIONAL extension without breaking existing parsers

---

## ✅ Strateji: Extension System Kullan (Spec v1.0'da zaten var!)

### BEVE Spec'deki Extension Mekanizması:

```
## 6 - Extensions

Extensions are considered to be a formal part of the BEVE specification, 
but are not expected to be as broadly implemented.

Following the first three HEADER bits, the next five bits denote various extensions.
```

**Mevcut Extension'lar:**
```c++
0 -> data delimiter
1 -> type tag (variants)
2 -> matrices
3 -> complex numbers
4 -> (RESERVED)
5 -> (RESERVED)
```

---

## 💡 Çözüm: Yeni Extension Tanımları

### Extension 4: Field Index
```
Header: 0x86 | (4 << 3) = 0xA6

Purpose: Single object with field offset table
Use case: Fast partial read for individual objects
```

### Extension 5: Typed Object Array
```
Header: 0x86 | (5 << 3) = 0xAE

Purpose: Array of objects with shared schema
Use case: Reduce key duplication in arrays
```

### Extension 6: Nested Typed Array
```
Header: 0x86 | (6 << 3) = 0xB6

Purpose: Array of nested objects with hierarchical schema
Use case: Deep nesting optimization
```

---

## 🔄 Backward Compatibility Stratejisi

### 1. Fallback Encoding (Default Behavior)

**Mevcut davranış korunur:**
```go
// Default: Generic encoding (spec v1.0 compliant)
data, _ := beve.Marshal(&users)
// → Generic array (type 5, header 0x85)
```

**Yeni davranış (opt-in):**
```go
// Opt-in: Typed schema encoding
data, _ := beve.MarshalWithOptions(&users, MarshalOptions{
    UseTypedSchema: true,
})
// → Typed array (extension 5, header 0xAE)
```

---

### 2. Parser Compatibility Layers

#### Old Parser (BEVE v1.0):

**Extension'ı tanımaz, ne yapar?**

**Spec'den:**
```
Extensions are not expected to be as broadly implemented.
```

**Önerilen davranış:**
```go
// Old parser extension görünce:
if header == 0xAE && !supportsTypedSchema {
    // Option 1: Reject (safe)
    return ErrUnsupportedExtension
    
    // Option 2: Skip and try decode (risky)
    // NOT RECOMMENDED - data format farklı
}
```

**Problem:** Old parser typed array'ı decode edemez çünkü format farklı!

---

#### Solution: Hybrid Encoding Flag

**Spec'e eklenecek:**
```
Extension Compatibility Flag (Optional):

A parser MAY include a compatibility flag in the extension header
to indicate whether the data can be decoded by non-extension parsers.

Bit 7 of extension header:
  0 -> Extension-only (requires parser support)
  1 -> Backward-compatible fallback available
```

**Example:**
```
0xAE (0b10101110)  // Typed array, no fallback
0xEE (0b11101110)  // Typed array, WITH fallback data
```

**With fallback:**
```
Layout:
  0xEE                    // Typed array with fallback
  [Typed Schema Data]     // Extension data
  0xFF                    // Fallback delimiter
  [Generic Array Data]    // Fallback for old parsers
```

**Parser behavior:**
```go
func Unmarshal(data []byte) {
    header := data[0]
    
    if header == 0xEE {
        if supportsTypedSchema {
            // Use typed schema (faster)
            return decodeTypedArray(data)
        } else {
            // Skip to fallback
            offset := findFallbackDelimiter(data)
            return decodeGenericArray(data[offset:])
        }
    }
}
```

---

### 3. Opt-In with Feature Detection

**API Design:**

```go
// Check if parser supports extension
if beve.SupportsExtension(beve.ExtTypedSchema) {
    data, _ = beve.MarshalTyped(&users)
} else {
    data, _ = beve.Marshal(&users)  // Fallback
}

// Auto-detect
data, _ = beve.MarshalAuto(&users)  // Uses typed if available
```

**Unmarshal auto-detection:**
```go
var users []User
err := beve.Unmarshal(data, &users)
// Automatically detects extension and uses appropriate decoder
```

---

## 📋 Spec Addition Proposal

### New Section: "Extension 4-6: Typed Schemas"

```markdown
## Extension 4: Field Index

**Header:** `0x86 | (4 << 3)` = `0xA6`

**Purpose:** Provide fast field access for individual objects via offset table.

**Layout:**
```
HEADER           1 byte    0xA6
OBJECT_TYPE      1 byte    0x03 (object)
FIELD_COUNT      varint    Number of fields
INDEX_TABLE      variable  Field offset table
FIELD_DATA       variable  Field values

Index Entry (7 bytes per field):
  offset:  4 bytes (uint32, relative to FIELD_DATA start)
  size:    2 bytes (uint16, 0 = variable)
  flags:   1 byte  (omitempty, nested, etc.)
```

**Compatibility:**
- Old parsers MUST reject data with header `0xA6`
- Implementers SHOULD provide fallback encoding for compatibility

**Use Case:**
```go
// Fast partial read
age, _ := beve.ReadField(data, "age")
```

---

## Extension 5: Typed Object Array

**Header:** `0x86 | (5 << 3)` = `0xAE`

**Purpose:** Reduce key duplication in arrays of homogeneous objects.

**Layout:**
```
HEADER           1 byte    0xAE
FIELD_COUNT      varint    Number of fields in schema
FIELD_SCHEMA     variable  Field name definitions (once)
ARRAY_SIZE       varint    Number of objects
OBJECT_DATA      variable  Values only (no keys per object)

Field Schema Entry:
  name_length:  varint
  name_data:    UTF-8 bytes
  type_hint:    1 byte (optional, for validation)
```

**Example:**
```
// User struct with 3 fields
0xAE              // Typed array header
0x0C              // 3 fields
  0x08 "id"       // Field 0: "id"
  0x10 "name"     // Field 1: "name"
  0x0c "age"      // Field 2: "age"
0x0C              // 3 objects
  [values only]   // Object 0
  [values only]   // Object 1
  [values only]   // Object 2
```

**Compatibility:**
- Old parsers MUST reject data with header `0xAE`
- For compatibility, use header `0xEE` with fallback data (see Hybrid Encoding)

**Performance:**
- Marshal: 2-3× faster
- Unmarshal: 3-4× faster
- Size: 50-56% reduction

---

## Extension 6: Nested Typed Array

**Header:** `0x86 | (6 << 3)` = `0xB6`

**Purpose:** Optimize deeply nested object arrays.

**Layout:**
```
HEADER           1 byte    0xB6
SCHEMA_DEPTH     varint    Nesting depth
SCHEMA_TABLE     variable  Hierarchical schema definitions
ARRAY_SIZE       varint    Number of root objects
OBJECT_DATA      variable  Nested values only

Schema Table Entry:
  schema_id:     varint    (0 = root, 1+ = nested)
  field_count:   varint
  field_schema:  variable  (name, type, nested_schema_id)
```

**Example:**
```
// User with nested Address
0xB6              // Nested typed array
0x08              // Depth = 2
  // Schema 0 (User)
  0x00            // Schema ID: 0 (root)
  0x0C            // 3 fields
    0x08 "id"     type=int64
    0x10 "name"   type=string
    0x1c "address" type=nested, schema_id=1
  
  // Schema 1 (Address)
  0x04            // Schema ID: 1
  0x0C            // 3 fields
    0x18 "street" type=string
    0x10 "city"   type=string
    0x0c "zip"    type=string

0x0C              // 3 objects
  [nested values only]
```

**Compatibility:**
- Old parsers MUST reject data with header `0xB6`
- For compatibility, use hybrid encoding with fallback

**Performance:**
- Marshal: 2.7-2.85× faster
- Unmarshal: 3.0-3.2× faster
- Size: 56-63% reduction (increases with depth)
```

---

## 🔧 Implementation Guidelines

### Phase 1: Core Extensions (Required)

```go
package beve

// Extension IDs
const (
    ExtFieldIndex      = 4
    ExtTypedArray      = 5
    ExtNestedTypedArray = 6
)

// Marshal options
type MarshalOptions struct {
    UseTypedSchema     bool  // Enable extension 5/6
    UseFieldIndex      bool  // Enable extension 4
    IncludeFallback    bool  // Include generic encoding for old parsers
}

// Default: Backward compatible
var DefaultOptions = MarshalOptions{
    UseTypedSchema:  false,  // Opt-in
    UseFieldIndex:   false,  // Opt-in
    IncludeFallback: true,   // Safe default
}

// Marshal with typed schema
func MarshalTyped(v interface{}) ([]byte, error) {
    return MarshalWithOptions(v, MarshalOptions{
        UseTypedSchema: true,
    })
}

// Auto-detect best encoding
func MarshalAuto(v interface{}) ([]byte, error) {
    // Heuristic: Use typed schema for arrays of structs
    if isArrayOfStructs(v) && arraySize(v) >= 5 {
        return MarshalTyped(v)
    }
    return Marshal(v)  // Default generic
}
```

### Phase 2: Feature Detection

```go
// Parser capabilities
type ParserCapabilities struct {
    SupportsFieldIndex      bool
    SupportsTypedArray      bool
    SupportsNestedTyped     bool
}

// Check capabilities
func GetCapabilities() ParserCapabilities {
    return ParserCapabilities{
        SupportsFieldIndex:  true,
        SupportsTypedArray:  true,
        SupportsNestedTyped: true,
    }
}

// Version negotiation
func NegotiateFormat(producer, consumer ParserCapabilities) MarshalOptions {
    return MarshalOptions{
        UseTypedSchema: producer.SupportsTypedArray && consumer.SupportsTypedArray,
        UseFieldIndex:  producer.SupportsFieldIndex && consumer.SupportsFieldIndex,
    }
}
```

### Phase 3: Hybrid Encoding (Advanced)

```go
// Encode with fallback for old parsers
func MarshalHybrid(v interface{}) ([]byte, error) {
    opts := MarshalOptions{
        UseTypedSchema:  true,
        IncludeFallback: true,
    }
    
    // Encode typed version
    typedData, _ := encodeTyped(v, opts)
    
    // Encode generic fallback
    genericData, _ := encodeGeneric(v)
    
    // Combine
    return appendHybrid(typedData, genericData), nil
}

// Unmarshal auto-detects format
func Unmarshal(data []byte, v interface{}) error {
    header := data[0]
    
    switch {
    case header == 0xAE:  // Typed array
        return unmarshalTyped(data, v)
    case header == 0xEE:  // Hybrid
        if SupportsExtension(ExtTypedArray) {
            return unmarshalTyped(data, v)
        }
        return unmarshalGenericFromHybrid(data, v)
    case header == 0x85:  // Generic array
        return unmarshalGeneric(data, v)
    default:
        return ErrUnsupportedFormat
    }
}
```

---

## 📊 Deployment Strategy

### Stage 1: Opt-In (Safe)

```
Week 1-4: Release with opt-in flag
- Default: Generic encoding (v1.0)
- Optional: Typed schema (extension 5/6)
- Users explicitly enable: beve.MarshalTyped()
```

### Stage 2: Auto-Detection (Smart)

```
Week 5-8: Add auto-detection heuristics
- Small arrays (N<5): Generic (overhead not worth it)
- Large arrays (N≥5): Typed (big gains)
- Users can override: beve.MarshalAuto()
```

### Stage 3: Hybrid Encoding (Compatible)

```
Week 9-12: Add hybrid mode
- Encode both typed + generic
- Old parsers: Use fallback
- New parsers: Use typed (faster)
- Size overhead: ~10% (worth it for compatibility)
```

### Stage 4: Default Switch (Future)

```
Version 2.0 (6+ months later):
- Default: Typed schema
- Fallback: Generic (deprecated but supported)
- Users can force generic: beve.MarshalLegacy()
```

---

## ⚠️ Migration Path

### For Library Authors:

```go
// v1.x (current) - Backward compatible
import "github.com/beve-org/beve-go"

data, _ := beve.Marshal(&users)  // Generic (v1.0)

// v1.x (opt-in) - After extension release
data, _ := beve.MarshalTyped(&users)  // Typed (extension)

// v2.0 (future) - Typed by default
data, _ := beve.Marshal(&users)  // Typed (default)
data, _ := beve.MarshalLegacy(&users)  // Generic (if needed)
```

### For End Users:

**No breaking changes!**
```go
// Existing code continues to work
var users []User
err := beve.Unmarshal(data, &users)
// Automatically handles both generic and typed formats
```

---

## ✅ Spec Amendment Proposal

### Add to BEVE Specification v1.1:

```markdown
# Extension Compatibility Guidelines

## Parser Requirements

**MUST:**
1. Parsers MUST reject unknown extension headers with error
2. Parsers MUST NOT attempt to decode unsupported extensions

**SHOULD:**
1. Parsers SHOULD provide feature detection API
2. Parsers SHOULD support graceful degradation

**MAY:**
1. Parsers MAY support hybrid encoding (0xEE header)
2. Parsers MAY implement auto-format detection

## Encoder Requirements

**MUST:**
1. Encoders MUST provide generic encoding (v1.0) by default
2. Encoders MUST document which extensions are supported

**SHOULD:**
1. Encoders SHOULD provide opt-in flags for extensions
2. Encoders SHOULD support version negotiation

**MAY:**
1. Encoders MAY implement auto-optimization heuristics
2. Encoders MAY provide hybrid encoding for compatibility

## Extension Registration

New extensions MUST be proposed via:
1. GitHub issue with specification draft
2. Reference implementation in at least one language
3. Performance benchmarks demonstrating benefit
4. Backward compatibility analysis
```

---

## 🎯 Final Recommendation

### ✅ EVET, Opsiyonel Ekleme Mümkün!

**Strateji:**
1. ✅ **Extension system kullan** (spec'de zaten var)
2. ✅ **Opt-in başla** (breaking change yok)
3. ✅ **Auto-detection ekle** (smart defaults)
4. ✅ **Hybrid encoding** (max compatibility)
5. ✅ **Version 2.0'da default yap** (uzun vadeli plan)

**Avantajlar:**
- ✅ Mevcut parser'lar break olmaz
- ✅ Yeni feature'lar opt-in
- ✅ Performance kazancı isteyenler hemen kullanır
- ✅ Yavaş yavaş adoption artar
- ✅ Eventually everyone benefits

**Implementation priority:**
```
Phase 1 (1 month):  Extension 5 (Typed Array)
Phase 2 (1 month):  Extension 4 (Field Index)
Phase 3 (2 months): Extension 6 (Nested Typed)
Phase 4 (2 months): Hybrid encoding
Phase 5 (Future):   Default switch in v2.0
```

**Sonuç:** Backward compatible, opt-in, gradual adoption! 🎯✅
