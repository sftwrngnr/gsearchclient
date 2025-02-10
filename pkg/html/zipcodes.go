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
	fmt.Printf("ZipCode qs: %s\n", qs)
	myZips, err := system.GetSystemParams().Dbc.GetZipsForState(qs)
	if err != nil {
		fmt.Printf("ZipCodes: %v\n", err)
		return nil
	}
	/*
		Div(ID("selectorupdate"),
		Table(
			Tr(Th(Text("Zip code")), Th(Text("Area code")), Th(Text("Keywords"))),
			Tr(
				Td(Div(Label(Name("zipcodes")), ID("zipcodes"))),
				Td(Div(Label(Name("areacodes")), ID("areacodes"))),
				Td(Div(Label(Name("keywords")), ID("keywords")))),
		)),

	*/
	myAreaCodes, _ := system.GetSystemParams().Dbc.GetAreaCodesForState(qs)
	fmt.Printf("There are %d area codes for qs: %s\n", len(myAreaCodes), qs)
	//rval = append(rval, (ID("selectorupdate")))
	for _, ac := range myAreaCodes {
		acval = append(acval, Option(Value(ac.Code), Text(ac.Code)))
	}
	for _, z := range myZips {
		zval = append(zval, Option(Value(strconv.Itoa(int(z.ID))), Text(z.Zipcode)))
	}
	rval = append(rval, Table(Tr(Th(Text("Zip code")), Th(Text("Area code")), Th(Text("Keywords"))),
		Tr(
			Td(Div(Label(Name("zipcodes")), ID("zipcodes")),
				Select(Name("zc"), ID("zc"), Multiple(), Var(zval...)),
			),
			Td(Div(Label(Name("areacodes")), ID("areacodes")),
				Select(Name("ac"), ID("ac"), Multiple(), Var(acval...)),
			),
		)))

	return Var(rval...)
}
