# AMD EPYC 7763 64-Core Processor — Linux

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE ZeroCopy | Marshal | 80112 | 52 | 0 |
| Large Payload | BEVE | Marshal | 116493 | 196679 | 1 |
| Large Payload | Sonic | Marshal | 156654 | 215871 | 3 |
| Large Payload | CBOR | Marshal | 206584 | 188642 | 1 |
| Large Payload | MessagePack | Marshal | 317233 | 526777 | 115 |
| Large Payload | JSON | Marshal | 430010 | 205090 | 8 |
| Large Payload | BEVE | Unmarshal | 233307 | 252924 | 417 |
| Large Payload | Sonic | Unmarshal | 358195 | 540002 | 579 |
| Large Payload | MessagePack | Unmarshal | 578981 | 352678 | 6430 |
| Large Payload | CBOR | Unmarshal | 681723 | 300522 | 6120 |
| Large Payload | JSON | Unmarshal | 2226475 | 515835 | 6722 |
| Medium Payload | BEVE ZeroCopy | Marshal | 8578 | 3 | 0 |
| Medium Payload | BEVE | Marshal | 11229 | 18439 | 1 |
| Medium Payload | Sonic | Marshal | 14756 | 19481 | 3 |
| Medium Payload | CBOR | Marshal | 22643 | 20499 | 1 |
| Medium Payload | MessagePack | Marshal | 36171 | 65782 | 22 |
| Medium Payload | JSON | Marshal | 46110 | 24808 | 8 |
| Medium Payload | BEVE | Unmarshal | 25278 | 28670 | 59 |
| Medium Payload | Sonic | Unmarshal | 43014 | 64261 | 79 |
| Medium Payload | CBOR | Unmarshal | 58785 | 25192 | 523 |
| Medium Payload | MessagePack | Unmarshal | 59055 | 36648 | 682 |
| Medium Payload | JSON | Unmarshal | 236247 | 53544 | 747 |
| Small Struct | BEVE ZeroCopy | Marshal | 293 | 0 | 0 |
| Small Struct | Sonic | Marshal | 887 | 763 | 2 |
| Small Struct | BEVE | Marshal | 1304 | 1792 | 1 |
| Small Struct | MessagePack | Marshal | 2159 | 2056 | 7 |
| Small Struct | JSON | Marshal | 2382 | 896 | 1 |
| Small Struct | CBOR | Marshal | 2408 | 1792 | 1 |
| Small Struct | BEVE | Unmarshal | 1514 | 1336 | 4 |
| Small Struct | Sonic | Unmarshal | 5246 | 7425 | 10 |
| Small Struct | MessagePack | Unmarshal | 6610 | 4000 | 85 |
| Small Struct | CBOR | Unmarshal | 6786 | 3560 | 76 |
| Small Struct | JSON | Unmarshal | 35485 | 7944 | 114 |
