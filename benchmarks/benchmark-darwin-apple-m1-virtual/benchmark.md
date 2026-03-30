# Apple M1 (Virtual) — Darwin

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE ZeroCopy | Marshal | 54622 | 26 | 0 |
| Large Payload | BEVE | Marshal | 99440 | 188457 | 1 |
| Large Payload | CBOR | Marshal | 134464 | 188551 | 1 |
| Large Payload | MessagePack | Marshal | 189668 | 526754 | 115 |
| Large Payload | JSON | Marshal | 401229 | 213306 | 8 |
| Large Payload | Sonic | Marshal | 448029 | 213686 | 3 |
| Large Payload | BEVE | Unmarshal | 256089 | 280083 | 418 |
| Large Payload | Sonic | Unmarshal | 304925 | 337465 | 209 |
| Large Payload | MessagePack | Unmarshal | 556981 | 331650 | 5998 |
| Large Payload | CBOR | Unmarshal | 635894 | 325961 | 6647 |
| Large Payload | JSON | Unmarshal | 2356302 | 527538 | 6847 |
| Medium Payload | BEVE ZeroCopy | Marshal | 11616 | 3 | 0 |
| Medium Payload | BEVE | Marshal | 11848 | 19075 | 1 |
| Medium Payload | CBOR | Marshal | 13356 | 16395 | 1 |
| Medium Payload | MessagePack | Marshal | 30922 | 65778 | 22 |
| Medium Payload | Sonic | Marshal | 45268 | 19342 | 3 |
| Medium Payload | JSON | Marshal | 52492 | 27495 | 8 |
| Medium Payload | BEVE | Unmarshal | 23923 | 28573 | 59 |
| Medium Payload | Sonic | Unmarshal | 29667 | 37626 | 33 |
| Medium Payload | MessagePack | Unmarshal | 36874 | 32701 | 603 |
| Medium Payload | CBOR | Unmarshal | 53614 | 33768 | 697 |
| Medium Payload | JSON | Unmarshal | 175053 | 51832 | 662 |
| Small Struct | BEVE ZeroCopy | Marshal | 390 | 0 | 0 |
| Small Struct | BEVE | Marshal | 779 | 640 | 1 |
| Small Struct | CBOR | Marshal | 1870 | 2304 | 1 |
| Small Struct | MessagePack | Marshal | 3638 | 4104 | 8 |
| Small Struct | JSON | Marshal | 5522 | 2688 | 1 |
| Small Struct | Sonic | Marshal | 6518 | 3103 | 2 |
| Small Struct | BEVE | Unmarshal | 1246 | 1720 | 4 |
| Small Struct | Sonic | Unmarshal | 2339 | 2726 | 6 |
| Small Struct | CBOR | Unmarshal | 3596 | 2048 | 45 |
| Small Struct | MessagePack | Unmarshal | 4398 | 3136 | 66 |
| Small Struct | JSON | Unmarshal | 28954 | 8040 | 117 |
