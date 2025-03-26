package html

import (
	"fmt"
	"github.com/sftwrngnr/gsearchclient/pkg/searcher"
	"github.com/sftwrngnr/gsearchclient/pkg/sqldb"
	"github.com/sftwrngnr/gsearchclient/pkg/system"
	. "maragu.dev/gomponents"
	"slices"
	"strconv"
	"strings"
)

func GenQry(mymap map[string][]string) (myOut Node, err error) {
	mySearchClient := getSearchClient(mymap)
	searchp := &searcher.SearchParms{}

	fmt.Printf("GenQry %v\n", mymap)
	searchp.SKeys = make([]string, len(mymap))
	err = searchp.ImportState(mymap["state"][0])
	if err != nil {
		fmt.Println(err)
		myOut = searchp.ErrorText("You must specify a state")
		err = nil
		return
	}

	//qrystr := fmt.Sprintf("%s", myst.Name)
	i := 0
	for k := range mymap {
		searchp.SKeys[i] = k
		i++
	}

	fmt.Printf("Received the following query keys: %v\n", searchp.SKeys)

	if slices.Contains(searchp.SKeys, "kw") {
		err = searchp.ImportKeywords(mymap["kw"])
		if err != nil {
			myOut = searchp.ErrorText(err.Error())
			fmt.Printf("%s\n", myOut)
			err = nil
			return
		}
	}
	if slices.Contains(searchp.SKeys, "zc") {
		err = searchp.ImportZipCodes(mymap["zc"])
		if err != nil {
			return
		}
	}
	if slices.Contains(searchp.SKeys, "ac") {
		err = searchp.ImportAreaCodes(mymap["ac"])
		if err != nil {
			return
		}
	}
	if slices.Contains(searchp.SKeys, "top10z") {
		fmt.Printf("Top 10 Zipcodes for %d\n", searchp.State)
		err = searchp.GetTop10Zips()

	}
	myOut, err = searcher.Search(searchp, mySearchClient)
	return
}

func getSearchClient(mymap map[string][]string) searcher.Searcher {
	fmt.Printf("getSearchClient %v\n", mymap)
	//mySearchClient := &searcher.GooglesearchClient{}
	return &searcher.DummySearchClient{}

}

func build_keywordqry(mykws []string) string {
	getKwd := func(kw string, kwr []sqldb.Keywords) string {
		for _, keywrdr := range kwr {
			if strconv.Itoa(int(keywrdr.ID)) == kw {
				if strings.Index(keywrdr.Keyword, " ") > -1 {
					return fmt.Sprintf("\"%s\"", keywrdr.Keyword)
				}
				return keywrdr.Keyword
			}
		}
		return ""
	}

	fmt.Println("Building keywordqry...")
	fmt.Printf("%v\n", mykws)
	var qrystring string
	// First let's see if a dentist or orthodontist. If not, insert dentist. Then... add other keywords
	keywordrecs, err := system.GetSystemParams().Dbc.GetKeywords()
	if err != nil {
		panic(err)
	}
	dstring, mydids := getDentistIdents(keywordrecs, mykws)
	qrystring = dstring
	for _, kw := range mykws {
		if kw != mydids {
			qrystring += "+" + getKwd(kw, keywordrecs)
		}
	}
	return qrystring
}

func build_zipcodeqry(zipcodes []string) string {
	return "in zip code(" + strings.Join(zipcodes, ",") + ")"
}

func build_areacodeqry(areacodes []string) string {
	return "in areacode(" + strings.Join(areacodes, ",") + ")"
}

func getDentistIdents(kwr []sqldb.Keywords, mykws []string) (rstring string, rval string) {
	dentkws := []string{"Dentist", "Orthodontist", "Dental clinic"}
	rstring = ""
	rval = ""
	getKWRec := func(inchk string) (string, string) {
		for _, krec := range kwr {
			if inchk == strconv.Itoa(int(krec.ID)) {
				if slices.Contains(dentkws, krec.Keyword) {
					return krec.Keyword, inchk
				}
			}
		}
		return "", ""
	}

	for _, kw := range mykws {
		rstring, rval = getKWRec(kw)
		if rval != "" {
			return
		}
	}
	return
}
