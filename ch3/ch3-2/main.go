package main

import (
	"fmt"
	"math"
)

func main() {
	p63_1()
	p63_2()
}

func p63_1() {
	fmt.Printf("%g\n", math.MaxFloat32)
	fmt.Printf("%g\n", math.MaxFloat64)
}
func p63_2() {
	var f float32 = 16777216
	fmt.Printf("%g\n", f)
	fmt.Printf("%t\n", f == f+1)
}
