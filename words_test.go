package main

import (
	"testing"
)

func Benchmark_lowerAndTrimDot(b *testing.B) {
	for n := 0; n < b.N; n++ {
		lowerAndTrimDot("Laborum.")
	}
}

func Benchmark_OldGetWordCount(b *testing.B) {
	for n := 0; n < b.N; n++ {
		OldGetWordCount("Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.")
	}
}

func Benchmark_SplitTextToMapNode(b *testing.B) {
	for n := 0; n < b.N; n++ {
		SplitTextToMapNode("Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.")
	}
}

func Benchmark_SplitTextToMapStr(b *testing.B) {
	for n := 0; n < b.N; n++ {
		SplitTextToMapStr("Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.")
	}
}

func Benchmark_SliceTextToMapNode(b *testing.B) {
	for n := 0; n < b.N; n++ {
		SliceTextToMapNode("Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.")
	}
}

func GenerateMapNode() map[string]*WordNode {
	result := map[string]*WordNode{
		"abc1": {str: "abc1", count: 2},
		"abc2": {str: "abc2", count: 3},
		"abc3": {str: "abc3", count: 9},
		"abc4": {str: "abc4", count: 6},
		"abc5": {str: "abc5", count: 1},
		"abc6": {str: "abc6", count: 8},
		"abc7": {str: "abc7", count: 3},
		"abc8": {str: "abc8", count: 3},
	}

	return result
}
func Benchmark_MapNodeToArray(b *testing.B) {
	for n := 0; n < b.N; n++ {
		MapNodeToArray(GenerateMapNode())
	}
}

func GenerateMapStr() map[string]int {
	result := map[string]int{
		"abc1": 2,
		"abc2": 3,
		"abc3": 9,
		"abc4": 6,
		"abc5": 1,
		"abc6": 8,
		"abc7": 3,
		"abc8": 3,
	}

	return result
}
func Benchmark_MapStrToArray(b *testing.B) {
	for n := 0; n < b.N; n++ {
		MapStrToArray(GenerateMapStr())
	}
}

func Benchmark_OldFunc(b *testing.B) {
	for n := 0; n < b.N; n++ {
		OldGetTopK(GenerateMapStr(), 3)
	}
}

func Benchmark_TopKBinarySearchMapStr(b *testing.B) {
	for n := 0; n < b.N; n++ {
		TopKBinarySearchMapStr(GenerateMapStr(), 3)
	}
}

func Benchmark_TopKBinarySearchMapNode(b *testing.B) {
	for n := 0; n < b.N; n++ {
		TopKBinarySearchMapNode(GenerateMapNode(), 3)
	}
}

func Benchmark_TopKSortedArrayFromMapStr(b *testing.B) {
	for n := 0; n < b.N; n++ {
		TopKSortedArrayFromMapStr(GenerateMapStr(), 3)
	}
}

func Benchmark_TopKSortedArrayFromMapNode(b *testing.B) {
	for n := 0; n < b.N; n++ {
		TopKSortedArrayFromMapNode(GenerateMapNode(), 3)
	}
}

/**
$ go test -bench . -benchmem
goos: windows
goarch: amd64
cpu: AMD Ryzen 9 4900HS with Radeon Graphics
Benchmark_lowerAndTrimDot-16                    21477931                52.17 ns/op            8 B/op          1 allocs/op
Benchmark_OldGetWordCount-16                      101523             11302 ns/op            9019 B/op         13 allocs/op
Benchmark_SplitTextToMapNode-16                    73424             16524 ns/op           10530 B/op         76 allocs/op
Benchmark_SplitTextToMapStr-16                    103392             11448 ns/op            9018 B/op         13 allocs/op
Benchmark_SliceTextToMapNode-16                    77559             15874 ns/op            9379 B/op         75 allocs/op
Benchmark_MapNodeToArray-16                      1293907               934.0 ns/op           312 B/op         12 allocs/op
Benchmark_MapStrToArray-16                       1261267               962.8 ns/op           312 B/op         12 allocs/op
Benchmark_OldFunc-16                              998718              1279 ns/op             368 B/op         14 allocs/op
Benchmark_TopKBinarySearchMapStr-16              1656986               728.8 ns/op           166 B/op          6 allocs/op
Benchmark_TopKBinarySearchMapNode-16             1357292               882.4 ns/op           224 B/op          9 allocs/op
Benchmark_TopKSortedArrayFromMapStr-16            999092              1233 ns/op             368 B/op         14 allocs/op
Benchmark_TopKSortedArrayFromMapNode-16           923396              1276 ns/op             368 B/op         14 allocs/op
**/
