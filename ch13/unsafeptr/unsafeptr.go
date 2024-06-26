package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var x struct {
		a bool
		b int16
		c []int
	}

	// equivalent to pb := &x.b
	pb := (*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&x)) + unsafe.Offsetof(x.b)))
	*pb = 42

	fmt.Println(x.b) // "42"
}
