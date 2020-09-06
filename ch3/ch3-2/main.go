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
	var f32 float32 = 16777216
	fmt.Printf("%g\n", f32)
	fmt.Printf("%t\n", f32 == f32+1)

	var f64 float64 = 16777216
	fmt.Printf("%g\n", f64)
	fmt.Printf("%t\n", f64 == f64+1)

}
