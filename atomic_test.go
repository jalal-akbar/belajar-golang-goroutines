package belajar_golang_goroutines

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
)

// Atomic
var group = sync.WaitGroup{}
var counter int64 = 0

func TestAtoomic(t *testing.T) {
	for i := 0; i < 10; i++ {
		group.Add(1)
		go func() {
			for j := 0; j < 10; j++ {
				atomic.AddInt64(&counter, 1)
			}
			defer group.Done()
		}()

	}
	group.Wait()
	fmt.Println("counter", counter)
}
