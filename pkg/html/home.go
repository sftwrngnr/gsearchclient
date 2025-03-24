package html

import (
	"fmt"
	"github.com/sftwrngnr/gsearchclient/pkg/system"
	. "maragu.dev/gomponents"
	hx "maragu.dev/gomponents-htmx"
	. "maragu.dev/gomponents/html"
)

func HomePage() Node {
	stateOpts := getStateOptions()
	States := Select(stateOpts...)
	Crawlers := crawlerlist()
	return page("Home",
		Head(Script(Src("https://cdn.tailwindcss.com?plugins=forms,typography")),
			Script(Src("https://unpkg.com/htmx.org"))),
		H1(Text("Market Research Crawler")),
		H2(Text("Select query options:")),
		Form(
			Table(
				Tr(
					Td(Text("State:"), States),
					Td(Text("Crawler:"), Select(Name("Crawler"), ID("Crawler"), Option(Crawlers...))),
				),
			),
			Hr(Style("border: 5px solid blue; border-radius: 5px")),
			Div(ID("selectorupdate"),
				Table(Tr(Th(Text("Zip code")), Th(Text("City")), Th(Text("Area code")), Th(Text("Keywords"))),
					Tr(
						Td(
							Select(Name("zc"), ID("zc"), Multiple()),
						),
						Td(
							Select(Name("city"), ID("city"), Multiple()),
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
						Td(Input(Type("Text"), Label(Text("City"))), Text("City")),
						Td(Input(Type("checkbox"), Label(Text("allac"))), Text("All Area Codes")),
						Td(Input(Type("checkbox"), Label(Text("allkw"))), Text("All Keywords")),
					),
					Tr(Td(Input(Type("checkbox"), Label(Text("top10z"))), Text("Top 10 Zipcodes")),
						Td(),
						Td(Input(Type("checkbox"), Label(Text("indivkw"))), Text("Individual Keywords")),
					),
				)),
			Hr(Style("border: 5px solid blue; border-radius: 5px")),
			GetSearchPostReq(),
		),
		GetQueryString("", ""),
	)

}

func crawlerlist() []Node {
	validCrawlers := []string{"Dummy", "Google", "Delta"}
	rval := []Node{}
	rval = append(rval)
	for i, crawler := range validCrawlers {
		rval = append(rval, Option(Value(string(i)), Text(crawler)))
	}
	return rval
}

func getStateOptions() []Node {
	rval := []Node{}
	myStates, err := system.GetSystemParams().Dbc.GetAllStates()
	if err != nil {
		fmt.Printf("getStateOptions: %v\n", err)
		return rval
	}
	rval = append(rval, Name("state"), hx.Get("/zipcodes"))
	rval = append(rval, Option(Value("Name"), Text("None")))
	for _, myState := range myStates {
		rval = append(rval, Option(Value(myState.Abbrev), Text(myState.Name)))
	}
	return rval
}
