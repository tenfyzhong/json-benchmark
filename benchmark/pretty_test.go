package benchmark

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/tidwall/pretty"
)

// ==================== Pretty Format Benchmarks ====================
// Compare tidwall/pretty with encoding/json.Indent for JSON formatting.

func BenchmarkPretty_Small_TidwallPretty(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = pretty.Pretty(smallStructJSON)
	}
}

func BenchmarkPretty_Small_StdLibIndent(b *testing.B) {
	var buf bytes.Buffer
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		buf.Reset()
		_ = json.Indent(&buf, smallStructJSON, "", "  ")
	}
}

func BenchmarkPretty_Large_TidwallPretty(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = pretty.Pretty(largeStructJSON)
	}
}

func BenchmarkPretty_Large_StdLibIndent(b *testing.B) {
	var buf bytes.Buffer
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		buf.Reset()
		_ = json.Indent(&buf, largeStructJSON, "", "  ")
	}
}

// ==================== Ugly (Compact) Benchmarks ====================
// Compare tidwall/pretty.Ugly with encoding/json.Compact.

func BenchmarkUgly_Small_TidwallUgly(b *testing.B) {
	// Pre-format JSON for uglify test.
	prettyJSON := pretty.Pretty(smallStructJSON)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = pretty.Ugly(prettyJSON)
	}
}

func BenchmarkUgly_Small_StdLibCompact(b *testing.B) {
	prettyJSON := pretty.Pretty(smallStructJSON)
	var buf bytes.Buffer
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		buf.Reset()
		_ = json.Compact(&buf, prettyJSON)
	}
}

func BenchmarkUgly_Large_TidwallUgly(b *testing.B) {
	prettyJSON := pretty.Pretty(largeStructJSON)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = pretty.Ugly(prettyJSON)
	}
}

func BenchmarkUgly_Large_StdLibCompact(b *testing.B) {
	prettyJSON := pretty.Pretty(largeStructJSON)
	var buf bytes.Buffer
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		buf.Reset()
		_ = json.Compact(&buf, prettyJSON)
	}
}
