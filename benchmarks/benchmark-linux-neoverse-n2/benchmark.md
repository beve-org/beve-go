# Neoverse-N2 — Linux

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE ZeroCopy | Marshal | 63656 | 65 | 0 |
| Large Payload | BEVE | Marshal | 102099 | 188510 | 1 |
| Large Payload | CBOR | Marshal | 182712 | 196778 | 1 |
| Large Payload | MessagePack | Marshal | 251788 | 526794 | 115 |
| Large Payload | Sonic | Marshal | 301656 | 214339 | 3 |
| Large Payload | JSON | Marshal | 379894 | 213291 | 8 |
| Large Payload | BEVE | Unmarshal | 226242 | 281102 | 419 |
| Large Payload | Sonic | Unmarshal | 271557 | 367878 | 211 |
| Large Payload | MessagePack | Unmarshal | 501817 | 345978 | 6289 |
| Large Payload | CBOR | Unmarshal | 638897 | 311657 | 6344 |
| Large Payload | JSON | Unmarshal | 1810862 | 475930 | 6297 |
| Medium Payload | BEVE ZeroCopy | Marshal | 6227 | 6 | 0 |
| Medium Payload | BEVE | Marshal | 7944 | 14339 | 1 |
| Medium Payload | CBOR | Marshal | 17006 | 18447 | 1 |
| Medium Payload | Sonic | Marshal | 26514 | 18827 | 3 |
| Medium Payload | MessagePack | Marshal | 29318 | 65782 | 22 |
| Medium Payload | JSON | Marshal | 36681 | 20710 | 8 |
| Medium Payload | BEVE | Unmarshal | 19991 | 22332 | 59 |
| Medium Payload | Sonic | Unmarshal | 28924 | 38946 | 33 |
| Medium Payload | MessagePack | Unmarshal | 58052 | 42833 | 806 |
| Medium Payload | CBOR | Unmarshal | 66688 | 33688 | 691 |
| Medium Payload | JSON | Unmarshal | 211308 | 61112 | 790 |
| Small Struct | BEVE | Marshal | 485 | 640 | 1 |
| Small Struct | BEVE ZeroCopy | Marshal | 665 | 0 | 0 |
| Small Struct | Sonic | Marshal | 1258 | 798 | 2 |
| Small Struct | CBOR | Marshal | 1845 | 2048 | 1 |
| Small Struct | MessagePack | Marshal | 2402 | 4104 | 8 |
| Small Struct | JSON | Marshal | 3821 | 2304 | 1 |
| Small Struct | BEVE | Unmarshal | 1179 | 1592 | 4 |
| Small Struct | MessagePack | Unmarshal | 2273 | 1472 | 33 |
| Small Struct | Sonic | Unmarshal | 2359 | 3896 | 6 |
| Small Struct | CBOR | Unmarshal | 4088 | 2240 | 49 |
| Small Struct | JSON | Unmarshal | 5349 | 1288 | 27 |
