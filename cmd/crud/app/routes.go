package app

import "net/http"

func (receiver *server) InitRoutes() {
	mux := receiver.router.(*http.ServeMux)
	mux.HandleFunc("/", receiver.handleIndex())
	mux.HandleFunc("/favicon.ico", receiver.handleFavicon())
}