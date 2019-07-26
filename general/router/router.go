package router

import (
	"artWebsite/art"
	"artWebsite/general"
	"artWebsite/general/types"
	"github.com/gorilla/mux"
	"net/http"
)

type allRoutes []types.RoutesArray

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	_allRoutes := allRoutes{art.Routes, general.Routes}

	// Loops through each app and then each route inside each app
	for _, appRoute := range _allRoutes {
		for _, route := range appRoute {
			var handler http.Handler
			handler = route.HandlerFunc
			handler = Logger(handler, route.Name)

			router.
				Methods(route.Method).
				Path(route.Pattern).
				Name(route.Name).
				Handler(handler)

		}
	}
	return router
}
