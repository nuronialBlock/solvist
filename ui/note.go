// Copyright 2016 The Solvist Author(s). All rights reserved.

package ui

import "net/http"

// ServeNoteNewForm serves a new note form.
func ServeNoteNewForm(w http.ResponseWriter, r *http.Request) {
	err := TplNoteNewForm.Execute(w, TplNoteNewFormValues{})
	if err != nil {
		ServeInternalServerError(w, r)
	}
}

func init() {
	Router.NewRoute().
		Methods("Get").
		Path("/notes/new").
		HandlerFunc(ServeNoteNewForm)
}
