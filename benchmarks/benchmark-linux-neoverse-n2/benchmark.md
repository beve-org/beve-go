# Neoverse-N2 — Linux

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Large Payload | BEVE ZeroCopy | Marshal | 65457 | 52 | 0 |
| Large Payload | BEVE | Marshal | 113298 | 204935 | 1 |
| Large Payload | CBOR | Marshal | 176415 | 180467 | 1 |
| Large Payload | MessagePack | Marshal | 282052 | 526807 | 115 |
| Large Payload | Sonic | Marshal | 313012 | 222793 | 3 |
| Large Payload | JSON | Marshal | 409809 | 229757 | 8 |
| Large Payload | BEVE | Unmarshal | 230552 | 275373 | 419 |
| Large Payload | Sonic | Unmarshal | 285967 | 382290 | 209 |
| Large Payload | MessagePack | Unmarshal | 499169 | 335045 | 6069 |
| Large Payload | CBOR | Unmarshal | 678444 | 332651 | 6781 |
| Large Payload | JSON | Unmarshal | 1863988 | 494467 | 6494 |
| Medium Payload | BEVE ZeroCopy | Marshal | 8009 | 6 | 0 |
| Medium Payload | BEVE | Marshal | 9365 | 18438 | 1 |
| Medium Payload | CBOR | Marshal | 18143 | 19090 | 1 |
| Medium Payload | MessagePack | Marshal | 22528 | 33006 | 21 |
| Medium Payload | Sonic | Marshal | 33906 | 25111 | 3 |
| Medium Payload | JSON | Marshal | 43584 | 24807 | 8 |
| Medium Payload | BEVE | Unmarshal | 23849 | 30942 | 59 |
| Medium Payload | Sonic | Unmarshal | 32167 | 45109 | 33 |
| Medium Payload | MessagePack | Unmarshal | 53358 | 37920 | 706 |
| Medium Payload | CBOR | Unmarshal | 59798 | 28664 | 588 |
| Medium Payload | JSON | Unmarshal | 199964 | 58440 | 732 |
| Small Struct | BEVE ZeroCopy | Marshal | 710 | 0 | 0 |
| Small Struct | JSON | Marshal | 734 | 288 | 1 |
| Small Struct | MessagePack | Marshal | 765 | 520 | 5 |
| Small Struct | BEVE | Marshal | 1345 | 2689 | 1 |
| Small Struct | Sonic | Marshal | 1737 | 1196 | 2 |
| Small Struct | CBOR | Marshal | 1899 | 2304 | 1 |
| Small Struct | BEVE | Unmarshal | 1656 | 3384 | 4 |
| Small Struct | Sonic | Unmarshal | 1888 | 2557 | 6 |
| Small Struct | MessagePack | Unmarshal | 3920 | 3168 | 67 |
| Small Struct | CBOR | Unmarshal | 4365 | 2440 | 53 |
| Small Struct | JSON | Unmarshal | 22598 | 7592 | 103 |
