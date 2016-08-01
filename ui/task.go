// Copyright 2016 The Solvist Author(s). All rights reserved.

package ui

import (
	"net/http"

	"labix.org/v2/mgo/bson"

	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
	"github.com/nuronialBlock/solvist/solvist/data"
)

// ServeTaskNewForm serves new task form requested by r.
func ServeTaskNewForm(w http.ResponseWriter, r *http.Request) {
	acc, ok := context.Get(r, "account").(*data.Account)
	if !ok {
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}
	err := TplTaskNewForm.Execute(w, TplTaskNewFormValues{
		Common: TplCommonValues{
			Account: acc,
		},
	})
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
	acc, ok := context.Get(r, "account").(*data.Account)
	if !ok {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	err := r.ParseForm()
	if err != nil {
		ServeInternalServerError(w, r)
		return
	}

	taskValues := TaskValues{}
	decoder := schema.NewDecoder()
	err = decoder.Decode(&taskValues, r.PostForm)
	if err != nil {
		ServeInternalServerError(w, r)
		return
	}

	task := data.Task{}
	task.ProblemOJ = taskValues.ProblemOJ
	task.ProblemID = taskValues.ProblemID
	task.ProblemName = taskValues.ProblemName
	task.ProblemURL = taskValues.ProblemURL
	task.AccountID = acc.ID

	note := data.Note{}
	note.ProblemOJ = taskValues.ProblemOJ
	note.ProblemID = taskValues.ProblemID
	note.ProblemName = taskValues.ProblemName
	note.ProblemURL = taskValues.ProblemURL
	note.AccountID = acc.ID

	err = note.Put()
	if err != nil {
		ServeInternalServerError(w, r)
		return
	}
	task.NoteID = note.ID

	err = task.Put()
	if err != nil {
		ServeInternalServerError(w, r)
		return
	}

	http.Redirect(w, r, "/tasks", http.StatusSeeOther)
}

// ServeTasksList serves the list of the tasks.
func ServeTasksList(w http.ResponseWriter, r *http.Request) {
	acc, ok := context.Get(r, "account").(*data.Account)
	if !ok {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	tasks, err := data.ListTasksByAccountID(acc.ID)
	if err != nil {
		ServeInternalServerError(w, r)
		return
	}
	err = TplTasksList.Execute(w, TplTasksListValues{
		Common: TplCommonValues{
			Account: acc,
		},
		Tasks: tasks,
	})
	if err != nil {
		ServeInternalServerError(w, r)
		return
	}
}

// HandleTaskRemove removes a task from the task list.
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

	acc, ok := context.Get(r, "account").(*data.Account)
	if !ok {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	if acc.ID != task.AccountID {
		ServeBadRequest(w, r)
		return
	}

	note, err := data.GetNote(task.NoteID)
	if err != nil {
		ServeInternalServerError(w, r)
		return
	}
	if note == nil {
		ServeInternalServerError(w, r)
		return
	}
	if note.ModifiedAt == note.CreatedAt {
		note.Remove()
		if err != nil {
			ServeInternalServerError(w, r)
			return
		}
	}

	err = task.Remove()
	if err != nil {
		ServeInternalServerError(w, r)
		return
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
	acc, ok := context.Get(r, "account").(*data.Account)
	if !ok {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	if acc.ID != task.AccountID {
		ServeBadRequest(w, r)
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

// EditedTaskValues stores the edited values of a task.
type EditedTaskValues struct {
	ProblemName string `schema:"name"`
	ProblemOJ   string `schema:"oj"`
	ProblemID   string `schema:"id"`
	ProblemURL  string `schema:"url"`
}

// HandleTaskSave restores the edited values of tasks
// to tasks database.
func HandleTaskSave(w http.ResponseWriter, r *http.Request) {
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

	err = r.ParseForm()
	if err != nil {
		ServeInternalServerError(w, r)
		return
	}

	acc, ok := context.Get(r, "account").(*data.Account)
	if !ok {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	if acc.ID != task.AccountID {
		ServeBadRequest(w, r)
		return
	}

	editValues := EditedTaskValues{}
	decoder := schema.NewDecoder()
	err = decoder.Decode(&editValues, r.PostForm)
	if err != nil {
		ServeInternalServerError(w, r)
		return
	}

	task.ProblemName = editValues.ProblemName
	task.ProblemID = editValues.ProblemID
	task.ProblemOJ = editValues.ProblemOJ
	task.ProblemURL = editValues.ProblemURL

	err = task.Put()
	if err != nil {
		ServeInternalServerError(w, r)
		return
	}

	http.Redirect(w, r, "/tasks", http.StatusSeeOther)
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
		Methods("GET").
		Path("/tasks/edit/{id}").
		HandlerFunc(ServeTaskEditForm)
	Router.NewRoute().
		Methods("POST").
		Path("/tasks/edit/{id}").
		HandlerFunc(HandleTaskSave)
}
