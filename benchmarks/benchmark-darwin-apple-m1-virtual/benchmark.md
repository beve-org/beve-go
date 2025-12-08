# Apple M1 (Virtual) — Darwin

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE ZeroCopy | Marshal | 93479 | 26 | 0 |
| Large Payload | BEVE | Marshal | 134444 | 188483 | 1 |
| Large Payload | CBOR | Marshal | 152693 | 196761 | 1 |
| Large Payload | MessagePack | Marshal | 336034 | 526753 | 115 |
| Large Payload | JSON | Marshal | 483725 | 213305 | 8 |
| Large Payload | Sonic | Marshal | 577802 | 206354 | 3 |
| Large Payload | BEVE | Unmarshal | 271825 | 286420 | 419 |
| Large Payload | MessagePack | Unmarshal | 407729 | 343205 | 6233 |
| Large Payload | Sonic | Unmarshal | 415329 | 381922 | 213 |
| Large Payload | CBOR | Unmarshal | 737173 | 324106 | 6616 |
| Large Payload | JSON | Unmarshal | 2548131 | 541618 | 7199 |
| Medium Payload | BEVE ZeroCopy | Marshal | 8049 | 3 | 0 |
| Medium Payload | BEVE | Marshal | 17135 | 21768 | 1 |
| Medium Payload | CBOR | Marshal | 23408 | 19088 | 1 |
| Medium Payload | MessagePack | Marshal | 33905 | 65778 | 22 |
| Medium Payload | JSON | Marshal | 47048 | 20710 | 8 |
| Medium Payload | Sonic | Marshal | 62385 | 24882 | 3 |
| Medium Payload | BEVE | Unmarshal | 34401 | 36318 | 59 |
| Medium Payload | Sonic | Unmarshal | 49832 | 41605 | 33 |
| Medium Payload | MessagePack | Unmarshal | 70393 | 43022 | 803 |
| Medium Payload | CBOR | Unmarshal | 95742 | 34872 | 718 |
| Medium Payload | JSON | Unmarshal | 239372 | 53304 | 711 |
| Small Struct | BEVE ZeroCopy | Marshal | 1055 | 0 | 0 |
| Small Struct | CBOR | Marshal | 1286 | 768 | 1 |
| Small Struct | JSON | Marshal | 1755 | 576 | 1 |
| Small Struct | BEVE | Marshal | 1832 | 2305 | 1 |
| Small Struct | Sonic | Marshal | 2191 | 656 | 2 |
| Small Struct | MessagePack | Marshal | 3359 | 4104 | 8 |
| Small Struct | BEVE | Unmarshal | 3756 | 2360 | 4 |
| Small Struct | CBOR | Unmarshal | 4496 | 1480 | 34 |
| Small Struct | Sonic | Unmarshal | 4902 | 4559 | 6 |
| Small Struct | MessagePack | Unmarshal | 8421 | 3840 | 80 |
| Small Struct | JSON | Unmarshal | 21256 | 4072 | 64 |
