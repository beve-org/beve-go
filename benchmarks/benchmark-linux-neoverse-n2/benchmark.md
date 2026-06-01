# Neoverse-N2 — Linux

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE ZeroCopy | Marshal | 68040 | 65 | 0 |
| Large Payload | BEVE | Marshal | 99966 | 180329 | 1 |
| Large Payload | CBOR | Marshal | 184859 | 196778 | 1 |
| Large Payload | MessagePack | Marshal | 274228 | 526799 | 115 |
| Large Payload | Sonic | Marshal | 309388 | 223265 | 3 |
| Large Payload | JSON | Marshal | 393541 | 213395 | 8 |
| Large Payload | BEVE | Unmarshal | 225185 | 276942 | 416 |
| Large Payload | Sonic | Unmarshal | 268825 | 351489 | 213 |
| Large Payload | MessagePack | Unmarshal | 530845 | 365939 | 6695 |
| Large Payload | CBOR | Unmarshal | 678388 | 333322 | 6797 |
| Large Payload | JSON | Unmarshal | 2032689 | 549988 | 7281 |
| Medium Payload | BEVE ZeroCopy | Marshal | 7522 | 6 | 0 |
| Medium Payload | BEVE | Marshal | 9057 | 16396 | 1 |
| Medium Payload | CBOR | Marshal | 16839 | 18453 | 1 |
| Medium Payload | Sonic | Marshal | 26676 | 18909 | 3 |
| Medium Payload | MessagePack | Marshal | 30537 | 65782 | 22 |
| Medium Payload | JSON | Marshal | 39834 | 21994 | 8 |
| Medium Payload | BEVE | Unmarshal | 20758 | 23805 | 59 |
| Medium Payload | Sonic | Unmarshal | 30926 | 44664 | 33 |
| Medium Payload | MessagePack | Unmarshal | 52177 | 35663 | 660 |
| Medium Payload | CBOR | Unmarshal | 56024 | 25784 | 529 |
| Medium Payload | JSON | Unmarshal | 215429 | 61896 | 802 |
| Small Struct | BEVE ZeroCopy | Marshal | 298 | 0 | 0 |
| Small Struct | Sonic | Marshal | 778 | 414 | 2 |
| Small Struct | BEVE | Marshal | 897 | 1536 | 1 |
| Small Struct | MessagePack | Marshal | 992 | 1032 | 6 |
| Small Struct | CBOR | Marshal | 2000 | 2304 | 1 |
| Small Struct | JSON | Marshal | 3154 | 1792 | 1 |
| Small Struct | Sonic | Unmarshal | 1068 | 1010 | 6 |
| Small Struct | BEVE | Unmarshal | 1208 | 1592 | 4 |
| Small Struct | MessagePack | Unmarshal | 5333 | 4608 | 96 |
| Small Struct | CBOR | Unmarshal | 5386 | 3144 | 67 |
| Small Struct | JSON | Unmarshal | 25012 | 7944 | 114 |
