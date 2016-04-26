// Copyright 2016 The Solvist Author(s). All rights reserved.

package ui

import (
	"html/template"

	"github.com/nuronialBlock/solvist/solvist/data"
)

// TplCreateNewFormValues passes the new form values.
type TplCreateNewFormValues struct{}

// TplCreateNewForm generates new form for user.
var TplCreateNewForm = template.Must(template.ParseFiles("ui/templates/layout.gohtml", "ui/templates/taskNewFormPage.gohtml", "ui/templates/taskNewForm.gohtml"))

// TplServeListTasksValues stores the values of ServeListTasks template.
type TplServeListTasksValues struct {
	Tasks []data.Task
}

// TplServeListTasks renders Tasks.
var TplServeListTasks = template.Must(template.ParseFiles("ui/templates/layout.gohtml", "ui/templates/listTasksPage.gohtml", "ui/templates/listTasks.gohtml", "ui/templates/listTask.gohtml"))
