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

	return s
}

func (s *Server) getRoutes() Routes {
	return s.routes
}
