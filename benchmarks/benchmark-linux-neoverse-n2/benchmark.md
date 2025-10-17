# Neoverse-N2 — Linux

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE ZeroCopy | Marshal | 66935 | 65 | 0 |
| Large Payload | BEVE | Marshal | 110508 | 188537 | 1 |
| Large Payload | CBOR | Marshal | 190197 | 196857 | 1 |
| Large Payload | MessagePack | Marshal | 279544 | 526804 | 115 |
| Large Payload | Sonic | Marshal | 307659 | 216865 | 3 |
| Large Payload | JSON | Marshal | 376074 | 205308 | 8 |
| Large Payload | BEVE | Unmarshal | 220811 | 251651 | 416 |
| Large Payload | Sonic | Unmarshal | 303882 | 412197 | 213 |
| Large Payload | MessagePack | Unmarshal | 531652 | 353403 | 6436 |
| Large Payload | CBOR | Unmarshal | 646501 | 305898 | 6230 |
| Large Payload | JSON | Unmarshal | 1950074 | 524420 | 6843 |
| Medium Payload | BEVE ZeroCopy | Marshal | 8067 | 3 | 0 |
| Medium Payload | BEVE | Marshal | 9667 | 18437 | 1 |
| Medium Payload | CBOR | Marshal | 17281 | 18447 | 1 |
| Medium Payload | Sonic | Marshal | 24665 | 16907 | 3 |
| Medium Payload | MessagePack | Marshal | 32203 | 65782 | 22 |
| Medium Payload | JSON | Marshal | 41444 | 24807 | 8 |
| Medium Payload | BEVE | Unmarshal | 22242 | 24317 | 59 |
| Medium Payload | Sonic | Unmarshal | 32190 | 45643 | 31 |
| Medium Payload | MessagePack | Unmarshal | 49547 | 32478 | 594 |
| Medium Payload | CBOR | Unmarshal | 71436 | 35224 | 727 |
| Medium Payload | JSON | Unmarshal | 178488 | 50472 | 633 |
| Small Struct | BEVE | Marshal | 688 | 1024 | 1 |
| Small Struct | BEVE ZeroCopy | Marshal | 777 | 0 | 0 |
| Small Struct | CBOR | Marshal | 1004 | 896 | 1 |
| Small Struct | MessagePack | Marshal | 1463 | 2056 | 7 |
| Small Struct | Sonic | Marshal | 1726 | 1203 | 2 |
| Small Struct | JSON | Marshal | 4474 | 2688 | 1 |
| Small Struct | BEVE | Unmarshal | 1440 | 2360 | 4 |
| Small Struct | Sonic | Unmarshal | 1976 | 2831 | 6 |
| Small Struct | MessagePack | Unmarshal | 5828 | 4833 | 103 |
| Small Struct | CBOR | Unmarshal | 6654 | 3880 | 82 |
| Small Struct | JSON | Unmarshal | 25895 | 8072 | 118 |
