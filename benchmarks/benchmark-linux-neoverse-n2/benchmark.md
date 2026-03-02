# Neoverse-N2 — Linux

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE ZeroCopy | Marshal | 67191 | 65 | 0 |
| Large Payload | BEVE | Marshal | 112492 | 196782 | 1 |
| Large Payload | CBOR | Marshal | 185438 | 188637 | 1 |
| Large Payload | MessagePack | Marshal | 282930 | 526803 | 115 |
| Large Payload | Sonic | Marshal | 300185 | 206844 | 3 |
| Large Payload | JSON | Marshal | 384165 | 205124 | 8 |
| Large Payload | BEVE | Unmarshal | 234979 | 270827 | 418 |
| Large Payload | Sonic | Unmarshal | 310655 | 413249 | 209 |
| Large Payload | MessagePack | Unmarshal | 541703 | 366323 | 6717 |
| Large Payload | CBOR | Unmarshal | 649962 | 311003 | 6339 |
| Large Payload | JSON | Unmarshal | 1964664 | 526756 | 6881 |
| Medium Payload | BEVE ZeroCopy | Marshal | 6528 | 6 | 0 |
| Medium Payload | BEVE | Marshal | 9085 | 16389 | 1 |
| Medium Payload | CBOR | Marshal | 16664 | 16395 | 1 |
| Medium Payload | MessagePack | Marshal | 23261 | 33006 | 21 |
| Medium Payload | Sonic | Marshal | 27995 | 20942 | 3 |
| Medium Payload | JSON | Marshal | 43453 | 24814 | 8 |
| Medium Payload | BEVE | Unmarshal | 20606 | 23196 | 58 |
| Medium Payload | Sonic | Unmarshal | 28817 | 37346 | 33 |
| Medium Payload | MessagePack | Unmarshal | 54487 | 37775 | 703 |
| Medium Payload | CBOR | Unmarshal | 72490 | 36584 | 750 |
| Medium Payload | JSON | Unmarshal | 198052 | 57280 | 717 |
| Small Struct | BEVE | Marshal | 337 | 320 | 1 |
| Small Struct | BEVE ZeroCopy | Marshal | 394 | 0 | 0 |
| Small Struct | CBOR | Marshal | 694 | 480 | 1 |
| Small Struct | MessagePack | Marshal | 1526 | 2056 | 7 |
| Small Struct | JSON | Marshal | 2913 | 1792 | 1 |
| Small Struct | Sonic | Marshal | 3578 | 2760 | 2 |
| Small Struct | BEVE | Unmarshal | 907 | 952 | 4 |
| Small Struct | Sonic | Unmarshal | 3224 | 5491 | 6 |
| Small Struct | CBOR | Unmarshal | 3551 | 1904 | 42 |
| Small Struct | MessagePack | Unmarshal | 4589 | 3904 | 82 |
| Small Struct | JSON | Unmarshal | 13873 | 4104 | 65 |
