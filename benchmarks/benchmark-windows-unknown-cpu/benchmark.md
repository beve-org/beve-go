# Unknown CPU — Windows

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE ZeroCopy | Marshal | 79781 | 65 | 0 |
| Large Payload | BEVE | Marshal | 120675 | 188504 | 1 |
| Large Payload | Sonic | Marshal | 168178 | 199206 | 3 |
| Large Payload | CBOR | Marshal | 270953 | 188555 | 1 |
| Large Payload | MessagePack | Marshal | 306517 | 526708 | 115 |
| Large Payload | JSON | Marshal | 453269 | 196881 | 8 |
| Large Payload | BEVE | Unmarshal | 329881 | 279301 | 418 |
| Large Payload | Sonic | Unmarshal | 513706 | 541462 | 574 |
| Large Payload | MessagePack | Unmarshal | 729643 | 343716 | 6246 |
| Large Payload | CBOR | Unmarshal | 847159 | 297994 | 6080 |
| Large Payload | JSON | Unmarshal | 2731779 | 546876 | 7164 |
| Medium Payload | BEVE ZeroCopy | Marshal | 8125 | 5 | 0 |
| Medium Payload | BEVE | Marshal | 13248 | 18439 | 1 |
| Medium Payload | Sonic | Marshal | 21278 | 27656 | 3 |
| Medium Payload | CBOR | Marshal | 22919 | 18436 | 1 |
| Medium Payload | MessagePack | Marshal | 39335 | 65772 | 22 |
| Medium Payload | JSON | Marshal | 48639 | 20709 | 8 |
| Medium Payload | BEVE | Unmarshal | 35367 | 30907 | 59 |
| Medium Payload | Sonic | Unmarshal | 61119 | 64701 | 75 |
| Medium Payload | MessagePack | Unmarshal | 71523 | 34557 | 638 |
| Medium Payload | CBOR | Unmarshal | 90697 | 31432 | 647 |
| Medium Payload | JSON | Unmarshal | 287751 | 50360 | 653 |
| Small Struct | BEVE ZeroCopy | Marshal | 345 | 0 | 0 |
| Small Struct | Sonic | Marshal | 467 | 293 | 2 |
| Small Struct | BEVE | Marshal | 2219 | 2688 | 1 |
| Small Struct | CBOR | Marshal | 2661 | 2688 | 1 |
| Small Struct | JSON | Marshal | 4723 | 2305 | 1 |
| Small Struct | MessagePack | Marshal | 5588 | 8200 | 9 |
| Small Struct | BEVE | Unmarshal | 1830 | 2360 | 4 |
| Small Struct | Sonic | Unmarshal | 2262 | 1888 | 8 |
| Small Struct | MessagePack | Unmarshal | 3581 | 2088 | 45 |
| Small Struct | CBOR | Unmarshal | 4897 | 1920 | 43 |
| Small Struct | JSON | Unmarshal | 9200 | 2024 | 35 |
