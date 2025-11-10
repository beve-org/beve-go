# Apple M1 (Virtual) — Darwin

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE ZeroCopy | Marshal | 66707 | 52 | 0 |
| Large Payload | BEVE | Marshal | 95246 | 188456 | 1 |
| Large Payload | CBOR | Marshal | 160756 | 204938 | 1 |
| Large Payload | MessagePack | Marshal | 199083 | 526753 | 115 |
| Large Payload | JSON | Marshal | 384615 | 213279 | 8 |
| Large Payload | Sonic | Marshal | 452693 | 213905 | 3 |
| Large Payload | BEVE | Unmarshal | 180274 | 274483 | 418 |
| Large Payload | Sonic | Unmarshal | 331812 | 356656 | 211 |
| Large Payload | MessagePack | Unmarshal | 453573 | 344804 | 6261 |
| Large Payload | CBOR | Unmarshal | 622505 | 326874 | 6662 |
| Large Payload | JSON | Unmarshal | 1906647 | 527387 | 6848 |
| Medium Payload | BEVE ZeroCopy | Marshal | 5806 | 3 | 0 |
| Medium Payload | BEVE | Marshal | 10258 | 21767 | 1 |
| Medium Payload | CBOR | Marshal | 20087 | 21778 | 1 |
| Medium Payload | MessagePack | Marshal | 20336 | 33005 | 21 |
| Medium Payload | JSON | Marshal | 32112 | 20710 | 8 |
| Medium Payload | Sonic | Marshal | 46918 | 24868 | 3 |
| Medium Payload | BEVE | Unmarshal | 20355 | 29821 | 59 |
| Medium Payload | Sonic | Unmarshal | 31972 | 37710 | 33 |
| Medium Payload | MessagePack | Unmarshal | 50071 | 37190 | 688 |
| Medium Payload | CBOR | Unmarshal | 50180 | 26664 | 552 |
| Medium Payload | JSON | Unmarshal | 169756 | 49295 | 634 |
| Small Struct | BEVE ZeroCopy | Marshal | 209 | 0 | 0 |
| Small Struct | JSON | Marshal | 398 | 224 | 1 |
| Small Struct | Sonic | Marshal | 639 | 279 | 2 |
| Small Struct | CBOR | Marshal | 729 | 1024 | 1 |
| Small Struct | BEVE | Marshal | 1145 | 2304 | 1 |
| Small Struct | MessagePack | Marshal | 3658 | 8201 | 9 |
| Small Struct | BEVE | Unmarshal | 1005 | 2360 | 4 |
| Small Struct | Sonic | Unmarshal | 2521 | 2123 | 6 |
| Small Struct | CBOR | Unmarshal | 3705 | 1632 | 37 |
| Small Struct | MessagePack | Unmarshal | 4767 | 5152 | 105 |
| Small Struct | JSON | Unmarshal | 14954 | 4552 | 79 |
