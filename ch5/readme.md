# Functions

``` go
func hypot(x, y float64) float64 {
    return math.Sqrt(x*x + y*y)
}
fmt.Println(hypot(3, 4)) // "5"
```

> Four ways to declare a function with two parameters and one result, all of type int:
``` go
func add(x int, y int) int { return x + y }
func sub(x, y int) (z int) { z = x - y; return }
func first(x int, _ int) int { return x }
func zero(int, int) int { return 0 }

fmt.Printf("%T\n", add) // "func(int, int) int"
fmt.Printf("%T\n", sub) // "func(int, int) int"
fmt.Printf("%T\n", first) // "func(int, int) int"
fmt.Printf("%T\n", zero) // "func(int, int) int"
```

### findlinks1

`golang.org/x/net/html`

``` sh
go run ~/learn-go/go-tpl/ch1/fetch/fetch.go https://golang.org | go run findlinks1.go
```

### outline

``` sh
go run ~/learn-go/go-tpl/ch1/fetch/fetch.go https://golang.org | go run outline.go
```

> variable-size function call stack

### Exercise 5.1

``` sh
go run ~/learn-go/go-tpl/ch1/fetch/fetch.go https://golang.org | go run findlinks1.go
```

### Exercise 5.2

``` sh
go run ~/learn-go/go-tpl/ch1/fetch/fetch.go https://golang.org | go run findlinks1.go
```

### Exercise 5.3

``` sh
go run ~/learn-go/go-tpl/ch1/fetch/fetch.go https://golang.org | go run findlinks1.go
```

### Exercise 5.4

``` sh
go run ~/learn-go/go-tpl/ch1/fetch/fetch.go https://golang.org | go run findlinks1.go
```

### findlinks2

``` sh
go run findlinks2.go https://golang.org
```

``` go
// convinient during debugging
log.Printf(findLinks(url))
// or
links, err := findLinks(url)
log.Printf(links, err)
```

> convention dictates that a final bool result indicates success

> bare return - return statement may be omitted in a function with named results

``` go
// CountWordsAndImages does an HTTP GET request for the HTML
// document url and returns the number of words and images in it.
func CountWordsAndImages(url string) (words, images int, err error) {
    resp, err := http.Get(url)
    if err != nil {
        return
    }
    doc, err := html.Parse(resp.Body)
    resp.Body.Close()
    if err != nil {
        err = fmt.Errorf("parsing HTML: %s", err)
        return
    }
    words, images = countWordsAndImages(doc)
    return
}

func countWordsAndImages(n *html.Node) (words, images int) { /* ... */ }
```

> each `return` statement is equivalent to `return words, images, err`

> bare returns are best using sparingly - reduce code duplication, but they rarely make code easier to understand

### Exercise 5.5

``` sh
go run findlinks2.go https://golang.org
```

### Exercise 5.6

``` sh
go run surface.go > surface.svg
```

> if a failure is an expected result, function returns a bool (meaning ok) or error conventionally as a last parameter

> five error-handling strategies:
> - propagate the error
> - retry the failed operation
> - print the error and stop the program gracefully
> - log the error and then continue, perhaps with reduced functionality
> - ignore an error entirely

### wait

``` sh
go run wait.go
```

``` go
package io

import "errors"

// EOF is the error returned by Read when no more input is available.
var EOF = errors.New("EOF")

...

in := bufio.NewReader(os.Stdin)
for {
    r, _ := in.ReadRune()
    if err == io.EOF {
        break // finished reading
    }
    if err != nil {
        return fmt.Errorf("read failed: %v", err)
    }
    // ...use r...
}
```

### outline2

``` sh
go run outline2.go https://golang.org
```

``` go
strings.Map(func(r rune) rune { return r + 1 }, "HAL-9000")
```

// TODO: Exercise 5.7, 5.8 and 5.9

### squares

``` sh
go run squares.go
```

### toposort

``` sh
go run toposort.go
```

### links

``` sh
go run links.go https://golang.org
```

### findlinks3

``` sh
go run findlinks3.go https://golang.org
```

// TODO: Exercise 5.10, 5.11, 5.12, 5.13 and 5.14

### sum

``` sh
go run sum.go
```

// TODO: Exercise 5.15, 5.16 and 5.17

### title1

``` sh
go run title1.go https://golang.org
```

> defer - executed in reverse order

### title2

``` sh
go run title2.go https://golang.org
```

``` go
package ioutil

func ReadFile(filename string) ([]byte, error) {
    f, err := os.Open(filename)
    if err != nil {
        return nil, err
    }
    defer f.Close()
    return ReadAll(f)
}
```

``` go
var mu sync.Mutex
var m = make(map[string]int)

func lookup(key string) int {
    mu.Lock()
    defer mu.Unlock()
    return m[key]
}
```

> deferred functions run after return statements have updated the function's result variables

### fetch

``` sh
go run fetch.go https://golang.org
```

// TODO: Exercise 5.18

### defer1

```sh
go run defer1.go
```

### title3

```sh
go run title3.go https://golang.org
```

// TODO: Exercise 5.19
