# Neoverse-N2 — Linux

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE ZeroCopy | Marshal | 70379 | 233 | 2 |
| Large Payload | BEVE | Marshal | 108806 | 181953 | 3 |
| Large Payload | CBOR | Marshal | 187732 | 190670 | 2 |
| Large Payload | MessagePack | Marshal | 278377 | 526862 | 115 |
| Large Payload | Sonic | Marshal | 310623 | 228472 | 4 |
| Large Payload | JSON | Marshal | 403222 | 224779 | 9 |
| Large Payload | BEVE | Unmarshal | 233922 | 278413 | 418 |
| Large Payload | Sonic | Unmarshal | 296076 | 405315 | 211 |
| Large Payload | MessagePack | Unmarshal | 523487 | 354350 | 6457 |
| Large Payload | CBOR | Unmarshal | 691184 | 333642 | 6802 |
| Large Payload | JSON | Unmarshal | 1886425 | 493859 | 6534 |
| Medium Payload | BEVE ZeroCopy | Marshal | 6622 | 141 | 2 |
| Medium Payload | BEVE | Marshal | 9762 | 16532 | 3 |
| Medium Payload | CBOR | Marshal | 19377 | 20623 | 2 |
| Medium Payload | MessagePack | Marshal | 30531 | 65838 | 22 |
| Medium Payload | Sonic | Marshal | 31403 | 25032 | 4 |
| Medium Payload | JSON | Marshal | 38579 | 22085 | 9 |
| Medium Payload | BEVE | Unmarshal | 21963 | 27422 | 59 |
| Medium Payload | Sonic | Unmarshal | 30078 | 39407 | 33 |
| Medium Payload | MessagePack | Unmarshal | 43826 | 27037 | 484 |
| Medium Payload | CBOR | Unmarshal | 61541 | 29320 | 602 |
| Medium Payload | JSON | Unmarshal | 237079 | 68216 | 892 |
| Small Struct | BEVE ZeroCopy | Marshal | 948 | 289 | 2 |
| Small Struct | MessagePack | Marshal | 965 | 1152 | 6 |
| Small Struct | BEVE | Marshal | 1172 | 1697 | 3 |
| Small Struct | CBOR | Marshal | 1286 | 1297 | 2 |
| Small Struct | JSON | Marshal | 2598 | 1552 | 2 |
| Small Struct | Sonic | Marshal | 4084 | 3318 | 3 |
| Small Struct | BEVE | Unmarshal | 1314 | 2104 | 4 |
| Small Struct | Sonic | Unmarshal | 1884 | 2607 | 6 |
| Small Struct | MessagePack | Unmarshal | 3706 | 2912 | 63 |
| Small Struct | CBOR | Unmarshal | 5132 | 2888 | 63 |
| Small Struct | JSON | Unmarshal | 23637 | 7720 | 107 |
