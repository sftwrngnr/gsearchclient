package html

import (
	. "maragu.dev/gomponents"
	hx "maragu.dev/gomponents-htmx"
	. "maragu.dev/gomponents/html"
)

func GetSearchPostReq() (rval Node) {
	rval = Button(Type("submit"), Text(`Generate search request`),
		Class("rounded-md border border-transparent bg-blue-600 px-4 py-2 text-sm font-medium text-white shadow-sm hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2"),
		//hx.Include("zc"),
		//hx.Include("ac"),
		hx.Include("kw"),
		//hx.Include("allzc"),
		//hx.Include("allac"),
		//hx.Include("allkw"),
		//hx.Include("indivkw"),
		hx.Post("/genqry"),
	)
	return
}
