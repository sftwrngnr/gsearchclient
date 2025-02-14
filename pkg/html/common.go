package html

import (
	"fmt"
	"github.com/sftwrngnr/gsearchclient/pkg/searcher"
	. "maragu.dev/gomponents"
	hx "maragu.dev/gomponents-htmx"
	. "maragu.dev/gomponents/html"
)

func GetSearchPostReq() (rval Node) {
	rval = Button(Type("submit"), Text(`Generate search request`),
		Class("rounded-md border border-transparent bg-blue-600 px-4 py-2 text-sm font-medium text-white shadow-sm hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2"),

		hx.Include("#zc"),
		hx.Include("#ac"),
		hx.Include("#kw"),
		hx.Include("#allzc"),
		hx.Include("#allac"),
		hx.Include("#allkw"),
		hx.Include("#indivkw"),
		hx.Include("#sonly"),
		hx.Target("#qrystring"),
		hx.Post("/genqry"),
	)
	return
}

func GetQueryString(qs string, st string) (rval Node) {
	tval := []Node{Div(ID("qrystring"),
		H2(Text("Query String")),
		Text(qs)),
	}
	if st != "" {
		mySp := &searcher.SearchParms{Query: qs, Location: fmt.Sprintf("%s, United States", st)}

		qrs, qerr := mySp.Searchdata()
		if qerr != nil {
			tval = append(tval, H2(Text(fmt.Sprintf("Query Error %s", qerr))))
		}
		for i, d := range qrs {
			fmt.Printf("Line %d, %v\n", i, d)
		}
		olE := []Node{}
		for _, q := range qrs {
			olE = append(olE, Li(Text(q)))
		}
		tval = append(tval, Ol(olE...))
	}
	rval = Var(tval...)
	return
}

func GetQueryResults(qr []string) (rval Node) {
	tval := []Node{
		Var(
			H3(Text("Raw query Results")),
		),
	}
	for _, q := range qr {
		tval = append(tval, Text(q), Br())
	}
	rval = Var(tval...)
	return
}
