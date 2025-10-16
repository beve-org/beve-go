# Apple M1 (Virtual) — Darwin

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE ZeroCopy | Marshal | 55558 | 233 | 2 |
| Large Payload | BEVE | Marshal | 89631 | 197026 | 3 |
| Large Payload | CBOR | Marshal | 138519 | 205138 | 2 |
| Large Payload | MessagePack | Marshal | 183673 | 526811 | 115 |
| Large Payload | JSON | Marshal | 306156 | 197403 | 9 |
| Large Payload | Sonic | Marshal | 399222 | 223021 | 4 |
| Large Payload | BEVE | Unmarshal | 164493 | 271091 | 418 |
| Large Payload | Sonic | Unmarshal | 233548 | 319496 | 209 |
| Large Payload | MessagePack | Unmarshal | 380781 | 360732 | 6589 |
| Large Payload | CBOR | Unmarshal | 510255 | 322187 | 6551 |
| Large Payload | JSON | Unmarshal | 1618681 | 526249 | 7042 |
| Medium Payload | BEVE ZeroCopy | Marshal | 5677 | 134 | 2 |
| Medium Payload | BEVE | Marshal | 8702 | 21895 | 3 |
| Medium Payload | CBOR | Marshal | 15932 | 21858 | 2 |
| Medium Payload | MessagePack | Marshal | 22588 | 65833 | 22 |
| Medium Payload | JSON | Marshal | 34126 | 24874 | 9 |
| Medium Payload | Sonic | Marshal | 39247 | 21994 | 4 |
| Medium Payload | BEVE | Unmarshal | 18011 | 25725 | 59 |
| Medium Payload | Sonic | Unmarshal | 33544 | 45905 | 33 |
| Medium Payload | MessagePack | Unmarshal | 33881 | 31501 | 577 |
| Medium Payload | CBOR | Unmarshal | 49245 | 32344 | 669 |
| Medium Payload | JSON | Unmarshal | 205393 | 70136 | 895 |
| Small Struct | BEVE ZeroCopy | Marshal | 649 | 290 | 2 |
| Small Struct | BEVE | Marshal | 812 | 2080 | 3 |
| Small Struct | MessagePack | Marshal | 1613 | 4224 | 8 |
| Small Struct | CBOR | Marshal | 1735 | 2832 | 2 |
| Small Struct | JSON | Marshal | 2207 | 1681 | 2 |
| Small Struct | Sonic | Marshal | 2456 | 1454 | 3 |
| Small Struct | BEVE | Unmarshal | 1187 | 3000 | 4 |
| Small Struct | Sonic | Unmarshal | 2256 | 3197 | 6 |
| Small Struct | MessagePack | Unmarshal | 3473 | 4032 | 86 |
| Small Struct | CBOR | Unmarshal | 4088 | 3496 | 74 |
| Small Struct | JSON | Unmarshal | 14707 | 4584 | 80 |
