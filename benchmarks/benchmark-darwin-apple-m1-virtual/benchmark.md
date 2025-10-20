# Apple M1 (Virtual) — Darwin

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE ZeroCopy | Marshal | 55677 | 26 | 0 |
| Large Payload | BEVE | Marshal | 97210 | 188454 | 1 |
| Large Payload | CBOR | Marshal | 161969 | 196735 | 1 |
| Large Payload | MessagePack | Marshal | 280597 | 526757 | 115 |
| Large Payload | JSON | Marshal | 403257 | 213359 | 8 |
| Large Payload | Sonic | Marshal | 645140 | 222609 | 3 |
| Large Payload | BEVE | Unmarshal | 253054 | 266832 | 419 |
| Large Payload | Sonic | Unmarshal | 343892 | 361804 | 213 |
| Large Payload | MessagePack | Unmarshal | 477174 | 328176 | 5932 |
| Large Payload | CBOR | Unmarshal | 641821 | 308026 | 6291 |
| Large Payload | JSON | Unmarshal | 2300370 | 485299 | 6387 |
| Medium Payload | BEVE ZeroCopy | Marshal | 7583 | 3 | 0 |
| Medium Payload | BEVE | Marshal | 15543 | 19079 | 1 |
| Medium Payload | CBOR | Marshal | 17509 | 20490 | 1 |
| Medium Payload | MessagePack | Marshal | 25354 | 33005 | 21 |
| Medium Payload | JSON | Marshal | 45021 | 21997 | 8 |
| Medium Payload | Sonic | Marshal | 46987 | 19286 | 3 |
| Medium Payload | BEVE | Unmarshal | 23004 | 24764 | 59 |
| Medium Payload | Sonic | Unmarshal | 40749 | 46100 | 33 |
| Medium Payload | MessagePack | Unmarshal | 50255 | 39742 | 743 |
| Medium Payload | CBOR | Unmarshal | 67157 | 32392 | 664 |
| Medium Payload | JSON | Unmarshal | 236901 | 56776 | 758 |
| Small Struct | BEVE | Marshal | 210 | 176 | 1 |
| Small Struct | BEVE ZeroCopy | Marshal | 373 | 0 | 0 |
| Small Struct | CBOR | Marshal | 2707 | 3072 | 1 |
| Small Struct | MessagePack | Marshal | 3178 | 8201 | 9 |
| Small Struct | JSON | Marshal | 3253 | 1792 | 1 |
| Small Struct | Sonic | Marshal | 5572 | 2718 | 2 |
| Small Struct | BEVE | Unmarshal | 999 | 888 | 4 |
| Small Struct | Sonic | Unmarshal | 2209 | 1722 | 6 |
| Small Struct | MessagePack | Unmarshal | 2259 | 1280 | 29 |
| Small Struct | CBOR | Unmarshal | 6465 | 3912 | 83 |
| Small Struct | JSON | Unmarshal | 6797 | 1352 | 29 |
