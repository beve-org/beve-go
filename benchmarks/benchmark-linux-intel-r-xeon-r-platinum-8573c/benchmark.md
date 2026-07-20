# INTEL(R) XEON(R) PLATINUM 8573C — Linux

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE ZeroCopy | Marshal | 54349 | 26 | 0 |
| Large Payload | BEVE | Marshal | 99856 | 196696 | 1 |
| Large Payload | Sonic | Marshal | 139212 | 217112 | 3 |
| Large Payload | CBOR | Marshal | 160780 | 196788 | 1 |
| Large Payload | MessagePack | Marshal | 252704 | 526774 | 115 |
| Large Payload | JSON | Marshal | 365094 | 229825 | 8 |
| Large Payload | BEVE | Unmarshal | 207681 | 269664 | 419 |
| Large Payload | Sonic | Unmarshal | 324230 | 549997 | 588 |
| Large Payload | MessagePack | Unmarshal | 476829 | 363405 | 6650 |
| Large Payload | CBOR | Unmarshal | 609651 | 329035 | 6712 |
| Large Payload | JSON | Unmarshal | 1886211 | 544674 | 7162 |
| Medium Payload | BEVE ZeroCopy | Marshal | 5987 | 3 | 0 |
| Medium Payload | BEVE | Marshal | 7892 | 14344 | 1 |
| Medium Payload | CBOR | Marshal | 14399 | 16396 | 1 |
| Medium Payload | Sonic | Marshal | 14826 | 25339 | 3 |
| Medium Payload | MessagePack | Marshal | 21416 | 33007 | 21 |
| Medium Payload | JSON | Marshal | 30142 | 19300 | 8 |
| Medium Payload | BEVE | Unmarshal | 23018 | 28478 | 59 |
| Medium Payload | Sonic | Unmarshal | 39598 | 71078 | 79 |
| Medium Payload | MessagePack | Unmarshal | 49698 | 41634 | 780 |
| Medium Payload | CBOR | Unmarshal | 56539 | 30600 | 627 |
| Medium Payload | JSON | Unmarshal | 156375 | 47656 | 607 |
| Small Struct | BEVE ZeroCopy | Marshal | 368 | 0 | 0 |
| Small Struct | BEVE | Marshal | 1077 | 2048 | 1 |
| Small Struct | MessagePack | Marshal | 1278 | 2056 | 7 |
| Small Struct | CBOR | Marshal | 1582 | 2048 | 1 |
| Small Struct | Sonic | Marshal | 1789 | 2789 | 2 |
| Small Struct | JSON | Marshal | 2025 | 1280 | 1 |
| Small Struct | Sonic | Unmarshal | 564 | 485 | 4 |
| Small Struct | BEVE | Unmarshal | 1635 | 3384 | 4 |
| Small Struct | MessagePack | Unmarshal | 2581 | 2144 | 47 |
| Small Struct | CBOR | Unmarshal | 4514 | 2856 | 62 |
| Small Struct | JSON | Unmarshal | 10603 | 3752 | 54 |
