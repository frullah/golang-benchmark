package main

import (
	"testing"
)

func Fn1(a, b string) string {
	return a + b
}

func Fn2(a, b string) string {
	return b + a
}

func BenchmarkCallClosure(b *testing.B) {
	var closure func(url, handler string) string
	arg1 := "arg1 value"
	arg2 := "arg2 value"
	for i := 0; i < b.N; i++ {
		if i%2 == 0 {
			closure = Fn1
		} else {
			closure = Fn2
		}

		closure(arg1, arg2)
	}
}

func BenchmarkDirect(b *testing.B) {
	arg1 := "arg1 value"
	arg2 := "arg2 value"
	for i := 0; i < b.N; i++ {
		if i%2 == 0 {
			Fn1(arg1, arg2)
		} else {
			Fn2(arg1, arg2)
		}
	}
}
