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
	http.Error(w, "Sorry! Someone has already snatched that handle or email.", http.StatusBadRequest)
}

// ServeNameShort complains about the short naming.
func ServeNameShort(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Full name needs to be at least 5 letters long.", http.StatusBadRequest)
}

// ServeHandleShort complains about short handle.
func ServeHandleShort(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Handle needs to be at least 3 letters long.", http.StatusBadRequest)
}

// ServeCountryShort complains about short country name.
func ServeCountryShort(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Country name needs to be at least 2 letters long.", http.StatusBadRequest)
}

// ServeUniversityShort complains about short university name.
func ServeUniversityShort(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "University name needs to be at least 2 letters long.", http.StatusBadRequest)
}

// ServeInvalidEmail complains about invalid email address.
func ServeInvalidEmail(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Invalid email address", http.StatusBadRequest)
}

// ServeHandlePatternNotMatch complains about user handle pattern while registering.
func ServeHandlePatternNotMatch(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Please, use 0-9 A-Z a-z _ for handle.", http.StatusBadRequest)
}
