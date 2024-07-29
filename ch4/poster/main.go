package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

type Poster struct {
	Title    string
	Released string
	Runtime  string
	Plot     string
}

func main() {
	q := strings.Join(os.Args[1:], " ")
	req, _ := http.NewRequest("GET", "https://www.omdbapi.com/?apikey=af518301&t="+q, nil)
	req.Header.Set("Accept", "application/json")
	resp, _ := http.DefaultClient.Do(req)

	var poster Poster
	body, _ := io.ReadAll(resp.Body)

	json.Unmarshal(body, &poster)

	fmt.Printf("Title: %s\n", poster.Title)
	fmt.Printf("Realsed: %s\n", poster.Released)
	fmt.Printf("Runtime: %s\n", poster.Runtime)
	fmt.Printf("Plot: %s\n", poster.Plot)
}
