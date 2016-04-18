// Copyright 2016 The Solvist Author(s). All rights reserved.

package ui

import "net/http"

// ServeTaskNewForm serves new task form requested by r.
func ServeTaskNewForm(w http.ResponseWriter, r *http.Request) {
	err := TplCreateNewForm.Execute(w, TplCreateNewFormValues{})
	if err != nil {
		ServeInternalServerError(w, r)
	}
}

//
// // HandleTaskCreate handles new task to create from the form submission.
// func HandleTaskCreate(w http.ResponseWriter, r *http.Request) {
//
// }

func init() {
	Router.NewRoute().
		Methods("GET").
		Path("/tasks/new").
		HandlerFunc(ServeTaskNewForm)
	// Router.NewRoute().
	// 	Methods("POST").
	// 	Path("/tasks/new").
	// 	HandlerFunc(HandleTaskCreate)
}
