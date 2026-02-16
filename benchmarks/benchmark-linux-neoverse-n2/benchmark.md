# Neoverse-N2 — Linux

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE ZeroCopy | Marshal | 68534 | 105 | 0 |
| Large Payload | BEVE | Marshal | 106451 | 196729 | 1 |
| Large Payload | CBOR | Marshal | 178528 | 188714 | 1 |
| Large Payload | MessagePack | Marshal | 273950 | 526804 | 115 |
| Large Payload | Sonic | Marshal | 290300 | 207318 | 3 |
| Large Payload | JSON | Marshal | 380764 | 213316 | 8 |
| Large Payload | BEVE | Unmarshal | 220959 | 263686 | 417 |
| Large Payload | Sonic | Unmarshal | 274145 | 369916 | 211 |
| Large Payload | MessagePack | Unmarshal | 519362 | 352703 | 6435 |
| Large Payload | CBOR | Unmarshal | 653818 | 317163 | 6468 |
| Large Payload | JSON | Unmarshal | 1992433 | 533979 | 7078 |
| Medium Payload | BEVE ZeroCopy | Marshal | 7300 | 6 | 0 |
| Medium Payload | BEVE | Marshal | 9543 | 16390 | 1 |
| Medium Payload | CBOR | Marshal | 21052 | 24594 | 1 |
| Medium Payload | MessagePack | Marshal | 31732 | 65782 | 22 |
| Medium Payload | Sonic | Marshal | 34081 | 25410 | 3 |
| Medium Payload | JSON | Marshal | 36316 | 19302 | 8 |
| Medium Payload | BEVE | Unmarshal | 23037 | 28958 | 59 |
| Medium Payload | Sonic | Unmarshal | 29359 | 39076 | 33 |
| Medium Payload | MessagePack | Unmarshal | 50606 | 34127 | 627 |
| Medium Payload | CBOR | Unmarshal | 57758 | 26680 | 549 |
| Medium Payload | JSON | Unmarshal | 167821 | 46888 | 588 |
| Small Struct | BEVE ZeroCopy | Marshal | 195 | 0 | 0 |
| Small Struct | BEVE | Marshal | 555 | 704 | 1 |
| Small Struct | JSON | Marshal | 2270 | 1280 | 1 |
| Small Struct | CBOR | Marshal | 2408 | 3072 | 1 |
| Small Struct | MessagePack | Marshal | 2613 | 4104 | 8 |
| Small Struct | Sonic | Marshal | 3900 | 3187 | 2 |
| Small Struct | BEVE | Unmarshal | 1062 | 1208 | 4 |
| Small Struct | MessagePack | Unmarshal | 1946 | 1088 | 25 |
| Small Struct | Sonic | Unmarshal | 3407 | 5619 | 6 |
| Small Struct | CBOR | Unmarshal | 3796 | 2016 | 44 |
| Small Struct | JSON | Unmarshal | 6287 | 1416 | 31 |
