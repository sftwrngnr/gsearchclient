package searcher

import (
	"errors"
	"fmt"
	"github.com/sftwrngnr/gsearchclient/pkg/sqldb"
	"github.com/sftwrngnr/gsearchclient/pkg/system"
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"

	"strconv"
)

type SearchParms struct {
	Dbcref       *sqldb.DBConnData
	State        sqldb.States
	KeywordList  []sqldb.Keywords
	ZipCodeList  []sqldb.Zipcode
	TopZipList   []sqldb.Zipcode
	AreaCodeList []string
	SKeys        []string
}

type Searcher interface {
	ValidateSearchParameters(*SearchParms) error
	BuildQuery(string) error
	ExecuteSearch() error
	SaveResults() error
	GetNodeResults() Node
}

// Helpers to populate Search Parameters
func (sp *SearchParms) ImportState(s string) (rval error) {
	fmt.Printf("ImportState: %s\n", s)
	sp.Dbcref = system.GetSystemParams().Dbc
	sp.State, rval = sp.Dbcref.GetStateByAbbr(s)
	return
}

func (sp *SearchParms) ImportKeywords(kw []string) (err error) {
	var kwl []uint
	if len(kw) > 0 {
		err = errors.New("At least one keyword is required")
	}
	for _, k := range kw {
		i, cerr := strconv.Atoi(k)
		if cerr != nil {
			err = cerr
			return
		}
		kwl = append(kwl, uint(i))
	}
	fmt.Printf("ImportKeywords: %v\n", kw)
	err = sp.Dbcref.GetMatchingKeywords(kwl, &sp.KeywordList)

	if err != nil {
		fmt.Printf("error is %s", err.Error())
	}
	return
}

func (sp *SearchParms) ImportZipCodes(zc []string) (err error) {
	fmt.Printf("ImportZipCodes: %v\n", zc)
	err = sp.Dbcref.GetZipcodesForList(zc, &sp.ZipCodeList)
	return
}

func (sp *SearchParms) ImportAreaCodes(ac []string) (err error) {
	sp.AreaCodeList = make([]string, len(ac))
	for i, a := range ac {
		sp.AreaCodeList[i] = a
	}
	return
}

// Helper to set error message
func (sp *SearchParms) ErrorText(errmsg string) Node {
	return Var(Style("color: red"), Text(errmsg))
}

// Search Function
func Search(searchParms *SearchParms, searcher Searcher) (rnode Node, err error) {
	err = searcher.ValidateSearchParameters(searchParms)
	if err != nil {
		rnode = searchParms.ErrorText(fmt.Sprintf("Validation error %s", err.Error()))
		err = nil
		return
	}
	if searchParms.TopZipList != nil {
		rslts := make([]Node, len(searchParms.TopZipList))
		for i, zip := range searchParms.TopZipList {
			fmt.Printf("Build query and executing for zip code %s\n", zip.Zipcode)
			err = searcher.BuildQuery(zip.Zipcode)
			if err != nil {
				rnode = searchParms.ErrorText(fmt.Sprintf("Build query error %s", err.Error()))
				err = nil
				return
			}
			err = searcher.ExecuteSearch()
			if err != nil {
				rnode = searchParms.ErrorText(fmt.Sprintf("Execute Search error %s", err.Error()))
				err = nil
				return
			}
			err = searcher.SaveResults()
			if err != nil {
				return
			}
			rslts[i] = searcher.GetNodeResults()
		}
		rnode = Div(rslts...)
	} else {
		err = searcher.BuildQuery("")
		if err != nil {
			rnode = searchParms.ErrorText(fmt.Sprintf("Build query error %s", err.Error()))
			err = nil
			return
		}
		err = searcher.ExecuteSearch()
		if err != nil {
			rnode = searchParms.ErrorText(fmt.Sprintf("Execute Search error %s", err.Error()))
			err = nil
			return
		}
		err = searcher.SaveResults()
		if err != nil {
			return
		}
		rnode = searcher.GetNodeResults()

	}
	return
}

func (sp *SearchParms) GetTop10Zips() (err error) {
	sp.TopZipList, err = sp.Dbcref.Top10Zipcodes(sp.State.ID)
	for _, zip := range sp.TopZipList {
		fmt.Printf("%s, %d\n", zip.Zipcode, zip.Population)
	}
	return
}
