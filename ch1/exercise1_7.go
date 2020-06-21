// e.g. go run exercise1_7.go http://kota.oue.me
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		_, err := io.Copy(resp, b)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s; %v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("%s", b)
	}
}
