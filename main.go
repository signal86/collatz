package main

import (
	"runtime"
	"sync"
	// "fmt"
)

func collatz(n uint64, cache *sync.Map, wg *sync.WaitGroup) {
	defer wg.Done()
	original := n
	for n != 1 {
		if _, ok := cache.Load(n); ok {
			cache.Store(original, true)
			return
		}
		if n % 2 == 0 {
			n /= 2
		} else {
			n = 3 * n + 1
		}
	}
	cache.Store(original, true)
}

func main() {
	runtime.GOMAXPROCS(20)
	// cache := make(map[uint64]bool)
	// var mutx sync.Mutex
	var cache sync.Map
	var wg sync.WaitGroup
	jobs := make(chan uint64, 100000)
	for i := 0; i < 100; i++ {
		go func() {
			for n := range jobs {
				collatz(n, &cache, &wg)
			}
		}()
	}
	for n := uint64(1); n < 50000001; n++ {
		wg.Add(1)
		jobs <- n
	}
	close(jobs)
	wg.Wait()
}
