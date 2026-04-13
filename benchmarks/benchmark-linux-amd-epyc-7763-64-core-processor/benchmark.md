# AMD EPYC 7763 64-Core Processor — Linux

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE ZeroCopy | Marshal | 80130 | 26 | 0 |
| Large Payload | BEVE | Marshal | 118412 | 196719 | 1 |
| Large Payload | Sonic | Marshal | 159183 | 224041 | 3 |
| Large Payload | CBOR | Marshal | 235413 | 213149 | 1 |
| Large Payload | MessagePack | Marshal | 329062 | 526781 | 115 |
| Large Payload | JSON | Marshal | 431997 | 205090 | 8 |
| Large Payload | BEVE | Unmarshal | 248183 | 262496 | 418 |
| Large Payload | Sonic | Unmarshal | 372463 | 525462 | 572 |
| Large Payload | MessagePack | Unmarshal | 576413 | 344663 | 6268 |
| Large Payload | CBOR | Unmarshal | 722086 | 312730 | 6385 |
| Large Payload | JSON | Unmarshal | 2321559 | 541514 | 7003 |
| Medium Payload | BEVE ZeroCopy | Marshal | 7843 | 3 | 0 |
| Medium Payload | BEVE | Marshal | 12388 | 19079 | 1 |
| Medium Payload | CBOR | Marshal | 19841 | 18453 | 1 |
| Medium Payload | Sonic | Marshal | 20178 | 28123 | 3 |
| Medium Payload | MessagePack | Marshal | 39020 | 65782 | 22 |
| Medium Payload | JSON | Marshal | 41736 | 20711 | 8 |
| Medium Payload | BEVE | Unmarshal | 23625 | 24990 | 59 |
| Medium Payload | Sonic | Unmarshal | 48325 | 67402 | 79 |
| Medium Payload | MessagePack | Unmarshal | 53677 | 33928 | 623 |
| Medium Payload | CBOR | Unmarshal | 72629 | 33640 | 689 |
| Medium Payload | JSON | Unmarshal | 210074 | 48280 | 662 |
| Small Struct | BEVE ZeroCopy | Marshal | 525 | 0 | 0 |
| Small Struct | Sonic | Marshal | 891 | 1061 | 2 |
| Small Struct | BEVE | Marshal | 1524 | 2688 | 1 |
| Small Struct | CBOR | Marshal | 2766 | 2688 | 1 |
| Small Struct | MessagePack | Marshal | 2821 | 4104 | 8 |
| Small Struct | JSON | Marshal | 3694 | 2048 | 1 |
| Small Struct | CBOR | Unmarshal | 1156 | 224 | 7 |
| Small Struct | BEVE | Unmarshal | 2073 | 3384 | 4 |
| Small Struct | Sonic | Unmarshal | 2499 | 3793 | 9 |
| Small Struct | MessagePack | Unmarshal | 5631 | 4296 | 90 |
| Small Struct | JSON | Unmarshal | 5872 | 1256 | 26 |
