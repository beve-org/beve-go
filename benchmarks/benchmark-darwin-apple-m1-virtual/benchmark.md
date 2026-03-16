# Apple M1 (Virtual) — Darwin

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE ZeroCopy | Marshal | 53916 | 52 | 0 |
| Large Payload | BEVE | Marshal | 79556 | 204856 | 1 |
| Large Payload | CBOR | Marshal | 144053 | 204939 | 1 |
| Large Payload | MessagePack | Marshal | 180064 | 526757 | 115 |
| Large Payload | JSON | Marshal | 300747 | 221473 | 8 |
| Large Payload | Sonic | Marshal | 377985 | 221497 | 3 |
| Large Payload | BEVE | Unmarshal | 157326 | 275891 | 418 |
| Large Payload | Sonic | Unmarshal | 249176 | 342230 | 211 |
| Large Payload | MessagePack | Unmarshal | 370765 | 361480 | 6605 |
| Large Payload | CBOR | Unmarshal | 442151 | 295353 | 6029 |
| Large Payload | JSON | Unmarshal | 1519396 | 519153 | 6707 |
| Medium Payload | BEVE ZeroCopy | Marshal | 5392 | 0 | 0 |
| Medium Payload | BEVE | Marshal | 8850 | 21765 | 1 |
| Medium Payload | CBOR | Marshal | 12836 | 18444 | 1 |
| Medium Payload | MessagePack | Marshal | 22543 | 65778 | 22 |
| Medium Payload | Sonic | Marshal | 34289 | 20678 | 3 |
| Medium Payload | JSON | Marshal | 34887 | 24808 | 8 |
| Medium Payload | BEVE | Unmarshal | 15281 | 29917 | 59 |
| Medium Payload | Sonic | Unmarshal | 23847 | 34173 | 33 |
| Medium Payload | MessagePack | Unmarshal | 34530 | 35885 | 664 |
| Medium Payload | CBOR | Unmarshal | 47258 | 35495 | 732 |
| Medium Payload | JSON | Unmarshal | 160650 | 58952 | 762 |
| Small Struct | BEVE ZeroCopy | Marshal | 296 | 0 | 0 |
| Small Struct | BEVE | Marshal | 524 | 1408 | 1 |
| Small Struct | MessagePack | Marshal | 925 | 2056 | 7 |
| Small Struct | Sonic | Marshal | 1006 | 517 | 2 |
| Small Struct | CBOR | Marshal | 1438 | 2305 | 1 |
| Small Struct | JSON | Marshal | 2313 | 1792 | 1 |
| Small Struct | BEVE | Unmarshal | 466 | 504 | 4 |
| Small Struct | Sonic | Unmarshal | 1342 | 1457 | 6 |
| Small Struct | CBOR | Unmarshal | 1527 | 1064 | 25 |
| Small Struct | MessagePack | Unmarshal | 2762 | 3648 | 78 |
| Small Struct | JSON | Unmarshal | 8757 | 3784 | 55 |
