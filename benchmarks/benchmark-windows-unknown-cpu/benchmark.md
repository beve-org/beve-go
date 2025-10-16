# Unknown CPU — Windows

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE ZeroCopy | Marshal | 85485 | 233 | 2 |
| Large Payload | BEVE | Marshal | 126748 | 197220 | 3 |
| Large Payload | Sonic | Marshal | 157106 | 220323 | 4 |
| Large Payload | CBOR | Marshal | 215099 | 189635 | 2 |
| Large Payload | MessagePack | Marshal | 294886 | 526763 | 115 |
| Large Payload | JSON | Marshal | 486931 | 206738 | 9 |
| Large Payload | BEVE | Unmarshal | 280005 | 279778 | 418 |
| Large Payload | Sonic | Unmarshal | 436268 | 565620 | 598 |
| Large Payload | MessagePack | Unmarshal | 698021 | 338819 | 6156 |
| Large Payload | CBOR | Unmarshal | 865811 | 312266 | 6361 |
| Large Payload | JSON | Unmarshal | 2658080 | 566916 | 7263 |
| Medium Payload | BEVE ZeroCopy | Marshal | 8868 | 138 | 2 |
| Medium Payload | BEVE | Marshal | 17837 | 27410 | 3 |
| Medium Payload | Sonic | Marshal | 18549 | 25152 | 4 |
| Medium Payload | CBOR | Marshal | 19471 | 18506 | 2 |
| Medium Payload | MessagePack | Marshal | 34737 | 65827 | 22 |
| Medium Payload | JSON | Marshal | 51841 | 24913 | 9 |
| Medium Payload | BEVE | Unmarshal | 27082 | 28635 | 59 |
| Medium Payload | Sonic | Unmarshal | 43656 | 47545 | 70 |
| Medium Payload | MessagePack | Unmarshal | 68818 | 36621 | 679 |
| Medium Payload | CBOR | Unmarshal | 95143 | 35528 | 734 |
| Medium Payload | JSON | Unmarshal | 251823 | 53352 | 715 |
| Small Struct | BEVE ZeroCopy | Marshal | 986 | 289 | 2 |
| Small Struct | CBOR | Marshal | 1246 | 1168 | 2 |
| Small Struct | Sonic | Marshal | 1795 | 2264 | 3 |
| Small Struct | BEVE | Marshal | 2282 | 2593 | 3 |
| Small Struct | MessagePack | Marshal | 5664 | 8320 | 9 |
| Small Struct | JSON | Marshal | 5831 | 2833 | 2 |
| Small Struct | MessagePack | Unmarshal | 1603 | 544 | 14 |
| Small Struct | Sonic | Unmarshal | 1962 | 1989 | 8 |
| Small Struct | BEVE | Unmarshal | 2210 | 3000 | 4 |
| Small Struct | CBOR | Unmarshal | 9431 | 3816 | 80 |
| Small Struct | JSON | Unmarshal | 32179 | 8008 | 116 |
