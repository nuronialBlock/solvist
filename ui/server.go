// Copyright 2016 The Solvist Author(s). All rights reserved.

package ui

import (
	"net/http"

	"labix.org/v2/mgo/bson"

	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/nuronialBlock/solvist/solvist/data"
)

const (
	cookieLength = 20
)

// Router registers routes to be matched and dispatches a handler.
var Router = mux.NewRouter()

var store = sessions.NewCookieStore([]byte("something-very-secret"))

// Server stores router.
type Server struct {
	router *mux.Router
}

// ServeHTTP calls the HandlerFunc based on request r.
func (s Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	sess, err := store.Get(r, "s")
	if err != nil {
		ServeInternalServerError(w, r)
		return
	}

	accID, ok := sess.Values["accountID"].(string)
	if ok {
		if !bson.IsObjectIdHex(accID) {
			ServeBadRequest(w, r)
			return
		}
		acc, err := data.GetAccount(bson.ObjectIdHex(accID))
		if err != nil {
			ServeInternalServerError(w, r)
			return
		}
		context.Set(r, "account", acc)
	}

	s.router.ServeHTTP(w, r)
}

// NewServer returs a Server.
func NewServer() *Server {
	return &Server{
		router: Router,
	}
}
