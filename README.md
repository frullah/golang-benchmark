# Golang Benchmark

## Results :
Machine specification :
```
OS:  Arch Linux x86_64 
CPU: Intel i7-8750H (12 cores) @ 4.100GHz 
Mem: 16GB
Go:  go1.12.5 linux/amd64
```

### Function call [[func_test.go](benchmark/func_test.go)]
```
BenchmarkFunction_closureCall-12                2000000000               1.26 ns/op
BenchmarkFunction_directCall-12                 2000000000               0.26 ns/op
BenchmarkFunction_variadicParameter-12          2000000000               0.51 ns/op
```

### Struct memory alignment [[struct-align_test.go](benchmark/struct-align_test.go)]
```
BenchmarkGoodStructAlign-12     100000000               15.1 ns/op            24 B/op          0 allocs/op
BenchmarkBadStructAlign-12      100000000               18.4 ns/op            32 B/op          0 allocs/op
```

### JSON [[func_test.go](benchmark/json_test.go)]
```
BenchmarkJSONEncode_fromStruct-12                5000000               261 ns/op              48 B/op          1 allocs/op
BenchmarkJSONDeDecode_toStruct-12                2000000               866 ns/op             256 B/op          5 allocs/op
BenchmarkJSONEncode_fromMap-12                   2000000               989 ns/op             464 B/op         11 allocs/op
BenchmarkJSONDeDecode_toMap-12                   1000000              1408 ns/op             424 B/op         15 allocs/op
```

### JSON with jsoniter library [[func_test.go](benchmark/jsoniter_test.go)]
```
BenchmarkJSONIterEncode_fromStruct-12            5000000               258 ns/op              56 B/op          2 allocs/op
BenchmarkJSONIterDeDecode_toStruct-12            5000000               282 ns/op              80 B/op          3 allocs/op
BenchmarkJSONIterEncode_fromMap-12               3000000               404 ns/op             184 B/op          4 allocs/op
BenchmarkJSONIterDeDecode_toMap-12               3000000               581 ns/op             192 B/op         11 allocs/op
```