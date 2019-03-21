package main

import "testing"

func BenchmarkAllocateStack(b *testing.B) {
	b.ReportAllocs()
	s := "string"
	for i := 0; i < b.N; i++ {
		ss := copy(s)
		_ = ss
	}
}

func BenchmarkAllocateHeap(b *testing.B) {
	b.ReportAllocs()
	s := "string"
	for i := 0; i < b.N; i++ {
		ss := reference(s)
		_ = ss
	}
}
