# Golang Benchmark

### Function call (func_test.go)
```
BenchmarkCallClosure-12         50000000                39.3 ns/op
BenchmarkDirect-12              100000000               17.2 ns/op
```

### Struct memory alignment (struct-align_test.go)
```
BenchmarkGoodStructAlign-12     100000000               15.1 ns/op            24 B/op          0 allocs/op
BenchmarkBadStructAlign-12      100000000               18.4 ns/op            32 B/op          0 allocs/op
```

### JSON (json_test.go)
```
BenchmarkJSONEncode_fromStruct-12                5000000               261 ns/op              48 B/op          1 allocs/op
BenchmarkJSONDeDecode_toStruct-12                2000000               866 ns/op             256 B/op          5 allocs/op
BenchmarkJSONEncode_fromMap-12                   2000000               989 ns/op             464 B/op         11 allocs/op
BenchmarkJSONDeDecode_toMap-12                   1000000              1408 ns/op             424 B/op         15 allocs/op
```

### JSON with jsoniter library (jsoniter_test.go)
```
BenchmarkJSONIterEncode_fromStruct-12            5000000               258 ns/op              56 B/op          2 allocs/op
BenchmarkJSONIterDeDecode_toStruct-12            5000000               282 ns/op              80 B/op          3 allocs/op
BenchmarkJSONIterEncode_fromMap-12               3000000               404 ns/op             184 B/op          4 allocs/op
BenchmarkJSONIterDeDecode_toMap-12               3000000               581 ns/op             192 B/op         11 allocs/op
```