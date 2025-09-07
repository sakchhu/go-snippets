package examples

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

const httpPrefix = "http://"

func Fetch(args []string) {
	for _, url := range args[1:] {

		if len(args) < 1 {
			fmt.Fprintf(os.Stderr, "usage: fetch <url>")
			os.Exit(0)
		}

		resp, err := http.Get(addHttpPrefix(url))
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

// utility function, maybe separate file? eh
func addHttpPrefix(url string) string {
	if !strings.HasPrefix(url, httpPrefix) {
		return httpPrefix + url
	}

	return url
}
