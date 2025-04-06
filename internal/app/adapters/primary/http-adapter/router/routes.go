package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mkorobovv/my-rest/internal/app/adapters/primary/http-adapter/controller"
)

func (r *Router) AppendRoutes(ctr *controller.Controller) {
	apiV1Subrouter := r.router.PathPrefix(apiV1Prefix).Subrouter()

	routes := []Route{
		{
			Name:    "/orders",
			Path:    "/orders/{trackNumber}",
			Method:  http.MethodGet,
			Handler: http.HandlerFunc(ctr.Get),
		},
		{
			Name:    "/orders",
			Path:    "/orders/{trackNumber}",
			Method:  http.MethodPut,
			Handler: http.HandlerFunc(ctr.Update),
		},
	}

	r.appendRoutesToRouter(apiV1Subrouter, routes)
}

func (r *Router) appendRoutesToRouter(subrouter *mux.Router, routes []Route) {
	for _, route := range routes {
		subrouter.
			Methods(route.Method).
			Name(route.Name).
			Path(route.Path).
			Handler(route.Handler)
	}
}
