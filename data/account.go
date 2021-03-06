// Copyright 2016 The Solvist Author(s). All rights reserved.

// Package data contains
// all the data structures and database related codes.
package data

import (
	"time"

	"gopkg.in/mgo.v2"

	"gopkg.in/mgo.v2/bson"
)

// Account stores the informations of an user.
type Account struct {
	ID         bson.ObjectId `bson:"_id"`
	Name       string        `bson:"name"`
	Handle     string        `bson:"handle"`
	University string        `bson:"university"`
	Country    string        `bson:"country"`

	Password AccountPassword `bson:"password"`
	Emails   []AccountEmail  `bson:"emails"`
	Banned   bool            `bson:"banned"`

	CreatedAt  time.Time `bson:"created_at"`
	ModifiedAt time.Time `bson:"modified_at"`
}

// GetAccount retrives an account by it's ID.
func GetAccount(id bson.ObjectId) (*Account, error) {
	acc := Account{}
	err := sess.DB("").C(accountC).FindId(id).One(&acc)
	if err != nil {
		return nil, err
	}

	return &acc, nil
}

// GetAccountByHandle funtion takes handle as input, and
// returns Account information,
// returns error if error occured.
func GetAccountByHandle(handle string) (*Account, error) {
	acc := Account{}
	err := sess.DB("").C(accountC).Find(bson.M{"handle": handle}).One(&acc)
	if err != nil {
		return nil, err
	}
	return &acc, nil
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

// IsDup checks for duplicate key error for account.
func (a *Account) IsDup() bool {
	err := a.Put()
	return mgo.IsDup(err)
}

// Put puts data in the database.
func (a *Account) Put() error {
	a.ModifiedAt = time.Now()
	if a.ID == "" {
		a.CreatedAt = a.ModifiedAt
		a.ID = bson.NewObjectId()
	}
	_, err := sess.DB("").C(accountC).UpsertId(a.ID, a)
	return err
}

// Remove removes an account form database.
func (a *Account) Remove() error {
	err := sess.DB("").C(accountC).RemoveId(a.ID)
	if err != nil {
		return err
	}

	return nil
}
