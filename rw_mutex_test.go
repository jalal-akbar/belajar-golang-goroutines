package belajar_golang_goroutines

import (
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type BankAccount struct {
	Balance int
	RWMutex sync.RWMutex
}

func (account *BankAccount) AddBalance(amount int) {
	account.RWMutex.Lock()
	account.Balance = account.Balance + amount
	account.RWMutex.Unlock()
}

func (account *BankAccount) GetBalance() int {
	account.RWMutex.RLock()
	get := account.Balance
	account.RWMutex.RUnlock()
	return get
}

func TestGetBalanceBankAccount(t *testing.T) {
	accountBalance := BankAccount{}

	for i := 0; i < 10000; i++ {
		go func() {
			for j := 0; j < 10; j++ {
				accountBalance.AddBalance(1)
			}
		}()
	}
	time.Sleep(5 * time.Second)
	assert.Equal(t, 100000, accountBalance.GetBalance())
}
