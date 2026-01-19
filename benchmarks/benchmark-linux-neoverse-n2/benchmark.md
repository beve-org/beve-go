# Neoverse-N2 — Linux

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE ZeroCopy | Marshal | 65807 | 65 | 0 |
| Large Payload | BEVE | Marshal | 116524 | 196757 | 1 |
| Large Payload | CBOR | Marshal | 179841 | 180731 | 1 |
| Large Payload | MessagePack | Marshal | 274856 | 526802 | 115 |
| Large Payload | Sonic | Marshal | 326597 | 235147 | 3 |
| Large Payload | JSON | Marshal | 410334 | 221666 | 8 |
| Large Payload | BEVE | Unmarshal | 234288 | 292976 | 419 |
| Large Payload | Sonic | Unmarshal | 294899 | 404656 | 209 |
| Large Payload | MessagePack | Unmarshal | 526423 | 349948 | 6380 |
| Large Payload | CBOR | Unmarshal | 651189 | 308266 | 6288 |
| Large Payload | JSON | Unmarshal | 2081204 | 569179 | 7413 |
| Medium Payload | BEVE ZeroCopy | Marshal | 7243 | 3 | 0 |
| Medium Payload | BEVE | Marshal | 10149 | 18437 | 1 |
| Medium Payload | CBOR | Marshal | 19778 | 20502 | 1 |
| Medium Payload | Sonic | Marshal | 26433 | 18980 | 3 |
| Medium Payload | MessagePack | Marshal | 32938 | 65782 | 22 |
| Medium Payload | JSON | Marshal | 36169 | 19302 | 8 |
| Medium Payload | BEVE | Unmarshal | 22795 | 27614 | 59 |
| Medium Payload | Sonic | Unmarshal | 27532 | 34990 | 33 |
| Medium Payload | MessagePack | Unmarshal | 49796 | 32847 | 602 |
| Medium Payload | CBOR | Unmarshal | 65345 | 31288 | 648 |
| Medium Payload | JSON | Unmarshal | 222715 | 63128 | 828 |
| Small Struct | BEVE ZeroCopy | Marshal | 237 | 0 | 0 |
| Small Struct | Sonic | Marshal | 754 | 407 | 2 |
| Small Struct | BEVE | Marshal | 818 | 1408 | 1 |
| Small Struct | CBOR | Marshal | 1296 | 1280 | 1 |
| Small Struct | MessagePack | Marshal | 2570 | 4105 | 8 |
| Small Struct | JSON | Marshal | 3940 | 2304 | 1 |
| Small Struct | BEVE | Unmarshal | 1682 | 3384 | 4 |
| Small Struct | MessagePack | Unmarshal | 2005 | 1176 | 27 |
| Small Struct | Sonic | Unmarshal | 3005 | 5107 | 6 |
| Small Struct | CBOR | Unmarshal | 7410 | 4424 | 95 |
| Small Struct | JSON | Unmarshal | 16204 | 4424 | 75 |
