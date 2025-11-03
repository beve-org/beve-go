# Neoverse-N2 — Linux

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE ZeroCopy | Marshal | 65380 | 65 | 0 |
| Large Payload | BEVE | Marshal | 101820 | 180369 | 1 |
| Large Payload | CBOR | Marshal | 168974 | 172270 | 1 |
| Large Payload | MessagePack | Marshal | 272238 | 526800 | 115 |
| Large Payload | Sonic | Marshal | 306533 | 214497 | 3 |
| Large Payload | JSON | Marshal | 391493 | 213343 | 8 |
| Large Payload | BEVE | Unmarshal | 215895 | 252099 | 417 |
| Large Payload | Sonic | Unmarshal | 297756 | 407688 | 205 |
| Large Payload | MessagePack | Unmarshal | 529721 | 359934 | 6587 |
| Large Payload | CBOR | Unmarshal | 656154 | 319353 | 6502 |
| Large Payload | JSON | Unmarshal | 1911684 | 519420 | 6734 |
| Medium Payload | BEVE ZeroCopy | Marshal | 7679 | 3 | 0 |
| Medium Payload | BEVE | Marshal | 9563 | 16394 | 1 |
| Medium Payload | CBOR | Marshal | 17267 | 18447 | 1 |
| Medium Payload | MessagePack | Marshal | 22596 | 33006 | 21 |
| Medium Payload | Sonic | Marshal | 28068 | 19547 | 3 |
| Medium Payload | JSON | Marshal | 39619 | 21987 | 8 |
| Medium Payload | BEVE | Unmarshal | 22345 | 26717 | 59 |
| Medium Payload | Sonic | Unmarshal | 29575 | 40159 | 33 |
| Medium Payload | MessagePack | Unmarshal | 57582 | 41344 | 775 |
| Medium Payload | CBOR | Unmarshal | 64988 | 31624 | 653 |
| Medium Payload | JSON | Unmarshal | 164924 | 44312 | 571 |
| Small Struct | BEVE ZeroCopy | Marshal | 392 | 0 | 0 |
| Small Struct | BEVE | Marshal | 800 | 1280 | 1 |
| Small Struct | CBOR | Marshal | 2366 | 2688 | 1 |
| Small Struct | JSON | Marshal | 2499 | 1408 | 1 |
| Small Struct | Sonic | Marshal | 2563 | 1850 | 2 |
| Small Struct | MessagePack | Marshal | 3616 | 8201 | 9 |
| Small Struct | Sonic | Unmarshal | 1634 | 2068 | 6 |
| Small Struct | BEVE | Unmarshal | 1727 | 3384 | 4 |
| Small Struct | MessagePack | Unmarshal | 2880 | 1952 | 43 |
| Small Struct | CBOR | Unmarshal | 5558 | 3176 | 68 |
| Small Struct | JSON | Unmarshal | 8675 | 2248 | 42 |
