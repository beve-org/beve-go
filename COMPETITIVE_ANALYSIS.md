# Rakip Analizi ve Eksik Özellikler

## 📊 Mevcut Durum

### Test Coverage
- **beve-go**: 50.3%
- **core**: 41.5%
- **Hedef**: >80%

---

## 🔍 Rakip Karşılaştırması

### 1. **MessagePack (msgpack)**

#### Bizde OLAN:
- ✅ Varint encoding
- ✅ Typed arrays
- ✅ Struct tags
- ✅ Custom marshaling

#### Bizde OLMAYAN:
- ❌ **Extension Types** - Custom type registration (msgpack ext format)
- ❌ **Streaming API** - Encoder/Decoder for io.Reader/Writer
- ❌ **Time encoding** - Native time.Time support
- ❌ **Binary type** - Explicit binary data type (vs string)
- ❌ **Query/Path API** - Access nested fields without full unmarshal
- ❌ **Schema Evolution** - Graceful handling of missing/extra fields
- ❌ **Append Mode** - Append to existing buffer

**Örnek**:
```go
// MessagePack Extension Types
msgpack.RegisterExt(1, (*MyCustomType)(nil))

// Streaming
enc := msgpack.NewEncoder(conn)
enc.Encode(data)

// Time support
type Event struct {
    Timestamp time.Time `msgpack:"ts"`
}
```

---

### 2. **CBOR (fxamacker/cbor)**

#### Bizde OLAN:
- ✅ Binary format
- ✅ Struct tags
- ✅ Fast encoding

#### Bizde OLMAYAN:
- ❌ **CBOR Tags** - Semantic tagging (dates, bignum, regex, etc.)
- ❌ **Deterministic Encoding** - Canonical sorting for signatures
- ❌ **Streaming Decode** - Partial decode of large payloads
- ❌ **Map Key Sorting** - For deterministic output
- ❌ **Indefinite Length** - Streaming arrays/maps without known size
- ❌ **Big Numbers** - Arbitrary precision integers
- ❌ **RawMessage** - Delayed decoding (like json.RawMessage)
- ❌ **Decode Options** - Strict vs lenient mode, max nesting, etc.

**Örnek**:
```go
// CBOR deterministic encoding
em, _ := cbor.EncOptions{Sort: cbor.SortCanonical}.EncMode()
data := em.Marshal(v)

// CBOR tags
type Timestamped struct {
    Time time.Time `cbor:"1,keyasint"` // CBOR tag 1
}

// RawMessage for delayed decode
type Container struct {
    Metadata map[string]string
    Payload  cbor.RawMessage // decode later
}
```

---

### 3. **Protocol Buffers (protobuf)**

#### Bizde OLAN:
- ✅ Binary format
- ✅ Fast encoding

#### Bizde OLMAYAN:
- ❌ **Schema Definition** - .proto files with versioning
- ❌ **Code Generation** - Type-safe generated code
- ❌ **Backward Compatibility** - Field number-based evolution
- ❌ **Oneof/Union Types** - Discriminated unions
- ❌ **Repeated Fields** - Packed vs unpacked
- ❌ **Well-Known Types** - Timestamp, Duration, Any, etc.
- ❌ **gRPC Integration** - Native RPC support
- ❌ **Reflection API** - Dynamic message manipulation

**Not**: Protobuf schema-based, BEVE schemaless. Farklı use case'ler.

---

### 4. **JSON (encoding/json & sonic)**

#### Bizde OLAN:
- ✅ Struct tags
- ✅ Custom marshaling
- ✅ omitempty

#### Bizde OLMAYAN:
- ❌ **Text Format** - Human-readable (bizde binary)
- ❌ **UseNumber** - Preserve number precision
- ❌ **DisallowUnknownFields** - Strict unmarshal
- ❌ **HTMLEscape** - Safe for HTML embedding
- ❌ **json.RawMessage** - Delayed decode ✅ **Var ama tam değil**
- ❌ **MarshalIndent** - Pretty printing (N/A for binary)
- ❌ **Encoder/Decoder streaming** - ✅ **Kısmen var**
- ❌ **JSONPath/jq-like queries** - Field access without full decode

**Örnek**:
```go
// JSON streaming
enc := json.NewEncoder(w)
enc.Encode(v1)
enc.Encode(v2) // sequential encoding

// RawMessage
type Flexible struct {
    Type    string          `json:"type"`
    Payload json.RawMessage `json:"payload"`
}
```

---

## 🎯 BEVE'de Eksik Olan Öncelikli Özellikler

### **Tier 1: Kritik (Hemen Eklenebilir)**

1. **Streaming API** ⭐⭐⭐⭐⭐
   - `NewEncoder(io.Writer)` / `NewDecoder(io.Reader)`
   - Sequential encoding without buffering
   - Use case: Network protocols, file I/O, large datasets
   ```go
   enc := beve.NewEncoder(conn)
   enc.Encode(msg1)
   enc.Encode(msg2) // no intermediate buffer
   ```

2. **time.Time Native Support** ⭐⭐⭐⭐⭐
   - Common in 99% of applications
   - Currently falls back to reflection
   - Could encode as varint (Unix nanoseconds) or RFC3339
   ```go
   type Event struct {
       Timestamp time.Time `beve:"ts"`
       Data      string    `beve:"data"`
   }
   ```

3. **RawMessage Support** ⭐⭐⭐⭐
   - Already exists (`rawmessage.go`) but unused
   - Delayed decoding for performance
   ```go
   type Container struct {
       Type    string           `beve:"type"`
       Payload beve.RawMessage  `beve:"payload"`
   }
   ```

4. **Decode Options** ⭐⭐⭐⭐
   - `DisallowUnknownFields` - Strict validation
   - `MaxDepth` - Prevent deep recursion attacks
   - `UseNumber` - Preserve numeric precision
   ```go
   dec := beve.NewDecoder(data)
   dec.DisallowUnknownFields()
   err := dec.Decode(&v)
   ```

### **Tier 2: Önemli (Medium Priority)**

5. **Deterministic Encoding** ⭐⭐⭐
   - Map key sorting for cryptographic signatures
   - Required for blockchain, digital signatures
   ```go
   data, _ := beve.MarshalDeterministic(v)
   // always same output for same input
   ```

6. **Binary vs String Type** ⭐⭐⭐
   - Explicit []byte type (not encoded as string)
   - Saves metadata overhead
   ```go
   type Image struct {
       Data []byte `beve:"data,binary"`
   }
   ```

7. **Extension Types** ⭐⭐⭐
   - Register custom type handlers
   - Like msgpack ext format
   ```go
   beve.RegisterType((*big.Int)(nil), 
       func(enc *Encoder, v *big.Int) error { ... })
   ```

8. **Append Mode** ⭐⭐
   - Append to existing buffer (avoid realloc)
   ```go
   buf := make([]byte, 0, 1024)
   buf = beve.MarshalAppend(buf, v1)
   buf = beve.MarshalAppend(buf, v2)
   ```

### **Tier 3: Nice to Have**

9. **Partial Decode** ⭐⭐
   - Decode only specific fields (like jq)
   ```go
   var name string
   beve.DecodeField(data, "user.name", &name)
   ```

10. **Schema Validation** ⭐
    - Optional schema checking
    - Better error messages
    ```go
    schema := beve.MustSchema(User{})
    err := beve.ValidateSchema(data, schema)
    ```

---

## 📈 Test Coverage İyileştirme Planı

### Mevcut Durum Analizi

**Coverage Gaps**:
1. Error paths (30% eksik)
2. Edge cases (nil, empty, overflow)
3. Decoder error handling
4. Custom marshaler edge cases
5. Buffer pooling stress tests

### Coverage Hedefleri

| Package | Current | Target | Gap |
|---------|---------|--------|-----|
| beve-go | 50.3% | 85% | +34.7% |
| core | 41.5% | 80% | +38.5% |

### Test Ekleme Stratejisi

1. **Encoder Edge Cases** (+15% coverage)
   - Nil pointers
   - Empty slices/maps
   - Max int/uint values
   - Circular references
   - Large nested structs

2. **Decoder Error Paths** (+20% coverage)
   - Corrupted data
   - Truncated buffers
   - Type mismatches
   - Invalid varint encoding
   - Buffer overflows

3. **Custom Marshaler Tests** (+10% coverage)
   - BinaryMarshaler/Unmarshaler
   - TextMarshaler fallback
   - Error returns

4. **Pool/Buffer Tests** (+8% coverage)
   - Buffer growth
   - Pool exhaustion
   - Large allocations (>1MB)
   - Concurrent access

5. **Integration Tests** (+5% coverage)
   - Round-trip all types
   - Comparison with JSON
   - Streaming scenarios

---

## 💡 Öneriler

### Kısa Vadede (1-2 hafta)
1. ✅ Streaming API ekle (en çok istenen özellik)
2. ✅ time.Time native support
3. ✅ Test coverage'ı 80%'e çıkar
4. ✅ RawMessage'ı aktif et

### Orta Vadede (1-2 ay)
5. ✅ Deterministic encoding
6. ✅ Decode options (DisallowUnknownFields, MaxDepth)
7. ✅ Extension type system
8. ✅ Binary type hint

### Uzun Vadede (3-6 ay)
9. ✅ Partial decode/query API
10. ✅ Schema validation
11. ✅ Performance improvements (SIMD, etc.)

---

## 🏆 Rekabet Avantajları (Korunmalı)

BEVE'nin **unique selling points** (rakiplerde yok):
- ✅ **Zero-copy mode** (MarshalZeroCopy)
- ✅ **Extreme performance** (5.6× faster)
- ✅ **Minimal allocations** (2-4 allocs)
- ✅ **Cache-optimized** structs
- ✅ **Platform-specific** optimizations
- ✅ **Go-native** design (no C dependencies)

Bu avantajlar korunmalı ve vurgulanmalı! 🎯
