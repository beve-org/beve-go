# AMD EPYC 7763 64-Core Processor — Linux

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE ZeroCopy | Marshal | 79739 | 26 | 0 |
| Large Payload | BEVE | Marshal | 116683 | 188514 | 1 |
| Large Payload | Sonic | Marshal | 156804 | 215795 | 3 |
| Large Payload | CBOR | Marshal | 208700 | 196812 | 1 |
| Large Payload | MessagePack | Marshal | 316485 | 526778 | 115 |
| Large Payload | JSON | Marshal | 452566 | 221476 | 8 |
| Large Payload | BEVE | Unmarshal | 241102 | 279329 | 416 |
| Large Payload | Sonic | Unmarshal | 337526 | 506615 | 561 |
| Large Payload | MessagePack | Unmarshal | 556917 | 344772 | 6267 |
| Large Payload | CBOR | Unmarshal | 718233 | 322539 | 6577 |
| Large Payload | JSON | Unmarshal | 2317363 | 544754 | 7187 |
| Medium Payload | BEVE ZeroCopy | Marshal | 8333 | 3 | 0 |
| Medium Payload | BEVE | Marshal | 11359 | 18437 | 1 |
| Medium Payload | Sonic | Marshal | 17323 | 25131 | 3 |
| Medium Payload | CBOR | Marshal | 24988 | 24598 | 1 |
| Medium Payload | MessagePack | Marshal | 26427 | 33007 | 21 |
| Medium Payload | JSON | Marshal | 45911 | 24808 | 8 |
| Medium Payload | BEVE | Unmarshal | 25021 | 29407 | 59 |
| Medium Payload | Sonic | Unmarshal | 38377 | 54301 | 75 |
| Medium Payload | MessagePack | Unmarshal | 49585 | 29854 | 540 |
| Medium Payload | CBOR | Unmarshal | 68056 | 29752 | 615 |
| Medium Payload | JSON | Unmarshal | 206630 | 48120 | 642 |
| Small Struct | BEVE | Marshal | 363 | 384 | 1 |
| Small Struct | BEVE ZeroCopy | Marshal | 618 | 0 | 0 |
| Small Struct | CBOR | Marshal | 1021 | 896 | 1 |
| Small Struct | Sonic | Marshal | 1083 | 1438 | 2 |
| Small Struct | MessagePack | Marshal | 1553 | 2056 | 7 |
| Small Struct | JSON | Marshal | 3165 | 1792 | 1 |
| Small Struct | BEVE | Unmarshal | 1044 | 1208 | 4 |
| Small Struct | MessagePack | Unmarshal | 1308 | 448 | 12 |
| Small Struct | Sonic | Unmarshal | 4289 | 7788 | 10 |
| Small Struct | CBOR | Unmarshal | 8136 | 4712 | 100 |
| Small Struct | JSON | Unmarshal | 17985 | 4392 | 74 |
