# Unknown CPU — Windows

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE ZeroCopy | Marshal | 77753 | 52 | 0 |
| Large Payload | BEVE | Marshal | 119213 | 188451 | 1 |
| Large Payload | Sonic | Marshal | 178442 | 223636 | 3 |
| Large Payload | CBOR | Marshal | 236599 | 188491 | 1 |
| Large Payload | MessagePack | Marshal | 290789 | 526710 | 115 |
| Large Payload | JSON | Marshal | 498184 | 205102 | 8 |
| Large Payload | BEVE | Unmarshal | 299499 | 275526 | 418 |
| Large Payload | Sonic | Unmarshal | 466123 | 556469 | 599 |
| Large Payload | MessagePack | Unmarshal | 684783 | 357912 | 6522 |
| Large Payload | CBOR | Unmarshal | 868442 | 306874 | 6263 |
| Large Payload | JSON | Unmarshal | 2495780 | 504652 | 6640 |
| Medium Payload | BEVE ZeroCopy | Marshal | 8361 | 3 | 0 |
| Medium Payload | BEVE | Marshal | 12430 | 16386 | 1 |
| Medium Payload | Sonic | Marshal | 19665 | 24917 | 3 |
| Medium Payload | CBOR | Marshal | 25760 | 20489 | 1 |
| Medium Payload | MessagePack | Marshal | 26484 | 33002 | 21 |
| Medium Payload | JSON | Marshal | 51867 | 24809 | 8 |
| Medium Payload | BEVE | Unmarshal | 32708 | 31388 | 59 |
| Medium Payload | Sonic | Unmarshal | 48167 | 55089 | 74 |
| Medium Payload | MessagePack | Unmarshal | 71868 | 36172 | 667 |
| Medium Payload | CBOR | Unmarshal | 88342 | 30296 | 627 |
| Medium Payload | JSON | Unmarshal | 251549 | 47688 | 654 |
| Small Struct | BEVE ZeroCopy | Marshal | 311 | 0 | 0 |
| Small Struct | BEVE | Marshal | 535 | 256 | 1 |
| Small Struct | JSON | Marshal | 891 | 288 | 1 |
| Small Struct | MessagePack | Marshal | 2175 | 2056 | 7 |
| Small Struct | Sonic | Marshal | 2365 | 2120 | 2 |
| Small Struct | CBOR | Marshal | 2518 | 2304 | 1 |
| Small Struct | BEVE | Unmarshal | 661 | 312 | 3 |
| Small Struct | CBOR | Unmarshal | 1635 | 328 | 10 |
| Small Struct | Sonic | Unmarshal | 3086 | 3654 | 9 |
| Small Struct | MessagePack | Unmarshal | 5124 | 2816 | 60 |
| Small Struct | JSON | Unmarshal | 15289 | 3848 | 57 |
