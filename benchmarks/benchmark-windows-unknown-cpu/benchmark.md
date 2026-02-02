# Unknown CPU — Windows

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE ZeroCopy | Marshal | 86680 | 65 | 0 |
| Large Payload | BEVE | Marshal | 116235 | 188477 | 1 |
| Large Payload | Sonic | Marshal | 165943 | 215292 | 3 |
| Large Payload | CBOR | Marshal | 258377 | 213105 | 1 |
| Large Payload | MessagePack | Marshal | 285363 | 526706 | 115 |
| Large Payload | JSON | Marshal | 479352 | 205100 | 8 |
| Large Payload | BEVE | Unmarshal | 302652 | 299751 | 418 |
| Large Payload | Sonic | Unmarshal | 527041 | 555831 | 590 |
| Large Payload | MessagePack | Unmarshal | 786178 | 354391 | 6461 |
| Large Payload | CBOR | Unmarshal | 969283 | 310474 | 6328 |
| Large Payload | JSON | Unmarshal | 2893602 | 529419 | 6908 |
| Medium Payload | BEVE ZeroCopy | Marshal | 9422 | 6 | 0 |
| Medium Payload | BEVE | Marshal | 14915 | 18435 | 1 |
| Medium Payload | CBOR | Marshal | 21865 | 18439 | 1 |
| Medium Payload | Sonic | Marshal | 22759 | 22085 | 3 |
| Medium Payload | MessagePack | Marshal | 50826 | 65776 | 22 |
| Medium Payload | JSON | Marshal | 56830 | 21990 | 8 |
| Medium Payload | BEVE | Unmarshal | 42224 | 34909 | 59 |
| Medium Payload | Sonic | Unmarshal | 44077 | 48789 | 69 |
| Medium Payload | MessagePack | Unmarshal | 62396 | 29708 | 540 |
| Medium Payload | CBOR | Unmarshal | 94787 | 33848 | 700 |
| Medium Payload | JSON | Unmarshal | 311666 | 61432 | 836 |
| Small Struct | BEVE ZeroCopy | Marshal | 457 | 0 | 0 |
| Small Struct | Sonic | Marshal | 727 | 556 | 2 |
| Small Struct | BEVE | Marshal | 1017 | 1024 | 1 |
| Small Struct | MessagePack | Marshal | 1571 | 1032 | 6 |
| Small Struct | JSON | Marshal | 2555 | 896 | 1 |
| Small Struct | CBOR | Marshal | 2622 | 2048 | 1 |
| Small Struct | Sonic | Unmarshal | 1918 | 1322 | 7 |
| Small Struct | BEVE | Unmarshal | 3217 | 3384 | 4 |
| Small Struct | MessagePack | Unmarshal | 6356 | 3200 | 68 |
| Small Struct | CBOR | Unmarshal | 9282 | 3496 | 74 |
| Small Struct | JSON | Unmarshal | 23058 | 4392 | 74 |
