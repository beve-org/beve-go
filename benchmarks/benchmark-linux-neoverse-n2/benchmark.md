# Neoverse-N2 — Linux

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE ZeroCopy | Marshal | 69953 | 65 | 0 |
| Large Payload | BEVE | Marshal | 109862 | 196797 | 1 |
| Large Payload | CBOR | Marshal | 197726 | 205133 | 1 |
| Large Payload | MessagePack | Marshal | 281053 | 526806 | 115 |
| Large Payload | Sonic | Marshal | 317141 | 225718 | 3 |
| Large Payload | JSON | Marshal | 405133 | 221693 | 8 |
| Large Payload | BEVE | Unmarshal | 225873 | 269034 | 417 |
| Large Payload | Sonic | Unmarshal | 276570 | 382490 | 213 |
| Large Payload | MessagePack | Unmarshal | 513543 | 344025 | 6246 |
| Large Payload | CBOR | Unmarshal | 635404 | 306522 | 6247 |
| Large Payload | JSON | Unmarshal | 1985389 | 537243 | 7068 |
| Medium Payload | BEVE ZeroCopy | Marshal | 6606 | 1 | 0 |
| Medium Payload | BEVE | Marshal | 12387 | 27277 | 1 |
| Medium Payload | CBOR | Marshal | 20132 | 21777 | 1 |
| Medium Payload | Sonic | Marshal | 27456 | 19418 | 3 |
| Medium Payload | MessagePack | Marshal | 30375 | 65782 | 22 |
| Medium Payload | JSON | Marshal | 45268 | 27496 | 8 |
| Medium Payload | BEVE | Unmarshal | 21548 | 26845 | 59 |
| Medium Payload | Sonic | Unmarshal | 35249 | 51725 | 33 |
| Medium Payload | MessagePack | Unmarshal | 55412 | 39200 | 730 |
| Medium Payload | CBOR | Unmarshal | 58759 | 27560 | 569 |
| Medium Payload | JSON | Unmarshal | 172686 | 45472 | 618 |
| Small Struct | BEVE | Marshal | 390 | 416 | 1 |
| Small Struct | BEVE ZeroCopy | Marshal | 631 | 0 | 0 |
| Small Struct | MessagePack | Marshal | 736 | 520 | 5 |
| Small Struct | CBOR | Marshal | 2011 | 2304 | 1 |
| Small Struct | Sonic | Marshal | 2036 | 1452 | 2 |
| Small Struct | JSON | Marshal | 4133 | 2688 | 1 |
| Small Struct | BEVE | Unmarshal | 966 | 1208 | 4 |
| Small Struct | Sonic | Unmarshal | 1239 | 1218 | 6 |
| Small Struct | MessagePack | Unmarshal | 2525 | 1760 | 39 |
| Small Struct | CBOR | Unmarshal | 7062 | 4360 | 93 |
| Small Struct | JSON | Unmarshal | 19042 | 4872 | 89 |
