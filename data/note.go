// Copyright 2016 The Solvist Author(s). All rights reserved.

package data

import (
	"time"

	"labix.org/v2/mgo/bson"
)

// Note stores note of a problem.
type Note struct {
	ID bson.ObjectId `bson:"_id"`

	ProblemName string `bson:"problem_name"`
	ProblemOJ   string `bson:"problem_oj"`
	ProblemID   string `bson:"problem_id"`
	ProblemURL  string `bson:"problem_url"`

	Text      string `bson:"text"`
	TopicName string `bson:"topic_name"`
	Catagory  string `bson:"catagory"`

	CreatedAt  time.Time `bson:"created_at"`
	ModifiedAt time.Time `bson:"modified_at"`
}

// ListNotes lits all notes.
func (n *Note) ListNotes() ([]Note, error) {
	notes := []Note{}
	err := sess.DB("").C(noteC).Find(nil).All(notes)
	if err != nil {
		return nil, err
	}

	return notes, nil
}
