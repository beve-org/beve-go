# Apple M1 (Virtual) — Darwin

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE ZeroCopy | Marshal | 61174 | 26 | 0 |
| Large Payload | BEVE | Marshal | 113286 | 204867 | 1 |
| Large Payload | CBOR | Marshal | 201499 | 196759 | 1 |
| Large Payload | MessagePack | Marshal | 272515 | 526756 | 115 |
| Large Payload | JSON | Marshal | 410187 | 221551 | 8 |
| Large Payload | Sonic | Marshal | 534184 | 222362 | 3 |
| Large Payload | BEVE | Unmarshal | 292835 | 258701 | 417 |
| Large Payload | Sonic | Unmarshal | 342728 | 340823 | 213 |
| Large Payload | MessagePack | Unmarshal | 585961 | 357212 | 6522 |
| Large Payload | CBOR | Unmarshal | 718251 | 312906 | 6377 |
| Large Payload | JSON | Unmarshal | 2221389 | 525894 | 6936 |
| Medium Payload | BEVE ZeroCopy | Marshal | 6640 | 0 | 0 |
| Medium Payload | BEVE | Marshal | 12347 | 19076 | 1 |
| Medium Payload | CBOR | Marshal | 18990 | 21774 | 1 |
| Medium Payload | MessagePack | Marshal | 28478 | 65778 | 22 |
| Medium Payload | JSON | Marshal | 29824 | 19302 | 8 |
| Medium Payload | Sonic | Marshal | 43670 | 24900 | 3 |
| Medium Payload | BEVE | Unmarshal | 16287 | 26300 | 58 |
| Medium Payload | Sonic | Unmarshal | 29722 | 31275 | 33 |
| Medium Payload | MessagePack | Unmarshal | 45922 | 37486 | 697 |
| Medium Payload | CBOR | Unmarshal | 59416 | 36023 | 744 |
| Medium Payload | JSON | Unmarshal | 149927 | 41056 | 530 |
| Small Struct | BEVE | Marshal | 225 | 320 | 1 |
| Small Struct | BEVE ZeroCopy | Marshal | 401 | 0 | 0 |
| Small Struct | CBOR | Marshal | 748 | 1152 | 1 |
| Small Struct | MessagePack | Marshal | 1718 | 2056 | 7 |
| Small Struct | JSON | Marshal | 3389 | 2688 | 1 |
| Small Struct | Sonic | Marshal | 5279 | 2740 | 2 |
| Small Struct | BEVE | Unmarshal | 932 | 2104 | 4 |
| Small Struct | MessagePack | Unmarshal | 3687 | 3936 | 83 |
| Small Struct | CBOR | Unmarshal | 4081 | 2792 | 60 |
| Small Struct | Sonic | Unmarshal | 4459 | 5863 | 6 |
| Small Struct | JSON | Unmarshal | 21689 | 7752 | 108 |
