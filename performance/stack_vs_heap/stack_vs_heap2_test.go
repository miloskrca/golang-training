package main

import "testing"

func BenchmarkRecursiveStack(b *testing.B) {
	b.ReportAllocs()
	var res int
	for i := 0; i < b.N; i++ {
		res += recursive1(1000)
	}
}

func BenchmarkRecursiveHeap(b *testing.B) {
	b.ReportAllocs()
	var res int
	for i := 0; i < b.N; i++ {
		recursive2(&res, 1000)
	}
}

func recursive1(count int) int {
	if count <= 0 {
		return 0
	}
	return recursive1(count-1) + 1
}

func recursive2(result *int, count int) {
	if count <= 0 {
		return
	}
	*result++
	recursive2(result, count-1)
}
