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
		P(Text("Select query options:")),
		Text("State:"), Form(States),
		Table(
			Tr(Th(Text("Zip code")), Th(Text("Area code")), Th(Text("Keywords"))),
			Tr(Td(Div(P(Text("Zip codes")), Label(Name("zipcodes")), ID("zipcodes"), Name("zipcodes"))),
				Td(Div(Label(Name("areacodes")), ID("areacodes"))),
				Td(Div(Label(Name("keywords")), ID("keywords"))))),
		Button(Type("submit"), Text(`Generate search request`),
			Class("rounded-md border border-transparent bg-blue-600 px-4 py-2 text-sm font-medium text-white shadow-sm hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2"),
		),
	)
}

func getStateOptions() []Node {
	rval := []Node{}
	myStates, err := system.GetSystemParams().Dbc.GetAllStates()
	if err != nil {
		fmt.Printf("getStateOptions: %v\n", err)
		return rval
	}
	rval = append(rval, Name("state"), hx.Get("/zipcodes"), hx.Target("#zipcodes"), hx.Indicator(".htmx-indicator"))
	rval = append(rval, Option(Value("Name"), Text("None")))
	for _, myState := range myStates {
		rval = append(rval, Option(Value(myState.Abbrev), Text(myState.Name)))
	}
	return rval
}
