# Apple M1 (Virtual) — Darwin

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE ZeroCopy | Marshal | 54493 | 39 | 0 |
| Large Payload | BEVE | Marshal | 82654 | 196662 | 1 |
| Large Payload | MessagePack | Marshal | 212988 | 526753 | 115 |
| Large Payload | CBOR | Marshal | 220727 | 204922 | 1 |
| Large Payload | JSON | Marshal | 371179 | 221472 | 8 |
| Large Payload | Sonic | Marshal | 445923 | 213849 | 3 |
| Large Payload | BEVE | Unmarshal | 222095 | 284052 | 418 |
| Large Payload | Sonic | Unmarshal | 472468 | 341114 | 209 |
| Large Payload | MessagePack | Unmarshal | 481230 | 349094 | 6354 |
| Large Payload | CBOR | Unmarshal | 891061 | 316937 | 6439 |
| Large Payload | JSON | Unmarshal | 1831812 | 506226 | 6541 |
| Medium Payload | BEVE ZeroCopy | Marshal | 6653 | 3 | 0 |
| Medium Payload | BEVE | Marshal | 12340 | 18436 | 1 |
| Medium Payload | CBOR | Marshal | 18438 | 21784 | 1 |
| Medium Payload | MessagePack | Marshal | 39594 | 65778 | 22 |
| Medium Payload | JSON | Marshal | 42703 | 20710 | 8 |
| Medium Payload | Sonic | Marshal | 51675 | 20712 | 3 |
| Medium Payload | BEVE | Unmarshal | 24410 | 27100 | 59 |
| Medium Payload | Sonic | Unmarshal | 33992 | 41622 | 31 |
| Medium Payload | MessagePack | Unmarshal | 37553 | 30092 | 548 |
| Medium Payload | CBOR | Unmarshal | 60216 | 37560 | 775 |
| Medium Payload | JSON | Unmarshal | 255939 | 65432 | 875 |
| Small Struct | BEVE | Marshal | 632 | 1152 | 1 |
| Small Struct | BEVE ZeroCopy | Marshal | 708 | 0 | 0 |
| Small Struct | MessagePack | Marshal | 1295 | 2056 | 7 |
| Small Struct | JSON | Marshal | 2015 | 1024 | 1 |
| Small Struct | CBOR | Marshal | 2125 | 2688 | 1 |
| Small Struct | Sonic | Marshal | 4333 | 2085 | 2 |
| Small Struct | BEVE | Unmarshal | 988 | 2360 | 4 |
| Small Struct | MessagePack | Unmarshal | 2844 | 2136 | 47 |
| Small Struct | Sonic | Unmarshal | 3263 | 4687 | 6 |
| Small Struct | CBOR | Unmarshal | 9308 | 2408 | 52 |
| Small Struct | JSON | Unmarshal | 25115 | 8040 | 117 |
