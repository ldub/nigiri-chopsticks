package router

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
)

// ProxyElectrs forwards every request to the /esplora endpoint to electrs HTTP server
func (r *Router) ProxyElectrs(res http.ResponseWriter, req *http.Request) {
	electrsURL := fmt.Sprintf("http://%s:%s", r.Config.Electrs.Host, r.Config.Electrs.Port)
	parsedURL, _ := url.Parse(electrsURL)

	req.Header.Set("X-Forwarded-Host", req.Header.Get("Host"))
	req.Host = parsedURL.Host
	req.URL.Host = parsedURL.Host
	req.URL.Scheme = parsedURL.Scheme

	proxy := httputil.NewSingleHostReverseProxy(parsedURL)
	proxy.ServeHTTP(res, req)
}
