package art

import (
	"artWebsite/art/handlers"
	"artWebsite/general/types"
)

var ArtRoutes = types.RoutesArray{
	types.Route{
		"Index",
		"GET",
		"/",
		art.Index,
	},
	types.Route{
		"TodoIndex",
		"GET",
		"/todos",
		art.TodoIndex,
	},
	types.Route{
		"TodoShow",
		"GET",
		"/todos/{todoId}",
		art.TodoShow,
	},
}
