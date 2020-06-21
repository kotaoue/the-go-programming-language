// e.g. go run exercise1_7.go http://kota.oue.me
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}

		r := strings.NewReader("some io.Reader stream to be read\n")
		_, err = io.Copy(r, resp.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s; %v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("%s", r)
	}
}
