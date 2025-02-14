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
	results, err := search.GetJSON()
	for k, v := range results {
		//fmt.Println(k, v)
		if k == "organic_results" {
			tres := v.([]interface{})
			for _, t := range tres {
				fmt.Printf("%v\n", t)
			}
			fmt.Println("\n")
		}
	}
	// Now unmarshal json

	if err != nil {
		fmt.Println(err)
		return rval, err
	}

	return rval, err
}
