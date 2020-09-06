package main

import (
	"fmt"
	"math"
)

const e = 2.71828

// Avogadro is example for number of very large.
const Avogadro = 6.02214129e23

// Planck is example for number of very small.
const Planck = 6.62606957e-34

func main() {
	p63_1()
	p63_2()
	p63_3()
	p63_4()
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

func p63_3() {
	fmt.Printf("%g\n", e)
}

func p63_4() {
	fmt.Printf("%g\n", Avogadro)
	fmt.Printf("%g\n", Planck)
	fmt.Printf("%t\n", Avogadro >= Planck)
	fmt.Printf("%t\n", Avogadro == Planck)
	fmt.Printf("%t\n", Avogadro <= Planck)
}
