# AMD EPYC 7763 64-Core Processor — Linux

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE ZeroCopy | Marshal | 79666 | 52 | 0 |
| Large Payload | BEVE | Marshal | 109947 | 172101 | 1 |
| Large Payload | Sonic | Marshal | 149036 | 198937 | 3 |
| Large Payload | CBOR | Marshal | 208247 | 180368 | 1 |
| Large Payload | MessagePack | Marshal | 331103 | 526779 | 115 |
| Large Payload | JSON | Marshal | 454897 | 213309 | 8 |
| Large Payload | BEVE | Unmarshal | 245410 | 268351 | 417 |
| Large Payload | Sonic | Unmarshal | 373849 | 556272 | 584 |
| Large Payload | MessagePack | Unmarshal | 562337 | 335892 | 6082 |
| Large Payload | CBOR | Unmarshal | 752365 | 332186 | 6769 |
| Large Payload | JSON | Unmarshal | 2277901 | 509290 | 6710 |
| Medium Payload | BEVE ZeroCopy | Marshal | 8857 | 5 | 0 |
| Medium Payload | BEVE | Marshal | 12592 | 20486 | 1 |
| Medium Payload | Sonic | Marshal | 16410 | 20937 | 3 |
| Medium Payload | CBOR | Marshal | 21469 | 19090 | 1 |
| Medium Payload | MessagePack | Marshal | 38615 | 65783 | 22 |
| Medium Payload | JSON | Marshal | 42614 | 20711 | 8 |
| Medium Payload | BEVE | Unmarshal | 24357 | 23934 | 59 |
| Medium Payload | Sonic | Unmarshal | 36460 | 50359 | 68 |
| Medium Payload | MessagePack | Unmarshal | 51515 | 31823 | 581 |
| Medium Payload | CBOR | Unmarshal | 76736 | 32904 | 677 |
| Medium Payload | JSON | Unmarshal | 223881 | 53208 | 685 |
| Small Struct | BEVE ZeroCopy | Marshal | 288 | 0 | 0 |
| Small Struct | BEVE | Marshal | 1107 | 1792 | 1 |
| Small Struct | Sonic | Marshal | 1530 | 2121 | 2 |
| Small Struct | JSON | Marshal | 2822 | 1408 | 1 |
| Small Struct | CBOR | Marshal | 3126 | 2689 | 1 |
| Small Struct | MessagePack | Marshal | 4660 | 8201 | 9 |
| Small Struct | BEVE | Unmarshal | 2247 | 3512 | 4 |
| Small Struct | MessagePack | Unmarshal | 2380 | 1376 | 31 |
| Small Struct | Sonic | Unmarshal | 2488 | 3670 | 9 |
| Small Struct | CBOR | Unmarshal | 7528 | 4264 | 90 |
| Small Struct | JSON | Unmarshal | 16382 | 4136 | 66 |
