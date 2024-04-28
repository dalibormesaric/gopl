# Goroutines and Channels

> communicating sequential processes or CSP

> shared memory multithreading

### spinner

``` sh
go run spinner.go
```

### clock1

``` sh
go run clock1.go &
nc localhost 8000
```

### netcat1

``` sh
go run ../clock1/clock1.go &
go run netcat1.go
killall clock1
```

### clock2

``` sh
go run clock2.go &
go run ../netcat1/netcat1.go
killall clock2
```

// TODO: Exercise 8.1 and 8.2

### reverb1 and netcat2

``` sh
go run ../reverb1/reverb1.go &
go run netcat2.go
Hello?
killall reverb1
```

### reverb2 and netcat2

``` sh
go run ../reverb2/reverb2.go &
go run netcat2.go
Hello?
killall reverb2
```

> channels are connections between goroutines

``` go
ch := make(chan int) // ch has type 'chan int'

ch <- x  // a send statement

x = <-ch // a receive expression in an assignment statement
<-ch     // a receive statement; result is discarded

close(ch)

ch = make(chan int)    // unbuffered channel
ch = make(chan int, 0) // unbuffered channel
ch = make(chan int, 3) // buffered channel with capacity 3
```

> unbuffered channels are sometimes called synchronous channels

### reverb2 and netcat3

``` sh
go run ../reverb2/reverb2.go &
go run netcat3.go
Hello?
killall reverb2
```

// TODO: Exercise 8.3

> pipeline is when output of one goroutine is input to another

### pipeline1

``` sh
go run pipeline1.go
```

### pipeline2

``` sh
go run pipeline2.go
```

### pipeline3

``` sh
go run pipeline3.go
```

> goroutine leak is a situation when a goroutine trys to send its response on a channel from which no goroutine will ever receive - solved by buffered channels

### thumbnail

``` sh
mkdir thumbnail
curl https://raw.githubusercontent.com/adonovan/gopl.io/master/ch8/thumbnail/thumbnail.go > thumbnail/thumbnail.go
curl https://go.dev/blog/go-brand/Go-Logo/JPG/Go-Logo_Aqua.jpg > Go-Logo_Aqua.jpg
curl https://go.dev/blog/go-brand/Go-Logo/JPG/Go-Logo_Black.jpg > Go-Logo_Black.jpg
curl https://go.dev/blog/go-brand/Go-Logo/JPG/Go-Logo_Blue.jpg > Go-Logo_Blue.jpg
curl https://go.dev/blog/go-brand/Go-Logo/JPG/Go-Logo_Fuchsia.jpg > Go-Logo_Fuchsia.jpg
curl https://go.dev/blog/go-brand/Go-Logo/JPG/Go-Logo_LightBlue.jpg > Go-Logo_LightBlue.jpg
curl https://go.dev/blog/go-brand/Go-Logo/JPG/Go-Logo_Yellow.jpg > Go-Logo_Yellow.jpg
go run thumbnail.go
```

> to remove thumbnails `ls | grep thumb.jpg | while read line; do rm "$line"; done;`

// TODO: Exercise 8.4 and 8.5

### crawl1

``` sh
go run crawl1.go http://gopl.io/
```

> limit parallelism using a buffered channel of capacity n to model a concurrency primitive called a *counting semaphore*

### crawl2

``` sh
go run crawl2.go http://gopl.io/
```

### crawl3

``` sh
go run crawl3.go http://gopl.io/
```

// TODO: Exercise 8.6 and 8.7

### countdown1

``` sh
go run countdown1.go
```

### countdown2

``` sh
go run countdown2.go
```

// TODO: Exercise 8.8

### du1

``` sh
go run du1.go $HOME /usr /bin /etc
```

### du2

``` sh
go run du2.go -v $HOME /usr /bin /etc
```

### du3

``` sh
go run du3.go -v $HOME /usr /bin /etc
```

// TODO: Exercise 8.9

### du4

``` sh
go run du4.go -v $HOME /usr /bin /etc
```

// TODO: Exercise 8.10 and 8.11

### chat

``` sh
go run chat.go &
go run ../netcat3/netcat3.go
```

// TODO: Exercise 8.12, 8.13, 8.14 and 8.15
