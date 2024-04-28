# Low-Level Programming

> ``` go
> import "unsafe"
> 
> fmt.Println(unsafe.Sizeof(float64(0))) // "8"
> ```

### unsafeptr

``` sh
go run unsafeptr.go
```

### equal

``` sh
go run equal.go
```

// TODO: Exercise 13.1 and 13.2

### bzip

``` sh
# https://stackoverflow.com/questions/63165283/seeing-bzlib-h-not-found-when-trying-the-bindgen-tutorial
sudo apt install libbz2-dev
# https://installati.one/install-bzip2-ubuntu-20-04/
sudo apt-get -y install bzip2
```

### bzipper

``` sh
go build bzipper.go
wc -c < /usr/share/dict/words
sha256sum < /usr/share/dict/words
./bzipper < /usr/share/dict/words | wc -c
./bzipper < /usr/share/dict/words | bunzip2 | sha256sum
```

// TODO: Exercise 13.3 and 13.4
