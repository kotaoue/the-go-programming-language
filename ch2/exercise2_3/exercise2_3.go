package main

import (
	"testing"

	popcount "github.com/kotaoue/the-go-programming-language/ch2/popcount"
)

func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCount(8)
	}
}
