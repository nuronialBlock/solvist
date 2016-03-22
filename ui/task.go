package ui

import "net/http"

func ServeNewTaskForm(w http.ResponseWriter, r *http.Request) {
	err := TplCreateNewForm.Execute(w, TplCreateNewFormValues{})
	if err != nil {
		ServeInternalServerError(w, r)
	}
}

func init() {
	Router.NewRoute().Methods("GET").Path("/tasks/new").HandlerFunc(ServeNewTaskForm)
}
