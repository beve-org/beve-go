# Unknown CPU — Windows

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE ZeroCopy | Marshal | 84101 | 259 | 2 |
| Large Payload | BEVE | Marshal | 123936 | 197195 | 3 |
| Large Payload | Sonic | Marshal | 190659 | 228446 | 4 |
| Large Payload | CBOR | Marshal | 219817 | 198398 | 2 |
| Large Payload | MessagePack | Marshal | 285691 | 526764 | 115 |
| Large Payload | JSON | Marshal | 479789 | 215350 | 9 |
| Large Payload | BEVE | Unmarshal | 280364 | 284772 | 418 |
| Large Payload | Sonic | Unmarshal | 425416 | 539603 | 590 |
| Large Payload | MessagePack | Unmarshal | 696188 | 366519 | 6713 |
| Large Payload | CBOR | Unmarshal | 840449 | 309482 | 6293 |
| Large Payload | JSON | Unmarshal | 2463851 | 531629 | 6941 |
| Medium Payload | BEVE ZeroCopy | Marshal | 9346 | 138 | 2 |
| Medium Payload | BEVE | Marshal | 12973 | 18565 | 3 |
| Medium Payload | Sonic | Marshal | 19913 | 27814 | 4 |
| Medium Payload | CBOR | Marshal | 21826 | 19167 | 2 |
| Medium Payload | MessagePack | Marshal | 34408 | 65827 | 22 |
| Medium Payload | JSON | Marshal | 42723 | 19395 | 9 |
| Medium Payload | BEVE | Unmarshal | 29221 | 30364 | 59 |
| Medium Payload | Sonic | Unmarshal | 45049 | 56230 | 73 |
| Medium Payload | MessagePack | Unmarshal | 60676 | 30555 | 557 |
| Medium Payload | CBOR | Unmarshal | 92731 | 34952 | 723 |
| Medium Payload | JSON | Unmarshal | 277943 | 63192 | 805 |
| Small Struct | BEVE ZeroCopy | Marshal | 642 | 290 | 2 |
| Small Struct | CBOR | Marshal | 721 | 528 | 2 |
| Small Struct | BEVE | Marshal | 1253 | 1056 | 3 |
| Small Struct | Sonic | Marshal | 1853 | 2236 | 3 |
| Small Struct | MessagePack | Marshal | 3240 | 4224 | 8 |
| Small Struct | JSON | Marshal | 5705 | 2834 | 2 |
| Small Struct | MessagePack | Unmarshal | 1950 | 784 | 19 |
| Small Struct | CBOR | Unmarshal | 2265 | 704 | 18 |
| Small Struct | BEVE | Unmarshal | 3038 | 3512 | 4 |
| Small Struct | Sonic | Unmarshal | 4016 | 4411 | 9 |
| Small Struct | JSON | Unmarshal | 27940 | 7656 | 105 |
