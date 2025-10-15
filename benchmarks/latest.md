# Apple M2 Max — Darwin

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE ZeroCopy | Marshal | 53252 | 5431 | 2 |
| Large Payload | BEVE | Marshal | 69970 | 191465 | 3 |
| Large Payload | CBOR | Marshal | 125702 | 218852 | 3 |
| Large Payload | MessagePack | Marshal | 159745 | 527190 | 115 |
| Large Payload | JSON | Marshal | 292380 | 226899 | 9 |
| Large Payload | Sonic | Marshal | 323524 | 223570 | 4 |
| Large Payload | BEVE | Unmarshal | 137960 | 282139 | 419 |
| Large Payload | Sonic | Unmarshal | 236778 | 352119 | 209 |
| Large Payload | MessagePack | Unmarshal | 331702 | 378887 | 6960 |
| Large Payload | CBOR | Unmarshal | 419377 | 316713 | 6462 |
| Large Payload | JSON | Unmarshal | 1402622 | 504716 | 6690 |
| Medium Payload | BEVE ZeroCopy | Marshal | 6201 | 843 | 2 |
| Medium Payload | BEVE | Marshal | 7604 | 18937 | 3 |
| Medium Payload | CBOR | Marshal | 11823 | 20598 | 2 |
| Medium Payload | MessagePack | Marshal | 18882 | 65931 | 22 |
| Medium Payload | Sonic | Marshal | 27559 | 20816 | 4 |
| Medium Payload | JSON | Marshal | 32763 | 24882 | 9 |
| Medium Payload | BEVE | Unmarshal | 14591 | 34461 | 59 |
| Medium Payload | Sonic | Unmarshal | 23365 | 55467 | 33 |
| Medium Payload | MessagePack | Unmarshal | 30924 | 37902 | 705 |
| Medium Payload | CBOR | Unmarshal | 42328 | 34664 | 715 |
| Medium Payload | JSON | Unmarshal | 107895 | 40456 | 538 |
| Small Struct | BEVE ZeroCopy | Marshal | 313 | 324 | 2 |
| Small Struct | BEVE | Marshal | 1030 | 1894 | 3 |
| Small Struct | Sonic | Marshal | 1199 | 1950 | 3 |
| Small Struct | CBOR | Marshal | 1406 | 2210 | 2 |
| Small Struct | JSON | Marshal | 1420 | 887 | 2 |
| Small Struct | MessagePack | Marshal | 3271 | 8339 | 9 |
| Small Struct | BEVE | Unmarshal | 757 | 1482 | 4 |
| Small Struct | MessagePack | Unmarshal | 1427 | 756 | 18 |
| Small Struct | CBOR | Unmarshal | 1600 | 856 | 21 |
| Small Struct | Sonic | Unmarshal | 2431 | 3737 | 6 |
| Small Struct | JSON | Unmarshal | 10475 | 3848 | 57 |
