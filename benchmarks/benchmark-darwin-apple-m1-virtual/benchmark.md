# Apple M1 (Virtual) — Darwin

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE ZeroCopy | Marshal | 58262 | 26 | 0 |
| Large Payload | BEVE | Marshal | 81073 | 196703 | 1 |
| Large Payload | CBOR | Marshal | 162307 | 196764 | 1 |
| Large Payload | MessagePack | Marshal | 180899 | 526752 | 115 |
| Large Payload | JSON | Marshal | 341079 | 221473 | 8 |
| Large Payload | Sonic | Marshal | 471737 | 213672 | 3 |
| Large Payload | BEVE | Unmarshal | 185892 | 259247 | 418 |
| Large Payload | Sonic | Unmarshal | 333540 | 363016 | 213 |
| Large Payload | MessagePack | Unmarshal | 402670 | 352550 | 6434 |
| Large Payload | CBOR | Unmarshal | 461033 | 290729 | 5932 |
| Large Payload | JSON | Unmarshal | 2452359 | 554227 | 7185 |
| Medium Payload | BEVE ZeroCopy | Marshal | 6777 | 0 | 0 |
| Medium Payload | BEVE | Marshal | 7743 | 18436 | 1 |
| Medium Payload | CBOR | Marshal | 13950 | 19091 | 1 |
| Medium Payload | MessagePack | Marshal | 23947 | 65778 | 22 |
| Medium Payload | JSON | Marshal | 35582 | 24806 | 8 |
| Medium Payload | Sonic | Marshal | 40553 | 20664 | 3 |
| Medium Payload | BEVE | Unmarshal | 26416 | 28796 | 59 |
| Medium Payload | Sonic | Unmarshal | 28476 | 32059 | 33 |
| Medium Payload | MessagePack | Unmarshal | 41566 | 37989 | 709 |
| Medium Payload | CBOR | Unmarshal | 63752 | 31976 | 656 |
| Medium Payload | JSON | Unmarshal | 172971 | 51496 | 714 |
| Small Struct | BEVE ZeroCopy | Marshal | 243 | 0 | 0 |
| Small Struct | CBOR | Marshal | 402 | 384 | 1 |
| Small Struct | BEVE | Marshal | 950 | 2689 | 1 |
| Small Struct | Sonic | Marshal | 2563 | 1175 | 2 |
| Small Struct | MessagePack | Marshal | 2613 | 8201 | 9 |
| Small Struct | JSON | Marshal | 3389 | 2689 | 1 |
| Small Struct | BEVE | Unmarshal | 853 | 1336 | 4 |
| Small Struct | MessagePack | Unmarshal | 2195 | 2272 | 49 |
| Small Struct | Sonic | Unmarshal | 2277 | 3220 | 6 |
| Small Struct | CBOR | Unmarshal | 2938 | 2216 | 48 |
| Small Struct | JSON | Unmarshal | 15935 | 4264 | 70 |
