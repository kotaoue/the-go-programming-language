package main

import "fmt"

func main() {
	// p59
	var u uint8 = 255
	fmt.Println(u, u+1, u*u)

	var i int8 = 127
	fmt.Println(1, 1+1, i*i)

	// p60
	var x uint8 = 1 << 1
	fmt.Printf("%08b\n", x)
	x = x | 1<<5
	fmt.Printf("%08b\n", x)
}
