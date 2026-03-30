# Neoverse-N2 — Linux

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE ZeroCopy | Marshal | 69449 | 65 | 0 |
| Large Payload | BEVE | Marshal | 105915 | 180357 | 1 |
| Large Payload | CBOR | Marshal | 193150 | 204977 | 1 |
| Large Payload | MessagePack | Marshal | 289062 | 526810 | 115 |
| Large Payload | Sonic | Marshal | 313398 | 218119 | 3 |
| Large Payload | JSON | Marshal | 414280 | 221640 | 8 |
| Large Payload | BEVE | Unmarshal | 230098 | 282767 | 417 |
| Large Payload | Sonic | Unmarshal | 279453 | 379307 | 211 |
| Large Payload | MessagePack | Unmarshal | 515958 | 350633 | 6381 |
| Large Payload | CBOR | Unmarshal | 663371 | 323482 | 6579 |
| Large Payload | JSON | Unmarshal | 1871880 | 494867 | 6508 |
| Medium Payload | BEVE ZeroCopy | Marshal | 6761 | 6 | 0 |
| Medium Payload | BEVE | Marshal | 10064 | 19077 | 1 |
| Medium Payload | CBOR | Marshal | 18320 | 19090 | 1 |
| Medium Payload | MessagePack | Marshal | 29987 | 65782 | 22 |
| Medium Payload | Sonic | Marshal | 30795 | 22104 | 3 |
| Medium Payload | JSON | Marshal | 39301 | 21991 | 8 |
| Medium Payload | BEVE | Unmarshal | 22794 | 30494 | 59 |
| Medium Payload | Sonic | Unmarshal | 34446 | 51117 | 33 |
| Medium Payload | MessagePack | Unmarshal | 60109 | 44097 | 828 |
| Medium Payload | CBOR | Unmarshal | 66270 | 32664 | 673 |
| Medium Payload | JSON | Unmarshal | 197923 | 57960 | 728 |
| Small Struct | BEVE ZeroCopy | Marshal | 696 | 0 | 0 |
| Small Struct | BEVE | Marshal | 1097 | 2049 | 1 |
| Small Struct | Sonic | Marshal | 1148 | 677 | 2 |
| Small Struct | CBOR | Marshal | 1784 | 2048 | 1 |
| Small Struct | JSON | Marshal | 2361 | 1280 | 1 |
| Small Struct | MessagePack | Marshal | 3813 | 8201 | 9 |
| Small Struct | BEVE | Unmarshal | 1514 | 2616 | 4 |
| Small Struct | Sonic | Unmarshal | 2143 | 3320 | 6 |
| Small Struct | MessagePack | Unmarshal | 4500 | 3680 | 79 |
| Small Struct | CBOR | Unmarshal | 6626 | 3976 | 85 |
| Small Struct | JSON | Unmarshal | 25233 | 7944 | 114 |
