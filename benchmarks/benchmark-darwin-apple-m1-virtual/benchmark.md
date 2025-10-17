# Apple M1 (Virtual) — Darwin

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE ZeroCopy | Marshal | 66443 | 39 | 0 |
| Large Payload | BEVE | Marshal | 103030 | 188493 | 1 |
| Large Payload | CBOR | Marshal | 195993 | 172146 | 1 |
| Large Payload | MessagePack | Marshal | 285064 | 526753 | 115 |
| Large Payload | JSON | Marshal | 448157 | 221499 | 8 |
| Large Payload | Sonic | Marshal | 547675 | 222443 | 3 |
| Large Payload | BEVE | Unmarshal | 253359 | 267312 | 419 |
| Large Payload | Sonic | Unmarshal | 360089 | 348275 | 213 |
| Large Payload | MessagePack | Unmarshal | 524601 | 338771 | 6153 |
| Large Payload | CBOR | Unmarshal | 773636 | 309097 | 6303 |
| Large Payload | JSON | Unmarshal | 2393677 | 566957 | 7325 |
| Medium Payload | BEVE ZeroCopy | Marshal | 6414 | 3 | 0 |
| Medium Payload | BEVE | Marshal | 12849 | 18438 | 1 |
| Medium Payload | CBOR | Marshal | 16773 | 18441 | 1 |
| Medium Payload | MessagePack | Marshal | 25640 | 33005 | 21 |
| Medium Payload | JSON | Marshal | 49604 | 24810 | 8 |
| Medium Payload | Sonic | Marshal | 52062 | 20777 | 3 |
| Medium Payload | BEVE | Unmarshal | 23409 | 24860 | 58 |
| Medium Payload | Sonic | Unmarshal | 36925 | 31627 | 33 |
| Medium Payload | MessagePack | Unmarshal | 61797 | 39710 | 745 |
| Medium Payload | CBOR | Unmarshal | 72447 | 29624 | 605 |
| Medium Payload | JSON | Unmarshal | 307250 | 69816 | 906 |
| Small Struct | BEVE ZeroCopy | Marshal | 1021 | 0 | 0 |
| Small Struct | CBOR | Marshal | 1710 | 1536 | 1 |
| Small Struct | BEVE | Marshal | 1774 | 2304 | 1 |
| Small Struct | JSON | Marshal | 2099 | 640 | 1 |
| Small Struct | Sonic | Marshal | 2201 | 933 | 2 |
| Small Struct | MessagePack | Marshal | 2446 | 4104 | 8 |
| Small Struct | BEVE | Unmarshal | 1265 | 1080 | 4 |
| Small Struct | Sonic | Unmarshal | 4032 | 5327 | 6 |
| Small Struct | CBOR | Unmarshal | 5128 | 2696 | 57 |
| Small Struct | MessagePack | Unmarshal | 5741 | 4296 | 90 |
| Small Struct | JSON | Unmarshal | 16599 | 4136 | 66 |
