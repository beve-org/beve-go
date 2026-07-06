# Neoverse-N2 — Linux

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE ZeroCopy | Marshal | 69507 | 65 | 0 |
| Large Payload | BEVE | Marshal | 114599 | 196822 | 1 |
| Large Payload | CBOR | Marshal | 184034 | 188715 | 1 |
| Large Payload | MessagePack | Marshal | 281565 | 526805 | 115 |
| Large Payload | Sonic | Marshal | 293375 | 207123 | 3 |
| Large Payload | JSON | Marshal | 410500 | 221562 | 8 |
| Large Payload | BEVE | Unmarshal | 216845 | 252387 | 417 |
| Large Payload | Sonic | Unmarshal | 282225 | 384533 | 209 |
| Large Payload | MessagePack | Unmarshal | 502375 | 335734 | 6092 |
| Large Payload | CBOR | Unmarshal | 687467 | 335338 | 6851 |
| Large Payload | JSON | Unmarshal | 1972644 | 532077 | 6971 |
| Medium Payload | BEVE ZeroCopy | Marshal | 7544 | 5 | 0 |
| Medium Payload | BEVE | Marshal | 10454 | 20489 | 1 |
| Medium Payload | CBOR | Marshal | 16577 | 18450 | 1 |
| Medium Payload | MessagePack | Marshal | 30279 | 65783 | 22 |
| Medium Payload | Sonic | Marshal | 30350 | 22124 | 3 |
| Medium Payload | JSON | Marshal | 39082 | 21991 | 8 |
| Medium Payload | BEVE | Unmarshal | 22124 | 28126 | 58 |
| Medium Payload | Sonic | Unmarshal | 33340 | 48775 | 33 |
| Medium Payload | MessagePack | Unmarshal | 54797 | 39376 | 730 |
| Medium Payload | CBOR | Unmarshal | 63954 | 31272 | 646 |
| Medium Payload | JSON | Unmarshal | 193442 | 55512 | 708 |
| Small Struct | BEVE | Marshal | 235 | 160 | 1 |
| Small Struct | BEVE ZeroCopy | Marshal | 491 | 0 | 0 |
| Small Struct | JSON | Marshal | 2116 | 1152 | 1 |
| Small Struct | CBOR | Marshal | 2267 | 2688 | 1 |
| Small Struct | MessagePack | Marshal | 2372 | 4104 | 8 |
| Small Struct | Sonic | Marshal | 3728 | 2761 | 2 |
| Small Struct | BEVE | Unmarshal | 669 | 440 | 4 |
| Small Struct | MessagePack | Unmarshal | 1636 | 736 | 18 |
| Small Struct | Sonic | Unmarshal | 3305 | 5724 | 6 |
| Small Struct | CBOR | Unmarshal | 5572 | 3248 | 70 |
| Small Struct | JSON | Unmarshal | 23452 | 7688 | 106 |
