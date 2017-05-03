package main

import (
	"net/http"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

type Server struct {
	routes Routes
}

func NewServer() *Server {
	s := &Server{
		routes: Routes{},
	}

	handlers := NewCalicoAPIHandlers()

	getPolicy := Route{
		"GetPolicy",
		"GET",
		"/policy",
		handlers.listPolicy,
	}
	s.routes = append(s.routes, getPolicy)

	getPolicies := Route{
		"GetPolicies",
		"POST",
		"/policies",
		handlers.listPolicy,
	}
	s.routes = append(s.routes, getPolicies)

	createPolicy := Route{
		"CreatePolicy",
		"PUT",
		"/policy",
		handlers.createPolicy,
	}
	s.routes = append(s.routes, createPolicy)

	updatePolicy := Route{
		"UpdatePolicy",
		"POST",
		"/policy",
		handlers.updatePolicy,
	}
	s.routes = append(s.routes, updatePolicy)

	deletePolicy := Route{
		"DeletePolicy",
		"DELETE",
		"/policy",
		handlers.deletePolicy,
	}
	s.routes = append(s.routes, deletePolicy)

	listWorkloadEndpoints := Route{
		"ListEndpoints",
		"GET",
		"/workloadEndpoints",
		handlers.listWorkloadEndpoints,
	}
	s.routes = append(s.routes, listWorkloadEndpoints)

	listNodes := Route{
		"ListNodes",
		"GET",
		"/nodes",
		handlers.listNodes,
	}
	s.routes = append(s.routes, listNodes)

	getTier := Route{
		"GetTier",
		"GET",
		"/tier",
		handlers.listTier,
	}
	s.routes = append(s.routes, getTier)

	createTier := Route{
		"CreateTier",
		"PUT",
		"/tier",
		handlers.createTier,
	}
	s.routes = append(s.routes, createTier)

	updateTier := Route{
		"UpdateTier",
		"POST",
		"/tier",
		handlers.updateTier,
	}
	s.routes = append(s.routes, updateTier)

	deleteTier := Route{
		"DeleteTier",
		"DELETE",
		"/tier",
		handlers.deleteTier,
	}
	s.routes = append(s.routes, deleteTier)

	return s
}

func (s *Server) getRoutes() Routes {
	return s.routes
}
