package examples

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
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

// Fetch Anime Information from an api
func JsonAnimeInfo() {
	const apiURL = "https://api.jikan.moe/v4/anime/"

	id := 1 + rand.Int()%10000 // idk how many MAL has, so limit the request id.

	type Genre struct {
		Name string `json:"name"`
	}

	type Data struct {
		MALid    int     `json:"mal_id"`
		Title    string  `json:"title"`
		Type     string  `json:"type"`
		Score    float64 `json:"score"`
		Synopsis string  `json:"synopsis"`
		Year     *int    `json:"year"`
		Genres   []Genre `json:"genres"`
	}

	type Response struct {
		Parent Data `json:"data"`
	}

	resp, _ := http.Get(apiURL + fmt.Sprint(id))
	var response Response

	b, _ := io.ReadAll(resp.Body)
	json.Unmarshal(b, &response)

	formatted, _ := json.MarshalIndent(response.Parent, "", "")
	fmt.Printf("%s\n", formatted)
}
