# Unknown CPU — Windows

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE ZeroCopy | Marshal | 81566 | 65 | 0 |
| Large Payload | BEVE | Marshal | 121897 | 196697 | 1 |
| Large Payload | Sonic | Marshal | 172597 | 223459 | 3 |
| Large Payload | CBOR | Marshal | 226941 | 204904 | 1 |
| Large Payload | MessagePack | Marshal | 280497 | 526705 | 115 |
| Large Payload | JSON | Marshal | 507397 | 213266 | 8 |
| Large Payload | BEVE | Unmarshal | 280473 | 271333 | 417 |
| Large Payload | Sonic | Unmarshal | 417214 | 523912 | 578 |
| Large Payload | MessagePack | Unmarshal | 657419 | 336833 | 6111 |
| Large Payload | CBOR | Unmarshal | 807576 | 290026 | 5909 |
| Large Payload | JSON | Unmarshal | 2554641 | 508923 | 6719 |
| Medium Payload | BEVE ZeroCopy | Marshal | 9839 | 3 | 0 |
| Medium Payload | Sonic | Marshal | 15294 | 18829 | 3 |
| Medium Payload | BEVE | Marshal | 16636 | 24577 | 1 |
| Medium Payload | CBOR | Marshal | 26541 | 21769 | 1 |
| Medium Payload | MessagePack | Marshal | 35070 | 65772 | 22 |
| Medium Payload | JSON | Marshal | 48161 | 21993 | 8 |
| Medium Payload | BEVE | Unmarshal | 23267 | 20346 | 59 |
| Medium Payload | Sonic | Unmarshal | 50146 | 60977 | 77 |
| Medium Payload | MessagePack | Unmarshal | 71042 | 36140 | 668 |
| Medium Payload | CBOR | Unmarshal | 92251 | 33664 | 689 |
| Medium Payload | JSON | Unmarshal | 213152 | 41176 | 536 |
| Small Struct | CBOR | Marshal | 440 | 224 | 1 |
| Small Struct | BEVE | Marshal | 469 | 384 | 1 |
| Small Struct | BEVE ZeroCopy | Marshal | 804 | 0 | 0 |
| Small Struct | Sonic | Marshal | 1844 | 2099 | 2 |
| Small Struct | MessagePack | Marshal | 2874 | 4104 | 8 |
| Small Struct | JSON | Marshal | 4873 | 2304 | 1 |
| Small Struct | MessagePack | Unmarshal | 1217 | 304 | 9 |
| Small Struct | BEVE | Unmarshal | 1825 | 2360 | 4 |
| Small Struct | Sonic | Unmarshal | 1844 | 2090 | 8 |
| Small Struct | JSON | Unmarshal | 3288 | 496 | 13 |
| Small Struct | CBOR | Unmarshal | 9095 | 4360 | 93 |
