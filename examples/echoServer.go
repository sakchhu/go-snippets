package examples

import (
	"fmt"
	"image/color"
	"log"
	"net/http"
	"strconv"
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
	http.HandleFunc("/animate", animateHandler)
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

func animateHandler(w http.ResponseWriter, r *http.Request) {
	palette := []color.Color{
		color.RGBA{0xF4, 0xDB, 0xD6, 0xFF}, // Rosewater
		color.RGBA{0xF0, 0xC6, 0xC6, 0xFF}, // Flamingo
		color.RGBA{0xF5, 0xBD, 0xE6, 0xFF}, // Pink
		color.RGBA{0xC6, 0xA0, 0xF6, 0xFF}, // Mauve
		color.RGBA{0xED, 0x87, 0x96, 0xFF}, // Red
		color.RGBA{0xEE, 0x99, 0xA0, 0xFF}, // Maroon
		color.RGBA{0xF5, 0xA9, 0x7F, 0xFF}, // Peach
		color.RGBA{0xEE, 0xD4, 0x9F, 0xFF}, // Yellow
		color.RGBA{0xA6, 0xDA, 0x95, 0xFF}, // Green
		color.RGBA{0x8B, 0xD5, 0xCA, 0xFF}, // Teal
		color.RGBA{0x91, 0xD7, 0xE3, 0xFF}, // Sky
		color.RGBA{0x7D, 0xC4, 0xE4, 0xFF}, // Sapphire
		color.RGBA{0x8A, 0xAD, 0xF4, 0xFF}, // Blue
		color.RGBA{0xB7, 0xBD, 0xF8, 0xFF}, // Lavender

		color.RGBA{0xCA, 0xD3, 0xF5, 0xFF}, // Text
		color.RGBA{0xB8, 0xC0, 0xE0, 0xFF}, // Subtext1
		color.RGBA{0xA5, 0xAD, 0xCB, 0xFF}, // Subtext0
		color.RGBA{0x93, 0x9A, 0xB7, 0xFF}, // Overlay2
		color.RGBA{0x80, 0x87, 0xA2, 0xFF}, // Overlay1
		color.RGBA{0x6E, 0x73, 0x8D, 0xFF}, // Overlay0
		color.RGBA{0x5B, 0x60, 0x78, 0xFF}, // Surface2
		color.RGBA{0x49, 0x4D, 0x64, 0xFF}, // Surface1
		color.RGBA{0x36, 0x3A, 0x4F, 0xFF}, // Surface0
		color.RGBA{0x24, 0x27, 0x3A, 0xFF}, // Base
		color.RGBA{0x1E, 0x20, 0x30, 0xFF}, // Mantle
		color.RGBA{0x18, 0x19, 0x26, 0xFF}, // Crust
	}

	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "animate: %q", err.Error())
	}

	checkAndPopulate := func(name string, or int) int {
		if r.Form.Has(name) {
			n, err := strconv.Atoi(r.Form.Get(name))

			if err != nil {
				fmt.Fprintf(w, "?%s=[%q]: %q", name, r.Form.Get(name), err.Error())
				return or
			}
			return n
		}

		return or
	}

	cycles := checkAndPopulate("cycles", 5)
	size := checkAndPopulate("size", 100)
	nframes := checkAndPopulate("nframes", 64)
	delay := checkAndPopulate("delay", 8)
	res := 0.001

	if r.Form.Has("res") {
		var err error
		res, err = strconv.ParseFloat(r.Form.Get("res"), 64)
		if err != nil {
			fmt.Fprintf(w, "?res=[%q]: %q", r.Form.Get("res"), err.Error())
		}

	}

	l, _ := NewLissajous(palette, cycles, res, size, nframes, delay)
	l.Animate(w)
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
