# Apple M1 (Virtual) — Darwin

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE ZeroCopy | Marshal | 65363 | 26 | 0 |
| Large Payload | BEVE | Marshal | 127754 | 204881 | 1 |
| Large Payload | CBOR | Marshal | 215839 | 213127 | 1 |
| Large Payload | MessagePack | Marshal | 334695 | 526756 | 115 |
| Large Payload | JSON | Marshal | 485031 | 213332 | 8 |
| Large Payload | Sonic | Marshal | 553781 | 222313 | 3 |
| Large Payload | BEVE | Unmarshal | 263904 | 264143 | 417 |
| Large Payload | Sonic | Unmarshal | 373869 | 349205 | 211 |
| Large Payload | MessagePack | Unmarshal | 563073 | 328946 | 5949 |
| Large Payload | CBOR | Unmarshal | 683782 | 312587 | 6373 |
| Large Payload | JSON | Unmarshal | 2180935 | 556483 | 7197 |
| Medium Payload | BEVE ZeroCopy | Marshal | 8339 | 3 | 0 |
| Medium Payload | BEVE | Marshal | 13608 | 19081 | 1 |
| Medium Payload | CBOR | Marshal | 25045 | 20493 | 1 |
| Medium Payload | MessagePack | Marshal | 34500 | 65778 | 22 |
| Medium Payload | JSON | Marshal | 48175 | 24806 | 8 |
| Medium Payload | Sonic | Marshal | 49353 | 20720 | 3 |
| Medium Payload | BEVE | Unmarshal | 31631 | 31325 | 59 |
| Medium Payload | Sonic | Unmarshal | 52328 | 45899 | 33 |
| Medium Payload | MessagePack | Unmarshal | 58817 | 33789 | 621 |
| Medium Payload | CBOR | Unmarshal | 66087 | 28456 | 583 |
| Medium Payload | JSON | Unmarshal | 201498 | 48471 | 676 |
| Small Struct | BEVE ZeroCopy | Marshal | 360 | 0 | 0 |
| Small Struct | BEVE | Marshal | 1040 | 2688 | 1 |
| Small Struct | CBOR | Marshal | 2697 | 2688 | 1 |
| Small Struct | JSON | Marshal | 3681 | 3073 | 1 |
| Small Struct | MessagePack | Marshal | 3698 | 4104 | 8 |
| Small Struct | Sonic | Marshal | 5534 | 2711 | 2 |
| Small Struct | BEVE | Unmarshal | 952 | 1336 | 4 |
| Small Struct | MessagePack | Unmarshal | 4321 | 3232 | 69 |
| Small Struct | Sonic | Unmarshal | 4693 | 5607 | 6 |
| Small Struct | CBOR | Unmarshal | 5298 | 3048 | 64 |
| Small Struct | JSON | Unmarshal | 10208 | 2376 | 46 |
