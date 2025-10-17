# AMD EPYC 7763 64-Core Processor — Linux

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE ZeroCopy | Marshal | 81778 | 26 | 0 |
| Large Payload | BEVE | Marshal | 115588 | 188526 | 1 |
| Large Payload | Sonic | Marshal | 163351 | 223779 | 3 |
| Large Payload | CBOR | Marshal | 210066 | 196839 | 1 |
| Large Payload | MessagePack | Marshal | 331359 | 526780 | 115 |
| Large Payload | JSON | Marshal | 453900 | 221476 | 8 |
| Large Payload | BEVE | Unmarshal | 239560 | 271488 | 418 |
| Large Payload | Sonic | Unmarshal | 344726 | 525346 | 548 |
| Large Payload | MessagePack | Unmarshal | 571663 | 361800 | 6623 |
| Large Payload | CBOR | Unmarshal | 728174 | 307529 | 6266 |
| Large Payload | JSON | Unmarshal | 2244962 | 533033 | 6922 |
| Medium Payload | BEVE ZeroCopy | Marshal | 7406 | 3 | 0 |
| Medium Payload | BEVE | Marshal | 12142 | 20486 | 1 |
| Medium Payload | Sonic | Marshal | 17745 | 25021 | 3 |
| Medium Payload | CBOR | Marshal | 20038 | 18453 | 1 |
| Medium Payload | MessagePack | Marshal | 36460 | 65782 | 22 |
| Medium Payload | JSON | Marshal | 40643 | 20712 | 8 |
| Medium Payload | BEVE | Unmarshal | 27608 | 29983 | 59 |
| Medium Payload | Sonic | Unmarshal | 46875 | 64927 | 77 |
| Medium Payload | MessagePack | Unmarshal | 56423 | 35760 | 658 |
| Medium Payload | CBOR | Unmarshal | 78160 | 33624 | 692 |
| Medium Payload | JSON | Unmarshal | 220199 | 52632 | 714 |
| Small Struct | BEVE ZeroCopy | Marshal | 329 | 0 | 0 |
| Small Struct | Sonic | Marshal | 463 | 428 | 2 |
| Small Struct | CBOR | Marshal | 790 | 640 | 1 |
| Small Struct | BEVE | Marshal | 1617 | 3073 | 1 |
| Small Struct | MessagePack | Marshal | 2771 | 4104 | 8 |
| Small Struct | JSON | Marshal | 4364 | 2304 | 1 |
| Small Struct | BEVE | Unmarshal | 1655 | 2616 | 4 |
| Small Struct | Sonic | Unmarshal | 2688 | 4166 | 9 |
| Small Struct | CBOR | Unmarshal | 3131 | 1384 | 32 |
| Small Struct | MessagePack | Unmarshal | 3355 | 2272 | 49 |
| Small Struct | JSON | Unmarshal | 14305 | 2440 | 48 |
