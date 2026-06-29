# Apple M1 (Virtual) — Darwin

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE ZeroCopy | Marshal | 53977 | 26 | 0 |
| Large Payload | BEVE | Marshal | 75472 | 188457 | 1 |
| Large Payload | CBOR | Marshal | 138955 | 196741 | 1 |
| Large Payload | MessagePack | Marshal | 186695 | 526751 | 115 |
| Large Payload | JSON | Marshal | 322556 | 213360 | 8 |
| Large Payload | Sonic | Marshal | 379025 | 205151 | 3 |
| Large Payload | BEVE | Unmarshal | 169891 | 269522 | 418 |
| Large Payload | Sonic | Unmarshal | 261044 | 317722 | 207 |
| Large Payload | MessagePack | Unmarshal | 385767 | 368536 | 6750 |
| Large Payload | CBOR | Unmarshal | 491178 | 308521 | 6282 |
| Large Payload | JSON | Unmarshal | 1777687 | 539692 | 7226 |
| Medium Payload | BEVE ZeroCopy | Marshal | 7177 | 3 | 0 |
| Medium Payload | BEVE | Marshal | 7408 | 18436 | 1 |
| Medium Payload | CBOR | Marshal | 12899 | 18448 | 1 |
| Medium Payload | MessagePack | Marshal | 22249 | 65778 | 22 |
| Medium Payload | JSON | Marshal | 30848 | 20710 | 8 |
| Medium Payload | Sonic | Marshal | 41779 | 24779 | 3 |
| Medium Payload | BEVE | Unmarshal | 18522 | 28796 | 59 |
| Medium Payload | Sonic | Unmarshal | 26595 | 37303 | 33 |
| Medium Payload | MessagePack | Unmarshal | 38830 | 39822 | 744 |
| Medium Payload | CBOR | Unmarshal | 47449 | 32248 | 662 |
| Medium Payload | JSON | Unmarshal | 176523 | 48280 | 662 |
| Small Struct | CBOR | Marshal | 316 | 256 | 1 |
| Small Struct | BEVE ZeroCopy | Marshal | 419 | 0 | 0 |
| Small Struct | BEVE | Marshal | 652 | 1536 | 1 |
| Small Struct | Sonic | Marshal | 2208 | 1168 | 2 |
| Small Struct | MessagePack | Marshal | 2356 | 8201 | 9 |
| Small Struct | JSON | Marshal | 3343 | 2688 | 1 |
| Small Struct | BEVE | Unmarshal | 877 | 1848 | 4 |
| Small Struct | CBOR | Unmarshal | 1265 | 616 | 16 |
| Small Struct | Sonic | Unmarshal | 3125 | 5863 | 6 |
| Small Struct | MessagePack | Unmarshal | 3869 | 4728 | 100 |
| Small Struct | JSON | Unmarshal | 9133 | 3656 | 51 |
