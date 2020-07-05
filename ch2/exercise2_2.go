// e.g go run ch2/exercise2_2.go 0 32
// e.g. echo -e "0\n32" | go run ch2/exercise2_2.go
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Celsius float64
type Fahrenheit float64
type Feet float64
type Meter float64

func CToF(c Celsius) Fahrenheit     { return Fahrenheit(c*9/5 + 32) }
func FToC(f Fahrenheit) Celsius     { return Celsius((f - 32)) * 5 / 9 }
func (c Celsius) String() string    { return fmt.Sprintf("%g°C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }

func FToM(f Feet) Meter        { return Meter(f / 3.2808) }
func MToF(m Meter) Feet        { return Feet(m * 3.2808) }
func (f Feet) String() string  { return fmt.Sprintf("%gft", f) }
func (m Meter) String() string { return fmt.Sprintf("%gm", m) }

func main() {
	options := os.Args[1:]

	if len(options) == 0 {
		input := bufio.NewScanner(os.Stdin)
		for input.Scan() {
			options = append(options, input.Text())
		}
	}

	for _, arg := range options {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			os.Exit(1)
		}

		f := Fahrenheit(t)
		c := Celsius(t)
		fmt.Printf("%s = %s, %s = %s\n", f, FToC(f), c, CToF(c))

		ft := Fahrenheit(t)
		m := Meter(t)
		fmt.Printf("%s = %s, %s = %s\n", ft, FToM(ft), m, MToF(m))
	}
}
