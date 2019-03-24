package main_test

import "testing"

func BenchmarkRecursiveStack(b *testing.B) {
	var res int
	for i := 0; i < b.N; i++ {
		res += recursiveStack(1000)
	}
}

func BenchmarkRecursiveHeap(b *testing.B) {
	var res int
	for i := 0; i < b.N; i++ {
		recursiveHeap(&res, 1000)
	}
}

func recursiveStack(count int) int {
	if count <= 0 {
		return 0
	}
	return recursiveStack(count-1) + 1
}

func recursiveHeap(result *int, count int) {
	if count <= 0 {
		return
	}
	*result++
	recursiveHeap(result, count-1)
}
