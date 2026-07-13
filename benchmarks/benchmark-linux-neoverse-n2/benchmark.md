# Neoverse-N2 — Linux

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE ZeroCopy | Marshal | 67796 | 65 | 0 |
| Large Payload | BEVE | Marshal | 106644 | 180369 | 1 |
| Large Payload | CBOR | Marshal | 192156 | 196884 | 1 |
| Large Payload | MessagePack | Marshal | 296622 | 526809 | 115 |
| Large Payload | Sonic | Marshal | 305224 | 214769 | 3 |
| Large Payload | JSON | Marshal | 377912 | 205177 | 8 |
| Large Payload | BEVE | Unmarshal | 238068 | 284719 | 419 |
| Large Payload | Sonic | Unmarshal | 290499 | 388068 | 213 |
| Large Payload | MessagePack | Unmarshal | 520213 | 339209 | 6165 |
| Large Payload | CBOR | Unmarshal | 698038 | 338170 | 6886 |
| Large Payload | JSON | Unmarshal | 2072014 | 561116 | 7360 |
| Medium Payload | BEVE ZeroCopy | Marshal | 7022 | 0 | 0 |
| Medium Payload | BEVE | Marshal | 11343 | 21767 | 1 |
| Medium Payload | CBOR | Marshal | 21812 | 24594 | 1 |
| Medium Payload | MessagePack | Marshal | 32990 | 65782 | 22 |
| Medium Payload | Sonic | Marshal | 33990 | 25040 | 3 |
| Medium Payload | JSON | Marshal | 38768 | 21991 | 8 |
| Medium Payload | BEVE | Unmarshal | 25623 | 34975 | 59 |
| Medium Payload | Sonic | Unmarshal | 33339 | 49197 | 33 |
| Medium Payload | MessagePack | Unmarshal | 51887 | 35999 | 665 |
| Medium Payload | CBOR | Unmarshal | 68090 | 33720 | 694 |
| Medium Payload | JSON | Unmarshal | 205748 | 57432 | 753 |
| Small Struct | BEVE | Marshal | 455 | 512 | 1 |
| Small Struct | BEVE ZeroCopy | Marshal | 505 | 0 | 0 |
| Small Struct | Sonic | Marshal | 768 | 421 | 2 |
| Small Struct | CBOR | Marshal | 2300 | 2688 | 1 |
| Small Struct | MessagePack | Marshal | 3723 | 8201 | 9 |
| Small Struct | JSON | Marshal | 4270 | 2689 | 1 |
| Small Struct | BEVE | Unmarshal | 786 | 568 | 4 |
| Small Struct | Sonic | Unmarshal | 2030 | 2854 | 6 |
| Small Struct | MessagePack | Unmarshal | 2066 | 1216 | 28 |
| Small Struct | CBOR | Unmarshal | 4175 | 2272 | 50 |
| Small Struct | JSON | Unmarshal | 23512 | 7720 | 107 |
