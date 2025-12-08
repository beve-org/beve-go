# Neoverse-N2 — Linux

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE ZeroCopy | Marshal | 67693 | 65 | 0 |
| Large Payload | BEVE | Marshal | 119407 | 221388 | 1 |
| Large Payload | CBOR | Marshal | 180717 | 188688 | 1 |
| Large Payload | MessagePack | Marshal | 288683 | 526810 | 115 |
| Large Payload | Sonic | Marshal | 318593 | 225749 | 3 |
| Large Payload | JSON | Marshal | 390754 | 213422 | 8 |
| Large Payload | BEVE | Unmarshal | 228760 | 272011 | 418 |
| Large Payload | Sonic | Unmarshal | 277002 | 376234 | 213 |
| Large Payload | MessagePack | Unmarshal | 520945 | 364818 | 6682 |
| Large Payload | CBOR | Unmarshal | 611991 | 291546 | 5948 |
| Large Payload | JSON | Unmarshal | 2050415 | 551716 | 7323 |
| Medium Payload | BEVE ZeroCopy | Marshal | 6894 | 3 | 0 |
| Medium Payload | BEVE | Marshal | 9658 | 18439 | 1 |
| Medium Payload | CBOR | Marshal | 22514 | 24594 | 1 |
| Medium Payload | MessagePack | Marshal | 33372 | 65782 | 22 |
| Medium Payload | Sonic | Marshal | 37756 | 27977 | 3 |
| Medium Payload | JSON | Marshal | 41534 | 24807 | 8 |
| Medium Payload | BEVE | Unmarshal | 20955 | 23197 | 59 |
| Medium Payload | Sonic | Unmarshal | 28264 | 36700 | 33 |
| Medium Payload | MessagePack | Unmarshal | 55386 | 39552 | 736 |
| Medium Payload | CBOR | Unmarshal | 68548 | 34360 | 704 |
| Medium Payload | JSON | Unmarshal | 156760 | 40344 | 540 |
| Small Struct | BEVE ZeroCopy | Marshal | 594 | 0 | 0 |
| Small Struct | CBOR | Marshal | 622 | 384 | 1 |
| Small Struct | MessagePack | Marshal | 752 | 520 | 5 |
| Small Struct | BEVE | Marshal | 950 | 1537 | 1 |
| Small Struct | Sonic | Marshal | 1093 | 670 | 2 |
| Small Struct | JSON | Marshal | 2881 | 1792 | 1 |
| Small Struct | BEVE | Unmarshal | 1538 | 2360 | 4 |
| Small Struct | Sonic | Unmarshal | 2242 | 3407 | 6 |
| Small Struct | JSON | Unmarshal | 3129 | 552 | 15 |
| Small Struct | MessagePack | Unmarshal | 3793 | 2912 | 63 |
| Small Struct | CBOR | Unmarshal | 6649 | 3944 | 84 |
