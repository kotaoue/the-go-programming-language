package main

import (
	"fmt"

	tempconv "github.com/kotaoue/the-go-programming-language/ch2/tempconv"
)

func main() {
	fmt.Printf("%g\n", tempconv.BoilingC-tempconv.FreezingC)
	fmt.Printf("%g\n", tempconv.CToF(BoilingC))
}
