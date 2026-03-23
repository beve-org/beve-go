# AMD EPYC 7763 64-Core Processor — Linux

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE ZeroCopy | Marshal | 79647 | 39 | 0 |
| Large Payload | BEVE | Marshal | 116573 | 180294 | 1 |
| Large Payload | Sonic | Marshal | 174197 | 232430 | 3 |
| Large Payload | CBOR | Marshal | 202240 | 180446 | 1 |
| Large Payload | MessagePack | Marshal | 344713 | 526780 | 115 |
| Large Payload | JSON | Marshal | 448177 | 213283 | 8 |
| Large Payload | BEVE | Unmarshal | 250495 | 267136 | 415 |
| Large Payload | Sonic | Unmarshal | 396428 | 587356 | 597 |
| Large Payload | MessagePack | Unmarshal | 563346 | 341285 | 6200 |
| Large Payload | CBOR | Unmarshal | 737205 | 331098 | 6739 |
| Large Payload | JSON | Unmarshal | 2183423 | 493163 | 6557 |
| Medium Payload | BEVE ZeroCopy | Marshal | 8436 | 1 | 0 |
| Medium Payload | BEVE | Marshal | 12704 | 20487 | 1 |
| Medium Payload | Sonic | Marshal | 18787 | 25137 | 3 |
| Medium Payload | CBOR | Marshal | 24148 | 21786 | 1 |
| Medium Payload | MessagePack | Marshal | 37681 | 65782 | 22 |
| Medium Payload | JSON | Marshal | 55652 | 27501 | 8 |
| Medium Payload | BEVE | Unmarshal | 22563 | 23421 | 59 |
| Medium Payload | Sonic | Unmarshal | 43124 | 65290 | 74 |
| Medium Payload | MessagePack | Unmarshal | 52693 | 32127 | 583 |
| Medium Payload | CBOR | Unmarshal | 71785 | 33016 | 680 |
| Medium Payload | JSON | Unmarshal | 237243 | 57520 | 749 |
| Small Struct | BEVE | Marshal | 357 | 288 | 1 |
| Small Struct | BEVE ZeroCopy | Marshal | 384 | 0 | 0 |
| Small Struct | CBOR | Marshal | 2150 | 2048 | 1 |
| Small Struct | Sonic | Marshal | 2186 | 3145 | 2 |
| Small Struct | MessagePack | Marshal | 2829 | 4104 | 8 |
| Small Struct | JSON | Marshal | 2873 | 1408 | 1 |
| Small Struct | BEVE | Unmarshal | 1094 | 1208 | 4 |
| Small Struct | Sonic | Unmarshal | 2882 | 4427 | 9 |
| Small Struct | CBOR | Unmarshal | 4318 | 2216 | 48 |
| Small Struct | MessagePack | Unmarshal | 4539 | 3448 | 72 |
| Small Struct | JSON | Unmarshal | 14712 | 3880 | 58 |
