# AMD EPYC 7763 64-Core Processor — Linux

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE ZeroCopy | Marshal | 78696 | 207 | 2 |
| Large Payload | BEVE | Marshal | 129261 | 197141 | 3 |
| Large Payload | Sonic | Marshal | 179380 | 233857 | 4 |
| Large Payload | CBOR | Marshal | 198980 | 189339 | 2 |
| Large Payload | MessagePack | Marshal | 321159 | 526837 | 115 |
| Large Payload | JSON | Marshal | 452586 | 221880 | 9 |
| Large Payload | BEVE | Unmarshal | 241791 | 276515 | 418 |
| Large Payload | Sonic | Unmarshal | 369174 | 561531 | 583 |
| Large Payload | MessagePack | Unmarshal | 548437 | 340951 | 6198 |
| Large Payload | CBOR | Unmarshal | 755358 | 303146 | 6188 |
| Large Payload | JSON | Unmarshal | 2357065 | 552860 | 7198 |
| Medium Payload | BEVE ZeroCopy | Marshal | 7418 | 134 | 2 |
| Medium Payload | BEVE | Marshal | 13429 | 20668 | 3 |
| Medium Payload | Sonic | Marshal | 16652 | 22677 | 4 |
| Medium Payload | CBOR | Marshal | 19333 | 18529 | 2 |
| Medium Payload | MessagePack | Marshal | 35977 | 65838 | 22 |
| Medium Payload | JSON | Marshal | 40692 | 19456 | 9 |
| Medium Payload | BEVE | Unmarshal | 24337 | 28895 | 59 |
| Medium Payload | Sonic | Unmarshal | 37092 | 54926 | 70 |
| Medium Payload | MessagePack | Unmarshal | 55190 | 35040 | 646 |
| Medium Payload | CBOR | Unmarshal | 84417 | 36456 | 748 |
| Medium Payload | JSON | Unmarshal | 213788 | 51544 | 690 |
| Small Struct | BEVE | Marshal | 615 | 672 | 3 |
| Small Struct | BEVE ZeroCopy | Marshal | 664 | 289 | 2 |
| Small Struct | CBOR | Marshal | 826 | 656 | 2 |
| Small Struct | Sonic | Marshal | 1443 | 2052 | 3 |
| Small Struct | MessagePack | Marshal | 2944 | 4225 | 8 |
| Small Struct | JSON | Marshal | 5400 | 3220 | 2 |
| Small Struct | MessagePack | Unmarshal | 1281 | 456 | 12 |
| Small Struct | Sonic | Unmarshal | 1615 | 1925 | 8 |
| Small Struct | BEVE | Unmarshal | 1803 | 3000 | 4 |
| Small Struct | CBOR | Unmarshal | 4824 | 2208 | 48 |
| Small Struct | JSON | Unmarshal | 27457 | 7912 | 113 |
