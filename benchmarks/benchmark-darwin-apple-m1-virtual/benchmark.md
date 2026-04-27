# Apple M1 (Virtual) — Darwin

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE ZeroCopy | Marshal | 54151 | 26 | 0 |
| Large Payload | BEVE | Marshal | 84804 | 196650 | 1 |
| Large Payload | CBOR | Marshal | 154086 | 204963 | 1 |
| Large Payload | MessagePack | Marshal | 256376 | 526757 | 115 |
| Large Payload | JSON | Marshal | 367817 | 213332 | 8 |
| Large Payload | Sonic | Marshal | 496031 | 213852 | 3 |
| Large Payload | BEVE | Unmarshal | 173755 | 268400 | 419 |
| Large Payload | Sonic | Unmarshal | 296915 | 348439 | 211 |
| Large Payload | MessagePack | Unmarshal | 418254 | 365240 | 6673 |
| Large Payload | CBOR | Unmarshal | 596348 | 340538 | 6935 |
| Large Payload | JSON | Unmarshal | 1651547 | 488553 | 6406 |
| Medium Payload | BEVE ZeroCopy | Marshal | 5320 | 3 | 0 |
| Medium Payload | BEVE | Marshal | 8222 | 20483 | 1 |
| Medium Payload | CBOR | Marshal | 16854 | 20493 | 1 |
| Medium Payload | MessagePack | Marshal | 26264 | 65778 | 22 |
| Medium Payload | JSON | Marshal | 27668 | 18662 | 8 |
| Medium Payload | Sonic | Marshal | 43606 | 24843 | 3 |
| Medium Payload | BEVE | Unmarshal | 20690 | 33181 | 59 |
| Medium Payload | Sonic | Unmarshal | 26014 | 31882 | 31 |
| Medium Payload | MessagePack | Unmarshal | 43869 | 38190 | 706 |
| Medium Payload | CBOR | Unmarshal | 51526 | 33752 | 693 |
| Medium Payload | JSON | Unmarshal | 167636 | 52128 | 684 |
| Small Struct | BEVE ZeroCopy | Marshal | 586 | 0 | 0 |
| Small Struct | BEVE | Marshal | 606 | 1536 | 1 |
| Small Struct | CBOR | Marshal | 685 | 896 | 1 |
| Small Struct | MessagePack | Marshal | 1000 | 2056 | 7 |
| Small Struct | JSON | Marshal | 1006 | 768 | 1 |
| Small Struct | Sonic | Marshal | 3686 | 2348 | 2 |
| Small Struct | CBOR | Unmarshal | 1001 | 376 | 11 |
| Small Struct | BEVE | Unmarshal | 1013 | 2616 | 4 |
| Small Struct | Sonic | Unmarshal | 1630 | 1736 | 6 |
| Small Struct | MessagePack | Unmarshal | 1807 | 1752 | 39 |
| Small Struct | JSON | Unmarshal | 18486 | 7528 | 101 |
