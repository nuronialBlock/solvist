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
func GetAccount(id bson.ObjectId) (Account, error) {
	acc := Account{}
	err := sess.DB("").C(accountC).FindId(id).One(&acc)
	if err != nil {
		return acc, err
	}

	return acc, nil
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

// Remove removes an account form database.
func (a *Account) Remove() error {
	err := sess.DB("").C(accountC).RemoveId(a.ID)
	if err != nil {
		return err
	}

	return nil
}
