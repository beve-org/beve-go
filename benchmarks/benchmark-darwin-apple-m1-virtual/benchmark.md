# Apple M1 (Virtual) — Darwin

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE ZeroCopy | Marshal | 75636 | 39 | 0 |
| Large Payload | BEVE | Marshal | 121236 | 196659 | 1 |
| Large Payload | CBOR | Marshal | 156985 | 196726 | 1 |
| Large Payload | MessagePack | Marshal | 263223 | 526757 | 115 |
| Large Payload | JSON | Marshal | 442972 | 213278 | 8 |
| Large Payload | Sonic | Marshal | 487161 | 222405 | 3 |
| Large Payload | BEVE | Unmarshal | 244614 | 266191 | 417 |
| Large Payload | Sonic | Unmarshal | 394040 | 353946 | 211 |
| Large Payload | MessagePack | Unmarshal | 418817 | 348995 | 6360 |
| Large Payload | CBOR | Unmarshal | 730138 | 325866 | 6654 |
| Large Payload | JSON | Unmarshal | 1789388 | 495825 | 6406 |
| Medium Payload | BEVE ZeroCopy | Marshal | 6785 | 0 | 0 |
| Medium Payload | CBOR | Marshal | 14282 | 16397 | 1 |
| Medium Payload | BEVE | Marshal | 17417 | 24585 | 1 |
| Medium Payload | MessagePack | Marshal | 27678 | 65778 | 22 |
| Medium Payload | Sonic | Marshal | 39261 | 19301 | 3 |
| Medium Payload | JSON | Marshal | 40859 | 21990 | 8 |
| Medium Payload | BEVE | Unmarshal | 18135 | 25820 | 59 |
| Medium Payload | Sonic | Unmarshal | 34543 | 37736 | 33 |
| Medium Payload | MessagePack | Unmarshal | 43969 | 40542 | 758 |
| Medium Payload | CBOR | Unmarshal | 48306 | 29896 | 614 |
| Medium Payload | JSON | Unmarshal | 227052 | 72152 | 934 |
| Small Struct | BEVE ZeroCopy | Marshal | 550 | 0 | 0 |
| Small Struct | BEVE | Marshal | 756 | 1280 | 1 |
| Small Struct | CBOR | Marshal | 840 | 896 | 1 |
| Small Struct | MessagePack | Marshal | 2708 | 4104 | 8 |
| Small Struct | Sonic | Marshal | 4279 | 1822 | 2 |
| Small Struct | JSON | Marshal | 5979 | 2688 | 1 |
| Small Struct | BEVE | Unmarshal | 2072 | 3000 | 4 |
| Small Struct | MessagePack | Unmarshal | 2796 | 1920 | 42 |
| Small Struct | Sonic | Unmarshal | 3514 | 4093 | 6 |
| Small Struct | CBOR | Unmarshal | 5957 | 3208 | 69 |
| Small Struct | JSON | Unmarshal | 12528 | 3832 | 56 |
