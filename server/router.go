package server

import (
	"log"
	"net/http"
	"strings"
)

type routeHandler func(string, http.ResponseWriter, *http.Request)

type route struct {
	Method  string
	Prefix  string
	Handler routeHandler
}

type router struct {
	Routes []*route
}

func newRouter() *router {
	return &router{
		Routes: make([]*route, 0),
	}
}

func (r *router) Add(method, prefix string, handler routeHandler) {
	r.Routes = append(
		r.Routes,
		&route{Method: method, Prefix: prefix, Handler: handler},
	)
}

func (r *router) GET(prefix string, handler routeHandler) {
	r.Add(http.MethodGet, prefix, handler)
}

func (r *router) OPTIONS(prefix string, handler routeHandler) {
	r.Add(http.MethodOptions, prefix, handler)
}

func (r *router) ServeHTTP(rw http.ResponseWriter, req *http.Request) {

	rw.Header().Set("Server", "imgproxy")

	for _, rr := range r.Routes {
		if rr.Method == req.Method && strings.HasPrefix(req.URL.Path, rr.Prefix) {
			log.Print("Running route handler for " + req.URL.Path)
			rr.Handler(rr.Prefix, rw, req)
			return
		}
	}

	rw.WriteHeader(404)
}
