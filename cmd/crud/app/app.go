package app

import (
	"errors"
	"net/http"
)

type server struct {
	router       http.Handler
	templatePath string
	assetsPath string
}

func (receiver *server) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	//  panic("implement me")
	receiver.router.ServeHTTP(writer, request)
}

func NewServer(router http.Handler, templatePath string, assetsPath string) *server {
	if router == nil {
		panic(errors.New("router can't be nil"))
	}
	if templatePath == "" {
		panic(errors.New("templatePath cant be empty"))
	}
	if assetsPath =="" {
		panic(errors.New("assetspath cant be empty"))
	}
	return &server{router: router, templatePath: templatePath, assetsPath: assetsPath}
}



