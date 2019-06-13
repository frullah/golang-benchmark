// Golang benchmark - Struct alignment memory usage
// Struct alignment in golang is 4 bytes for 32-bit, and 8 bytes for 64-bit
//
// To check struct memory alignment,
// You can use the tools in http://golang-sizeof.tips
package benchmark

import (
	"testing"
)

type BadStructAlign struct {
	IsActive bool   // 1 bytes
	Name     string // Pointer size
	Age      uint16 // 2 bytes
	Number   int8   //1 bytes
}

type GoodStructAlign struct {
	Name     string // Pointer size
	Age      uint16 // 2 bytes
	Number   int8   //1 bytes
	IsActive bool   //1 bytes
}

func BenchmarkGoodStructAlign(b *testing.B) {
	instances := make([]GoodStructAlign, b.N)
	for i := 0; i < b.N; i++ {
		instances[i] = GoodStructAlign{
			Number:   127,
			Name:     "Hello world",
			Age:      0xFF,
			IsActive: false,
		}
	}
}

func BenchmarkBadStructAlign(b *testing.B) {
	instances := make([]BadStructAlign, b.N)
	for i := 0; i < b.N; i++ {
		instances[i] = BadStructAlign{
			Number:   127,
			Name:     "Hello world",
			Age:      0xFF,
			IsActive: false,
		}
	}
}
