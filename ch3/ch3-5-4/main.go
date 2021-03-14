// $ go run ch3/ch3-5-4/main.go
package main

import (
	"fmt"
	"os"
)

func main() {
	{
		i := "a/b/c.go"
		fmt.Fprintf(os.Stdout, "%s\t->\t%s\n", i, basename(i))
	}
	{
		i := "c.d.go"
		fmt.Fprintf(os.Stdout, "%s\t->\t%s\n", i, basename(i))
	}
	{
		i := "abc"
		fmt.Fprintf(os.Stdout, "%s\t->\t%s\n", i, basename(i))
	}
}

func basename(s string) string {
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '/' {
			s = s[i+1:]
			break
		}
	}

	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			s = s[i+1:]
			break
		}
	}
	return s
}
