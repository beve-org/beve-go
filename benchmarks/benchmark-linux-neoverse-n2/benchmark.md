# Neoverse-N2 — Linux

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE ZeroCopy | Marshal | 69499 | 65 | 0 |
| Large Payload | BEVE | Marshal | 107189 | 188550 | 1 |
| Large Payload | CBOR | Marshal | 189284 | 197041 | 1 |
| Large Payload | MessagePack | Marshal | 271371 | 526805 | 115 |
| Large Payload | Sonic | Marshal | 302428 | 209719 | 3 |
| Large Payload | JSON | Marshal | 391936 | 213500 | 8 |
| Large Payload | BEVE | Unmarshal | 217183 | 257540 | 418 |
| Large Payload | Sonic | Unmarshal | 288043 | 397652 | 213 |
| Large Payload | MessagePack | Unmarshal | 506826 | 338712 | 6143 |
| Large Payload | CBOR | Unmarshal | 657145 | 318841 | 6502 |
| Large Payload | JSON | Unmarshal | 1991359 | 540812 | 7067 |
| Medium Payload | BEVE ZeroCopy | Marshal | 6300 | 6 | 0 |
| Medium Payload | BEVE | Marshal | 10978 | 21771 | 1 |
| Medium Payload | CBOR | Marshal | 16121 | 16396 | 1 |
| Medium Payload | Sonic | Marshal | 26068 | 18779 | 3 |
| Medium Payload | MessagePack | Marshal | 29894 | 65782 | 22 |
| Medium Payload | JSON | Marshal | 37992 | 20711 | 8 |
| Medium Payload | BEVE | Unmarshal | 22966 | 30430 | 58 |
| Medium Payload | Sonic | Unmarshal | 31870 | 45095 | 33 |
| Medium Payload | MessagePack | Unmarshal | 55892 | 40160 | 749 |
| Medium Payload | CBOR | Unmarshal | 65569 | 32376 | 664 |
| Medium Payload | JSON | Unmarshal | 190276 | 49848 | 697 |
| Small Struct | BEVE ZeroCopy | Marshal | 494 | 0 | 0 |
| Small Struct | CBOR | Marshal | 775 | 640 | 1 |
| Small Struct | BEVE | Marshal | 1128 | 2305 | 1 |
| Small Struct | Sonic | Marshal | 1173 | 734 | 2 |
| Small Struct | MessagePack | Marshal | 2195 | 4104 | 8 |
| Small Struct | JSON | Marshal | 2792 | 1792 | 1 |
| Small Struct | BEVE | Unmarshal | 1010 | 1208 | 4 |
| Small Struct | CBOR | Unmarshal | 2467 | 1104 | 26 |
| Small Struct | Sonic | Unmarshal | 3215 | 5281 | 6 |
| Small Struct | MessagePack | Unmarshal | 3545 | 2880 | 62 |
| Small Struct | JSON | Unmarshal | 12813 | 3944 | 60 |
