# Unknown CPU — Windows

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE ZeroCopy | Marshal | 88985 | 65 | 0 |
| Large Payload | BEVE | Marshal | 162500 | 196735 | 1 |
| Large Payload | CBOR | Marshal | 238467 | 188552 | 1 |
| Large Payload | Sonic | Marshal | 247571 | 227018 | 3 |
| Large Payload | MessagePack | Marshal | 482254 | 526753 | 115 |
| Large Payload | JSON | Marshal | 609746 | 221498 | 8 |
| Large Payload | BEVE | Unmarshal | 423630 | 263629 | 417 |
| Large Payload | Sonic | Unmarshal | 492900 | 549241 | 567 |
| Large Payload | MessagePack | Unmarshal | 761711 | 364013 | 6669 |
| Large Payload | CBOR | Unmarshal | 849693 | 300569 | 6133 |
| Large Payload | JSON | Unmarshal | 2918592 | 560404 | 7266 |
| Medium Payload | BEVE ZeroCopy | Marshal | 7815 | 5 | 0 |
| Medium Payload | BEVE | Marshal | 19651 | 21772 | 1 |
| Medium Payload | Sonic | Marshal | 20570 | 20866 | 3 |
| Medium Payload | CBOR | Marshal | 38246 | 24593 | 1 |
| Medium Payload | MessagePack | Marshal | 51769 | 65777 | 22 |
| Medium Payload | JSON | Marshal | 56288 | 24810 | 8 |
| Medium Payload | BEVE | Unmarshal | 31648 | 21980 | 59 |
| Medium Payload | Sonic | Unmarshal | 60495 | 64523 | 79 |
| Medium Payload | MessagePack | Unmarshal | 76099 | 31518 | 574 |
| Medium Payload | CBOR | Unmarshal | 83312 | 32024 | 659 |
| Medium Payload | JSON | Unmarshal | 280265 | 61784 | 832 |
| Small Struct | BEVE ZeroCopy | Marshal | 851 | 0 | 0 |
| Small Struct | BEVE | Marshal | 894 | 1024 | 1 |
| Small Struct | CBOR | Marshal | 2677 | 2688 | 1 |
| Small Struct | Sonic | Marshal | 2694 | 2761 | 2 |
| Small Struct | MessagePack | Marshal | 2899 | 4104 | 8 |
| Small Struct | JSON | Marshal | 4951 | 2305 | 1 |
| Small Struct | BEVE | Unmarshal | 1877 | 1848 | 4 |
| Small Struct | CBOR | Unmarshal | 2435 | 816 | 20 |
| Small Struct | MessagePack | Unmarshal | 5248 | 3136 | 66 |
| Small Struct | Sonic | Unmarshal | 6632 | 7457 | 10 |
| Small Struct | JSON | Unmarshal | 11327 | 2408 | 47 |
