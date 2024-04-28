// Package bank provides a concurrency-safe bank with one account
package bank

var deposits = make(chan int) // send amount to deposint
var balances = make(chan int) // receive balance

var (
	sema    = make(chan struct{}, 1) // a binary semaphore balance
	balance int
)

func Deposit(amount int) {
	sema <- struct{}{} // acquire token
	balance = balance + amount
	<-sema // release token
}
func Balance() int {
	sema <- struct{}{} // acquire token
	b := balance
	<-sema // release token
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
