// Copyright 2016 The Solvist Author(s). All rights reserved.

package ui

import (
	"net/http"

	"github.com/nuronialBlock/solvist/solvist/data"
)

// ServeTaskNewForm serves new task form requested by r.
func ServeTaskNewForm(w http.ResponseWriter, r *http.Request) {
	err := TplCreateNewForm.Execute(w, TplCreateNewFormValues{})
	if err != nil {
		ServeInternalServerError(w, r)
	}
}

// HandleTaskCreate handles new task to create from the form submission.
func HandleTaskCreate(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		ServeInternalServerError(w, r)
	}
	task := data.Task{}
	task.ProblemOJ = r.FormValue("ProblemOJ")
	task.ProblemID = r.FormValue("ProblemID")

	err = task.Put()
	if err != nil {
		ServeInternalServerError(w, r)
	}

	http.Redirect(w, r, "/tasks/new", http.StatusSeeOther)
}

// ServeTasksList serves the list of the tasks.
func ServeTasksList(w http.ResponseWriter, r *http.Request) {
	tasks, err := data.ListTasks()
	if err != nil {
		ServeInternalServerError(w, r)
	}
	err = TplServeListTasks.Execute(w, TplServeListTasksValues{
		Tasks: tasks,
	})
	if err != nil {
		ServeInternalServerError(w, r)
	}
}

func init() {
	Router.NewRoute().
		Methods("GET").
		Path("/tasks/new").
		HandlerFunc(ServeTaskNewForm)
	Router.NewRoute().
		Methods("POST").
		Path("/tasks/new").
		HandlerFunc(HandleTaskCreate)
	Router.NewRoute().
		Methods("GET").
		Path("/tasks").
		HandlerFunc(ServeTasksList)
}
