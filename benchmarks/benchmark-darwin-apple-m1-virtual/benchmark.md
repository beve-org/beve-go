# Apple M1 (Virtual) — Darwin

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE ZeroCopy | Marshal | 58161 | 39 | 0 |
| Large Payload | BEVE | Marshal | 104579 | 188496 | 1 |
| Large Payload | CBOR | Marshal | 156987 | 196743 | 1 |
| Large Payload | MessagePack | Marshal | 249507 | 526751 | 115 |
| Large Payload | JSON | Marshal | 370185 | 213306 | 8 |
| Large Payload | Sonic | Marshal | 534028 | 222382 | 3 |
| Large Payload | BEVE | Unmarshal | 176818 | 280051 | 418 |
| Large Payload | Sonic | Unmarshal | 312026 | 337034 | 209 |
| Large Payload | MessagePack | Unmarshal | 484216 | 345810 | 6283 |
| Large Payload | CBOR | Unmarshal | 595654 | 303098 | 6176 |
| Large Payload | JSON | Unmarshal | 1790340 | 493937 | 6470 |
| Medium Payload | BEVE ZeroCopy | Marshal | 6131 | 1 | 0 |
| Medium Payload | BEVE | Marshal | 8811 | 20483 | 1 |
| Medium Payload | CBOR | Marshal | 24033 | 24592 | 1 |
| Medium Payload | MessagePack | Marshal | 26159 | 65779 | 22 |
| Medium Payload | JSON | Marshal | 32152 | 21994 | 8 |
| Medium Payload | Sonic | Marshal | 37210 | 20671 | 3 |
| Medium Payload | BEVE | Unmarshal | 16713 | 28765 | 59 |
| Medium Payload | Sonic | Unmarshal | 36784 | 41579 | 33 |
| Medium Payload | MessagePack | Unmarshal | 39129 | 33821 | 622 |
| Medium Payload | CBOR | Unmarshal | 63823 | 39680 | 814 |
| Medium Payload | JSON | Unmarshal | 192865 | 62664 | 814 |
| Small Struct | BEVE ZeroCopy | Marshal | 273 | 0 | 0 |
| Small Struct | CBOR | Marshal | 444 | 448 | 1 |
| Small Struct | BEVE | Marshal | 830 | 2304 | 1 |
| Small Struct | MessagePack | Marshal | 968 | 2056 | 7 |
| Small Struct | JSON | Marshal | 1029 | 704 | 1 |
| Small Struct | Sonic | Marshal | 3633 | 2071 | 2 |
| Small Struct | BEVE | Unmarshal | 571 | 696 | 4 |
| Small Struct | Sonic | Unmarshal | 3046 | 5304 | 6 |
| Small Struct | CBOR | Unmarshal | 3317 | 2504 | 55 |
| Small Struct | MessagePack | Unmarshal | 3886 | 1928 | 42 |
| Small Struct | JSON | Unmarshal | 17580 | 4840 | 88 |
