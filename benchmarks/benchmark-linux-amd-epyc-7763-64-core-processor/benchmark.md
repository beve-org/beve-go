# AMD EPYC 7763 64-Core Processor — Linux

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE ZeroCopy | Marshal | 81535 | 52 | 0 |
| Large Payload | BEVE | Marshal | 113494 | 188487 | 1 |
| Large Payload | Sonic | Marshal | 162935 | 223904 | 3 |
| Large Payload | CBOR | Marshal | 202436 | 180420 | 1 |
| Large Payload | MessagePack | Marshal | 337995 | 526780 | 115 |
| Large Payload | JSON | Marshal | 436831 | 205142 | 8 |
| Large Payload | BEVE | Unmarshal | 245807 | 277058 | 418 |
| Large Payload | Sonic | Unmarshal | 343564 | 508211 | 563 |
| Large Payload | MessagePack | Unmarshal | 588739 | 359130 | 6543 |
| Large Payload | CBOR | Unmarshal | 739735 | 323562 | 6595 |
| Large Payload | JSON | Unmarshal | 2290351 | 514874 | 6736 |
| Medium Payload | BEVE ZeroCopy | Marshal | 7215 | 0 | 0 |
| Medium Payload | Sonic | Marshal | 15110 | 20935 | 3 |
| Medium Payload | BEVE | Marshal | 17369 | 24584 | 1 |
| Medium Payload | CBOR | Marshal | 21268 | 19089 | 1 |
| Medium Payload | JSON | Marshal | 36718 | 18663 | 8 |
| Medium Payload | MessagePack | Marshal | 38066 | 65783 | 22 |
| Medium Payload | BEVE | Unmarshal | 25021 | 27934 | 59 |
| Medium Payload | Sonic | Unmarshal | 39139 | 57737 | 72 |
| Medium Payload | MessagePack | Unmarshal | 64193 | 38817 | 729 |
| Medium Payload | CBOR | Unmarshal | 78119 | 36760 | 755 |
| Medium Payload | JSON | Unmarshal | 227684 | 52824 | 716 |
| Small Struct | Sonic | Marshal | 636 | 350 | 2 |
| Small Struct | BEVE ZeroCopy | Marshal | 980 | 0 | 0 |
| Small Struct | MessagePack | Marshal | 1889 | 2056 | 7 |
| Small Struct | BEVE | Marshal | 2058 | 2689 | 1 |
| Small Struct | CBOR | Marshal | 3279 | 2688 | 1 |
| Small Struct | JSON | Marshal | 6512 | 3073 | 1 |
| Small Struct | BEVE | Unmarshal | 1918 | 2104 | 4 |
| Small Struct | Sonic | Unmarshal | 3542 | 4160 | 9 |
| Small Struct | MessagePack | Unmarshal | 4488 | 2816 | 60 |
| Small Struct | CBOR | Unmarshal | 6895 | 2888 | 63 |
| Small Struct | JSON | Unmarshal | 8596 | 2056 | 36 |
