// Copyright 2016 The Solvist Author(s). All rights reserved.

// Package ui handles the user interface.
package ui

import "net/http"

// ServeInternalServerError writes internal server error requested by r.
func ServeInternalServerError(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Internal Server Error", http.StatusInternalServerError)
}

// ServeNotFound writes not found error requested by r.
func ServeNotFound(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Not Found", http.StatusNotFound)
}

// ServeBadRequest writes bad request requestd by r.
func ServeBadRequest(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Bad Request", http.StatusBadRequest)
}

// ServeHandleIncorrect writes handle incorrect requested by r.
func ServeHandleIncorrect(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Handle Incorrect", http.StatusNotFound)
}

// ServeHandleOREmailDuplicate informs about the duplicacy.
func ServeHandleOREmailDuplicate(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Sorry! Someone has already snacthed that handle ", http.StatusBadRequest)
}
