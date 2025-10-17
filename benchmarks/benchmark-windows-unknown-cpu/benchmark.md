# Unknown CPU — Windows

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE ZeroCopy | Marshal | 78288 | 52 | 0 |
| Large Payload | BEVE | Marshal | 113305 | 196642 | 1 |
| Large Payload | Sonic | Marshal | 155984 | 206763 | 3 |
| Large Payload | CBOR | Marshal | 213966 | 188543 | 1 |
| Large Payload | MessagePack | Marshal | 275794 | 526706 | 115 |
| Large Payload | JSON | Marshal | 467944 | 213265 | 8 |
| Large Payload | BEVE | Unmarshal | 272015 | 273797 | 417 |
| Large Payload | Sonic | Unmarshal | 404565 | 529590 | 570 |
| Large Payload | MessagePack | Unmarshal | 679461 | 350647 | 6398 |
| Large Payload | CBOR | Unmarshal | 822770 | 311225 | 6350 |
| Large Payload | JSON | Unmarshal | 2542790 | 505763 | 6689 |
| Medium Payload | BEVE ZeroCopy | Marshal | 7910 | 5 | 0 |
| Medium Payload | BEVE | Marshal | 12430 | 18437 | 1 |
| Medium Payload | Sonic | Marshal | 13364 | 16625 | 3 |
| Medium Payload | CBOR | Marshal | 22300 | 20491 | 1 |
| Medium Payload | MessagePack | Marshal | 25228 | 33001 | 21 |
| Medium Payload | JSON | Marshal | 48475 | 22000 | 8 |
| Medium Payload | BEVE | Unmarshal | 29216 | 30843 | 59 |
| Medium Payload | Sonic | Unmarshal | 52062 | 66684 | 80 |
| Medium Payload | MessagePack | Unmarshal | 72539 | 38253 | 710 |
| Medium Payload | CBOR | Unmarshal | 93435 | 37536 | 767 |
| Medium Payload | JSON | Unmarshal | 243516 | 49400 | 652 |
| Small Struct | BEVE ZeroCopy | Marshal | 592 | 0 | 0 |
| Small Struct | BEVE | Marshal | 903 | 1024 | 1 |
| Small Struct | Sonic | Marshal | 1027 | 1203 | 2 |
| Small Struct | CBOR | Marshal | 2419 | 2688 | 1 |
| Small Struct | MessagePack | Marshal | 2699 | 4104 | 8 |
| Small Struct | JSON | Marshal | 4034 | 2048 | 1 |
| Small Struct | BEVE | Unmarshal | 2490 | 3512 | 4 |
| Small Struct | MessagePack | Unmarshal | 3102 | 1664 | 37 |
| Small Struct | Sonic | Unmarshal | 5498 | 7419 | 10 |
| Small Struct | CBOR | Unmarshal | 9428 | 4584 | 96 |
| Small Struct | JSON | Unmarshal | 16854 | 4048 | 63 |
