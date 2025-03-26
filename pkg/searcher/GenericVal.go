package searcher

import (
	"fmt"
	"github.com/sftwrngnr/gsearchclient/pkg/sqldb"
	"github.com/sftwrngnr/gsearchclient/pkg/system"
	"slices"
	"strings"
)

type GenericValidator struct {
	KeywordList []sqldb.Keywords
}

func (g GenericValidator) Validate(sp *SearchParms) (rval error) {
	fmt.Printf("Validating search parameters\n")
	fmt.Printf("Validating search parameters\n")

	if sp.State.ID == 0 {
		rval = fmt.Errorf("State ID must be set")
		return
	}
	if len(sp.KeywordList) == 0 {
		rval = fmt.Errorf("At least one keyword must be specified")
		return
	}
	fmt.Printf("Checking to see if at least one required keyword has been selected\n")

	rval = g.CheckRequiredKeywords(g.KeywordList)
	if rval != nil {
		rval = fmt.Errorf("At least one of the following keywords needs to be selected: %s", rval)
		return
	}
	if (slices.Contains(sp.SKeys, "ac") || slices.Contains(sp.SKeys, "allac")) && (slices.Contains(sp.SKeys, "allzc") || slices.Contains(sp.SKeys, "zc") ||
		slices.Contains(sp.SKeys, "top10z")) {
		rval = fmt.Errorf("Only zip code or area code may be selected. Not both.")
		return
	}
	return

}

func (g GenericValidator) CheckRequiredKeywords(kwds []sqldb.Keywords) (rval error) {
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
