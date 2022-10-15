package belajar_golang_goroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

type UserBalance struct {
	sync.Mutex
	Name    string
	Balance int
}

func (user *UserBalance) Lock() {
	user.Mutex.Lock()
}

func (user *UserBalance) Unlock() {
	user.Mutex.Unlock()
}

func (user *UserBalance) Change(amount int) {
	user.Balance = user.Balance + amount
}

func Transfer(user1 *UserBalance, user2 *UserBalance, amount int) {
	user1.Lock()
	fmt.Println("Lock User 1", user1.Name)
	user1.Change(-amount)

	time.Sleep(2 * time.Second)

	user2.Lock()
	fmt.Println("Lock User 2", user2.Name)
	user2.Change(amount)

	time.Sleep(2 * time.Second)

	user1.Unlock()
	user2.Unlock()
}

func TestDeadlockTransfer(t *testing.T) {
	user1 := UserBalance{
		Name:    "Jalal",
		Balance: 100000,
	}
	user2 := UserBalance{
		Name:    "Akbar",
		Balance: 100000,
	}
	go Transfer(&user1, &user2, 20000)
	go Transfer(&user2, &user1, 20000)
	time.Sleep(5 * time.Second)

	fmt.Println("User 1 :", user1.Name, "Balance", user1.Balance)
	fmt.Println("User 2 :", user2.Name, "Balance", user2.Balance)
}