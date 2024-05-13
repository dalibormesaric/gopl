// Exercise 2.1: Add types, constants, and functions to tempconv for processing temperatures in
// the Kelvin scale, where zero Kelvin is -273.15oC and a difference of 1K has the same magnitude
// as 1oC.
package tempconv

import "fmt"

type Celsius float64
type Fahrenheit float64
type Kelvin float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
	FreezingK     Kelvin  = 273.15
)

func (c Celsius) String() string    { return fmt.Sprintf("%goC", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%goF", f) }
func (k Kelvin) String() string     { return fmt.Sprintf("%goK", k) }
