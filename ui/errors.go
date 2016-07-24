// Copyright 2016 The Solvist Author(s). All rights reserved.

// Package ui handles the user interface.
package ui

import "net/http"

// ServeInternalServerError writes Internal server error requested by r.
func ServeInternalServerError(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Internal Server Error", http.StatusInternalServerError)
}

// ServeNotFound writes Not Found error requested by r.
func ServeNotFound(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Not Found", http.StatusNotFound)
}

// ServeBadRequest writes serve bad request by r.
func ServeBadRequest(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Bad Request", http.StatusBadRequest)
}

// ServeHandleIncorrect writes serve bad request by r.
func ServeHandleIncorrect(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Handle Incorrect", http.StatusNotFound)
}
