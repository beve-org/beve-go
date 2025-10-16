# Neoverse-N2 — Linux

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE ZeroCopy | Marshal | 67055 | 259 | 2 |
| Large Payload | BEVE | Marshal | 114482 | 197813 | 3 |
| Large Payload | CBOR | Marshal | 195884 | 191039 | 2 |
| Large Payload | MessagePack | Marshal | 277572 | 526859 | 115 |
| Large Payload | Sonic | Marshal | 306272 | 227433 | 4 |
| Large Payload | JSON | Marshal | 355900 | 198306 | 9 |
| Large Payload | BEVE | Unmarshal | 233443 | 295376 | 417 |
| Large Payload | Sonic | Unmarshal | 322058 | 408490 | 211 |
| Large Payload | MessagePack | Unmarshal | 531069 | 350072 | 6369 |
| Large Payload | CBOR | Unmarshal | 624183 | 290026 | 5928 |
| Large Payload | JSON | Unmarshal | 1911908 | 512219 | 6704 |
| Medium Payload | BEVE ZeroCopy | Marshal | 8021 | 141 | 2 |
| Medium Payload | BEVE | Marshal | 9676 | 16518 | 3 |
| Medium Payload | CBOR | Marshal | 18978 | 20610 | 2 |
| Medium Payload | Sonic | Marshal | 29982 | 22325 | 4 |
| Medium Payload | MessagePack | Marshal | 32008 | 65837 | 22 |
| Medium Payload | JSON | Marshal | 40947 | 24908 | 9 |
| Medium Payload | BEVE | Unmarshal | 21510 | 25885 | 58 |
| Medium Payload | Sonic | Unmarshal | 27921 | 37845 | 33 |
| Medium Payload | MessagePack | Unmarshal | 55153 | 39568 | 738 |
| Medium Payload | CBOR | Unmarshal | 55295 | 25528 | 525 |
| Medium Payload | JSON | Unmarshal | 216009 | 60952 | 806 |
| Small Struct | BEVE ZeroCopy | Marshal | 719 | 290 | 2 |
| Small Struct | Sonic | Marshal | 955 | 608 | 3 |
| Small Struct | CBOR | Marshal | 973 | 848 | 2 |
| Small Struct | MessagePack | Marshal | 1477 | 2176 | 7 |
| Small Struct | BEVE | Marshal | 1552 | 2980 | 3 |
| Small Struct | JSON | Marshal | 4772 | 3219 | 2 |
| Small Struct | BEVE | Unmarshal | 1052 | 1336 | 4 |
| Small Struct | CBOR | Unmarshal | 1311 | 328 | 10 |
| Small Struct | Sonic | Unmarshal | 2927 | 4595 | 6 |
| Small Struct | MessagePack | Unmarshal | 3547 | 2688 | 56 |
| Small Struct | JSON | Unmarshal | 4484 | 904 | 22 |
