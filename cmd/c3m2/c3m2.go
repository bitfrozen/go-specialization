package main

import (
	"fmt"
	"os"
	"sync"
)

// A race condition is a condition when programs outcome
// depends on nondeterministic interleaving of multiple goroutines, threads, or processes.

// In this case, we modify the same memory location from 2 different simultaneous goroutines.
// If goroutine addOne() runs first and increases the value at memory location to non-zero,
// division will succeed. But if goroutine divide() runs first,
// the goroutine will panic (print error).

// Since goroutine (or thread) interleaving is nondeterministic we run this test several times
// to increase the probability, that different behavior will manifest.

func addOne(n *int, w *sync.WaitGroup) {
	*n = *n + 1
	w.Done()
}

func divide(n *int, w *sync.WaitGroup) {
	defer func() {
		if err := recover(); err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "%v\n", err)
		}
		w.Done()
	}()
	_ = 5 / (*n)
	fmt.Println("Division correct")
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(2)
		num := 0
		go addOne(&num, &wg)
		go divide(&num, &wg)
		wg.Wait()
	}
}
