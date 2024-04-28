# Packages and the Go Tool

### jpeg

``` sh
go run ../../ch3/mandelbrot/mandelbrot.go | go run jpeg.go > mandelbrot.jpg
```

// TODO: Exercise 10.1 and 10.2

// TODO: Exercise 10.3

### cross

``` sh
go run cross.go
GOARCH=386 go run cross.go
```

> build tags are special comments that give more fine-grained control

> doc.go is a file used for longer package comments

``` sh
go list ...xml...
go list -f '{{join .Deps " "}}' strconv
go list -f '{{.ImportPath}} -> {{join .Imports " "}}' compress/...
```

// TODO: Exercise 10.4
