# Neoverse-N2 — Linux

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE ZeroCopy | Marshal | 67896 | 286 | 2 |
| Large Payload | BEVE | Marshal | 116101 | 206268 | 3 |
| Large Payload | CBOR | Marshal | 188697 | 199761 | 3 |
| Large Payload | MessagePack | Marshal | 265820 | 526856 | 115 |
| Large Payload | Sonic | Marshal | 292768 | 219654 | 4 |
| Large Payload | JSON | Marshal | 382216 | 214956 | 9 |
| Large Payload | BEVE | Unmarshal | 221942 | 272008 | 417 |
| Large Payload | Sonic | Unmarshal | 278501 | 388509 | 211 |
| Large Payload | MessagePack | Unmarshal | 498996 | 337206 | 6112 |
| Large Payload | CBOR | Unmarshal | 653246 | 316394 | 6448 |
| Large Payload | JSON | Unmarshal | 1896598 | 504427 | 6613 |
| Medium Payload | BEVE ZeroCopy | Marshal | 7128 | 141 | 2 |
| Medium Payload | BEVE | Marshal | 11273 | 20650 | 3 |
| Medium Payload | CBOR | Marshal | 17750 | 19214 | 2 |
| Medium Payload | JSON | Marshal | 30096 | 16692 | 9 |
| Medium Payload | MessagePack | Marshal | 31799 | 65839 | 22 |
| Medium Payload | Sonic | Marshal | 33933 | 25132 | 4 |
| Medium Payload | BEVE | Unmarshal | 23158 | 29470 | 59 |
| Medium Payload | Sonic | Unmarshal | 25039 | 32973 | 33 |
| Medium Payload | MessagePack | Unmarshal | 43290 | 27501 | 496 |
| Medium Payload | CBOR | Unmarshal | 62042 | 30088 | 620 |
| Medium Payload | JSON | Unmarshal | 194637 | 52984 | 707 |
| Small Struct | BEVE ZeroCopy | Marshal | 423 | 289 | 2 |
| Small Struct | CBOR | Marshal | 506 | 352 | 2 |
| Small Struct | BEVE | Marshal | 1693 | 2978 | 3 |
| Small Struct | Sonic | Marshal | 1696 | 1241 | 3 |
| Small Struct | JSON | Marshal | 3432 | 2193 | 2 |
| Small Struct | MessagePack | Marshal | 3869 | 8321 | 9 |
| Small Struct | BEVE | Unmarshal | 1722 | 3384 | 4 |
| Small Struct | JSON | Unmarshal | 2548 | 464 | 12 |
| Small Struct | Sonic | Unmarshal | 3222 | 5282 | 6 |
| Small Struct | CBOR | Unmarshal | 4880 | 2688 | 57 |
| Small Struct | MessagePack | Unmarshal | 5832 | 5161 | 105 |
