# Neoverse-N2 — Linux

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE ZeroCopy | Marshal | 69503 | 65 | 0 |
| Large Payload | BEVE | Marshal | 111257 | 196730 | 1 |
| Large Payload | CBOR | Marshal | 182479 | 188768 | 1 |
| Large Payload | MessagePack | Marshal | 269534 | 526803 | 115 |
| Large Payload | Sonic | Marshal | 314475 | 225098 | 3 |
| Large Payload | JSON | Marshal | 402074 | 221615 | 8 |
| Large Payload | BEVE | Unmarshal | 229837 | 285679 | 418 |
| Large Payload | Sonic | Unmarshal | 296696 | 409784 | 211 |
| Large Payload | MessagePack | Unmarshal | 504081 | 334615 | 6078 |
| Large Payload | CBOR | Unmarshal | 646052 | 310427 | 6326 |
| Large Payload | JSON | Unmarshal | 2067255 | 569469 | 7407 |
| Medium Payload | BEVE ZeroCopy | Marshal | 7719 | 3 | 0 |
| Medium Payload | BEVE | Marshal | 9590 | 18441 | 1 |
| Medium Payload | CBOR | Marshal | 16479 | 18447 | 1 |
| Medium Payload | Sonic | Marshal | 27501 | 19417 | 3 |
| Medium Payload | MessagePack | Marshal | 30860 | 65782 | 22 |
| Medium Payload | JSON | Marshal | 42462 | 24807 | 8 |
| Medium Payload | BEVE | Unmarshal | 22803 | 28510 | 59 |
| Medium Payload | Sonic | Unmarshal | 28082 | 37207 | 33 |
| Medium Payload | MessagePack | Unmarshal | 53907 | 37536 | 696 |
| Medium Payload | CBOR | Unmarshal | 64535 | 31720 | 651 |
| Medium Payload | JSON | Unmarshal | 240437 | 72248 | 911 |
| Small Struct | BEVE ZeroCopy | Marshal | 455 | 0 | 0 |
| Small Struct | JSON | Marshal | 944 | 416 | 1 |
| Small Struct | CBOR | Marshal | 976 | 896 | 1 |
| Small Struct | BEVE | Marshal | 977 | 1536 | 1 |
| Small Struct | MessagePack | Marshal | 2403 | 4104 | 8 |
| Small Struct | Sonic | Marshal | 2818 | 2113 | 2 |
| Small Struct | MessagePack | Unmarshal | 1264 | 448 | 12 |
| Small Struct | BEVE | Unmarshal | 1678 | 3000 | 4 |
| Small Struct | Sonic | Unmarshal | 2872 | 4513 | 6 |
| Small Struct | CBOR | Unmarshal | 7154 | 4360 | 93 |
| Small Struct | JSON | Unmarshal | 23539 | 7720 | 107 |
