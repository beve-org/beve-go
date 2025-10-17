# Apple M1 (Virtual) — Darwin

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE ZeroCopy | Marshal | 65787 | 26 | 0 |
| Large Payload | BEVE | Marshal | 96918 | 188470 | 1 |
| Large Payload | CBOR | Marshal | 165546 | 196750 | 1 |
| Large Payload | MessagePack | Marshal | 303319 | 526751 | 115 |
| Large Payload | JSON | Marshal | 415268 | 221523 | 8 |
| Large Payload | Sonic | Marshal | 484381 | 206073 | 3 |
| Large Payload | BEVE | Unmarshal | 274564 | 270449 | 418 |
| Large Payload | Sonic | Unmarshal | 414229 | 352411 | 211 |
| Large Payload | MessagePack | Unmarshal | 425607 | 330674 | 5993 |
| Large Payload | CBOR | Unmarshal | 803928 | 317147 | 6469 |
| Large Payload | JSON | Unmarshal | 2174389 | 519443 | 6822 |
| Medium Payload | BEVE ZeroCopy | Marshal | 7644 | 3 | 0 |
| Medium Payload | BEVE | Marshal | 14374 | 27269 | 1 |
| Medium Payload | CBOR | Marshal | 17763 | 20490 | 1 |
| Medium Payload | MessagePack | Marshal | 31875 | 33004 | 21 |
| Medium Payload | JSON | Marshal | 43020 | 21993 | 8 |
| Medium Payload | Sonic | Marshal | 46983 | 18714 | 3 |
| Medium Payload | BEVE | Unmarshal | 22996 | 27357 | 59 |
| Medium Payload | Sonic | Unmarshal | 27544 | 32379 | 31 |
| Medium Payload | MessagePack | Unmarshal | 43115 | 34525 | 637 |
| Medium Payload | CBOR | Unmarshal | 44542 | 27768 | 574 |
| Medium Payload | JSON | Unmarshal | 233174 | 54639 | 704 |
| Small Struct | BEVE ZeroCopy | Marshal | 347 | 0 | 0 |
| Small Struct | BEVE | Marshal | 1069 | 1152 | 1 |
| Small Struct | MessagePack | Marshal | 1135 | 2056 | 7 |
| Small Struct | CBOR | Marshal | 1170 | 1280 | 1 |
| Small Struct | JSON | Marshal | 3471 | 1280 | 1 |
| Small Struct | Sonic | Marshal | 4845 | 2327 | 2 |
| Small Struct | MessagePack | Unmarshal | 1784 | 640 | 16 |
| Small Struct | BEVE | Unmarshal | 1875 | 3000 | 4 |
| Small Struct | Sonic | Unmarshal | 2683 | 2691 | 6 |
| Small Struct | CBOR | Unmarshal | 5305 | 2408 | 52 |
| Small Struct | JSON | Unmarshal | 19373 | 4496 | 77 |
