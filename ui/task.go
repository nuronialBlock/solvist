// Copyright 2016 The Solvist Author(s). All rights reserved.

package ui

import (
	"net/http"

	"labix.org/v2/mgo/bson"

	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
	"github.com/nuronialBlock/solvist/solvist/data"
)

// ServeTaskNewForm serves new task form requested by r.
func ServeTaskNewForm(w http.ResponseWriter, r *http.Request) {
	err := TplTaskNewForm.Execute(w, TplTaskNewFormValues{})
	if err != nil {
		ServeInternalServerError(w, r)
	}
}

// TaskValues stores the values of the task.
type TaskValues struct {
	ProblemName string `schema:"name"`
	ProblemOJ   string `schema:"oj"`
	ProblemID   string `schema:"id"`
	ProblemURL  string `schema:"url"`
}

// HandleTaskCreate handles new task to create from the form submission.
func HandleTaskCreate(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		ServeInternalServerError(w, r)
	}

	taskValues := TaskValues{}
	decoder := schema.NewDecoder()
	err = decoder.Decode(&taskValues, r.PostForm)
	if err != nil {
		ServeInternalServerError(w, r)
	}

	task := data.Task{}
	task.ProblemOJ = taskValues.ProblemOJ
	task.ProblemID = taskValues.ProblemID
	task.ProblemName = taskValues.ProblemName
	task.ProblemURL = taskValues.ProblemURL

	err = task.Put()
	if err != nil {
		ServeInternalServerError(w, r)
	}

	http.Redirect(w, r, "/tasks", http.StatusSeeOther)
}

// ServeTasksList serves the list of the tasks.
func ServeTasksList(w http.ResponseWriter, r *http.Request) {
	tasks, err := data.ListTasks()
	if err != nil {
		ServeInternalServerError(w, r)
	}
	err = TplListTasks.Execute(w, TplListTasksValues{
		Tasks: tasks,
	})
	if err != nil {
		ServeInternalServerError(w, r)
	}
}

// HandleTaskRemove handler removes a task from the task list.
func HandleTaskRemove(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	if !bson.IsObjectIdHex(idStr) {
		ServeNotFound(w, r)
		return
	}

	id := bson.ObjectIdHex(idStr)
	task, err := data.GetTask(id)
	if err != nil {
		ServeInternalServerError(w, r)
		return
	}
	if task == nil {
		ServeNotFound(w, r)
		return
	}

	err = task.Remove()
	if err != nil {
		ServeInternalServerError(w, r)
	}

	http.Redirect(w, r, "/tasks", http.StatusSeeOther)
}

// ServeTaskEditForm serves edit page for a task.
func ServeTaskEditForm(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	if !bson.IsObjectIdHex(idStr) {
		ServeNotFound(w, r)
		return
	}

	id := bson.ObjectIdHex(idStr)
	task, err := data.GetTask(id)
	if err != nil {
		ServeInternalServerError(w, r)
		return
	}
	if task == nil {
		ServeNotFound(w, r)
		return
	}

	err = TplTaskEditForm.Execute(w, TplTaskEditFormValues{
		Task: *task,
	})
	if err != nil {
		ServeInternalServerError(w, r)
		return
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
	Router.NewRoute().
		Methods("POST").
		Path("/tasks/remove/{id}").
		HandlerFunc(HandleTaskRemove)
	Router.NewRoute().
		Methods("POST").
		Path("/tasks/edit/{id}").
		HandlerFunc(ServeTaskEditForm)
}
