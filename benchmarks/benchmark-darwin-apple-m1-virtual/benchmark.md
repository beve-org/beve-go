# Apple M1 (Virtual) — Darwin

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE ZeroCopy | Marshal | 75111 | 26 | 0 |
| Large Payload | BEVE | Marshal | 138656 | 188469 | 1 |
| Large Payload | CBOR | Marshal | 153771 | 196743 | 1 |
| Large Payload | MessagePack | Marshal | 399824 | 526754 | 115 |
| Large Payload | JSON | Marshal | 567380 | 213305 | 8 |
| Large Payload | Sonic | Marshal | 637296 | 214510 | 3 |
| Large Payload | BEVE | Unmarshal | 282531 | 271120 | 418 |
| Large Payload | Sonic | Unmarshal | 362497 | 365762 | 211 |
| Large Payload | MessagePack | Unmarshal | 607443 | 380746 | 7001 |
| Large Payload | CBOR | Unmarshal | 678474 | 327865 | 6660 |
| Large Payload | JSON | Unmarshal | 2578309 | 556979 | 7418 |
| Medium Payload | BEVE ZeroCopy | Marshal | 8741 | 3 | 0 |
| Medium Payload | BEVE | Marshal | 16067 | 18436 | 1 |
| Medium Payload | CBOR | Marshal | 24075 | 19085 | 1 |
| Medium Payload | MessagePack | Marshal | 39688 | 65778 | 22 |
| Medium Payload | JSON | Marshal | 45494 | 20710 | 8 |
| Medium Payload | Sonic | Marshal | 62212 | 22063 | 3 |
| Medium Payload | BEVE | Unmarshal | 29530 | 29180 | 59 |
| Medium Payload | Sonic | Unmarshal | 43345 | 42080 | 33 |
| Medium Payload | CBOR | Unmarshal | 59321 | 23592 | 486 |
| Medium Payload | MessagePack | Unmarshal | 61077 | 31213 | 571 |
| Medium Payload | JSON | Unmarshal | 225949 | 50040 | 679 |
| Small Struct | BEVE ZeroCopy | Marshal | 597 | 0 | 0 |
| Small Struct | BEVE | Marshal | 663 | 640 | 1 |
| Small Struct | CBOR | Marshal | 1118 | 640 | 1 |
| Small Struct | MessagePack | Marshal | 2232 | 2056 | 7 |
| Small Struct | JSON | Marshal | 2378 | 1280 | 1 |
| Small Struct | Sonic | Marshal | 7528 | 3102 | 2 |
| Small Struct | BEVE | Unmarshal | 2324 | 3000 | 4 |
| Small Struct | Sonic | Unmarshal | 3174 | 2883 | 6 |
| Small Struct | CBOR | Unmarshal | 6676 | 3168 | 68 |
| Small Struct | MessagePack | Unmarshal | 7748 | 4608 | 96 |
| Small Struct | JSON | Unmarshal | 33640 | 7464 | 99 |
