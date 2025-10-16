# AMD EPYC 7763 64-Core Processor — Linux

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE ZeroCopy | Marshal | 81493 | 233 | 2 |
| Large Payload | BEVE | Marshal | 120747 | 196905 | 3 |
| Large Payload | Sonic | Marshal | 159237 | 208881 | 4 |
| Large Payload | CBOR | Marshal | 203414 | 189182 | 2 |
| Large Payload | MessagePack | Marshal | 291611 | 526834 | 115 |
| Large Payload | JSON | Marshal | 419425 | 205335 | 9 |
| Large Payload | BEVE | Unmarshal | 230925 | 269055 | 419 |
| Large Payload | Sonic | Unmarshal | 341857 | 525139 | 559 |
| Large Payload | MessagePack | Unmarshal | 571717 | 349079 | 6355 |
| Large Payload | CBOR | Unmarshal | 709310 | 310777 | 6333 |
| Large Payload | JSON | Unmarshal | 2497613 | 602410 | 7797 |
| Medium Payload | BEVE ZeroCopy | Marshal | 8765 | 134 | 2 |
| Medium Payload | Sonic | Marshal | 13976 | 19152 | 4 |
| Medium Payload | BEVE | Marshal | 14574 | 19235 | 3 |
| Medium Payload | CBOR | Marshal | 22652 | 21860 | 2 |
| Medium Payload | JSON | Marshal | 32683 | 16725 | 9 |
| Medium Payload | MessagePack | Marshal | 35973 | 65838 | 22 |
| Medium Payload | BEVE | Unmarshal | 23989 | 27774 | 59 |
| Medium Payload | Sonic | Unmarshal | 39305 | 60281 | 75 |
| Medium Payload | MessagePack | Unmarshal | 52711 | 31903 | 585 |
| Medium Payload | CBOR | Unmarshal | 68284 | 29192 | 599 |
| Medium Payload | JSON | Unmarshal | 197961 | 48968 | 625 |
| Small Struct | BEVE ZeroCopy | Marshal | 742 | 289 | 2 |
| Small Struct | Sonic | Marshal | 811 | 985 | 3 |
| Small Struct | BEVE | Marshal | 878 | 1313 | 3 |
| Small Struct | CBOR | Marshal | 2207 | 2193 | 2 |
| Small Struct | MessagePack | Marshal | 4370 | 8321 | 9 |
| Small Struct | JSON | Marshal | 5198 | 2835 | 2 |
| Small Struct | BEVE | Unmarshal | 943 | 952 | 4 |
| Small Struct | MessagePack | Unmarshal | 1732 | 832 | 20 |
| Small Struct | Sonic | Unmarshal | 4189 | 7399 | 10 |
| Small Struct | CBOR | Unmarshal | 4740 | 2416 | 52 |
| Small Struct | JSON | Unmarshal | 15829 | 4168 | 67 |
