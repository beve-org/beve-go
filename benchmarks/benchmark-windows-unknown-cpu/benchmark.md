# Unknown CPU — Windows

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE ZeroCopy | Marshal | 75257 | 259 | 2 |
| Large Payload | BEVE | Marshal | 117585 | 180836 | 3 |
| Large Payload | Sonic | Marshal | 171858 | 237046 | 4 |
| Large Payload | CBOR | Marshal | 227698 | 205855 | 2 |
| Large Payload | MessagePack | Marshal | 306947 | 526768 | 115 |
| Large Payload | JSON | Marshal | 526218 | 224228 | 9 |
| Large Payload | BEVE | Unmarshal | 294377 | 284998 | 418 |
| Large Payload | Sonic | Unmarshal | 420248 | 536552 | 566 |
| Large Payload | MessagePack | Unmarshal | 666410 | 342371 | 6231 |
| Large Payload | CBOR | Unmarshal | 853149 | 313675 | 6385 |
| Large Payload | JSON | Unmarshal | 2584664 | 532123 | 6830 |
| Medium Payload | BEVE ZeroCopy | Marshal | 7128 | 138 | 2 |
| Medium Payload | BEVE | Marshal | 15186 | 21892 | 3 |
| Medium Payload | Sonic | Marshal | 20225 | 25229 | 4 |
| Medium Payload | CBOR | Marshal | 22280 | 18522 | 2 |
| Medium Payload | MessagePack | Marshal | 36412 | 65829 | 22 |
| Medium Payload | JSON | Marshal | 46236 | 20809 | 9 |
| Medium Payload | BEVE | Unmarshal | 28485 | 29083 | 59 |
| Medium Payload | Sonic | Unmarshal | 46454 | 56080 | 71 |
| Medium Payload | MessagePack | Unmarshal | 60495 | 29036 | 524 |
| Medium Payload | CBOR | Unmarshal | 87845 | 31048 | 642 |
| Medium Payload | JSON | Unmarshal | 282695 | 61128 | 784 |
| Small Struct | BEVE ZeroCopy | Marshal | 580 | 290 | 2 |
| Small Struct | CBOR | Marshal | 628 | 496 | 2 |
| Small Struct | Sonic | Marshal | 993 | 1113 | 3 |
| Small Struct | BEVE | Marshal | 1453 | 1698 | 3 |
| Small Struct | MessagePack | Marshal | 2947 | 4224 | 8 |
| Small Struct | JSON | Marshal | 5826 | 2833 | 2 |
| Small Struct | BEVE | Unmarshal | 1001 | 728 | 4 |
| Small Struct | Sonic | Unmarshal | 2556 | 2363 | 8 |
| Small Struct | CBOR | Unmarshal | 3756 | 1344 | 31 |
| Small Struct | MessagePack | Unmarshal | 3871 | 1920 | 42 |
| Small Struct | JSON | Unmarshal | 8993 | 2056 | 36 |
