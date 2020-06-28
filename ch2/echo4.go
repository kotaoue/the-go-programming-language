// e.g. go run ch2/echo4.go --help
// e.g. go run echo4.go a b c d e f
// go run ch2/echo4.go -n a b c d e f
package main

import (
	"flag"
	"fmt"
	"strings"
)

var n = flag.Bool("n", false, "omit trailing newline")
var sep = flag.String("s", "", "separator")

func main() {
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
	}
}
