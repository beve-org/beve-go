# AMD EPYC 7763 64-Core Processor — Linux

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE ZeroCopy | Marshal | 76971 | 39 | 0 |
| Large Payload | BEVE | Marshal | 114608 | 180308 | 1 |
| Large Payload | Sonic | Marshal | 165238 | 215825 | 3 |
| Large Payload | CBOR | Marshal | 223529 | 205008 | 1 |
| Large Payload | MessagePack | Marshal | 338451 | 526781 | 115 |
| Large Payload | JSON | Marshal | 451693 | 213309 | 8 |
| Large Payload | BEVE | Unmarshal | 244607 | 267456 | 418 |
| Large Payload | Sonic | Unmarshal | 374278 | 549588 | 592 |
| Large Payload | MessagePack | Unmarshal | 570416 | 338229 | 6140 |
| Large Payload | CBOR | Unmarshal | 758643 | 336395 | 6854 |
| Large Payload | JSON | Unmarshal | 2282733 | 524722 | 6780 |
| Medium Payload | BEVE ZeroCopy | Marshal | 7429 | 3 | 0 |
| Medium Payload | BEVE | Marshal | 12210 | 18437 | 1 |
| Medium Payload | Sonic | Marshal | 18370 | 25119 | 3 |
| Medium Payload | CBOR | Marshal | 22342 | 20502 | 1 |
| Medium Payload | JSON | Marshal | 37629 | 18670 | 8 |
| Medium Payload | MessagePack | Marshal | 43194 | 65782 | 22 |
| Medium Payload | BEVE | Unmarshal | 24211 | 25758 | 59 |
| Medium Payload | Sonic | Unmarshal | 43551 | 64661 | 74 |
| Medium Payload | MessagePack | Unmarshal | 59893 | 38881 | 728 |
| Medium Payload | CBOR | Unmarshal | 73666 | 31864 | 653 |
| Medium Payload | JSON | Unmarshal | 260894 | 65096 | 864 |
| Small Struct | BEVE ZeroCopy | Marshal | 242 | 0 | 0 |
| Small Struct | CBOR | Marshal | 949 | 256 | 1 |
| Small Struct | BEVE | Marshal | 1950 | 2689 | 1 |
| Small Struct | MessagePack | Marshal | 2144 | 2056 | 7 |
| Small Struct | JSON | Marshal | 2181 | 640 | 1 |
| Small Struct | Sonic | Marshal | 2938 | 3180 | 2 |
| Small Struct | BEVE | Unmarshal | 1527 | 1016 | 4 |
| Small Struct | Sonic | Unmarshal | 2062 | 1957 | 8 |
| Small Struct | MessagePack | Unmarshal | 2937 | 1504 | 34 |
| Small Struct | CBOR | Unmarshal | 9175 | 3880 | 82 |
| Small Struct | JSON | Unmarshal | 19931 | 4200 | 68 |
