# Apple M1 (Virtual) — Darwin

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE ZeroCopy | Marshal | 64797 | 26 | 0 |
| Large Payload | BEVE | Marshal | 112776 | 213049 | 1 |
| Large Payload | CBOR | Marshal | 206179 | 180374 | 1 |
| Large Payload | MessagePack | Marshal | 265235 | 526756 | 115 |
| Large Payload | JSON | Marshal | 383841 | 196945 | 8 |
| Large Payload | Sonic | Marshal | 554921 | 222247 | 3 |
| Large Payload | BEVE | Unmarshal | 199011 | 252204 | 415 |
| Large Payload | Sonic | Unmarshal | 367838 | 374620 | 211 |
| Large Payload | MessagePack | Unmarshal | 467291 | 356616 | 6507 |
| Large Payload | CBOR | Unmarshal | 579761 | 276522 | 5635 |
| Large Payload | JSON | Unmarshal | 2301150 | 546763 | 7054 |
| Medium Payload | BEVE ZeroCopy | Marshal | 8389 | 0 | 0 |
| Medium Payload | BEVE | Marshal | 13170 | 27276 | 1 |
| Medium Payload | CBOR | Marshal | 18672 | 18448 | 1 |
| Medium Payload | MessagePack | Marshal | 24847 | 33005 | 21 |
| Medium Payload | JSON | Marshal | 44483 | 20706 | 8 |
| Medium Payload | Sonic | Marshal | 48065 | 20716 | 3 |
| Medium Payload | BEVE | Unmarshal | 21698 | 29468 | 59 |
| Medium Payload | Sonic | Unmarshal | 27115 | 28856 | 31 |
| Medium Payload | MessagePack | Unmarshal | 46661 | 32365 | 589 |
| Medium Payload | CBOR | Unmarshal | 69581 | 39176 | 806 |
| Medium Payload | JSON | Unmarshal | 204908 | 54312 | 684 |
| Small Struct | BEVE ZeroCopy | Marshal | 387 | 0 | 0 |
| Small Struct | BEVE | Marshal | 1207 | 1792 | 1 |
| Small Struct | Sonic | Marshal | 1599 | 407 | 2 |
| Small Struct | CBOR | Marshal | 2036 | 3072 | 1 |
| Small Struct | JSON | Marshal | 2803 | 1536 | 1 |
| Small Struct | MessagePack | Marshal | 2869 | 4104 | 8 |
| Small Struct | BEVE | Unmarshal | 1530 | 3000 | 4 |
| Small Struct | Sonic | Unmarshal | 3296 | 5304 | 6 |
| Small Struct | CBOR | Unmarshal | 3787 | 1928 | 43 |
| Small Struct | MessagePack | Unmarshal | 4108 | 3576 | 76 |
| Small Struct | JSON | Unmarshal | 15248 | 4072 | 64 |
