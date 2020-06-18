// e.g. go run dup2.go ./dup_target.txt ./dup_target2.txt
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type dup struct {
	Count     int
	FileNames string
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
			fmt.Printf("%d\t%s\t%v\n", n.Count, line, n.FileNames)
		}
	}
}

func countLines(f *os.File, counts map[string]*dup) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		c := 1
		n := f.Name()
		if _, ok := counts[input.Text()]; ok {
			c = counts[input.Text()].Count + 1

			if !strings.Contains(counts[input.Text()].FileNames, f.Name()) {
				n = counts[input.Text()].FileNames + "," + f.Name()
			}
		}
		counts[input.Text()] = &dup{
			Count:     c,
			FileNames: n,
		}
	}
}
