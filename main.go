package main

import (
	"fmt"
	g "github.com/serpapi/google-search-results-golang"
)

func main() {
	parameter := map[string]string{
		"q":             "dentists",
		"location":      "Henderson, Nevada, United States",
		"hl":            "en",
		"gl":            "us",
		"google_domain": "google.com",
	}

	search := g.NewGoogleSearch(parameter, "330b3fbe3c76655dfa738ffade25801d699dce3d45384a9a0643d89ba90c01df")
	results, err := search.GetJSON()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(results["organic_results"])
}
