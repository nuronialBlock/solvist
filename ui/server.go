// Copyright 2016 The Solvist Author(s). All rights reserved.

package ui

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Router registers routes to be matched and dispatches a handler.
var Router = mux.NewRouter()

// Server stores router.
type Server struct {
	router *mux.Router
}

// ServeHTTP calls the HandlerFunc based on request r.
func (s Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

// NewServer returs a Server.
func NewServer() *Server {
	return &Server{
		router: Router,
	}
}
