// Copyright 2016 The Solvist Author(s). All rights reserved.

package data

import (
	"gopkg.in/mgo.v2"
)

var sess *mgo.Session

// Collection codes for database collection.
const (
	accountC = "accounts"
	taskC    = "tasks"
	noteC    = "notes"
)

// OpenDBSession reads an URL
func OpenDBSession(url string) (err error) {
	sess, err = mgo.Dial(url)
	return err
}
