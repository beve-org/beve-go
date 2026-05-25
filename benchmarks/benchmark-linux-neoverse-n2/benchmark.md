# Neoverse-N2 — Linux

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE ZeroCopy | Marshal | 70050 | 65 | 0 |
| Large Payload | BEVE | Marshal | 122424 | 196851 | 1 |
| Large Payload | CBOR | Marshal | 188803 | 196858 | 1 |
| Large Payload | MessagePack | Marshal | 287906 | 526802 | 115 |
| Large Payload | Sonic | Marshal | 318830 | 225730 | 3 |
| Large Payload | JSON | Marshal | 417086 | 221614 | 8 |
| Large Payload | BEVE | Unmarshal | 231735 | 272269 | 418 |
| Large Payload | Sonic | Unmarshal | 287313 | 374205 | 207 |
| Large Payload | MessagePack | Unmarshal | 523121 | 352703 | 6435 |
| Large Payload | CBOR | Unmarshal | 690303 | 335211 | 6843 |
| Large Payload | JSON | Unmarshal | 2003047 | 548909 | 7102 |
| Medium Payload | BEVE ZeroCopy | Marshal | 7635 | 3 | 0 |
| Medium Payload | BEVE | Marshal | 9899 | 18439 | 1 |
| Medium Payload | CBOR | Marshal | 22284 | 24594 | 1 |
| Medium Payload | JSON | Marshal | 31857 | 16612 | 8 |
| Medium Payload | Sonic | Marshal | 34181 | 25413 | 3 |
| Medium Payload | MessagePack | Marshal | 35654 | 65782 | 22 |
| Medium Payload | BEVE | Unmarshal | 23521 | 28670 | 59 |
| Medium Payload | Sonic | Unmarshal | 32120 | 44913 | 33 |
| Medium Payload | MessagePack | Unmarshal | 42649 | 25917 | 459 |
| Medium Payload | CBOR | Unmarshal | 55699 | 24728 | 512 |
| Medium Payload | JSON | Unmarshal | 232190 | 65720 | 863 |
| Small Struct | CBOR | Marshal | 737 | 576 | 1 |
| Small Struct | BEVE ZeroCopy | Marshal | 773 | 0 | 0 |
| Small Struct | BEVE | Marshal | 1108 | 2304 | 1 |
| Small Struct | JSON | Marshal | 1265 | 640 | 1 |
| Small Struct | Sonic | Marshal | 2354 | 1602 | 2 |
| Small Struct | MessagePack | Marshal | 2397 | 4104 | 8 |
| Small Struct | BEVE | Unmarshal | 1200 | 1464 | 4 |
| Small Struct | MessagePack | Unmarshal | 1703 | 872 | 21 |
| Small Struct | Sonic | Unmarshal | 3528 | 5747 | 6 |
| Small Struct | CBOR | Unmarshal | 7714 | 4648 | 98 |
| Small Struct | JSON | Unmarshal | 16020 | 4424 | 75 |
