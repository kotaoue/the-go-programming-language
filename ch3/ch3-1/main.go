package main

import "fmt"

func main() {
	p59()
	p60()
	p61_1()
}

func p59() {
	fmt.Println("--P59--")
	var u uint8 = 255
	fmt.Println(u, u+1, u*u)

	var i int8 = 127
	fmt.Println(1, 1+1, i*i)
}

func p60() {
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
	fmt.Printf("%08b\n", ^(x | y))

	for i := uint(0); i < 8; i++ {
		if x&(1<<i) != 0 {
			fmt.Println(i)
		}
	}

	fmt.Printf("%08b\n", x<<1)
	fmt.Printf("%08b\n", x>>1)
}

func p61_1() {
	fmt.Println("--P61_1--")
	medals := []string{"gold", "silver", "bronze"}
	for i := len(medals) - 1; i >= 0; i-- {
		fmt.Println(medals[i])
	}
}
