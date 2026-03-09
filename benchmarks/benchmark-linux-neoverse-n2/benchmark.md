# Neoverse-N2 — Linux

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE ZeroCopy | Marshal | 69334 | 52 | 0 |
| Large Payload | BEVE | Marshal | 103874 | 180383 | 1 |
| Large Payload | CBOR | Marshal | 185452 | 196858 | 1 |
| Large Payload | MessagePack | Marshal | 291634 | 526804 | 115 |
| Large Payload | Sonic | Marshal | 325019 | 231064 | 3 |
| Large Payload | JSON | Marshal | 387399 | 213395 | 8 |
| Large Payload | BEVE | Unmarshal | 227137 | 266921 | 417 |
| Large Payload | Sonic | Unmarshal | 297804 | 406296 | 211 |
| Large Payload | MessagePack | Unmarshal | 486569 | 319843 | 5757 |
| Large Payload | CBOR | Unmarshal | 647880 | 310971 | 6351 |
| Large Payload | JSON | Unmarshal | 2108192 | 573027 | 7550 |
| Medium Payload | BEVE ZeroCopy | Marshal | 7595 | 6 | 0 |
| Medium Payload | BEVE | Marshal | 9923 | 16392 | 1 |
| Medium Payload | CBOR | Marshal | 17931 | 19087 | 1 |
| Medium Payload | MessagePack | Marshal | 32989 | 65782 | 22 |
| Medium Payload | Sonic | Marshal | 35375 | 25090 | 3 |
| Medium Payload | JSON | Marshal | 42815 | 24804 | 8 |
| Medium Payload | BEVE | Unmarshal | 25271 | 30911 | 59 |
| Medium Payload | Sonic | Unmarshal | 29667 | 37619 | 33 |
| Medium Payload | MessagePack | Unmarshal | 53886 | 36831 | 681 |
| Medium Payload | CBOR | Unmarshal | 54481 | 24440 | 503 |
| Medium Payload | JSON | Unmarshal | 271861 | 81912 | 1036 |
| Small Struct | BEVE ZeroCopy | Marshal | 589 | 0 | 0 |
| Small Struct | BEVE | Marshal | 1041 | 1792 | 1 |
| Small Struct | CBOR | Marshal | 2105 | 2304 | 1 |
| Small Struct | MessagePack | Marshal | 2444 | 4104 | 8 |
| Small Struct | Sonic | Marshal | 3476 | 2754 | 2 |
| Small Struct | JSON | Marshal | 4851 | 3072 | 1 |
| Small Struct | BEVE | Unmarshal | 1717 | 3384 | 4 |
| Small Struct | MessagePack | Unmarshal | 3047 | 2144 | 47 |
| Small Struct | Sonic | Unmarshal | 3207 | 5002 | 6 |
| Small Struct | CBOR | Unmarshal | 5993 | 3496 | 74 |
| Small Struct | JSON | Unmarshal | 12727 | 3912 | 59 |
