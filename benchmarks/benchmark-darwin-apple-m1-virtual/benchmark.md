# Apple M1 (Virtual) — Darwin

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE ZeroCopy | Marshal | 91210 | 52 | 0 |
| Large Payload | BEVE | Marshal | 145435 | 180275 | 1 |
| Large Payload | CBOR | Marshal | 183189 | 188593 | 1 |
| Large Payload | MessagePack | Marshal | 283606 | 526754 | 115 |
| Large Payload | Sonic | Marshal | 495295 | 197688 | 3 |
| Large Payload | JSON | Marshal | 509434 | 221525 | 8 |
| Large Payload | BEVE | Unmarshal | 273246 | 277969 | 419 |
| Large Payload | Sonic | Unmarshal | 347350 | 360719 | 211 |
| Large Payload | MessagePack | Unmarshal | 599547 | 369946 | 6771 |
| Large Payload | CBOR | Unmarshal | 671635 | 333594 | 6804 |
| Large Payload | JSON | Unmarshal | 1995001 | 543748 | 7100 |
| Medium Payload | BEVE ZeroCopy | Marshal | 6612 | 0 | 0 |
| Medium Payload | BEVE | Marshal | 8173 | 18434 | 1 |
| Medium Payload | CBOR | Marshal | 29886 | 18448 | 1 |
| Medium Payload | MessagePack | Marshal | 36421 | 65778 | 22 |
| Medium Payload | JSON | Marshal | 37820 | 20706 | 8 |
| Medium Payload | Sonic | Marshal | 53088 | 20706 | 3 |
| Medium Payload | BEVE | Unmarshal | 28998 | 32029 | 59 |
| Medium Payload | Sonic | Unmarshal | 37066 | 32257 | 33 |
| Medium Payload | MessagePack | Unmarshal | 72956 | 40863 | 763 |
| Medium Payload | CBOR | Unmarshal | 74931 | 33192 | 682 |
| Medium Payload | JSON | Unmarshal | 256555 | 57752 | 756 |
| Small Struct | CBOR | Marshal | 251 | 144 | 1 |
| Small Struct | BEVE ZeroCopy | Marshal | 392 | 0 | 0 |
| Small Struct | BEVE | Marshal | 394 | 896 | 1 |
| Small Struct | JSON | Marshal | 2331 | 1792 | 1 |
| Small Struct | MessagePack | Marshal | 2961 | 8201 | 9 |
| Small Struct | Sonic | Marshal | 4779 | 2341 | 2 |
| Small Struct | BEVE | Unmarshal | 1017 | 2104 | 4 |
| Small Struct | MessagePack | Unmarshal | 1186 | 688 | 17 |
| Small Struct | CBOR | Unmarshal | 1552 | 712 | 18 |
| Small Struct | Sonic | Unmarshal | 2915 | 4280 | 6 |
| Small Struct | JSON | Unmarshal | 10314 | 3688 | 52 |
