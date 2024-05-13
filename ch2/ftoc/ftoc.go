// Ftoc prints two Fahrenheit-to-Celsius conversions.
package main

import "fmt"

func main() {
	const freezingF, boilingF = 32.0, 212.0
	fmt.Printf("%goF = %goC\n", freezingF, fToC(freezingF)) // "32oF = 0oC"
	fmt.Printf("%goF = %goC\n", boilingF, fToC(boilingF))   // "212oF = 100oC"
}

func fToC(f float64) float64 {
	return (f - 32) * 5 / 9
}
