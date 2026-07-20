# Unknown CPU — Windows

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE ZeroCopy | Marshal | 82632 | 92 | 0 |
| Large Payload | BEVE | Marshal | 120345 | 188451 | 1 |
| Large Payload | Sonic | Marshal | 178435 | 223344 | 3 |
| Large Payload | CBOR | Marshal | 271078 | 196761 | 1 |
| Large Payload | MessagePack | Marshal | 288255 | 526706 | 115 |
| Large Payload | JSON | Marshal | 470871 | 205074 | 8 |
| Large Payload | BEVE | Unmarshal | 296913 | 265604 | 417 |
| Large Payload | Sonic | Unmarshal | 524394 | 573279 | 597 |
| Large Payload | MessagePack | Unmarshal | 869561 | 347132 | 6317 |
| Large Payload | CBOR | Unmarshal | 878312 | 306491 | 6255 |
| Large Payload | JSON | Unmarshal | 2725592 | 561187 | 7280 |
| Medium Payload | BEVE ZeroCopy | Marshal | 8659 | 5 | 0 |
| Medium Payload | BEVE | Marshal | 12780 | 18433 | 1 |
| Medium Payload | Sonic | Marshal | 18190 | 22140 | 3 |
| Medium Payload | CBOR | Marshal | 21330 | 18436 | 1 |
| Medium Payload | MessagePack | Marshal | 37687 | 65772 | 22 |
| Medium Payload | JSON | Marshal | 53353 | 24809 | 8 |
| Medium Payload | BEVE | Unmarshal | 31987 | 32476 | 59 |
| Medium Payload | Sonic | Unmarshal | 52260 | 60200 | 73 |
| Medium Payload | MessagePack | Unmarshal | 70633 | 31564 | 578 |
| Medium Payload | CBOR | Unmarshal | 101739 | 36760 | 752 |
| Medium Payload | JSON | Unmarshal | 273104 | 54504 | 689 |
| Small Struct | BEVE ZeroCopy | Marshal | 553 | 0 | 0 |
| Small Struct | BEVE | Marshal | 880 | 896 | 1 |
| Small Struct | CBOR | Marshal | 1990 | 1536 | 1 |
| Small Struct | Sonic | Marshal | 2909 | 2789 | 2 |
| Small Struct | MessagePack | Marshal | 3596 | 4104 | 8 |
| Small Struct | JSON | Marshal | 5166 | 2305 | 1 |
| Small Struct | BEVE | Unmarshal | 2173 | 2104 | 4 |
| Small Struct | Sonic | Unmarshal | 5281 | 7350 | 10 |
| Small Struct | CBOR | Unmarshal | 6165 | 2504 | 55 |
| Small Struct | MessagePack | Unmarshal | 6178 | 3584 | 76 |
| Small Struct | JSON | Unmarshal | 10507 | 2152 | 39 |
