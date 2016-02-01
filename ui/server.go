package ui

import (
	"net/http"

	"github.com/gorilla/mux"
)

var router = mux.NewRouter()

type Server struct {
	router *mux.Router
}

func (s Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func NewServer() *Server {
	return &Server{
		router: router,
	}
}
