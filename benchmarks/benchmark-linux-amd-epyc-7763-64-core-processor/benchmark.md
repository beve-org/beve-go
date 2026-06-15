# AMD EPYC 7763 64-Core Processor — Linux

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE ZeroCopy | Marshal | 79600 | 26 | 0 |
| Large Payload | BEVE | Marshal | 123516 | 204900 | 1 |
| Large Payload | Sonic | Marshal | 163435 | 223873 | 3 |
| Large Payload | CBOR | Marshal | 219349 | 205008 | 1 |
| Large Payload | MessagePack | Marshal | 343517 | 526781 | 115 |
| Large Payload | JSON | Marshal | 421072 | 196896 | 8 |
| Large Payload | BEVE | Unmarshal | 248101 | 279969 | 419 |
| Large Payload | Sonic | Unmarshal | 346853 | 500085 | 564 |
| Large Payload | MessagePack | Unmarshal | 601880 | 372380 | 6826 |
| Large Payload | CBOR | Unmarshal | 742715 | 333771 | 6812 |
| Large Payload | JSON | Unmarshal | 2518375 | 545155 | 7191 |
| Medium Payload | BEVE ZeroCopy | Marshal | 8885 | 5 | 0 |
| Medium Payload | BEVE | Marshal | 9797 | 14340 | 1 |
| Medium Payload | CBOR | Marshal | 20546 | 18456 | 1 |
| Medium Payload | Sonic | Marshal | 20774 | 28062 | 3 |
| Medium Payload | JSON | Marshal | 40039 | 19303 | 8 |
| Medium Payload | MessagePack | Marshal | 41114 | 65783 | 22 |
| Medium Payload | BEVE | Unmarshal | 28616 | 26462 | 59 |
| Medium Payload | Sonic | Unmarshal | 40357 | 58105 | 76 |
| Medium Payload | MessagePack | Unmarshal | 45676 | 26894 | 482 |
| Medium Payload | CBOR | Unmarshal | 67437 | 27624 | 570 |
| Medium Payload | JSON | Unmarshal | 252482 | 59896 | 763 |
| Small Struct | BEVE ZeroCopy | Marshal | 865 | 0 | 0 |
| Small Struct | Sonic | Marshal | 1642 | 1595 | 2 |
| Small Struct | BEVE | Marshal | 2060 | 2688 | 1 |
| Small Struct | CBOR | Marshal | 3540 | 3074 | 1 |
| Small Struct | JSON | Marshal | 3889 | 1536 | 1 |
| Small Struct | MessagePack | Marshal | 6552 | 8201 | 9 |
| Small Struct | BEVE | Unmarshal | 1542 | 1464 | 4 |
| Small Struct | Sonic | Unmarshal | 1776 | 2251 | 8 |
| Small Struct | MessagePack | Unmarshal | 1872 | 880 | 21 |
| Small Struct | CBOR | Unmarshal | 2756 | 1192 | 28 |
| Small Struct | JSON | Unmarshal | 30600 | 7912 | 113 |
