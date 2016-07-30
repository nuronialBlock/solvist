// Copyright 2016 The Solvist Author(s). All rights reserved.

package ui

import (
	"html/template"

	"github.com/nuronialBlock/solvist/solvist/data"
)

// TplTaskNewFormValues passes the new form values.
type TplTaskNewFormValues struct {
	Common TplCommonValues
}

// TplTaskNewForm generates new form for user.
var TplTaskNewForm = template.Must(template.ParseFiles("ui/templates/layout.gohtml", "ui/templates/taskNewFormPage.gohtml", "ui/templates/taskNewForm.gohtml"))

// TplTasksListValues stores the values of tasks.
type TplTasksListValues struct {
	Common TplCommonValues
	Tasks  []data.Task
}

// TplTasksList renders Tasks.
var TplTasksList = template.Must(template.ParseFiles("ui/templates/layout.gohtml", "ui/templates/tasksListPage.gohtml", "ui/templates/tasksList.gohtml", "ui/templates/tasksListItem.gohtml"))

// TplTaskEditFormValues stores the values of the edited task.
type TplTaskEditFormValues struct {
	Task data.Task
}

// TplTaskEditForm renders task edit.
var TplTaskEditForm = template.Must(template.ParseFiles("ui/templates/layout.gohtml", "ui/templates/editTaskFormPage.gohtml", "ui/templates/editTaskForm.gohtml"))

// TplNoteNewFormValues stores the new form values.
type TplNoteNewFormValues struct{}

// TplNoteNewForm generates new note form for user.
var TplNoteNewForm = template.Must(template.ParseFiles("ui/templates/layout.gohtml", "ui/templates/noteNewFormPage.gohtml", "ui/templates/noteNewForm.gohtml"))

// TplNotesListValues stores the value of notes.
type TplNotesListValues struct {
	Notes []data.Note
}

// TplNotesList serves the notes list page.
var TplNotesList = template.Must(template.ParseFiles("ui/templates/layout.gohtml", "ui/templates/notesListPage.gohtml", "ui/templates/notesList.gohtml", "ui/templates/notesListItem.gohtml"))

// TplNoteValues stores the note value for rendering a note.
type TplNoteValues struct {
	Note data.Note
}

// TplNote stores note's template.
var TplNote = template.Must(template.ParseFiles("ui/templates/layout.gohtml", "ui/templates/notePage.gohtml"))

// TplNoteView renders a note.
var TplNoteView = template.Must(TplNote.Funcs(template.FuncMap{"Markdown": Markdown}).ParseFiles("ui/templates/noteView.gohtml"))

// TplNoteEditFormValues stores edited note values.
type TplNoteEditFormValues struct {
	Note data.Note
}

// TplNoteEditForm render note edit form page.
var TplNoteEditForm = template.Must(template.ParseFiles("ui/templates/layout.gohtml", "ui/templates/noteEditFormPage.gohtml", "ui/templates/noteEditForm.gohtml"))

// TplLogInValues stores template log in values.
type TplLogInValues struct {
	Common TplCommonValues
}

// TplLogIn renders log in page.
var TplLogIn = template.Must(template.ParseFiles("ui/templates/layout.gohtml", "ui/templates/loginPage.gohtml", "ui/templates/login.gohtml"))

// TplCommonValues stores common values for templates.
type TplCommonValues struct {
	Account *data.Account
}

// TplRegisterValues stores values while rendering register form.
type TplRegisterValues struct {
	Common TplCommonValues
}

// TplRegister renders the register form.
var TplRegister = template.Must(template.ParseFiles("ui/templates/layout.gohtml", "ui/templates/registerPage.gohtml", "ui/templates/register.gohtml"))
