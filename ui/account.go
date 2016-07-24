// Copyright 2016 The Solvist Author(s). All rights reserved.

package ui

import (
	"net/http"

	"labix.org/v2/mgo"

	"github.com/gorilla/context"
	"github.com/gorilla/schema"
	"github.com/nuronialBlock/solvist/solvist/data"
)

// ServeLogInForm serves login page.
func ServeLogInForm(w http.ResponseWriter, r *http.Request) {
	acc, ok := context.Get(r, "account").(*data.Account)
	if ok {
		ServeBadRequest(w, r)
		return
	}
	err := TplLogIn.Execute(w, TplLogInValues{
		Common: TplCommonValues{
			Account: acc,
		},
	})
	if err != nil {
		ServeInternalServerError(w, r)
		return
	}
}

// ServeRegisterForm serves the register page.
func ServeRegisterForm(w http.ResponseWriter, r *http.Request) {
	acc, ok := context.Get(r, "account").(*data.Account)
	if ok {
		ServeBadRequest(w, r)
		return
	}
	if acc != nil {
		ServeBadRequest(w, r)
		return
	}

	err := TplRegister.Execute(w, TplRegisterValues{
		Common: TplCommonValues{
			Account: acc,
		},
	})
	if err != nil {
		ServeInternalServerError(w, r)
		return
	}
}

// LoginFormValues stroes form values while logging in.
type LoginFormValues struct {
	Handle   string `schema:"handle"`
	Password string `schema:"pass"`
}

// HandleLogin handles login of a user.
func HandleLogin(w http.ResponseWriter, r *http.Request) {
	acc, ok := context.Get(r, "account").(*data.Account)
	if ok {
		ServeBadRequest(w, r)
		return
	}
	if acc != nil {
		ServeBadRequest(w, r)
		return
	}

	err := r.ParseForm()
	if err != nil {
		ServeHandleIncorrect(w, r)
		return
	}
	values := LoginFormValues{}
	decoder := schema.NewDecoder()
	err = decoder.Decode(&values, r.PostForm)
	if err != nil {
		ServeInternalServerError(w, r)
		return
	}

	acc, err = data.GetAccountByHandle(values.Handle)
	if err == mgo.ErrNotFound {
		ServeHandleIncorrect(w, r)
		return
	}
	if err != nil {
		ServeInternalServerError(w, r)
		return
	}
	m := acc.Password.Match(values.Password)
	if !m {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	sess, err := store.Get(r, "very-secret")
	if err != nil {
		ServeInternalServerError(w, r)
		return
	}
	sess.Values["accountID"] = acc.ID.Hex()
	sess.Save(r, w)
	http.Redirect(w, r, "/tasks", http.StatusSeeOther)
}

func init() {
	Router.NewRoute().
		Methods("Get").
		Path("/login").
		HandlerFunc(ServeLogInForm)
	Router.NewRoute().
		Methods("Post").
		Path("/login").
		HandlerFunc(HandleLogin)
	Router.NewRoute().
		Methods("Get").
		Path("/register").
		HandlerFunc(ServeRegisterForm)
}
