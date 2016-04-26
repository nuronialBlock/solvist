// Copyright 2016 The Solvist Author(s). All rights reserved.

package data

import (
	"time"

	"labix.org/v2/mgo/bson"
)

// Task stores the task informations of an user.
type Task struct {
	ID          bson.ObjectId `bson:"_id"`
	ProblemName string        `bson:"problem_name"`
	ProblemOJ   string        `bson:"problem_oj"`
	ProblemID   string        `bson:"problem_id"`

	ModifiedAt time.Time `bson:"modified_at"`
	CreatedAt  time.Time `bson:"created_at"`

	Solved bool `bson:"solved"`
}

// ListTasksAccount retrieves a list of tasks form database for ID and returns,
// if successful, an array of Task object that can be used for showing tasks list.
func ListTasksAccount(id bson.ObjectId) ([]Task, error) {
	tasks := []Task{}
	err := sess.DB("").C(accountC).FindId(id).All(&tasks)
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

// Put puts data into database.
func (t *Task) Put() error {
	t.ModifiedAt = time.Now()
	if t.ID == "" {
		t.CreatedAt = t.ModifiedAt
		t.ID = bson.NewObjectId()
	}

	_, err := sess.DB("").C(taskC).UpsertId(t.ID, t)
	return err
}
