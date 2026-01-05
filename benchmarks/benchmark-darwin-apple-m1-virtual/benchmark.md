# Apple M1 (Virtual) — Darwin

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE ZeroCopy | Marshal | 66984 | 26 | 0 |
| Large Payload | BEVE | Marshal | 123556 | 196675 | 1 |
| Large Payload | CBOR | Marshal | 231988 | 196755 | 1 |
| Large Payload | MessagePack | Marshal | 306778 | 526754 | 115 |
| Large Payload | JSON | Marshal | 545373 | 221497 | 8 |
| Large Payload | Sonic | Marshal | 565709 | 214141 | 3 |
| Large Payload | BEVE | Unmarshal | 304162 | 284754 | 419 |
| Large Payload | Sonic | Unmarshal | 377977 | 350641 | 207 |
| Large Payload | MessagePack | Unmarshal | 641813 | 363882 | 6662 |
| Large Payload | CBOR | Unmarshal | 773253 | 316060 | 6444 |
| Large Payload | JSON | Unmarshal | 2376814 | 527620 | 6843 |
| Medium Payload | BEVE ZeroCopy | Marshal | 7033 | 3 | 0 |
| Medium Payload | BEVE | Marshal | 13795 | 20486 | 1 |
| Medium Payload | CBOR | Marshal | 22878 | 24588 | 1 |
| Medium Payload | MessagePack | Marshal | 37972 | 65778 | 22 |
| Medium Payload | Sonic | Marshal | 46922 | 16609 | 3 |
| Medium Payload | JSON | Marshal | 57445 | 21993 | 8 |
| Medium Payload | BEVE | Unmarshal | 25507 | 26940 | 59 |
| Medium Payload | MessagePack | Unmarshal | 35553 | 26668 | 478 |
| Medium Payload | Sonic | Unmarshal | 37499 | 42661 | 33 |
| Medium Payload | CBOR | Unmarshal | 61186 | 35224 | 724 |
| Medium Payload | JSON | Unmarshal | 232986 | 57944 | 732 |
| Small Struct | BEVE ZeroCopy | Marshal | 529 | 0 | 0 |
| Small Struct | BEVE | Marshal | 1004 | 1408 | 1 |
| Small Struct | CBOR | Marshal | 1909 | 1792 | 1 |
| Small Struct | MessagePack | Marshal | 2474 | 4104 | 8 |
| Small Struct | Sonic | Marshal | 3185 | 1310 | 2 |
| Small Struct | JSON | Marshal | 3689 | 2048 | 1 |
| Small Struct | BEVE | Unmarshal | 1140 | 824 | 4 |
| Small Struct | Sonic | Unmarshal | 1629 | 961 | 6 |
| Small Struct | MessagePack | Unmarshal | 1864 | 832 | 20 |
| Small Struct | CBOR | Unmarshal | 6943 | 3888 | 82 |
| Small Struct | JSON | Unmarshal | 16236 | 3976 | 61 |
