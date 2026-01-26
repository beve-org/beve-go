# Unknown CPU — Windows

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE ZeroCopy | Marshal | 77136 | 52 | 0 |
| Large Payload | BEVE | Marshal | 120326 | 196671 | 1 |
| Large Payload | Sonic | Marshal | 152949 | 198664 | 3 |
| Large Payload | CBOR | Marshal | 210976 | 188517 | 1 |
| Large Payload | MessagePack | Marshal | 323286 | 526709 | 115 |
| Large Payload | JSON | Marshal | 484106 | 213267 | 8 |
| Large Payload | BEVE | Unmarshal | 292095 | 270502 | 418 |
| Large Payload | Sonic | Unmarshal | 441135 | 538895 | 572 |
| Large Payload | MessagePack | Unmarshal | 664049 | 341041 | 6180 |
| Large Payload | CBOR | Unmarshal | 899426 | 329913 | 6729 |
| Large Payload | JSON | Unmarshal | 2785071 | 571508 | 7476 |
| Medium Payload | BEVE ZeroCopy | Marshal | 8835 | 5 | 0 |
| Medium Payload | BEVE | Marshal | 11674 | 16388 | 1 |
| Medium Payload | Sonic | Marshal | 20452 | 24957 | 3 |
| Medium Payload | CBOR | Marshal | 22091 | 18437 | 1 |
| Medium Payload | MessagePack | Marshal | 33143 | 65771 | 22 |
| Medium Payload | JSON | Marshal | 54205 | 24809 | 8 |
| Medium Payload | BEVE | Unmarshal | 27935 | 28027 | 59 |
| Medium Payload | Sonic | Unmarshal | 43627 | 48779 | 71 |
| Medium Payload | MessagePack | Unmarshal | 77680 | 41293 | 774 |
| Medium Payload | CBOR | Unmarshal | 100832 | 38600 | 791 |
| Medium Payload | JSON | Unmarshal | 248182 | 46264 | 646 |
| Small Struct | BEVE ZeroCopy | Marshal | 248 | 0 | 0 |
| Small Struct | CBOR | Marshal | 551 | 352 | 1 |
| Small Struct | Sonic | Marshal | 1063 | 1310 | 2 |
| Small Struct | MessagePack | Marshal | 1073 | 1032 | 6 |
| Small Struct | JSON | Marshal | 1606 | 704 | 1 |
| Small Struct | BEVE | Marshal | 1625 | 1536 | 1 |
| Small Struct | BEVE | Unmarshal | 820 | 472 | 4 |
| Small Struct | Sonic | Unmarshal | 3283 | 3637 | 9 |
| Small Struct | CBOR | Unmarshal | 4039 | 1728 | 39 |
| Small Struct | MessagePack | Unmarshal | 7009 | 4256 | 89 |
| Small Struct | JSON | Unmarshal | 27781 | 7304 | 94 |
