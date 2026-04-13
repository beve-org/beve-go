# Neoverse-N2 — Linux

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE ZeroCopy | Marshal | 66679 | 65 | 0 |
| Large Payload | BEVE | Marshal | 110054 | 188524 | 1 |
| Large Payload | CBOR | Marshal | 186460 | 188742 | 1 |
| Large Payload | MessagePack | Marshal | 275187 | 526804 | 115 |
| Large Payload | Sonic | Marshal | 299819 | 215581 | 3 |
| Large Payload | JSON | Marshal | 406412 | 221720 | 8 |
| Large Payload | BEVE | Unmarshal | 229712 | 275246 | 419 |
| Large Payload | Sonic | Unmarshal | 292047 | 399531 | 213 |
| Large Payload | MessagePack | Unmarshal | 521098 | 355901 | 6491 |
| Large Payload | CBOR | Unmarshal | 645631 | 310618 | 6327 |
| Large Payload | JSON | Unmarshal | 2026342 | 552907 | 7229 |
| Medium Payload | BEVE ZeroCopy | Marshal | 6587 | 3 | 0 |
| Medium Payload | BEVE | Marshal | 10772 | 20484 | 1 |
| Medium Payload | CBOR | Marshal | 18028 | 19090 | 1 |
| Medium Payload | Sonic | Marshal | 25061 | 18824 | 3 |
| Medium Payload | MessagePack | Marshal | 31698 | 65782 | 22 |
| Medium Payload | JSON | Marshal | 37814 | 20711 | 8 |
| Medium Payload | BEVE | Unmarshal | 23475 | 30527 | 59 |
| Medium Payload | Sonic | Unmarshal | 31725 | 44099 | 33 |
| Medium Payload | MessagePack | Unmarshal | 52567 | 36768 | 682 |
| Medium Payload | CBOR | Unmarshal | 62833 | 30424 | 627 |
| Medium Payload | JSON | Unmarshal | 188440 | 50520 | 684 |
| Small Struct | BEVE ZeroCopy | Marshal | 333 | 0 | 0 |
| Small Struct | CBOR | Marshal | 853 | 768 | 1 |
| Small Struct | BEVE | Marshal | 886 | 1792 | 1 |
| Small Struct | JSON | Marshal | 1678 | 896 | 1 |
| Small Struct | MessagePack | Marshal | 2310 | 4104 | 8 |
| Small Struct | Sonic | Marshal | 3413 | 2802 | 2 |
| Small Struct | MessagePack | Unmarshal | 1104 | 304 | 9 |
| Small Struct | BEVE | Unmarshal | 1682 | 3000 | 4 |
| Small Struct | Sonic | Unmarshal | 2118 | 2708 | 6 |
| Small Struct | CBOR | Unmarshal | 2336 | 1000 | 24 |
| Small Struct | JSON | Unmarshal | 12390 | 3880 | 58 |
