// Boiling prints the boiling point of water.
package main

import "fmt"

const boilingF = 212.0

func main() {
	var f = boilingF
	var c = (f - 32) * 5 / 9
	fmt.Printf("boiling point = %goF or %goC\n", f, c)
	// Output:
	// boiling point = 212oF or 100oC
}
