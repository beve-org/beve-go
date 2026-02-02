# Neoverse-N2 — Linux

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE ZeroCopy | Marshal | 65290 | 52 | 0 |
| Large Payload | BEVE | Marshal | 99335 | 180343 | 1 |
| Large Payload | CBOR | Marshal | 183295 | 188873 | 1 |
| Large Payload | MessagePack | Marshal | 268476 | 526804 | 115 |
| Large Payload | Sonic | Marshal | 310174 | 222771 | 3 |
| Large Payload | JSON | Marshal | 406519 | 229784 | 8 |
| Large Payload | BEVE | Unmarshal | 214534 | 257092 | 418 |
| Large Payload | Sonic | Unmarshal | 296499 | 421927 | 213 |
| Large Payload | MessagePack | Unmarshal | 525817 | 360927 | 6577 |
| Large Payload | CBOR | Unmarshal | 647842 | 313067 | 6391 |
| Large Payload | JSON | Unmarshal | 1872402 | 507132 | 6526 |
| Medium Payload | BEVE ZeroCopy | Marshal | 8232 | 0 | 0 |
| Medium Payload | BEVE | Marshal | 10119 | 19081 | 1 |
| Medium Payload | CBOR | Marshal | 16635 | 18443 | 1 |
| Medium Payload | JSON | Marshal | 30957 | 16610 | 8 |
| Medium Payload | MessagePack | Marshal | 31471 | 65782 | 22 |
| Medium Payload | Sonic | Marshal | 34033 | 25061 | 3 |
| Medium Payload | BEVE | Unmarshal | 23635 | 32735 | 59 |
| Medium Payload | Sonic | Unmarshal | 26531 | 33618 | 33 |
| Medium Payload | CBOR | Unmarshal | 53344 | 24104 | 499 |
| Medium Payload | MessagePack | Unmarshal | 55722 | 40321 | 751 |
| Medium Payload | JSON | Unmarshal | 219567 | 62296 | 828 |
| Small Struct | BEVE ZeroCopy | Marshal | 338 | 0 | 0 |
| Small Struct | BEVE | Marshal | 887 | 1536 | 1 |
| Small Struct | CBOR | Marshal | 1519 | 1792 | 1 |
| Small Struct | MessagePack | Marshal | 2320 | 4104 | 8 |
| Small Struct | Sonic | Marshal | 2730 | 2120 | 2 |
| Small Struct | JSON | Marshal | 2935 | 1792 | 1 |
| Small Struct | BEVE | Unmarshal | 1774 | 3512 | 4 |
| Small Struct | MessagePack | Unmarshal | 3398 | 2688 | 56 |
| Small Struct | Sonic | Unmarshal | 3402 | 5712 | 6 |
| Small Struct | CBOR | Unmarshal | 7768 | 4776 | 102 |
| Small Struct | JSON | Unmarshal | 18458 | 4776 | 86 |
