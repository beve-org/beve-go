# Apple M1 (Virtual) — Darwin

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE ZeroCopy | Marshal | 66650 | 26 | 0 |
| Large Payload | BEVE | Marshal | 113494 | 196649 | 1 |
| Large Payload | CBOR | Marshal | 154646 | 188570 | 1 |
| Large Payload | MessagePack | Marshal | 186156 | 526759 | 115 |
| Large Payload | JSON | Marshal | 422473 | 205166 | 8 |
| Large Payload | Sonic | Marshal | 494092 | 213778 | 3 |
| Large Payload | BEVE | Unmarshal | 164630 | 269779 | 417 |
| Large Payload | Sonic | Unmarshal | 261321 | 353997 | 211 |
| Large Payload | MessagePack | Unmarshal | 411500 | 341906 | 6203 |
| Large Payload | CBOR | Unmarshal | 528154 | 320106 | 6528 |
| Large Payload | JSON | Unmarshal | 1768461 | 546524 | 7197 |
| Medium Payload | BEVE ZeroCopy | Marshal | 9132 | 3 | 0 |
| Medium Payload | BEVE | Marshal | 16384 | 24583 | 1 |
| Medium Payload | CBOR | Marshal | 20098 | 21774 | 1 |
| Medium Payload | MessagePack | Marshal | 31828 | 65778 | 22 |
| Medium Payload | JSON | Marshal | 42085 | 20713 | 8 |
| Medium Payload | Sonic | Marshal | 67163 | 29005 | 3 |
| Medium Payload | BEVE | Unmarshal | 23352 | 22011 | 59 |
| Medium Payload | Sonic | Unmarshal | 35539 | 37342 | 33 |
| Medium Payload | MessagePack | Unmarshal | 55325 | 33389 | 614 |
| Medium Payload | CBOR | Unmarshal | 68357 | 32072 | 659 |
| Medium Payload | JSON | Unmarshal | 220408 | 56648 | 725 |
| Small Struct | BEVE ZeroCopy | Marshal | 415 | 0 | 0 |
| Small Struct | BEVE | Marshal | 451 | 896 | 1 |
| Small Struct | CBOR | Marshal | 2235 | 2048 | 1 |
| Small Struct | JSON | Marshal | 2592 | 2048 | 1 |
| Small Struct | MessagePack | Marshal | 2693 | 4104 | 8 |
| Small Struct | Sonic | Marshal | 7001 | 3109 | 2 |
| Small Struct | BEVE | Unmarshal | 1535 | 1720 | 4 |
| Small Struct | CBOR | Unmarshal | 3521 | 1352 | 31 |
| Small Struct | Sonic | Unmarshal | 4431 | 4664 | 6 |
| Small Struct | MessagePack | Unmarshal | 6707 | 3584 | 76 |
| Small Struct | JSON | Unmarshal | 19588 | 4616 | 81 |
