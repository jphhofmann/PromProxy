package main

import (
	"fmt"
	"net"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/valyala/fasthttp"
)

/* Display root route */
func routesRoot(ctx *fasthttp.RequestCtx) {
	fmt.Fprintf(ctx, "PromProxy service ready.")
}

/* Respond with proxied content */
func proxyRespond(ctx *fasthttp.RequestCtx) bool {
	ipAddr, _, err := net.SplitHostPort(ctx.RemoteAddr().String())
	if err != nil {
		log.Errorf("Failed to split client ip-address, %v", err)
		ctx.SetStatusCode(500)
	} else {
		for _, ip := range Cfg.Whitelist {
			if ip == ipAddr {
				path := strings.SplitAfter(string(ctx.URI().Path()), "/")[1]
				ctx.Request.SetRequestURI(Cfg.Routes[path].Path)
				var proxyClient = &fasthttp.HostClient{
					Addr:          Cfg.Routes[path].Target,
					ReadTimeout:   30 * time.Second,
					WriteTimeout:  30 * time.Second,
					DialDualStack: true,
					IsTLS:         false,
				}
				req := &ctx.Request
				resp := &ctx.Response
				req.Header.Del("Connection")
				resp.Header.SetServer("PromProxy")
				if err := proxyClient.Do(req, resp); err != nil {
					ctx.SetStatusCode(fasthttp.StatusBadGateway)
					if Cfg.Debug {
						log.Errorf("Failed to contact exporter %v on %v, %v", path, Cfg.Routes[path].Path, err)
					}
				}
				return true
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
