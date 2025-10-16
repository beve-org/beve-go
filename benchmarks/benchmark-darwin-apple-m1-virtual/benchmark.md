# Apple M1 (Virtual) — Darwin

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE ZeroCopy | Marshal | 63805 | 180 | 2 |
| Large Payload | BEVE | Marshal | 128775 | 189496 | 3 |
| Large Payload | CBOR | Marshal | 178243 | 197360 | 2 |
| Large Payload | MessagePack | Marshal | 288062 | 526814 | 115 |
| Large Payload | JSON | Marshal | 443967 | 214157 | 9 |
| Large Payload | Sonic | Marshal | 524893 | 223300 | 4 |
| Large Payload | BEVE | Unmarshal | 257082 | 283219 | 419 |
| Large Payload | Sonic | Unmarshal | 337235 | 351503 | 211 |
| Large Payload | MessagePack | Unmarshal | 527954 | 353095 | 6448 |
| Large Payload | CBOR | Unmarshal | 766718 | 317354 | 6469 |
| Large Payload | JSON | Unmarshal | 2374683 | 542722 | 7082 |
| Medium Payload | BEVE ZeroCopy | Marshal | 7237 | 134 | 2 |
| Medium Payload | BEVE | Marshal | 12910 | 20625 | 3 |
| Medium Payload | CBOR | Marshal | 22867 | 20581 | 2 |
| Medium Payload | JSON | Marshal | 44206 | 19369 | 9 |
| Medium Payload | MessagePack | Marshal | 46825 | 65834 | 22 |
| Medium Payload | Sonic | Marshal | 67386 | 27690 | 4 |
| Medium Payload | BEVE | Unmarshal | 34737 | 31581 | 59 |
| Medium Payload | Sonic | Unmarshal | 45376 | 36324 | 33 |
| Medium Payload | MessagePack | Unmarshal | 58419 | 32205 | 590 |
| Medium Payload | CBOR | Unmarshal | 71727 | 33720 | 692 |
| Medium Payload | JSON | Unmarshal | 209768 | 44472 | 584 |
| Small Struct | CBOR | Marshal | 364 | 352 | 2 |
| Small Struct | BEVE ZeroCopy | Marshal | 507 | 290 | 2 |
| Small Struct | JSON | Marshal | 1488 | 1040 | 2 |
| Small Struct | MessagePack | Marshal | 1897 | 2176 | 7 |
| Small Struct | BEVE | Marshal | 2155 | 2978 | 3 |
| Small Struct | Sonic | Marshal | 3953 | 2208 | 3 |
| Small Struct | MessagePack | Unmarshal | 852 | 352 | 10 |
| Small Struct | BEVE | Unmarshal | 1922 | 2360 | 4 |
| Small Struct | Sonic | Unmarshal | 2898 | 4140 | 6 |
| Small Struct | CBOR | Unmarshal | 6397 | 4720 | 100 |
| Small Struct | JSON | Unmarshal | 26637 | 7816 | 110 |
