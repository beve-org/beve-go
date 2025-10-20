# Neoverse-N2 — Linux

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE ZeroCopy | Marshal | 67286 | 65 | 0 |
| Large Payload | BEVE | Marshal | 100741 | 180316 | 1 |
| Large Payload | CBOR | Marshal | 178981 | 188714 | 1 |
| Large Payload | MessagePack | Marshal | 287030 | 526808 | 115 |
| Large Payload | Sonic | Marshal | 316541 | 231099 | 3 |
| Large Payload | JSON | Marshal | 389565 | 213396 | 8 |
| Large Payload | BEVE | Unmarshal | 231499 | 281775 | 418 |
| Large Payload | Sonic | Unmarshal | 293022 | 405761 | 211 |
| Large Payload | MessagePack | Unmarshal | 535162 | 364162 | 6665 |
| Large Payload | CBOR | Unmarshal | 687845 | 341946 | 6952 |
| Large Payload | JSON | Unmarshal | 2014155 | 546237 | 7195 |
| Medium Payload | BEVE ZeroCopy | Marshal | 6454 | 6 | 0 |
| Medium Payload | BEVE | Marshal | 9844 | 18437 | 1 |
| Medium Payload | CBOR | Marshal | 19225 | 20502 | 1 |
| Medium Payload | MessagePack | Marshal | 32677 | 65782 | 22 |
| Medium Payload | Sonic | Marshal | 36061 | 27947 | 3 |
| Medium Payload | JSON | Marshal | 37265 | 20707 | 8 |
| Medium Payload | BEVE | Unmarshal | 25017 | 33759 | 59 |
| Medium Payload | Sonic | Unmarshal | 31435 | 44006 | 33 |
| Medium Payload | MessagePack | Unmarshal | 49560 | 33087 | 608 |
| Medium Payload | CBOR | Unmarshal | 65364 | 32360 | 663 |
| Medium Payload | JSON | Unmarshal | 225741 | 64328 | 841 |
| Small Struct | BEVE ZeroCopy | Marshal | 288 | 0 | 0 |
| Small Struct | CBOR | Marshal | 978 | 896 | 1 |
| Small Struct | BEVE | Marshal | 1088 | 2048 | 1 |
| Small Struct | JSON | Marshal | 1257 | 640 | 1 |
| Small Struct | Sonic | Marshal | 1730 | 1196 | 2 |
| Small Struct | MessagePack | Marshal | 2406 | 4104 | 8 |
| Small Struct | BEVE | Unmarshal | 898 | 824 | 4 |
| Small Struct | CBOR | Unmarshal | 3062 | 1487 | 34 |
| Small Struct | MessagePack | Unmarshal | 3296 | 2464 | 53 |
| Small Struct | Sonic | Unmarshal | 3308 | 5514 | 6 |
| Small Struct | JSON | Unmarshal | 19807 | 7208 | 91 |
