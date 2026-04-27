# Neoverse-N2 — Linux

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE ZeroCopy | Marshal | 71885 | 52 | 0 |
| Large Payload | BEVE | Marshal | 108447 | 188616 | 1 |
| Large Payload | CBOR | Marshal | 186654 | 196805 | 1 |
| Large Payload | MessagePack | Marshal | 281575 | 526802 | 115 |
| Large Payload | Sonic | Marshal | 305526 | 214719 | 3 |
| Large Payload | JSON | Marshal | 383764 | 205177 | 8 |
| Large Payload | BEVE | Unmarshal | 229211 | 280591 | 417 |
| Large Payload | Sonic | Unmarshal | 289170 | 392480 | 211 |
| Large Payload | MessagePack | Unmarshal | 492153 | 322981 | 5828 |
| Large Payload | CBOR | Unmarshal | 645605 | 310122 | 6318 |
| Large Payload | JSON | Unmarshal | 1919752 | 513516 | 6762 |
| Medium Payload | BEVE ZeroCopy | Marshal | 7886 | 6 | 0 |
| Medium Payload | BEVE | Marshal | 10414 | 19084 | 1 |
| Medium Payload | CBOR | Marshal | 21785 | 24597 | 1 |
| Medium Payload | MessagePack | Marshal | 31238 | 65782 | 22 |
| Medium Payload | Sonic | Marshal | 38154 | 27801 | 3 |
| Medium Payload | JSON | Marshal | 44829 | 24814 | 8 |
| Medium Payload | BEVE | Unmarshal | 22446 | 26494 | 59 |
| Medium Payload | Sonic | Unmarshal | 33016 | 45205 | 33 |
| Medium Payload | MessagePack | Unmarshal | 49967 | 33615 | 620 |
| Medium Payload | CBOR | Unmarshal | 73942 | 38056 | 781 |
| Medium Payload | JSON | Unmarshal | 219759 | 61960 | 817 |
| Small Struct | BEVE ZeroCopy | Marshal | 342 | 0 | 0 |
| Small Struct | BEVE | Marshal | 979 | 1792 | 1 |
| Small Struct | JSON | Marshal | 1196 | 576 | 1 |
| Small Struct | Sonic | Marshal | 1313 | 812 | 2 |
| Small Struct | CBOR | Marshal | 1981 | 2304 | 1 |
| Small Struct | MessagePack | Marshal | 3744 | 8202 | 9 |
| Small Struct | BEVE | Unmarshal | 1500 | 2616 | 4 |
| Small Struct | Sonic | Unmarshal | 2841 | 4490 | 6 |
| Small Struct | MessagePack | Unmarshal | 5096 | 4288 | 90 |
| Small Struct | CBOR | Unmarshal | 7312 | 4392 | 94 |
| Small Struct | JSON | Unmarshal | 24179 | 7816 | 110 |
