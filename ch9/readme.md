# Concurrency with Shared Variables

> Do not communicate by sharing memory;  instead, share memory by communicating.

> three ways to avoid a data race:
>   - not to write the variable
>   - avoid accessing the variable from multiple goroutines
>   - allow many goroutines to access the variable, but only one at the time - mutual exclusion

### bank1

> a package

> serial confinement
> ``` go
> type Cake struct{ state string }
>
> func baker(cooked chan<- *Cake) {
>   for {
>       cake := new(Cake)
>       cake.state = "cooked"
>       cooked <- cake // baker never touches this cake again
>   }
> }
>
> func icer(iced chan<- *Cake, cooked <-chan *Cake) {
>   for cake := range cooked {
>       cake.state = "iced"
>       iced <- cake // icer never touches this cake again
>   }
> }
> ```

// TODO: Exercise 9.1

> bindary semaphore - channel of capacity 1

### bank2

### bank3

> variables guarded by a mutex are declared immediately after the declaration of the mutex itself

> critical section - region of code between Lock and Unlock in which a goroutine is free to read and modify the shared variables

> monitor - arrangement of functions, mutex lock, and variables

// TODO: Exercise 9.2

> race detector - add -race flag to go build, go run, or go test

> The Go Memory Model document - https://go.dev/ref/mem

### memo1

// TODO: Testing

### memo2

### memo3

### memo4

### memo5

// TODO: Exercise 9.3

// TODO: Exercise 9.4

// TODO: Exercise 9.5

// TODO: Exercise 9.6
