package concurency

import (
	"sync"
	"testing"
)

func BenchmarkNoGoroutines(b *testing.B) {
	results := make(map[int]int, 1024)
	for i := 0; i < b.N; i++ {
		for index := 0; index < 1024; index++ {
			results[index] = index * index
		}
	}
}

func BenchmarkGoroutines(b *testing.B) {
	results := make(map[int]int, 1024)
	for i := 0; i < b.N; i++ {
		var mutex sync.Mutex
		var wg sync.WaitGroup
		wg.Add(1024)
		for index := 0; index < 1024; index++ {
			go func(j int) {
				mutex.Lock()
				results[j] = j * j
				mutex.Unlock()
				wg.Done()
			}(index)
		}
		wg.Wait()
	}
}

func BenchmarkChannels(b *testing.B) {
	type res struct {
		key, value int
	}
	results := make(map[int]int, 1024)
	resultCh := make(chan res)
	// resultCh := make(chan res, 1024)
	for i := 0; i < b.N; i++ {
		for index := 0; index < 1024; index++ {
			go func(j int) {
				resultCh <- res{j, j * j}
			}(index)
		}
		for index := 0; index < 1024; index++ {
			r := <-resultCh
			results[r.key] = r.value
		}
	}
}

func BenchmarkSyncMap(b *testing.B) {
	var results sync.Map
	for i := 0; i < b.N; i++ {
		for index := 0; index < 1024; index++ {
			results.Store(index, index*index)
		}
	}
}
