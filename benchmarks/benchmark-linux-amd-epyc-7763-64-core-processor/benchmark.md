# AMD EPYC 7763 64-Core Processor — Linux

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE ZeroCopy | Marshal | 95859 | 259 | 2 |
| Large Payload | BEVE | Marshal | 130717 | 188845 | 3 |
| Large Payload | Sonic | Marshal | 162877 | 216625 | 4 |
| Large Payload | CBOR | Marshal | 215579 | 205573 | 2 |
| Large Payload | MessagePack | Marshal | 301385 | 526836 | 115 |
| Large Payload | JSON | Marshal | 432415 | 213793 | 9 |
| Large Payload | BEVE | Unmarshal | 235102 | 267680 | 419 |
| Large Payload | Sonic | Unmarshal | 367213 | 555581 | 595 |
| Large Payload | MessagePack | Unmarshal | 587056 | 364460 | 6665 |
| Large Payload | CBOR | Unmarshal | 781213 | 316105 | 6441 |
| Large Payload | JSON | Unmarshal | 2295163 | 522371 | 6942 |
| Medium Payload | BEVE ZeroCopy | Marshal | 11114 | 128 | 2 |
| Medium Payload | Sonic | Marshal | 18766 | 25758 | 4 |
| Medium Payload | BEVE | Marshal | 20109 | 21903 | 3 |
| Medium Payload | CBOR | Marshal | 20709 | 18589 | 2 |
| Medium Payload | MessagePack | Marshal | 35973 | 65838 | 22 |
| Medium Payload | JSON | Marshal | 53016 | 22059 | 9 |
| Medium Payload | BEVE | Unmarshal | 24575 | 28895 | 59 |
| Medium Payload | MessagePack | Unmarshal | 41483 | 22621 | 395 |
| Medium Payload | Sonic | Unmarshal | 41709 | 64272 | 79 |
| Medium Payload | CBOR | Unmarshal | 87893 | 38200 | 785 |
| Medium Payload | JSON | Unmarshal | 214468 | 50624 | 641 |
| Small Struct | BEVE ZeroCopy | Marshal | 626 | 290 | 2 |
| Small Struct | Sonic | Marshal | 1558 | 1995 | 3 |
| Small Struct | CBOR | Marshal | 2016 | 1552 | 2 |
| Small Struct | BEVE | Marshal | 2309 | 2979 | 3 |
| Small Struct | MessagePack | Marshal | 3187 | 4225 | 8 |
| Small Struct | JSON | Marshal | 3589 | 1424 | 2 |
| Small Struct | BEVE | Unmarshal | 2714 | 3384 | 4 |
| Small Struct | Sonic | Unmarshal | 2950 | 3654 | 9 |
| Small Struct | MessagePack | Unmarshal | 3426 | 1792 | 40 |
| Small Struct | CBOR | Unmarshal | 6702 | 2304 | 51 |
| Small Struct | JSON | Unmarshal | 10467 | 2120 | 38 |
