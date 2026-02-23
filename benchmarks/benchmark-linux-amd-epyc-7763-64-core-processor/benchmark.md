# AMD EPYC 7763 64-Core Processor — Linux

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE ZeroCopy | Marshal | 78582 | 26 | 0 |
| Large Payload | BEVE | Marshal | 115562 | 196666 | 1 |
| Large Payload | Sonic | Marshal | 160614 | 223825 | 3 |
| Large Payload | CBOR | Marshal | 214147 | 196759 | 1 |
| Large Payload | MessagePack | Marshal | 321552 | 526778 | 115 |
| Large Payload | JSON | Marshal | 443938 | 213282 | 8 |
| Large Payload | BEVE | Unmarshal | 254394 | 279906 | 419 |
| Large Payload | Sonic | Unmarshal | 364257 | 532338 | 572 |
| Large Payload | MessagePack | Unmarshal | 535467 | 314463 | 5646 |
| Large Payload | CBOR | Unmarshal | 735968 | 324715 | 6613 |
| Large Payload | JSON | Unmarshal | 2349939 | 540170 | 7102 |
| Medium Payload | BEVE ZeroCopy | Marshal | 9108 | 3 | 0 |
| Medium Payload | BEVE | Marshal | 12133 | 19081 | 1 |
| Medium Payload | Sonic | Marshal | 14202 | 19499 | 3 |
| Medium Payload | CBOR | Marshal | 24800 | 24601 | 1 |
| Medium Payload | MessagePack | Marshal | 26214 | 33007 | 21 |
| Medium Payload | JSON | Marshal | 41474 | 20708 | 8 |
| Medium Payload | BEVE | Unmarshal | 24715 | 28543 | 59 |
| Medium Payload | Sonic | Unmarshal | 31674 | 44064 | 63 |
| Medium Payload | MessagePack | Unmarshal | 51943 | 32591 | 596 |
| Medium Payload | CBOR | Unmarshal | 70379 | 31912 | 662 |
| Medium Payload | JSON | Unmarshal | 223776 | 55000 | 720 |
| Small Struct | BEVE ZeroCopy | Marshal | 670 | 0 | 0 |
| Small Struct | Sonic | Marshal | 704 | 791 | 2 |
| Small Struct | BEVE | Marshal | 823 | 1152 | 1 |
| Small Struct | MessagePack | Marshal | 1034 | 1032 | 6 |
| Small Struct | CBOR | Marshal | 1378 | 1280 | 1 |
| Small Struct | JSON | Marshal | 3511 | 1792 | 1 |
| Small Struct | BEVE | Unmarshal | 855 | 728 | 4 |
| Small Struct | Sonic | Unmarshal | 4660 | 7815 | 10 |
| Small Struct | MessagePack | Unmarshal | 6109 | 4801 | 102 |
| Small Struct | CBOR | Unmarshal | 6481 | 3616 | 78 |
| Small Struct | JSON | Unmarshal | 10439 | 2344 | 45 |
