package http

import (
	"fmt"
	"github.com/sftwrngnr/gsearchclient/pkg/html"
	. "maragu.dev/gomponents"
	ghttp "maragu.dev/gomponents/http"
	"net/http"
)

func Home(mux *http.ServeMux) {
	mux.Handle("GET /", ghttp.Adapt(func(w http.ResponseWriter, r *http.Request) (Node, error) {
		// Let's pretend this comes from a db or something.
		//return html.HomePage(items), nil
		return html.HomePage(nil), nil
	}))
}

func ZipCodes(mux *http.ServeMux) {
	mux.Handle("GET /zipcodes", ghttp.Adapt(func(w http.ResponseWriter, r *http.Request) (Node, error) {
		qs := r.URL.Query().Get("state")
		fmt.Printf("Received zipcodes request for state %s\n", qs)
		return html.ZipCodes(qs), nil
	}))
}

func GenQry(mux *http.ServeMux) {
	mux.Handle("POST /genqry", ghttp.Adapt(func(w http.ResponseWriter, r *http.Request) (Node, error) {
		err := r.ParseForm()
		if err != nil {

			fmt.Printf("Error with ParseForm %s\n", err.Error())
			return nil, err
		}
		return html.GenQry(r.Form)
	}))
}

func About(mux *http.ServeMux) {
	mux.Handle("GET /about", ghttp.Adapt(func(w http.ResponseWriter, r *http.Request) (Node, error) {
		return html.AboutPage(), nil
	}))
}
