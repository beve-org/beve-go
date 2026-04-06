# Apple M1 (Virtual) — Darwin

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE ZeroCopy | Marshal | 55470 | 26 | 0 |
| Large Payload | BEVE | Marshal | 94073 | 188509 | 1 |
| Large Payload | CBOR | Marshal | 180150 | 188570 | 1 |
| Large Payload | MessagePack | Marshal | 214039 | 526756 | 115 |
| Large Payload | JSON | Marshal | 334983 | 213306 | 8 |
| Large Payload | Sonic | Marshal | 382089 | 197516 | 3 |
| Large Payload | BEVE | Unmarshal | 200371 | 266513 | 419 |
| Large Payload | Sonic | Unmarshal | 324395 | 323471 | 209 |
| Large Payload | MessagePack | Unmarshal | 593295 | 358778 | 6557 |
| Large Payload | CBOR | Unmarshal | 771394 | 339691 | 6916 |
| Large Payload | JSON | Unmarshal | 2030255 | 516267 | 6729 |
| Medium Payload | BEVE ZeroCopy | Marshal | 6519 | 0 | 0 |
| Medium Payload | BEVE | Marshal | 11902 | 16386 | 1 |
| Medium Payload | CBOR | Marshal | 14680 | 19081 | 1 |
| Medium Payload | MessagePack | Marshal | 32675 | 65778 | 22 |
| Medium Payload | JSON | Marshal | 40632 | 20710 | 8 |
| Medium Payload | Sonic | Marshal | 55639 | 21971 | 3 |
| Medium Payload | BEVE | Unmarshal | 26568 | 33117 | 59 |
| Medium Payload | Sonic | Unmarshal | 41470 | 42108 | 33 |
| Medium Payload | MessagePack | Unmarshal | 45031 | 36622 | 679 |
| Medium Payload | CBOR | Unmarshal | 52155 | 31160 | 644 |
| Medium Payload | JSON | Unmarshal | 242719 | 62328 | 815 |
| Small Struct | BEVE ZeroCopy | Marshal | 538 | 0 | 0 |
| Small Struct | CBOR | Marshal | 620 | 640 | 1 |
| Small Struct | BEVE | Marshal | 1743 | 2048 | 1 |
| Small Struct | MessagePack | Marshal | 2571 | 4104 | 8 |
| Small Struct | Sonic | Marshal | 3534 | 1438 | 2 |
| Small Struct | JSON | Marshal | 5232 | 2688 | 1 |
| Small Struct | MessagePack | Unmarshal | 1373 | 640 | 16 |
| Small Struct | BEVE | Unmarshal | 1640 | 3000 | 4 |
| Small Struct | Sonic | Unmarshal | 4146 | 5292 | 6 |
| Small Struct | CBOR | Unmarshal | 4341 | 2120 | 47 |
| Small Struct | JSON | Unmarshal | 6412 | 1440 | 32 |
