# Intel(R) Xeon(R) Platinum 8370C CPU @ 2.80GHz — Linux

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE ZeroCopy | Marshal | 76018 | 233 | 2 |
| Large Payload | BEVE | Marshal | 113479 | 180677 | 3 |
| Large Payload | Sonic | Marshal | 168201 | 217780 | 4 |
| Large Payload | CBOR | Marshal | 191134 | 188918 | 2 |
| Large Payload | MessagePack | Marshal | 301772 | 526835 | 115 |
| Large Payload | JSON | Marshal | 404771 | 205232 | 9 |
| Large Payload | BEVE | Unmarshal | 235127 | 270912 | 417 |
| Large Payload | Sonic | Unmarshal | 401520 | 570121 | 602 |
| Large Payload | MessagePack | Unmarshal | 559010 | 360623 | 6588 |
| Large Payload | CBOR | Unmarshal | 622911 | 283866 | 5783 |
| Large Payload | JSON | Unmarshal | 1963207 | 498187 | 6633 |
| Medium Payload | BEVE ZeroCopy | Marshal | 7408 | 134 | 2 |
| Medium Payload | BEVE | Marshal | 13903 | 21907 | 3 |
| Medium Payload | Sonic | Marshal | 21874 | 28312 | 4 |
| Medium Payload | CBOR | Marshal | 23386 | 24685 | 2 |
| Medium Payload | MessagePack | Marshal | 25416 | 33064 | 21 |
| Medium Payload | JSON | Marshal | 44990 | 24902 | 9 |
| Medium Payload | BEVE | Unmarshal | 23623 | 26494 | 59 |
| Medium Payload | Sonic | Unmarshal | 42220 | 60911 | 77 |
| Medium Payload | MessagePack | Unmarshal | 50476 | 33231 | 610 |
| Medium Payload | CBOR | Unmarshal | 65410 | 32104 | 664 |
| Medium Payload | JSON | Unmarshal | 212508 | 58680 | 772 |
| Small Struct | BEVE ZeroCopy | Marshal | 550 | 289 | 2 |
| Small Struct | CBOR | Marshal | 1259 | 1297 | 2 |
| Small Struct | BEVE | Marshal | 1386 | 1826 | 3 |
| Small Struct | MessagePack | Marshal | 1532 | 2176 | 7 |
| Small Struct | JSON | Marshal | 1636 | 912 | 2 |
| Small Struct | Sonic | Marshal | 1712 | 2280 | 3 |
| Small Struct | BEVE | Unmarshal | 1713 | 3001 | 4 |
| Small Struct | Sonic | Unmarshal | 2560 | 3782 | 9 |
| Small Struct | MessagePack | Unmarshal | 5026 | 3976 | 84 |
| Small Struct | CBOR | Unmarshal | 7294 | 4424 | 95 |
| Small Struct | JSON | Unmarshal | 12039 | 3784 | 55 |
