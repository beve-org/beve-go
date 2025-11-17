# Apple M1 (Virtual) — Darwin

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE ZeroCopy | Marshal | 59804 | 39 | 0 |
| Large Payload | BEVE | Marshal | 110540 | 196661 | 1 |
| Large Payload | CBOR | Marshal | 199755 | 188588 | 1 |
| Large Payload | MessagePack | Marshal | 304203 | 526757 | 115 |
| Large Payload | JSON | Marshal | 393157 | 205113 | 8 |
| Large Payload | Sonic | Marshal | 524462 | 214548 | 3 |
| Large Payload | BEVE | Unmarshal | 265605 | 257326 | 417 |
| Large Payload | Sonic | Unmarshal | 365971 | 340106 | 213 |
| Large Payload | MessagePack | Unmarshal | 572856 | 355383 | 6484 |
| Large Payload | CBOR | Unmarshal | 720305 | 314586 | 6422 |
| Large Payload | JSON | Unmarshal | 2073708 | 509867 | 6748 |
| Medium Payload | BEVE ZeroCopy | Marshal | 6971 | 0 | 0 |
| Medium Payload | BEVE | Marshal | 10409 | 16387 | 1 |
| Medium Payload | CBOR | Marshal | 16514 | 18444 | 1 |
| Medium Payload | MessagePack | Marshal | 33106 | 65778 | 22 |
| Medium Payload | JSON | Marshal | 45980 | 20707 | 8 |
| Medium Payload | Sonic | Marshal | 61891 | 24857 | 3 |
| Medium Payload | BEVE | Unmarshal | 24208 | 29373 | 59 |
| Medium Payload | MessagePack | Unmarshal | 35591 | 23260 | 407 |
| Medium Payload | Sonic | Unmarshal | 37720 | 42740 | 33 |
| Medium Payload | CBOR | Unmarshal | 57383 | 26200 | 542 |
| Medium Payload | JSON | Unmarshal | 217552 | 56840 | 735 |
| Small Struct | BEVE ZeroCopy | Marshal | 851 | 0 | 0 |
| Small Struct | BEVE | Marshal | 1136 | 1792 | 1 |
| Small Struct | Sonic | Marshal | 1907 | 926 | 2 |
| Small Struct | MessagePack | Marshal | 2013 | 4104 | 8 |
| Small Struct | CBOR | Marshal | 2836 | 2688 | 1 |
| Small Struct | JSON | Marshal | 6684 | 2689 | 1 |
| Small Struct | BEVE | Unmarshal | 1058 | 1336 | 4 |
| Small Struct | Sonic | Unmarshal | 3789 | 4967 | 6 |
| Small Struct | MessagePack | Unmarshal | 4160 | 3848 | 80 |
| Small Struct | CBOR | Unmarshal | 4363 | 2304 | 51 |
| Small Struct | JSON | Unmarshal | 10936 | 2312 | 44 |
