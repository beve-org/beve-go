# Neoverse-N2 — Linux

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE ZeroCopy | Marshal | 70557 | 52 | 0 |
| Large Payload | BEVE | Marshal | 108738 | 188590 | 1 |
| Large Payload | CBOR | Marshal | 184991 | 188899 | 1 |
| Large Payload | MessagePack | Marshal | 283492 | 526805 | 115 |
| Large Payload | Sonic | Marshal | 303783 | 214522 | 3 |
| Large Payload | JSON | Marshal | 401136 | 221615 | 8 |
| Large Payload | BEVE | Unmarshal | 232258 | 271565 | 419 |
| Large Payload | Sonic | Unmarshal | 296765 | 400972 | 213 |
| Large Payload | MessagePack | Unmarshal | 490616 | 316498 | 5700 |
| Large Payload | CBOR | Unmarshal | 681457 | 329722 | 6724 |
| Large Payload | JSON | Unmarshal | 2067796 | 559132 | 7351 |
| Medium Payload | BEVE ZeroCopy | Marshal | 6771 | 3 | 0 |
| Medium Payload | BEVE | Marshal | 9263 | 18441 | 1 |
| Medium Payload | CBOR | Marshal | 23254 | 27284 | 1 |
| Medium Payload | Sonic | Marshal | 29725 | 22217 | 3 |
| Medium Payload | MessagePack | Marshal | 31500 | 65782 | 22 |
| Medium Payload | JSON | Marshal | 39092 | 21988 | 8 |
| Medium Payload | BEVE | Unmarshal | 24113 | 31263 | 59 |
| Medium Payload | Sonic | Unmarshal | 26924 | 34533 | 33 |
| Medium Payload | MessagePack | Unmarshal | 50101 | 32527 | 598 |
| Medium Payload | CBOR | Unmarshal | 75334 | 38712 | 796 |
| Medium Payload | JSON | Unmarshal | 181379 | 49432 | 652 |
| Small Struct | BEVE ZeroCopy | Marshal | 762 | 0 | 0 |
| Small Struct | BEVE | Marshal | 1346 | 2688 | 1 |
| Small Struct | CBOR | Marshal | 1878 | 2048 | 1 |
| Small Struct | Sonic | Marshal | 3433 | 2774 | 2 |
| Small Struct | MessagePack | Marshal | 3591 | 8201 | 9 |
| Small Struct | JSON | Marshal | 4384 | 2688 | 1 |
| Small Struct | BEVE | Unmarshal | 1446 | 2360 | 4 |
| Small Struct | Sonic | Unmarshal | 2867 | 3989 | 6 |
| Small Struct | MessagePack | Unmarshal | 3198 | 2304 | 50 |
| Small Struct | CBOR | Unmarshal | 7904 | 4808 | 103 |
| Small Struct | JSON | Unmarshal | 11202 | 3688 | 52 |
