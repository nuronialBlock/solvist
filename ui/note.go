// Copyright 2016 The Solvist Author(s). All rights reserved.

package ui

import (
	"bytes"
	"net/http"

	"github.com/gorilla/schema"
	"github.com/microcosm-cc/bluemonday"
	"github.com/nuronialBlock/solvist/solvist/data"
	"github.com/russross/blackfriday"
)

// ServeNoteNewForm serves a new note form.
func ServeNoteNewForm(w http.ResponseWriter, r *http.Request) {
	err := TplNoteNewForm.Execute(w, TplNoteNewFormValues{})
	if err != nil {
		ServeInternalServerError(w, r)
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

	textBytes := bytes.NewBufferString(formValues.Text).Bytes()
	unsafe := blackfriday.MarkdownCommon(textBytes)
	htmlBuf := bluemonday.UGCPolicy().SanitizeBytes(unsafe)
	html := bytes.NewBuffer(htmlBuf).String()
	note.Text = html

	err = note.Put()
	if err != nil {
		ServeInternalServerError(w, r)
	}

	http.Redirect(w, r, "/tasks", http.StatusSeeOther)
}

// ServeNotesList renders the notes.
func ServeNotesList(w http.ResponseWriter, r *http.Request) {
	notes := []data.Note{}
	notes, err := data.ListNotes()
	if err != nil {
		ServeInternalServerError(w, r)
		return
	}

	err = TplNotesList.Execute(w, TplNotesListValues{
		Notes: notes,
	})
	if err != nil {
		ServeInternalServerError(w, r)
		return
	}
}

func init() {
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
}