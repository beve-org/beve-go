# AMD EPYC 7763 64-Core Processor — Linux

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE ZeroCopy | Marshal | 79240 | 26 | 0 |
| Large Payload | BEVE | Marshal | 128983 | 204900 | 1 |
| Large Payload | Sonic | Marshal | 162957 | 215717 | 3 |
| Large Payload | CBOR | Marshal | 215152 | 188589 | 1 |
| Large Payload | MessagePack | Marshal | 348110 | 526783 | 115 |
| Large Payload | JSON | Marshal | 423792 | 196896 | 8 |
| Large Payload | BEVE | Unmarshal | 257424 | 269730 | 418 |
| Large Payload | Sonic | Unmarshal | 410646 | 580116 | 602 |
| Large Payload | MessagePack | Unmarshal | 583526 | 356104 | 6500 |
| Large Payload | CBOR | Unmarshal | 683212 | 290281 | 5908 |
| Large Payload | JSON | Unmarshal | 2485167 | 541411 | 7090 |
| Medium Payload | BEVE ZeroCopy | Marshal | 10372 | 3 | 0 |
| Medium Payload | BEVE | Marshal | 14697 | 21767 | 1 |
| Medium Payload | Sonic | Marshal | 16697 | 22213 | 3 |
| Medium Payload | CBOR | Marshal | 23618 | 21783 | 1 |
| Medium Payload | MessagePack | Marshal | 38205 | 65783 | 22 |
| Medium Payload | JSON | Marshal | 50318 | 24808 | 8 |
| Medium Payload | BEVE | Unmarshal | 22745 | 23069 | 59 |
| Medium Payload | Sonic | Unmarshal | 43300 | 63946 | 72 |
| Medium Payload | MessagePack | Unmarshal | 61638 | 39185 | 730 |
| Medium Payload | CBOR | Unmarshal | 72934 | 33496 | 688 |
| Medium Payload | JSON | Unmarshal | 207678 | 44504 | 588 |
| Small Struct | BEVE ZeroCopy | Marshal | 404 | 0 | 0 |
| Small Struct | CBOR | Marshal | 541 | 352 | 1 |
| Small Struct | BEVE | Marshal | 996 | 1024 | 1 |
| Small Struct | MessagePack | Marshal | 1204 | 1032 | 6 |
| Small Struct | Sonic | Marshal | 1695 | 2391 | 2 |
| Small Struct | JSON | Marshal | 3398 | 1536 | 1 |
| Small Struct | BEVE | Unmarshal | 916 | 952 | 4 |
| Small Struct | Sonic | Unmarshal | 1579 | 1957 | 8 |
| Small Struct | MessagePack | Unmarshal | 5758 | 3904 | 82 |
| Small Struct | CBOR | Unmarshal | 9784 | 4296 | 91 |
| Small Struct | JSON | Unmarshal | 16301 | 4040 | 63 |
