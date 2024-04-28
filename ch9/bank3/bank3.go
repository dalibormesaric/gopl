// Package bank provides a concurrency-safe bank with one account
package bank

import "sync"

var deposits = make(chan int) // send amount to deposint
var balances = make(chan int) // receive balance

var (
	mu      sync.Mutex // guards balance
	balance int
)

func Deposit(amount int) {
	mu.Lock()
	balance = balance + amount
	mu.Unlock()
}
func Balance() int {
	mu.Lock()
	b := balance
	mu.Unlock()
	return b
}

func teller() {
	var balance int // balance is confined to teller goroutine
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case balances <- balance:
		}
	}
}

func init() {
	go teller() // start the monitor goroutine
}
