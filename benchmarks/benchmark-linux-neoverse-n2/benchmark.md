# Neoverse-N2 — Linux

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE ZeroCopy | Marshal | 69064 | 52 | 0 |
| Large Payload | BEVE | Marshal | 106974 | 196729 | 1 |
| Large Payload | CBOR | Marshal | 193493 | 196882 | 1 |
| Large Payload | MessagePack | Marshal | 269260 | 526800 | 115 |
| Large Payload | Sonic | Marshal | 298828 | 214866 | 3 |
| Large Payload | JSON | Marshal | 376551 | 205124 | 8 |
| Large Payload | BEVE | Unmarshal | 229759 | 273326 | 419 |
| Large Payload | Sonic | Unmarshal | 299909 | 412861 | 213 |
| Large Payload | MessagePack | Unmarshal | 513925 | 348637 | 6344 |
| Large Payload | CBOR | Unmarshal | 658859 | 319739 | 6521 |
| Large Payload | JSON | Unmarshal | 1815970 | 480226 | 6282 |
| Medium Payload | BEVE ZeroCopy | Marshal | 7813 | 6 | 0 |
| Medium Payload | BEVE | Marshal | 10424 | 20486 | 1 |
| Medium Payload | CBOR | Marshal | 18271 | 19087 | 1 |
| Medium Payload | MessagePack | Marshal | 23253 | 33007 | 21 |
| Medium Payload | Sonic | Marshal | 28509 | 20927 | 3 |
| Medium Payload | JSON | Marshal | 41680 | 24807 | 8 |
| Medium Payload | BEVE | Unmarshal | 21274 | 25661 | 59 |
| Medium Payload | Sonic | Unmarshal | 28549 | 38022 | 33 |
| Medium Payload | MessagePack | Unmarshal | 55444 | 39840 | 742 |
| Medium Payload | CBOR | Unmarshal | 65498 | 32424 | 664 |
| Medium Payload | JSON | Unmarshal | 196514 | 55752 | 719 |
| Small Struct | BEVE | Marshal | 306 | 256 | 1 |
| Small Struct | Sonic | Marshal | 617 | 286 | 2 |
| Small Struct | BEVE ZeroCopy | Marshal | 670 | 0 | 0 |
| Small Struct | CBOR | Marshal | 947 | 896 | 1 |
| Small Struct | MessagePack | Marshal | 2424 | 4104 | 8 |
| Small Struct | JSON | Marshal | 3868 | 2305 | 1 |
| Small Struct | BEVE | Unmarshal | 1485 | 2616 | 4 |
| Small Struct | MessagePack | Unmarshal | 1586 | 832 | 20 |
| Small Struct | Sonic | Unmarshal | 1600 | 1963 | 6 |
| Small Struct | CBOR | Unmarshal | 6398 | 3872 | 82 |
| Small Struct | JSON | Unmarshal | 21922 | 7496 | 100 |
