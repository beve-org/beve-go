# Intel(R) Xeon(R) Platinum 8370C CPU @ 2.80GHz — Linux

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE ZeroCopy | Marshal | 69235 | 26 | 0 |
| Large Payload | BEVE | Marshal | 125250 | 196708 | 1 |
| Large Payload | Sonic | Marshal | 168164 | 209245 | 3 |
| Large Payload | CBOR | Marshal | 206677 | 205009 | 1 |
| Large Payload | MessagePack | Marshal | 303715 | 526775 | 115 |
| Large Payload | JSON | Marshal | 423007 | 213283 | 8 |
| Large Payload | BEVE | Unmarshal | 245840 | 264703 | 419 |
| Large Payload | Sonic | Unmarshal | 370369 | 521649 | 578 |
| Large Payload | MessagePack | Unmarshal | 525397 | 328694 | 5931 |
| Large Payload | CBOR | Unmarshal | 712369 | 329785 | 6726 |
| Large Payload | JSON | Unmarshal | 2037040 | 527740 | 6893 |
| Medium Payload | BEVE ZeroCopy | Marshal | 7188 | 3 | 0 |
| Medium Payload | BEVE | Marshal | 13101 | 21768 | 1 |
| Medium Payload | Sonic | Marshal | 15263 | 19028 | 3 |
| Medium Payload | CBOR | Marshal | 18641 | 18452 | 1 |
| Medium Payload | JSON | Marshal | 35270 | 18663 | 8 |
| Medium Payload | MessagePack | Marshal | 36551 | 65782 | 22 |
| Medium Payload | BEVE | Unmarshal | 23235 | 25438 | 59 |
| Medium Payload | Sonic | Unmarshal | 41820 | 61300 | 70 |
| Medium Payload | MessagePack | Unmarshal | 53442 | 35520 | 663 |
| Medium Payload | CBOR | Unmarshal | 61102 | 27928 | 571 |
| Medium Payload | JSON | Unmarshal | 169369 | 44424 | 577 |
| Small Struct | BEVE ZeroCopy | Marshal | 457 | 0 | 0 |
| Small Struct | BEVE | Marshal | 1524 | 2048 | 1 |
| Small Struct | CBOR | Marshal | 1963 | 1792 | 1 |
| Small Struct | Sonic | Marshal | 2956 | 2775 | 2 |
| Small Struct | JSON | Marshal | 3356 | 1408 | 1 |
| Small Struct | MessagePack | Marshal | 3379 | 4104 | 8 |
| Small Struct | BEVE | Unmarshal | 2268 | 2360 | 4 |
| Small Struct | MessagePack | Unmarshal | 3454 | 1504 | 34 |
| Small Struct | Sonic | Unmarshal | 4010 | 4160 | 9 |
| Small Struct | CBOR | Unmarshal | 6407 | 3144 | 67 |
| Small Struct | JSON | Unmarshal | 23845 | 4840 | 88 |
