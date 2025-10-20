# Unknown CPU — Windows

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE ZeroCopy | Marshal | 80778 | 79 | 0 |
| Large Payload | BEVE | Marshal | 117110 | 188464 | 1 |
| Large Payload | Sonic | Marshal | 165912 | 215223 | 3 |
| Large Payload | CBOR | Marshal | 217180 | 196682 | 1 |
| Large Payload | MessagePack | Marshal | 333155 | 526711 | 115 |
| Large Payload | JSON | Marshal | 486861 | 213267 | 8 |
| Large Payload | BEVE | Unmarshal | 330625 | 263656 | 418 |
| Large Payload | Sonic | Unmarshal | 442076 | 547131 | 586 |
| Large Payload | MessagePack | Unmarshal | 674689 | 340447 | 6189 |
| Large Payload | CBOR | Unmarshal | 872604 | 314505 | 6403 |
| Large Payload | JSON | Unmarshal | 2721366 | 557435 | 7297 |
| Medium Payload | BEVE ZeroCopy | Marshal | 8491 | 5 | 0 |
| Medium Payload | Sonic | Marshal | 15224 | 18760 | 3 |
| Medium Payload | BEVE | Marshal | 15881 | 24584 | 1 |
| Medium Payload | CBOR | Marshal | 22009 | 18443 | 1 |
| Medium Payload | MessagePack | Marshal | 35326 | 65772 | 22 |
| Medium Payload | JSON | Marshal | 54781 | 24806 | 8 |
| Medium Payload | BEVE | Unmarshal | 28993 | 29691 | 58 |
| Medium Payload | Sonic | Unmarshal | 47353 | 56032 | 67 |
| Medium Payload | MessagePack | Unmarshal | 76997 | 41533 | 782 |
| Medium Payload | CBOR | Unmarshal | 79148 | 27816 | 573 |
| Medium Payload | JSON | Unmarshal | 266280 | 52872 | 731 |
| Small Struct | Sonic | Marshal | 582 | 485 | 2 |
| Small Struct | BEVE | Marshal | 688 | 576 | 1 |
| Small Struct | BEVE ZeroCopy | Marshal | 751 | 0 | 0 |
| Small Struct | CBOR | Marshal | 1372 | 1280 | 1 |
| Small Struct | JSON | Marshal | 2538 | 1152 | 1 |
| Small Struct | MessagePack | Marshal | 5400 | 8200 | 9 |
| Small Struct | Sonic | Unmarshal | 999 | 800 | 6 |
| Small Struct | BEVE | Unmarshal | 1828 | 2360 | 4 |
| Small Struct | CBOR | Unmarshal | 5417 | 2312 | 51 |
| Small Struct | MessagePack | Unmarshal | 7592 | 4832 | 103 |
| Small Struct | JSON | Unmarshal | 22095 | 4616 | 81 |
