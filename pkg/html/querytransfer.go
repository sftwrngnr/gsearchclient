package html

import (
	. "maragu.dev/gomponents"
	hx "maragu.dev/gomponents-htmx"
	. "maragu.dev/gomponents/html"
)

func QueryTransfer(items []string) Node {
	return page("QueryTransfer",
		Head(Script(Src("https://cdn.tailwindcss.com?plugins=forms,typography")),
			Script(Src("https://unpkg.com/htmx.org"))),
		H1(Text("Market Research Crawler")),
		Form(
			H2(Text("Query Transfer")),
			QueryButton(),
			Img(ID("spinner"), Class("htmx-indicator"), Src("https://unpkg.com/html-spinner")),
			Hr(Style("border: 5px solid blue; border-radius: 5px")),
			Br(),
			Div(ID("transfer_res")),
			Hr(Style("border: 5px solid blue; border-radius: 5px")),
		),
	)

}

func QueryButton() (rval Node) {
	rval = Button(Type("submit"), Text(`Transfer`),
		Class("rounded-md border border-transparent bg-blue-600 px-4 py-2 text-sm font-medium text-white shadow-sm hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2"),
		hx.Target("#transfer_res"),
		hx.Post("/exectransfer"),
	)
	return
}
