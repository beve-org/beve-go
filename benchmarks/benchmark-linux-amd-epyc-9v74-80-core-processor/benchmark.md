# AMD EPYC 9V74 80-Core Processor — Linux

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE ZeroCopy | Marshal | 72071 | 26 | 0 |
| Large Payload | BEVE | Marshal | 123963 | 196684 | 1 |
| Large Payload | Sonic | Marshal | 173074 | 224435 | 3 |
| Large Payload | CBOR | Marshal | 206334 | 188616 | 1 |
| Large Payload | MessagePack | Marshal | 338928 | 526780 | 115 |
| Large Payload | JSON | Marshal | 441946 | 213282 | 8 |
| Large Payload | BEVE | Unmarshal | 248935 | 277891 | 419 |
| Large Payload | Sonic | Unmarshal | 417109 | 567916 | 582 |
| Large Payload | MessagePack | Unmarshal | 561112 | 346680 | 6310 |
| Large Payload | CBOR | Unmarshal | 761241 | 305274 | 6228 |
| Large Payload | JSON | Unmarshal | 2085192 | 531163 | 6973 |
| Medium Payload | BEVE ZeroCopy | Marshal | 6545 | 3 | 0 |
| Medium Payload | BEVE | Marshal | 10287 | 16389 | 1 |
| Medium Payload | Sonic | Marshal | 16785 | 22238 | 3 |
| Medium Payload | CBOR | Marshal | 19762 | 18449 | 1 |
| Medium Payload | MessagePack | Marshal | 39694 | 65782 | 22 |
| Medium Payload | JSON | Marshal | 44990 | 24812 | 8 |
| Medium Payload | BEVE | Unmarshal | 24218 | 27230 | 59 |
| Medium Payload | Sonic | Unmarshal | 43422 | 62127 | 71 |
| Medium Payload | MessagePack | Unmarshal | 58154 | 39777 | 744 |
| Medium Payload | CBOR | Unmarshal | 75255 | 32376 | 665 |
| Medium Payload | JSON | Unmarshal | 234847 | 64840 | 847 |
| Small Struct | BEVE ZeroCopy | Marshal | 242 | 0 | 0 |
| Small Struct | BEVE | Marshal | 585 | 896 | 1 |
| Small Struct | Sonic | Marshal | 1090 | 1324 | 2 |
| Small Struct | CBOR | Marshal | 1806 | 1792 | 1 |
| Small Struct | JSON | Marshal | 2195 | 1152 | 1 |
| Small Struct | MessagePack | Marshal | 4979 | 8201 | 9 |
| Small Struct | BEVE | Unmarshal | 1313 | 1848 | 4 |
| Small Struct | Sonic | Unmarshal | 2833 | 4171 | 9 |
| Small Struct | MessagePack | Unmarshal | 5522 | 4384 | 93 |
| Small Struct | CBOR | Unmarshal | 6596 | 3208 | 69 |
| Small Struct | JSON | Unmarshal | 7617 | 1448 | 32 |
