# AMD EPYC 9V74 80-Core Processor — Linux

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE ZeroCopy | Marshal | 73801 | 52 | 0 |
| Large Payload | BEVE | Marshal | 113142 | 188515 | 1 |
| Large Payload | Sonic | Marshal | 166130 | 216506 | 3 |
| Large Payload | CBOR | Marshal | 200304 | 196865 | 1 |
| Large Payload | MessagePack | Marshal | 327476 | 526781 | 115 |
| Large Payload | JSON | Marshal | 436331 | 213283 | 8 |
| Large Payload | BEVE | Unmarshal | 228655 | 268383 | 418 |
| Large Payload | Sonic | Unmarshal | 395819 | 550077 | 575 |
| Large Payload | MessagePack | Unmarshal | 529176 | 337445 | 6121 |
| Large Payload | CBOR | Unmarshal | 789376 | 328411 | 6694 |
| Large Payload | JSON | Unmarshal | 1994292 | 509250 | 6696 |
| Medium Payload | BEVE ZeroCopy | Marshal | 6556 | 5 | 0 |
| Medium Payload | BEVE | Marshal | 11166 | 18439 | 1 |
| Medium Payload | Sonic | Marshal | 17083 | 22280 | 3 |
| Medium Payload | CBOR | Marshal | 23637 | 21776 | 1 |
| Medium Payload | MessagePack | Marshal | 39076 | 65783 | 22 |
| Medium Payload | JSON | Marshal | 46110 | 24808 | 8 |
| Medium Payload | BEVE | Unmarshal | 26041 | 29215 | 59 |
| Medium Payload | Sonic | Unmarshal | 46200 | 59475 | 75 |
| Medium Payload | MessagePack | Unmarshal | 54259 | 34512 | 639 |
| Medium Payload | CBOR | Unmarshal | 76517 | 31512 | 649 |
| Medium Payload | JSON | Unmarshal | 208975 | 53320 | 724 |
| Small Struct | BEVE ZeroCopy | Marshal | 450 | 0 | 0 |
| Small Struct | BEVE | Marshal | 744 | 704 | 1 |
| Small Struct | CBOR | Marshal | 781 | 640 | 1 |
| Small Struct | Sonic | Marshal | 1090 | 1467 | 2 |
| Small Struct | MessagePack | Marshal | 2567 | 4104 | 8 |
| Small Struct | JSON | Marshal | 5245 | 2689 | 1 |
| Small Struct | BEVE | Unmarshal | 1058 | 1336 | 4 |
| Small Struct | MessagePack | Unmarshal | 1280 | 496 | 13 |
| Small Struct | Sonic | Unmarshal | 2458 | 3675 | 9 |
| Small Struct | CBOR | Unmarshal | 5239 | 2472 | 54 |
| Small Struct | JSON | Unmarshal | 24273 | 7664 | 105 |
