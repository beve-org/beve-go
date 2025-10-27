# Neoverse-N2 — Linux

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE ZeroCopy | Marshal | 66879 | 65 | 0 |
| Large Payload | BEVE | Marshal | 108491 | 196756 | 1 |
| Large Payload | CBOR | Marshal | 182760 | 196805 | 1 |
| Large Payload | MessagePack | Marshal | 285817 | 526805 | 115 |
| Large Payload | Sonic | Marshal | 311610 | 223110 | 3 |
| Large Payload | JSON | Marshal | 362528 | 196902 | 8 |
| Large Payload | BEVE | Unmarshal | 230870 | 294480 | 418 |
| Large Payload | Sonic | Unmarshal | 282657 | 394220 | 205 |
| Large Payload | MessagePack | Unmarshal | 505379 | 346063 | 6287 |
| Large Payload | CBOR | Unmarshal | 651296 | 318091 | 6490 |
| Large Payload | JSON | Unmarshal | 1908739 | 515356 | 6686 |
| Medium Payload | BEVE ZeroCopy | Marshal | 7560 | 3 | 0 |
| Medium Payload | BEVE | Marshal | 9342 | 18439 | 1 |
| Medium Payload | CBOR | Marshal | 18941 | 20496 | 1 |
| Medium Payload | MessagePack | Marshal | 22789 | 33007 | 21 |
| Medium Payload | Sonic | Marshal | 28765 | 20852 | 3 |
| Medium Payload | JSON | Marshal | 44140 | 24808 | 8 |
| Medium Payload | BEVE | Unmarshal | 22361 | 28606 | 59 |
| Medium Payload | Sonic | Unmarshal | 27996 | 37239 | 33 |
| Medium Payload | CBOR | Unmarshal | 56317 | 26280 | 543 |
| Medium Payload | MessagePack | Unmarshal | 57285 | 42529 | 797 |
| Medium Payload | JSON | Unmarshal | 175614 | 46984 | 633 |
| Small Struct | BEVE ZeroCopy | Marshal | 563 | 0 | 0 |
| Small Struct | BEVE | Marshal | 954 | 1792 | 1 |
| Small Struct | JSON | Marshal | 1377 | 704 | 1 |
| Small Struct | CBOR | Marshal | 1772 | 2048 | 1 |
| Small Struct | Sonic | Marshal | 3046 | 2376 | 2 |
| Small Struct | MessagePack | Marshal | 3601 | 8201 | 9 |
| Small Struct | Sonic | Unmarshal | 1129 | 1002 | 6 |
| Small Struct | BEVE | Unmarshal | 1197 | 1720 | 4 |
| Small Struct | MessagePack | Unmarshal | 4995 | 4384 | 93 |
| Small Struct | CBOR | Unmarshal | 5331 | 3144 | 67 |
| Small Struct | JSON | Unmarshal | 16030 | 4416 | 75 |
