# Neoverse-N2 — Linux

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE ZeroCopy | Marshal | 70094 | 52 | 0 |
| Large Payload | BEVE | Marshal | 118600 | 188696 | 1 |
| Large Payload | CBOR | Marshal | 178499 | 180549 | 1 |
| Large Payload | MessagePack | Marshal | 312135 | 526813 | 115 |
| Large Payload | Sonic | Marshal | 325355 | 226503 | 3 |
| Large Payload | JSON | Marshal | 395449 | 213500 | 8 |
| Large Payload | BEVE | Unmarshal | 236000 | 269804 | 417 |
| Large Payload | Sonic | Unmarshal | 298188 | 400247 | 211 |
| Large Payload | MessagePack | Unmarshal | 503780 | 333001 | 6031 |
| Large Payload | CBOR | Unmarshal | 653662 | 312986 | 6362 |
| Large Payload | JSON | Unmarshal | 1882082 | 499483 | 6521 |
| Medium Payload | BEVE ZeroCopy | Marshal | 6952 | 5 | 0 |
| Medium Payload | BEVE | Marshal | 11388 | 20487 | 1 |
| Medium Payload | CBOR | Marshal | 17957 | 18453 | 1 |
| Medium Payload | Sonic | Marshal | 27655 | 19665 | 3 |
| Medium Payload | MessagePack | Marshal | 34468 | 65782 | 22 |
| Medium Payload | JSON | Marshal | 34497 | 18662 | 8 |
| Medium Payload | BEVE | Unmarshal | 22460 | 26077 | 58 |
| Medium Payload | Sonic | Unmarshal | 34843 | 50208 | 33 |
| Medium Payload | MessagePack | Unmarshal | 53137 | 35407 | 659 |
| Medium Payload | CBOR | Unmarshal | 78409 | 39784 | 814 |
| Medium Payload | JSON | Unmarshal | 234968 | 68344 | 875 |
| Small Struct | BEVE ZeroCopy | Marshal | 287 | 0 | 0 |
| Small Struct | BEVE | Marshal | 385 | 416 | 1 |
| Small Struct | CBOR | Marshal | 527 | 288 | 1 |
| Small Struct | Sonic | Marshal | 1440 | 933 | 2 |
| Small Struct | MessagePack | Marshal | 1542 | 2056 | 7 |
| Small Struct | JSON | Marshal | 1863 | 1024 | 1 |
| Small Struct | BEVE | Unmarshal | 1728 | 3000 | 4 |
| Small Struct | MessagePack | Unmarshal | 2687 | 1696 | 38 |
| Small Struct | CBOR | Unmarshal | 3148 | 1552 | 35 |
| Small Struct | Sonic | Unmarshal | 3325 | 5281 | 6 |
| Small Struct | JSON | Unmarshal | 7699 | 2088 | 37 |
