# Unknown CPU — Windows

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE ZeroCopy | Marshal | 69037 | 65 | 0 |
| Large Payload | BEVE | Marshal | 140406 | 196706 | 1 |
| Large Payload | Sonic | Marshal | 201717 | 217668 | 3 |
| Large Payload | CBOR | Marshal | 231675 | 188608 | 1 |
| Large Payload | MessagePack | Marshal | 365232 | 526743 | 115 |
| Large Payload | JSON | Marshal | 456491 | 197048 | 8 |
| Large Payload | BEVE | Unmarshal | 329839 | 265332 | 419 |
| Large Payload | Sonic | Unmarshal | 530226 | 549681 | 574 |
| Large Payload | MessagePack | Unmarshal | 749556 | 342218 | 6214 |
| Large Payload | CBOR | Unmarshal | 857038 | 325353 | 6623 |
| Large Payload | JSON | Unmarshal | 2384405 | 528922 | 7039 |
| Medium Payload | BEVE ZeroCopy | Marshal | 10765 | 5 | 0 |
| Medium Payload | BEVE | Marshal | 14151 | 18445 | 1 |
| Medium Payload | Sonic | Marshal | 22445 | 19572 | 3 |
| Medium Payload | CBOR | Marshal | 22738 | 16395 | 1 |
| Medium Payload | MessagePack | Marshal | 48572 | 65777 | 22 |
| Medium Payload | JSON | Marshal | 54486 | 21994 | 8 |
| Medium Payload | BEVE | Unmarshal | 33853 | 32862 | 59 |
| Medium Payload | Sonic | Unmarshal | 66997 | 63495 | 76 |
| Medium Payload | MessagePack | Unmarshal | 72508 | 37887 | 707 |
| Medium Payload | CBOR | Unmarshal | 93631 | 32471 | 664 |
| Medium Payload | JSON | Unmarshal | 219338 | 42224 | 587 |
| Small Struct | Sonic | Marshal | 831 | 613 | 2 |
| Small Struct | BEVE ZeroCopy | Marshal | 1383 | 0 | 0 |
| Small Struct | CBOR | Marshal | 1733 | 1152 | 1 |
| Small Struct | BEVE | Marshal | 2234 | 2305 | 1 |
| Small Struct | MessagePack | Marshal | 2513 | 2056 | 7 |
| Small Struct | JSON | Marshal | 5922 | 2305 | 1 |
| Small Struct | BEVE | Unmarshal | 2932 | 3000 | 4 |
| Small Struct | Sonic | Unmarshal | 3348 | 3648 | 9 |
| Small Struct | MessagePack | Unmarshal | 6553 | 4224 | 88 |
| Small Struct | CBOR | Unmarshal | 6695 | 3240 | 70 |
| Small Struct | JSON | Unmarshal | 9465 | 2216 | 41 |
