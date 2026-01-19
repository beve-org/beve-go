# Unknown CPU — Windows

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE ZeroCopy | Marshal | 80538 | 52 | 0 |
| Large Payload | BEVE | Marshal | 113573 | 188465 | 1 |
| Large Payload | Sonic | Marshal | 160860 | 222870 | 3 |
| Large Payload | CBOR | Marshal | 232859 | 204878 | 1 |
| Large Payload | MessagePack | Marshal | 286475 | 526709 | 115 |
| Large Payload | JSON | Marshal | 523049 | 221460 | 8 |
| Large Payload | BEVE | Unmarshal | 273528 | 273287 | 419 |
| Large Payload | Sonic | Unmarshal | 432607 | 543836 | 574 |
| Large Payload | MessagePack | Unmarshal | 681535 | 345704 | 6297 |
| Large Payload | CBOR | Unmarshal | 846918 | 302809 | 6186 |
| Large Payload | JSON | Unmarshal | 2588704 | 519141 | 6847 |
| Medium Payload | BEVE ZeroCopy | Marshal | 8340 | 5 | 0 |
| Medium Payload | BEVE | Marshal | 14496 | 21763 | 1 |
| Medium Payload | Sonic | Marshal | 16346 | 19405 | 3 |
| Medium Payload | CBOR | Marshal | 23483 | 20491 | 1 |
| Medium Payload | MessagePack | Marshal | 38375 | 65771 | 22 |
| Medium Payload | JSON | Marshal | 45818 | 18668 | 8 |
| Medium Payload | BEVE | Unmarshal | 29199 | 33019 | 59 |
| Medium Payload | Sonic | Unmarshal | 46096 | 54342 | 72 |
| Medium Payload | MessagePack | Unmarshal | 69264 | 35805 | 666 |
| Medium Payload | CBOR | Unmarshal | 75343 | 27176 | 560 |
| Medium Payload | JSON | Unmarshal | 271361 | 57640 | 768 |
| Small Struct | BEVE ZeroCopy | Marshal | 349 | 0 | 0 |
| Small Struct | BEVE | Marshal | 628 | 640 | 1 |
| Small Struct | Sonic | Marshal | 2092 | 2781 | 2 |
| Small Struct | CBOR | Marshal | 2418 | 1792 | 1 |
| Small Struct | MessagePack | Marshal | 3029 | 4104 | 8 |
| Small Struct | JSON | Marshal | 6083 | 2688 | 1 |
| Small Struct | BEVE | Unmarshal | 762 | 344 | 4 |
| Small Struct | CBOR | Unmarshal | 1406 | 320 | 10 |
| Small Struct | MessagePack | Unmarshal | 5080 | 3288 | 71 |
| Small Struct | Sonic | Unmarshal | 5181 | 6976 | 10 |
| Small Struct | JSON | Unmarshal | 18491 | 4200 | 68 |
