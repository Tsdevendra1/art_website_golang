package general

import (
	"artWebsite/general/handlers"
	"artWebsite/general/types"
)

var Routes = types.RoutesArray{
	types.Route{
		"Index",
		"GET",
		"/users/create",
		general.UserCreate,
	},
}
