# Apple M1 (Virtual) — Darwin

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE ZeroCopy | Marshal | 56525 | 39 | 0 |
| Large Payload | BEVE | Marshal | 79248 | 180290 | 1 |
| Large Payload | CBOR | Marshal | 140355 | 188550 | 1 |
| Large Payload | MessagePack | Marshal | 166955 | 526751 | 115 |
| Large Payload | JSON | Marshal | 318281 | 221474 | 8 |
| Large Payload | Sonic | Marshal | 385085 | 213355 | 3 |
| Large Payload | BEVE | Unmarshal | 164738 | 275411 | 419 |
| Large Payload | Sonic | Unmarshal | 275202 | 368903 | 211 |
| Large Payload | MessagePack | Unmarshal | 365028 | 342148 | 6203 |
| Large Payload | CBOR | Unmarshal | 504697 | 334185 | 6807 |
| Large Payload | JSON | Unmarshal | 1744158 | 552274 | 7233 |
| Medium Payload | BEVE ZeroCopy | Marshal | 4820 | 0 | 0 |
| Medium Payload | BEVE | Marshal | 7688 | 18435 | 1 |
| Medium Payload | CBOR | Marshal | 12917 | 16393 | 1 |
| Medium Payload | MessagePack | Marshal | 21843 | 65778 | 22 |
| Medium Payload | JSON | Marshal | 29563 | 20710 | 8 |
| Medium Payload | Sonic | Marshal | 48432 | 27437 | 3 |
| Medium Payload | BEVE | Unmarshal | 15860 | 27196 | 59 |
| Medium Payload | Sonic | Unmarshal | 28028 | 37566 | 33 |
| Medium Payload | MessagePack | Unmarshal | 35254 | 34333 | 633 |
| Medium Payload | CBOR | Unmarshal | 60237 | 36120 | 744 |
| Medium Payload | JSON | Unmarshal | 129101 | 39448 | 514 |
| Small Struct | BEVE ZeroCopy | Marshal | 491 | 0 | 0 |
| Small Struct | BEVE | Marshal | 836 | 2304 | 1 |
| Small Struct | JSON | Marshal | 916 | 704 | 1 |
| Small Struct | CBOR | Marshal | 1434 | 2304 | 1 |
| Small Struct | MessagePack | Marshal | 1620 | 4104 | 8 |
| Small Struct | Sonic | Marshal | 2389 | 1431 | 2 |
| Small Struct | BEVE | Unmarshal | 1233 | 3384 | 4 |
| Small Struct | Sonic | Unmarshal | 2160 | 3372 | 6 |
| Small Struct | MessagePack | Unmarshal | 2913 | 3296 | 71 |
| Small Struct | CBOR | Unmarshal | 5345 | 4584 | 96 |
| Small Struct | JSON | Unmarshal | 9066 | 3656 | 51 |
