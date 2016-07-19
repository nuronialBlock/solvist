// Copyright 2016 The Solvist Author(s). All rights reserved.

package ui

import (
	"net/http"

	"github.com/gorilla/context"
	"github.com/nuronialBlock/solvist/solvist/data"
)

func serveLogIn(w http.ResponseWriter, r *http.Request) {
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

// ServeRegister serves the register page.
func ServeRegister(w http.ResponseWriter, r *http.Request) {
	acc, ok := context.Get(r, "account").(*data.Account)
	if !ok {
		ServeBadRequest(w, r)
		return
	}
	if acc != nil {
		ServeBadRequest(w, r)
		return
	}

	err := TplRegister.Execute(w, TplRegisterValues{
		CommonValues: TplCommonValues{
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
		HandlerFunc(serveLogIn)
	Router.NewRoute().
		Methods("Get").
		Path("/register").
		HandlerFunc(ServeRegister)
}
