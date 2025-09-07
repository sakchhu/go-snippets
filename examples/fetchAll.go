// FetchAll demonstrates goroutines(subroutines in go).
// My understanding as of now (I'll try my best to explain it below):
// We create a channel which is where all our subroutines will send their messages and communicate (hence `chan string`)
// The arguments supplied to the function are then all fetched concurrently:
// Loop over all the urls, and start the subroutines to fetch each url. If you don't know what subroutines are,
// it's just invoking a function and not blocking on it's return value! So a means to achieve concurrency basically.
// Think async functions.
// Every goroutine receives a reference to the channel.
// When a subroutine sends/tries to receive messages to the channel, it blocks execution(waits) until
// another goroutine attempts to receive/send messages (sender waits for receiver, receiver waits for sender).
// Here, FetchAll keeps waiting for the `fetch` goroutines to send their messages in the loop,
// and ends execution after all the subroutines have sent messages (either due to errors or successful execution).

package examples

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func FetchAll(args []string) {
	start := time.Now()
	ch := make(chan string) // channel is where every goroutine can communicate with each other?

	for _, url := range args[1:] {
		go fetch(url, ch) // goroutine. hmm...
	}

	for range args[1:] {
		fmt.Println(<-ch)
	}

	fmt.Printf("fetch-all: %.2fs elapsed.\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(addHttpPrefix(url))
	if err != nil {
		ch <- fmt.Sprint(err) // send the error to channel ch
		return
	}

	nbytes, err := io.Copy(io.Discard, resp.Body) // pour the response body down the sink
	resp.Body.Close()                             // no leaks, eh?
	if err != nil {
		ch <- fmt.Sprintf("fetch-all: while reading %s: %v", url, err)
		return
	}

	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}
