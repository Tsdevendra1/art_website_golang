package general

import "artWebsite/general/types"

var Routes = types.RoutesArray{
	types.Route{
		"Index",
		"GET",
		"/",
		art.Index,
	},
}
