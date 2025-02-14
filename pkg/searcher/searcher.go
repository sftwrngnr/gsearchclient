package searcher

import (
	"fmt"
	g "github.com/serpapi/google-search-results-golang"
	"github.com/sftwrngnr/gsearchclient/pkg/system"
)

type SearchParms struct {
	Query        string
	Location     string
	Language     string `default:"en"`
	Country      string `default:"us"`
	SearchDomain string `default:"google.com"`
}

func (sp *SearchParms) Searchdata() ([]string, error) {
	var rval []string
	parameter := map[string]string{
		"q":             sp.Query,
		"location":      sp.Location,
		"hl":            sp.Language,
		"gl":            sp.Country,
		"google_domain": sp.SearchDomain,
	}

	search := g.NewGoogleSearch(parameter, system.GetSystemParams().GQKey)
	data, err := search.GetJSON()
	if err != nil {
		fmt.Printf("Fuck chocolate shakes.\n")
		return rval, err
	}
	//fmt.Printf("%v\n", data)
	if data["organic_results"] != nil {
		results := data["organic_results"].([]interface{})
		fmt.Printf("Got organic results: \n")
		for i := range len(results) {
			if results[i] == nil {
				continue
			}
			myMap := results[i].(map[string]interface{})
			fmt.Printf("%s\n", myMap["title"].(string))
			fmt.Printf("%s\n", myMap["link"].(string))
			//fmt.Printf("%i\n", myMap["display_link"].(string))
			fmt.Printf("\n")
			rval = append(rval, myMap["link"].(string))
		}
	}

	/*
		results := data["organic_results"].([]interface{})
		for i := range len(results) {
			if results[i] != nil {
				for k, v := range results[i].(map[string]interface{}) {
					fmt.Printf("%s :: %v\n", k, v)
				}
			}
			//rval = append(rval, results[i].(map[string]interface{})["displayed_link"].(string))
		}

	*/
	return rval, err
}
