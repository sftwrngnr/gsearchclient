package http

import (
	"fmt"
	"net/http"
	"strconv"
)

func ServerStart(host string, port int16) error {
	fmt.Printf("Starting server on %s:%d\n", host, port)
	return http.ListenAndServe(host+":"+strconv.Itoa(int(port)), setupRoutes())

}

func setupRoutes() http.Handler {
	mux := http.NewServeMux()
	Home(mux)
	Home2(mux)
	ZipCodes(mux)
	GenQry(mux)
	QueryTransfer(mux)
	GetCrawlers(mux)
	ExecTransfer(mux)
	CrawlerExec(mux)
	CrawlerSetup(mux)
	Crawler3Exec(mux)
	About(mux)
	return mux
}
