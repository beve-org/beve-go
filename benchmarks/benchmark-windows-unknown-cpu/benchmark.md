# Unknown CPU — Windows

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE ZeroCopy | Marshal | 77123 | 65 | 0 |
| Large Payload | BEVE | Marshal | 121263 | 196670 | 1 |
| Large Payload | Sonic | Marshal | 171252 | 223233 | 3 |
| Large Payload | CBOR | Marshal | 230122 | 188523 | 1 |
| Large Payload | MessagePack | Marshal | 293190 | 526709 | 115 |
| Large Payload | JSON | Marshal | 499167 | 221460 | 8 |
| Large Payload | BEVE | Unmarshal | 265397 | 248641 | 419 |
| Large Payload | Sonic | Unmarshal | 437103 | 555680 | 579 |
| Large Payload | MessagePack | Unmarshal | 720014 | 373081 | 6842 |
| Large Payload | CBOR | Unmarshal | 984555 | 350105 | 7144 |
| Large Payload | JSON | Unmarshal | 2707809 | 536891 | 7092 |
| Medium Payload | BEVE ZeroCopy | Marshal | 7689 | 5 | 0 |
| Medium Payload | BEVE | Marshal | 13076 | 16388 | 1 |
| Medium Payload | Sonic | Marshal | 21210 | 24968 | 3 |
| Medium Payload | CBOR | Marshal | 26153 | 24581 | 1 |
| Medium Payload | MessagePack | Marshal | 42949 | 65774 | 22 |
| Medium Payload | JSON | Marshal | 52467 | 20710 | 8 |
| Medium Payload | BEVE | Unmarshal | 32301 | 30843 | 59 |
| Medium Payload | Sonic | Unmarshal | 55182 | 64825 | 78 |
| Medium Payload | MessagePack | Unmarshal | 73184 | 39629 | 743 |
| Medium Payload | CBOR | Unmarshal | 102724 | 36488 | 754 |
| Medium Payload | JSON | Unmarshal | 326775 | 70104 | 894 |
| Small Struct | CBOR | Marshal | 448 | 256 | 1 |
| Small Struct | BEVE ZeroCopy | Marshal | 560 | 0 | 0 |
| Small Struct | Sonic | Marshal | 1369 | 1850 | 2 |
| Small Struct | BEVE | Marshal | 2032 | 2689 | 1 |
| Small Struct | JSON | Marshal | 2533 | 1152 | 1 |
| Small Struct | MessagePack | Marshal | 4864 | 8200 | 9 |
| Small Struct | Sonic | Unmarshal | 625 | 384 | 3 |
| Small Struct | BEVE | Unmarshal | 974 | 728 | 4 |
| Small Struct | CBOR | Unmarshal | 2662 | 904 | 22 |
| Small Struct | JSON | Unmarshal | 2883 | 464 | 12 |
| Small Struct | MessagePack | Unmarshal | 4135 | 2496 | 54 |
