# Apple M1 (Virtual) — Darwin

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE ZeroCopy | Marshal | 62536 | 39 | 0 |
| Large Payload | BEVE | Marshal | 114541 | 188483 | 1 |
| Large Payload | CBOR | Marshal | 154064 | 196817 | 1 |
| Large Payload | MessagePack | Marshal | 222309 | 526755 | 115 |
| Large Payload | JSON | Marshal | 384899 | 205192 | 8 |
| Large Payload | Sonic | Marshal | 395093 | 205362 | 3 |
| Large Payload | BEVE | Unmarshal | 182475 | 262416 | 416 |
| Large Payload | Sonic | Unmarshal | 304643 | 331809 | 207 |
| Large Payload | MessagePack | Unmarshal | 384066 | 348917 | 6358 |
| Large Payload | CBOR | Unmarshal | 541918 | 324825 | 6617 |
| Large Payload | JSON | Unmarshal | 1995906 | 562091 | 7308 |
| Medium Payload | BEVE ZeroCopy | Marshal | 6758 | 3 | 0 |
| Medium Payload | BEVE | Marshal | 11097 | 27271 | 1 |
| Medium Payload | CBOR | Marshal | 23691 | 21777 | 1 |
| Medium Payload | MessagePack | Marshal | 27925 | 65778 | 22 |
| Medium Payload | JSON | Marshal | 32193 | 24803 | 8 |
| Medium Payload | Sonic | Marshal | 39881 | 22016 | 3 |
| Medium Payload | BEVE | Unmarshal | 21386 | 26108 | 59 |
| Medium Payload | Sonic | Unmarshal | 34614 | 35013 | 33 |
| Medium Payload | MessagePack | Unmarshal | 65367 | 35325 | 657 |
| Medium Payload | CBOR | Unmarshal | 67011 | 34152 | 699 |
| Medium Payload | JSON | Unmarshal | 245327 | 66648 | 879 |
| Small Struct | BEVE | Marshal | 342 | 416 | 1 |
| Small Struct | BEVE ZeroCopy | Marshal | 348 | 0 | 0 |
| Small Struct | CBOR | Marshal | 1039 | 1152 | 1 |
| Small Struct | JSON | Marshal | 2827 | 2305 | 1 |
| Small Struct | MessagePack | Marshal | 2907 | 8201 | 9 |
| Small Struct | Sonic | Marshal | 5449 | 2725 | 2 |
| Small Struct | BEVE | Unmarshal | 913 | 2104 | 4 |
| Small Struct | CBOR | Unmarshal | 1186 | 656 | 17 |
| Small Struct | MessagePack | Unmarshal | 3190 | 3488 | 73 |
| Small Struct | Sonic | Unmarshal | 3278 | 5199 | 6 |
| Small Struct | JSON | Unmarshal | 4594 | 1320 | 28 |
