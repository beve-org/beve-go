# AMD EPYC 7763 64-Core Processor — Linux

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE ZeroCopy | Marshal | 78140 | 39 | 0 |
| Large Payload | BEVE | Marshal | 115734 | 188486 | 1 |
| Large Payload | Sonic | Marshal | 156651 | 215137 | 3 |
| Large Payload | CBOR | Marshal | 214509 | 196759 | 1 |
| Large Payload | MessagePack | Marshal | 329898 | 526781 | 115 |
| Large Payload | JSON | Marshal | 439239 | 213283 | 8 |
| Large Payload | BEVE | Unmarshal | 255968 | 284707 | 419 |
| Large Payload | Sonic | Unmarshal | 369529 | 525868 | 568 |
| Large Payload | MessagePack | Unmarshal | 572997 | 339702 | 6162 |
| Large Payload | CBOR | Unmarshal | 692313 | 298666 | 6096 |
| Large Payload | JSON | Unmarshal | 2213500 | 505418 | 6637 |
| Medium Payload | BEVE ZeroCopy | Marshal | 8929 | 3 | 0 |
| Medium Payload | BEVE | Marshal | 17762 | 27271 | 1 |
| Medium Payload | Sonic | Marshal | 18627 | 25278 | 3 |
| Medium Payload | CBOR | Marshal | 23651 | 21780 | 1 |
| Medium Payload | MessagePack | Marshal | 37584 | 65782 | 22 |
| Medium Payload | JSON | Marshal | 51195 | 24808 | 8 |
| Medium Payload | BEVE | Unmarshal | 28902 | 34784 | 59 |
| Medium Payload | Sonic | Unmarshal | 37071 | 43621 | 67 |
| Medium Payload | MessagePack | Unmarshal | 59742 | 37648 | 703 |
| Medium Payload | CBOR | Unmarshal | 76465 | 35768 | 732 |
| Medium Payload | JSON | Unmarshal | 238595 | 60592 | 753 |
| Small Struct | BEVE | Marshal | 888 | 1408 | 1 |
| Small Struct | BEVE ZeroCopy | Marshal | 908 | 0 | 0 |
| Small Struct | CBOR | Marshal | 1788 | 1792 | 1 |
| Small Struct | Sonic | Marshal | 1853 | 2768 | 2 |
| Small Struct | MessagePack | Marshal | 2811 | 4104 | 8 |
| Small Struct | JSON | Marshal | 3740 | 2048 | 1 |
| Small Struct | CBOR | Unmarshal | 1519 | 424 | 12 |
| Small Struct | Sonic | Unmarshal | 1603 | 1968 | 8 |
| Small Struct | BEVE | Unmarshal | 1878 | 3384 | 4 |
| Small Struct | MessagePack | Unmarshal | 2928 | 1888 | 41 |
| Small Struct | JSON | Unmarshal | 27361 | 7912 | 113 |
