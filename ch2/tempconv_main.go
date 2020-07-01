package main

import (
	"fmt"

	tempconv "github.com/kotaoue/the-go-programming-language/ch2/tempconv"
)

func main() {
	fmt.Printf("%g\n", tempconv.BoilingC-tempconv.FreezingC)
	BoilingF := tempconv.CToF(tempconv.BoilingC)
	fmt.Printf("%g\n", BoilingF-tempconv.CToF(tempconv.BoilingC))
	// fmt.Printf("%g\n", BoilingF-tempconv.FreezingC)

	var c tempconv.Celsius
	var f tempconv.Fahrenheit

	fmt.Println(c == 0)
	fmt.Println(f >= 0)
	// fmt.Println(c == f)
	fmt.Println(c == tempconv.Celsius(f))
}
