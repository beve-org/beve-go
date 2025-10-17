# Neoverse-N2 — Linux

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE ZeroCopy | Marshal | 68269 | 65 | 0 |
| Large Payload | BEVE | Marshal | 103251 | 180355 | 1 |
| Large Payload | CBOR | Marshal | 170262 | 172454 | 1 |
| Large Payload | MessagePack | Marshal | 274553 | 526804 | 115 |
| Large Payload | Sonic | Marshal | 317780 | 223634 | 3 |
| Large Payload | JSON | Marshal | 380397 | 205597 | 8 |
| Large Payload | BEVE | Unmarshal | 230095 | 270474 | 417 |
| Large Payload | Sonic | Unmarshal | 286179 | 393343 | 213 |
| Large Payload | MessagePack | Unmarshal | 527263 | 353311 | 6448 |
| Large Payload | CBOR | Unmarshal | 637669 | 306810 | 6254 |
| Large Payload | JSON | Unmarshal | 2102624 | 576139 | 7475 |
| Medium Payload | BEVE ZeroCopy | Marshal | 6807 | 7 | 0 |
| Medium Payload | BEVE | Marshal | 9343 | 16387 | 1 |
| Medium Payload | CBOR | Marshal | 16890 | 18447 | 1 |
| Medium Payload | Sonic | Marshal | 25959 | 18932 | 3 |
| Medium Payload | MessagePack | Marshal | 31985 | 65782 | 22 |
| Medium Payload | JSON | Marshal | 40508 | 21987 | 8 |
| Medium Payload | BEVE | Unmarshal | 24151 | 30878 | 59 |
| Medium Payload | Sonic | Unmarshal | 32068 | 44474 | 33 |
| Medium Payload | MessagePack | Unmarshal | 60436 | 43649 | 818 |
| Medium Payload | CBOR | Unmarshal | 63424 | 31224 | 634 |
| Medium Payload | JSON | Unmarshal | 155831 | 40456 | 529 |
| Small Struct | BEVE ZeroCopy | Marshal | 388 | 0 | 0 |
| Small Struct | BEVE | Marshal | 694 | 1024 | 1 |
| Small Struct | MessagePack | Marshal | 2295 | 4104 | 8 |
| Small Struct | CBOR | Marshal | 2401 | 3072 | 1 |
| Small Struct | Sonic | Marshal | 3220 | 2362 | 2 |
| Small Struct | JSON | Marshal | 4784 | 3073 | 1 |
| Small Struct | BEVE | Unmarshal | 805 | 600 | 4 |
| Small Struct | Sonic | Unmarshal | 3657 | 6282 | 6 |
| Small Struct | MessagePack | Unmarshal | 5688 | 4760 | 101 |
| Small Struct | CBOR | Unmarshal | 7926 | 4808 | 103 |
| Small Struct | JSON | Unmarshal | 8069 | 2152 | 39 |
