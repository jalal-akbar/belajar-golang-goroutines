package belajar_golang_goroutines

import (
	"fmt"
	"testing"
	"time"
)

// Membuat Goroutines
func RunHelloWorld() {
	fmt.Println("Hello World")
}

func TestCreateGoroutines(t *testing.T) {
	go RunHelloWorld()
	fmt.Println("Ups")

	time.Sleep(1 * time.Second)
}

// Membuat Banyak Goroutines
func DisplayNumber(number int) {
	fmt.Println("Nomor Ke :", number)
}

func TestManyGoroutines(t *testing.T) {
	for i := 0; i < 100000; i++ {
		go DisplayNumber(i)
	}

	time.Sleep(10 * time.Second)
}
