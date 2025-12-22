# Apple M1 (Virtual) — Darwin

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE ZeroCopy | Marshal | 61944 | 26 | 0 |
| Large Payload | BEVE | Marshal | 112351 | 204850 | 1 |
| Large Payload | CBOR | Marshal | 142622 | 188551 | 1 |
| Large Payload | MessagePack | Marshal | 226082 | 526754 | 115 |
| Large Payload | JSON | Marshal | 352453 | 196893 | 8 |
| Large Payload | Sonic | Marshal | 465862 | 222095 | 3 |
| Large Payload | BEVE | Unmarshal | 190625 | 274772 | 419 |
| Large Payload | Sonic | Unmarshal | 289771 | 347006 | 211 |
| Large Payload | MessagePack | Unmarshal | 445072 | 356150 | 6495 |
| Large Payload | CBOR | Unmarshal | 553590 | 310570 | 6332 |
| Large Payload | JSON | Unmarshal | 1903928 | 570549 | 7465 |
| Medium Payload | BEVE ZeroCopy | Marshal | 5833 | 0 | 0 |
| Medium Payload | BEVE | Marshal | 13071 | 21763 | 1 |
| Medium Payload | CBOR | Marshal | 22878 | 20490 | 1 |
| Medium Payload | JSON | Marshal | 32709 | 18662 | 8 |
| Medium Payload | MessagePack | Marshal | 33768 | 65778 | 22 |
| Medium Payload | Sonic | Marshal | 55147 | 24841 | 3 |
| Medium Payload | BEVE | Unmarshal | 16580 | 26652 | 59 |
| Medium Payload | Sonic | Unmarshal | 26164 | 31295 | 33 |
| Medium Payload | CBOR | Unmarshal | 48478 | 29896 | 619 |
| Medium Payload | MessagePack | Unmarshal | 63909 | 40814 | 764 |
| Medium Payload | JSON | Unmarshal | 191283 | 50104 | 673 |
| Small Struct | BEVE ZeroCopy | Marshal | 590 | 0 | 0 |
| Small Struct | CBOR | Marshal | 644 | 640 | 1 |
| Small Struct | JSON | Marshal | 650 | 384 | 1 |
| Small Struct | BEVE | Marshal | 734 | 1793 | 1 |
| Small Struct | MessagePack | Marshal | 1300 | 2056 | 7 |
| Small Struct | Sonic | Marshal | 3995 | 2085 | 2 |
| Small Struct | CBOR | Unmarshal | 1548 | 712 | 18 |
| Small Struct | BEVE | Unmarshal | 1785 | 3384 | 4 |
| Small Struct | Sonic | Unmarshal | 2624 | 3791 | 6 |
| Small Struct | MessagePack | Unmarshal | 5818 | 4352 | 92 |
| Small Struct | JSON | Unmarshal | 18836 | 7240 | 92 |
