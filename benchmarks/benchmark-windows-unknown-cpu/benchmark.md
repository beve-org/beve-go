# Unknown CPU — Windows

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE ZeroCopy | Marshal | 79050 | 233 | 2 |
| Large Payload | BEVE | Marshal | 116486 | 189109 | 3 |
| Large Payload | Sonic | Marshal | 162119 | 228300 | 4 |
| Large Payload | CBOR | Marshal | 225006 | 190054 | 2 |
| Large Payload | MessagePack | Marshal | 288051 | 526765 | 115 |
| Large Payload | JSON | Marshal | 487549 | 215774 | 9 |
| Large Payload | BEVE | Unmarshal | 271590 | 270982 | 419 |
| Large Payload | Sonic | Unmarshal | 435012 | 553582 | 588 |
| Large Payload | MessagePack | Unmarshal | 640117 | 324929 | 5861 |
| Large Payload | CBOR | Unmarshal | 874978 | 333147 | 6784 |
| Large Payload | JSON | Unmarshal | 2637990 | 566581 | 7312 |
| Medium Payload | BEVE ZeroCopy | Marshal | 7545 | 138 | 2 |
| Medium Payload | BEVE | Marshal | 13798 | 19226 | 3 |
| Medium Payload | Sonic | Marshal | 17822 | 22092 | 4 |
| Medium Payload | CBOR | Marshal | 23242 | 19175 | 2 |
| Medium Payload | MessagePack | Marshal | 37577 | 65828 | 22 |
| Medium Payload | JSON | Marshal | 42594 | 18775 | 9 |
| Medium Payload | BEVE | Unmarshal | 29446 | 29787 | 59 |
| Medium Payload | Sonic | Unmarshal | 51538 | 63814 | 78 |
| Medium Payload | MessagePack | Unmarshal | 75545 | 40189 | 749 |
| Medium Payload | CBOR | Unmarshal | 80602 | 30728 | 636 |
| Medium Payload | JSON | Unmarshal | 242066 | 51616 | 668 |
| Small Struct | BEVE ZeroCopy | Marshal | 568 | 289 | 2 |
| Small Struct | Sonic | Marshal | 602 | 522 | 3 |
| Small Struct | BEVE | Marshal | 928 | 800 | 3 |
| Small Struct | MessagePack | Marshal | 1287 | 1152 | 6 |
| Small Struct | JSON | Marshal | 1443 | 624 | 2 |
| Small Struct | CBOR | Marshal | 2602 | 2450 | 2 |
| Small Struct | BEVE | Unmarshal | 3739 | 3384 | 4 |
| Small Struct | Sonic | Unmarshal | 4786 | 7409 | 10 |
| Small Struct | CBOR | Unmarshal | 6533 | 3048 | 64 |
| Small Struct | MessagePack | Unmarshal | 6950 | 3873 | 81 |
| Small Struct | JSON | Unmarshal | 32209 | 7968 | 115 |
