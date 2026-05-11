# AMD EPYC 7763 64-Core Processor — Linux

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE ZeroCopy | Marshal | 81045 | 52 | 0 |
| Large Payload | BEVE | Marshal | 113140 | 180281 | 1 |
| Large Payload | Sonic | Marshal | 169414 | 224497 | 3 |
| Large Payload | CBOR | Marshal | 206522 | 188563 | 1 |
| Large Payload | MessagePack | Marshal | 331197 | 526780 | 115 |
| Large Payload | JSON | Marshal | 483678 | 229669 | 8 |
| Large Payload | BEVE | Unmarshal | 241517 | 269856 | 419 |
| Large Payload | Sonic | Unmarshal | 371756 | 556564 | 583 |
| Large Payload | MessagePack | Unmarshal | 572625 | 349223 | 6350 |
| Large Payload | CBOR | Unmarshal | 662949 | 285802 | 5833 |
| Large Payload | JSON | Unmarshal | 2561802 | 559234 | 7295 |
| Medium Payload | BEVE ZeroCopy | Marshal | 9315 | 3 | 0 |
| Medium Payload | BEVE | Marshal | 11053 | 16389 | 1 |
| Medium Payload | Sonic | Marshal | 15071 | 19450 | 3 |
| Medium Payload | CBOR | Marshal | 22832 | 20499 | 1 |
| Medium Payload | JSON | Marshal | 36425 | 18663 | 8 |
| Medium Payload | MessagePack | Marshal | 37486 | 65782 | 22 |
| Medium Payload | BEVE | Unmarshal | 27739 | 31199 | 59 |
| Medium Payload | Sonic | Unmarshal | 35634 | 49676 | 70 |
| Medium Payload | MessagePack | Unmarshal | 45578 | 26670 | 477 |
| Medium Payload | CBOR | Unmarshal | 80659 | 38936 | 794 |
| Medium Payload | JSON | Unmarshal | 280233 | 65080 | 843 |
| Small Struct | BEVE ZeroCopy | Marshal | 821 | 0 | 0 |
| Small Struct | BEVE | Marshal | 1106 | 1792 | 1 |
| Small Struct | Sonic | Marshal | 1218 | 1581 | 2 |
| Small Struct | CBOR | Marshal | 2818 | 3072 | 1 |
| Small Struct | MessagePack | Marshal | 2917 | 4104 | 8 |
| Small Struct | JSON | Marshal | 4699 | 2688 | 1 |
| Small Struct | BEVE | Unmarshal | 806 | 600 | 4 |
| Small Struct | Sonic | Unmarshal | 1871 | 2379 | 8 |
| Small Struct | MessagePack | Unmarshal | 4908 | 3648 | 78 |
| Small Struct | CBOR | Unmarshal | 6110 | 3272 | 71 |
| Small Struct | JSON | Unmarshal | 13348 | 3688 | 52 |
