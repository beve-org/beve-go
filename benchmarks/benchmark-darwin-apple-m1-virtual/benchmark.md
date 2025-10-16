# Apple M1 (Virtual) — Darwin

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE ZeroCopy | Marshal | 63789 | 79 | 0 |
| Large Payload | BEVE | Marshal | 101207 | 188469 | 1 |
| Large Payload | CBOR | Marshal | 195684 | 196758 | 1 |
| Large Payload | MessagePack | Marshal | 339807 | 526757 | 115 |
| Large Payload | JSON | Marshal | 442237 | 213306 | 8 |
| Large Payload | Sonic | Marshal | 556672 | 214615 | 3 |
| Large Payload | BEVE | Unmarshal | 209520 | 275953 | 418 |
| Large Payload | Sonic | Unmarshal | 342305 | 379902 | 213 |
| Large Payload | CBOR | Unmarshal | 531409 | 291274 | 5938 |
| Large Payload | MessagePack | Unmarshal | 683294 | 359110 | 6562 |
| Large Payload | JSON | Unmarshal | 2067981 | 516996 | 6814 |
| Medium Payload | BEVE ZeroCopy | Marshal | 6212 | 6 | 0 |
| Medium Payload | BEVE | Marshal | 12038 | 20486 | 1 |
| Medium Payload | CBOR | Marshal | 19160 | 21777 | 1 |
| Medium Payload | MessagePack | Marshal | 21639 | 33005 | 21 |
| Medium Payload | JSON | Marshal | 49123 | 24817 | 8 |
| Medium Payload | Sonic | Marshal | 54761 | 24839 | 3 |
| Medium Payload | BEVE | Unmarshal | 20585 | 29533 | 59 |
| Medium Payload | Sonic | Unmarshal | 26509 | 25093 | 33 |
| Medium Payload | MessagePack | Unmarshal | 53011 | 37086 | 691 |
| Medium Payload | CBOR | Unmarshal | 69600 | 30664 | 631 |
| Medium Payload | JSON | Unmarshal | 219624 | 52904 | 745 |
| Small Struct | BEVE ZeroCopy | Marshal | 537 | 0 | 0 |
| Small Struct | JSON | Marshal | 856 | 448 | 1 |
| Small Struct | BEVE | Marshal | 1810 | 2689 | 1 |
| Small Struct | CBOR | Marshal | 2062 | 2048 | 1 |
| Small Struct | MessagePack | Marshal | 2610 | 2056 | 7 |
| Small Struct | Sonic | Marshal | 4135 | 1851 | 2 |
| Small Struct | BEVE | Unmarshal | 523 | 312 | 3 |
| Small Struct | CBOR | Unmarshal | 3138 | 2120 | 47 |
| Small Struct | Sonic | Unmarshal | 3171 | 4140 | 6 |
| Small Struct | MessagePack | Unmarshal | 8280 | 4216 | 88 |
| Small Struct | JSON | Unmarshal | 30383 | 7880 | 112 |
