package examples

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func Fetch(args []string) {
	for _, url := range args[1:] {

		if len(args) < 1 {
			fmt.Fprintf(os.Stderr, "usage: fetch <url>")
			os.Exit(0)
		}

		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err.Error())
			os.Exit(1)
		}

		b, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err.Error())
			os.Exit(1)
		}

		fmt.Printf("%s", b)
	}
}
