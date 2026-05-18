# Neoverse-N2 — Linux

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE ZeroCopy | Marshal | 70234 | 52 | 0 |
| Large Payload | BEVE | Marshal | 106237 | 180397 | 1 |
| Large Payload | CBOR | Marshal | 188810 | 188769 | 1 |
| Large Payload | MessagePack | Marshal | 293026 | 526811 | 115 |
| Large Payload | Sonic | Marshal | 326088 | 233978 | 3 |
| Large Payload | JSON | Marshal | 378447 | 205203 | 8 |
| Large Payload | BEVE | Unmarshal | 229362 | 274605 | 418 |
| Large Payload | Sonic | Unmarshal | 298194 | 407962 | 209 |
| Large Payload | MessagePack | Unmarshal | 508137 | 337799 | 6124 |
| Large Payload | CBOR | Unmarshal | 615001 | 287402 | 5863 |
| Large Payload | JSON | Unmarshal | 1982936 | 540468 | 6995 |
| Medium Payload | BEVE ZeroCopy | Marshal | 7673 | 5 | 0 |
| Medium Payload | BEVE | Marshal | 11980 | 24588 | 1 |
| Medium Payload | CBOR | Marshal | 20182 | 21776 | 1 |
| Medium Payload | Sonic | Marshal | 31814 | 25062 | 3 |
| Medium Payload | MessagePack | Marshal | 32463 | 65782 | 22 |
| Medium Payload | JSON | Marshal | 35765 | 19302 | 8 |
| Medium Payload | BEVE | Unmarshal | 23107 | 30110 | 59 |
| Medium Payload | Sonic | Unmarshal | 29245 | 38462 | 33 |
| Medium Payload | MessagePack | Unmarshal | 44224 | 27309 | 491 |
| Medium Payload | CBOR | Unmarshal | 70296 | 34872 | 714 |
| Medium Payload | JSON | Unmarshal | 193042 | 53112 | 703 |
| Small Struct | BEVE | Marshal | 320 | 288 | 1 |
| Small Struct | BEVE ZeroCopy | Marshal | 722 | 0 | 0 |
| Small Struct | CBOR | Marshal | 830 | 704 | 1 |
| Small Struct | MessagePack | Marshal | 1533 | 2056 | 7 |
| Small Struct | Sonic | Marshal | 2408 | 1864 | 2 |
| Small Struct | JSON | Marshal | 3905 | 2304 | 1 |
| Small Struct | BEVE | Unmarshal | 1081 | 1336 | 4 |
| Small Struct | Sonic | Unmarshal | 2101 | 2918 | 6 |
| Small Struct | MessagePack | Unmarshal | 4508 | 3680 | 79 |
| Small Struct | CBOR | Unmarshal | 5628 | 3144 | 67 |
| Small Struct | JSON | Unmarshal | 21104 | 7368 | 96 |
