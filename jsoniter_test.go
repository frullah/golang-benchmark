package golangbenchmark

import (
	"testing"

	"github.com/frullah/go-test/testdata"
	jsoniter "github.com/json-iterator/go"
)

func BenchmarkJSONIterEncode_fromStruct(b *testing.B) {
	for i := 0; i < b.N; i++ {
		jsoniter.Marshal(&testdata.JSONDataStruct)
	}
}

func BenchmarkJSONIterDeDecode_toStruct(b *testing.B) {
	var data testdata.JSONObj
	for i := 0; i < b.N; i++ {
		jsoniter.Unmarshal([]byte(testdata.JSONStr), &data)
	}
}

func BenchmarkJSONIterEncode_fromMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		jsoniter.Marshal(&testdata.JSONDataMap)
	}
}

func BenchmarkJSONIterDeDecode_toMap(b *testing.B) {
	var data map[string]interface{}
	for i := 0; i < b.N; i++ {
		jsoniter.Unmarshal([]byte(testdata.JSONStr), &data)
	}
}
