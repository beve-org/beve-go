# Intel(R) Xeon(R) Platinum 8370C CPU @ 2.80GHz — Linux

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE ZeroCopy | Marshal | 72004 | 26 | 0 |
| Large Payload | BEVE | Marshal | 110375 | 196679 | 1 |
| Large Payload | Sonic | Marshal | 162385 | 216152 | 3 |
| Large Payload | CBOR | Marshal | 193305 | 188643 | 1 |
| Large Payload | MessagePack | Marshal | 290323 | 526773 | 115 |
| Large Payload | JSON | Marshal | 435697 | 221503 | 8 |
| Large Payload | BEVE | Unmarshal | 243006 | 282178 | 419 |
| Large Payload | Sonic | Unmarshal | 398027 | 589914 | 607 |
| Large Payload | MessagePack | Unmarshal | 538392 | 335927 | 6079 |
| Large Payload | CBOR | Unmarshal | 708761 | 314634 | 6413 |
| Large Payload | JSON | Unmarshal | 2161667 | 564309 | 7381 |
| Medium Payload | BEVE ZeroCopy | Marshal | 7655 | 3 | 0 |
| Medium Payload | BEVE | Marshal | 10922 | 16394 | 1 |
| Medium Payload | Sonic | Marshal | 13193 | 16862 | 3 |
| Medium Payload | CBOR | Marshal | 21983 | 21780 | 1 |
| Medium Payload | MessagePack | Marshal | 38829 | 65782 | 22 |
| Medium Payload | JSON | Marshal | 41367 | 20718 | 8 |
| Medium Payload | BEVE | Unmarshal | 23454 | 26942 | 59 |
| Medium Payload | Sonic | Unmarshal | 33710 | 45803 | 67 |
| Medium Payload | MessagePack | Unmarshal | 59418 | 40561 | 763 |
| Medium Payload | CBOR | Unmarshal | 69178 | 33016 | 681 |
| Medium Payload | JSON | Unmarshal | 192507 | 49432 | 673 |
| Small Struct | BEVE ZeroCopy | Marshal | 300 | 0 | 0 |
| Small Struct | BEVE | Marshal | 984 | 1536 | 1 |
| Small Struct | Sonic | Marshal | 1622 | 2100 | 2 |
| Small Struct | CBOR | Marshal | 1985 | 2048 | 1 |
| Small Struct | MessagePack | Marshal | 2405 | 4104 | 8 |
| Small Struct | JSON | Marshal | 4949 | 2689 | 1 |
| Small Struct | BEVE | Unmarshal | 2228 | 3384 | 4 |
| Small Struct | Sonic | Unmarshal | 3213 | 4673 | 9 |
| Small Struct | MessagePack | Unmarshal | 5013 | 3904 | 82 |
| Small Struct | CBOR | Unmarshal | 7231 | 4192 | 88 |
| Small Struct | JSON | Unmarshal | 17655 | 4680 | 83 |
