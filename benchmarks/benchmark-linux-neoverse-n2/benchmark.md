# Neoverse-N2 — Linux

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE ZeroCopy | Marshal | 67135 | 65 | 0 |
| Large Payload | BEVE | Marshal | 114641 | 188643 | 1 |
| Large Payload | CBOR | Marshal | 184325 | 188741 | 1 |
| Large Payload | MessagePack | Marshal | 285805 | 526809 | 115 |
| Large Payload | Sonic | Marshal | 306035 | 217676 | 3 |
| Large Payload | JSON | Marshal | 427590 | 229731 | 8 |
| Large Payload | BEVE | Unmarshal | 224758 | 268456 | 419 |
| Large Payload | Sonic | Unmarshal | 295978 | 404514 | 211 |
| Large Payload | MessagePack | Unmarshal | 515216 | 352095 | 6411 |
| Large Payload | CBOR | Unmarshal | 647787 | 313882 | 6398 |
| Large Payload | JSON | Unmarshal | 1940162 | 516284 | 6821 |
| Medium Payload | BEVE ZeroCopy | Marshal | 7934 | 1 | 0 |
| Medium Payload | BEVE | Marshal | 9346 | 18442 | 1 |
| Medium Payload | CBOR | Marshal | 19790 | 20502 | 1 |
| Medium Payload | MessagePack | Marshal | 21384 | 33007 | 21 |
| Medium Payload | Sonic | Marshal | 32389 | 25085 | 3 |
| Medium Payload | JSON | Marshal | 48315 | 27496 | 8 |
| Medium Payload | BEVE | Unmarshal | 22641 | 28638 | 59 |
| Medium Payload | Sonic | Unmarshal | 32184 | 45431 | 33 |
| Medium Payload | MessagePack | Unmarshal | 46766 | 30014 | 545 |
| Medium Payload | CBOR | Unmarshal | 68490 | 33736 | 696 |
| Medium Payload | JSON | Unmarshal | 196157 | 53304 | 724 |
| Small Struct | BEVE ZeroCopy | Marshal | 674 | 0 | 0 |
| Small Struct | BEVE | Marshal | 921 | 1792 | 1 |
| Small Struct | Sonic | Marshal | 1909 | 1331 | 2 |
| Small Struct | CBOR | Marshal | 1942 | 2304 | 1 |
| Small Struct | MessagePack | Marshal | 2268 | 4104 | 8 |
| Small Struct | JSON | Marshal | 2693 | 1536 | 1 |
| Small Struct | BEVE | Unmarshal | 1472 | 2616 | 4 |
| Small Struct | MessagePack | Unmarshal | 3320 | 2496 | 54 |
| Small Struct | Sonic | Unmarshal | 3434 | 5433 | 6 |
| Small Struct | CBOR | Unmarshal | 3671 | 1928 | 43 |
| Small Struct | JSON | Unmarshal | 17446 | 4616 | 81 |
