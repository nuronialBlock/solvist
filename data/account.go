// Copyright 2016 The Solvist Author(s). All rights reserved.

// Package data contains
// all the data structures and database related codes.
package data

import (
	"time"

	"labix.org/v2/mgo/bson"
)

// Account stores the informations of an user.
type Account struct {
	ID     bson.ObjectId `bson:"_id"`
	Name   string        `bson:"name"`
	Handle string        `bson:"handle"`

	Password AccountPassword `bson:"password"`
	Emails   []AccountEmail  `bson:"emails"`
	Banned   bool            `bson:"banned"`

	CreatedAt  time.Time `bson:"created_at"`
	ModifiedAt time.Time `bson:"modified_at"`
}

// ListAccounts function gives an array of all accounts,
// and give an error if any occurs.
func ListAccounts() ([]Account, error) {
	accs := []Account{}
	err := sess.DB("").C(accountC).Find(nil).All(&accs)
	if err != nil {
		return nil, err
	}

	return accs, nil
}
