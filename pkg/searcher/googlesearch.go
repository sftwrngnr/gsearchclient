package searcher

import (
	"fmt"
	. "maragu.dev/gomponents"
)

type GooglesearchClient struct {
	Query        string
	Location     string
	Language     string `default:"en"`
	Country      string `default:"us"`
	SearchDomain string `default:"google.com"`
	sParms       *SearchParms
}

func (gsc *GooglesearchClient) ValidateSearchParameters(sp *SearchParms) (rval error) {
	fmt.Printf("Validating search parameters\n")
	gsc.sParms = sp
	if sp.State.ID == 0 {
		rval = fmt.Errorf("State ID must be set")
		return
	}
	if len(sp.KeywordList) == 0 {
		rval = fmt.Errorf("At least one keyword must be specified")
		return
	}
	gsc.Location = fmt.Sprintf("%s, United Statess", sp.State.Name)
	return
}

func (gsc *GooglesearchClient) BuildQuery() (rval error) {
	return
}

func (gsc *GooglesearchClient) SaveResults() (rval error) {

	return
}

func (gsc *GooglesearchClient) ExecuteSearch() (rval error) {
	myRes := NewSearchResults()
	myRes = myRes.AddResult(OrganicResultType, nil).AddResult(KnowledgeGraph, nil)
	//search := g.NewGoogleSearch(gsc.sParms, system.GetSystemParams().GQKey)
	/*
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
	return
}

func (gsc *GooglesearchClient) GetNodeResults() (rval Node) {
	return
}
