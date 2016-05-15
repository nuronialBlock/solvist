// Copyright 2016 The Solvist Author(s). All rights reserved.

package ui

import (
	"html/template"

	"github.com/nuronialBlock/solvist/solvist/data"
)

// TplTaskNewFormValues passes the new form values.
type TplTaskNewFormValues struct{}

// TplTaskNewForm generates new form for user.
var TplTaskNewForm = template.Must(template.ParseFiles("ui/templates/layout.gohtml", "ui/templates/taskNewFormPage.gohtml", "ui/templates/taskNewForm.gohtml"))

// TplListTasksValues stores the values of tasks.
type TplListTasksValues struct {
	Tasks []data.Task
}

// TplListTasks renders Tasks.
var TplListTasks = template.Must(template.ParseFiles("ui/templates/layout.gohtml", "ui/templates/listTasksPage.gohtml", "ui/templates/listTasks.gohtml", "ui/templates/listTask.gohtml"))

// TplTaskEditFormValues stores the values of the edited task.
type TplTaskEditFormValues struct {
	Task data.Task
}

// TplTaskEditForm renders task edit.
var TplTaskEditForm = template.Must(template.ParseFiles("ui/templates/layout.gohtml", "ui/templates/editTaskFormPage.gohtml", "ui/templates/editTaskForm.gohtml"))
