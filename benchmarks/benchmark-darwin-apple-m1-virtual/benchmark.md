# Apple M1 (Virtual) — Darwin

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE ZeroCopy | Marshal | 59168 | 180 | 2 |
| Large Payload | BEVE | Marshal | 116538 | 189205 | 3 |
| Large Payload | CBOR | Marshal | 250783 | 206172 | 2 |
| Large Payload | MessagePack | Marshal | 331588 | 526812 | 115 |
| Large Payload | JSON | Marshal | 474835 | 221773 | 9 |
| Large Payload | Sonic | Marshal | 528826 | 215657 | 4 |
| Large Payload | BEVE | Unmarshal | 264983 | 272879 | 419 |
| Large Payload | Sonic | Unmarshal | 406897 | 380973 | 211 |
| Large Payload | MessagePack | Unmarshal | 668984 | 356519 | 6511 |
| Large Payload | CBOR | Unmarshal | 671021 | 286506 | 5831 |
| Large Payload | JSON | Unmarshal | 2390962 | 560005 | 7316 |
| Medium Payload | BEVE ZeroCopy | Marshal | 9926 | 134 | 2 |
| Medium Payload | CBOR | Marshal | 16583 | 21848 | 2 |
| Medium Payload | BEVE | Marshal | 16764 | 20624 | 3 |
| Medium Payload | MessagePack | Marshal | 33302 | 65834 | 22 |
| Medium Payload | JSON | Marshal | 46705 | 19376 | 9 |
| Medium Payload | Sonic | Marshal | 47853 | 25004 | 4 |
| Medium Payload | BEVE | Unmarshal | 22714 | 29436 | 59 |
| Medium Payload | Sonic | Unmarshal | 35503 | 45974 | 33 |
| Medium Payload | MessagePack | Unmarshal | 37878 | 36045 | 668 |
| Medium Payload | CBOR | Unmarshal | 51322 | 33384 | 685 |
| Medium Payload | JSON | Unmarshal | 208328 | 62104 | 808 |
| Small Struct | BEVE ZeroCopy | Marshal | 471 | 288 | 2 |
| Small Struct | CBOR | Marshal | 2088 | 2833 | 2 |
| Small Struct | MessagePack | Marshal | 2336 | 4224 | 8 |
| Small Struct | BEVE | Marshal | 2764 | 2979 | 3 |
| Small Struct | JSON | Marshal | 3859 | 1936 | 2 |
| Small Struct | Sonic | Marshal | 4710 | 2507 | 3 |
| Small Struct | MessagePack | Unmarshal | 1684 | 680 | 17 |
| Small Struct | BEVE | Unmarshal | 2078 | 3384 | 4 |
| Small Struct | Sonic | Unmarshal | 4485 | 4780 | 6 |
| Small Struct | CBOR | Unmarshal | 9371 | 3944 | 84 |
| Small Struct | JSON | Unmarshal | 24505 | 7208 | 91 |
