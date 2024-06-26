// Nonempty is an example of an in-place slice alogrithm.
package main

import "fmt"

func main() {
	data := []string{"one", "", "three"}
	fmt.Printf("%q\n", nonempty(data)) // `["one" "three"]`
	fmt.Printf("%q\n", data)           // `["one" "three" "three"]`

	data = []string{"one", "", "three"}
	data = nonempty2(data)
	fmt.Printf("%q\n", data)

	s := []int{5, 6, 6, 8, 9}
	fmt.Println(remove(s, 2))

	s = []int{5, 6, 7, 8, 9}
	fmt.Println(remove2(s, 2))
}

// nonempty returns a slice holding only the non-empty strings.
// The underlying array is modified during the call.
func nonempty(strings []string) []string {
	i := 0
	for _, s := range strings {
		if s != "" {
			strings[i] = s
			i++
		}
	}
	return strings[:i]
}

func nonempty2(strings []string) []string {
	out := strings[:0] // zero-length slice of original
	for _, s := range strings {
		if s != "" {
			out = append(out, s)
		}
	}
	return out
}

// remove removes element at position i with perserving the order.
func remove(slice []int, i int) []int {
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}

// remove2 removes element at position i without perserving the order.
func remove2(slice []int, i int) []int {
	slice[i] = slice[len(slice)-1]
	return slice[:len(slice)-1]
}
