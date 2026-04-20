# AMD EPYC 7763 64-Core Processor — Linux

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE ZeroCopy | Marshal | 83533 | 65 | 0 |
| Large Payload | BEVE | Marshal | 132096 | 196682 | 1 |
| Large Payload | Sonic | Marshal | 164149 | 207957 | 3 |
| Large Payload | CBOR | Marshal | 207460 | 188589 | 1 |
| Large Payload | MessagePack | Marshal | 356303 | 526781 | 115 |
| Large Payload | JSON | Marshal | 475446 | 221476 | 8 |
| Large Payload | BEVE | Unmarshal | 268371 | 280420 | 417 |
| Large Payload | Sonic | Unmarshal | 401806 | 553260 | 590 |
| Large Payload | MessagePack | Unmarshal | 583136 | 351465 | 6406 |
| Large Payload | CBOR | Unmarshal | 738928 | 319690 | 6513 |
| Large Payload | JSON | Unmarshal | 2320219 | 537852 | 6955 |
| Medium Payload | BEVE ZeroCopy | Marshal | 6979 | 3 | 0 |
| Medium Payload | Sonic | Marshal | 15577 | 20925 | 3 |
| Medium Payload | BEVE | Marshal | 19586 | 24586 | 1 |
| Medium Payload | CBOR | Marshal | 20748 | 18455 | 1 |
| Medium Payload | MessagePack | Marshal | 27488 | 33007 | 21 |
| Medium Payload | JSON | Marshal | 52169 | 24808 | 8 |
| Medium Payload | BEVE | Unmarshal | 26334 | 30079 | 59 |
| Medium Payload | Sonic | Unmarshal | 45281 | 64570 | 77 |
| Medium Payload | MessagePack | Unmarshal | 58340 | 37000 | 686 |
| Medium Payload | CBOR | Unmarshal | 77869 | 36248 | 748 |
| Medium Payload | JSON | Unmarshal | 192677 | 45288 | 591 |
| Small Struct | Sonic | Marshal | 615 | 414 | 2 |
| Small Struct | BEVE ZeroCopy | Marshal | 620 | 0 | 0 |
| Small Struct | BEVE | Marshal | 800 | 1024 | 1 |
| Small Struct | MessagePack | Marshal | 2269 | 2056 | 7 |
| Small Struct | CBOR | Marshal | 3184 | 2689 | 1 |
| Small Struct | JSON | Marshal | 4173 | 1792 | 1 |
| Small Struct | CBOR | Unmarshal | 2305 | 480 | 13 |
| Small Struct | BEVE | Unmarshal | 2862 | 3000 | 4 |
| Small Struct | MessagePack | Unmarshal | 3177 | 1120 | 26 |
| Small Struct | JSON | Unmarshal | 4014 | 544 | 15 |
| Small Struct | Sonic | Unmarshal | 5548 | 7020 | 10 |
