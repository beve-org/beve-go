# Neoverse-N2 — Linux

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE ZeroCopy | Marshal | 73438 | 65 | 0 |
| Large Payload | BEVE | Marshal | 107836 | 188564 | 1 |
| Large Payload | CBOR | Marshal | 188093 | 188819 | 1 |
| Large Payload | MessagePack | Marshal | 288156 | 526808 | 115 |
| Large Payload | Sonic | Marshal | 326278 | 232415 | 3 |
| Large Payload | JSON | Marshal | 378972 | 205150 | 8 |
| Large Payload | BEVE | Unmarshal | 229858 | 274381 | 417 |
| Large Payload | Sonic | Unmarshal | 287776 | 375964 | 213 |
| Large Payload | MessagePack | Unmarshal | 522078 | 352495 | 6421 |
| Large Payload | CBOR | Unmarshal | 659194 | 318219 | 6467 |
| Large Payload | JSON | Unmarshal | 1992577 | 542532 | 7041 |
| Medium Payload | BEVE ZeroCopy | Marshal | 7088 | 6 | 0 |
| Medium Payload | BEVE | Marshal | 9528 | 16387 | 1 |
| Medium Payload | CBOR | Marshal | 16569 | 16397 | 1 |
| Medium Payload | Sonic | Marshal | 30198 | 22322 | 3 |
| Medium Payload | MessagePack | Marshal | 32874 | 65782 | 22 |
| Medium Payload | JSON | Marshal | 37998 | 20711 | 8 |
| Medium Payload | BEVE | Unmarshal | 22898 | 29214 | 59 |
| Medium Payload | Sonic | Unmarshal | 32432 | 45267 | 33 |
| Medium Payload | MessagePack | Unmarshal | 49407 | 34015 | 624 |
| Medium Payload | CBOR | Unmarshal | 57139 | 26808 | 549 |
| Medium Payload | JSON | Unmarshal | 169454 | 45480 | 597 |
| Small Struct | BEVE ZeroCopy | Marshal | 262 | 0 | 0 |
| Small Struct | BEVE | Marshal | 1380 | 2688 | 1 |
| Small Struct | MessagePack | Marshal | 1427 | 2056 | 7 |
| Small Struct | Sonic | Marshal | 1484 | 933 | 2 |
| Small Struct | JSON | Marshal | 2400 | 1408 | 1 |
| Small Struct | CBOR | Marshal | 2438 | 3072 | 1 |
| Small Struct | Sonic | Unmarshal | 1145 | 1122 | 6 |
| Small Struct | BEVE | Unmarshal | 1153 | 1464 | 4 |
| Small Struct | MessagePack | Unmarshal | 1874 | 1024 | 24 |
| Small Struct | CBOR | Unmarshal | 2385 | 1064 | 25 |
| Small Struct | JSON | Unmarshal | 21772 | 7464 | 99 |
