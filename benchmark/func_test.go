package benchmark

import (
	"testing"
)

type EmptyStruct struct{}

var noopClosure = noop

func BenchmarkFunction_closureCall(b *testing.B) {
	for i := 0; i < b.N; i++ {
		noopClosure()
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
