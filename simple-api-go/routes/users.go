package routes

import (
	uc "../Domains/Users/Controllers"
)

var UserRoutes = Routes{
	Route{
		"Index",
		"GET",
		"/users",
		uc.Index,
	},
	Route{
		"GetById",
		"GET",
		"/users/{userId}",
		uc.Show,
	},
	Route{
		"Store",
		"POST",
		"/users",
		uc.Store,
	},
	Route{
		"Update",
		"PUT",
		"/users/{userId}",
		uc.Update,
	},
	Route{
		"Remove",
		"DELETE",
		"/users/{userId}",
		uc.Remove,
	},

}
