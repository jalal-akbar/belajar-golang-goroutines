package sync

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestPool(t *testing.T) {
	p := sync.Pool{
		New: func() any {
			return "Contain if Empty"
		},
	}
	var s = map[string]interface{}{
		"Name1": "Jalal",
		"Name2": "Muh",
		"Name3": "Akbar",
	}
	// Contain Data
	p.Put(s["Name1"])
	p.Put(s["Name2"])
	p.Put(s["Name3"])

	for i := 0; i < 10; i++ {
		go func() {
			data := p.Get() // Untuk Mendapatkan Data
			fmt.Println(data)
			time.Sleep(2 * time.Second)
			p.Put(data) // Untuk Meletekkan kembali data yang telah di ammbil
		}()
	}
	time.Sleep(4 * time.Second)
}
