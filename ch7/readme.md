# Interfaces

### bytecounter

``` sh
go run bytecounter.go
```

// TODO: Exercise 7.1, 7.2 and 7.3

// TODO: Exercise 7.4 and 7.5

> declaration below asserts at compile time that a value of type *bytes.Buffer satisfies io.Writer

``` go
// *bytes.Buffer must satisfy io.Writer
var w io.Writer = new(bytes.Buffer)

// *bytes.Buffer must satisfy io.Writer
var _ io.Writer = (*bytes.Buffer)(nil)
```

### sleep

``` sh
go run sleep.go
go run sleep.go -period 50ms
go run sleep.go -period 2m30s
```

### tempconv / tempflag

``` sh
go run tempflag.go
go run tempflag.go -temp -18C
go run tempflag.go -temp 212F
go run tempflag.go -temp 273.15K
go run tempflag.go -help
```

### Exercise 7.6

``` sh
go run tempflag.go -temp 273.15K
```

### Exercise 7.7

``` md
> Explan why the help message contains oC when the default value of 20.0 does not.
Because of the String() method implementation.
```

### sorting

``` sh
go run sorting.go
```

// TODO: Exercise 7.8, 7.9 and 7.10

### http1

``` sh
go run http1.go &
go run ../../ch1/fetch/fetch.go http://localhost:8000
```

### http2

``` sh
go run http2.go &
go run ../../ch1/fetch/fetch.go http://localhost:8000/list
go run ../../ch1/fetch/fetch.go http://localhost:8000/price?item=socks
go run ../../ch1/fetch/fetch.go http://localhost:8000/price?item=shoes
go run ../../ch1/fetch/fetch.go http://localhost:8000/price?item=hat
go run ../../ch1/fetch/fetch.go http://localhost:8000/help
```

### Exercise 7.11

``` sh
go run http3.go &
go run ../../ch1/fetch/fetch.go "http://localhost:8000/create?item=shoes&price=6"
go run ../../ch1/fetch/fetch.go "http://localhost:8000/create?item=hat&price=hat"
go run ../../ch1/fetch/fetch.go "http://localhost:8000/create?item=hat&price=6"
go run ../../ch1/fetch/fetch.go http://localhost:8000/list
go run ../../ch1/fetch/fetch.go http://localhost:8000/read?item=shirt
go run ../../ch1/fetch/fetch.go http://localhost:8000/read?item=socks
go run ../../ch1/fetch/fetch.go "http://localhost:8000/update?item=shirt&price=6"
go run ../../ch1/fetch/fetch.go "http://localhost:8000/update?item=hat&price=hat"
go run ../../ch1/fetch/fetch.go "http://localhost:8000/update?item=hat&price=10"
go run ../../ch1/fetch/fetch.go http://localhost:8000/price?item=hat
go run ../../ch1/fetch/fetch.go http://localhost:8000/delete?item=shirt
go run ../../ch1/fetch/fetch.go http://localhost:8000/delete?item=socks
go run ../../ch1/fetch/fetch.go http://localhost:8000/list
```

### Exercise 7.12

``` sh
go run http3.go &
go run ../../ch1/fetch/fetch.go http://localhost:8000/list
```

### eval

``` sh
curl https://raw.githubusercontent.com/adonovan/gopl.io/master/ch7/eval/parse.go > parse.go

go test -v eval
```

### surface

``` sh
go run surface.go
```

``` txt
Open in browser:
http://localhost:8000/plot?expr=sin(-x)*pow(1.5,-r)
http://localhost:8000/plot?expr=pow(2,sin(y))*pow(2,sin(x))/12
http://localhost:8000/plot?expr=sin(x*y/10)/10
```

// TODO: Exercise 7.13, 7.14, 7.15 and 7.16

> Interface Type Assertions
> ``` go
> // writeString writes s to w.
> // Is w has a WriteString method, it is invoked instead of w.Write.
> func writeString(w io.Writer, s string) (n int, err error) {
>   type stringWriter interface {
>       WriteString(string) (n int, err error)
>   }
>   if sw, ok := w.(stringWriter); ok {
>       return w.WriteString(s) // avoid a copy
>   }
>   return w.Write([]byte(s)) // allocate temporary copy
> }
>
> func writeHeader(w io.Writer, contentType string) error {
>   if _, err := writeString(w, "Content-Type: "); err != nil {
>       return err
>   }
>   if _, err := writeString(w, contentType); err != nil {
>       return err
>   }
>   // ...
> }
> ```

### xmlselect

``` sh
go run ../../ch1/fetch/fetch.go https://www.w3.org/TR/2006/REC-xml11-20060816/ | go run xmlselect.go div div h2
```

// TODO: Exercise 7.17 and 7.18
