# AMD EPYC 7763 64-Core Processor — Linux

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE ZeroCopy | Marshal | 79760 | 52 | 0 |
| Large Payload | BEVE | Marshal | 115289 | 188499 | 1 |
| Large Payload | Sonic | Marshal | 163823 | 223735 | 3 |
| Large Payload | CBOR | Marshal | 209552 | 196785 | 1 |
| Large Payload | MessagePack | Marshal | 304806 | 526778 | 115 |
| Large Payload | JSON | Marshal | 422518 | 205090 | 8 |
| Large Payload | BEVE | Unmarshal | 235341 | 262686 | 418 |
| Large Payload | Sonic | Unmarshal | 365298 | 557865 | 585 |
| Large Payload | MessagePack | Unmarshal | 555833 | 343892 | 6252 |
| Large Payload | CBOR | Unmarshal | 739186 | 339146 | 6898 |
| Large Payload | JSON | Unmarshal | 2152315 | 503258 | 6535 |
| Medium Payload | BEVE ZeroCopy | Marshal | 8728 | 3 | 0 |
| Medium Payload | BEVE | Marshal | 11211 | 18439 | 1 |
| Medium Payload | Sonic | Marshal | 15865 | 21004 | 3 |
| Medium Payload | CBOR | Marshal | 21654 | 20499 | 1 |
| Medium Payload | MessagePack | Marshal | 37165 | 65783 | 22 |
| Medium Payload | JSON | Marshal | 39722 | 19303 | 8 |
| Medium Payload | BEVE | Unmarshal | 24507 | 27262 | 59 |
| Medium Payload | Sonic | Unmarshal | 47336 | 67547 | 74 |
| Medium Payload | MessagePack | Unmarshal | 56800 | 36704 | 683 |
| Medium Payload | CBOR | Unmarshal | 71063 | 32824 | 669 |
| Medium Payload | JSON | Unmarshal | 262029 | 66072 | 860 |
| Small Struct | BEVE ZeroCopy | Marshal | 268 | 0 | 0 |
| Small Struct | BEVE | Marshal | 1045 | 1792 | 1 |
| Small Struct | Sonic | Marshal | 1898 | 2775 | 2 |
| Small Struct | CBOR | Marshal | 2257 | 2304 | 1 |
| Small Struct | JSON | Marshal | 4275 | 2304 | 1 |
| Small Struct | MessagePack | Marshal | 4345 | 8201 | 9 |
| Small Struct | BEVE | Unmarshal | 1282 | 1592 | 4 |
| Small Struct | MessagePack | Unmarshal | 3199 | 2112 | 46 |
| Small Struct | Sonic | Unmarshal | 4270 | 7452 | 10 |
| Small Struct | JSON | Unmarshal | 5064 | 936 | 23 |
| Small Struct | CBOR | Unmarshal | 6909 | 3912 | 83 |
