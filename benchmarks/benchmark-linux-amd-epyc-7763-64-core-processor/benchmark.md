# AMD EPYC 7763 64-Core Processor — Linux

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE ZeroCopy | Marshal | 76011 | 52 | 0 |
| Large Payload | BEVE | Marshal | 115931 | 196667 | 1 |
| Large Payload | Sonic | Marshal | 151914 | 207843 | 3 |
| Large Payload | CBOR | Marshal | 197863 | 188641 | 1 |
| Large Payload | MessagePack | Marshal | 308334 | 526778 | 115 |
| Large Payload | JSON | Marshal | 454126 | 221503 | 8 |
| Large Payload | BEVE | Unmarshal | 228403 | 264447 | 418 |
| Large Payload | Sonic | Unmarshal | 388372 | 585586 | 593 |
| Large Payload | MessagePack | Unmarshal | 546035 | 329474 | 5953 |
| Large Payload | CBOR | Unmarshal | 754078 | 301161 | 6133 |
| Large Payload | JSON | Unmarshal | 2420195 | 570427 | 7462 |
| Medium Payload | BEVE ZeroCopy | Marshal | 7966 | 6 | 0 |
| Medium Payload | BEVE | Marshal | 15410 | 27273 | 1 |
| Medium Payload | Sonic | Marshal | 17937 | 25145 | 3 |
| Medium Payload | CBOR | Marshal | 24829 | 24601 | 1 |
| Medium Payload | MessagePack | Marshal | 36501 | 65783 | 22 |
| Medium Payload | JSON | Marshal | 42018 | 20715 | 8 |
| Medium Payload | BEVE | Unmarshal | 20453 | 21085 | 58 |
| Medium Payload | Sonic | Unmarshal | 40751 | 60815 | 77 |
| Medium Payload | MessagePack | Unmarshal | 64148 | 42273 | 793 |
| Medium Payload | CBOR | Unmarshal | 75104 | 31096 | 639 |
| Medium Payload | JSON | Unmarshal | 229238 | 58048 | 741 |
| Small Struct | BEVE ZeroCopy | Marshal | 324 | 0 | 0 |
| Small Struct | BEVE | Marshal | 1007 | 1537 | 1 |
| Small Struct | CBOR | Marshal | 1749 | 1792 | 1 |
| Small Struct | Sonic | Marshal | 1969 | 3174 | 2 |
| Small Struct | MessagePack | Marshal | 2639 | 4104 | 8 |
| Small Struct | JSON | Marshal | 3965 | 2048 | 1 |
| Small Struct | BEVE | Unmarshal | 1291 | 1720 | 4 |
| Small Struct | MessagePack | Unmarshal | 3943 | 2816 | 60 |
| Small Struct | Sonic | Unmarshal | 4331 | 7740 | 10 |
| Small Struct | CBOR | Unmarshal | 6555 | 3208 | 69 |
| Small Struct | JSON | Unmarshal | 23814 | 7464 | 99 |
