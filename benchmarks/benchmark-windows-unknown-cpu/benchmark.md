# Unknown CPU — Windows

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE ZeroCopy | Marshal | 78802 | 259 | 2 |
| Large Payload | BEVE | Marshal | 124850 | 180892 | 3 |
| Large Payload | Sonic | Marshal | 187884 | 218668 | 4 |
| Large Payload | CBOR | Marshal | 221958 | 198136 | 2 |
| Large Payload | MessagePack | Marshal | 329577 | 526766 | 115 |
| Large Payload | JSON | Marshal | 537825 | 215354 | 9 |
| Large Payload | BEVE | Unmarshal | 332514 | 281990 | 419 |
| Large Payload | Sonic | Unmarshal | 404301 | 497948 | 558 |
| Large Payload | MessagePack | Unmarshal | 651910 | 321663 | 5808 |
| Large Payload | CBOR | Unmarshal | 885135 | 311787 | 6361 |
| Large Payload | JSON | Unmarshal | 2820221 | 553355 | 7258 |
| Medium Payload | BEVE ZeroCopy | Marshal | 9094 | 138 | 2 |
| Medium Payload | BEVE | Marshal | 16640 | 20644 | 3 |
| Medium Payload | Sonic | Marshal | 22197 | 25071 | 4 |
| Medium Payload | CBOR | Marshal | 26454 | 24693 | 2 |
| Medium Payload | MessagePack | Marshal | 43845 | 65830 | 22 |
| Medium Payload | JSON | Marshal | 51138 | 20837 | 9 |
| Medium Payload | BEVE | Unmarshal | 29081 | 23675 | 59 |
| Medium Payload | Sonic | Unmarshal | 52308 | 62269 | 74 |
| Medium Payload | MessagePack | Unmarshal | 77761 | 38574 | 724 |
| Medium Payload | CBOR | Unmarshal | 102687 | 36600 | 751 |
| Medium Payload | JSON | Unmarshal | 236646 | 44880 | 628 |
| Small Struct | BEVE ZeroCopy | Marshal | 490 | 290 | 2 |
| Small Struct | BEVE | Marshal | 2029 | 2339 | 3 |
| Small Struct | CBOR | Marshal | 2159 | 1680 | 2 |
| Small Struct | Sonic | Marshal | 2901 | 3374 | 3 |
| Small Struct | JSON | Marshal | 3856 | 1682 | 2 |
| Small Struct | MessagePack | Marshal | 4003 | 4224 | 8 |
| Small Struct | BEVE | Unmarshal | 2697 | 3512 | 4 |
| Small Struct | MessagePack | Unmarshal | 2870 | 1280 | 29 |
| Small Struct | Sonic | Unmarshal | 4312 | 4400 | 9 |
| Small Struct | CBOR | Unmarshal | 9985 | 4008 | 86 |
| Small Struct | JSON | Unmarshal | 15068 | 3848 | 57 |
