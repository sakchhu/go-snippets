// Dupes prints the text of each line that appears more than once
// in the input, preceded by its count. It reads from standard input or the supplied files.
package examples

import (
	"bufio"
	"fmt"
	"os"
)

func Dupes() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) < 1 {
		input := bufio.NewScanner(os.Stdin)

		for input.Scan() {
			counts[input.Text()]++
		}

		for in, count := range counts {
			if count > 1 {
				fmt.Printf("repeated %d times: %q\n", count, in)
			}
		}
	}

	for _, arg := range files {
		f, err := os.Open(arg)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dupes: %v\n", err)
			continue
		}

		countRepeatedLines(f, counts)
		f.Close()
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("repeated %d times: %q\n", n, line)
		}
	}
}

// !NOTE: passing around file descriptors is probably not what you want to do
func countRepeatedLines(f *os.File, counts map[string]int) {
	line := bufio.NewScanner(f)
	for line.Scan() {
		counts[line.Text()]++
	}
}
