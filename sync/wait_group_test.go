package sync

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func WaitGroup(waitGroup *sync.WaitGroup) {
	defer waitGroup.Done()

	waitGroup.Add(1)

	fmt.Println("Hello This is WaitGroup")
	time.Sleep(5 * time.Second)
}

func TestWaitGroup(t *testing.T) {
	waitGroup := new(sync.WaitGroup)

	for i := 0; i < 100; i++ {
		go WaitGroup(waitGroup)
	}

	waitGroup.Wait()
	fmt.Println("Complete")
}
