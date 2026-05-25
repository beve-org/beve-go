# AMD EPYC 9V74 80-Core Processor — Linux

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE ZeroCopy | Marshal | 59254 | 52 | 0 |
| Large Payload | BEVE | Marshal | 99131 | 196708 | 1 |
| Large Payload | Sonic | Marshal | 126405 | 208091 | 3 |
| Large Payload | CBOR | Marshal | 156340 | 188590 | 1 |
| Large Payload | MessagePack | Marshal | 271575 | 526784 | 115 |
| Large Payload | JSON | Marshal | 333681 | 205089 | 8 |
| Large Payload | BEVE | Unmarshal | 193376 | 271427 | 417 |
| Large Payload | Sonic | Unmarshal | 321134 | 542656 | 574 |
| Large Payload | MessagePack | Unmarshal | 442279 | 359676 | 6568 |
| Large Payload | CBOR | Unmarshal | 600591 | 316922 | 6474 |
| Large Payload | JSON | Unmarshal | 1542639 | 502162 | 6611 |
| Medium Payload | BEVE ZeroCopy | Marshal | 6627 | 3 | 0 |
| Medium Payload | BEVE | Marshal | 7685 | 16389 | 1 |
| Medium Payload | Sonic | Marshal | 13142 | 22415 | 3 |
| Medium Payload | CBOR | Marshal | 13958 | 16399 | 1 |
| Medium Payload | MessagePack | Marshal | 29057 | 65783 | 22 |
| Medium Payload | JSON | Marshal | 35586 | 24811 | 8 |
| Medium Payload | BEVE | Unmarshal | 19444 | 31039 | 59 |
| Medium Payload | Sonic | Unmarshal | 30032 | 49825 | 69 |
| Medium Payload | MessagePack | Unmarshal | 39900 | 33263 | 610 |
| Medium Payload | CBOR | Unmarshal | 67969 | 39496 | 811 |
| Medium Payload | JSON | Unmarshal | 153836 | 53776 | 688 |
| Small Struct | BEVE ZeroCopy | Marshal | 396 | 0 | 0 |
| Small Struct | CBOR | Marshal | 526 | 576 | 1 |
| Small Struct | BEVE | Marshal | 733 | 1536 | 1 |
| Small Struct | MessagePack | Marshal | 1147 | 2056 | 7 |
| Small Struct | Sonic | Marshal | 1655 | 3180 | 2 |
| Small Struct | JSON | Marshal | 3777 | 2305 | 1 |
| Small Struct | BEVE | Unmarshal | 986 | 1720 | 4 |
| Small Struct | MessagePack | Unmarshal | 1140 | 592 | 15 |
| Small Struct | Sonic | Unmarshal | 2504 | 4694 | 9 |
| Small Struct | CBOR | Unmarshal | 2771 | 1488 | 34 |
| Small Struct | JSON | Unmarshal | 14794 | 4680 | 83 |
