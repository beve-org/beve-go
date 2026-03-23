# Unknown CPU — Windows

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE ZeroCopy | Marshal | 80894 | 79 | 0 |
| Large Payload | BEVE | Marshal | 119731 | 196656 | 1 |
| Large Payload | Sonic | Marshal | 177869 | 223816 | 3 |
| Large Payload | CBOR | Marshal | 221022 | 196683 | 1 |
| Large Payload | MessagePack | Marshal | 287021 | 526705 | 115 |
| Large Payload | JSON | Marshal | 538661 | 237845 | 8 |
| Large Payload | BEVE | Unmarshal | 284929 | 267042 | 417 |
| Large Payload | Sonic | Unmarshal | 447238 | 547108 | 588 |
| Large Payload | MessagePack | Unmarshal | 705503 | 348678 | 6347 |
| Large Payload | CBOR | Unmarshal | 832409 | 289657 | 5902 |
| Large Payload | JSON | Unmarshal | 2519719 | 507163 | 6507 |
| Medium Payload | BEVE ZeroCopy | Marshal | 7710 | 1 | 0 |
| Medium Payload | BEVE | Marshal | 14780 | 24579 | 1 |
| Medium Payload | Sonic | Marshal | 19898 | 24916 | 3 |
| Medium Payload | CBOR | Marshal | 25323 | 21768 | 1 |
| Medium Payload | MessagePack | Marshal | 37006 | 65772 | 22 |
| Medium Payload | JSON | Marshal | 52079 | 24810 | 8 |
| Medium Payload | BEVE | Unmarshal | 25249 | 22107 | 58 |
| Medium Payload | Sonic | Unmarshal | 57901 | 67203 | 79 |
| Medium Payload | MessagePack | Unmarshal | 67786 | 35213 | 650 |
| Medium Payload | CBOR | Unmarshal | 81615 | 28744 | 594 |
| Medium Payload | JSON | Unmarshal | 254807 | 51112 | 679 |
| Small Struct | BEVE ZeroCopy | Marshal | 1039 | 0 | 0 |
| Small Struct | Sonic | Marshal | 1302 | 1310 | 2 |
| Small Struct | MessagePack | Marshal | 1890 | 2056 | 7 |
| Small Struct | BEVE | Marshal | 2578 | 2048 | 1 |
| Small Struct | CBOR | Marshal | 2984 | 3072 | 1 |
| Small Struct | JSON | Marshal | 5715 | 2304 | 1 |
| Small Struct | BEVE | Unmarshal | 1576 | 1592 | 4 |
| Small Struct | Sonic | Unmarshal | 3487 | 3632 | 9 |
| Small Struct | CBOR | Unmarshal | 4057 | 1640 | 37 |
| Small Struct | MessagePack | Unmarshal | 4624 | 2520 | 55 |
| Small Struct | JSON | Unmarshal | 17375 | 4008 | 62 |
