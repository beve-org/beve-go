# Neoverse-N2 — Linux

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE ZeroCopy | Marshal | 70362 | 259 | 2 |
| Large Payload | BEVE | Marshal | 124623 | 198684 | 3 |
| Large Payload | CBOR | Marshal | 189399 | 190515 | 2 |
| Large Payload | MessagePack | Marshal | 284404 | 526859 | 115 |
| Large Payload | Sonic | Marshal | 312984 | 218778 | 4 |
| Large Payload | JSON | Marshal | 399138 | 224989 | 9 |
| Large Payload | BEVE | Unmarshal | 227604 | 259972 | 419 |
| Large Payload | Sonic | Unmarshal | 301957 | 402632 | 211 |
| Large Payload | MessagePack | Unmarshal | 521145 | 346253 | 6303 |
| Large Payload | CBOR | Unmarshal | 659553 | 320506 | 6542 |
| Large Payload | JSON | Unmarshal | 2016778 | 545902 | 7063 |
| Medium Payload | BEVE ZeroCopy | Marshal | 7590 | 141 | 2 |
| Medium Payload | BEVE | Marshal | 10913 | 19231 | 3 |
| Medium Payload | CBOR | Marshal | 17231 | 18567 | 2 |
| Medium Payload | JSON | Marshal | 30771 | 16688 | 9 |
| Medium Payload | Sonic | Marshal | 33182 | 25227 | 4 |
| Medium Payload | MessagePack | Marshal | 33662 | 65838 | 22 |
| Medium Payload | BEVE | Unmarshal | 24334 | 30046 | 59 |
| Medium Payload | Sonic | Unmarshal | 32813 | 44826 | 33 |
| Medium Payload | MessagePack | Unmarshal | 52291 | 36447 | 679 |
| Medium Payload | CBOR | Unmarshal | 73041 | 38000 | 776 |
| Medium Payload | JSON | Unmarshal | 202217 | 54328 | 739 |
| Small Struct | BEVE ZeroCopy | Marshal | 383 | 289 | 2 |
| Small Struct | CBOR | Marshal | 900 | 720 | 2 |
| Small Struct | BEVE | Marshal | 1207 | 1825 | 3 |
| Small Struct | Sonic | Marshal | 1504 | 1085 | 3 |
| Small Struct | JSON | Marshal | 2232 | 1296 | 2 |
| Small Struct | MessagePack | Marshal | 4117 | 8321 | 9 |
| Small Struct | BEVE | Unmarshal | 657 | 312 | 3 |
| Small Struct | MessagePack | Unmarshal | 1549 | 688 | 17 |
| Small Struct | Sonic | Unmarshal | 2395 | 3791 | 6 |
| Small Struct | CBOR | Unmarshal | 5309 | 2888 | 63 |
| Small Struct | JSON | Unmarshal | 14217 | 4104 | 65 |
