package golangbenchmark

import (
	"encoding/json"
	"testing"

	"github.com/frullah/go-test/testdata"
)

func BenchmarkJSONEncode_fromStruct(b *testing.B) {
	for i := 0; i < b.N; i++ {
		json.Marshal(&testdata.JSONDataStruct)
	}
}

func BenchmarkJSONDeDecode_toStruct(b *testing.B) {
	var data testdata.JSONObj
	for i := 0; i < b.N; i++ {
		json.Unmarshal([]byte(testdata.JSONStr), &data)
	}
}

func BenchmarkJSONEncode_fromMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		json.Marshal(&testdata.JSONDataMap)
	}
}

func BenchmarkJSONDeDecode_toMap(b *testing.B) {
	var data map[string]interface{}
	for i := 0; i < b.N; i++ {
		json.Unmarshal([]byte(testdata.JSONStr), &data)
	}
}
