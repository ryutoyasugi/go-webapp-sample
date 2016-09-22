package router

import (
	"net/http"

	"github.com/ryutoyasugi/go-webapp-sample/handler"
)

func Router() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler.IndexHandler)
	return mux
}
