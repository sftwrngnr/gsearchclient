package html

import (
	"fmt"
	"github.com/sftwrngnr/gsearchclient/pkg/system"
	. "maragu.dev/gomponents"
	hx "maragu.dev/gomponents-htmx"
	. "maragu.dev/gomponents/components"
	. "maragu.dev/gomponents/html"
)

func AltHPage(items []string) Node {
	stateOpts := altgetStateOptions()
	States := Select(stateOpts...)
	return HTML5(HTML5Props{
		Title: "Alt Home",

		Head: []Node{
			Script(Src("https://cdn.tailwindcss.com?plugins=forms,typography")),
			Script(Src("https://unpkg.com/htmx.org")),
		},
		Body: []Node{
			H1(Text("Market Research Crawler")),
			P(Text("Select query options:")),
			Text("State:"), Form(States),
			Img(ID("spinner"), Class("htmx-indicator"), Src("https://unpkg.com/spinner")),
			P(Text("Zip codes")),
			Div(Label(Name("zipcodes")),
				ID("zipcodes"), Name("zipcodes")),
		},
	})

}

func altgetStateOptions() []Node {
	rval := []Node{}
	myStates, err := system.GetSystemParams().Dbc.GetAllStates()
	if err != nil {
		fmt.Printf("getStateOptions: %v\n", err)
		return rval
	}
	rval = append(rval, Name("state"), hx.Get("/zipcodes"), hx.Target("#zipcodes"), hx.Indicator("#spinner"))
	for _, myState := range myStates {
		rval = append(rval, Option(Value(myState.Abbrev), Text(myState.Name)))
	}

	return rval
}
