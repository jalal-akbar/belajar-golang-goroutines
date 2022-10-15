package belajar_golang_goroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestMutex(t *testing.T) {
	m := new(sync.Mutex)
	x := 0
	for i := 1; i <= 1000; i++ {
		go func() {
			for j := 1; j <= 100; j++ {
				m.Lock() // Locking
				x = x + 1
				m.Unlock() // Unlock
			}
		}()
	}
	time.Sleep(2 * time.Second)
	fmt.Println("x :", x)
}
