# AMD EPYC 7763 64-Core Processor — Linux

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE ZeroCopy | Marshal | 80129 | 286 | 2 |
| Large Payload | BEVE | Marshal | 118936 | 188712 | 3 |
| Large Payload | Sonic | Marshal | 167997 | 216897 | 4 |
| Large Payload | CBOR | Marshal | 222334 | 213767 | 2 |
| Large Payload | MessagePack | Marshal | 304225 | 526837 | 115 |
| Large Payload | JSON | Marshal | 436804 | 213635 | 9 |
| Large Payload | BEVE | Unmarshal | 227100 | 265758 | 416 |
| Large Payload | Sonic | Unmarshal | 360641 | 551438 | 577 |
| Large Payload | MessagePack | Unmarshal | 582470 | 359930 | 6576 |
| Large Payload | CBOR | Unmarshal | 727623 | 317210 | 6470 |
| Large Payload | JSON | Unmarshal | 2348153 | 541866 | 7064 |
| Medium Payload | BEVE ZeroCopy | Marshal | 8584 | 138 | 2 |
| Medium Payload | BEVE | Marshal | 15868 | 24782 | 3 |
| Medium Payload | Sonic | Marshal | 18772 | 25569 | 4 |
| Medium Payload | CBOR | Marshal | 21794 | 20599 | 2 |
| Medium Payload | JSON | Marshal | 36645 | 16708 | 9 |
| Medium Payload | MessagePack | Marshal | 38172 | 65839 | 22 |
| Medium Payload | BEVE | Unmarshal | 20613 | 21149 | 57 |
| Medium Payload | Sonic | Unmarshal | 43237 | 65573 | 79 |
| Medium Payload | MessagePack | Unmarshal | 62517 | 41105 | 769 |
| Medium Payload | CBOR | Unmarshal | 80124 | 37144 | 761 |
| Medium Payload | JSON | Unmarshal | 223332 | 50944 | 687 |
| Small Struct | BEVE ZeroCopy | Marshal | 476 | 289 | 2 |
| Small Struct | Sonic | Marshal | 650 | 686 | 3 |
| Small Struct | MessagePack | Marshal | 1090 | 1152 | 6 |
| Small Struct | BEVE | Marshal | 1309 | 2081 | 3 |
| Small Struct | JSON | Marshal | 2001 | 1040 | 2 |
| Small Struct | CBOR | Marshal | 2548 | 2835 | 2 |
| Small Struct | BEVE | Unmarshal | 1551 | 2360 | 4 |
| Small Struct | Sonic | Unmarshal | 4108 | 7409 | 10 |
| Small Struct | MessagePack | Unmarshal | 4652 | 3464 | 72 |
| Small Struct | JSON | Unmarshal | 8200 | 2048 | 36 |
| Small Struct | CBOR | Unmarshal | 8964 | 4744 | 101 |
