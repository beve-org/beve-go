# Neoverse-N2 — Linux

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE ZeroCopy | Marshal | 67472 | 65 | 0 |
| Large Payload | BEVE | Marshal | 105053 | 188510 | 1 |
| Large Payload | CBOR | Marshal | 181114 | 188663 | 1 |
| Large Payload | MessagePack | Marshal | 275574 | 526802 | 115 |
| Large Payload | Sonic | Marshal | 299382 | 214497 | 3 |
| Large Payload | JSON | Marshal | 389952 | 213317 | 8 |
| Large Payload | BEVE | Unmarshal | 223623 | 269579 | 419 |
| Large Payload | Sonic | Unmarshal | 286667 | 389720 | 209 |
| Large Payload | MessagePack | Unmarshal | 536740 | 368099 | 6746 |
| Large Payload | CBOR | Unmarshal | 635499 | 302522 | 6179 |
| Large Payload | JSON | Unmarshal | 1979793 | 533578 | 7040 |
| Medium Payload | BEVE ZeroCopy | Marshal | 6774 | 3 | 0 |
| Medium Payload | BEVE | Marshal | 9545 | 18435 | 1 |
| Medium Payload | CBOR | Marshal | 19148 | 20499 | 1 |
| Medium Payload | Sonic | Marshal | 29463 | 20813 | 3 |
| Medium Payload | MessagePack | Marshal | 31558 | 65782 | 22 |
| Medium Payload | JSON | Marshal | 39601 | 21991 | 8 |
| Medium Payload | BEVE | Unmarshal | 23414 | 30046 | 59 |
| Medium Payload | Sonic | Unmarshal | 25580 | 34125 | 33 |
| Medium Payload | MessagePack | Unmarshal | 47912 | 31983 | 586 |
| Medium Payload | CBOR | Unmarshal | 63597 | 31160 | 641 |
| Medium Payload | JSON | Unmarshal | 218063 | 63096 | 814 |
| Small Struct | Sonic | Marshal | 665 | 311 | 2 |
| Small Struct | BEVE ZeroCopy | Marshal | 666 | 0 | 0 |
| Small Struct | CBOR | Marshal | 983 | 896 | 1 |
| Small Struct | BEVE | Marshal | 1068 | 2048 | 1 |
| Small Struct | MessagePack | Marshal | 1397 | 2056 | 7 |
| Small Struct | JSON | Marshal | 3694 | 2304 | 1 |
| Small Struct | BEVE | Unmarshal | 1821 | 3512 | 4 |
| Small Struct | Sonic | Unmarshal | 2331 | 4106 | 6 |
| Small Struct | MessagePack | Unmarshal | 2430 | 1600 | 36 |
| Small Struct | CBOR | Unmarshal | 4366 | 2440 | 53 |
| Small Struct | JSON | Unmarshal | 7099 | 2024 | 35 |
