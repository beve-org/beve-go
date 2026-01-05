# Neoverse-N2 — Linux

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE ZeroCopy | Marshal | 67982 | 65 | 0 |
| Large Payload | BEVE | Marshal | 105612 | 204921 | 1 |
| Large Payload | CBOR | Marshal | 192309 | 205001 | 1 |
| Large Payload | MessagePack | Marshal | 269065 | 526801 | 115 |
| Large Payload | Sonic | Marshal | 280458 | 198534 | 3 |
| Large Payload | JSON | Marshal | 380827 | 213369 | 8 |
| Large Payload | BEVE | Unmarshal | 216446 | 275148 | 417 |
| Large Payload | Sonic | Unmarshal | 271238 | 372788 | 213 |
| Large Payload | MessagePack | Unmarshal | 514508 | 358288 | 6544 |
| Large Payload | CBOR | Unmarshal | 637271 | 310585 | 6343 |
| Large Payload | JSON | Unmarshal | 1897280 | 510076 | 6746 |
| Medium Payload | BEVE ZeroCopy | Marshal | 6433 | 6 | 0 |
| Medium Payload | BEVE | Marshal | 9499 | 18439 | 1 |
| Medium Payload | CBOR | Marshal | 19435 | 21780 | 1 |
| Medium Payload | MessagePack | Marshal | 22303 | 33007 | 21 |
| Medium Payload | Sonic | Marshal | 31347 | 25045 | 3 |
| Medium Payload | JSON | Marshal | 33485 | 18659 | 8 |
| Medium Payload | BEVE | Unmarshal | 22667 | 30878 | 59 |
| Medium Payload | Sonic | Unmarshal | 34552 | 50749 | 33 |
| Medium Payload | MessagePack | Unmarshal | 56162 | 40976 | 772 |
| Medium Payload | CBOR | Unmarshal | 63517 | 31224 | 642 |
| Medium Payload | JSON | Unmarshal | 220867 | 63000 | 835 |
| Small Struct | BEVE ZeroCopy | Marshal | 201 | 0 | 0 |
| Small Struct | CBOR | Marshal | 592 | 384 | 1 |
| Small Struct | BEVE | Marshal | 711 | 1024 | 1 |
| Small Struct | Sonic | Marshal | 1130 | 663 | 2 |
| Small Struct | JSON | Marshal | 1456 | 704 | 1 |
| Small Struct | MessagePack | Marshal | 2444 | 4104 | 8 |
| Small Struct | BEVE | Unmarshal | 1053 | 1336 | 4 |
| Small Struct | Sonic | Unmarshal | 1421 | 1643 | 6 |
| Small Struct | MessagePack | Unmarshal | 2922 | 2144 | 47 |
| Small Struct | CBOR | Unmarshal | 3841 | 2056 | 45 |
| Small Struct | JSON | Unmarshal | 23194 | 7688 | 106 |
