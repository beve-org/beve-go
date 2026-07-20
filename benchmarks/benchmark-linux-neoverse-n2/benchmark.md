# Neoverse-N2 — Linux

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE ZeroCopy | Marshal | 69865 | 65 | 0 |
| Large Payload | BEVE | Marshal | 106476 | 196756 | 1 |
| Large Payload | CBOR | Marshal | 192118 | 205083 | 1 |
| Large Payload | MessagePack | Marshal | 277205 | 526805 | 115 |
| Large Payload | Sonic | Marshal | 306827 | 217439 | 3 |
| Large Payload | JSON | Marshal | 396154 | 221720 | 8 |
| Large Payload | BEVE | Unmarshal | 218174 | 273037 | 416 |
| Large Payload | Sonic | Unmarshal | 289824 | 404816 | 211 |
| Large Payload | MessagePack | Unmarshal | 512539 | 357259 | 6521 |
| Large Payload | CBOR | Unmarshal | 634423 | 307978 | 6275 |
| Large Payload | JSON | Unmarshal | 2191666 | 601771 | 7990 |
| Medium Payload | BEVE ZeroCopy | Marshal | 7039 | 6 | 0 |
| Medium Payload | BEVE | Marshal | 11968 | 24583 | 1 |
| Medium Payload | CBOR | Marshal | 18578 | 20499 | 1 |
| Medium Payload | Sonic | Marshal | 30135 | 22123 | 3 |
| Medium Payload | MessagePack | Marshal | 30797 | 65782 | 22 |
| Medium Payload | JSON | Marshal | 35199 | 19302 | 8 |
| Medium Payload | BEVE | Unmarshal | 22057 | 28638 | 59 |
| Medium Payload | Sonic | Unmarshal | 27860 | 38026 | 33 |
| Medium Payload | CBOR | Unmarshal | 52938 | 24200 | 499 |
| Medium Payload | MessagePack | Unmarshal | 57008 | 41569 | 781 |
| Medium Payload | JSON | Unmarshal | 193062 | 54072 | 712 |
| Small Struct | BEVE ZeroCopy | Marshal | 575 | 0 | 0 |
| Small Struct | BEVE | Marshal | 1424 | 2688 | 1 |
| Small Struct | Sonic | Marshal | 2118 | 1473 | 2 |
| Small Struct | MessagePack | Marshal | 2301 | 4104 | 8 |
| Small Struct | CBOR | Marshal | 2413 | 3073 | 1 |
| Small Struct | JSON | Marshal | 2822 | 1536 | 1 |
| Small Struct | MessagePack | Unmarshal | 1181 | 352 | 10 |
| Small Struct | BEVE | Unmarshal | 1299 | 1848 | 4 |
| Small Struct | CBOR | Unmarshal | 2958 | 1384 | 32 |
| Small Struct | Sonic | Unmarshal | 2966 | 4874 | 6 |
| Small Struct | JSON | Unmarshal | 9412 | 2344 | 45 |
