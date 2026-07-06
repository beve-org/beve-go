# Unknown CPU — Windows

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE ZeroCopy | Marshal | 78845 | 52 | 0 |
| Large Payload | BEVE | Marshal | 119319 | 196669 | 1 |
| Large Payload | Sonic | Marshal | 161336 | 207072 | 3 |
| Large Payload | CBOR | Marshal | 239258 | 204932 | 1 |
| Large Payload | MessagePack | Marshal | 287444 | 526706 | 115 |
| Large Payload | JSON | Marshal | 499084 | 213266 | 8 |
| Large Payload | BEVE | Unmarshal | 273116 | 267751 | 418 |
| Large Payload | Sonic | Unmarshal | 434123 | 538728 | 576 |
| Large Payload | MessagePack | Unmarshal | 655117 | 337426 | 6112 |
| Large Payload | CBOR | Unmarshal | 843667 | 309177 | 6304 |
| Large Payload | JSON | Unmarshal | 2591411 | 535781 | 7014 |
| Medium Payload | BEVE ZeroCopy | Marshal | 8202 | 5 | 0 |
| Medium Payload | BEVE | Marshal | 13021 | 20483 | 1 |
| Medium Payload | Sonic | Marshal | 15937 | 20721 | 3 |
| Medium Payload | CBOR | Marshal | 23436 | 20492 | 1 |
| Medium Payload | MessagePack | Marshal | 36809 | 65771 | 22 |
| Medium Payload | JSON | Marshal | 52153 | 24805 | 8 |
| Medium Payload | BEVE | Unmarshal | 30828 | 33756 | 59 |
| Medium Payload | Sonic | Unmarshal | 54692 | 67832 | 80 |
| Medium Payload | MessagePack | Unmarshal | 77500 | 43110 | 804 |
| Medium Payload | CBOR | Unmarshal | 89355 | 33160 | 682 |
| Medium Payload | JSON | Unmarshal | 212728 | 42392 | 585 |
| Small Struct | BEVE ZeroCopy | Marshal | 401 | 0 | 0 |
| Small Struct | BEVE | Marshal | 607 | 256 | 1 |
| Small Struct | CBOR | Marshal | 780 | 576 | 1 |
| Small Struct | JSON | Marshal | 1877 | 704 | 1 |
| Small Struct | Sonic | Marshal | 2525 | 3138 | 2 |
| Small Struct | MessagePack | Marshal | 3244 | 4104 | 8 |
| Small Struct | BEVE | Unmarshal | 1114 | 1208 | 4 |
| Small Struct | MessagePack | Unmarshal | 4177 | 2464 | 53 |
| Small Struct | Sonic | Unmarshal | 5865 | 7750 | 10 |
| Small Struct | CBOR | Unmarshal | 10301 | 5128 | 105 |
| Small Struct | JSON | Unmarshal | 24501 | 4768 | 86 |
