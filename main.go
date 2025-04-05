package main

import (
	"runtime"
	"sync"
	// "fmt"
)

func collatz(n uint64, cache map[uint64]bool, mutx *sync.Mutex) bool {
	original := n
	for n != 1 {
		mutx.Lock()
		if cache[n] {
			// fmt.Println("found", n)
			cache[original] = true
			mutx.Unlock()
			break
		}
		mutx.Unlock()
		if n % 2 == 0 {
			n /= 2
		} else {
			n = 3 * n + 1
		}
	}
	mutx.Lock()
	cache[original] = true
	// fmt.Println("cached", original)
	mutx.Unlock()
	return true
}

func main() {
	runtime.GOMAXPROCS(16)
	cache := make(map[uint64]bool)
	var mutx sync.Mutex
	for n := uint64(1); n < 5000001; n++ {
		go collatz(n, cache, &mutx)
	}
}
