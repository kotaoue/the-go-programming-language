// e.g $ pwd
// /go/src/github.com/kotaoue/the-go-programming-language/ch2/exercise2_3
// go test -bench .
package main

import (
	"testing"

	popcount "github.com/kotaoue/the-go-programming-language/ch2/popcount"
)

func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCount(i)
	}
}

func BenchmarkPopCountLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCountLoop(i)
	}
}
