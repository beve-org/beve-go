# Unknown CPU — Windows

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE ZeroCopy | Marshal | 81045 | 65 | 0 |
| Large Payload | BEVE | Marshal | 111978 | 188463 | 1 |
| Large Payload | Sonic | Marshal | 151111 | 205971 | 3 |
| Large Payload | CBOR | Marshal | 221152 | 196707 | 1 |
| Large Payload | MessagePack | Marshal | 276934 | 526706 | 115 |
| Large Payload | JSON | Marshal | 492707 | 229650 | 8 |
| Large Payload | BEVE | Unmarshal | 268294 | 268293 | 418 |
| Large Payload | Sonic | Unmarshal | 424423 | 546677 | 586 |
| Large Payload | MessagePack | Unmarshal | 662303 | 350051 | 6368 |
| Large Payload | CBOR | Unmarshal | 840849 | 313849 | 6394 |
| Large Payload | JSON | Unmarshal | 2550980 | 519260 | 6800 |
| Medium Payload | BEVE ZeroCopy | Marshal | 8346 | 1 | 0 |
| Medium Payload | BEVE | Marshal | 15496 | 24579 | 1 |
| Medium Payload | Sonic | Marshal | 16789 | 20779 | 3 |
| Medium Payload | MessagePack | Marshal | 25359 | 33002 | 21 |
| Medium Payload | CBOR | Marshal | 27421 | 24584 | 1 |
| Medium Payload | JSON | Marshal | 44285 | 20710 | 8 |
| Medium Payload | BEVE | Unmarshal | 26395 | 26235 | 59 |
| Medium Payload | Sonic | Unmarshal | 51393 | 58845 | 74 |
| Medium Payload | MessagePack | Unmarshal | 64675 | 33932 | 627 |
| Medium Payload | CBOR | Unmarshal | 87647 | 32824 | 672 |
| Medium Payload | JSON | Unmarshal | 265412 | 54648 | 731 |
| Small Struct | BEVE ZeroCopy | Marshal | 357 | 0 | 0 |
| Small Struct | MessagePack | Marshal | 1082 | 1032 | 6 |
| Small Struct | BEVE | Marshal | 1118 | 1152 | 1 |
| Small Struct | CBOR | Marshal | 1666 | 1792 | 1 |
| Small Struct | Sonic | Marshal | 2282 | 2746 | 2 |
| Small Struct | JSON | Marshal | 2546 | 1152 | 1 |
| Small Struct | MessagePack | Unmarshal | 1055 | 256 | 7 |
| Small Struct | Sonic | Unmarshal | 1673 | 1333 | 7 |
| Small Struct | BEVE | Unmarshal | 2247 | 2360 | 4 |
| Small Struct | CBOR | Unmarshal | 5949 | 2472 | 54 |
| Small Struct | JSON | Unmarshal | 33234 | 8040 | 117 |
