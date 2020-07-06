module github.com/kotaoue/the-go-programming-language

go 1.14

replace (
	github.com/kotaoue/the-go-programming-language/ch2/popcount => ./ch2/popcount
	github.com/kotaoue/the-go-programming-language/ch2/tempconv => ./ch2/tempconv
)

require (
	github.com/kotaoue/the-go-programming-language/ch2/popcount v0.0.0-00010101000000-000000000000 // indirect
	github.com/kotaoue/the-go-programming-language/ch2/tempconv v0.0.0-00010101000000-000000000000 // indirect
	gopl.io v0.0.0-20200323155855-65c318dde95e // indirect
)
