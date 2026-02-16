# Unknown CPU — Windows

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE ZeroCopy | Marshal | 87589 | 65 | 0 |
| Large Payload | BEVE | Marshal | 133350 | 196645 | 1 |
| Large Payload | Sonic | Marshal | 167823 | 215308 | 3 |
| Large Payload | CBOR | Marshal | 244016 | 188546 | 1 |
| Large Payload | MessagePack | Marshal | 290160 | 526709 | 115 |
| Large Payload | JSON | Marshal | 458839 | 196933 | 8 |
| Large Payload | BEVE | Unmarshal | 291493 | 274244 | 419 |
| Large Payload | Sonic | Unmarshal | 471632 | 549234 | 572 |
| Large Payload | MessagePack | Unmarshal | 689737 | 346259 | 6304 |
| Large Payload | CBOR | Unmarshal | 950270 | 298042 | 6085 |
| Large Payload | JSON | Unmarshal | 2749452 | 532084 | 6980 |
| Medium Payload | BEVE ZeroCopy | Marshal | 7178 | 1 | 0 |
| Medium Payload | BEVE | Marshal | 15208 | 19076 | 1 |
| Medium Payload | Sonic | Marshal | 19482 | 22114 | 3 |
| Medium Payload | CBOR | Marshal | 31947 | 24587 | 1 |
| Medium Payload | MessagePack | Marshal | 40522 | 65773 | 22 |
| Medium Payload | JSON | Marshal | 61044 | 27509 | 8 |
| Medium Payload | BEVE | Unmarshal | 34561 | 29533 | 59 |
| Medium Payload | Sonic | Unmarshal | 56250 | 63533 | 75 |
| Medium Payload | MessagePack | Unmarshal | 80900 | 39390 | 734 |
| Medium Payload | CBOR | Unmarshal | 91977 | 30792 | 635 |
| Medium Payload | JSON | Unmarshal | 254058 | 46920 | 627 |
| Small Struct | BEVE | Marshal | 546 | 416 | 1 |
| Small Struct | BEVE ZeroCopy | Marshal | 593 | 0 | 0 |
| Small Struct | MessagePack | Marshal | 1125 | 1032 | 6 |
| Small Struct | CBOR | Marshal | 2012 | 2048 | 1 |
| Small Struct | Sonic | Marshal | 2507 | 2753 | 2 |
| Small Struct | JSON | Marshal | 3745 | 1536 | 1 |
| Small Struct | BEVE | Unmarshal | 2594 | 3000 | 4 |
| Small Struct | Sonic | Unmarshal | 3619 | 3653 | 9 |
| Small Struct | MessagePack | Unmarshal | 7064 | 3832 | 80 |
| Small Struct | CBOR | Unmarshal | 7189 | 3144 | 67 |
| Small Struct | JSON | Unmarshal | 7567 | 1352 | 29 |
