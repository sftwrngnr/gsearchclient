package html

import (
	"fmt"
	"github.com/sftwrngnr/gsearchclient/pkg/system"
	. "maragu.dev/gomponents"
	"strconv"

	//hx "maragu.dev/gomponents-htmx"
	. "maragu.dev/gomponents/html"
)

func ZipCodes(qs string) Node {
	fmt.Printf("ZipCode handler called")
	var rval []Node
	var zval []Node
	var acval []Node
	var kwval []Node
	fmt.Printf("ZipCode qs: %s\n", qs)
	myZips, err := system.GetSystemParams().Dbc.GetZipsForState(qs)
	if err != nil {
		fmt.Printf("ZipCodes: %v\n", err)
		return nil
	}
	myAreaCodes, _ := system.GetSystemParams().Dbc.GetAreaCodesForState(qs)
	fmt.Printf("There are %d area codes for qs: %s\n", len(myAreaCodes), qs)
	myKeywords, _ := system.GetSystemParams().Dbc.GetKeywords()

	for _, ac := range myAreaCodes {
		acval = append(acval, Option(Value(ac.Code), Text(ac.Code)))
	}
	for _, z := range myZips {
		zval = append(zval, Option(Value(strconv.Itoa(int(z.ID))), Text(z.Zipcode)))
	}
	for _, k := range myKeywords {
		kwval = append(kwval, Option(Value(k.Keyword), Text(k.Keyword)))
	}
	rval = append(rval, Table(Tr(Th(Text("Zip code")), Th(Text("Area code")), Th(Text("Keywords"))),
		Tr(
			Td(Div(Label(Name("zipcodes")), ID("zipcodes")),
				Select(Name("zc"), ID("zc"), Multiple(), Var(zval...)),
			),
			Td(Div(Label(Name("areacodes")), ID("areacodes")),
				Select(Name("ac"), ID("ac"), Multiple(), Var(acval...)),
			),
			Td(Div(Label(Name("keywords")), ID("keywords")),
				Select(Name("kw"), ID("kw"), Multiple(), Var(kwval...)),
			),
		),
		Tr(
			Td(Input(Type("checkbox"), Label(Text("allzc"))), Text("All Zipcodes")),
			Td(Input(Type("checkbox"), Label(Text("allac"))), Text("All Area Codes")),
			Td(Input(Type("checkbox"), Label(Text("allkw"))), Text("All Keywords"),
				Input(Type("checkbox"), Label(Text("indivkw"))), Text("Individual Keywords")),
		),
	))

	return Var(rval...)
}
