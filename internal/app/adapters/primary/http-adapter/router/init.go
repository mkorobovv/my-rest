package router

import (
	"net/http"

	"github.com/gorilla/mux"
)

const (
	apiV1Prefix = "/api/v1"
)

type Router struct {
	router *mux.Router
}

func New() *Router {
	return &Router{
		router: mux.NewRouter(),
	}
}

func (r *Router) Router() http.Handler {
	return r.router
}

type Route struct {
	Name    string
	Method  string
	Path    string
	Handler http.Handler
}
