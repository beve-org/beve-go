# Unknown CPU — Windows

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE ZeroCopy | Marshal | 83009 | 92 | 0 |
| Large Payload | BEVE | Marshal | 115916 | 188477 | 1 |
| Large Payload | Sonic | Marshal | 160737 | 214812 | 3 |
| Large Payload | CBOR | Marshal | 212032 | 196678 | 1 |
| Large Payload | MessagePack | Marshal | 272856 | 526704 | 115 |
| Large Payload | JSON | Marshal | 473650 | 213291 | 8 |
| Large Payload | BEVE | Unmarshal | 273069 | 272322 | 419 |
| Large Payload | Sonic | Unmarshal | 428343 | 545160 | 585 |
| Large Payload | MessagePack | Unmarshal | 651716 | 337004 | 6107 |
| Large Payload | CBOR | Unmarshal | 881472 | 329386 | 6712 |
| Large Payload | JSON | Unmarshal | 2378693 | 484051 | 6408 |
| Medium Payload | BEVE ZeroCopy | Marshal | 7877 | 5 | 0 |
| Medium Payload | BEVE | Marshal | 13476 | 20483 | 1 |
| Medium Payload | Sonic | Marshal | 19636 | 24954 | 3 |
| Medium Payload | CBOR | Marshal | 20322 | 18439 | 1 |
| Medium Payload | MessagePack | Marshal | 24965 | 33001 | 21 |
| Medium Payload | JSON | Marshal | 49881 | 24806 | 8 |
| Medium Payload | BEVE | Unmarshal | 28369 | 27387 | 59 |
| Medium Payload | Sonic | Unmarshal | 52545 | 66007 | 77 |
| Medium Payload | MessagePack | Unmarshal | 73764 | 38157 | 711 |
| Medium Payload | CBOR | Unmarshal | 86804 | 30264 | 623 |
| Medium Payload | JSON | Unmarshal | 282422 | 59192 | 765 |
| Small Struct | BEVE ZeroCopy | Marshal | 379 | 0 | 0 |
| Small Struct | BEVE | Marshal | 864 | 704 | 1 |
| Small Struct | CBOR | Marshal | 1398 | 1408 | 1 |
| Small Struct | Sonic | Marshal | 2483 | 3130 | 2 |
| Small Struct | MessagePack | Marshal | 3099 | 4104 | 8 |
| Small Struct | JSON | Marshal | 3379 | 1408 | 1 |
| Small Struct | BEVE | Unmarshal | 1857 | 2616 | 4 |
| Small Struct | Sonic | Unmarshal | 2454 | 2357 | 8 |
| Small Struct | MessagePack | Unmarshal | 5316 | 3200 | 68 |
| Small Struct | CBOR | Unmarshal | 11115 | 5192 | 107 |
| Small Struct | JSON | Unmarshal | 27077 | 7424 | 98 |
