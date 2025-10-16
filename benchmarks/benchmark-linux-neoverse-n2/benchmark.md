# Neoverse-N2 — Linux

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE ZeroCopy | Marshal | 65782 | 233 | 2 |
| Large Payload | BEVE | Marshal | 111334 | 182218 | 3 |
| Large Payload | CBOR | Marshal | 182197 | 190672 | 2 |
| Large Payload | MessagePack | Marshal | 278342 | 526860 | 115 |
| Large Payload | Sonic | Marshal | 289568 | 209945 | 4 |
| Large Payload | JSON | Marshal | 396162 | 223572 | 9 |
| Large Payload | BEVE | Unmarshal | 229203 | 277389 | 418 |
| Large Payload | Sonic | Unmarshal | 285415 | 374874 | 213 |
| Large Payload | MessagePack | Unmarshal | 517091 | 354943 | 6477 |
| Large Payload | CBOR | Unmarshal | 673677 | 333706 | 6793 |
| Large Payload | JSON | Unmarshal | 2070322 | 563611 | 7366 |
| Medium Payload | BEVE ZeroCopy | Marshal | 7277 | 141 | 2 |
| Medium Payload | BEVE | Marshal | 11583 | 21917 | 3 |
| Medium Payload | CBOR | Marshal | 18049 | 19214 | 2 |
| Medium Payload | Sonic | Marshal | 29645 | 22251 | 4 |
| Medium Payload | MessagePack | Marshal | 31354 | 65839 | 22 |
| Medium Payload | JSON | Marshal | 35788 | 20824 | 9 |
| Medium Payload | BEVE | Unmarshal | 23481 | 33247 | 59 |
| Medium Payload | Sonic | Unmarshal | 25677 | 33683 | 33 |
| Medium Payload | MessagePack | Unmarshal | 49594 | 34191 | 631 |
| Medium Payload | CBOR | Unmarshal | 67328 | 34040 | 698 |
| Medium Payload | JSON | Unmarshal | 196541 | 56296 | 718 |
| Small Struct | BEVE ZeroCopy | Marshal | 539 | 290 | 2 |
| Small Struct | BEVE | Marshal | 1184 | 2081 | 3 |
| Small Struct | CBOR | Marshal | 2015 | 2192 | 2 |
| Small Struct | JSON | Marshal | 2238 | 1297 | 2 |
| Small Struct | MessagePack | Marshal | 3532 | 8322 | 9 |
| Small Struct | Sonic | Marshal | 3888 | 3275 | 3 |
| Small Struct | Sonic | Unmarshal | 1232 | 1323 | 6 |
| Small Struct | CBOR | Unmarshal | 1740 | 616 | 16 |
| Small Struct | BEVE | Unmarshal | 1741 | 3384 | 4 |
| Small Struct | MessagePack | Unmarshal | 4109 | 3296 | 71 |
| Small Struct | JSON | Unmarshal | 17740 | 4648 | 82 |
