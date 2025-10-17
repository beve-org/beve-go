# Unknown CPU — Windows

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE ZeroCopy | Marshal | 78181 | 52 | 0 |
| Large Payload | BEVE | Marshal | 120427 | 196656 | 1 |
| Large Payload | Sonic | Marshal | 166373 | 215088 | 3 |
| Large Payload | CBOR | Marshal | 222710 | 204899 | 1 |
| Large Payload | MessagePack | Marshal | 271406 | 526707 | 115 |
| Large Payload | JSON | Marshal | 459662 | 205124 | 8 |
| Large Payload | BEVE | Unmarshal | 267331 | 262306 | 417 |
| Large Payload | Sonic | Unmarshal | 415226 | 538328 | 581 |
| Large Payload | MessagePack | Unmarshal | 682345 | 354884 | 6479 |
| Large Payload | CBOR | Unmarshal | 824112 | 301786 | 6167 |
| Large Payload | JSON | Unmarshal | 2598279 | 535564 | 6991 |
| Medium Payload | BEVE ZeroCopy | Marshal | 7509 | 5 | 0 |
| Medium Payload | BEVE | Marshal | 13861 | 20484 | 1 |
| Medium Payload | Sonic | Marshal | 15023 | 18805 | 3 |
| Medium Payload | CBOR | Marshal | 23697 | 19079 | 1 |
| Medium Payload | MessagePack | Marshal | 25174 | 33002 | 21 |
| Medium Payload | JSON | Marshal | 53811 | 24809 | 8 |
| Medium Payload | BEVE | Unmarshal | 25218 | 23515 | 59 |
| Medium Payload | Sonic | Unmarshal | 44642 | 54108 | 72 |
| Medium Payload | MessagePack | Unmarshal | 76495 | 40749 | 760 |
| Medium Payload | CBOR | Unmarshal | 95601 | 35528 | 726 |
| Medium Payload | JSON | Unmarshal | 256477 | 55496 | 722 |
| Small Struct | BEVE ZeroCopy | Marshal | 652 | 0 | 0 |
| Small Struct | BEVE | Marshal | 1989 | 2688 | 1 |
| Small Struct | Sonic | Marshal | 2503 | 2725 | 2 |
| Small Struct | CBOR | Marshal | 2947 | 2688 | 1 |
| Small Struct | JSON | Marshal | 3695 | 1792 | 1 |
| Small Struct | MessagePack | Marshal | 3884 | 8200 | 9 |
| Small Struct | BEVE | Unmarshal | 1650 | 1848 | 4 |
| Small Struct | Sonic | Unmarshal | 5974 | 7777 | 10 |
| Small Struct | MessagePack | Unmarshal | 6817 | 3904 | 82 |
| Small Struct | CBOR | Unmarshal | 8272 | 3616 | 78 |
| Small Struct | JSON | Unmarshal | 27292 | 7304 | 94 |
