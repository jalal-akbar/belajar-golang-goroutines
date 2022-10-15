package belajar_golang_goroutines

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

func TestGOMAXPROCS(t *testing.T) {
	for i := 0; i < 10; i++ {
		group.Add(1)
		go func() {
			time.Sleep(3 * time.Second)
			group.Done()
		}()
	}

	totalCPU := runtime.NumCPU()
	fmt.Println("Total CPU", totalCPU)

	totalGOMAXPROCS := runtime.GOMAXPROCS(-1)
	fmt.Println("Total Thread", totalGOMAXPROCS)

	totalGoroutines := runtime.NumGoroutine()
	fmt.Println("Total GoRoutines", totalGoroutines)

	group.Wait()
}
