# Unknown CPU — Windows

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE ZeroCopy | Marshal | 85955 | 79 | 0 |
| Large Payload | BEVE | Marshal | 115188 | 188464 | 1 |
| Large Payload | Sonic | Marshal | 169229 | 223540 | 3 |
| Large Payload | CBOR | Marshal | 220868 | 204931 | 1 |
| Large Payload | MessagePack | Marshal | 291536 | 526706 | 115 |
| Large Payload | JSON | Marshal | 440032 | 196879 | 8 |
| Large Payload | BEVE | Unmarshal | 296069 | 277985 | 417 |
| Large Payload | Sonic | Unmarshal | 441440 | 521541 | 560 |
| Large Payload | MessagePack | Unmarshal | 679656 | 340836 | 6191 |
| Large Payload | CBOR | Unmarshal | 837421 | 301721 | 6164 |
| Large Payload | JSON | Unmarshal | 2642525 | 527086 | 6984 |
| Medium Payload | BEVE ZeroCopy | Marshal | 9601 | 5 | 0 |
| Medium Payload | BEVE | Marshal | 15609 | 18437 | 1 |
| Medium Payload | Sonic | Marshal | 18462 | 24901 | 3 |
| Medium Payload | CBOR | Marshal | 22455 | 18440 | 1 |
| Medium Payload | MessagePack | Marshal | 35061 | 65772 | 22 |
| Medium Payload | JSON | Marshal | 53187 | 24806 | 8 |
| Medium Payload | BEVE | Unmarshal | 29681 | 30619 | 59 |
| Medium Payload | Sonic | Unmarshal | 47894 | 56887 | 76 |
| Medium Payload | MessagePack | Unmarshal | 66252 | 32748 | 600 |
| Medium Payload | CBOR | Unmarshal | 97154 | 34520 | 712 |
| Medium Payload | JSON | Unmarshal | 300433 | 61376 | 806 |
| Small Struct | BEVE ZeroCopy | Marshal | 810 | 0 | 0 |
| Small Struct | MessagePack | Marshal | 834 | 520 | 5 |
| Small Struct | BEVE | Marshal | 1115 | 1408 | 1 |
| Small Struct | Sonic | Marshal | 1707 | 1843 | 2 |
| Small Struct | CBOR | Marshal | 2766 | 3072 | 1 |
| Small Struct | JSON | Marshal | 6307 | 3073 | 1 |
| Small Struct | MessagePack | Unmarshal | 1137 | 304 | 9 |
| Small Struct | BEVE | Unmarshal | 1251 | 1464 | 4 |
| Small Struct | Sonic | Unmarshal | 3071 | 3525 | 9 |
| Small Struct | CBOR | Unmarshal | 7101 | 2888 | 63 |
| Small Struct | JSON | Unmarshal | 30111 | 7656 | 105 |
