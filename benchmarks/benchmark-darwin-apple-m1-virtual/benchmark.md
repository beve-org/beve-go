# Apple M1 (Virtual) — Darwin

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE ZeroCopy | Marshal | 68151 | 26 | 0 |
| Large Payload | BEVE | Marshal | 88342 | 188453 | 1 |
| Large Payload | CBOR | Marshal | 137094 | 180366 | 1 |
| Large Payload | MessagePack | Marshal | 208430 | 526750 | 115 |
| Large Payload | JSON | Marshal | 322728 | 205113 | 8 |
| Large Payload | Sonic | Marshal | 414710 | 213824 | 3 |
| Large Payload | BEVE | Unmarshal | 160268 | 277203 | 418 |
| Large Payload | Sonic | Unmarshal | 276865 | 350289 | 211 |
| Large Payload | MessagePack | Unmarshal | 385638 | 365290 | 6677 |
| Large Payload | CBOR | Unmarshal | 501016 | 311641 | 6355 |
| Large Payload | JSON | Unmarshal | 1740353 | 535850 | 6972 |
| Medium Payload | BEVE ZeroCopy | Marshal | 5067 | 0 | 0 |
| Medium Payload | BEVE | Marshal | 7440 | 16388 | 1 |
| Medium Payload | CBOR | Marshal | 16254 | 20494 | 1 |
| Medium Payload | MessagePack | Marshal | 23514 | 65778 | 22 |
| Medium Payload | JSON | Marshal | 28299 | 19301 | 8 |
| Medium Payload | Sonic | Marshal | 37146 | 20731 | 3 |
| Medium Payload | BEVE | Unmarshal | 15381 | 22651 | 59 |
| Medium Payload | Sonic | Unmarshal | 30611 | 42100 | 33 |
| Medium Payload | MessagePack | Unmarshal | 38195 | 35934 | 671 |
| Medium Payload | CBOR | Unmarshal | 48432 | 30456 | 626 |
| Medium Payload | JSON | Unmarshal | 204002 | 69752 | 903 |
| Small Struct | BEVE ZeroCopy | Marshal | 235 | 0 | 0 |
| Small Struct | BEVE | Marshal | 268 | 512 | 1 |
| Small Struct | CBOR | Marshal | 460 | 384 | 1 |
| Small Struct | MessagePack | Marshal | 1581 | 4104 | 8 |
| Small Struct | JSON | Marshal | 2702 | 2304 | 1 |
| Small Struct | Sonic | Marshal | 5097 | 3103 | 2 |
| Small Struct | BEVE | Unmarshal | 881 | 1592 | 4 |
| Small Struct | Sonic | Unmarshal | 1440 | 1466 | 6 |
| Small Struct | MessagePack | Unmarshal | 1863 | 2080 | 45 |
| Small Struct | CBOR | Unmarshal | 4403 | 3880 | 82 |
| Small Struct | JSON | Unmarshal | 14490 | 4720 | 84 |
