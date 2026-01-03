package benchmark

import (
	"encoding/json"
	"testing"

	gojson "github.com/goccy/go-json"
	jsoniter "github.com/json-iterator/go"
	"github.com/mailru/easyjson"
	"github.com/tenfyzhong/json-benchmark/internal/testdata"
)

// jsonIterator is configured to be compatible with standard library.
var jsonIterator = jsoniter.ConfigCompatibleWithStandardLibrary

// Pre-generated test data to avoid regeneration during benchmark.
var (
	smallStruct     testdata.SmallStruct
	largeStruct     testdata.LargeStruct
	smallStructJSON []byte
	largeStructJSON []byte
)

func init() {
	// Use fixed seed for reproducible test data.
	smallStruct = testdata.GenerateSmallStruct(42)
	largeStruct = testdata.GenerateLargeStruct(42)
	smallStructJSON = testdata.GenerateSmallStructJSON(42)
	largeStructJSON = testdata.GenerateLargeStructJSON(42)
}

// ==================== Small Struct Encode Benchmarks ====================

func BenchmarkEncode_Small_StdLib(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = json.Marshal(&smallStruct)
	}
}

func BenchmarkEncode_Small_GoJSON(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = gojson.Marshal(&smallStruct)
	}
}

func BenchmarkEncode_Small_JSONIterator(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = jsonIterator.Marshal(&smallStruct)
	}
}

func BenchmarkEncode_Small_EasyJSON(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = easyjson.Marshal(&smallStruct)
	}
}

// ==================== Large Struct Encode Benchmarks ====================

func BenchmarkEncode_Large_StdLib(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = json.Marshal(&largeStruct)
	}
}

func BenchmarkEncode_Large_GoJSON(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = gojson.Marshal(&largeStruct)
	}
}

func BenchmarkEncode_Large_JSONIterator(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = jsonIterator.Marshal(&largeStruct)
	}
}

func BenchmarkEncode_Large_EasyJSON(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = easyjson.Marshal(&largeStruct)
	}
}
