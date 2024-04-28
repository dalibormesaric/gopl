# Methods

### geometry

``` sh
go run geometry.go
```

> extra parameter is called method's *receiver*

> p.Distance is called a *selector*

``` go
// An IntList is a linked list of integers.
// A nil *IntList represents the empty list.
type IntList struct {
    Value int
    Tail  *IntList
}

// Sum return s the sum of the list elements.
func (list *IntList) Sum() int {
    if list == nil {
        return 0
    }
    return list.Value + list.Tail.Sum()
}
```

### urlvalues

``` sh
go run urlvalues.go
```

### coloredpoint

``` sh
go run coloredpoint.go
```

``` go
var (
    mu sync.Mutex // guards mapping
    mapping = make(map[string]string)
)

func Lookup(key string) string {
    mu.Lock()
    v := mapping[key]
    mu.Unlock()
    return v
}
```

``` go
var cache = struct {
    sync.Mutex
    mapping map[string]string
} {
    mapping: make(map[string]string),
}

func Lookup(key string) string {
    cache.Lock()
    v := cache.mapping[key]
    cache.Unlock()
    return v
}
```

### intset

``` sh
go run intset.go
```

// TODO: Exercise 6.1, 6.2, 6.3, 6.4 and 6.5
