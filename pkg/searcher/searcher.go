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

type ResultType int

const (
	OrganicResultType ResultType = iota
	KnowledgeGraph
	LocalPack
	Pagination
)

type SearchResults struct {
	RsltType ResultType
	Results  inter1face{}
}

type Searcher interface {
	BuildQueryString(SearchParms) (string, error)
	ExecuteQuery(SearchParms) ([]SearchResults, error)
}

func (sp *SearchParms) BuildSearch() string {

}

func (sp *SearchParms) Searchdata() (rval []string, err error) {
	parameter := map[string]string{
		"q":             sp.Query,
		"location":      sp.Location,
		"hl":            sp.Language,
		"gl":            sp.Country,
		"google_domain": sp.SearchDomain,
	}

	/*
		search := g.NewGoogleSearch(parameter, system.GetSystemParams().GQKey)
		data, err := search.GetJSON()
		if err != nil {
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
				fmt.Printf("\n")
				rval = append(rval, myMap["link"].(string))
			}
		}
	*/
	return rval, err
}
