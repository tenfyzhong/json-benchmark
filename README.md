# Go JSON Library Benchmark

Performance comparison of popular Go JSON libraries.

## Test Environment

- **Go Version**: go1.25.5
- **OS/Arch**: darwin/arm64
- **CPU Cores**: 8
- **Date**: 2026-01-03

## Libraries Tested

| Library | Description |
|---------|-------------|
| `encoding/json` | Go standard library |
| `goccy/go-json` | High-performance JSON encoder/decoder |
| `json-iterator/go` | Drop-in replacement with better performance |
| `mailru/easyjson` | Code generation based, fastest but requires build step |
| `tidwall/pretty` | JSON formatting/beautifying only (not full encoder/decoder) |

## Benchmark Results

### Encode - Small Struct (~10 fields)

| Rank | Library | ns/op | B/op | allocs/op | Speedup vs Slowest |
|:----:|---------|------:|-----:|----------:|-------------------:|
| 1 | easyjson | 342.60 | 912 | 5 | 3.46x |
| 2 | json-iterator | 383.20 | 1032 | 6 | 3.10x |
| 3 | go-json | 883.10 | 1024 | 5 | 1.34x |
| 4 | encoding/json | 1187.00 | 1024 | 5 | 1.00x |

```
Performance Comparison (lower is better):

easyjson        ███████████ 343 ns/op
json-iterator   ████████████ 383 ns/op
go-json         █████████████████████████████ 883 ns/op
encoding/json   ████████████████████████████████████████ 1187 ns/op
```

### Encode - Large Struct (~50 fields, nested)

| Rank | Library | ns/op | B/op | allocs/op | Speedup vs Slowest |
|:----:|---------|------:|-----:|----------:|-------------------:|
| 1 | easyjson | 3186.00 | 3070 | 8 | 3.88x |
| 2 | json-iterator | 3378.00 | 5277 | 9 | 3.66x |
| 3 | go-json | 7489.00 | 5270 | 8 | 1.65x |
| 4 | encoding/json | 12351.00 | 5273 | 8 | 1.00x |

```
Performance Comparison (lower is better):

easyjson        ██████████ 3186 ns/op
json-iterator   ██████████ 3378 ns/op
go-json         ████████████████████████ 7489 ns/op
encoding/json   ████████████████████████████████████████ 12351 ns/op
```

### Decode - Small Struct

| Rank | Library | ns/op | B/op | allocs/op | Speedup vs Slowest |
|:----:|---------|------:|-----:|----------:|-------------------:|
| 1 | easyjson | 452.20 | 168 | 8 | 3.80x |
| 2 | go-json | 694.60 | 744 | 11 | 2.47x |
| 3 | json-iterator | 1059.00 | 640 | 22 | 1.62x |
| 4 | encoding/json | 1718.00 | 472 | 12 | 1.00x |

```
Performance Comparison (lower is better):

easyjson        ██████████ 452 ns/op
go-json         ████████████████ 695 ns/op
json-iterator   ████████████████████████ 1059 ns/op
encoding/json   ████████████████████████████████████████ 1718 ns/op
```

### Decode - Large Struct

| Rank | Library | ns/op | B/op | allocs/op | Speedup vs Slowest |
|:----:|---------|------:|-----:|----------:|-------------------:|
| 1 | easyjson | 5362.00 | 2944 | 55 | 3.07x |
| 2 | go-json | 7071.00 | 8130 | 58 | 2.32x |
| 3 | json-iterator | 9511.00 | 7162 | 168 | 1.73x |
| 4 | encoding/json | 16435.00 | 3712 | 60 | 1.00x |

```
Performance Comparison (lower is better):

easyjson        █████████████ 5362 ns/op
go-json         █████████████████ 7071 ns/op
json-iterator   ███████████████████████ 9511 ns/op
encoding/json   ████████████████████████████████████████ 16435 ns/op
```

### Pretty Print - Small JSON

| Rank | Library | ns/op | B/op | allocs/op | Speedup vs Slowest |
|:----:|---------|------:|-----:|----------:|-------------------:|
| 1 | pretty.Pretty | 317.10 | 672 | 2 | 2.82x |
| 2 | json.Indent | 894.70 | 0 | 0 | 1.00x |

```
Performance Comparison (lower is better):

pretty.Pretty   ██████████████ 317 ns/op
json.Indent     ████████████████████████████████████████ 895 ns/op
```

### Pretty Print - Large JSON

| Rank | Library | ns/op | B/op | allocs/op | Speedup vs Slowest |
|:----:|---------|------:|-----:|----------:|-------------------:|
| 1 | pretty.Pretty | 3389.00 | 9472 | 3 | 2.92x |
| 2 | json.Indent | 9883.00 | 0 | 0 | 1.00x |

```
Performance Comparison (lower is better):

pretty.Pretty   █████████████ 3389 ns/op
json.Indent     ████████████████████████████████████████ 9883 ns/op
```

### Compact (Ugly) - Small JSON

| Rank | Library | ns/op | B/op | allocs/op | Speedup vs Slowest |
|:----:|---------|------:|-----:|----------:|-------------------:|
| 1 | pretty.Ugly | 219.20 | 288 | 1 | 4.58x |
| 2 | json.Compact | 1004.00 | 0 | 0 | 1.00x |

```
Performance Comparison (lower is better):

pretty.Ugly     ████████ 219 ns/op
json.Compact    ████████████████████████████████████████ 1004 ns/op
```

### Compact (Ugly) - Large JSON

| Rank | Library | ns/op | B/op | allocs/op | Speedup vs Slowest |
|:----:|---------|------:|-----:|----------:|-------------------:|
| 1 | pretty.Ugly | 2673.00 | 3200 | 1 | 4.67x |
| 2 | json.Compact | 12475.00 | 0 | 0 | 1.00x |

```
Performance Comparison (lower is better):

pretty.Ugly     ████████ 2673 ns/op
json.Compact    ████████████████████████████████████████ 12475 ns/op
```

## How to Run

```bash
# Run all benchmarks
make benchmark

# Generate report
make report

# Or manually
go test -bench=. -benchmem ./benchmark/... | tee benchmark_results.txt
go run ./cmd/report/main.go benchmark_results.txt
```

## License

MIT License
