# Apple M1 (Virtual) — Darwin

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE ZeroCopy | Marshal | 77594 | 39 | 0 |
| Large Payload | BEVE | Marshal | 144313 | 196675 | 1 |
| Large Payload | CBOR | Marshal | 262537 | 204931 | 1 |
| Large Payload | MessagePack | Marshal | 431270 | 526757 | 115 |
| Large Payload | JSON | Marshal | 523181 | 205165 | 8 |
| Large Payload | Sonic | Marshal | 634278 | 214323 | 3 |
| Large Payload | BEVE | Unmarshal | 304874 | 278258 | 419 |
| Large Payload | Sonic | Unmarshal | 405677 | 327623 | 207 |
| Large Payload | MessagePack | Unmarshal | 660024 | 354391 | 6473 |
| Large Payload | CBOR | Unmarshal | 752458 | 312858 | 6367 |
| Large Payload | JSON | Unmarshal | 2537211 | 554234 | 7306 |
| Medium Payload | BEVE ZeroCopy | Marshal | 8714 | 0 | 0 |
| Medium Payload | BEVE | Marshal | 16704 | 19076 | 1 |
| Medium Payload | CBOR | Marshal | 29488 | 21774 | 1 |
| Medium Payload | MessagePack | Marshal | 37518 | 65778 | 22 |
| Medium Payload | JSON | Marshal | 46238 | 21987 | 8 |
| Medium Payload | Sonic | Marshal | 58035 | 24955 | 3 |
| Medium Payload | BEVE | Unmarshal | 28654 | 25724 | 59 |
| Medium Payload | Sonic | Unmarshal | 47028 | 43547 | 33 |
| Medium Payload | MessagePack | Unmarshal | 62965 | 29805 | 542 |
| Medium Payload | CBOR | Unmarshal | 92434 | 35288 | 726 |
| Medium Payload | JSON | Unmarshal | 291367 | 72952 | 933 |
| Small Struct | BEVE ZeroCopy | Marshal | 186 | 0 | 0 |
| Small Struct | Sonic | Marshal | 952 | 215 | 2 |
| Small Struct | JSON | Marshal | 1614 | 1024 | 1 |
| Small Struct | BEVE | Marshal | 1753 | 2304 | 1 |
| Small Struct | CBOR | Marshal | 2724 | 2689 | 1 |
| Small Struct | MessagePack | Marshal | 3866 | 8201 | 9 |
| Small Struct | BEVE | Unmarshal | 1044 | 1208 | 4 |
| Small Struct | Sonic | Unmarshal | 2276 | 2621 | 6 |
| Small Struct | MessagePack | Unmarshal | 2330 | 1216 | 28 |
| Small Struct | CBOR | Unmarshal | 9171 | 4264 | 90 |
| Small Struct | JSON | Unmarshal | 28567 | 7304 | 94 |
