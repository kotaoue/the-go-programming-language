package main

import (
	"fmt"

	popcount "github.com/kotaoue/the-go-programming-language/ch2/popcount"
)

func main() {
	for i := 0; i < 64; i++ {
		fmt.Printf("%d %d %d\n", i, popcount.PopCount(uint64(i)), popcount.PopCountLoop(uint64(i)))
	}
}
