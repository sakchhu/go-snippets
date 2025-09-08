package examples

import (
	"encoding/json"
	"fmt"
	"os"
)

// Demonstrates JSON serialization/deserialization
func JsonSimple() {

	type Simple struct { // json schema
		Title  string
		Author string
		URL    string
	}

	data, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "ExampleJSON: %q\n", err.Error())
		os.Exit(1)
	}

	var info Simple
	json.Unmarshal(data, &info) // deserialized

	fmt.Printf("%+v\n", info) // unformatted

	// pretty printing as well as serializing
	formatted, _ := json.MarshalIndent(info, "", "    ") // can use `Marshal` if formatted json is not a necessity
	fmt.Printf("%s\n", formatted)
}
