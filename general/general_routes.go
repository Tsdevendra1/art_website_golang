package general

import (
	"artWebsite/general/handlers"
	"artWebsite/general/types"
)

var Routes = types.RoutesArray{
	types.Route{
		"UserCreate",
		"POST",
		"/users/create",
		general.UserCreate,
	},
	types.Route{
		"UserSelect",
		"GET",
		"/users/select",
		general.UserSelect,
	},
}
