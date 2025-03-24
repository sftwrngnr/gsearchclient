package html

import (
	"fmt"
	"github.com/sftwrngnr/gsearchclient/pkg/system"
	. "maragu.dev/gomponents"
	"strconv"

	//hx "maragu.dev/gomponents-htmx"
	. "maragu.dev/gomponents/html"
)

func ZipCodes(qs string) (rval Node) {
	fmt.Printf("ZipCode handler called")
	var tval []Node
	var zval []Node
	var acval []Node
	var kwval []Node
	var cval []Node
	fmt.Printf("ZipCode qs: %s\n", qs)
	myZips, err := system.GetSystemParams().Dbc.GetZipsForState(qs)
	myCities, cerr := system.GetSystemParams().Dbc.GetCitiesForState(qs)
	if err != nil {
		fmt.Printf("ZipCodes: %v\n", err)
		return nil
	}
	if cerr != nil {
		fmt.Printf("Cities: %v\n", cerr)
		return nil
	}
	myAreaCodes, _ := system.GetSystemParams().Dbc.GetAreaCodesForState(qs)
	fmt.Printf("There are %d area codes for qs: %s\n", len(myAreaCodes), qs)
	myKeywords, _ := system.GetSystemParams().Dbc.GetKeywords()

	for _, ac := range myAreaCodes {
		acval = append(acval, Option(Value(ac.Code), Text(ac.Code)))
	}
	for _, z := range myZips {
		zval = append(zval, Option(Value(z.Zipcode), Text(z.Zipcode)))
	}
	for _, c := range myCities {
		cval = append(cval, Option(Value(string(c.ID)), Text(c.Name)))
	}
	fmt.Printf("There are %d cities returned for qs: %s\n", len(myCities), qs)

	fmt.Printf("There are %d keywords for qs: %s\n", len(myKeywords), qs)
	for _, k := range myKeywords {
		kwval = append(kwval, Option(Value(strconv.Itoa(int(k.ID))), Text(k.Keyword)))
	}
	tval = append(tval,
		Table(
			Table(Tr(Th(Text("Zip code")), Th(Text("City")), Th(Text("Area code")), Th(Text("Keywords"))),
				Tr(
					Td(Div(Label(Name("zipcodes")), ID("zipcodes")),
						Select(Name("zc"), ID("zc"), Multiple(), Option(zval...)),
					),
					Td(
						Select(Name("city"), ID("city"), Multiple(), Option(cval...)),
					),
					Td(Div(Label(Name("areacodes")), ID("areacodes")),
						Select(Name("ac"), ID("ac"), Multiple(), Option(acval...)),
					),
					Td(Div(Label(Name("keywords")), ID("keywords")),
						Select(Name("kw"), ID("kw"), Multiple(), Option(kwval...)),
					),
				),
				Tr(
					Td(Input(Type("checkbox"), Name("allzc"), ID("allzc")), Text("All Zipcodes")),
					Td(Input(Type("Text"), Label(Text("City"))), Text("City")),
					Td(Input(Type("checkbox"), Name("allac"), ID("allac")), Text("All Area Codes")),
					Td(Input(Type("checkbox"), Name("allkw"), ID("allkw")), Text("All Keywords"))),
				Tr(
					Td(Input(Type("checkbox"), Name("top10z"), ID("top10z")), Text("Top 10 Zipcodes")),
					Td(),
					Td(Input(Type("checkbox"), Label(Text("indivkw")), ID("indivkw")), Text("Individual Keywords"))),
			),
			H2(Text("Pagination:")),
			Table(
				Tr(
					Td(Text("Start page:"), Input(Type("text"), ID("spage"), Value("1"))),
				),
				Tr(
					Td(Text("Maximum pages:"), Input(Type("text"), ID("maxpages"), Value("10"))),
				),
				Tr(
					Td(Text("Results per page:"), Input(Type("text"), ID("resper"), Value("10"))),
				),
			),
		))

	rval = Var(tval...)
	return
}
