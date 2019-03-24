package deferperformance

import (
	"sync"
	"testing"
)

type Defer struct {
	sync.Mutex
	value int
}

func (r *Defer) withDefer() {
	r.Lock()
	defer r.Unlock()
	r.value++
}

func (r *Defer) noDefer() {
	r.Lock()
	r.value++
	r.Unlock()
}

func BenchmarkDefer(b *testing.B) {
	var d Defer
	for i := 0; i < b.N; i++ {
		d.withDefer()
	}
}

func BenchmarkNoDefer(b *testing.B) {
	var d Defer
	for i := 0; i < b.N; i++ {
		d.noDefer()
	}
}
