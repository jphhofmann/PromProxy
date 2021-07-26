package main

import (
	"github.com/fasthttp/router"
	log "github.com/sirupsen/logrus"
	"github.com/valyala/fasthttp"
)

func handleRequests() {

	r := router.New()

	r.GET("/", routesRoot)

	for path, _ := range Cfg.Routes {
		r.GET("/"+path, routesProxy)
	}

	log.Infof("Listening on %v", Cfg.Listen)

	log.Fatal(fasthttp.ListenAndServe(Cfg.Listen, r.Handler))
}

func main() {
	configLoad()
	log.Info("Started PromProxy")
	handleRequests()
}
