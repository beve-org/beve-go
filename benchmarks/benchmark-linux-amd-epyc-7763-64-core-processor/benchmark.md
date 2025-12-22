# AMD EPYC 7763 64-Core Processor — Linux

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE ZeroCopy | Marshal | 83429 | 26 | 0 |
| Large Payload | BEVE | Marshal | 113427 | 188460 | 1 |
| Large Payload | Sonic | Marshal | 155972 | 215724 | 3 |
| Large Payload | CBOR | Marshal | 218181 | 196890 | 1 |
| Large Payload | MessagePack | Marshal | 324413 | 526780 | 115 |
| Large Payload | JSON | Marshal | 447231 | 213309 | 8 |
| Large Payload | BEVE | Unmarshal | 259621 | 289860 | 415 |
| Large Payload | Sonic | Unmarshal | 366160 | 541908 | 573 |
| Large Payload | MessagePack | Unmarshal | 573795 | 356745 | 6501 |
| Large Payload | CBOR | Unmarshal | 700491 | 304521 | 6207 |
| Large Payload | JSON | Unmarshal | 2339388 | 519298 | 6803 |
| Medium Payload | BEVE ZeroCopy | Marshal | 8539 | 3 | 0 |
| Medium Payload | BEVE | Marshal | 13148 | 21768 | 1 |
| Medium Payload | Sonic | Marshal | 17489 | 25234 | 3 |
| Medium Payload | CBOR | Marshal | 24903 | 24601 | 1 |
| Medium Payload | MessagePack | Marshal | 35983 | 65782 | 22 |
| Medium Payload | JSON | Marshal | 46002 | 24808 | 8 |
| Medium Payload | BEVE | Unmarshal | 23398 | 27198 | 59 |
| Medium Payload | Sonic | Unmarshal | 40177 | 61876 | 70 |
| Medium Payload | MessagePack | Unmarshal | 52079 | 32623 | 598 |
| Medium Payload | CBOR | Unmarshal | 73222 | 34264 | 706 |
| Medium Payload | JSON | Unmarshal | 275276 | 66680 | 879 |
| Small Struct | Sonic | Marshal | 548 | 613 | 2 |
| Small Struct | BEVE ZeroCopy | Marshal | 616 | 0 | 0 |
| Small Struct | BEVE | Marshal | 720 | 1024 | 1 |
| Small Struct | MessagePack | Marshal | 1348 | 1032 | 6 |
| Small Struct | CBOR | Marshal | 1861 | 1280 | 1 |
| Small Struct | JSON | Marshal | 4413 | 2304 | 1 |
| Small Struct | BEVE | Unmarshal | 1205 | 1592 | 4 |
| Small Struct | CBOR | Unmarshal | 1583 | 424 | 12 |
| Small Struct | Sonic | Unmarshal | 1651 | 2112 | 8 |
| Small Struct | MessagePack | Unmarshal | 1934 | 976 | 23 |
| Small Struct | JSON | Unmarshal | 5810 | 1256 | 26 |
