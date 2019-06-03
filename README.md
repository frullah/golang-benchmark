# Golang benchmark

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