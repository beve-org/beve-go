# Apple M1 (Virtual) — Darwin

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE ZeroCopy | Marshal | 62205 | 39 | 0 |
| Large Payload | BEVE | Marshal | 114942 | 196661 | 1 |
| Large Payload | CBOR | Marshal | 145131 | 188568 | 1 |
| Large Payload | MessagePack | Marshal | 291631 | 526751 | 115 |
| Large Payload | JSON | Marshal | 434299 | 205139 | 8 |
| Large Payload | Sonic | Marshal | 508280 | 214186 | 3 |
| Large Payload | BEVE | Unmarshal | 210793 | 278291 | 417 |
| Large Payload | Sonic | Unmarshal | 316134 | 358506 | 211 |
| Large Payload | MessagePack | Unmarshal | 405944 | 318335 | 5726 |
| Large Payload | CBOR | Unmarshal | 527856 | 310392 | 6315 |
| Large Payload | JSON | Unmarshal | 1961703 | 524972 | 6859 |
| Medium Payload | BEVE ZeroCopy | Marshal | 7752 | 3 | 0 |
| Medium Payload | BEVE | Marshal | 10787 | 20486 | 1 |
| Medium Payload | CBOR | Marshal | 19735 | 27292 | 1 |
| Medium Payload | MessagePack | Marshal | 27987 | 65778 | 22 |
| Medium Payload | JSON | Marshal | 37819 | 19302 | 8 |
| Medium Payload | Sonic | Marshal | 46973 | 20745 | 3 |
| Medium Payload | BEVE | Unmarshal | 19533 | 30205 | 59 |
| Medium Payload | Sonic | Unmarshal | 31651 | 36593 | 33 |
| Medium Payload | CBOR | Unmarshal | 53111 | 32584 | 671 |
| Medium Payload | MessagePack | Unmarshal | 55532 | 41126 | 770 |
| Medium Payload | JSON | Unmarshal | 224987 | 70504 | 893 |
| Small Struct | BEVE ZeroCopy | Marshal | 310 | 0 | 0 |
| Small Struct | MessagePack | Marshal | 709 | 1032 | 6 |
| Small Struct | CBOR | Marshal | 1407 | 1792 | 1 |
| Small Struct | BEVE | Marshal | 1796 | 2689 | 1 |
| Small Struct | JSON | Marshal | 2388 | 1280 | 1 |
| Small Struct | Sonic | Marshal | 4099 | 1836 | 2 |
| Small Struct | BEVE | Unmarshal | 751 | 728 | 4 |
| Small Struct | MessagePack | Unmarshal | 1423 | 1024 | 24 |
| Small Struct | CBOR | Unmarshal | 3076 | 1928 | 43 |
| Small Struct | Sonic | Unmarshal | 3959 | 5397 | 6 |
| Small Struct | JSON | Unmarshal | 22477 | 7256 | 92 |
