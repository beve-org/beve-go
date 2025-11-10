# Neoverse-N2 — Linux

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE ZeroCopy | Marshal | 70467 | 65 | 0 |
| Large Payload | BEVE | Marshal | 120964 | 213169 | 1 |
| Large Payload | CBOR | Marshal | 189828 | 197014 | 1 |
| Large Payload | MessagePack | Marshal | 299817 | 526810 | 115 |
| Large Payload | Sonic | Marshal | 304055 | 217351 | 3 |
| Large Payload | JSON | Marshal | 396279 | 213421 | 8 |
| Large Payload | BEVE | Unmarshal | 229116 | 266472 | 419 |
| Large Payload | Sonic | Unmarshal | 295245 | 399074 | 213 |
| Large Payload | MessagePack | Unmarshal | 518680 | 342490 | 6213 |
| Large Payload | CBOR | Unmarshal | 666764 | 319035 | 6499 |
| Large Payload | JSON | Unmarshal | 2029889 | 544900 | 7186 |
| Medium Payload | BEVE ZeroCopy | Marshal | 6788 | 5 | 0 |
| Medium Payload | BEVE | Marshal | 10892 | 20492 | 1 |
| Medium Payload | CBOR | Marshal | 18006 | 18453 | 1 |
| Medium Payload | MessagePack | Marshal | 23433 | 33007 | 21 |
| Medium Payload | Sonic | Marshal | 32977 | 25041 | 3 |
| Medium Payload | JSON | Marshal | 36899 | 20714 | 8 |
| Medium Payload | BEVE | Unmarshal | 21973 | 26013 | 59 |
| Medium Payload | Sonic | Unmarshal | 27733 | 34086 | 33 |
| Medium Payload | MessagePack | Unmarshal | 52146 | 34623 | 642 |
| Medium Payload | CBOR | Unmarshal | 66935 | 32728 | 674 |
| Medium Payload | JSON | Unmarshal | 187899 | 49512 | 684 |
| Small Struct | BEVE ZeroCopy | Marshal | 645 | 0 | 0 |
| Small Struct | JSON | Marshal | 829 | 352 | 1 |
| Small Struct | BEVE | Marshal | 920 | 1536 | 1 |
| Small Struct | CBOR | Marshal | 1061 | 1024 | 1 |
| Small Struct | MessagePack | Marshal | 1660 | 2056 | 7 |
| Small Struct | Sonic | Marshal | 3387 | 2754 | 2 |
| Small Struct | BEVE | Unmarshal | 1325 | 1848 | 4 |
| Small Struct | Sonic | Unmarshal | 1966 | 2662 | 6 |
| Small Struct | MessagePack | Unmarshal | 4676 | 3880 | 81 |
| Small Struct | CBOR | Unmarshal | 7693 | 4672 | 99 |
| Small Struct | JSON | Unmarshal | 16030 | 4392 | 74 |
