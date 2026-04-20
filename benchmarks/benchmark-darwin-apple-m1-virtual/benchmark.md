# Apple M1 (Virtual) — Darwin

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE ZeroCopy | Marshal | 50390 | 26 | 0 |
| Large Payload | BEVE | Marshal | 76259 | 188458 | 1 |
| Large Payload | CBOR | Marshal | 132358 | 188577 | 1 |
| Large Payload | MessagePack | Marshal | 187344 | 526755 | 115 |
| Large Payload | JSON | Marshal | 287060 | 196893 | 8 |
| Large Payload | Sonic | Marshal | 380254 | 221529 | 3 |
| Large Payload | BEVE | Unmarshal | 165896 | 254804 | 419 |
| Large Payload | Sonic | Unmarshal | 258660 | 354314 | 213 |
| Large Payload | MessagePack | Unmarshal | 349030 | 329792 | 5981 |
| Large Payload | CBOR | Unmarshal | 430306 | 289402 | 5903 |
| Large Payload | JSON | Unmarshal | 1652453 | 564578 | 7383 |
| Medium Payload | BEVE ZeroCopy | Marshal | 5837 | 1 | 0 |
| Medium Payload | BEVE | Marshal | 7602 | 19075 | 1 |
| Medium Payload | CBOR | Marshal | 14134 | 20494 | 1 |
| Medium Payload | MessagePack | Marshal | 20877 | 65778 | 22 |
| Medium Payload | JSON | Marshal | 28953 | 21994 | 8 |
| Medium Payload | Sonic | Marshal | 41477 | 24876 | 3 |
| Medium Payload | BEVE | Unmarshal | 15765 | 31005 | 59 |
| Medium Payload | Sonic | Unmarshal | 29545 | 43869 | 33 |
| Medium Payload | MessagePack | Unmarshal | 33562 | 35037 | 646 |
| Medium Payload | CBOR | Unmarshal | 52264 | 36360 | 747 |
| Medium Payload | JSON | Unmarshal | 167889 | 58512 | 780 |
| Small Struct | BEVE | Marshal | 355 | 704 | 1 |
| Small Struct | BEVE ZeroCopy | Marshal | 496 | 0 | 0 |
| Small Struct | MessagePack | Marshal | 917 | 2056 | 7 |
| Small Struct | CBOR | Marshal | 1284 | 2048 | 1 |
| Small Struct | JSON | Marshal | 1463 | 1152 | 1 |
| Small Struct | Sonic | Marshal | 5148 | 3102 | 2 |
| Small Struct | MessagePack | Unmarshal | 648 | 256 | 7 |
| Small Struct | BEVE | Unmarshal | 999 | 2616 | 4 |
| Small Struct | CBOR | Unmarshal | 1669 | 1064 | 25 |
| Small Struct | Sonic | Unmarshal | 1868 | 2778 | 6 |
| Small Struct | JSON | Unmarshal | 8610 | 3752 | 54 |
