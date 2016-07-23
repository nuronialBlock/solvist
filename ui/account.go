// Copyright 2016 The Solvist Author(s). All rights reserved.

package ui

import (
	"net/http"

	"github.com/gorilla/context"
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

func init() {
	Router.NewRoute().
		Methods("Get").
		Path("/login").
		HandlerFunc(ServeLogInForm)
	Router.NewRoute().
		Methods("Get").
		Path("/register").
		HandlerFunc(ServeRegisterForm)
}
