package main

import (
	"fmt"
	"net"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/valyala/fasthttp"
	proxy "github.com/yeqown/fasthttp-reverse-proxy"
)

/* Display root route */
func routesRoot(ctx *fasthttp.RequestCtx) {
	fmt.Fprintf(ctx, "promproxy service ready.")
}

/* Respond with proxied content */
func proxyRespond(ctx *fasthttp.RequestCtx) bool {
	ipAddr, _, err := net.SplitHostPort(ctx.RemoteAddr().String())
	if err != nil {
		ctx.SetStatusCode(500)
	} else {
		for _, ip := range Cfg.Whitelist {
			if ip == ipAddr {
				path := strings.SplitAfter(string(ctx.URI().Path()), "/")[1]
				config := Cfg.Routes[path]
				if len(config.Target) != 0 {
					ctx.Request.SetRequestURI(config.Path)
					if !Cfg.Debug {
						proxy.SetProduction()
					}
					proxyServer := proxy.NewReverseProxy(config.Target)
					proxyServer.ServeHTTP(ctx)
					return true
				} else {
					ctx.SetStatusCode(500)
					fmt.Fprintf(ctx, "Failed to find route")
					return true
				}
			}
		}
	}
	if Cfg.Debug {
		log.Infof("Unauthorized connection from %v", ipAddr)
	}
	return false
}

/* Proxy route */
func routesProxy(ctx *fasthttp.RequestCtx) {
	if !proxyRespond(ctx) {
		fmt.Fprintf(ctx, "Unauthorized")
		ctx.SetStatusCode(401)
	}
}
