# Apple M1 (Virtual) — Darwin

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE ZeroCopy | Marshal | 56835 | 26 | 0 |
| Large Payload | BEVE | Marshal | 80999 | 196650 | 1 |
| Large Payload | CBOR | Marshal | 130499 | 188548 | 1 |
| Large Payload | MessagePack | Marshal | 185679 | 526757 | 115 |
| Large Payload | JSON | Marshal | 329153 | 221473 | 8 |
| Large Payload | Sonic | Marshal | 382677 | 205294 | 3 |
| Large Payload | BEVE | Unmarshal | 158555 | 275028 | 419 |
| Large Payload | Sonic | Unmarshal | 282609 | 346043 | 211 |
| Large Payload | MessagePack | Unmarshal | 369964 | 350680 | 6394 |
| Large Payload | CBOR | Unmarshal | 470836 | 314521 | 6408 |
| Large Payload | JSON | Unmarshal | 1653064 | 526124 | 6795 |
| Medium Payload | BEVE ZeroCopy | Marshal | 4578 | 0 | 0 |
| Medium Payload | BEVE | Marshal | 8010 | 18437 | 1 |
| Medium Payload | CBOR | Marshal | 13690 | 19085 | 1 |
| Medium Payload | MessagePack | Marshal | 24246 | 65778 | 22 |
| Medium Payload | JSON | Marshal | 34673 | 21993 | 8 |
| Medium Payload | Sonic | Marshal | 36083 | 18626 | 3 |
| Medium Payload | BEVE | Unmarshal | 14515 | 20539 | 59 |
| Medium Payload | Sonic | Unmarshal | 27341 | 34920 | 33 |
| Medium Payload | MessagePack | Unmarshal | 40616 | 38334 | 710 |
| Medium Payload | CBOR | Unmarshal | 47151 | 33208 | 685 |
| Medium Payload | JSON | Unmarshal | 197210 | 64408 | 858 |
| Small Struct | CBOR | Marshal | 415 | 416 | 1 |
| Small Struct | BEVE ZeroCopy | Marshal | 540 | 0 | 0 |
| Small Struct | BEVE | Marshal | 692 | 1408 | 1 |
| Small Struct | JSON | Marshal | 1161 | 704 | 1 |
| Small Struct | MessagePack | Marshal | 2736 | 8201 | 9 |
| Small Struct | Sonic | Marshal | 3226 | 1808 | 2 |
| Small Struct | BEVE | Unmarshal | 643 | 952 | 4 |
| Small Struct | MessagePack | Unmarshal | 1429 | 1176 | 27 |
| Small Struct | Sonic | Unmarshal | 2142 | 3302 | 6 |
| Small Struct | CBOR | Unmarshal | 5317 | 4712 | 100 |
| Small Struct | JSON | Unmarshal | 14232 | 4456 | 76 |
