# Intel(R) Xeon(R) Platinum 8370C CPU @ 2.80GHz — Linux

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE ZeroCopy | Marshal | 69485 | 26 | 0 |
| Large Payload | BEVE | Marshal | 107326 | 188500 | 1 |
| Large Payload | Sonic | Marshal | 160973 | 215956 | 3 |
| Large Payload | CBOR | Marshal | 210990 | 213176 | 1 |
| Large Payload | MessagePack | Marshal | 290345 | 526774 | 115 |
| Large Payload | JSON | Marshal | 422388 | 213309 | 8 |
| Large Payload | BEVE | Unmarshal | 243923 | 282626 | 419 |
| Large Payload | Sonic | Unmarshal | 383470 | 562249 | 593 |
| Large Payload | MessagePack | Unmarshal | 542316 | 341577 | 6195 |
| Large Payload | CBOR | Unmarshal | 666072 | 296618 | 6058 |
| Large Payload | JSON | Unmarshal | 2079687 | 550249 | 7153 |
| Medium Payload | BEVE ZeroCopy | Marshal | 7570 | 3 | 0 |
| Medium Payload | BEVE | Marshal | 11134 | 18448 | 1 |
| Medium Payload | Sonic | Marshal | 18918 | 25259 | 3 |
| Medium Payload | CBOR | Marshal | 20642 | 20502 | 1 |
| Medium Payload | MessagePack | Marshal | 34639 | 65783 | 22 |
| Medium Payload | JSON | Marshal | 46559 | 24805 | 8 |
| Medium Payload | BEVE | Unmarshal | 24062 | 28318 | 59 |
| Medium Payload | Sonic | Unmarshal | 44463 | 63033 | 73 |
| Medium Payload | MessagePack | Unmarshal | 59439 | 40433 | 753 |
| Medium Payload | CBOR | Unmarshal | 65945 | 30712 | 631 |
| Medium Payload | JSON | Unmarshal | 188869 | 49656 | 669 |
| Small Struct | BEVE ZeroCopy | Marshal | 385 | 0 | 0 |
| Small Struct | BEVE | Marshal | 398 | 416 | 1 |
| Small Struct | CBOR | Marshal | 1315 | 1280 | 1 |
| Small Struct | MessagePack | Marshal | 1486 | 2056 | 7 |
| Small Struct | Sonic | Marshal | 2091 | 2775 | 2 |
| Small Struct | JSON | Marshal | 3873 | 2048 | 1 |
| Small Struct | BEVE | Unmarshal | 850 | 632 | 4 |
| Small Struct | Sonic | Unmarshal | 1568 | 1904 | 8 |
| Small Struct | MessagePack | Unmarshal | 4143 | 3168 | 67 |
| Small Struct | CBOR | Unmarshal | 7127 | 4232 | 89 |
| Small Struct | JSON | Unmarshal | 7570 | 2056 | 36 |
