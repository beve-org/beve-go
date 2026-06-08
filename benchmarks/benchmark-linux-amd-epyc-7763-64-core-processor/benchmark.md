# AMD EPYC 7763 64-Core Processor — Linux

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE ZeroCopy | Marshal | 79494 | 39 | 0 |
| Large Payload | BEVE | Marshal | 113430 | 188473 | 1 |
| Large Payload | Sonic | Marshal | 152110 | 207289 | 3 |
| Large Payload | CBOR | Marshal | 210488 | 196758 | 1 |
| Large Payload | MessagePack | Marshal | 322014 | 526778 | 115 |
| Large Payload | JSON | Marshal | 433218 | 205116 | 8 |
| Large Payload | BEVE | Unmarshal | 241093 | 272480 | 418 |
| Large Payload | Sonic | Unmarshal | 388433 | 586950 | 603 |
| Large Payload | MessagePack | Unmarshal | 595889 | 368267 | 6737 |
| Large Payload | CBOR | Unmarshal | 667565 | 290009 | 5907 |
| Large Payload | JSON | Unmarshal | 2348521 | 536129 | 7021 |
| Medium Payload | BEVE ZeroCopy | Marshal | 7693 | 3 | 0 |
| Medium Payload | BEVE | Marshal | 11833 | 18437 | 1 |
| Medium Payload | Sonic | Marshal | 15696 | 20837 | 3 |
| Medium Payload | CBOR | Marshal | 19021 | 16401 | 1 |
| Medium Payload | JSON | Marshal | 37311 | 18666 | 8 |
| Medium Payload | MessagePack | Marshal | 38290 | 65783 | 22 |
| Medium Payload | BEVE | Unmarshal | 26810 | 31583 | 59 |
| Medium Payload | Sonic | Unmarshal | 34012 | 46587 | 66 |
| Medium Payload | MessagePack | Unmarshal | 54986 | 33968 | 624 |
| Medium Payload | CBOR | Unmarshal | 61076 | 25528 | 529 |
| Medium Payload | JSON | Unmarshal | 265000 | 62296 | 849 |
| Small Struct | BEVE ZeroCopy | Marshal | 465 | 0 | 0 |
| Small Struct | Sonic | Marshal | 583 | 606 | 2 |
| Small Struct | JSON | Marshal | 910 | 352 | 1 |
| Small Struct | BEVE | Marshal | 1294 | 2048 | 1 |
| Small Struct | CBOR | Marshal | 1382 | 1024 | 1 |
| Small Struct | MessagePack | Marshal | 3494 | 4104 | 8 |
| Small Struct | BEVE | Unmarshal | 1112 | 1080 | 4 |
| Small Struct | CBOR | Unmarshal | 2015 | 760 | 19 |
| Small Struct | MessagePack | Unmarshal | 3599 | 2496 | 54 |
| Small Struct | JSON | Unmarshal | 3932 | 680 | 18 |
| Small Struct | Sonic | Unmarshal | 4251 | 7420 | 10 |
