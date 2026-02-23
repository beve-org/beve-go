# Neoverse-N2 — Linux

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE ZeroCopy | Marshal | 66850 | 65 | 0 |
| Large Payload | BEVE | Marshal | 104109 | 180435 | 1 |
| Large Payload | CBOR | Marshal | 195656 | 205109 | 2 |
| Large Payload | MessagePack | Marshal | 274804 | 526801 | 115 |
| Large Payload | Sonic | Marshal | 298256 | 206755 | 3 |
| Large Payload | JSON | Marshal | 369334 | 196902 | 8 |
| Large Payload | BEVE | Unmarshal | 236239 | 281327 | 419 |
| Large Payload | Sonic | Unmarshal | 294912 | 396845 | 209 |
| Large Payload | MessagePack | Unmarshal | 545481 | 363936 | 6664 |
| Large Payload | CBOR | Unmarshal | 674162 | 322666 | 6581 |
| Large Payload | JSON | Unmarshal | 2116421 | 576869 | 7562 |
| Medium Payload | BEVE ZeroCopy | Marshal | 6861 | 5 | 0 |
| Medium Payload | BEVE | Marshal | 10148 | 18442 | 1 |
| Medium Payload | CBOR | Marshal | 17905 | 18447 | 1 |
| Medium Payload | Sonic | Marshal | 25271 | 18866 | 3 |
| Medium Payload | MessagePack | Marshal | 32071 | 65782 | 22 |
| Medium Payload | JSON | Marshal | 37822 | 20714 | 8 |
| Medium Payload | BEVE | Unmarshal | 21942 | 25629 | 59 |
| Medium Payload | Sonic | Unmarshal | 32994 | 44885 | 33 |
| Medium Payload | MessagePack | Unmarshal | 52009 | 34383 | 635 |
| Medium Payload | CBOR | Unmarshal | 60252 | 28440 | 581 |
| Medium Payload | JSON | Unmarshal | 219037 | 59352 | 817 |
| Small Struct | BEVE ZeroCopy | Marshal | 496 | 0 | 0 |
| Small Struct | BEVE | Marshal | 973 | 1792 | 1 |
| Small Struct | MessagePack | Marshal | 1492 | 2056 | 7 |
| Small Struct | CBOR | Marshal | 2214 | 2688 | 1 |
| Small Struct | Sonic | Marshal | 3919 | 3145 | 2 |
| Small Struct | JSON | Marshal | 4667 | 2688 | 1 |
| Small Struct | Sonic | Unmarshal | 1311 | 1346 | 6 |
| Small Struct | BEVE | Unmarshal | 1555 | 2616 | 4 |
| Small Struct | JSON | Unmarshal | 3038 | 552 | 15 |
| Small Struct | CBOR | Unmarshal | 3178 | 1576 | 36 |
| Small Struct | MessagePack | Unmarshal | 3440 | 2528 | 55 |
