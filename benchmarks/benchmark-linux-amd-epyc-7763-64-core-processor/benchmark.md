# AMD EPYC 7763 64-Core Processor — Linux

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE ZeroCopy | Marshal | 80110 | 233 | 2 |
| Large Payload | BEVE | Marshal | 116195 | 188923 | 3 |
| Large Payload | Sonic | Marshal | 164958 | 216868 | 4 |
| Large Payload | CBOR | Marshal | 223081 | 213452 | 2 |
| Large Payload | MessagePack | Marshal | 302226 | 526835 | 115 |
| Large Payload | JSON | Marshal | 455155 | 221933 | 9 |
| Large Payload | BEVE | Unmarshal | 224847 | 252859 | 417 |
| Large Payload | Sonic | Unmarshal | 380206 | 579292 | 601 |
| Large Payload | MessagePack | Unmarshal | 595636 | 375198 | 6877 |
| Large Payload | CBOR | Unmarshal | 802452 | 324857 | 6627 |
| Large Payload | JSON | Unmarshal | 2296186 | 538113 | 7015 |
| Medium Payload | BEVE ZeroCopy | Marshal | 8471 | 134 | 2 |
| Medium Payload | BEVE | Marshal | 14280 | 21973 | 3 |
| Medium Payload | CBOR | Marshal | 19981 | 18562 | 2 |
| Medium Payload | Sonic | Marshal | 20169 | 21190 | 4 |
| Medium Payload | MessagePack | Marshal | 28422 | 33063 | 21 |
| Medium Payload | JSON | Marshal | 42260 | 18809 | 9 |
| Medium Payload | BEVE | Unmarshal | 25266 | 30815 | 59 |
| Medium Payload | Sonic | Unmarshal | 42515 | 66182 | 77 |
| Medium Payload | MessagePack | Unmarshal | 55598 | 34384 | 629 |
| Medium Payload | CBOR | Unmarshal | 75803 | 31480 | 646 |
| Medium Payload | JSON | Unmarshal | 208869 | 51704 | 660 |
| Small Struct | BEVE ZeroCopy | Marshal | 566 | 289 | 2 |
| Small Struct | BEVE | Marshal | 1199 | 1825 | 3 |
| Small Struct | Sonic | Marshal | 1330 | 2023 | 3 |
| Small Struct | MessagePack | Marshal | 2464 | 4225 | 8 |
| Small Struct | JSON | Marshal | 2742 | 1425 | 2 |
| Small Struct | CBOR | Marshal | 2875 | 2837 | 2 |
| Small Struct | Sonic | Unmarshal | 1676 | 2075 | 8 |
| Small Struct | BEVE | Unmarshal | 1779 | 3384 | 4 |
| Small Struct | CBOR | Unmarshal | 2185 | 712 | 18 |
| Small Struct | MessagePack | Unmarshal | 3111 | 2080 | 45 |
| Small Struct | JSON | Unmarshal | 5800 | 1256 | 26 |
