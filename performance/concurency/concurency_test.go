package concurency

import (
	"sync"
	"testing"
)

var size = 24

func fib(n int) int {
	if n <= 1 {
		return n
	}
	return fib(n-1) + fib(n-2)
}

func BenchmarkSequential(b *testing.B) {
	results := make(map[int]int, size)
	for i := 0; i < b.N; i++ {
		for j := 0; j < size; j++ {
			results[j] = fib(j)
		}
	}
}

func BenchmarkMutex(b *testing.B) {
	results := make(map[int]int, size)
	for i := 0; i < b.N; i++ {
		var mutex sync.Mutex
		var wg sync.WaitGroup
		wg.Add(size)
		for j := 0; j < size; j++ {
			go func(j int) {
				mutex.Lock()
				results[j] = fib(j)
				mutex.Unlock()
				wg.Done()
			}(j)
		}
		wg.Wait()
	}
}

func BenchmarkChannels(b *testing.B) {
	type res struct {
		key   int
		value int
	}
	results := make(map[int]int, size)
	resultCh := make(chan res)
	for i := 0; i < b.N; i++ {
		for j := 0; j < size; j++ {
			go func(j int) {
				resultCh <- res{j, fib(j)}
			}(j)
		}
		for index := 0; index < size; index++ {
			r := <-resultCh
			results[r.key] = r.value
		}
	}
}

func BenchmarkSyncMap(b *testing.B) {
	var results sync.Map
	for i := 0; i < b.N; i++ {
		var wg sync.WaitGroup
		wg.Add(size)
		for j := 0; j < size; j++ {
			go func(j int) {
				results.Store(j, fib(j))
				wg.Done()
			}(j)
		}
		wg.Wait()
	}
}
