package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"runtime"
	"sort"
	"strconv"
	"strings"
	"time"
)

// BenchResult holds the result of a single benchmark.
type BenchResult struct {
	Name        string
	Category    string  // Encode, Decode, Pretty, Ugly
	Size        string  // Small, Large
	Library     string  // StdLib, GoJSON, JSONIterator, EasyJSON, TidwallPretty, etc.
	NsPerOp     float64
	BytesPerOp  int64
	AllocsPerOp int64
	Iterations  int64
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: report <benchmark_output_file>")
		os.Exit(1)
	}

	results := parseBenchmarkOutput(os.Args[1])
	if len(results) == 0 {
		fmt.Fprintln(os.Stderr, "No benchmark results found")
		os.Exit(1)
	}

	generateREADME(results)
}

func parseBenchmarkOutput(filename string) []BenchResult {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	var results []BenchResult
	// Match benchmark output lines:
	// BenchmarkEncode_Small_StdLib-8    1000000    1234 ns/op    256 B/op    3 allocs/op
	re := regexp.MustCompile(`^Benchmark(\w+)_(\w+)_(\w+)-\d+\s+(\d+)\s+([\d.]+)\s+ns/op\s+(\d+)\s+B/op\s+(\d+)\s+allocs/op`)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		matches := re.FindStringSubmatch(line)
		if len(matches) == 8 {
			nsPerOp, _ := strconv.ParseFloat(matches[5], 64)
			bytesPerOp, _ := strconv.ParseInt(matches[6], 10, 64)
			allocsPerOp, _ := strconv.ParseInt(matches[7], 10, 64)
			iterations, _ := strconv.ParseInt(matches[4], 10, 64)

			results = append(results, BenchResult{
				Name:        matches[0],
				Category:    matches[1],
				Size:        matches[2],
				Library:     matches[3],
				NsPerOp:     nsPerOp,
				BytesPerOp:  bytesPerOp,
				AllocsPerOp: allocsPerOp,
				Iterations:  iterations,
			})
		}
	}
	return results
}

func generateREADME(results []BenchResult) {
	fmt.Println("# Go JSON Library Benchmark")
	fmt.Println()
	fmt.Println("Performance comparison of popular Go JSON libraries.")
	fmt.Println()

	printEnvironmentInfo()

	fmt.Println("## Libraries Tested")
	fmt.Println()
	fmt.Println("| Library | Description |")
	fmt.Println("|---------|-------------|")
	fmt.Println("| `encoding/json` | Go standard library |")
	fmt.Println("| `goccy/go-json` | High-performance JSON encoder/decoder |")
	fmt.Println("| `json-iterator/go` | Drop-in replacement with better performance |")
	fmt.Println("| `mailru/easyjson` | Code generation based, fastest but requires build step |")
	fmt.Println("| `tidwall/pretty` | JSON formatting/beautifying only (not full encoder/decoder) |")
	fmt.Println()

	// Group results by category and size.
	groups := make(map[string][]BenchResult)
	for _, r := range results {
		key := r.Category + "_" + r.Size
		groups[key] = append(groups[key], r)
	}

	// Print results in order.
	fmt.Println("## Benchmark Results")
	fmt.Println()

	sections := []struct {
		category string
		size     string
		title    string
	}{
		{"Encode", "Small", "Encode - Small Struct (~10 fields)"},
		{"Encode", "Large", "Encode - Large Struct (~50 fields, nested)"},
		{"Decode", "Small", "Decode - Small Struct"},
		{"Decode", "Large", "Decode - Large Struct"},
		{"Pretty", "Small", "Pretty Print - Small JSON"},
		{"Pretty", "Large", "Pretty Print - Large JSON"},
		{"Ugly", "Small", "Compact (Ugly) - Small JSON"},
		{"Ugly", "Large", "Compact (Ugly) - Large JSON"},
	}

	for _, sec := range sections {
		key := sec.category + "_" + sec.size
		if group, ok := groups[key]; ok && len(group) > 0 {
			fmt.Printf("### %s\n\n", sec.title)
			printTable(group)
			printASCIIChart(group)
			fmt.Println()
		}
	}

	fmt.Println("## How to Run")
	fmt.Println()
	fmt.Println("```bash")
	fmt.Println("# Run all benchmarks")
	fmt.Println("make benchmark")
	fmt.Println()
	fmt.Println("# Generate report")
	fmt.Println("make report")
	fmt.Println()
	fmt.Println("# Or manually")
	fmt.Println("go test -bench=. -benchmem ./benchmark/... | tee benchmark_results.txt")
	fmt.Println("go run ./cmd/report/main.go benchmark_results.txt")
	fmt.Println("```")
	fmt.Println()

	fmt.Println("## License")
	fmt.Println()
	fmt.Println("MIT License")
}

func printEnvironmentInfo() {
	fmt.Println("## Test Environment")
	fmt.Println()
	fmt.Printf("- **Go Version**: %s\n", runtime.Version())
	fmt.Printf("- **OS/Arch**: %s/%s\n", runtime.GOOS, runtime.GOARCH)
	fmt.Printf("- **CPU Cores**: %d\n", runtime.NumCPU())
	fmt.Printf("- **Date**: %s\n", time.Now().Format("2006-01-02"))
	fmt.Println()
}

func printTable(results []BenchResult) {
	// Sort by performance (fastest first).
	sort.Slice(results, func(i, j int) bool {
		return results[i].NsPerOp < results[j].NsPerOp
	})

	// Find the slowest as baseline for speedup calculation.
	baseline := results[len(results)-1].NsPerOp

	fmt.Println("| Rank | Library | ns/op | B/op | allocs/op | Speedup vs Slowest |")
	fmt.Println("|:----:|---------|------:|-----:|----------:|-------------------:|")

	for rank, r := range results {
		speedup := baseline / r.NsPerOp
		libName := formatLibraryName(r.Library)
		fmt.Printf("| %d | %s | %.2f | %d | %d | %.2fx |\n",
			rank+1, libName, r.NsPerOp, r.BytesPerOp, r.AllocsPerOp, speedup)
	}
	fmt.Println()
}

func printASCIIChart(results []BenchResult) {
	// Find max value for scaling.
	maxNs := 0.0
	for _, r := range results {
		if r.NsPerOp > maxNs {
			maxNs = r.NsPerOp
		}
	}

	const maxWidth = 40

	// Sort by performance (fastest first).
	sort.Slice(results, func(i, j int) bool {
		return results[i].NsPerOp < results[j].NsPerOp
	})

	fmt.Println("```")
	fmt.Println("Performance Comparison (lower is better):")
	fmt.Println()

	for _, r := range results {
		barWidth := int((r.NsPerOp / maxNs) * maxWidth)
		if barWidth < 1 {
			barWidth = 1
		}
		bar := strings.Repeat("█", barWidth)
		libName := formatLibraryName(r.Library)
		fmt.Printf("%-15s %s %.0f ns/op\n", libName, bar, r.NsPerOp)
	}
	fmt.Println("```")
}

func formatLibraryName(lib string) string {
	switch lib {
	case "StdLib":
		return "encoding/json"
	case "StdLibIndent":
		return "json.Indent"
	case "StdLibCompact":
		return "json.Compact"
	case "GoJSON":
		return "go-json"
	case "JSONIterator":
		return "json-iterator"
	case "EasyJSON":
		return "easyjson"
	case "TidwallPretty":
		return "pretty.Pretty"
	case "TidwallUgly":
		return "pretty.Ugly"
	default:
		return lib
	}
}
