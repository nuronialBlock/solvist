// Copyright 2016 The Solvist Author(s). All rights reserved.

package ui

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"

	"labix.org/v2/mgo/bson"

	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
	"github.com/microcosm-cc/bluemonday"
	"github.com/nuronialBlock/solvist/solvist/data"
	"github.com/russross/blackfriday"
)

// ServeNoteNewForm serves a new note form.
func ServeNoteNewForm(w http.ResponseWriter, r *http.Request) {
	acc, ok := context.Get(r, "account").(*data.Account)
	if !ok {
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}

	err := TplNoteNewForm.Execute(w, TplNoteNewFormValues{
		Common: TplCommonValues{
			Account: acc,
		},
	})
	if err != nil {
		ServeInternalServerError(w, r)
		return
	}
}

// NoteValues stores the values of the note.
type NoteValues struct {
	ProblemName string `schema:"name"`
	ProblemOJ   string `schema:"oj"`
	ProblemID   string `schema:"id"`
	ProblemURL  string `schema:"url"`
	TopicName   string `schema:"topic"`
	Catagory    string `schema:"catagory"`
	Text        string `schema:"text"`
}

// HandleNoteCreate handles to create new note.
func HandleNoteCreate(w http.ResponseWriter, r *http.Request) {
	acc, ok := context.Get(r, "account").(*data.Account)
	if !ok {
		ServeBadRequest(w, r)
		return
	}

	err := r.ParseForm()
	if err != nil {
		ServeInternalServerError(w, r)
	}

	formValues := NoteValues{}
	decoder := schema.NewDecoder()
	err = decoder.Decode(&formValues, r.PostForm)
	if err != nil {
		ServeInternalServerError(w, r)
	}

	note := data.Note{}
	note.ProblemName = formValues.ProblemName
	note.ProblemOJ = formValues.ProblemOJ
	note.ProblemID = formValues.ProblemID
	note.ProblemURL = formValues.ProblemURL
	note.TopicName = formValues.TopicName
	note.Catagory = formValues.Catagory
	note.Text = formValues.Text
	note.AccountID = acc.ID

	err = note.Put()
	if err != nil {
		ServeInternalServerError(w, r)
		return
	}

	http.Redirect(w, r, "/notes", http.StatusSeeOther)
}

// ServeNotesList renders the notes.
func ServeNotesList(w http.ResponseWriter, r *http.Request) {
	acc, ok := context.Get(r, "account").(*data.Account)
	if !ok {
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}

	notes := []data.Note{}
	notes, err := data.ListNotesByAccountID(acc.ID)
	if err != nil {
		ServeInternalServerError(w, r)
		return
	}

	err = TplNotesList.Execute(w, TplNotesListValues{
		Common: TplCommonValues{
			Account: acc,
		},
		Notes: notes,
	})
	if err != nil {
		ServeInternalServerError(w, r)
		return
	}
}

// ServeNote renders note of a given note ID.
func ServeNote(w http.ResponseWriter, r *http.Request) {
	acc, ok := context.Get(r, "account").(*data.Account)
	if !ok {
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}

	vars := mux.Vars(r)
	idStr := vars["id"]
	if !bson.IsObjectIdHex(idStr) {
		ServeNotFound(w, r)
		return
	}

	id := bson.ObjectIdHex(idStr)
	note, err := data.GetNote(id)
	if err != nil {
		ServeInternalServerError(w, r)
		return
	}
	if note.Public == 0 {
		if note.AccountID != acc.ID {
			ServeBadRequest(w, r)
			return
		}
	}

	err = TplNoteView.Execute(w, TplNoteValues{
		Common: TplCommonValues{
			Account: acc,
		},
		Note: note,
	})
	if err != nil {
		ServeInternalServerError(w, r)
		return
	}
}

// ServeNoteEditForm serves the edit form page.
func ServeNoteEditForm(w http.ResponseWriter, r *http.Request) {
	acc, ok := context.Get(r, "account").(*data.Account)
	if !ok {
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}

	vars := mux.Vars(r)
	idStr := vars["id"]
	if !bson.IsObjectIdHex(idStr) {
		ServeNotFound(w, r)
		return
	}

	id := bson.ObjectIdHex(idStr)
	note, err := data.GetNote(id)
	if err != nil {
		ServeInternalServerError(w, r)
		return
	}
	if acc.ID != note.AccountID {
		ServeBadRequest(w, r)
		return
	}
	if note == nil {
		ServeNotFound(w, r)
		return
	}

	err = TplNoteEditForm.Execute(w, TplNoteEditFormValues{
		Common: TplCommonValues{
			Account: acc,
		},
		Note: note,
	})
	if err != nil {
		ServeInternalServerError(w, r)
		return
	}
}

// EditedNoteValues stores edited values of a note.
type EditedNoteValues struct {
	ProblemName string `schema:"name"`
	ProblemOJ   string `schema:"oj"`
	ProblemID   string `schema:"id"`
	ProblemURL  string `schema:"url"`
	TopicName   string `schema:"topic"`
	Catagory    string `schema:"catagory"`
	Text        string `schema:"text"`
	Public      int    `schema:"public"`
}

// HandleNoteSave saves the edited info of a note.
func HandleNoteSave(w http.ResponseWriter, r *http.Request) {
	acc, ok := context.Get(r, "account").(*data.Account)
	if !ok {
		ServeBadRequest(w, r)
		return
	}

	vars := mux.Vars(r)
	idStr := vars["id"]
	if !bson.IsObjectIdHex(idStr) {
		ServeNotFound(w, r)
		return
	}
	id := bson.ObjectIdHex(idStr)
	note, err := data.GetNote(id)
	if err != nil {
		ServeInternalServerError(w, r)
		return
	}
	if note.AccountID != acc.ID {
		ServeBadRequest(w, r)
		return
	}
	if note == nil {
		ServeNotFound(w, r)
		return
	}

	err = r.ParseForm()
	if err != nil {
		ServeInternalServerError(w, r)
		return
	}
	formValues := EditedNoteValues{}
	decoder := schema.NewDecoder()
	err = decoder.Decode(&formValues, r.PostForm)
	if err != nil {
		ServeInternalServerError(w, r)
		return
	}

	note.ProblemName = formValues.ProblemName
	note.ProblemOJ = formValues.ProblemOJ
	note.ProblemID = formValues.ProblemID
	note.ProblemURL = formValues.ProblemURL
	note.TopicName = formValues.TopicName
	note.Catagory = formValues.Catagory
	note.Text = formValues.Text
	note.Public = formValues.Public
	note.Writer = acc.Handle

	err = note.Put()
	if err != nil {
		ServeInternalServerError(w, r)
		return
	}

	http.Redirect(w, r, "/notes", http.StatusSeeOther)
}

// HandleNoteRemove removes a note from database.
func HandleNoteRemove(w http.ResponseWriter, r *http.Request) {
	acc, ok := context.Get(r, "account").(*data.Account)
	if !ok {
		ServeBadRequest(w, r)
		return
	}

	vars := mux.Vars(r)
	idStr := vars["id"]
	if !bson.IsObjectIdHex(idStr) {
		ServeNotFound(w, r)
		return
	}

	id := bson.ObjectIdHex(idStr)
	note, err := data.GetNote(id)
	if err != nil {
		ServeInternalServerError(w, r)
		return
	}
	if note.AccountID != acc.ID {
		ServeBadRequest(w, r)
		return
	}
	if note == nil {
		ServeNotFound(w, r)
		return
	}

	err = note.Remove()
	if err != nil {
		ServeInternalServerError(w, r)
		return
	}

	http.Redirect(w, r, "/notes", http.StatusSeeOther)
}

// ServePublicNotes serves notes which have been
// made public by the users.
func ServePublicNotes(w http.ResponseWriter, r *http.Request) {
	acc, ok := context.Get(r, "account").(*data.Account)
	if !ok || acc == nil {
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}
	notes, err := data.ListNotesByPublic()
	if err != nil {
		ServeInternalServerError(w, r)
		return
	}

	err = TplPublicNotes.Execute(w, TplPublicNotesValues{
		Common: TplCommonValues{
			Account: acc,
		},
		Notes: notes,
	})
	if err != nil {
		fmt.Println(err)
		ServeInternalServerError(w, r)
		return
	}
}

// Markdown parse a Markdown to HTML.
func Markdown(m string) template.HTML {
	textBytes := bytes.NewBufferString(m).Bytes()
	unsafe := blackfriday.MarkdownCommon(textBytes)
	html := bluemonday.UGCPolicy().SanitizeBytes(unsafe)
	return template.HTML(html)
}

func init() {
	Router.NewRoute().
		Methods("Get").
		Path("/notes/edit/{id}").
		HandlerFunc(ServeNoteEditForm)
	Router.NewRoute().
		Methods("Post").
		Path("/notes/edit/{id}").
		HandlerFunc(HandleNoteSave)
	Router.NewRoute().
		Methods("Get").
		Path("/notes/new").
		HandlerFunc(ServeNoteNewForm)
	Router.NewRoute().
		Methods("Post").
		Path("/notes/new").
		HandlerFunc(HandleNoteCreate)
	Router.NewRoute().
		Methods("Get").
		Path("/notes").
		HandlerFunc(ServeNotesList)
	Router.NewRoute().
		Methods("Get").
		Path("/notes/{id}").
		HandlerFunc(ServeNote)
	Router.NewRoute().
		Methods("Post").
		Path("/notes/remove/{id}").
		HandlerFunc(HandleNoteRemove)
	Router.NewRoute().
		Methods("Get").
		Path("/public").
		HandlerFunc(ServePublicNotes)
}
