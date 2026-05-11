# Apple M1 (Virtual) — Darwin

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE ZeroCopy | Marshal | 57528 | 26 | 0 |
| Large Payload | BEVE | Marshal | 99118 | 196662 | 1 |
| Large Payload | CBOR | Marshal | 154658 | 196742 | 1 |
| Large Payload | MessagePack | Marshal | 337496 | 526755 | 115 |
| Large Payload | JSON | Marshal | 394347 | 221525 | 8 |
| Large Payload | Sonic | Marshal | 498205 | 214365 | 3 |
| Large Payload | BEVE | Unmarshal | 228141 | 280947 | 419 |
| Large Payload | Sonic | Unmarshal | 343399 | 347647 | 209 |
| Large Payload | MessagePack | Unmarshal | 408563 | 331123 | 5999 |
| Large Payload | CBOR | Unmarshal | 602194 | 320667 | 6520 |
| Large Payload | JSON | Unmarshal | 2036351 | 549067 | 7187 |
| Medium Payload | BEVE ZeroCopy | Marshal | 5975 | 0 | 0 |
| Medium Payload | BEVE | Marshal | 10773 | 19074 | 1 |
| Medium Payload | CBOR | Marshal | 18705 | 20497 | 1 |
| Medium Payload | MessagePack | Marshal | 21885 | 33005 | 21 |
| Medium Payload | JSON | Marshal | 28463 | 18662 | 8 |
| Medium Payload | Sonic | Marshal | 40376 | 20744 | 3 |
| Medium Payload | BEVE | Unmarshal | 20711 | 27420 | 59 |
| Medium Payload | Sonic | Unmarshal | 32591 | 38567 | 33 |
| Medium Payload | MessagePack | Unmarshal | 44659 | 35949 | 664 |
| Medium Payload | CBOR | Unmarshal | 61999 | 30375 | 624 |
| Medium Payload | JSON | Unmarshal | 168223 | 50776 | 659 |
| Small Struct | BEVE ZeroCopy | Marshal | 497 | 0 | 0 |
| Small Struct | BEVE | Marshal | 787 | 1408 | 1 |
| Small Struct | CBOR | Marshal | 1604 | 2048 | 1 |
| Small Struct | JSON | Marshal | 2183 | 1280 | 1 |
| Small Struct | Sonic | Marshal | 3016 | 1210 | 2 |
| Small Struct | MessagePack | Marshal | 4285 | 8201 | 9 |
| Small Struct | CBOR | Unmarshal | 837 | 232 | 7 |
| Small Struct | BEVE | Unmarshal | 886 | 1336 | 4 |
| Small Struct | MessagePack | Unmarshal | 1966 | 1312 | 30 |
| Small Struct | Sonic | Unmarshal | 2538 | 3674 | 6 |
| Small Struct | JSON | Unmarshal | 18637 | 7504 | 100 |
