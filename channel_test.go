package belajar_golang_goroutines

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

// Membuat Channel
func TestCreateChannel(t *testing.T) {
	// Membuat Channel
	channel := make(chan string)
	// Membuat Goroutines
	go func() {
		time.Sleep(2 * time.Second)
		// Mengirim Data Ke Channel
		channel <- "Jalaluddin Muh Akbar"
		fmt.Println("Sukses Mengirim Data")
	}()
	// Menerima Data Dari Channel
	data := <-channel
	fmt.Println("Sukses Menerima Data")
	fmt.Println(data)
	// Menutup Channel
	defer close(channel) // Wajib Menutup Channel
}

// Channel Sebagai Parameter
func GiveMeResponse(channel chan string) {
	time.Sleep(2 * time.Second)
	channel <- "Jalaluddin Muh Akbar"
}

func TestChannelAsParameter(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go GiveMeResponse(channel)
	fmt.Println("Success Send Data to Channel")

	time.Sleep(5 * time.Second)
	data := <-channel
	fmt.Println("Receive Data From Channel")

	fmt.Println("Data :", data)

}

// Channel In/Out
func SendChannel(c chan<- string) { // Use <- After chan
	time.Sleep(1 * time.Second)
	c <- "JalaluddinMuh Akbar"
}
func ReceiveChannel(c <-chan string) { // Use <- Before chan
	data := <-c
	fmt.Println(data)
}
func TestChannelInOut(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go SendChannel(channel)
	go ReceiveChannel(channel)

	time.Sleep(5 * time.Second)
	fmt.Println("Success")
}

// Buffered Channel
func TestBufferedChannel(t *testing.T) {
	channnel := make(chan string, 3) // Default make(chan string)
	defer close(channnel)

	channnel <- "Jalaluddin"
	channnel <- "Muh"
	channnel <- "Akbar"

	data1 := <-channnel
	data2 := <-channnel
	data3 := <-channnel
	fmt.Println(data1)
	fmt.Println(data2)
	fmt.Println(data3)

}

// Range Channel
func TestRangeChannel(t *testing.T) {
	channel := make(chan string)

	go func() {
		for i := 0; i < 10; i++ {
			channel <- "Perulangan Ke :" + strconv.Itoa(i)
		}
		defer close(channel)
	}()

	for data := range channel {
		fmt.Println(data)
	}
	time.Sleep(5 * time.Second)
}

// Select Channel
func TestSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)

	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	count := 0
	for {
		select {
		case data1 := <-channel1:
			fmt.Println("Data Dari Channel 1", data1)
			count++
		case data2 := <-channel2:
			fmt.Println("Data Dari Channel 2", data2)
			count++
		}
		if count == 2 {
			break
		}
	}

	defer close(channel1)
	defer close(channel2)
}

// Default Channel
func TestDefaultSelect(t *testing.T) {
	chan1 := make(chan string)
	chan2 := make(chan string)

	go GiveMeResponse(chan1)
	go GiveMeResponse(chan2)

	count := 0
	for {
		select {
		case data1 := <-chan1:
			fmt.Println("Data Dari Channel 1", data1)
			count++
		case data2 := <-chan2:
			fmt.Println("Data Dari Channel 2", data2)
			count++
		default:
			fmt.Println("Menunggu Data")
		}
		if count == 2 {
			break
		}
	}

	defer close(chan1)
	defer close(chan2)
}
