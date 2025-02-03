package searcher

import (
	"fmt"
	g "github.com/serpapi/google-search-results-golang"
)

func searchdata() {
	parameter := map[string]string{
		"q":             "dentists",
		"location":      "Henderson, Nevada, United States",
		"hl":            "en",
		"gl":            "us",
		"google_domain": "google.com",
	}

	search := g.NewGoogleSearch(parameter, "mykey")
	results, err := search.GetJSON()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(results["organic_results"])

}
