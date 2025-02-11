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
	ZipCodes(mux)
	About(mux)
	return mux
}
