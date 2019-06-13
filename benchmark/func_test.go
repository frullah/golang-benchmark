package benchmark

import (
	"testing"
)

type EmptyStruct struct{}

func BenchmarkCallClosure(b *testing.B) {
	closure := noop
	for i := 0; i < b.N; i++ {
		closure()
	}
}

func BenchmarkCallFunction(b *testing.B) {
	for i := 0; i < b.N; i++ {
		noop()
	}
}

func BenchmarkSpreadOperator(b *testing.B) {
	slices := []EmptyStruct{{}, {}, {}, {}, {}, {}}
	for i := 0; i < b.N; i++ {
		fnWithVariadicParameter(slices...)
	}
}

func noop() {
}

func fnWithVariadicParameter(values ...EmptyStruct) {
}
