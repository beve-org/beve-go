# Apple M1 (Virtual) — Darwin

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE ZeroCopy | Marshal | 63156 | 39 | 0 |
| Large Payload | BEVE | Marshal | 89771 | 196648 | 1 |
| Large Payload | MessagePack | Marshal | 206123 | 526753 | 115 |
| Large Payload | CBOR | Marshal | 290415 | 188534 | 1 |
| Large Payload | JSON | Marshal | 326646 | 213307 | 8 |
| Large Payload | Sonic | Marshal | 416657 | 213869 | 3 |
| Large Payload | BEVE | Unmarshal | 178961 | 280693 | 418 |
| Large Payload | Sonic | Unmarshal | 298680 | 348137 | 209 |
| Large Payload | CBOR | Unmarshal | 668896 | 318426 | 6494 |
| Large Payload | MessagePack | Unmarshal | 722654 | 356453 | 6502 |
| Large Payload | JSON | Unmarshal | 1993269 | 537531 | 7137 |
| Medium Payload | BEVE ZeroCopy | Marshal | 5592 | 3 | 0 |
| Medium Payload | BEVE | Marshal | 10281 | 20484 | 1 |
| Medium Payload | CBOR | Marshal | 17190 | 20497 | 1 |
| Medium Payload | MessagePack | Marshal | 21964 | 65778 | 22 |
| Medium Payload | Sonic | Marshal | 31656 | 16594 | 3 |
| Medium Payload | JSON | Marshal | 36691 | 21990 | 8 |
| Medium Payload | BEVE | Unmarshal | 14863 | 24348 | 59 |
| Medium Payload | Sonic | Unmarshal | 27407 | 36647 | 31 |
| Medium Payload | MessagePack | Unmarshal | 38224 | 32333 | 596 |
| Medium Payload | CBOR | Unmarshal | 45001 | 29792 | 613 |
| Medium Payload | JSON | Unmarshal | 185177 | 62312 | 814 |
| Small Struct | BEVE ZeroCopy | Marshal | 566 | 0 | 0 |
| Small Struct | BEVE | Marshal | 585 | 1280 | 1 |
| Small Struct | JSON | Marshal | 985 | 704 | 1 |
| Small Struct | CBOR | Marshal | 1496 | 1152 | 1 |
| Small Struct | MessagePack | Marshal | 1824 | 2056 | 7 |
| Small Struct | Sonic | Marshal | 4150 | 2092 | 2 |
| Small Struct | Sonic | Unmarshal | 1863 | 2644 | 6 |
| Small Struct | BEVE | Unmarshal | 1956 | 1720 | 4 |
| Small Struct | CBOR | Unmarshal | 2108 | 1096 | 26 |
| Small Struct | MessagePack | Unmarshal | 2898 | 3104 | 65 |
| Small Struct | JSON | Unmarshal | 20104 | 7400 | 97 |
