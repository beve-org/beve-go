# Neoverse-N2 — Linux

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE ZeroCopy | Marshal | 66887 | 65 | 0 |
| Large Payload | BEVE | Marshal | 106482 | 180450 | 1 |
| Large Payload | CBOR | Marshal | 193552 | 205003 | 1 |
| Large Payload | MessagePack | Marshal | 272970 | 526804 | 115 |
| Large Payload | Sonic | Marshal | 300741 | 208154 | 3 |
| Large Payload | JSON | Marshal | 390097 | 213422 | 8 |
| Large Payload | BEVE | Unmarshal | 223942 | 270121 | 418 |
| Large Payload | Sonic | Unmarshal | 282297 | 387529 | 213 |
| Large Payload | MessagePack | Unmarshal | 530656 | 360159 | 6582 |
| Large Payload | CBOR | Unmarshal | 641861 | 308394 | 6296 |
| Large Payload | JSON | Unmarshal | 2032335 | 550308 | 7233 |
| Medium Payload | BEVE ZeroCopy | Marshal | 7152 | 3 | 0 |
| Medium Payload | BEVE | Marshal | 9747 | 18441 | 1 |
| Medium Payload | CBOR | Marshal | 17721 | 18453 | 1 |
| Medium Payload | Sonic | Marshal | 25105 | 18808 | 3 |
| Medium Payload | MessagePack | Marshal | 30673 | 65783 | 22 |
| Medium Payload | JSON | Marshal | 33027 | 18662 | 8 |
| Medium Payload | Sonic | Unmarshal | 23710 | 28828 | 33 |
| Medium Payload | BEVE | Unmarshal | 24198 | 32735 | 59 |
| Medium Payload | MessagePack | Unmarshal | 43347 | 27517 | 492 |
| Medium Payload | CBOR | Unmarshal | 58796 | 27688 | 570 |
| Medium Payload | JSON | Unmarshal | 164037 | 42264 | 580 |
| Small Struct | BEVE ZeroCopy | Marshal | 709 | 0 | 0 |
| Small Struct | BEVE | Marshal | 1314 | 2689 | 1 |
| Small Struct | CBOR | Marshal | 1637 | 1792 | 1 |
| Small Struct | Sonic | Marshal | 3475 | 2746 | 2 |
| Small Struct | MessagePack | Marshal | 3789 | 8201 | 9 |
| Small Struct | JSON | Marshal | 4260 | 2688 | 1 |
| Small Struct | BEVE | Unmarshal | 1204 | 1592 | 4 |
| Small Struct | MessagePack | Unmarshal | 2688 | 1880 | 41 |
| Small Struct | Sonic | Unmarshal | 2785 | 4350 | 6 |
| Small Struct | CBOR | Unmarshal | 6067 | 3592 | 77 |
| Small Struct | JSON | Unmarshal | 22586 | 7592 | 103 |
