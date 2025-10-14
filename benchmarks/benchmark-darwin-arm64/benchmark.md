# Apple M2 Max — Darwin

![Benchmark Chart](benchmark-darwin-arm64/benchmark.png)

_Performance visualization: lower is better. Chart shows relative performance across different codecs and scenarios._

### Detailed Results

| Scenario | Codec | Operation | ns/op | B/op | allocs/op |
|----------|-------|-----------|-------|------|-----------|
| Small Struct | BEVE ZeroCopy | Marshal | 530 | 290 | 2 |
| Small Struct | BEVE | Marshal | 1012 | 2596 | 3 |
| Small Struct | CBOR | Marshal | 1043 | 1681 | 2 |
| Small Struct | JSON | Marshal | 1434 | 1168 | 2 |
| Small Struct | MessagePack | Marshal | 2500 | 8325 | 9 |
| Small Struct | Sonic | Marshal | 3845 | 2525 | 3 |
| Small Struct | BEVE | Unmarshal | 719 | 1593 | 4 |
| Small Struct | MessagePack | Unmarshal | 1593 | 1697 | 38 |
| Small Struct | Sonic | Unmarshal | 1950 | 3793 | 6 |
| Small Struct | CBOR | Unmarshal | 3161 | 2800 | 60 |
| Small Struct | JSON | Unmarshal | 7092 | 2408 | 47 |