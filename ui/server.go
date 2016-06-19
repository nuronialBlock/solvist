// Copyright 2016 The Solvist Author(s). All rights reserved.

package ui

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
)

const (
	cookieLength = 20
)

// Router registers routes to be matched and dispatches a handler.
var Router = mux.NewRouter()

var store = sessions.NewCookieStore(securecookie.GenerateRandomKey(cookieLength))

// Server stores router.
type Server struct {
	router *mux.Router
}

// ServeHTTP calls the HandlerFunc based on request r.
func (s Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	sess, err := store.Get(r, "very-secret")
	if err != nil {
		ServeInternalServerError(w, r)
	}
	s.router.ServeHTTP(w, r)
}

// NewServer returs a Server.
func NewServer() *Server {
	return &Server{
		router: Router,
	}
}
