package main

import (
	"runtime"
	"sync"
	// "fmt"
)

func collatz(n uint64, cache *sync.Map) bool {
	original := n
	for n != 1 {
		if _, ok := cache.Load(n); ok {
			cache.Store(original, true)
			break
		}
		if n % 2 == 0 {
			n /= 2
		} else {
			n = 3 * n + 1
		}
	}
	cache.Store(original, true)
	return true
}

func main() {
	runtime.GOMAXPROCS(20)
	// cache := make(map[uint64]bool)
	// var mutx sync.Mutex
	var cache sync.Map
	for n := uint64(1); n < 5000000001; n++ {
		go collatz(n, &cache)
	}
}
