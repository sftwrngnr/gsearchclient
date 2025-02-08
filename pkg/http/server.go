package http

import (
	"net/http"
	"strconv"
)

func ServerStart(host string, port int16) error {
	return http.ListenAndServe(host+":"+strconv.Itoa(int(port)), setupRoutes())

}

func setupRoutes() http.Handler {
	mux := http.NewServeMux()
	Home(mux)
	//Query(mux)
	About(mux)
	return mux
}
