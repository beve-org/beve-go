# Neoverse-N2 — Linux

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE ZeroCopy | Marshal | 71630 | 259 | 2 |
| Large Payload | BEVE | Marshal | 105847 | 173813 | 3 |
| Large Payload | CBOR | Marshal | 197458 | 207277 | 3 |
| Large Payload | MessagePack | Marshal | 278005 | 526858 | 115 |
| Large Payload | Sonic | Marshal | 311542 | 227204 | 4 |
| Large Payload | JSON | Marshal | 390182 | 225358 | 9 |
| Large Payload | BEVE | Unmarshal | 232116 | 282640 | 417 |
| Large Payload | Sonic | Unmarshal | 276754 | 365251 | 211 |
| Large Payload | MessagePack | Unmarshal | 518273 | 347034 | 6310 |
| Large Payload | CBOR | Unmarshal | 659259 | 320555 | 6530 |
| Large Payload | JSON | Unmarshal | 2020509 | 545477 | 7129 |
| Medium Payload | BEVE ZeroCopy | Marshal | 7922 | 128 | 2 |
| Medium Payload | BEVE | Marshal | 10529 | 18581 | 3 |
| Medium Payload | CBOR | Marshal | 21929 | 24675 | 2 |
| Medium Payload | MessagePack | Marshal | 30628 | 65839 | 22 |
| Medium Payload | Sonic | Marshal | 31401 | 25087 | 4 |
| Medium Payload | JSON | Marshal | 35167 | 19383 | 9 |
| Medium Payload | BEVE | Unmarshal | 23826 | 30622 | 59 |
| Medium Payload | Sonic | Unmarshal | 26741 | 34351 | 29 |
| Medium Payload | MessagePack | Unmarshal | 46505 | 30270 | 551 |
| Medium Payload | CBOR | Unmarshal | 62646 | 30432 | 623 |
| Medium Payload | JSON | Unmarshal | 184349 | 51288 | 660 |
| Small Struct | BEVE ZeroCopy | Marshal | 438 | 289 | 2 |
| Small Struct | Sonic | Marshal | 824 | 572 | 3 |
| Small Struct | CBOR | Marshal | 1175 | 1168 | 2 |
| Small Struct | JSON | Marshal | 1467 | 848 | 2 |
| Small Struct | BEVE | Marshal | 1598 | 2980 | 3 |
| Small Struct | MessagePack | Marshal | 1621 | 2176 | 7 |
| Small Struct | BEVE | Unmarshal | 669 | 376 | 4 |
| Small Struct | Sonic | Unmarshal | 2948 | 4769 | 6 |
| Small Struct | CBOR | Unmarshal | 4262 | 2312 | 51 |
| Small Struct | MessagePack | Unmarshal | 4443 | 3520 | 74 |
| Small Struct | JSON | Unmarshal | 20594 | 7272 | 93 |
