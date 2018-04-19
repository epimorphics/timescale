package main

import (
	"github.com/epimorphics/timescale/backend/auth"
	"net/http"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Token",
		"POST",
		"/token",
		auth.CreateTokenEndpoint,
	},
	Route{
		"Get Projects",
		"GET",
		"/projects",
		auth.ValidateMiddleware(GetProjects),
	},
	Route{
		"Create Project",
		"POST",
		"/projects",
		auth.ValidateMiddleware(CreateProject),
	},
	Route{
		"Update Project",
		"PUT",
		"/projects",
		auth.ValidateMiddleware(UpdateProject),
	},
	Route{
		"Delete Project",
		"DELETE",
		"/projects/{id}",
		auth.ValidateMiddleware(DeleteProject),
	},
}
