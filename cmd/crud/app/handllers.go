package app

import (
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"
)

func (receiver *server) handleIndex() func(writer http.ResponseWriter, request *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		log.Print(request.URL)
		_, err := writer.Write([]byte("ok"))
		if err != nil {
			log.Print(err)
		}
	}
}


func (receiver *server) handleFavicon() func(http.ResponseWriter, *http.Request) {
	bytes, err := ioutil.ReadFile(filepath.Join(receiver.assetsPath, "favicon.ico"))
	if err != nil {
		panic(err)
	}

	return func(writer http.ResponseWriter, request *http.Request) {
		_, err := writer.Write(bytes)
		if err != nil {
			log.Print(err)
		}
	}
}