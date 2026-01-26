# Apple M1 (Virtual) — Darwin

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE ZeroCopy | Marshal | 69428 | 26 | 0 |
| Large Payload | BEVE | Marshal | 129040 | 188469 | 1 |
| Large Payload | CBOR | Marshal | 154567 | 188548 | 1 |
| Large Payload | MessagePack | Marshal | 429006 | 526747 | 115 |
| Large Payload | JSON | Marshal | 489726 | 205111 | 8 |
| Large Payload | Sonic | Marshal | 620065 | 206328 | 3 |
| Large Payload | BEVE | Unmarshal | 316742 | 294196 | 419 |
| Large Payload | Sonic | Unmarshal | 365710 | 366934 | 213 |
| Large Payload | MessagePack | Unmarshal | 402141 | 342276 | 6214 |
| Large Payload | CBOR | Unmarshal | 731172 | 332171 | 6772 |
| Large Payload | JSON | Unmarshal | 2531305 | 569163 | 7474 |
| Medium Payload | BEVE ZeroCopy | Marshal | 7285 | 0 | 0 |
| Medium Payload | BEVE | Marshal | 16651 | 21768 | 1 |
| Medium Payload | CBOR | Marshal | 26794 | 24591 | 1 |
| Medium Payload | MessagePack | Marshal | 30316 | 33005 | 21 |
| Medium Payload | Sonic | Marshal | 52929 | 19393 | 3 |
| Medium Payload | JSON | Marshal | 58466 | 24807 | 8 |
| Medium Payload | BEVE | Unmarshal | 31162 | 28252 | 59 |
| Medium Payload | MessagePack | Unmarshal | 42985 | 22395 | 388 |
| Medium Payload | Sonic | Unmarshal | 43853 | 42761 | 33 |
| Medium Payload | CBOR | Unmarshal | 73382 | 33992 | 695 |
| Medium Payload | JSON | Unmarshal | 231532 | 47944 | 642 |
| Small Struct | BEVE ZeroCopy | Marshal | 352 | 0 | 0 |
| Small Struct | CBOR | Marshal | 683 | 320 | 1 |
| Small Struct | BEVE | Marshal | 1533 | 2048 | 1 |
| Small Struct | MessagePack | Marshal | 1750 | 2056 | 7 |
| Small Struct | Sonic | Marshal | 1782 | 432 | 2 |
| Small Struct | JSON | Marshal | 3115 | 1408 | 1 |
| Small Struct | BEVE | Unmarshal | 845 | 408 | 4 |
| Small Struct | Sonic | Unmarshal | 1419 | 1170 | 6 |
| Small Struct | MessagePack | Unmarshal | 2433 | 1176 | 27 |
| Small Struct | CBOR | Unmarshal | 2937 | 1096 | 26 |
| Small Struct | JSON | Unmarshal | 9188 | 2248 | 42 |
