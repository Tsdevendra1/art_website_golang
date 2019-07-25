package router

import (
	"github.com/gorilla/mux"
	"net/http"
	"artWebsite/art"
)



func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range art.ArtRoutes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)

	}
	return router
}
