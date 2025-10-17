# AMD EPYC 7763 64-Core Processor — Linux

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE ZeroCopy | Marshal | 80280 | 26 | 0 |
| Large Payload | BEVE | Marshal | 116958 | 196666 | 1 |
| Large Payload | Sonic | Marshal | 156462 | 215523 | 3 |
| Large Payload | CBOR | Marshal | 207543 | 196811 | 1 |
| Large Payload | MessagePack | Marshal | 309170 | 526778 | 115 |
| Large Payload | JSON | Marshal | 443264 | 213309 | 8 |
| Large Payload | BEVE | Unmarshal | 241710 | 273441 | 418 |
| Large Payload | Sonic | Unmarshal | 342142 | 512131 | 566 |
| Large Payload | MessagePack | Unmarshal | 565804 | 350391 | 6361 |
| Large Payload | CBOR | Unmarshal | 705516 | 316234 | 6451 |
| Large Payload | JSON | Unmarshal | 2269856 | 538241 | 6975 |
| Medium Payload | BEVE ZeroCopy | Marshal | 7687 | 0 | 0 |
| Medium Payload | BEVE | Marshal | 11277 | 18437 | 1 |
| Medium Payload | Sonic | Marshal | 16442 | 22182 | 3 |
| Medium Payload | CBOR | Marshal | 18463 | 16397 | 1 |
| Medium Payload | MessagePack | Marshal | 37087 | 65783 | 22 |
| Medium Payload | JSON | Marshal | 40307 | 20708 | 8 |
| Medium Payload | BEVE | Unmarshal | 25384 | 30623 | 59 |
| Medium Payload | Sonic | Unmarshal | 35554 | 50919 | 73 |
| Medium Payload | MessagePack | Unmarshal | 57360 | 36944 | 686 |
| Medium Payload | CBOR | Unmarshal | 65319 | 29704 | 606 |
| Medium Payload | JSON | Unmarshal | 278778 | 72120 | 922 |
| Small Struct | CBOR | Marshal | 630 | 448 | 1 |
| Small Struct | BEVE ZeroCopy | Marshal | 762 | 0 | 0 |
| Small Struct | Sonic | Marshal | 1305 | 1857 | 2 |
| Small Struct | BEVE | Marshal | 1452 | 2689 | 1 |
| Small Struct | MessagePack | Marshal | 1614 | 2056 | 7 |
| Small Struct | JSON | Marshal | 3347 | 1792 | 1 |
| Small Struct | Sonic | Unmarshal | 1224 | 1338 | 7 |
| Small Struct | BEVE | Unmarshal | 1259 | 1592 | 4 |
| Small Struct | CBOR | Unmarshal | 2561 | 1064 | 25 |
| Small Struct | MessagePack | Unmarshal | 5206 | 3968 | 84 |
| Small Struct | JSON | Unmarshal | 17595 | 4424 | 75 |
