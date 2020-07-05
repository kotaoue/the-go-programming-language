// e.g go run ch2/exercise2_2.go 0 32
package main

import (
	"fmt"
	"os"
	"strconv"
)

type Celsius float64
type Fahrenheit float64

func CToF(c Celsius) Fahrenheit     { return Fahrenheit(c*9/5 + 32) }
func FToC(f Fahrenheit) Celsius     { return Celsius((f - 32)) * 5 / 9 }
func (c Celsius) String() string    { return fmt.Sprintf("%g°C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }

func main() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			os.Exit(1)
		}

		f := Fahrenheit(t)
		c := Celsius(t)
		fmt.Printf("%s = %s, %s = %s\n", f, FToC(f), c, CToF(c))
	}
}
