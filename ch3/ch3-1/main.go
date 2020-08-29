package main

import "fmt"

func main() {
	fmt.Println("--P59--")
	var u uint8 = 255
	fmt.Println(u, u+1, u*u)

	var i int8 = 127
	fmt.Println(1, 1+1, i*i)

	fmt.Println("--P60--")
	var x uint8 = 1 << 1
	fmt.Printf("%08b\n", x)
	x |= 1 << 5
	fmt.Printf("%08b\n", x)

	var y uint8 = 1<<1 | 1<<2
	fmt.Printf("%08b\n", y)

	fmt.Printf("%08b\n", x&y)
	fmt.Printf("%08b\n", x|y)
	fmt.Printf("%08b\n", x^y)
	fmt.Printf("%08b\n", x&^y)
}
