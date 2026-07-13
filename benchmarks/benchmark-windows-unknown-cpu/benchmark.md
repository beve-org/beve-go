# Unknown CPU — Windows

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE ZeroCopy | Marshal | 80803 | 65 | 0 |
| Large Payload | BEVE | Marshal | 119352 | 188438 | 1 |
| Large Payload | Sonic | Marshal | 172824 | 222728 | 3 |
| Large Payload | CBOR | Marshal | 212794 | 188493 | 1 |
| Large Payload | MessagePack | Marshal | 310401 | 526711 | 115 |
| Large Payload | JSON | Marshal | 477073 | 213268 | 8 |
| Large Payload | BEVE | Unmarshal | 300736 | 260484 | 417 |
| Large Payload | Sonic | Unmarshal | 530285 | 591310 | 601 |
| Large Payload | MessagePack | Unmarshal | 717434 | 319683 | 5762 |
| Large Payload | CBOR | Unmarshal | 925045 | 295721 | 6030 |
| Large Payload | JSON | Unmarshal | 2543690 | 518820 | 6753 |
| Medium Payload | BEVE ZeroCopy | Marshal | 6505 | 5 | 0 |
| Medium Payload | BEVE | Marshal | 14500 | 20486 | 1 |
| Medium Payload | Sonic | Marshal | 17098 | 18905 | 3 |
| Medium Payload | CBOR | Marshal | 26382 | 21775 | 1 |
| Medium Payload | MessagePack | Marshal | 38493 | 65771 | 22 |
| Medium Payload | JSON | Marshal | 52337 | 20716 | 8 |
| Medium Payload | BEVE | Unmarshal | 31062 | 25371 | 59 |
| Medium Payload | Sonic | Unmarshal | 62895 | 69711 | 80 |
| Medium Payload | CBOR | Unmarshal | 75183 | 23480 | 486 |
| Medium Payload | MessagePack | Unmarshal | 76044 | 36765 | 680 |
| Medium Payload | JSON | Unmarshal | 289917 | 55304 | 735 |
| Small Struct | BEVE ZeroCopy | Marshal | 794 | 0 | 0 |
| Small Struct | MessagePack | Marshal | 1213 | 1032 | 6 |
| Small Struct | CBOR | Marshal | 1263 | 1024 | 1 |
| Small Struct | BEVE | Marshal | 1528 | 1536 | 1 |
| Small Struct | Sonic | Marshal | 2094 | 2383 | 2 |
| Small Struct | JSON | Marshal | 4760 | 2304 | 1 |
| Small Struct | BEVE | Unmarshal | 1933 | 1848 | 4 |
| Small Struct | CBOR | Unmarshal | 3161 | 1160 | 27 |
| Small Struct | Sonic | Unmarshal | 3284 | 3509 | 9 |
| Small Struct | MessagePack | Unmarshal | 4624 | 2424 | 52 |
| Small Struct | JSON | Unmarshal | 29770 | 7432 | 98 |
