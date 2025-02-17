package searcher

import (
	"fmt"
	g "github.com/serpapi/google-search-results-golang"
	"github.com/sftwrngnr/gsearchclient/pkg/sqldb"
	"github.com/sftwrngnr/gsearchclient/pkg/system"
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
	"slices"
	"strings"
)

type GooglesearchClient struct {
	Query        string
	Location     string
	Language     string `default:"en"`
	Country      string `default:"us"`
	SearchDomain string `default:"google.com"`
	sParms       *SearchParms
	searchParms  map[string]string
	sResults     g.SearchResult
}

func (gsc *GooglesearchClient) GetQueryStateId() uint {
	return gsc.sParms.State.ID
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
	fmt.Printf("Checking to see if at least one required keyword has been selected\n")
	rval = gsc.CheckRequiredKeywords(sp.KeywordList)
	if rval != nil {
		rval = fmt.Errorf("At least one of the following keywords needs to be selected: %s", rval)
		return
	}
	if slices.Contains(sp.SKeys, "ac") && (slices.Contains(sp.SKeys, "zc")) {
		rval = fmt.Errorf("Only zip code or area code may be selected. Not both.")
		return
	}
	gsc.Location = fmt.Sprintf("%s, United Statess", sp.State.Name)
	return
}

func (gsc *GooglesearchClient) BuildQuery() (rval error) {
	fmt.Printf("Building search query\n")
	gsc.Query = gsc.sParms.State.Name
	gsc.Query += " + " + gsc.GetFirstReqKwd()
	gsc.Query += "+" + gsc.GetAddtlKwds()
	if slices.Contains(gsc.sParms.SKeys, "ac") {
		gsc.Query += "+ in area code (" + strings.Join(gsc.sParms.AreaCodeList, ",") + ")"
	}
	if slices.Contains(gsc.sParms.SKeys, "zc") {
		zcl := make([]string, 0)
		for _, v := range gsc.sParms.ZipCodeList {
			zcl = append(zcl, v.Zipcode)
		}
		gsc.Query += "+ in area code (" + strings.Join(zcl, ",") + ")"
	}
	gsc.searchParms = make(map[string]string)
	gsc.searchParms["q"] = gsc.Query
	gsc.searchParms["location"] = gsc.sParms.State.Name
	return
}

func (gsc *GooglesearchClient) SaveResults() (rval error) {
	fmt.Printf("Saving results\n")
	rval = gsc.sParms.Dbcref.Gsearch_SaveQueryData(gsc.sParms.State.ID, gsc.sParms.KeywordList, gsc.sParms.ZipCodeList,
		gsc.sParms.AreaCodeList, gsc.Query)
	return
}

func (gsc *GooglesearchClient) ExecuteSearch() (rval error) {
	myRes := NewSearchResults()
	myRes.GetResults()
	/*
		search := g.NewGoogleSearch(gsc.searchParms, system.GetSystemParams().GQKey)
		gsc.sResults, rval = search.GetJSON()
		if rval != nil {
			return
		}
		fmt.Printf("Search results are: %v\n", gsc.sResults)
		myRes.StoreResults(gsc.sResults)

	*/
	//myRes.ProcessSearchData(gsc.sResults)

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
	fmt.Printf("Getting node results\n")
	rval = Div(ID("qrystring"),
		H2(Text("Query String")),
		Text(gsc.Query))
	return
}

func (gsc *GooglesearchClient) CheckRequiredKeywords(kwds []sqldb.Keywords) (rval error) {
	reqlist, err := system.GetSystemParams().Dbc.GetReqKeywords()
	if err != nil {
		return err
	}
	for _, kwd := range kwds {
		if kwd.Req {
			return
		}
	}
	rval = fmt.Errorf("Keywords selected must include one of the following: [%s]", strings.Join(reqlist, ", "))
	return
}

func (gsc *GooglesearchClient) GetFirstReqKwd() string {
	for _, kwd := range gsc.sParms.KeywordList {
		if kwd.Req {
			return fmt.Sprintf("\"%s\"", kwd.Keyword)
		}
	}
	return ""
}

func (gsc *GooglesearchClient) GetAddtlKwds() (rval string) {
	for _, kwd := range gsc.sParms.KeywordList {
		if !kwd.Req {
			if len(rval) > 0 {
				rval += " + "
			}
			rval += fmt.Sprintf("\"%s\"", kwd.Keyword)
		}
	}
	return
}
