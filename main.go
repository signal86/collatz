package main

import (
	"runtime"
)

func collatz(n uint64) bool {
	for n != 1 {
		if n % 2 == 0 {
			n /= 2
		} else {
			n = 3 * n + 1
		}
	}
	return true
}

func main() {
	runtime.GOMAXPROCS(16)
	for n := uint64(1); n < 5000000001; n++ {
		go collatz(n)
	}
}
