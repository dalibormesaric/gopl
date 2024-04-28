# Testing

> within *_test.go files, three kinds of functions are treated specially: tests, benchmarks and examples

### word1

``` sh
go test
go test -v
go test -v -run="French|Canal"
```

### word2

``` sh
go test
```

// TODO: Exercise 11.1 and 11.2

// TODO: Exercise 11.3 and 11.4

### echo

``` sh
go test
```

### storage1

### storage2

``` sh
go test
```

> - List of files in fmt package: go list -f={{.GoFiles}} fmt
> - External tests in fmt package: go list -f={{.TestGoFiles}} fmt
> - Files that constitnue external tests: go list -f={{.XTestGoFiles}} fmt

### Exercise 11.5

``` sh
go test
```

### eval (copied from ch7)

``` sh
go test -v -run=Coverage
go tool cover

go test -run=Coverage -coverprofile=c.out
go tool cover -html=c.out

go test -run=Coverage -coverprofile=c.out -covermode=count
go tool cover -html=c.out
```

> list of tools: ls /usr/local/go/pkg/tool/linux_amd64/

### word3 (copy of word2)

``` sh
go test -bench=.
go test -bench=. -benchmem
```

// TODO: Exercise 11.6 and 11.7

> profiling
> ``` sh
> go test -cpuprofile=cpu.out
> go test -blockprofile=block.out
> go test -memprofile=mem.out
>
> go test -run=NONE -bench=ClientServerParallel/64/h2 -cpuprofile=cpu.log net/http
> go tool pprof -text -nodecount=10 ./http.test cpu.log
> ```
> https://graphviz.org/ and "Profiling Go Programs" article on Go Blog
