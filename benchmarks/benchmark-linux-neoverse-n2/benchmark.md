# Neoverse-N2 — Linux

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE ZeroCopy | Marshal | 68481 | 65 | 0 |
| Large Payload | BEVE | Marshal | 119642 | 188683 | 1 |
| Large Payload | CBOR | Marshal | 189590 | 196805 | 1 |
| Large Payload | MessagePack | Marshal | 289618 | 526802 | 115 |
| Large Payload | Sonic | Marshal | 325949 | 218216 | 3 |
| Large Payload | JSON | Marshal | 417131 | 221641 | 8 |
| Large Payload | BEVE | Unmarshal | 233219 | 278351 | 419 |
| Large Payload | Sonic | Unmarshal | 302565 | 407985 | 213 |
| Large Payload | MessagePack | Unmarshal | 510791 | 337513 | 6125 |
| Large Payload | CBOR | Unmarshal | 652878 | 311562 | 6336 |
| Large Payload | JSON | Unmarshal | 1889136 | 501884 | 6629 |
| Medium Payload | BEVE ZeroCopy | Marshal | 6739 | 0 | 0 |
| Medium Payload | BEVE | Marshal | 10788 | 20486 | 1 |
| Medium Payload | CBOR | Marshal | 19461 | 20498 | 1 |
| Medium Payload | Sonic | Marshal | 30568 | 22142 | 3 |
| Medium Payload | MessagePack | Marshal | 32538 | 65782 | 22 |
| Medium Payload | JSON | Marshal | 34371 | 18662 | 8 |
| Medium Payload | Sonic | Unmarshal | 24767 | 30626 | 33 |
| Medium Payload | BEVE | Unmarshal | 25429 | 35583 | 59 |
| Medium Payload | MessagePack | Unmarshal | 56159 | 36479 | 676 |
| Medium Payload | CBOR | Unmarshal | 67817 | 33272 | 681 |
| Medium Payload | JSON | Unmarshal | 217074 | 59288 | 811 |
| Small Struct | BEVE ZeroCopy | Marshal | 536 | 0 | 0 |
| Small Struct | BEVE | Marshal | 613 | 896 | 1 |
| Small Struct | Sonic | Marshal | 925 | 517 | 2 |
| Small Struct | JSON | Marshal | 2122 | 1152 | 1 |
| Small Struct | MessagePack | Marshal | 2376 | 4104 | 8 |
| Small Struct | CBOR | Marshal | 2527 | 3072 | 1 |
| Small Struct | BEVE | Unmarshal | 1301 | 2104 | 4 |
| Small Struct | Sonic | Unmarshal | 1357 | 1249 | 6 |
| Small Struct | CBOR | Unmarshal | 1444 | 424 | 12 |
| Small Struct | MessagePack | Unmarshal | 5393 | 4616 | 96 |
| Small Struct | JSON | Unmarshal | 15180 | 4296 | 71 |
