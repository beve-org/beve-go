# Unknown CPU — Windows

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE ZeroCopy | Marshal | 81019 | 79 | 0 |
| Large Payload | BEVE | Marshal | 111433 | 188476 | 1 |
| Large Payload | Sonic | Marshal | 142919 | 189850 | 3 |
| Large Payload | CBOR | Marshal | 212584 | 196753 | 1 |
| Large Payload | MessagePack | Marshal | 269291 | 526703 | 115 |
| Large Payload | JSON | Marshal | 487143 | 221456 | 8 |
| Large Payload | BEVE | Unmarshal | 273026 | 271204 | 418 |
| Large Payload | Sonic | Unmarshal | 438834 | 572980 | 608 |
| Large Payload | MessagePack | Unmarshal | 616119 | 318380 | 5738 |
| Large Payload | CBOR | Unmarshal | 816171 | 294890 | 6006 |
| Large Payload | JSON | Unmarshal | 2627341 | 532956 | 7052 |
| Medium Payload | BEVE ZeroCopy | Marshal | 9237 | 1 | 0 |
| Medium Payload | BEVE | Marshal | 13943 | 21764 | 1 |
| Medium Payload | Sonic | Marshal | 16844 | 22033 | 3 |
| Medium Payload | CBOR | Marshal | 23222 | 20492 | 1 |
| Medium Payload | MessagePack | Marshal | 34689 | 65771 | 22 |
| Medium Payload | JSON | Marshal | 49849 | 21993 | 8 |
| Medium Payload | BEVE | Unmarshal | 25911 | 27450 | 59 |
| Medium Payload | Sonic | Unmarshal | 53560 | 70045 | 82 |
| Medium Payload | MessagePack | Unmarshal | 58946 | 30283 | 549 |
| Medium Payload | CBOR | Unmarshal | 91296 | 35000 | 721 |
| Medium Payload | JSON | Unmarshal | 241130 | 48808 | 648 |
| Small Struct | BEVE ZeroCopy | Marshal | 486 | 0 | 0 |
| Small Struct | BEVE | Marshal | 796 | 896 | 1 |
| Small Struct | Sonic | Marshal | 1504 | 1580 | 2 |
| Small Struct | JSON | Marshal | 1894 | 768 | 1 |
| Small Struct | CBOR | Marshal | 2292 | 2304 | 1 |
| Small Struct | MessagePack | Marshal | 3043 | 4104 | 8 |
| Small Struct | BEVE | Unmarshal | 1051 | 952 | 4 |
| Small Struct | CBOR | Unmarshal | 3386 | 1184 | 28 |
| Small Struct | MessagePack | Unmarshal | 4071 | 2432 | 52 |
| Small Struct | Sonic | Unmarshal | 5927 | 7760 | 10 |
| Small Struct | JSON | Unmarshal | 23853 | 4680 | 83 |
