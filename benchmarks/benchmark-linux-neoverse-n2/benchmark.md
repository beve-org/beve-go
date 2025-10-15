# Neoverse-N2 — Linux

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE ZeroCopy | Marshal | 75393 | 233 | 2 |
| Large Payload | BEVE | Marshal | 123766 | 189644 | 3 |
| Large Payload | CBOR | Marshal | 183106 | 200286 | 3 |
| Large Payload | MessagePack | Marshal | 288945 | 526866 | 115 |
| Large Payload | Sonic | Marshal | 316839 | 231427 | 4 |
| Large Payload | JSON | Marshal | 368410 | 206186 | 9 |
| Large Payload | BEVE | Unmarshal | 218288 | 265576 | 417 |
| Large Payload | Sonic | Unmarshal | 278891 | 384466 | 213 |
| Large Payload | MessagePack | Unmarshal | 484506 | 328727 | 5955 |
| Large Payload | CBOR | Unmarshal | 658741 | 321177 | 6551 |
| Large Payload | JSON | Unmarshal | 2040182 | 559292 | 7256 |
| Medium Payload | BEVE ZeroCopy | Marshal | 8928 | 141 | 2 |
| Medium Payload | BEVE | Marshal | 10040 | 16536 | 3 |
| Medium Payload | CBOR | Marshal | 19328 | 20589 | 2 |
| Medium Payload | Sonic | Marshal | 28311 | 20953 | 4 |
| Medium Payload | MessagePack | Marshal | 28969 | 65838 | 22 |
| Medium Payload | JSON | Marshal | 42095 | 24901 | 9 |
| Medium Payload | BEVE | Unmarshal | 22225 | 28638 | 59 |
| Medium Payload | Sonic | Unmarshal | 28778 | 39586 | 33 |
| Medium Payload | MessagePack | Unmarshal | 58275 | 43120 | 813 |
| Medium Payload | CBOR | Unmarshal | 73127 | 36648 | 755 |
| Medium Payload | JSON | Unmarshal | 212825 | 59896 | 794 |
| Small Struct | BEVE ZeroCopy | Marshal | 497 | 288 | 2 |
| Small Struct | Sonic | Marshal | 1294 | 906 | 3 |
| Small Struct | BEVE | Marshal | 1402 | 2337 | 3 |
| Small Struct | CBOR | Marshal | 1427 | 1553 | 2 |
| Small Struct | MessagePack | Marshal | 2380 | 4224 | 8 |
| Small Struct | JSON | Marshal | 3808 | 2449 | 2 |
| Small Struct | BEVE | Unmarshal | 1735 | 3384 | 4 |
| Small Struct | Sonic | Unmarshal | 2009 | 2604 | 6 |
| Small Struct | MessagePack | Unmarshal | 2059 | 1216 | 28 |
| Small Struct | CBOR | Unmarshal | 2809 | 1280 | 30 |
| Small Struct | JSON | Unmarshal | 10395 | 2472 | 49 |
