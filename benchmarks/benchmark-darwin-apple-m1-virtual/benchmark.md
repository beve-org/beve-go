# Apple M1 (Virtual) — Darwin

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE ZeroCopy | Marshal | 77516 | 207 | 2 |
| Large Payload | BEVE | Marshal | 135130 | 189626 | 3 |
| Large Payload | CBOR | Marshal | 171856 | 197666 | 2 |
| Large Payload | MessagePack | Marshal | 326382 | 526813 | 115 |
| Large Payload | JSON | Marshal | 516647 | 221931 | 9 |
| Large Payload | Sonic | Marshal | 575604 | 223596 | 4 |
| Large Payload | BEVE | Unmarshal | 249258 | 280722 | 417 |
| Large Payload | Sonic | Unmarshal | 378530 | 347961 | 211 |
| Large Payload | MessagePack | Unmarshal | 483353 | 344130 | 6247 |
| Large Payload | CBOR | Unmarshal | 554651 | 305195 | 6226 |
| Large Payload | JSON | Unmarshal | 2266746 | 514957 | 6726 |
| Medium Payload | BEVE ZeroCopy | Marshal | 6736 | 134 | 2 |
| Medium Payload | BEVE | Marshal | 15011 | 24732 | 3 |
| Medium Payload | CBOR | Marshal | 16404 | 20561 | 2 |
| Medium Payload | JSON | Marshal | 30220 | 18755 | 9 |
| Medium Payload | MessagePack | Marshal | 34261 | 65834 | 22 |
| Medium Payload | Sonic | Marshal | 56034 | 24856 | 4 |
| Medium Payload | BEVE | Unmarshal | 16996 | 24027 | 57 |
| Medium Payload | Sonic | Unmarshal | 43833 | 33259 | 33 |
| Medium Payload | MessagePack | Unmarshal | 75836 | 47871 | 907 |
| Medium Payload | CBOR | Unmarshal | 104023 | 32151 | 667 |
| Medium Payload | JSON | Unmarshal | 229118 | 65720 | 843 |
| Small Struct | BEVE | Marshal | 560 | 736 | 3 |
| Small Struct | BEVE ZeroCopy | Marshal | 1167 | 289 | 2 |
| Small Struct | CBOR | Marshal | 1369 | 1424 | 2 |
| Small Struct | JSON | Marshal | 2366 | 1681 | 2 |
| Small Struct | Sonic | Marshal | 2764 | 1326 | 3 |
| Small Struct | MessagePack | Marshal | 4227 | 4224 | 8 |
| Small Struct | BEVE | Unmarshal | 1369 | 1336 | 4 |
| Small Struct | Sonic | Unmarshal | 2419 | 2187 | 6 |
| Small Struct | MessagePack | Unmarshal | 3991 | 2144 | 47 |
| Small Struct | CBOR | Unmarshal | 5322 | 3560 | 76 |
| Small Struct | JSON | Unmarshal | 7530 | 2024 | 35 |
