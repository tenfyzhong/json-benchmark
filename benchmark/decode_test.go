package benchmark

import (
	"encoding/json"
	"testing"

	gojson "github.com/goccy/go-json"
	"github.com/mailru/easyjson"
	"github.com/tenfyzhong/json-benchmark/internal/testdata"
)

// ==================== Small Struct Decode Benchmarks ====================

func BenchmarkDecode_Small_StdLib(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var s testdata.SmallStruct
		_ = json.Unmarshal(smallStructJSON, &s)
	}
}

func BenchmarkDecode_Small_GoJSON(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var s testdata.SmallStruct
		_ = gojson.Unmarshal(smallStructJSON, &s)
	}
}

func BenchmarkDecode_Small_JSONIterator(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var s testdata.SmallStruct
		_ = jsonIterator.Unmarshal(smallStructJSON, &s)
	}
}

func BenchmarkDecode_Small_EasyJSON(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var s testdata.SmallStruct
		_ = easyjson.Unmarshal(smallStructJSON, &s)
	}
}

// ==================== Large Struct Decode Benchmarks ====================

func BenchmarkDecode_Large_StdLib(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var s testdata.LargeStruct
		_ = json.Unmarshal(largeStructJSON, &s)
	}
}

func BenchmarkDecode_Large_GoJSON(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var s testdata.LargeStruct
		_ = gojson.Unmarshal(largeStructJSON, &s)
	}
}

func BenchmarkDecode_Large_JSONIterator(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var s testdata.LargeStruct
		_ = jsonIterator.Unmarshal(largeStructJSON, &s)
	}
}

func BenchmarkDecode_Large_EasyJSON(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var s testdata.LargeStruct
		_ = easyjson.Unmarshal(largeStructJSON, &s)
	}
}
