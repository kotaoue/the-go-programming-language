module github.com/kotaoue/the-go-programming-language/ch2

go 1.14

replace github.com/kotaoue/the-go-programming-language/ch2/tempconv => ./tempconv

replace github.com/kotaoue/the-go-programming-language/ch2/poocount => ./popcount

require github.com/kotaoue/the-go-programming-language/ch2/tempconv v0.0.0-00010101000000-000000000000 // indirect

require github.com/kotaoue/the-go-programming-language/ch2/popcount v0.0.0-00010101000000-000000000000 // indirect
