package main

import (
	"fmt"

	"github.com/fasthttp/router"
	log "github.com/sirupsen/logrus"
	"github.com/valyala/fasthttp"
)

func handleRequests() {

	r := router.New()

	r.GET("/", routesRoot)

	for _, route := range Cfg.Routes {
		r.GET(fmt.Sprintf("/%s", route.Location), routesProxy)
	}

	log.Infof("Listening on %v", Cfg.Listen)

	log.Fatal(fasthttp.ListenAndServe(Cfg.Listen, r.Handler))
}

func main() {
	configLoad()
	log.Info("Started promproxy")
	handleRequests()
}
