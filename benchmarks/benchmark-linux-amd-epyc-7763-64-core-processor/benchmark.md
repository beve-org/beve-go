# AMD EPYC 7763 64-Core Processor — Linux

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE ZeroCopy | Marshal | 79135 | 26 | 0 |
| Large Payload | BEVE | Marshal | 123807 | 204847 | 1 |
| Large Payload | Sonic | Marshal | 159658 | 215445 | 3 |
| Large Payload | CBOR | Marshal | 196433 | 180367 | 1 |
| Large Payload | MessagePack | Marshal | 311873 | 526778 | 115 |
| Large Payload | JSON | Marshal | 438223 | 213335 | 8 |
| Large Payload | BEVE | Unmarshal | 245231 | 289956 | 416 |
| Large Payload | Sonic | Unmarshal | 349680 | 540562 | 571 |
| Large Payload | MessagePack | Unmarshal | 548317 | 331188 | 5984 |
| Large Payload | CBOR | Unmarshal | 702342 | 315625 | 6434 |
| Large Payload | JSON | Unmarshal | 2200810 | 520132 | 6728 |
| Medium Payload | BEVE ZeroCopy | Marshal | 8780 | 5 | 0 |
| Medium Payload | BEVE | Marshal | 14153 | 18439 | 1 |
| Medium Payload | Sonic | Marshal | 17843 | 25023 | 3 |
| Medium Payload | CBOR | Marshal | 21978 | 20499 | 1 |
| Medium Payload | MessagePack | Marshal | 35612 | 65782 | 22 |
| Medium Payload | JSON | Marshal | 55181 | 24808 | 8 |
| Medium Payload | BEVE | Unmarshal | 26428 | 33759 | 59 |
| Medium Payload | Sonic | Unmarshal | 38741 | 59017 | 75 |
| Medium Payload | MessagePack | Unmarshal | 56021 | 35520 | 658 |
| Medium Payload | CBOR | Unmarshal | 62600 | 27824 | 572 |
| Medium Payload | JSON | Unmarshal | 196494 | 47608 | 615 |
| Small Struct | BEVE ZeroCopy | Marshal | 376 | 0 | 0 |
| Small Struct | Sonic | Marshal | 888 | 684 | 2 |
| Small Struct | BEVE | Marshal | 919 | 1024 | 1 |
| Small Struct | JSON | Marshal | 1641 | 704 | 1 |
| Small Struct | CBOR | Marshal | 1894 | 896 | 1 |
| Small Struct | MessagePack | Marshal | 2727 | 2056 | 7 |
| Small Struct | BEVE | Unmarshal | 952 | 312 | 3 |
| Small Struct | Sonic | Unmarshal | 2173 | 2096 | 8 |
| Small Struct | JSON | Unmarshal | 3025 | 488 | 13 |
| Small Struct | MessagePack | Unmarshal | 7722 | 4833 | 103 |
| Small Struct | CBOR | Unmarshal | 10813 | 5184 | 107 |
