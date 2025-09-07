package examples

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var mu sync.Mutex
var count int

func EchoServer() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		mu.Lock()
		count++
		mu.Unlock()
		fmt.Fprintf(w, "You requested for: %s\n", r.URL.Path)
	})

	http.HandleFunc("/count", countHandler)
	http.HandleFunc("/header", headerHandler)
	log.Fatal(http.ListenAndServe(":6970", nil))
}

func headerHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}

	fmt.Fprintf(w, "Host = %q\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}

	for k, v := range r.Form {
		fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
	}
}

func countHandler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "/ was visited %d times", count)
	mu.Unlock()
}

// Example Output for http://localhost:6970/header?s=%22My%20Love%20Life%22&o=asc
// GET /header?s=%22My%20Love%20Life%22&o=asc HTTP/1.1
// Header["Connection"] = ["keep-alive"]
// Header["Upgrade-Insecure-Requests"] = ["1"]
// Header["User-Agent"] = ["Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/139.0.0.0 Safari/537.36"]
// Header["Accept"] = ["text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8"]
// Header["Accept-Language"] = ["en-US,en;q=0.9"]
// Header["Sec-Fetch-Mode"] = ["navigate"]
// Header["Sec-Fetch-User"] = ["?1"]
// Header["Sec-Fetch-Dest"] = ["document"]
// Header["Sec-Ch-Ua"] = ["\"Not;A=Brand\";v=\"99\", \"Brave\";v=\"139\", \"Chromium\";v=\"139\""]
// Header["Sec-Ch-Ua-Mobile"] = ["?0"]
// Header["Sec-Ch-Ua-Platform"] = ["\"Linux\""]
// Header["Sec-Gpc"] = ["1"]
// Header["Sec-Fetch-Site"] = ["none"]
// Header["Accept-Encoding"] = ["gzip, deflate, br, zstd"]
// Host = "localhost:6970"
// RemoteAddr = "[::1]:36588"
// Form["o"] = ["asc"]
// Form["s"] = ["\"My Love Life\""]
