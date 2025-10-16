# AMD EPYC 7763 64-Core Processor — Linux

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE ZeroCopy | Marshal | 85369 | 259 | 2 |
| Large Payload | BEVE | Marshal | 114981 | 180518 | 3 |
| Large Payload | Sonic | Marshal | 163775 | 217522 | 4 |
| Large Payload | CBOR | Marshal | 202663 | 189075 | 2 |
| Large Payload | MessagePack | Marshal | 299172 | 526835 | 115 |
| Large Payload | JSON | Marshal | 444549 | 222090 | 9 |
| Large Payload | BEVE | Unmarshal | 223492 | 258749 | 418 |
| Large Payload | Sonic | Unmarshal | 364359 | 553039 | 581 |
| Large Payload | MessagePack | Unmarshal | 569178 | 351019 | 6390 |
| Large Payload | CBOR | Unmarshal | 707384 | 305593 | 6224 |
| Large Payload | JSON | Unmarshal | 2279158 | 534691 | 6995 |
| Medium Payload | BEVE ZeroCopy | Marshal | 7265 | 134 | 2 |
| Medium Payload | BEVE | Marshal | 16621 | 19214 | 3 |
| Medium Payload | CBOR | Marshal | 20856 | 19203 | 2 |
| Medium Payload | Sonic | Marshal | 23856 | 25387 | 4 |
| Medium Payload | MessagePack | Marshal | 44357 | 65838 | 22 |
| Medium Payload | JSON | Marshal | 62107 | 24888 | 9 |
| Medium Payload | BEVE | Unmarshal | 31483 | 31711 | 59 |
| Medium Payload | Sonic | Unmarshal | 36226 | 51906 | 69 |
| Medium Payload | MessagePack | Unmarshal | 60206 | 39280 | 738 |
| Medium Payload | CBOR | Unmarshal | 66605 | 28056 | 581 |
| Medium Payload | JSON | Unmarshal | 205250 | 49592 | 633 |
| Small Struct | BEVE | Marshal | 777 | 576 | 3 |
| Small Struct | Sonic | Marshal | 830 | 956 | 3 |
| Small Struct | MessagePack | Marshal | 1555 | 1152 | 6 |
| Small Struct | BEVE ZeroCopy | Marshal | 1731 | 291 | 2 |
| Small Struct | CBOR | Marshal | 1970 | 1937 | 2 |
| Small Struct | JSON | Marshal | 7002 | 2833 | 2 |
| Small Struct | BEVE | Unmarshal | 1816 | 2360 | 4 |
| Small Struct | CBOR | Unmarshal | 2605 | 1000 | 24 |
| Small Struct | Sonic | Unmarshal | 3444 | 3910 | 9 |
| Small Struct | MessagePack | Unmarshal | 8346 | 5217 | 107 |
| Small Struct | JSON | Unmarshal | 31654 | 7592 | 103 |
