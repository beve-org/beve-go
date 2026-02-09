# AMD EPYC 7763 64-Core Processor — Linux

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE ZeroCopy | Marshal | 83674 | 26 | 0 |
| Large Payload | BEVE | Marshal | 116516 | 196706 | 1 |
| Large Payload | Sonic | Marshal | 158920 | 223650 | 3 |
| Large Payload | CBOR | Marshal | 212512 | 196812 | 1 |
| Large Payload | MessagePack | Marshal | 315182 | 526778 | 115 |
| Large Payload | JSON | Marshal | 450946 | 221476 | 8 |
| Large Payload | BEVE | Unmarshal | 229242 | 259198 | 418 |
| Large Payload | Sonic | Unmarshal | 356746 | 552064 | 590 |
| Large Payload | MessagePack | Unmarshal | 550878 | 339237 | 6170 |
| Large Payload | CBOR | Unmarshal | 699979 | 316586 | 6443 |
| Large Payload | JSON | Unmarshal | 2260306 | 541611 | 7050 |
| Medium Payload | BEVE ZeroCopy | Marshal | 8885 | 0 | 0 |
| Medium Payload | BEVE | Marshal | 13991 | 24584 | 1 |
| Medium Payload | Sonic | Marshal | 16023 | 22217 | 3 |
| Medium Payload | CBOR | Marshal | 25353 | 24601 | 1 |
| Medium Payload | MessagePack | Marshal | 35348 | 65782 | 22 |
| Medium Payload | JSON | Marshal | 39465 | 19305 | 8 |
| Medium Payload | BEVE | Unmarshal | 21915 | 23389 | 59 |
| Medium Payload | Sonic | Unmarshal | 40004 | 58430 | 73 |
| Medium Payload | MessagePack | Unmarshal | 58369 | 39153 | 732 |
| Medium Payload | CBOR | Unmarshal | 75641 | 36216 | 744 |
| Medium Payload | JSON | Unmarshal | 224475 | 58536 | 732 |
| Small Struct | BEVE ZeroCopy | Marshal | 727 | 0 | 0 |
| Small Struct | Sonic | Marshal | 830 | 926 | 2 |
| Small Struct | BEVE | Marshal | 1388 | 2048 | 1 |
| Small Struct | CBOR | Marshal | 1426 | 1280 | 1 |
| Small Struct | MessagePack | Marshal | 4316 | 8201 | 9 |
| Small Struct | JSON | Marshal | 4660 | 2688 | 1 |
| Small Struct | BEVE | Unmarshal | 2047 | 3512 | 4 |
| Small Struct | MessagePack | Unmarshal | 2063 | 1176 | 27 |
| Small Struct | Sonic | Unmarshal | 4007 | 7415 | 10 |
| Small Struct | CBOR | Unmarshal | 8274 | 4800 | 103 |
| Small Struct | JSON | Unmarshal | 21585 | 4840 | 88 |
