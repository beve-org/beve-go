# AMD EPYC 7763 64-Core Processor — Linux

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE ZeroCopy | Marshal | 81171 | 39 | 0 |
| Large Payload | BEVE | Marshal | 116739 | 196653 | 1 |
| Large Payload | Sonic | Marshal | 170341 | 240333 | 3 |
| Large Payload | CBOR | Marshal | 203026 | 180366 | 1 |
| Large Payload | MessagePack | Marshal | 325349 | 526780 | 115 |
| Large Payload | JSON | Marshal | 431252 | 205116 | 8 |
| Large Payload | BEVE | Unmarshal | 246828 | 259999 | 417 |
| Large Payload | Sonic | Unmarshal | 377489 | 544864 | 571 |
| Large Payload | MessagePack | Unmarshal | 559520 | 332244 | 6008 |
| Large Payload | CBOR | Unmarshal | 769447 | 347306 | 7063 |
| Large Payload | JSON | Unmarshal | 2225278 | 527979 | 6842 |
| Medium Payload | BEVE ZeroCopy | Marshal | 7074 | 0 | 0 |
| Medium Payload | BEVE | Marshal | 13074 | 21769 | 1 |
| Medium Payload | Sonic | Marshal | 15211 | 20932 | 3 |
| Medium Payload | CBOR | Marshal | 20384 | 18449 | 1 |
| Medium Payload | MessagePack | Marshal | 36208 | 65782 | 22 |
| Medium Payload | JSON | Marshal | 42736 | 20708 | 8 |
| Medium Payload | BEVE | Unmarshal | 25301 | 29087 | 59 |
| Medium Payload | Sonic | Unmarshal | 40840 | 60495 | 72 |
| Medium Payload | MessagePack | Unmarshal | 55767 | 34880 | 642 |
| Medium Payload | CBOR | Unmarshal | 66370 | 29544 | 610 |
| Medium Payload | JSON | Unmarshal | 268146 | 68840 | 876 |
| Small Struct | BEVE ZeroCopy | Marshal | 543 | 0 | 0 |
| Small Struct | Sonic | Marshal | 1233 | 1601 | 2 |
| Small Struct | CBOR | Marshal | 1252 | 1152 | 1 |
| Small Struct | BEVE | Marshal | 1620 | 2688 | 1 |
| Small Struct | JSON | Marshal | 2212 | 1024 | 1 |
| Small Struct | MessagePack | Marshal | 2425 | 4104 | 8 |
| Small Struct | BEVE | Unmarshal | 706 | 440 | 4 |
| Small Struct | Sonic | Unmarshal | 1744 | 2229 | 8 |
| Small Struct | MessagePack | Unmarshal | 1913 | 976 | 23 |
| Small Struct | CBOR | Unmarshal | 5537 | 3080 | 65 |
| Small Struct | JSON | Unmarshal | 25024 | 7560 | 102 |
