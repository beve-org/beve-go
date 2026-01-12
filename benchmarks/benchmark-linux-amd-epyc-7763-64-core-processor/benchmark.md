# AMD EPYC 7763 64-Core Processor — Linux

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE ZeroCopy | Marshal | 79954 | 52 | 0 |
| Large Payload | BEVE | Marshal | 112709 | 188486 | 1 |
| Large Payload | Sonic | Marshal | 166082 | 231866 | 3 |
| Large Payload | CBOR | Marshal | 203873 | 188563 | 1 |
| Large Payload | MessagePack | Marshal | 330090 | 526779 | 115 |
| Large Payload | JSON | Marshal | 438686 | 205090 | 8 |
| Large Payload | BEVE | Unmarshal | 261702 | 285731 | 417 |
| Large Payload | Sonic | Unmarshal | 358742 | 538013 | 568 |
| Large Payload | MessagePack | Unmarshal | 553604 | 345495 | 6288 |
| Large Payload | CBOR | Unmarshal | 726126 | 324842 | 6625 |
| Large Payload | JSON | Unmarshal | 2211237 | 503362 | 6555 |
| Medium Payload | BEVE ZeroCopy | Marshal | 8067 | 3 | 0 |
| Medium Payload | BEVE | Marshal | 11758 | 18436 | 1 |
| Medium Payload | Sonic | Marshal | 18385 | 25117 | 3 |
| Medium Payload | CBOR | Marshal | 18836 | 16398 | 1 |
| Medium Payload | MessagePack | Marshal | 36940 | 65782 | 22 |
| Medium Payload | JSON | Marshal | 49571 | 24805 | 8 |
| Medium Payload | BEVE | Unmarshal | 26103 | 28959 | 59 |
| Medium Payload | Sonic | Unmarshal | 36187 | 51655 | 73 |
| Medium Payload | MessagePack | Unmarshal | 65053 | 43874 | 828 |
| Medium Payload | CBOR | Unmarshal | 69837 | 31992 | 655 |
| Medium Payload | JSON | Unmarshal | 229855 | 55256 | 736 |
| Small Struct | BEVE ZeroCopy | Marshal | 476 | 0 | 0 |
| Small Struct | BEVE | Marshal | 796 | 768 | 1 |
| Small Struct | JSON | Marshal | 1851 | 896 | 1 |
| Small Struct | Sonic | Marshal | 1958 | 2747 | 2 |
| Small Struct | CBOR | Marshal | 2648 | 2688 | 1 |
| Small Struct | MessagePack | Marshal | 4417 | 8201 | 9 |
| Small Struct | BEVE | Unmarshal | 1931 | 3384 | 4 |
| Small Struct | Sonic | Unmarshal | 2697 | 4192 | 9 |
| Small Struct | CBOR | Unmarshal | 2899 | 1288 | 30 |
| Small Struct | MessagePack | Unmarshal | 6924 | 5185 | 106 |
| Small Struct | JSON | Unmarshal | 14468 | 3976 | 61 |
