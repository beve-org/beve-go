# Apple M1 (Virtual) — Darwin

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE ZeroCopy | Marshal | 84277 | 26 | 0 |
| Large Payload | BEVE | Marshal | 120374 | 188483 | 1 |
| Large Payload | CBOR | Marshal | 234669 | 204985 | 1 |
| Large Payload | MessagePack | Marshal | 366303 | 526758 | 115 |
| Large Payload | JSON | Marshal | 536479 | 221499 | 8 |
| Large Payload | Sonic | Marshal | 587467 | 214034 | 3 |
| Large Payload | BEVE | Unmarshal | 290562 | 273682 | 419 |
| Large Payload | Sonic | Unmarshal | 386632 | 337281 | 213 |
| Large Payload | MessagePack | Unmarshal | 674940 | 346613 | 6304 |
| Large Payload | CBOR | Unmarshal | 796452 | 310250 | 6324 |
| Large Payload | JSON | Unmarshal | 2508434 | 526996 | 6836 |
| Medium Payload | BEVE ZeroCopy | Marshal | 7893 | 4 | 0 |
| Medium Payload | BEVE | Marshal | 12221 | 16386 | 1 |
| Medium Payload | CBOR | Marshal | 19333 | 14345 | 1 |
| Medium Payload | MessagePack | Marshal | 33487 | 65778 | 22 |
| Medium Payload | JSON | Marshal | 53018 | 21993 | 8 |
| Medium Payload | Sonic | Marshal | 60225 | 24830 | 3 |
| Medium Payload | BEVE | Unmarshal | 25482 | 23515 | 59 |
| Medium Payload | Sonic | Unmarshal | 40998 | 35699 | 33 |
| Medium Payload | MessagePack | Unmarshal | 65055 | 41726 | 788 |
| Medium Payload | CBOR | Unmarshal | 86362 | 36488 | 750 |
| Medium Payload | JSON | Unmarshal | 196209 | 43792 | 567 |
| Small Struct | BEVE ZeroCopy | Marshal | 513 | 0 | 0 |
| Small Struct | BEVE | Marshal | 789 | 768 | 1 |
| Small Struct | MessagePack | Marshal | 1248 | 1032 | 6 |
| Small Struct | CBOR | Marshal | 2131 | 2048 | 1 |
| Small Struct | Sonic | Marshal | 3312 | 1168 | 2 |
| Small Struct | JSON | Marshal | 4002 | 2304 | 1 |
| Small Struct | BEVE | Unmarshal | 1441 | 2104 | 4 |
| Small Struct | MessagePack | Unmarshal | 1632 | 440 | 12 |
| Small Struct | CBOR | Unmarshal | 2370 | 904 | 22 |
| Small Struct | Sonic | Unmarshal | 3309 | 3372 | 6 |
| Small Struct | JSON | Unmarshal | 7519 | 1448 | 32 |
