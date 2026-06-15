# Unknown CPU — Windows

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE ZeroCopy | Marshal | 81929 | 65 | 0 |
| Large Payload | BEVE | Marshal | 108529 | 180257 | 1 |
| Large Payload | Sonic | Marshal | 154045 | 198210 | 3 |
| Large Payload | CBOR | Marshal | 216929 | 196704 | 1 |
| Large Payload | MessagePack | Marshal | 279249 | 526703 | 115 |
| Large Payload | JSON | Marshal | 475738 | 213291 | 8 |
| Large Payload | BEVE | Unmarshal | 273327 | 267203 | 417 |
| Large Payload | Sonic | Unmarshal | 453837 | 580714 | 595 |
| Large Payload | MessagePack | Unmarshal | 686261 | 361510 | 6631 |
| Large Payload | CBOR | Unmarshal | 852600 | 310874 | 6323 |
| Large Payload | JSON | Unmarshal | 2768864 | 564364 | 7347 |
| Medium Payload | BEVE ZeroCopy | Marshal | 7961 | 3 | 0 |
| Medium Payload | BEVE | Marshal | 12684 | 18434 | 1 |
| Medium Payload | Sonic | Marshal | 19856 | 24974 | 3 |
| Medium Payload | CBOR | Marshal | 28630 | 27270 | 1 |
| Medium Payload | MessagePack | Marshal | 33383 | 65770 | 22 |
| Medium Payload | JSON | Marshal | 40366 | 18658 | 8 |
| Medium Payload | BEVE | Unmarshal | 28493 | 29851 | 59 |
| Medium Payload | Sonic | Unmarshal | 46567 | 57383 | 70 |
| Medium Payload | MessagePack | Unmarshal | 73223 | 39165 | 728 |
| Medium Payload | CBOR | Unmarshal | 88393 | 33928 | 699 |
| Medium Payload | JSON | Unmarshal | 228074 | 46008 | 596 |
| Small Struct | BEVE ZeroCopy | Marshal | 565 | 0 | 0 |
| Small Struct | Sonic | Marshal | 1068 | 1324 | 2 |
| Small Struct | BEVE | Marshal | 1104 | 1280 | 1 |
| Small Struct | CBOR | Marshal | 1368 | 1280 | 1 |
| Small Struct | MessagePack | Marshal | 4716 | 8200 | 9 |
| Small Struct | JSON | Marshal | 5346 | 2689 | 1 |
| Small Struct | BEVE | Unmarshal | 2250 | 3384 | 4 |
| Small Struct | MessagePack | Unmarshal | 3683 | 1592 | 36 |
| Small Struct | Sonic | Unmarshal | 3741 | 4421 | 9 |
| Small Struct | JSON | Unmarshal | 6159 | 1224 | 25 |
| Small Struct | CBOR | Unmarshal | 10188 | 4712 | 100 |
