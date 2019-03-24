package main_test

import "testing"

func copy(sc string) string {
	ss := sc
	return ss
}

func reference(sr string) *string {
	ss := &sr
	return ss
}
func BenchmarkAllocateStack(b *testing.B) {
	s := "string"
	for i := 0; i < b.N; i++ {
		copy(s)
	}
}

func BenchmarkAllocateHeap(b *testing.B) {
	s := "string"
	for i := 0; i < b.N; i++ {
		reference(s)
	}
}
