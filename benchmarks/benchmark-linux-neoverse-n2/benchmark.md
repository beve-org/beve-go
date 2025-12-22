# Neoverse-N2 — Linux

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE ZeroCopy | Marshal | 67323 | 65 | 0 |
| Large Payload | BEVE | Marshal | 104299 | 188577 | 1 |
| Large Payload | CBOR | Marshal | 187948 | 196909 | 1 |
| Large Payload | MessagePack | Marshal | 273620 | 526804 | 115 |
| Large Payload | Sonic | Marshal | 298672 | 216591 | 3 |
| Large Payload | JSON | Marshal | 382296 | 213447 | 8 |
| Large Payload | BEVE | Unmarshal | 218588 | 280751 | 417 |
| Large Payload | Sonic | Unmarshal | 281625 | 399977 | 209 |
| Large Payload | MessagePack | Unmarshal | 498889 | 341850 | 6211 |
| Large Payload | CBOR | Unmarshal | 672288 | 334377 | 6820 |
| Large Payload | JSON | Unmarshal | 1796521 | 481154 | 6307 |
| Medium Payload | BEVE ZeroCopy | Marshal | 6135 | 6 | 0 |
| Medium Payload | BEVE | Marshal | 10258 | 20489 | 1 |
| Medium Payload | CBOR | Marshal | 17062 | 18450 | 1 |
| Medium Payload | Sonic | Marshal | 22864 | 16927 | 3 |
| Medium Payload | MessagePack | Marshal | 29989 | 65782 | 22 |
| Medium Payload | JSON | Marshal | 39400 | 21994 | 8 |
| Medium Payload | BEVE | Unmarshal | 20523 | 25245 | 59 |
| Medium Payload | Sonic | Unmarshal | 34796 | 50176 | 33 |
| Medium Payload | MessagePack | Unmarshal | 57280 | 41568 | 778 |
| Medium Payload | CBOR | Unmarshal | 61113 | 28888 | 596 |
| Medium Payload | JSON | Unmarshal | 250041 | 72856 | 965 |
| Small Struct | BEVE ZeroCopy | Marshal | 630 | 0 | 0 |
| Small Struct | CBOR | Marshal | 870 | 768 | 1 |
| Small Struct | BEVE | Marshal | 1039 | 2048 | 1 |
| Small Struct | MessagePack | Marshal | 1487 | 2056 | 7 |
| Small Struct | Sonic | Marshal | 2735 | 2120 | 2 |
| Small Struct | JSON | Marshal | 4606 | 3072 | 1 |
| Small Struct | BEVE | Unmarshal | 1103 | 1464 | 4 |
| Small Struct | Sonic | Unmarshal | 1429 | 1675 | 6 |
| Small Struct | CBOR | Unmarshal | 3981 | 2216 | 48 |
| Small Struct | MessagePack | Unmarshal | 5655 | 5209 | 107 |
| Small Struct | JSON | Unmarshal | 21101 | 7400 | 97 |
