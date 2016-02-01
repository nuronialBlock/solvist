// Copyright @nuronial_block
package data

import (
	"time"

	"labix.org/v2/mgo/bson"
)

type Task struct {
	ID          bson.ObjectId `bson:"_id"`
	ProblemName string        `bson:"problem_name"`
	ProblemOj   string        `bson:"problem_oj"`
	ProblemID   string        `bson:"problem_id"`

	ModifiedAt time.Time `bson:"modified_at"`
	CreatedAt  time.Time `bson:"created_at"`

	Solved bool `bson:"solved"`
}

func (t *Task) Put() error {
	t.ModifiedAt = time.Now()
	if t.ID == "" {
		t.CreatedAt = t.ModifiedAt
		t.ID = bson.NewObjectId()
	}

	_, err := sess.DB("").C(taskC).UpsertId(t.ID, t)
	return err
}
