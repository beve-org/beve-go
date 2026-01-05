# Unknown CPU — Windows

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE ZeroCopy | Marshal | 82582 | 52 | 0 |
| Large Payload | BEVE | Marshal | 116064 | 188464 | 1 |
| Large Payload | Sonic | Marshal | 153528 | 206109 | 3 |
| Large Payload | CBOR | Marshal | 221133 | 196684 | 1 |
| Large Payload | MessagePack | Marshal | 265558 | 526705 | 115 |
| Large Payload | JSON | Marshal | 506381 | 213269 | 8 |
| Large Payload | BEVE | Unmarshal | 259588 | 263842 | 417 |
| Large Payload | Sonic | Unmarshal | 448230 | 558701 | 586 |
| Large Payload | MessagePack | Unmarshal | 695455 | 379847 | 6988 |
| Large Payload | CBOR | Unmarshal | 853221 | 320969 | 6548 |
| Large Payload | JSON | Unmarshal | 2429961 | 502027 | 6595 |
| Medium Payload | BEVE ZeroCopy | Marshal | 7323 | 5 | 0 |
| Medium Payload | BEVE | Marshal | 13967 | 20483 | 1 |
| Medium Payload | Sonic | Marshal | 15135 | 16637 | 3 |
| Medium Payload | CBOR | Marshal | 20615 | 18439 | 1 |
| Medium Payload | MessagePack | Marshal | 34787 | 65771 | 22 |
| Medium Payload | JSON | Marshal | 53077 | 24809 | 8 |
| Medium Payload | BEVE | Unmarshal | 25502 | 24698 | 59 |
| Medium Payload | Sonic | Unmarshal | 50215 | 60305 | 78 |
| Medium Payload | CBOR | Unmarshal | 63916 | 21176 | 437 |
| Medium Payload | MessagePack | Unmarshal | 88575 | 44023 | 831 |
| Medium Payload | JSON | Unmarshal | 290653 | 55776 | 706 |
| Small Struct | BEVE ZeroCopy | Marshal | 252 | 0 | 0 |
| Small Struct | MessagePack | Marshal | 1070 | 1032 | 6 |
| Small Struct | BEVE | Marshal | 1136 | 1408 | 1 |
| Small Struct | Sonic | Marshal | 1488 | 1473 | 2 |
| Small Struct | CBOR | Marshal | 1620 | 1408 | 1 |
| Small Struct | JSON | Marshal | 2501 | 1152 | 1 |
| Small Struct | BEVE | Unmarshal | 1145 | 1208 | 4 |
| Small Struct | MessagePack | Unmarshal | 1791 | 688 | 17 |
| Small Struct | Sonic | Unmarshal | 4019 | 4683 | 9 |
| Small Struct | CBOR | Unmarshal | 8940 | 4264 | 90 |
| Small Struct | JSON | Unmarshal | 23123 | 4744 | 85 |
