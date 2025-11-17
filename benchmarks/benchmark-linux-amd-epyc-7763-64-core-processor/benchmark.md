# AMD EPYC 7763 64-Core Processor — Linux

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE ZeroCopy | Marshal | 79264 | 26 | 0 |
| Large Payload | BEVE | Marshal | 114229 | 180281 | 1 |
| Large Payload | Sonic | Marshal | 159945 | 215777 | 3 |
| Large Payload | CBOR | Marshal | 203753 | 180393 | 1 |
| Large Payload | MessagePack | Marshal | 321910 | 526779 | 115 |
| Large Payload | JSON | Marshal | 450537 | 213283 | 8 |
| Large Payload | BEVE | Unmarshal | 249999 | 276322 | 418 |
| Large Payload | Sonic | Unmarshal | 370810 | 544521 | 572 |
| Large Payload | MessagePack | Unmarshal | 566603 | 337558 | 6131 |
| Large Payload | CBOR | Unmarshal | 700169 | 302858 | 6178 |
| Large Payload | JSON | Unmarshal | 2352388 | 542297 | 7158 |
| Medium Payload | BEVE ZeroCopy | Marshal | 7012 | 3 | 0 |
| Medium Payload | BEVE | Marshal | 12378 | 20487 | 1 |
| Medium Payload | Sonic | Marshal | 16112 | 22189 | 3 |
| Medium Payload | CBOR | Marshal | 24111 | 21777 | 1 |
| Medium Payload | MessagePack | Marshal | 36816 | 65783 | 22 |
| Medium Payload | JSON | Marshal | 45122 | 21988 | 8 |
| Medium Payload | BEVE | Unmarshal | 23879 | 27486 | 59 |
| Medium Payload | Sonic | Unmarshal | 39440 | 57752 | 74 |
| Medium Payload | MessagePack | Unmarshal | 58883 | 37488 | 698 |
| Medium Payload | CBOR | Unmarshal | 62207 | 23608 | 486 |
| Medium Payload | JSON | Unmarshal | 244082 | 57816 | 768 |
| Small Struct | BEVE | Marshal | 819 | 704 | 1 |
| Small Struct | BEVE ZeroCopy | Marshal | 850 | 0 | 0 |
| Small Struct | CBOR | Marshal | 1303 | 1024 | 1 |
| Small Struct | Sonic | Marshal | 1427 | 1843 | 2 |
| Small Struct | MessagePack | Marshal | 2860 | 4104 | 8 |
| Small Struct | JSON | Marshal | 4077 | 1792 | 1 |
| Small Struct | BEVE | Unmarshal | 861 | 760 | 4 |
| Small Struct | MessagePack | Unmarshal | 1055 | 256 | 7 |
| Small Struct | Sonic | Unmarshal | 4428 | 7783 | 10 |
| Small Struct | JSON | Unmarshal | 5624 | 1224 | 25 |
| Small Struct | CBOR | Unmarshal | 6415 | 3215 | 69 |
