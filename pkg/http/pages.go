package http

import (
	"fmt"
	"github.com/sftwrngnr/gsearchclient/pkg/crawler"
	"github.com/sftwrngnr/gsearchclient/pkg/html"
	. "maragu.dev/gomponents"
	html2 "maragu.dev/gomponents/html"
	ghttp "maragu.dev/gomponents/http"
	"net/http"
)

func Home2(mux *http.ServeMux) {
	mux.Handle("GET /H2", ghttp.Adapt(func(w http.ResponseWriter, r *http.Request) (Node, error) {
		return html.HomePage2(), nil
	}))

}

func Home(mux *http.ServeMux) {
	mux.Handle("GET /", ghttp.Adapt(func(w http.ResponseWriter, r *http.Request) (Node, error) {
		return html.HomePage(), nil
	}))
}

func ZipCodes(mux *http.ServeMux) {
	mux.Handle("GET /zipcodes", ghttp.Adapt(func(w http.ResponseWriter, r *http.Request) (Node, error) {
		qs := r.URL.Query().Get("state")
		fmt.Printf("Received zipcodes request for state %s\n", qs)
		return html.ZipCodes(qs), nil
	}))
}

func ExecTransfer(mux *http.ServeMux) {
	mux.Handle("POST /exectransfer", ghttp.Adapt(func(w http.ResponseWriter, r *http.Request) (Node, error) {
		turls, err := crawler.TransferURLS()
		if err != nil {
			return html2.Div(), err
		}

		var nArr []Node
		for _, turl := range turls {
			nArr = append(nArr, html2.Li(Text(turl)))
		}
		return html2.Nav(html2.Ul(nArr...)), err
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

func QueryTransfer(mux *http.ServeMux) {
	mux.Handle("GET /qrytransfer", ghttp.Adapt(func(w http.ResponseWriter, r *http.Request) (Node, error) {
		return html.QueryTransfer(nil), nil
	}))
}

func CrawlerExec(mux *http.ServeMux) {
	mux.Handle("GET /crawlerexec", ghttp.Adapt(func(w http.ResponseWriter, r *http.Request) (Node, error) {
		return nil, nil
	}))
}

func CrawlerSetup(mux *http.ServeMux) {
	mux.Handle("GET /crawlsetup", ghttp.Adapt(func(w http.ResponseWriter, r *http.Request) (Node, error) {
		return nil, nil
	}))
}

func About(mux *http.ServeMux) {
	mux.Handle("GET /about", ghttp.Adapt(func(w http.ResponseWriter, r *http.Request) (Node, error) {
		return html.AboutPage(), nil
	}))
}
