# Neoverse-N2 — Linux

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE ZeroCopy | Marshal | 68233 | 52 | 0 |
| Large Payload | BEVE | Marshal | 111781 | 196742 | 1 |
| Large Payload | CBOR | Marshal | 195247 | 196935 | 1 |
| Large Payload | MessagePack | Marshal | 275090 | 526803 | 115 |
| Large Payload | Sonic | Marshal | 280483 | 198261 | 3 |
| Large Payload | JSON | Marshal | 385391 | 213290 | 8 |
| Large Payload | BEVE | Unmarshal | 227061 | 270667 | 418 |
| Large Payload | Sonic | Unmarshal | 298194 | 400355 | 211 |
| Large Payload | MessagePack | Unmarshal | 528181 | 351293 | 6398 |
| Large Payload | CBOR | Unmarshal | 628737 | 293994 | 5996 |
| Large Payload | JSON | Unmarshal | 1918037 | 512340 | 6743 |
| Medium Payload | BEVE ZeroCopy | Marshal | 7355 | 6 | 0 |
| Medium Payload | BEVE | Marshal | 9246 | 16389 | 1 |
| Medium Payload | CBOR | Marshal | 17486 | 18447 | 1 |
| Medium Payload | MessagePack | Marshal | 33093 | 65782 | 22 |
| Medium Payload | Sonic | Marshal | 35180 | 27889 | 3 |
| Medium Payload | JSON | Marshal | 39561 | 21988 | 8 |
| Medium Payload | BEVE | Unmarshal | 21477 | 26109 | 59 |
| Medium Payload | Sonic | Unmarshal | 25263 | 30805 | 33 |
| Medium Payload | MessagePack | Unmarshal | 48294 | 31886 | 580 |
| Medium Payload | CBOR | Unmarshal | 67342 | 33384 | 684 |
| Medium Payload | JSON | Unmarshal | 148518 | 39032 | 505 |
| Small Struct | BEVE ZeroCopy | Marshal | 394 | 0 | 0 |
| Small Struct | CBOR | Marshal | 791 | 576 | 1 |
| Small Struct | BEVE | Marshal | 842 | 1280 | 1 |
| Small Struct | Sonic | Marshal | 3200 | 2356 | 2 |
| Small Struct | JSON | Marshal | 3460 | 2048 | 1 |
| Small Struct | MessagePack | Marshal | 4012 | 8201 | 9 |
| Small Struct | Sonic | Unmarshal | 1203 | 1170 | 6 |
| Small Struct | BEVE | Unmarshal | 1264 | 1080 | 4 |
| Small Struct | CBOR | Unmarshal | 3742 | 2024 | 44 |
| Small Struct | MessagePack | Unmarshal | 5582 | 4832 | 103 |
| Small Struct | JSON | Unmarshal | 25756 | 8040 | 117 |
