# AMD EPYC 7763 64-Core Processor — Linux

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE ZeroCopy | Marshal | 80549 | 26 | 0 |
| Large Payload | BEVE | Marshal | 118773 | 188462 | 1 |
| Large Payload | Sonic | Marshal | 164731 | 224013 | 3 |
| Large Payload | CBOR | Marshal | 212035 | 196811 | 1 |
| Large Payload | MessagePack | Marshal | 324034 | 526779 | 115 |
| Large Payload | JSON | Marshal | 440641 | 213309 | 8 |
| Large Payload | BEVE | Unmarshal | 240905 | 260253 | 417 |
| Large Payload | Sonic | Unmarshal | 363403 | 525034 | 570 |
| Large Payload | MessagePack | Unmarshal | 596648 | 373853 | 6861 |
| Large Payload | CBOR | Unmarshal | 694653 | 302489 | 6177 |
| Large Payload | JSON | Unmarshal | 2426940 | 539435 | 6932 |
| Medium Payload | BEVE ZeroCopy | Marshal | 7574 | 5 | 0 |
| Medium Payload | BEVE | Marshal | 12423 | 19082 | 1 |
| Medium Payload | Sonic | Marshal | 17715 | 25110 | 3 |
| Medium Payload | CBOR | Marshal | 21104 | 19086 | 1 |
| Medium Payload | MessagePack | Marshal | 37969 | 65782 | 22 |
| Medium Payload | JSON | Marshal | 39924 | 19303 | 8 |
| Medium Payload | BEVE | Unmarshal | 24764 | 27422 | 59 |
| Medium Payload | Sonic | Unmarshal | 40273 | 50764 | 69 |
| Medium Payload | MessagePack | Unmarshal | 55364 | 35248 | 649 |
| Medium Payload | CBOR | Unmarshal | 63244 | 27784 | 575 |
| Medium Payload | JSON | Unmarshal | 194185 | 43800 | 558 |
| Small Struct | BEVE ZeroCopy | Marshal | 587 | 0 | 0 |
| Small Struct | JSON | Marshal | 988 | 384 | 1 |
| Small Struct | BEVE | Marshal | 1095 | 1280 | 1 |
| Small Struct | CBOR | Marshal | 1402 | 1280 | 1 |
| Small Struct | Sonic | Marshal | 2231 | 3194 | 2 |
| Small Struct | MessagePack | Marshal | 4710 | 8201 | 9 |
| Small Struct | MessagePack | Unmarshal | 1160 | 304 | 9 |
| Small Struct | BEVE | Unmarshal | 1750 | 2616 | 4 |
| Small Struct | Sonic | Unmarshal | 3132 | 4438 | 9 |
| Small Struct | CBOR | Unmarshal | 7804 | 4392 | 94 |
| Small Struct | JSON | Unmarshal | 21673 | 4680 | 83 |
