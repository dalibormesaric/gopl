# Tutorial

### helloworld

``` sh
go run helloworld.go
go build helloworld.go
./helloworld
```

### echo1

``` sh
go run echo1.go hello world
```

### echo2

``` sh
go run echo2.go hello world
```

### echo3

``` sh
go run echo3.go hello world
```

### Exercise 1.1

``` sh
go run echo3.go hello world
```

### Exercise 1.2

``` sh
go run echo2.go hello world
```

// TODO: Exercise 1.3

### dup1

``` sh
printf "a\nb\na" | go run dup1.go
```

### dup2

``` sh
printf "a\nb\na" | go run dup2.go
go run dup2.go tmp1.txt tmp2.txt tmp3.txt
```

### dup3

``` sh
go run dup3.go tmp1.txt tmp2.txt tmp3.txt
```

### Exercise 1.4

``` sh
go run dup2.go tmp1.txt tmp2.txt tmp3.txt
```

### lissajous

``` sh
go run lissajous.go > out.gif
```

### Exercise 1.5

``` sh
go run lissajous.go > out.gif
```

### Exercise 1.6

``` sh
go run lissajous.go > out.gif
```

### fetch

``` sh
go run fetch.go http://gopl.io
go run fetch.go http://bad.gopl.io
```

### Exercise 1.7

``` sh
go run fetch.go http://gopl.io
```

### Exercise 1.8

``` sh
go run fetch.go gopl.io
```

### Exercise 1.9

``` sh
go run fetch.go http://gopl.io
go run fetch.go http://gopl.io/notfound
```

### fetchall

``` sh
go run fetchall.go https://golang.org http://gopl.io https://godoc.org
```

// TODO: Exercise 1.10 and 1.11

### server1

``` sh
go run server1.go &
go run ../fetch/fetch.go http://localhost:8000
go run ../fetch/fetch.go http://localhost:8000/help
killall server1
```

### server2

``` sh
go run server2.go &
go run ../fetch/fetch.go http://localhost:8000/count
go run ../fetch/fetch.go http://localhost:8000
go run ../fetch/fetch.go http://localhost:8000
go run ../fetch/fetch.go http://localhost:8000/count
killall server2
```

### server3

``` sh
go run server3.go &
go run ../fetch/fetch.go http://localhost:8000/?q=query
killall server3
```

// TODO: Exercise 1.12
