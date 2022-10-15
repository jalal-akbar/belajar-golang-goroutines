package sync

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func syncMap(m *sync.Map, value int) {
	m.Store(value, value)
}
func TestMap(t *testing.T) {
	m := &sync.Map{}
	for i := 0; i < 10; i++ {
		go syncMap(m, i)
		time.Sleep(2 * time.Second)
	}
	time.Sleep(2 * time.Second)
	m.Range(func(key, value any) bool {
		fmt.Println(key, ":", value)
		return true
	})
}
