package html

import (
	"fmt"
	"github.com/sftwrngnr/gsearchclient/pkg/sqldb"
	"github.com/sftwrngnr/gsearchclient/pkg/system"
	. "maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
	"slices"
	"strconv"
	"strings"
)

func GenQry(mymap map[string][]string) (myOut Node, err error) {
	keys := make([]string, len(mymap))
	myst, err := system.GetSystemParams().Dbc.GetStateByAbbr(mymap["state"][0])
	if err != nil {
		return
	}
	qrystr := fmt.Sprintf("%s", myst.Name)

	i := 0
	for k := range mymap {
		keys[i] = k
		i++
	}
	fmt.Printf("Received the following query keys: %v\n", keys)
	if slices.Contains(keys, "kw") {
		qrystr += "+" + build_keywordqry(mymap["kw"])
	}
	if slices.Contains(keys, "zc") {
		qrystr += "+" + build_zipcodeqry(mymap["zc"])
	}
	if slices.Contains(keys, "ac") {
		qrystr += "+" + build_areacodeqry(mymap["ac"])
	}
	fmt.Printf("Generating a query with qry string: %v\n", qrystr)
	tOut := []Node{}
	tOut = append(tOut, GetQueryString(qrystr, myst.Name))
	myOut = html.Var(tOut...)
	return
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
