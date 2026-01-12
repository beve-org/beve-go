# Unknown CPU — Windows

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE ZeroCopy | Marshal | 77800 | 65 | 0 |
| Large Payload | BEVE | Marshal | 121321 | 196670 | 1 |
| Large Payload | Sonic | Marshal | 162315 | 206953 | 3 |
| Large Payload | CBOR | Marshal | 252173 | 196716 | 1 |
| Large Payload | MessagePack | Marshal | 286865 | 526708 | 115 |
| Large Payload | JSON | Marshal | 466130 | 205074 | 8 |
| Large Payload | BEVE | Unmarshal | 270459 | 264834 | 416 |
| Large Payload | Sonic | Unmarshal | 546648 | 528205 | 571 |
| Large Payload | MessagePack | Unmarshal | 712845 | 366858 | 6716 |
| Large Payload | CBOR | Unmarshal | 1078752 | 348457 | 7109 |
| Large Payload | JSON | Unmarshal | 2836034 | 535404 | 7010 |
| Medium Payload | BEVE ZeroCopy | Marshal | 8185 | 1 | 0 |
| Medium Payload | BEVE | Marshal | 11188 | 14338 | 1 |
| Medium Payload | Sonic | Marshal | 19424 | 24900 | 3 |
| Medium Payload | CBOR | Marshal | 21738 | 16389 | 1 |
| Medium Payload | MessagePack | Marshal | 37420 | 65772 | 22 |
| Medium Payload | JSON | Marshal | 43874 | 19308 | 8 |
| Medium Payload | BEVE | Unmarshal | 29532 | 28444 | 58 |
| Medium Payload | Sonic | Unmarshal | 44381 | 50294 | 71 |
| Medium Payload | MessagePack | Unmarshal | 62499 | 30155 | 549 |
| Medium Payload | CBOR | Unmarshal | 94342 | 34872 | 720 |
| Medium Payload | JSON | Unmarshal | 286739 | 60152 | 796 |
| Small Struct | BEVE | Marshal | 455 | 352 | 1 |
| Small Struct | BEVE ZeroCopy | Marshal | 508 | 0 | 0 |
| Small Struct | Sonic | Marshal | 1174 | 1224 | 2 |
| Small Struct | MessagePack | Marshal | 1755 | 2056 | 7 |
| Small Struct | CBOR | Marshal | 3107 | 3073 | 1 |
| Small Struct | JSON | Marshal | 5800 | 2688 | 1 |
| Small Struct | BEVE | Unmarshal | 1473 | 1464 | 4 |
| Small Struct | CBOR | Unmarshal | 2141 | 704 | 18 |
| Small Struct | MessagePack | Unmarshal | 5522 | 3544 | 75 |
| Small Struct | Sonic | Unmarshal | 6229 | 7755 | 10 |
| Small Struct | JSON | Unmarshal | 11426 | 2312 | 44 |
