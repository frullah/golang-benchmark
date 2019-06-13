package benchmark

import (
	"testing"
)

type EmptyStruct struct{}

func BenchmarkFunction_closureCall(b *testing.B) {
	closure := noop
	for i := 0; i < b.N; i++ {
		closure()
	}
}

func BenchmarkFunction_directCall(b *testing.B) {
	for i := 0; i < b.N; i++ {
		noop()
	}
}

func BenchmarkFunction_variadicParameter(b *testing.B) {
	slices := []EmptyStruct{{}, {}, {}, {}, {}, {}}
	for i := 0; i < b.N; i++ {
		fnWithVariadicParameter(slices...)
	}
}

func noop() {
}

func fnWithVariadicParameter(values ...EmptyStruct) {
}
