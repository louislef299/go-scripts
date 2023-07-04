package router

import (
	"net/http"
	"regexp"
)

type Route struct {
	Method  string
	Pattern string
	Handler http.Handler
}

type Router struct {
	routes []Route
}

func NewRouter() *Router {
	return &Router{}
}

func (r *Router) AddRoute(method, path string, handler http.Handler) {
	r.routes = append(r.routes, Route{Method: method, Pattern: path, Handler: handler})
}

// Methods that all call AddRoute:

func (r *Router) GET(path string, handler Handler) {
	r.AddRoute("GET", path, handler)
}

func (r *Router) POST(path string, handler Handler) {
	r.AddRoute("POST", path, handler)
}

func (r *Router) PUT(path string, handler Handler) {
	r.AddRoute("PUT", path, handler)
}

func (r *Router) DELETE(path string, handler Handler) {
	r.AddRoute("DELETE", path, handler)
}

func (r *Router) getHandler(method, path string) http.Handler {
	for _, route := range r.routes {
		re := regexp.MustCompile(route.Pattern)
		if route.Method == method && re.MatchString(path) {
			return route.Handler
		}
	}
	return http.NotFoundHandler()
}

// Implements the Handler interface
// https://pkg.go.dev/net/http#Handler
func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	path := req.URL.Path
	method := req.Method
	handler := r.getHandler(method, path)

	// handler middlewares go here
	handler.ServeHTTP(w, req)
}

// resource: https://medium.com/@sanjeevsiva/custom-golang-http-router-970a309531d7
