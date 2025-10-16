# Apple M1 (Virtual) — Darwin

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE ZeroCopy | Marshal | 55563 | 233 | 2 |
| Large Payload | BEVE | Marshal | 113718 | 197239 | 3 |
| Large Payload | CBOR | Marshal | 156003 | 189271 | 2 |
| Large Payload | MessagePack | Marshal | 181396 | 526814 | 115 |
| Large Payload | JSON | Marshal | 314182 | 197509 | 9 |
| Large Payload | Sonic | Marshal | 376348 | 214739 | 4 |
| Large Payload | BEVE | Unmarshal | 153048 | 263953 | 418 |
| Large Payload | Sonic | Unmarshal | 344074 | 335533 | 207 |
| Large Payload | MessagePack | Unmarshal | 448125 | 357109 | 6516 |
| Large Payload | CBOR | Unmarshal | 589305 | 325369 | 6625 |
| Large Payload | JSON | Unmarshal | 1627206 | 489259 | 6408 |
| Medium Payload | BEVE ZeroCopy | Marshal | 6877 | 134 | 2 |
| Medium Payload | CBOR | Marshal | 15224 | 18558 | 2 |
| Medium Payload | BEVE | Marshal | 15591 | 21898 | 3 |
| Medium Payload | MessagePack | Marshal | 29069 | 65834 | 22 |
| Medium Payload | JSON | Marshal | 36706 | 19376 | 9 |
| Medium Payload | Sonic | Marshal | 51114 | 27745 | 4 |
| Medium Payload | BEVE | Unmarshal | 22622 | 25692 | 58 |
| Medium Payload | Sonic | Unmarshal | 32900 | 37987 | 33 |
| Medium Payload | MessagePack | Unmarshal | 47911 | 37534 | 700 |
| Medium Payload | CBOR | Unmarshal | 51201 | 28744 | 593 |
| Medium Payload | JSON | Unmarshal | 224287 | 66248 | 866 |
| Small Struct | BEVE ZeroCopy | Marshal | 506 | 289 | 2 |
| Small Struct | BEVE | Marshal | 695 | 1568 | 3 |
| Small Struct | MessagePack | Marshal | 1834 | 4224 | 8 |
| Small Struct | CBOR | Marshal | 1861 | 2834 | 2 |
| Small Struct | JSON | Marshal | 2286 | 1680 | 2 |
| Small Struct | Sonic | Marshal | 5453 | 3261 | 3 |
| Small Struct | CBOR | Unmarshal | 897 | 232 | 7 |
| Small Struct | BEVE | Unmarshal | 1319 | 3000 | 4 |
| Small Struct | Sonic | Unmarshal | 3549 | 5967 | 6 |
| Small Struct | MessagePack | Unmarshal | 4304 | 2144 | 47 |
| Small Struct | JSON | Unmarshal | 23056 | 8040 | 117 |
