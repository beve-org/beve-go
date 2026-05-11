# Unknown CPU — Windows

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE ZeroCopy | Marshal | 72575 | 65 | 0 |
| Large Payload | BEVE | Marshal | 111248 | 188450 | 1 |
| Large Payload | Sonic | Marshal | 161778 | 223546 | 3 |
| Large Payload | CBOR | Marshal | 200811 | 188537 | 1 |
| Large Payload | MessagePack | Marshal | 282197 | 526708 | 115 |
| Large Payload | JSON | Marshal | 511071 | 229651 | 8 |
| Large Payload | BEVE | Unmarshal | 267100 | 262403 | 419 |
| Large Payload | Sonic | Unmarshal | 462060 | 523899 | 570 |
| Large Payload | MessagePack | Unmarshal | 634571 | 316796 | 5702 |
| Large Payload | CBOR | Unmarshal | 824069 | 312746 | 6382 |
| Large Payload | JSON | Unmarshal | 2366806 | 509810 | 6501 |
| Medium Payload | BEVE ZeroCopy | Marshal | 7949 | 5 | 0 |
| Medium Payload | BEVE | Marshal | 11913 | 18435 | 1 |
| Medium Payload | Sonic | Marshal | 15503 | 19409 | 3 |
| Medium Payload | CBOR | Marshal | 21440 | 18439 | 1 |
| Medium Payload | MessagePack | Marshal | 35579 | 65772 | 22 |
| Medium Payload | JSON | Marshal | 48193 | 24809 | 8 |
| Medium Payload | BEVE | Unmarshal | 27740 | 26330 | 58 |
| Medium Payload | Sonic | Unmarshal | 54036 | 60258 | 75 |
| Medium Payload | MessagePack | Unmarshal | 64413 | 33324 | 613 |
| Medium Payload | CBOR | Unmarshal | 79736 | 30528 | 631 |
| Medium Payload | JSON | Unmarshal | 240569 | 51656 | 674 |
| Small Struct | BEVE | Marshal | 336 | 192 | 1 |
| Small Struct | BEVE ZeroCopy | Marshal | 766 | 0 | 0 |
| Small Struct | MessagePack | Marshal | 814 | 520 | 5 |
| Small Struct | CBOR | Marshal | 1303 | 1152 | 1 |
| Small Struct | Sonic | Marshal | 2338 | 2341 | 2 |
| Small Struct | JSON | Marshal | 3040 | 1280 | 1 |
| Small Struct | BEVE | Unmarshal | 1350 | 1464 | 4 |
| Small Struct | Sonic | Unmarshal | 3443 | 3626 | 9 |
| Small Struct | CBOR | Unmarshal | 5580 | 2784 | 60 |
| Small Struct | MessagePack | Unmarshal | 7651 | 5144 | 105 |
| Small Struct | JSON | Unmarshal | 20765 | 4488 | 77 |
