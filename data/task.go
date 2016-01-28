// Copyright @nuronial_block
package data

import "labix.org/v2/mgo/bson"

type Task struct {
	ID          bson.ObjectId `bson:"_id"`
	ProblemName string        `bson:"problem_name"`
	ProblemOj   string        `bson:"problem_oj"`
	ProblemID   string        `bson:"problem_id"`

	Solved bool `bson:"solved"`
}
