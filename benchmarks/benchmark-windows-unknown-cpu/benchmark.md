# Unknown CPU — Windows

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE ZeroCopy | Marshal | 75655 | 65 | 0 |
| Large Payload | BEVE | Marshal | 157505 | 196764 | 1 |
| Large Payload | Sonic | Marshal | 219672 | 211089 | 3 |
| Large Payload | CBOR | Marshal | 248536 | 188619 | 1 |
| Large Payload | MessagePack | Marshal | 395712 | 526756 | 115 |
| Large Payload | JSON | Marshal | 505671 | 213358 | 8 |
| Large Payload | BEVE | Unmarshal | 323274 | 261492 | 418 |
| Large Payload | Sonic | Unmarshal | 527067 | 522210 | 558 |
| Large Payload | MessagePack | Unmarshal | 709897 | 340406 | 6192 |
| Large Payload | CBOR | Unmarshal | 1006294 | 335705 | 6827 |
| Large Payload | JSON | Unmarshal | 2406578 | 522762 | 6792 |
| Medium Payload | BEVE ZeroCopy | Marshal | 7959 | 5 | 0 |
| Medium Payload | BEVE | Marshal | 16354 | 18445 | 1 |
| Medium Payload | Sonic | Marshal | 25675 | 25039 | 3 |
| Medium Payload | CBOR | Marshal | 29039 | 19083 | 1 |
| Medium Payload | JSON | Marshal | 46269 | 18669 | 8 |
| Medium Payload | MessagePack | Marshal | 52021 | 65778 | 22 |
| Medium Payload | BEVE | Unmarshal | 36359 | 28285 | 59 |
| Medium Payload | Sonic | Unmarshal | 64490 | 61302 | 76 |
| Medium Payload | MessagePack | Unmarshal | 73868 | 33967 | 623 |
| Medium Payload | CBOR | Unmarshal | 109540 | 42152 | 864 |
| Medium Payload | JSON | Unmarshal | 265637 | 59752 | 794 |
| Small Struct | BEVE ZeroCopy | Marshal | 952 | 0 | 0 |
| Small Struct | BEVE | Marshal | 970 | 576 | 1 |
| Small Struct | CBOR | Marshal | 1641 | 640 | 1 |
| Small Struct | MessagePack | Marshal | 3256 | 2056 | 7 |
| Small Struct | Sonic | Marshal | 3363 | 2768 | 2 |
| Small Struct | JSON | Marshal | 5848 | 1792 | 1 |
| Small Struct | Sonic | Unmarshal | 1613 | 1312 | 7 |
| Small Struct | BEVE | Unmarshal | 3551 | 3000 | 4 |
| Small Struct | MessagePack | Unmarshal | 4390 | 2080 | 45 |
| Small Struct | CBOR | Unmarshal | 6751 | 2728 | 58 |
| Small Struct | JSON | Unmarshal | 15456 | 2440 | 48 |
