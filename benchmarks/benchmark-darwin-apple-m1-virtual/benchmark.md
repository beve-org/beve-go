# Apple M1 (Virtual) — Darwin

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE ZeroCopy | Marshal | 74265 | 26 | 0 |
| Large Payload | BEVE | Marshal | 156838 | 196687 | 1 |
| Large Payload | MessagePack | Marshal | 217022 | 526752 | 115 |
| Large Payload | CBOR | Marshal | 227978 | 196729 | 1 |
| Large Payload | JSON | Marshal | 404672 | 213276 | 8 |
| Large Payload | Sonic | Marshal | 422698 | 213678 | 3 |
| Large Payload | BEVE | Unmarshal | 186303 | 266417 | 418 |
| Large Payload | Sonic | Unmarshal | 450653 | 356120 | 209 |
| Large Payload | MessagePack | Unmarshal | 528440 | 356613 | 6506 |
| Large Payload | CBOR | Unmarshal | 777057 | 309738 | 6327 |
| Large Payload | JSON | Unmarshal | 1977073 | 528090 | 6970 |
| Medium Payload | BEVE ZeroCopy | Marshal | 6298 | 3 | 0 |
| Medium Payload | BEVE | Marshal | 11440 | 18441 | 1 |
| Medium Payload | CBOR | Marshal | 20835 | 20494 | 1 |
| Medium Payload | MessagePack | Marshal | 35397 | 65778 | 22 |
| Medium Payload | JSON | Marshal | 40660 | 24806 | 8 |
| Medium Payload | Sonic | Marshal | 41574 | 20764 | 3 |
| Medium Payload | BEVE | Unmarshal | 26169 | 26908 | 59 |
| Medium Payload | Sonic | Unmarshal | 33000 | 36391 | 33 |
| Medium Payload | MessagePack | Unmarshal | 38483 | 28124 | 509 |
| Medium Payload | CBOR | Unmarshal | 69886 | 30648 | 631 |
| Medium Payload | JSON | Unmarshal | 206455 | 49640 | 644 |
| Small Struct | BEVE ZeroCopy | Marshal | 329 | 0 | 0 |
| Small Struct | BEVE | Marshal | 925 | 1536 | 1 |
| Small Struct | JSON | Marshal | 1097 | 768 | 1 |
| Small Struct | CBOR | Marshal | 1372 | 2304 | 1 |
| Small Struct | MessagePack | Marshal | 1490 | 4104 | 8 |
| Small Struct | Sonic | Marshal | 2416 | 1438 | 2 |
| Small Struct | BEVE | Unmarshal | 953 | 2360 | 4 |
| Small Struct | MessagePack | Unmarshal | 2015 | 2112 | 46 |
| Small Struct | Sonic | Unmarshal | 2490 | 4035 | 6 |
| Small Struct | CBOR | Unmarshal | 4260 | 3656 | 79 |
| Small Struct | JSON | Unmarshal | 17019 | 7272 | 93 |
