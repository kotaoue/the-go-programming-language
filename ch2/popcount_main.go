package main

import (
	"fmt"

	popcount "github.com/kotaoue/the-go-programming-language/ch2/popcount"
)

func main() {
	fmt.Printf("%d\n", popcount.PopCount(1))
	fmt.Printf("%d\n", popcount.PopCount(2))
	fmt.Printf("%d\n", popcount.PopCount(3))
	fmt.Printf("%d\n", popcount.PopCount(4))
	fmt.Printf("%d\n", popcount.PopCount(5))
	fmt.Printf("%d\n", popcount.PopCount(6))
	fmt.Printf("%d\n", popcount.PopCount(7))
	fmt.Printf("%d\n", popcount.PopCount(8))
}
