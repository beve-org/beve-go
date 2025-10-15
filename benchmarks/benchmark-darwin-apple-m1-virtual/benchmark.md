# Apple M1 (Virtual) — Darwin

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE ZeroCopy | Marshal | 66836 | 233 | 2 |
| Large Payload | BEVE | Marshal | 108335 | 181330 | 3 |
| Large Payload | CBOR | Marshal | 167339 | 189479 | 2 |
| Large Payload | MessagePack | Marshal | 278054 | 526812 | 115 |
| Large Payload | JSON | Marshal | 336762 | 213580 | 9 |
| Large Payload | Sonic | Marshal | 463848 | 215175 | 4 |
| Large Payload | BEVE | Unmarshal | 255745 | 279218 | 419 |
| Large Payload | Sonic | Unmarshal | 382150 | 325447 | 211 |
| Large Payload | MessagePack | Unmarshal | 575873 | 359158 | 6572 |
| Large Payload | CBOR | Unmarshal | 760657 | 334523 | 6808 |
| Large Payload | JSON | Unmarshal | 1842787 | 497924 | 6503 |
| Medium Payload | BEVE ZeroCopy | Marshal | 6382 | 128 | 2 |
| Medium Payload | BEVE | Marshal | 9469 | 20611 | 3 |
| Medium Payload | CBOR | Marshal | 22682 | 20581 | 2 |
| Medium Payload | MessagePack | Marshal | 30595 | 65834 | 22 |
| Medium Payload | JSON | Marshal | 40940 | 27595 | 9 |
| Medium Payload | Sonic | Marshal | 51440 | 25093 | 4 |
| Medium Payload | BEVE | Unmarshal | 21529 | 30685 | 59 |
| Medium Payload | Sonic | Unmarshal | 37075 | 36874 | 33 |
| Medium Payload | MessagePack | Unmarshal | 60398 | 37166 | 693 |
| Medium Payload | CBOR | Unmarshal | 71307 | 32088 | 662 |
| Medium Payload | JSON | Unmarshal | 192709 | 55336 | 726 |
| Small Struct | BEVE ZeroCopy | Marshal | 736 | 288 | 2 |
| Small Struct | BEVE | Marshal | 791 | 1824 | 3 |
| Small Struct | CBOR | Marshal | 2682 | 2834 | 2 |
| Small Struct | MessagePack | Marshal | 3529 | 8321 | 9 |
| Small Struct | JSON | Marshal | 6352 | 3218 | 2 |
| Small Struct | Sonic | Marshal | 8409 | 3317 | 3 |
| Small Struct | CBOR | Unmarshal | 1735 | 896 | 22 |
| Small Struct | BEVE | Unmarshal | 1773 | 2616 | 4 |
| Small Struct | MessagePack | Unmarshal | 1978 | 1496 | 34 |
| Small Struct | Sonic | Unmarshal | 3432 | 1963 | 6 |
| Small Struct | JSON | Unmarshal | 27858 | 7528 | 101 |
