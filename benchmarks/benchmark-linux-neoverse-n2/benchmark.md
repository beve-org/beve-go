# Neoverse-N2 — Linux

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE ZeroCopy | Marshal | 67863 | 65 | 0 |
| Large Payload | BEVE | Marshal | 109356 | 180502 | 1 |
| Large Payload | CBOR | Marshal | 200728 | 205212 | 1 |
| Large Payload | MessagePack | Marshal | 274875 | 526799 | 115 |
| Large Payload | Sonic | Marshal | 298142 | 201731 | 3 |
| Large Payload | JSON | Marshal | 406428 | 213421 | 8 |
| Large Payload | BEVE | Unmarshal | 232305 | 273708 | 418 |
| Large Payload | Sonic | Unmarshal | 291462 | 395901 | 213 |
| Large Payload | MessagePack | Unmarshal | 520497 | 353820 | 6455 |
| Large Payload | CBOR | Unmarshal | 665595 | 319738 | 6507 |
| Large Payload | JSON | Unmarshal | 2103633 | 565229 | 7526 |
| Medium Payload | BEVE ZeroCopy | Marshal | 6776 | 3 | 0 |
| Medium Payload | BEVE | Marshal | 9398 | 16390 | 1 |
| Medium Payload | CBOR | Marshal | 18516 | 18450 | 1 |
| Medium Payload | MessagePack | Marshal | 22520 | 33007 | 21 |
| Medium Payload | Sonic | Marshal | 28616 | 20971 | 3 |
| Medium Payload | JSON | Marshal | 42709 | 24807 | 8 |
| Medium Payload | BEVE | Unmarshal | 20790 | 24221 | 59 |
| Medium Payload | Sonic | Unmarshal | 29768 | 40248 | 33 |
| Medium Payload | MessagePack | Unmarshal | 51114 | 34207 | 633 |
| Medium Payload | CBOR | Unmarshal | 63568 | 30824 | 629 |
| Medium Payload | JSON | Unmarshal | 188909 | 52728 | 690 |
| Small Struct | BEVE ZeroCopy | Marshal | 299 | 0 | 0 |
| Small Struct | BEVE | Marshal | 1020 | 1792 | 1 |
| Small Struct | JSON | Marshal | 1460 | 704 | 1 |
| Small Struct | Sonic | Marshal | 1701 | 1075 | 2 |
| Small Struct | CBOR | Marshal | 2345 | 2688 | 1 |
| Small Struct | MessagePack | Marshal | 3755 | 8201 | 9 |
| Small Struct | BEVE | Unmarshal | 678 | 376 | 4 |
| Small Struct | Sonic | Unmarshal | 2125 | 3110 | 6 |
| Small Struct | MessagePack | Unmarshal | 5107 | 4360 | 92 |
| Small Struct | CBOR | Unmarshal | 6496 | 3880 | 82 |
| Small Struct | JSON | Unmarshal | 22363 | 7560 | 102 |
