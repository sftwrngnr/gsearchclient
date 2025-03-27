package searcher

import (
	"encoding/json"
	"fmt"
	"github.com/sftwrngnr/gsearchclient/pkg/sqldb"
	"github.com/sftwrngnr/gsearchclient/pkg/system"
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"

	"slices"
	"strings"
)

type DummySearchClient struct {
	Query        string
	Location     string
	Language     string `default:"en"`
	Country      string `default:"us"`
	SearchDomain string `default:"google.com"`
	sParms       *SearchParms
	searchParms  map[string]string
	SResults     *SearchResults
	sRawResults  map[string]interface{}
	genericVal   GenericValidator
}

func (dsc *DummySearchClient) GetQueryStateId() uint {
	return dsc.sParms.State.ID
}

func (dsc *DummySearchClient) ValidateSearchParameters(sp *SearchParms) (rval error) {
	dsc.sParms = sp
	dsc.Location = fmt.Sprintf("%s, United States", sp.State.Name)
	return dsc.genericVal.Validate(sp)
}

func (dsc *DummySearchClient) BuildQuery(zc string) (rval error) {
	fmt.Printf("Building search query\n")
	dsc.Query = dsc.sParms.State.Name
	dsc.Query += " + " + dsc.GetFirstReqKwd()
	addkw := dsc.GetAddtlKwds()
	if addkw != "" {
		dsc.Query += "+" + addkw
	}
	if slices.Contains(dsc.sParms.SKeys, "ac") {
		dsc.Query += " + in area code (" + strings.Join(dsc.sParms.AreaCodeList, ",") + ")"
	}
	if slices.Contains(dsc.sParms.SKeys, "zc") {
		zcl := make([]string, 0)
		for _, v := range dsc.sParms.ZipCodeList {
			zcl = append(zcl, v.Zipcode)
		}
		dsc.Query += " + in zip code (" + strings.Join(zcl, ",") + ")"
	}
	if zc != "" {
		dsc.Query += " + in zip code " + zc
		fmt.Printf("Query: %s\n", dsc.Query)
	}
	dsc.searchParms = make(map[string]string)
	dsc.searchParms["q"] = dsc.Query
	dsc.searchParms["location"] = dsc.sParms.State.Name
	return
}

func (dsc *DummySearchClient) SaveResults() (rval error) {
	fmt.Printf("Saving results\n")
	var qryid uint
	qryid, rval = dsc.sParms.Dbcref.SaveQueryData(dsc.sParms.State.ID, dsc.sParms.KeywordList, dsc.sParms.ZipCodeList,
		dsc.sParms.AreaCodeList, dsc.Query)
	if rval != nil {
		fmt.Printf("Blew chow after SaveQueryData: %s\n", rval.Error())
		return
	}
	dsc.SResults.ProcessSearchData(qryid, dsc.sRawResults)

	fmt.Printf("sResults.Results len is %d\n", len(dsc.SResults.Results))
	for k, rslt := range dsc.SResults.Results {
		rbyte, _ := json.Marshal(rslt)
		rval = dsc.sParms.Dbcref.ProcessQry_results(qryid, 0, uint(k), rbyte)
		if rval != nil {
			continue
		}
	}

	return
}

func (dsc *DummySearchClient) ExecuteSearch(pgidx int) (rval error) {
	dsc.SResults = NewSearchResults()
	//search := g.NewGoogleSearch(dsc.searchParms, system.GetSystemParams().GQKey)
	dsc.sRawResults, rval = dsc.SResults.GetResults()
	if rval != nil {
		return
	}
	return
}

func (dsc *DummySearchClient) GetNodeResults() (rval Node) {
	fmt.Printf("Getting node results\n")
	rval = Div(ID("qrystring"),
		H2(Text("Query String")),
		Text(dsc.Query))
	return
}

func (dsc *DummySearchClient) CheckRequiredKeywords(kwds []sqldb.Keywords) (rval error) {
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

func (dsc *DummySearchClient) GetFirstReqKwd() string {
	for _, kwd := range dsc.sParms.KeywordList {
		if kwd.Req {
			return fmt.Sprintf("\"%s\"", kwd.Keyword)
		}
	}
	return ""
}

func (dsc *DummySearchClient) GetAddtlKwds() (rval string) {
	for _, kwd := range dsc.sParms.KeywordList {
		if !kwd.Req {
			if len(rval) > 0 {
				rval += " + "
			}
			rval += fmt.Sprintf("\"%s\"", kwd.Keyword)
		}
	}
	return
}
