package time

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

var group = sync.WaitGroup{}

// Ticker
func TestTicker(t *testing.T) {
	ticker := time.NewTicker(2 * time.Second)

	go func() {
		time.Sleep(2 * time.Second)
		ticker.Stop()
	}()

	for tick := range ticker.C {
		fmt.Println(tick)
	}
}

// time.Tick()
func TestTick(t *testing.T) {
	channel := time.Tick(1 * time.Second)
	group.Add(1)
	defer group.Done()
	for tick := range channel {
		fmt.Println(tick)
	}
	group.Wait()
}
