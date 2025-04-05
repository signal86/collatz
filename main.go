package main

import (
	"runtime"
	"sync"
	// "fmt"
)

func collatz(n uint64, cache *sync.Map, wg *sync.WaitGroup) {
	defer wg.Done()
	original := n
	var cache2 []uint64
	for n != 1 {
		if _, ok := cache.Load(n); ok {
			for _, i := range cache2 {
				cache.Store(i, true)
			}
			cache.Store(original, true)
			return
		}
		cache2 = append(cache2, n)
		if n % 2 == 0 {
			n /= 2
		} else {
			n = 3 * n + 1
		}
	}
	for _, i := range cache2 {
		cache.Store(i, true)
	}
	cache.Store(original, true)
}

func main() {
	runtime.GOMAXPROCS(20)
	// cache := make(map[uint64]bool)
	// var mutx sync.Mutex
	var cache sync.Map
	var wg sync.WaitGroup
	for n := uint64(1); n < 5000000001; n++ {
		wg.Add(1)
		go collatz(n, &cache, &wg)
	}
	wg.Wait()
}
