# Apple M1 (Virtual) — Darwin

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE ZeroCopy | Marshal | 78392 | 39 | 0 |
| Large Payload | BEVE | Marshal | 148536 | 196648 | 1 |
| Large Payload | CBOR | Marshal | 229320 | 196804 | 1 |
| Large Payload | MessagePack | Marshal | 309589 | 526755 | 115 |
| Large Payload | JSON | Marshal | 526052 | 205139 | 8 |
| Large Payload | Sonic | Marshal | 562946 | 206160 | 3 |
| Large Payload | BEVE | Unmarshal | 261716 | 279729 | 419 |
| Large Payload | Sonic | Unmarshal | 421640 | 340590 | 213 |
| Large Payload | MessagePack | Unmarshal | 691003 | 354965 | 6471 |
| Large Payload | CBOR | Unmarshal | 841434 | 320363 | 6524 |
| Large Payload | JSON | Unmarshal | 2279992 | 506883 | 6580 |
| Medium Payload | BEVE ZeroCopy | Marshal | 8972 | 0 | 0 |
| Medium Payload | BEVE | Marshal | 14545 | 13573 | 1 |
| Medium Payload | CBOR | Marshal | 31486 | 21777 | 1 |
| Medium Payload | MessagePack | Marshal | 46635 | 65778 | 22 |
| Medium Payload | Sonic | Marshal | 59760 | 18668 | 3 |
| Medium Payload | JSON | Marshal | 77285 | 24807 | 8 |
| Medium Payload | BEVE | Unmarshal | 34078 | 31709 | 59 |
| Medium Payload | Sonic | Unmarshal | 43615 | 45611 | 33 |
| Medium Payload | MessagePack | Unmarshal | 67267 | 35614 | 658 |
| Medium Payload | CBOR | Unmarshal | 67865 | 33416 | 685 |
| Medium Payload | JSON | Unmarshal | 219005 | 56008 | 724 |
| Small Struct | BEVE ZeroCopy | Marshal | 1131 | 0 | 0 |
| Small Struct | CBOR | Marshal | 1664 | 1792 | 1 |
| Small Struct | BEVE | Marshal | 1902 | 2305 | 1 |
| Small Struct | MessagePack | Marshal | 1945 | 2056 | 7 |
| Small Struct | JSON | Marshal | 4321 | 2304 | 1 |
| Small Struct | Sonic | Marshal | 9442 | 2732 | 2 |
| Small Struct | BEVE | Unmarshal | 719 | 600 | 4 |
| Small Struct | MessagePack | Unmarshal | 3685 | 1472 | 33 |
| Small Struct | Sonic | Unmarshal | 5086 | 5071 | 6 |
| Small Struct | CBOR | Unmarshal | 10243 | 3496 | 74 |
| Small Struct | JSON | Unmarshal | 12047 | 2088 | 37 |
