# Apple M1 (Virtual) — Darwin

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE ZeroCopy | Marshal | 50788 | 39 | 0 |
| Large Payload | BEVE | Marshal | 78664 | 180277 | 1 |
| Large Payload | CBOR | Marshal | 161557 | 213157 | 1 |
| Large Payload | MessagePack | Marshal | 228326 | 526755 | 115 |
| Large Payload | JSON | Marshal | 357268 | 221499 | 8 |
| Large Payload | Sonic | Marshal | 405580 | 197294 | 3 |
| Large Payload | BEVE | Unmarshal | 210868 | 278610 | 418 |
| Large Payload | Sonic | Unmarshal | 310000 | 359207 | 211 |
| Large Payload | MessagePack | Unmarshal | 402865 | 335683 | 6092 |
| Large Payload | CBOR | Unmarshal | 481598 | 308201 | 6286 |
| Large Payload | JSON | Unmarshal | 1803701 | 544282 | 7126 |
| Medium Payload | BEVE ZeroCopy | Marshal | 5940 | 0 | 0 |
| Medium Payload | BEVE | Marshal | 9539 | 20487 | 1 |
| Medium Payload | CBOR | Marshal | 17602 | 18444 | 1 |
| Medium Payload | MessagePack | Marshal | 24111 | 65778 | 22 |
| Medium Payload | JSON | Marshal | 29157 | 20710 | 8 |
| Medium Payload | Sonic | Marshal | 44908 | 24836 | 3 |
| Medium Payload | BEVE | Unmarshal | 16440 | 24764 | 59 |
| Medium Payload | Sonic | Unmarshal | 31252 | 41164 | 33 |
| Medium Payload | MessagePack | Unmarshal | 52281 | 34909 | 644 |
| Medium Payload | CBOR | Unmarshal | 61522 | 27368 | 565 |
| Medium Payload | JSON | Unmarshal | 212551 | 71000 | 943 |
| Small Struct | BEVE ZeroCopy | Marshal | 143 | 0 | 0 |
| Small Struct | BEVE | Marshal | 742 | 2048 | 1 |
| Small Struct | CBOR | Marshal | 1393 | 2048 | 1 |
| Small Struct | MessagePack | Marshal | 2373 | 4104 | 8 |
| Small Struct | Sonic | Marshal | 2591 | 1431 | 2 |
| Small Struct | JSON | Marshal | 2999 | 2305 | 1 |
| Small Struct | BEVE | Unmarshal | 1061 | 696 | 4 |
| Small Struct | MessagePack | Unmarshal | 1472 | 1472 | 33 |
| Small Struct | Sonic | Unmarshal | 2136 | 3477 | 6 |
| Small Struct | CBOR | Unmarshal | 5329 | 4648 | 98 |
| Small Struct | JSON | Unmarshal | 15093 | 4776 | 86 |
