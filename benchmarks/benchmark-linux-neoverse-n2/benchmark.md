# Neoverse-N2 — Linux

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE ZeroCopy | Marshal | 68898 | 65 | 0 |
| Large Payload | BEVE | Marshal | 107502 | 188576 | 1 |
| Large Payload | CBOR | Marshal | 194059 | 196884 | 1 |
| Large Payload | MessagePack | Marshal | 273301 | 526803 | 115 |
| Large Payload | Sonic | Marshal | 315554 | 222599 | 3 |
| Large Payload | JSON | Marshal | 397652 | 221668 | 8 |
| Large Payload | BEVE | Unmarshal | 222844 | 268619 | 418 |
| Large Payload | Sonic | Unmarshal | 277031 | 359889 | 211 |
| Large Payload | MessagePack | Unmarshal | 510095 | 335111 | 6076 |
| Large Payload | CBOR | Unmarshal | 688475 | 334683 | 6811 |
| Large Payload | JSON | Unmarshal | 2057129 | 563619 | 7346 |
| Medium Payload | BEVE ZeroCopy | Marshal | 6776 | 6 | 0 |
| Medium Payload | BEVE | Marshal | 10545 | 20489 | 1 |
| Medium Payload | CBOR | Marshal | 18870 | 20496 | 1 |
| Medium Payload | Sonic | Marshal | 22802 | 16896 | 3 |
| Medium Payload | MessagePack | Marshal | 31244 | 65782 | 22 |
| Medium Payload | JSON | Marshal | 41775 | 24814 | 8 |
| Medium Payload | BEVE | Unmarshal | 22031 | 27037 | 58 |
| Medium Payload | Sonic | Unmarshal | 25244 | 32004 | 33 |
| Medium Payload | MessagePack | Unmarshal | 56608 | 40864 | 766 |
| Medium Payload | CBOR | Unmarshal | 69863 | 35352 | 726 |
| Medium Payload | JSON | Unmarshal | 229168 | 66616 | 863 |
| Small Struct | BEVE ZeroCopy | Marshal | 464 | 0 | 0 |
| Small Struct | BEVE | Marshal | 925 | 1792 | 1 |
| Small Struct | CBOR | Marshal | 1549 | 1792 | 1 |
| Small Struct | MessagePack | Marshal | 3546 | 8201 | 9 |
| Small Struct | Sonic | Marshal | 3684 | 2782 | 2 |
| Small Struct | JSON | Marshal | 3842 | 2304 | 1 |
| Small Struct | BEVE | Unmarshal | 1657 | 3000 | 4 |
| Small Struct | Sonic | Unmarshal | 3562 | 6387 | 6 |
| Small Struct | CBOR | Unmarshal | 4834 | 2760 | 59 |
| Small Struct | MessagePack | Unmarshal | 5630 | 5121 | 104 |
| Small Struct | JSON | Unmarshal | 10120 | 2472 | 49 |
