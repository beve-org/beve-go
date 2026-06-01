# AMD EPYC 7763 64-Core Processor — Linux

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE ZeroCopy | Marshal | 80509 | 39 | 0 |
| Large Payload | BEVE | Marshal | 127478 | 188490 | 1 |
| Large Payload | Sonic | Marshal | 171568 | 224444 | 3 |
| Large Payload | CBOR | Marshal | 211929 | 196759 | 1 |
| Large Payload | MessagePack | Marshal | 329959 | 526780 | 115 |
| Large Payload | JSON | Marshal | 445062 | 205090 | 8 |
| Large Payload | BEVE | Unmarshal | 237071 | 259102 | 418 |
| Large Payload | Sonic | Unmarshal | 378394 | 545260 | 584 |
| Large Payload | MessagePack | Unmarshal | 558205 | 336324 | 6091 |
| Large Payload | CBOR | Unmarshal | 740801 | 331867 | 6766 |
| Large Payload | JSON | Unmarshal | 2425300 | 579725 | 7544 |
| Medium Payload | BEVE ZeroCopy | Marshal | 8409 | 3 | 0 |
| Medium Payload | BEVE | Marshal | 15893 | 18438 | 1 |
| Medium Payload | Sonic | Marshal | 19941 | 25381 | 3 |
| Medium Payload | CBOR | Marshal | 26604 | 21781 | 1 |
| Medium Payload | MessagePack | Marshal | 42923 | 65783 | 22 |
| Medium Payload | JSON | Marshal | 48342 | 21992 | 8 |
| Medium Payload | BEVE | Unmarshal | 30217 | 34208 | 59 |
| Medium Payload | Sonic | Unmarshal | 36263 | 46449 | 63 |
| Medium Payload | MessagePack | Unmarshal | 60324 | 35552 | 656 |
| Medium Payload | CBOR | Unmarshal | 74834 | 33112 | 680 |
| Medium Payload | JSON | Unmarshal | 204083 | 46552 | 620 |
| Small Struct | BEVE ZeroCopy | Marshal | 574 | 0 | 0 |
| Small Struct | Sonic | Marshal | 822 | 755 | 2 |
| Small Struct | BEVE | Marshal | 871 | 1024 | 1 |
| Small Struct | JSON | Marshal | 3104 | 1408 | 1 |
| Small Struct | CBOR | Marshal | 3117 | 3073 | 1 |
| Small Struct | MessagePack | Marshal | 3309 | 4104 | 8 |
| Small Struct | BEVE | Unmarshal | 2309 | 2616 | 4 |
| Small Struct | Sonic | Unmarshal | 3355 | 3787 | 9 |
| Small Struct | CBOR | Unmarshal | 4885 | 1760 | 40 |
| Small Struct | MessagePack | Unmarshal | 7708 | 4697 | 99 |
| Small Struct | JSON | Unmarshal | 35422 | 8008 | 116 |
