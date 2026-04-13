# Apple M1 (Virtual) — Darwin

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE ZeroCopy | Marshal | 59257 | 39 | 0 |
| Large Payload | BEVE | Marshal | 106654 | 196649 | 1 |
| Large Payload | CBOR | Marshal | 176548 | 196786 | 1 |
| Large Payload | MessagePack | Marshal | 333786 | 526753 | 115 |
| Large Payload | JSON | Marshal | 348024 | 196946 | 8 |
| Large Payload | Sonic | Marshal | 435189 | 213944 | 3 |
| Large Payload | BEVE | Unmarshal | 265919 | 264815 | 419 |
| Large Payload | Sonic | Unmarshal | 338890 | 347343 | 211 |
| Large Payload | MessagePack | Unmarshal | 476257 | 354006 | 6466 |
| Large Payload | CBOR | Unmarshal | 626491 | 311530 | 6360 |
| Large Payload | JSON | Unmarshal | 2068569 | 535882 | 7000 |
| Medium Payload | BEVE ZeroCopy | Marshal | 6375 | 5 | 0 |
| Medium Payload | BEVE | Marshal | 10484 | 24585 | 1 |
| Medium Payload | CBOR | Marshal | 14870 | 19091 | 1 |
| Medium Payload | MessagePack | Marshal | 18031 | 33005 | 21 |
| Medium Payload | JSON | Marshal | 29194 | 20710 | 8 |
| Medium Payload | Sonic | Marshal | 43275 | 24815 | 3 |
| Medium Payload | BEVE | Unmarshal | 18482 | 25884 | 59 |
| Medium Payload | Sonic | Unmarshal | 32184 | 42055 | 33 |
| Medium Payload | MessagePack | Unmarshal | 40690 | 36389 | 674 |
| Medium Payload | CBOR | Unmarshal | 49158 | 30487 | 628 |
| Medium Payload | JSON | Unmarshal | 159138 | 50360 | 667 |
| Small Struct | BEVE | Marshal | 264 | 448 | 1 |
| Small Struct | BEVE ZeroCopy | Marshal | 494 | 0 | 0 |
| Small Struct | JSON | Marshal | 796 | 512 | 1 |
| Small Struct | CBOR | Marshal | 1367 | 2048 | 1 |
| Small Struct | MessagePack | Marshal | 1475 | 4104 | 8 |
| Small Struct | Sonic | Marshal | 1497 | 798 | 2 |
| Small Struct | BEVE | Unmarshal | 1236 | 3384 | 4 |
| Small Struct | CBOR | Unmarshal | 2257 | 1576 | 36 |
| Small Struct | Sonic | Unmarshal | 2824 | 4990 | 6 |
| Small Struct | MessagePack | Unmarshal | 3374 | 4416 | 94 |
| Small Struct | JSON | Unmarshal | 12944 | 4200 | 68 |
