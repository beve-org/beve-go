# Neoverse-N2 — Linux

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE ZeroCopy | Marshal | 70332 | 286 | 2 |
| Large Payload | BEVE | Marshal | 107962 | 181451 | 3 |
| Large Payload | CBOR | Marshal | 199036 | 206910 | 3 |
| Large Payload | MessagePack | Marshal | 270882 | 526860 | 115 |
| Large Payload | Sonic | Marshal | 291442 | 209041 | 4 |
| Large Payload | JSON | Marshal | 396194 | 223047 | 9 |
| Large Payload | BEVE | Unmarshal | 227218 | 271369 | 419 |
| Large Payload | Sonic | Unmarshal | 293417 | 390664 | 213 |
| Large Payload | MessagePack | Unmarshal | 529026 | 361759 | 6612 |
| Large Payload | CBOR | Unmarshal | 585172 | 278297 | 5673 |
| Large Payload | JSON | Unmarshal | 2073252 | 561476 | 7328 |
| Medium Payload | BEVE ZeroCopy | Marshal | 7422 | 141 | 2 |
| Medium Payload | BEVE | Marshal | 10113 | 18581 | 3 |
| Medium Payload | CBOR | Marshal | 20703 | 21876 | 2 |
| Medium Payload | Sonic | Marshal | 25870 | 18775 | 4 |
| Medium Payload | MessagePack | Marshal | 30870 | 65838 | 22 |
| Medium Payload | JSON | Marshal | 35487 | 20811 | 9 |
| Medium Payload | BEVE | Unmarshal | 25263 | 35519 | 59 |
| Medium Payload | Sonic | Unmarshal | 33264 | 48396 | 33 |
| Medium Payload | MessagePack | Unmarshal | 47557 | 31359 | 571 |
| Medium Payload | CBOR | Unmarshal | 75657 | 40584 | 833 |
| Medium Payload | JSON | Unmarshal | 164881 | 43808 | 572 |
| Small Struct | BEVE ZeroCopy | Marshal | 722 | 290 | 2 |
| Small Struct | Sonic | Marshal | 944 | 590 | 3 |
| Small Struct | JSON | Marshal | 1606 | 1040 | 2 |
| Small Struct | BEVE | Marshal | 1733 | 2980 | 3 |
| Small Struct | MessagePack | Marshal | 2187 | 4224 | 8 |
| Small Struct | CBOR | Marshal | 2578 | 3219 | 2 |
| Small Struct | BEVE | Unmarshal | 1477 | 2616 | 4 |
| Small Struct | CBOR | Unmarshal | 2128 | 856 | 21 |
| Small Struct | Sonic | Unmarshal | 3109 | 4897 | 6 |
| Small Struct | MessagePack | Unmarshal | 4110 | 3288 | 71 |
| Small Struct | JSON | Unmarshal | 6510 | 1448 | 32 |
