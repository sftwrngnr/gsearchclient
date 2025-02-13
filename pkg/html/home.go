package html

import (
	"fmt"
	"github.com/sftwrngnr/gsearchclient/pkg/system"
	. "maragu.dev/gomponents"
	hx "maragu.dev/gomponents-htmx"
	. "maragu.dev/gomponents/html"
)

func HomePage(items []string) Node {
	stateOpts := getStateOptions()
	States := Select(stateOpts...)
	return page("Home",
		Head(Script(Src("https://cdn.tailwindcss.com?plugins=forms,typography")),
			Script(Src("https://unpkg.com/htmx.org"))),
		H1(Text("Market Research Crawler")),
		H2(Text("Select query options:")),
		Img(ID("spinner"), Class("htmx-indicator"), Src("https://unpkg.com/html-spinner")),
		Form(Text("State:"), Br(), States, Br(),
			Input(Type("checkbox"), Name("sonly"), ID("sonly")), Text("State Only"),
			Hr(Style("border: 5px solid blue; border-radius: 5px")),
			Div(ID("selectorupdate"),
				Table(Tr(Th(Text("Zip code")), Th(Text("Area code")), Th(Text("Keywords"))),
					Tr(
						Td(
							Select(Name("zc"), ID("zc"), Multiple()),
						),
						Td(
							Select(Name("ac"), ID("ac"), Multiple()),
						),
						Td(
							Select(Name("kw"), ID("kw"), Multiple()),
						),
					),
					Tr(
						Td(Input(Type("checkbox"), Label(Text("allzc"))), Text("All Zipcodes")),
						Td(Input(Type("checkbox"), Label(Text("allac"))), Text("All Area Codes")),
						Td(Input(Type("checkbox"), Label(Text("allkw"))), Text("All Keywords")),
					),
					Tr(Td(),
						Td(),
						Td(Input(Type("checkbox"), Label(Text("indivkw"))), Text("Individual Keywords")),
					),
				)),
			Hr(Style("border: 5px solid blue; border-radius: 5px")),
			GetSearchPostReq(),
		),
		GetQueryString(""),
		GetQueryResults([]string{""}),
	)

}

func searchreqscript() (rval []Node) {
	rval = append(rval, Script())
	return
}

func getStateOptions() []Node {
	rval := []Node{}
	myStates, err := system.GetSystemParams().Dbc.GetAllStates()
	if err != nil {
		fmt.Printf("getStateOptions: %v\n", err)
		return rval
	}
	rval = append(rval, Name("state"), hx.Get("/zipcodes"), hx.Target("#selectorupdate"), hx.Indicator("#spinner"))
	rval = append(rval, Option(Value("Name"), Text("None")))
	for _, myState := range myStates {
		rval = append(rval, Option(Value(myState.Abbrev), Text(myState.Name)))
	}
	return rval
}
