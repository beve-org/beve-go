# Unknown CPU — Windows

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE ZeroCopy | Marshal | 79536 | 65 | 0 |
| Large Payload | BEVE | Marshal | 108983 | 180285 | 1 |
| Large Payload | Sonic | Marshal | 166187 | 222703 | 3 |
| Large Payload | CBOR | Marshal | 207863 | 188511 | 1 |
| Large Payload | MessagePack | Marshal | 291139 | 526709 | 115 |
| Large Payload | JSON | Marshal | 511520 | 229650 | 8 |
| Large Payload | BEVE | Unmarshal | 270583 | 265860 | 419 |
| Large Payload | Sonic | Unmarshal | 415537 | 526106 | 575 |
| Large Payload | MessagePack | Unmarshal | 682993 | 359560 | 6571 |
| Large Payload | CBOR | Unmarshal | 846397 | 317017 | 6462 |
| Large Payload | JSON | Unmarshal | 2515653 | 524523 | 6813 |
| Medium Payload | BEVE ZeroCopy | Marshal | 7650 | 5 | 0 |
| Medium Payload | BEVE | Marshal | 13939 | 20486 | 1 |
| Medium Payload | Sonic | Marshal | 19667 | 24965 | 3 |
| Medium Payload | CBOR | Marshal | 22333 | 18445 | 1 |
| Medium Payload | MessagePack | Marshal | 33597 | 65772 | 22 |
| Medium Payload | JSON | Marshal | 43084 | 19308 | 8 |
| Medium Payload | BEVE | Unmarshal | 26957 | 27355 | 59 |
| Medium Payload | Sonic | Unmarshal | 46790 | 57468 | 71 |
| Medium Payload | MessagePack | Unmarshal | 67304 | 35189 | 652 |
| Medium Payload | CBOR | Unmarshal | 70131 | 23928 | 496 |
| Medium Payload | JSON | Unmarshal | 261968 | 56008 | 738 |
| Small Struct | BEVE ZeroCopy | Marshal | 452 | 0 | 0 |
| Small Struct | CBOR | Marshal | 1584 | 1280 | 1 |
| Small Struct | BEVE | Marshal | 2117 | 2688 | 1 |
| Small Struct | Sonic | Marshal | 2179 | 3130 | 2 |
| Small Struct | MessagePack | Marshal | 3697 | 4104 | 8 |
| Small Struct | JSON | Marshal | 4268 | 2048 | 1 |
| Small Struct | BEVE | Unmarshal | 2740 | 3000 | 4 |
| Small Struct | Sonic | Unmarshal | 5332 | 7739 | 10 |
| Small Struct | MessagePack | Unmarshal | 6559 | 4352 | 92 |
| Small Struct | CBOR | Unmarshal | 7625 | 3568 | 76 |
| Small Struct | JSON | Unmarshal | 14263 | 3720 | 53 |
