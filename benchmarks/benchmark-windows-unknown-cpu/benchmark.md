# Unknown CPU — Windows

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE ZeroCopy | Marshal | 72778 | 65 | 0 |
| Large Payload | BEVE | Marshal | 128918 | 180334 | 1 |
| Large Payload | Sonic | Marshal | 206421 | 225466 | 3 |
| Large Payload | CBOR | Marshal | 266685 | 204990 | 1 |
| Large Payload | MessagePack | Marshal | 337608 | 526736 | 115 |
| Large Payload | JSON | Marshal | 497371 | 221547 | 8 |
| Large Payload | BEVE | Unmarshal | 298779 | 264496 | 419 |
| Large Payload | Sonic | Unmarshal | 465831 | 550758 | 576 |
| Large Payload | MessagePack | Unmarshal | 669339 | 344324 | 6250 |
| Large Payload | CBOR | Unmarshal | 765405 | 290121 | 5910 |
| Large Payload | JSON | Unmarshal | 2226548 | 497202 | 6562 |
| Medium Payload | BEVE ZeroCopy | Marshal | 7144 | 1 | 0 |
| Medium Payload | BEVE | Marshal | 13659 | 16394 | 1 |
| Medium Payload | Sonic | Marshal | 22955 | 25055 | 3 |
| Medium Payload | CBOR | Marshal | 28264 | 24598 | 1 |
| Medium Payload | MessagePack | Marshal | 43109 | 65776 | 22 |
| Medium Payload | JSON | Marshal | 49938 | 24817 | 8 |
| Medium Payload | BEVE | Unmarshal | 28722 | 26428 | 59 |
| Medium Payload | Sonic | Unmarshal | 50258 | 55807 | 76 |
| Medium Payload | MessagePack | Unmarshal | 62607 | 33054 | 609 |
| Medium Payload | CBOR | Unmarshal | 80752 | 32808 | 678 |
| Medium Payload | JSON | Unmarshal | 270083 | 70264 | 885 |
| Small Struct | BEVE ZeroCopy | Marshal | 366 | 0 | 0 |
| Small Struct | Sonic | Marshal | 484 | 421 | 2 |
| Small Struct | BEVE | Marshal | 727 | 768 | 1 |
| Small Struct | CBOR | Marshal | 2195 | 2048 | 1 |
| Small Struct | MessagePack | Marshal | 3414 | 4104 | 8 |
| Small Struct | JSON | Marshal | 4134 | 2048 | 1 |
| Small Struct | BEVE | Unmarshal | 1442 | 1592 | 4 |
| Small Struct | CBOR | Unmarshal | 3025 | 1160 | 27 |
| Small Struct | Sonic | Unmarshal | 5538 | 7052 | 10 |
| Small Struct | MessagePack | Unmarshal | 6230 | 3936 | 83 |
| Small Struct | JSON | Unmarshal | 12841 | 3752 | 54 |
