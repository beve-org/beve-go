# Intel(R) Xeon(R) Platinum 8370C CPU @ 2.80GHz — Linux

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE ZeroCopy | Marshal | 70498 | 52 | 0 |
| Large Payload | BEVE | Marshal | 108785 | 188460 | 1 |
| Large Payload | Sonic | Marshal | 164076 | 215978 | 3 |
| Large Payload | CBOR | Marshal | 204823 | 205034 | 1 |
| Large Payload | MessagePack | Marshal | 310095 | 526777 | 115 |
| Large Payload | JSON | Marshal | 431750 | 213335 | 8 |
| Large Payload | BEVE | Unmarshal | 233785 | 259037 | 419 |
| Large Payload | Sonic | Unmarshal | 378950 | 536570 | 590 |
| Large Payload | MessagePack | Unmarshal | 563112 | 354012 | 6452 |
| Large Payload | CBOR | Unmarshal | 737232 | 330938 | 6733 |
| Large Payload | JSON | Unmarshal | 2024714 | 510417 | 6661 |
| Medium Payload | BEVE ZeroCopy | Marshal | 8027 | 3 | 0 |
| Medium Payload | BEVE | Marshal | 12193 | 19085 | 1 |
| Medium Payload | Sonic | Marshal | 14521 | 18954 | 3 |
| Medium Payload | CBOR | Marshal | 19283 | 18449 | 1 |
| Medium Payload | MessagePack | Marshal | 35261 | 65783 | 22 |
| Medium Payload | JSON | Marshal | 43561 | 21992 | 8 |
| Medium Payload | BEVE | Unmarshal | 27105 | 34272 | 59 |
| Medium Payload | Sonic | Unmarshal | 48068 | 70487 | 81 |
| Medium Payload | MessagePack | Unmarshal | 60905 | 41425 | 773 |
| Medium Payload | CBOR | Unmarshal | 70099 | 29944 | 619 |
| Medium Payload | JSON | Unmarshal | 194088 | 51784 | 650 |
| Small Struct | BEVE ZeroCopy | Marshal | 495 | 0 | 0 |
| Small Struct | JSON | Marshal | 616 | 192 | 1 |
| Small Struct | BEVE | Marshal | 1048 | 1536 | 1 |
| Small Struct | Sonic | Marshal | 1163 | 1332 | 2 |
| Small Struct | MessagePack | Marshal | 1668 | 2056 | 7 |
| Small Struct | CBOR | Marshal | 1740 | 1792 | 1 |
| Small Struct | BEVE | Unmarshal | 1573 | 2360 | 4 |
| Small Struct | Sonic | Unmarshal | 2475 | 3664 | 9 |
| Small Struct | MessagePack | Unmarshal | 5452 | 4288 | 90 |
| Small Struct | CBOR | Unmarshal | 8712 | 5136 | 105 |
| Small Struct | JSON | Unmarshal | 8741 | 2216 | 41 |
