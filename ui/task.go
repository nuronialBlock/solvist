// Copyright 2016 The Solvist Author(s). All rights reserved.

package ui

import "net/http"

// ServeNewTaskForm serves new task form requested by r.
func ServeNewTaskForm(w http.ResponseWriter, r *http.Request) {
	err := TplCreateNewForm.Execute(w, TplCreateNewFormValues{})
	if err != nil {
		ServeInternalServerError(w, r)
	}
}

func init() {
	Router.NewRoute().Methods("GET").Path("/tasks/new").HandlerFunc(ServeNewTaskForm)
}
