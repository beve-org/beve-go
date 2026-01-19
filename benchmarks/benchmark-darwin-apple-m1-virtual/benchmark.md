# Apple M1 (Virtual) — Darwin

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE ZeroCopy | Marshal | 72884 | 26 | 0 |
| Large Payload | BEVE | Marshal | 138497 | 188469 | 1 |
| Large Payload | CBOR | Marshal | 199433 | 188560 | 1 |
| Large Payload | MessagePack | Marshal | 369384 | 526751 | 115 |
| Large Payload | JSON | Marshal | 489050 | 213279 | 8 |
| Large Payload | Sonic | Marshal | 632256 | 206421 | 3 |
| Large Payload | BEVE | Unmarshal | 296348 | 261454 | 415 |
| Large Payload | Sonic | Unmarshal | 412267 | 353480 | 213 |
| Large Payload | MessagePack | Unmarshal | 638740 | 352377 | 6431 |
| Large Payload | CBOR | Unmarshal | 844278 | 315803 | 6429 |
| Large Payload | JSON | Unmarshal | 2533356 | 546757 | 7212 |
| Medium Payload | BEVE ZeroCopy | Marshal | 7967 | 3 | 0 |
| Medium Payload | BEVE | Marshal | 13074 | 16390 | 1 |
| Medium Payload | CBOR | Marshal | 28337 | 20490 | 1 |
| Medium Payload | MessagePack | Marshal | 30745 | 33005 | 21 |
| Medium Payload | JSON | Marshal | 44502 | 21990 | 8 |
| Medium Payload | Sonic | Marshal | 53636 | 22044 | 3 |
| Medium Payload | BEVE | Unmarshal | 28340 | 32381 | 59 |
| Medium Payload | Sonic | Unmarshal | 49532 | 42777 | 33 |
| Medium Payload | MessagePack | Unmarshal | 54041 | 27948 | 502 |
| Medium Payload | CBOR | Unmarshal | 62019 | 25736 | 530 |
| Medium Payload | JSON | Unmarshal | 229873 | 51016 | 659 |
| Small Struct | CBOR | Marshal | 428 | 352 | 1 |
| Small Struct | BEVE ZeroCopy | Marshal | 526 | 0 | 0 |
| Small Struct | MessagePack | Marshal | 1454 | 2056 | 7 |
| Small Struct | BEVE | Marshal | 1911 | 2688 | 1 |
| Small Struct | JSON | Marshal | 2178 | 1280 | 1 |
| Small Struct | Sonic | Marshal | 2690 | 1303 | 2 |
| Small Struct | BEVE | Unmarshal | 1158 | 1720 | 4 |
| Small Struct | Sonic | Unmarshal | 1815 | 1233 | 6 |
| Small Struct | MessagePack | Unmarshal | 4642 | 3584 | 76 |
| Small Struct | CBOR | Unmarshal | 5561 | 3176 | 68 |
| Small Struct | JSON | Unmarshal | 12041 | 3920 | 59 |
