package searcher

import (
	"encoding/json"
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
	SResults     *SearchResults
	gqrySr       g.SearchResult
	sRawResults  map[string]interface{}
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
		gsc.Query += "+ in zip code (" + strings.Join(zcl, ",") + ")"
	}
	gsc.searchParms = make(map[string]string)
	gsc.searchParms["q"] = gsc.Query
	gsc.searchParms["location"] = gsc.sParms.State.Name
	return
}

func (gsc *GooglesearchClient) SaveResults() (rval error) {
	fmt.Printf("Saving results\n")
	var qryid uint
	qryid, rval = gsc.sParms.Dbcref.SaveQueryData(gsc.sParms.State.ID, gsc.sParms.KeywordList, gsc.sParms.ZipCodeList,
		gsc.sParms.AreaCodeList, gsc.Query)
	if rval != nil {
		fmt.Printf("Blew chowafter SaveQueryData: %s\n", rval.Error())
		return
	}

	for k, rslt := range gsc.SResults.Results {
		rbyte, _ := json.Marshal(rslt)
		rval = gsc.sParms.Dbcref.ProcessQry_results(qryid, 0, uint(k), rbyte)
		if rval != nil {
			continue
		}
	}
	gsc.SResults.ProcessSearchData(qryid, gsc.gqrySr)

	return
}

func (gsc *GooglesearchClient) ExecuteSearch() (rval error) {
	gsc.SResults = NewSearchResults()
	search := g.NewGoogleSearch(gsc.searchParms, system.GetSystemParams().GQKey)
	gsc.gqrySr, rval = search.GetJSON()
	if rval != nil {
		return
	}
	fmt.Printf("Search results are (gsc.gqrySr): %v\n", gsc.gqrySr)
	gsc.SResults.StoreResults(gsc.gqrySr)
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
