# AMD EPYC 7763 64-Core Processor — Linux

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE ZeroCopy | Marshal | 79489 | 39 | 0 |
| Large Payload | BEVE | Marshal | 120876 | 180322 | 1 |
| Large Payload | Sonic | Marshal | 173914 | 232543 | 3 |
| Large Payload | CBOR | Marshal | 213882 | 188616 | 1 |
| Large Payload | MessagePack | Marshal | 343808 | 526782 | 115 |
| Large Payload | JSON | Marshal | 453261 | 213283 | 8 |
| Large Payload | BEVE | Unmarshal | 251735 | 273442 | 417 |
| Large Payload | Sonic | Unmarshal | 369122 | 546155 | 593 |
| Large Payload | MessagePack | Unmarshal | 608681 | 371114 | 6803 |
| Large Payload | CBOR | Unmarshal | 738073 | 325403 | 6638 |
| Large Payload | JSON | Unmarshal | 2372758 | 551738 | 7180 |
| Medium Payload | BEVE ZeroCopy | Marshal | 7089 | 3 | 0 |
| Medium Payload | BEVE | Marshal | 11081 | 16389 | 1 |
| Medium Payload | Sonic | Marshal | 18249 | 25126 | 3 |
| Medium Payload | CBOR | Marshal | 19920 | 16398 | 1 |
| Medium Payload | MessagePack | Marshal | 38211 | 65783 | 22 |
| Medium Payload | JSON | Marshal | 38418 | 18663 | 8 |
| Medium Payload | BEVE | Unmarshal | 23627 | 26430 | 59 |
| Medium Payload | Sonic | Unmarshal | 33950 | 48631 | 71 |
| Medium Payload | MessagePack | Unmarshal | 55003 | 35120 | 649 |
| Medium Payload | CBOR | Unmarshal | 76882 | 36872 | 753 |
| Medium Payload | JSON | Unmarshal | 185848 | 42456 | 576 |
| Small Struct | BEVE ZeroCopy | Marshal | 800 | 0 | 0 |
| Small Struct | Sonic | Marshal | 1224 | 940 | 2 |
| Small Struct | CBOR | Marshal | 1542 | 1024 | 1 |
| Small Struct | BEVE | Marshal | 1859 | 2688 | 1 |
| Small Struct | MessagePack | Marshal | 2964 | 4104 | 8 |
| Small Struct | JSON | Marshal | 3823 | 1536 | 1 |
| Small Struct | CBOR | Unmarshal | 1612 | 328 | 10 |
| Small Struct | BEVE | Unmarshal | 1983 | 2360 | 4 |
| Small Struct | Sonic | Unmarshal | 2147 | 1957 | 8 |
| Small Struct | MessagePack | Unmarshal | 7241 | 4697 | 99 |
| Small Struct | JSON | Unmarshal | 28633 | 7368 | 96 |
