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
		zval = append(zval, Option(Value(z.Zipcode), Text(z.Zipcode)))
	}
	fmt.Printf("There are %d keywords for qs: %s\n", len(myKeywords), qs)
	for _, k := range myKeywords {
		kwval = append(kwval, Option(Value(strconv.Itoa(int(k.ID))), Text(k.Keyword)))
	}
	tval = append(tval,
		Table(
			Tr(Th(Text("Zip code")), Th(Text("Area code")), Th(Text("Keywords"))),
			Tr(
				Td(Div(Label(Name("zipcodes")), ID("zipcodes")),
					Select(Name("zc"), ID("zc"), Multiple(), Option(zval...)),
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
				Td(Input(Type("checkbox"), Name("allac"), ID("allac")), Text("All Area Codes")),
				Td(Input(Type("checkbox"), Name("allkw"), ID("allkw")), Text("All Keywords"))),
			Tr(Td(),
				Td(),
				Td(Input(Type("checkbox"), Label(Text("indivkw")), ID("indivkw")), Text("Individual Keywords"))),
		),
	)

	rval = Var(tval...)
	return
}
