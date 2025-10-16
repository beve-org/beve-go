# Apple M1 (Virtual) — Darwin

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE ZeroCopy | Marshal | 74787 | 233 | 2 |
| Large Payload | BEVE | Marshal | 136334 | 181355 | 3 |
| Large Payload | CBOR | Marshal | 189512 | 189374 | 2 |
| Large Payload | MessagePack | Marshal | 203392 | 526811 | 115 |
| Large Payload | JSON | Marshal | 378468 | 213842 | 9 |
| Large Payload | Sonic | Marshal | 413973 | 214227 | 4 |
| Large Payload | BEVE | Unmarshal | 172364 | 251212 | 419 |
| Large Payload | Sonic | Unmarshal | 271064 | 336501 | 211 |
| Large Payload | MessagePack | Unmarshal | 454066 | 355348 | 6498 |
| Large Payload | CBOR | Unmarshal | 555595 | 339947 | 6942 |
| Large Payload | JSON | Unmarshal | 1951902 | 524350 | 6967 |
| Medium Payload | BEVE ZeroCopy | Marshal | 6589 | 128 | 2 |
| Medium Payload | BEVE | Marshal | 13834 | 21919 | 3 |
| Medium Payload | CBOR | Marshal | 15665 | 14410 | 2 |
| Medium Payload | MessagePack | Marshal | 32318 | 65834 | 22 |
| Medium Payload | JSON | Marshal | 37077 | 20777 | 9 |
| Medium Payload | Sonic | Marshal | 39662 | 18741 | 4 |
| Medium Payload | BEVE | Unmarshal | 31044 | 23260 | 59 |
| Medium Payload | Sonic | Unmarshal | 39861 | 42487 | 33 |
| Medium Payload | MessagePack | Unmarshal | 55826 | 33197 | 608 |
| Medium Payload | CBOR | Unmarshal | 57943 | 31408 | 644 |
| Medium Payload | JSON | Unmarshal | 237600 | 53184 | 695 |
| Small Struct | JSON | Marshal | 721 | 560 | 2 |
| Small Struct | BEVE | Marshal | 995 | 1697 | 3 |
| Small Struct | BEVE ZeroCopy | Marshal | 1007 | 288 | 2 |
| Small Struct | CBOR | Marshal | 1941 | 1680 | 2 |
| Small Struct | MessagePack | Marshal | 3403 | 8321 | 9 |
| Small Struct | Sonic | Marshal | 5680 | 3303 | 3 |
| Small Struct | BEVE | Unmarshal | 1740 | 2616 | 4 |
| Small Struct | Sonic | Unmarshal | 3371 | 5071 | 6 |
| Small Struct | CBOR | Unmarshal | 3672 | 3080 | 65 |
| Small Struct | MessagePack | Unmarshal | 4625 | 3168 | 67 |
| Small Struct | JSON | Unmarshal | 6625 | 1320 | 28 |
