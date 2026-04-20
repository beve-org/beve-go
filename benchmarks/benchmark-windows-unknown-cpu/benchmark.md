# Unknown CPU — Windows

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE ZeroCopy | Marshal | 57312 | 65 | 0 |
| Large Payload | BEVE | Marshal | 85375 | 188502 | 1 |
| Large Payload | Sonic | Marshal | 125336 | 223582 | 3 |
| Large Payload | CBOR | Marshal | 163865 | 204870 | 1 |
| Large Payload | MessagePack | Marshal | 211055 | 526701 | 115 |
| Large Payload | JSON | Marshal | 350627 | 205069 | 8 |
| Large Payload | BEVE | Unmarshal | 213560 | 273635 | 418 |
| Large Payload | Sonic | Unmarshal | 345708 | 539365 | 569 |
| Large Payload | MessagePack | Unmarshal | 555445 | 360994 | 6596 |
| Large Payload | CBOR | Unmarshal | 643175 | 323065 | 6577 |
| Large Payload | JSON | Unmarshal | 1961857 | 527045 | 6936 |
| Medium Payload | BEVE ZeroCopy | Marshal | 5922 | 5 | 0 |
| Medium Payload | BEVE | Marshal | 10845 | 21763 | 1 |
| Medium Payload | Sonic | Marshal | 12828 | 22142 | 3 |
| Medium Payload | CBOR | Marshal | 15292 | 18441 | 1 |
| Medium Payload | MessagePack | Marshal | 25619 | 65771 | 22 |
| Medium Payload | JSON | Marshal | 29700 | 18661 | 8 |
| Medium Payload | BEVE | Unmarshal | 21337 | 31834 | 59 |
| Medium Payload | Sonic | Unmarshal | 39918 | 59270 | 75 |
| Medium Payload | MessagePack | Unmarshal | 49154 | 32028 | 586 |
| Medium Payload | CBOR | Unmarshal | 58354 | 29944 | 614 |
| Medium Payload | JSON | Unmarshal | 236948 | 66840 | 878 |
| Small Struct | BEVE | Marshal | 300 | 256 | 1 |
| Small Struct | BEVE ZeroCopy | Marshal | 492 | 0 | 0 |
| Small Struct | Sonic | Marshal | 610 | 933 | 2 |
| Small Struct | CBOR | Marshal | 839 | 1024 | 1 |
| Small Struct | JSON | Marshal | 1926 | 1152 | 1 |
| Small Struct | MessagePack | Marshal | 2067 | 4104 | 8 |
| Small Struct | MessagePack | Unmarshal | 1253 | 544 | 14 |
| Small Struct | BEVE | Unmarshal | 1884 | 3000 | 4 |
| Small Struct | Sonic | Unmarshal | 2316 | 3509 | 9 |
| Small Struct | CBOR | Unmarshal | 7707 | 4640 | 98 |
| Small Struct | JSON | Unmarshal | 14010 | 4072 | 64 |
