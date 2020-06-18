// e.g. go run dup2.go ./dup_target.txt ./dup_target2.txt
package main

import (
	"bufio"
	"fmt"
	"os"
)

type dup struct {
	Count     int
	FineNames string
}

func main() {
	counts := make(map[string]*dup)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
			}
			countLines(f, counts)
			f.Close()
		}
	}
	for line, n := range counts {
		if n.Count > 1 {
			fmt.Printf("%d\t%s\n", n.Count, line)
		}
	}
}

func countLines(f *os.File, counts map[string]*dup) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()] = &dup{
			Count:     2,
			FineNames: "aaaa",
		}
	}

	// fmt.Sprintf("%s %s", f.Name(),
}
