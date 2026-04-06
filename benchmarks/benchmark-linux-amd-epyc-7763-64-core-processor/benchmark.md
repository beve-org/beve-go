# AMD EPYC 7763 64-Core Processor — Linux

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE ZeroCopy | Marshal | 77915 | 26 | 0 |
| Large Payload | BEVE | Marshal | 120504 | 188568 | 1 |
| Large Payload | Sonic | Marshal | 168308 | 224313 | 3 |
| Large Payload | CBOR | Marshal | 227151 | 205007 | 1 |
| Large Payload | MessagePack | Marshal | 350038 | 526783 | 115 |
| Large Payload | JSON | Marshal | 444477 | 205116 | 8 |
| Large Payload | BEVE | Unmarshal | 265919 | 252830 | 418 |
| Large Payload | Sonic | Unmarshal | 378394 | 550592 | 583 |
| Large Payload | MessagePack | Unmarshal | 615781 | 377790 | 6937 |
| Large Payload | CBOR | Unmarshal | 712659 | 313419 | 6394 |
| Large Payload | JSON | Unmarshal | 2315731 | 539812 | 7053 |
| Medium Payload | BEVE ZeroCopy | Marshal | 8795 | 0 | 0 |
| Medium Payload | BEVE | Marshal | 9109 | 13572 | 1 |
| Medium Payload | Sonic | Marshal | 15407 | 20897 | 3 |
| Medium Payload | CBOR | Marshal | 22425 | 20499 | 1 |
| Medium Payload | MessagePack | Marshal | 26303 | 33008 | 21 |
| Medium Payload | JSON | Marshal | 37295 | 18663 | 8 |
| Medium Payload | BEVE | Unmarshal | 21935 | 23325 | 58 |
| Medium Payload | Sonic | Unmarshal | 41571 | 64607 | 77 |
| Medium Payload | MessagePack | Unmarshal | 52161 | 29294 | 533 |
| Medium Payload | CBOR | Unmarshal | 74611 | 34600 | 708 |
| Medium Payload | JSON | Unmarshal | 214236 | 52592 | 676 |
| Small Struct | BEVE ZeroCopy | Marshal | 571 | 0 | 0 |
| Small Struct | CBOR | Marshal | 814 | 640 | 1 |
| Small Struct | BEVE | Marshal | 823 | 896 | 1 |
| Small Struct | JSON | Marshal | 901 | 352 | 1 |
| Small Struct | Sonic | Marshal | 1644 | 2106 | 2 |
| Small Struct | MessagePack | Marshal | 1727 | 2056 | 7 |
| Small Struct | BEVE | Unmarshal | 1118 | 1464 | 4 |
| Small Struct | MessagePack | Unmarshal | 4194 | 3104 | 65 |
| Small Struct | Sonic | Unmarshal | 4255 | 7414 | 10 |
| Small Struct | CBOR | Unmarshal | 6889 | 3912 | 83 |
| Small Struct | JSON | Unmarshal | 17292 | 4328 | 72 |
