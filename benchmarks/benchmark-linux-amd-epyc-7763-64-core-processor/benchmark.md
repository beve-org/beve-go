# AMD EPYC 7763 64-Core Processor — Linux

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE ZeroCopy | Marshal | 77657 | 26 | 0 |
| Large Payload | BEVE | Marshal | 111710 | 180294 | 1 |
| Large Payload | Sonic | Marshal | 157060 | 215492 | 3 |
| Large Payload | CBOR | Marshal | 199915 | 188615 | 1 |
| Large Payload | MessagePack | Marshal | 309818 | 526776 | 115 |
| Large Payload | JSON | Marshal | 467733 | 229669 | 8 |
| Large Payload | BEVE | Unmarshal | 239066 | 275937 | 418 |
| Large Payload | Sonic | Unmarshal | 376212 | 572763 | 589 |
| Large Payload | MessagePack | Unmarshal | 548124 | 330371 | 5983 |
| Large Payload | CBOR | Unmarshal | 715691 | 318537 | 6499 |
| Large Payload | JSON | Unmarshal | 2385115 | 562074 | 7449 |
| Medium Payload | BEVE ZeroCopy | Marshal | 8783 | 0 | 0 |
| Medium Payload | BEVE | Marshal | 11524 | 18440 | 1 |
| Medium Payload | Sonic | Marshal | 16122 | 22247 | 3 |
| Medium Payload | CBOR | Marshal | 22283 | 20502 | 1 |
| Medium Payload | MessagePack | Marshal | 35914 | 65783 | 22 |
| Medium Payload | JSON | Marshal | 37162 | 18660 | 8 |
| Medium Payload | BEVE | Unmarshal | 22150 | 24509 | 59 |
| Medium Payload | Sonic | Unmarshal | 33085 | 47564 | 65 |
| Medium Payload | MessagePack | Unmarshal | 62680 | 40145 | 748 |
| Medium Payload | CBOR | Unmarshal | 68011 | 31112 | 642 |
| Medium Payload | JSON | Unmarshal | 205289 | 50520 | 651 |
| Small Struct | BEVE ZeroCopy | Marshal | 422 | 0 | 0 |
| Small Struct | BEVE | Marshal | 805 | 1280 | 1 |
| Small Struct | Sonic | Marshal | 1087 | 1466 | 2 |
| Small Struct | CBOR | Marshal | 2142 | 2304 | 1 |
| Small Struct | JSON | Marshal | 2543 | 1280 | 1 |
| Small Struct | MessagePack | Marshal | 2759 | 4104 | 8 |
| Small Struct | BEVE | Unmarshal | 698 | 376 | 4 |
| Small Struct | MessagePack | Unmarshal | 2604 | 1568 | 35 |
| Small Struct | Sonic | Unmarshal | 2991 | 4438 | 9 |
| Small Struct | CBOR | Unmarshal | 4183 | 2128 | 47 |
| Small Struct | JSON | Unmarshal | 21097 | 4840 | 88 |
