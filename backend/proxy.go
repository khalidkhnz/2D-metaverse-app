package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/gorilla/mux"
	"github.com/khalidkhnz/2D-metaverse-app/backend/lib"
)

func (s *APIServer) FileServer(folderPath string,serveUrl string, router *mux.Router) {
	router.Handle(serveUrl, http.FileServer(http.Dir(folderPath)))
}

func (s *APIServer) ProxyServer(proxyURL string, router *mux.Router) {
	url, err := url.Parse(proxyURL)
	if err != nil {
		log.Fatal(err)
	}

	proxy := httputil.NewSingleHostReverseProxy(url)

	proxy.Director = func(r *http.Request) {
		r.Host = url.Host
		r.URL.Scheme = url.Scheme
		r.URL.Host = url.Host
		r.Header.Set("X-Forwarded-Host", lib.GetAPIBase())
	}

	router.PathPrefix("/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		proxy.ServeHTTP(w, r)
	})
}
