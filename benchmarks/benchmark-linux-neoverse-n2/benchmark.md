# Neoverse-N2 — Linux

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE ZeroCopy | Marshal | 68751 | 65 | 0 |
| Large Payload | BEVE | Marshal | 99537 | 180369 | 1 |
| Large Payload | CBOR | Marshal | 194273 | 205214 | 1 |
| Large Payload | MessagePack | Marshal | 257768 | 526794 | 115 |
| Large Payload | Sonic | Marshal | 316630 | 231138 | 3 |
| Large Payload | JSON | Marshal | 382145 | 213316 | 8 |
| Large Payload | BEVE | Unmarshal | 228514 | 281326 | 418 |
| Large Payload | Sonic | Unmarshal | 295584 | 406589 | 209 |
| Large Payload | MessagePack | Unmarshal | 492801 | 326452 | 5897 |
| Large Payload | CBOR | Unmarshal | 665135 | 322426 | 6574 |
| Large Payload | JSON | Unmarshal | 1883676 | 504700 | 6624 |
| Medium Payload | BEVE ZeroCopy | Marshal | 6857 | 3 | 0 |
| Medium Payload | BEVE | Marshal | 9879 | 18441 | 1 |
| Medium Payload | CBOR | Marshal | 17302 | 18447 | 1 |
| Medium Payload | Sonic | Marshal | 26056 | 18827 | 3 |
| Medium Payload | MessagePack | Marshal | 31796 | 65783 | 22 |
| Medium Payload | JSON | Marshal | 42301 | 24807 | 8 |
| Medium Payload | BEVE | Unmarshal | 21851 | 26781 | 59 |
| Medium Payload | Sonic | Unmarshal | 31647 | 44490 | 33 |
| Medium Payload | MessagePack | Unmarshal | 60506 | 44785 | 840 |
| Medium Payload | CBOR | Unmarshal | 65301 | 32152 | 661 |
| Medium Payload | JSON | Unmarshal | 227017 | 68184 | 854 |
| Small Struct | BEVE | Marshal | 363 | 352 | 1 |
| Small Struct | Sonic | Marshal | 772 | 389 | 2 |
| Small Struct | BEVE ZeroCopy | Marshal | 790 | 0 | 0 |
| Small Struct | CBOR | Marshal | 890 | 768 | 1 |
| Small Struct | JSON | Marshal | 1273 | 640 | 1 |
| Small Struct | MessagePack | Marshal | 1588 | 2056 | 7 |
| Small Struct | BEVE | Unmarshal | 1197 | 1720 | 4 |
| Small Struct | MessagePack | Unmarshal | 1963 | 1096 | 25 |
| Small Struct | Sonic | Unmarshal | 2238 | 3215 | 6 |
| Small Struct | CBOR | Unmarshal | 4135 | 2248 | 49 |
| Small Struct | JSON | Unmarshal | 15946 | 4392 | 74 |
