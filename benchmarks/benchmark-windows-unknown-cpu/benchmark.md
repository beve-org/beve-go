# Unknown CPU — Windows

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE ZeroCopy | Marshal | 84225 | 79 | 0 |
| Large Payload | BEVE | Marshal | 123337 | 196631 | 1 |
| Large Payload | Sonic | Marshal | 175792 | 214915 | 3 |
| Large Payload | CBOR | Marshal | 237986 | 188560 | 1 |
| Large Payload | MessagePack | Marshal | 296834 | 526710 | 115 |
| Large Payload | JSON | Marshal | 468688 | 205127 | 8 |
| Large Payload | BEVE | Unmarshal | 338993 | 273481 | 418 |
| Large Payload | Sonic | Unmarshal | 445017 | 506933 | 557 |
| Large Payload | MessagePack | Unmarshal | 710847 | 332214 | 6000 |
| Large Payload | CBOR | Unmarshal | 885891 | 299579 | 6104 |
| Large Payload | JSON | Unmarshal | 2804637 | 542924 | 7089 |
| Medium Payload | BEVE ZeroCopy | Marshal | 8147 | 5 | 0 |
| Medium Payload | BEVE | Marshal | 14231 | 18435 | 1 |
| Medium Payload | Sonic | Marshal | 17195 | 18927 | 3 |
| Medium Payload | CBOR | Marshal | 30814 | 27285 | 1 |
| Medium Payload | MessagePack | Marshal | 38672 | 65771 | 22 |
| Medium Payload | JSON | Marshal | 46576 | 18665 | 8 |
| Medium Payload | BEVE | Unmarshal | 31595 | 28507 | 59 |
| Medium Payload | Sonic | Unmarshal | 54212 | 63852 | 77 |
| Medium Payload | MessagePack | Unmarshal | 75162 | 37166 | 693 |
| Medium Payload | CBOR | Unmarshal | 80837 | 27328 | 564 |
| Medium Payload | JSON | Unmarshal | 260561 | 46456 | 655 |
| Small Struct | BEVE ZeroCopy | Marshal | 723 | 0 | 0 |
| Small Struct | JSON | Marshal | 1082 | 384 | 1 |
| Small Struct | MessagePack | Marshal | 1208 | 1032 | 6 |
| Small Struct | Sonic | Marshal | 1256 | 1445 | 2 |
| Small Struct | CBOR | Marshal | 1356 | 1024 | 1 |
| Small Struct | BEVE | Marshal | 2404 | 3072 | 1 |
| Small Struct | BEVE | Unmarshal | 1377 | 1080 | 4 |
| Small Struct | MessagePack | Unmarshal | 3332 | 1688 | 38 |
| Small Struct | CBOR | Unmarshal | 4111 | 1384 | 32 |
| Small Struct | Sonic | Unmarshal | 4435 | 4443 | 9 |
| Small Struct | JSON | Unmarshal | 28717 | 7400 | 97 |
