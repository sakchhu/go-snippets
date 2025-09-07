package main

import (
	"fmt"
	"maps"
	"os"
	"slices"

	"github.com/sakchhu/go-snippets/examples"
)

var funcMap = map[string]func(){
	"echo":        examples.Echo,
	"dupes":       examples.Dupes,
	"fetch":       examples.Fetch,
	"fetch-all":   examples.FetchAll,
	"echo-server": examples.EchoServer,
}

func main() {
	args := os.Args[1:]

	printUsageAndExit := func() {
		fmt.Fprintf(os.Stderr, "usage: paila <example> [arguments...]\n%v\n", slices.Collect(maps.Keys(funcMap)))
		os.Exit(0)
	}

	// paila <example name> [arguments...]
	if len(args) < 1 {
		printUsageAndExit()
	}

	chosen, exists := funcMap[args[0]]

	if !exists {
		printUsageAndExit()
	}

	os.Args = args
	chosen()
}
