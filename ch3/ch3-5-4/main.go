// $ go run ch3/ch3-5-4/main.go
package main

import (
	"fmt"
	"os"
	"strings"
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
	{
		i := "a/b/c.go"
		fmt.Fprintf(os.Stdout, "%s\t->\t%s\n", i, basename2(i))
	}
	{
		i := "c.d.go"
		fmt.Fprintf(os.Stdout, "%s\t->\t%s\n", i, basename2(i))
	}
	{
		i := "abc"
		fmt.Fprintf(os.Stdout, "%s\t->\t%s\n", i, basename2(i))
	}
	{
		i := "12345"
		fmt.Fprintf(os.Stdout, "%s\t->\t%s\n", i, comma(i))
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

func basename2(s string) string {
	slash := strings.LastIndex(s, "/")
	s = s[slash+1:]
	if dot := strings.LastIndex(s, "."); dot >= 0 {
		s = s[:dot]
	}
	return s
}

func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}
