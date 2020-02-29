package main

import (
	"flag"
	"github.com/FRahimov84/awesomeProject/cmd/crud/app"
	"net"
	"net/http"
	"path/filepath"
)
var (
	host = flag.String("name", "0.0.0.0", "server host")
	port = flag.String("port", "9999", "Server port")
)

func main() {
	flag.Parse()
	addr := net.JoinHostPort(*host, *port)
	start(addr)
}

func start(addr string) {
	router := http.NewServeMux()
	server := app.NewServer(router, filepath.Join("web", "templates"), filepath.Join("web", "assets"))
	server.InitRoutes()
	panic(http.ListenAndServe(addr, server))
}
