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
		H1(Text("Market Research Crawler")),
		P(Text("Select query options:")),
		Text("State:"), Form(States),
		P(Text("Zip codes")),
		Div(Label(Name("zipcodes")),
			ID("zipcodes"), Name("zipcodes")),
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
	for _, myState := range myStates {
		rval = append(rval, Option(Value(myState.Abbrev), Text(myState.Name)))
	}
	return rval
}
