// Copyright 2016 The Solvist Author(s). All rights reserved.

package data

import (
	"labix.org/v2/mgo"
)

var sess *mgo.Session

// Collection codes for database collection.
const (
	accountC = "accounts"
	taskC    = "tasks"
)

// OpenDBSession reads an URL
func OpenDBSession(url string) (err error) {
	sess, err = mgo.Dial(url)
	return err
}
