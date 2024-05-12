# Composite Types

> arrays, slices, maps and structs

### sha256

``` sh
go run sha256.go
```

// TODO: Exercise 4.1 and 4.2

### rev

``` sh
go run rev.go
```

### append

``` sh
go run append.go
```

### nonempty

``` sh
go run nonempty.go
```

> a slice can be used to implement a stack
> ``` go
> stack = append(stack, v)     // push v
> top := stack[len(stack)-1]   // top of stack
> stack = stack[:len(stack)-1] // pop
> ```

// TODO: Exercise 4.3, 4.4, 4.5, 4.6 and 4.7

### dedup

``` sh
echo -e "asd\nqwe\nasd" | go run dedup.go
printf "asd\nqwe\nasd" | go run dedup.go
```

### charcount

``` sh
printf "asd\nqwe\nasd" | go run charcount.go
```

### graph

``` sh
go run graph.go
```

// TODO: Exercise 4.8 and 4.9

### employee - some examples from page 100

``` sh
go run employee.go
```

### treesort

``` sh
go run treesort.go
```

> empty struct is sometimes used instead of bool to emphasize that only keys are significant
> ``` go
> seen := make(map[string[]struct{}) // set of strings
> // ...
> if _, ok := seen[s]; !ok {
>     seen[s] = struct{}{}
>     // ...first time seeing s...
> }
> ```

> struct literal
> ``` go
> type Point struct { X, Y int }
> 
> p1 := Point{1, 2}
> 
> p2 := Point{X: 1} // Y is omitted and set to zero value
> ```

> if all the fields of a struct are comparable, the struct itself is comparable

### embeda (embed was giving a warning)

``` sh
go run embeda.go
```

> outer struct type gains not just the fields of the embedded type buts methods too

### movie

``` sh
go run movie.go
```

### github (just a package used in issues)

### issues

``` sh
go run issues.go repo:golang/go is:open json decoder
```

// TODO: Exercise 4.10, 4.11, 4.12 and 4.13

### issuesreport

``` sh
go run issuesreport.go repo:golang/go is:open json decoder
```

### issueshtml

``` sh
go run issueshtml.go repo:golang/go commenter:gopherbot json decoder > issues.html
go run issueshtml.go repo:golang/go 3133 10535 > issues2.html
```

### autoescape

``` sh
go run autoescape.go > autoescape.html
```

> for more information
> ``` sh
> go doc text/template
> go doc html/template
> ```

// TODO: Exercise 4.14
