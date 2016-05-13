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
	ProblemURL  string        `bson:"problem_url"`

	ModifiedAt time.Time `bson:"modified_at"`
	CreatedAt  time.Time `bson:"created_at"`

	Solved bool `bson:"solved"`
}

// ListTasks retrieves a list of tasks form database, if successful,
// an array of Task object that can be used for showing tasks list.
func ListTasks() ([]Task, error) {
	tasks := []Task{}
	err := sess.DB("").C(taskC).Find(nil).All(&tasks)
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

// Remove removes a task from database.
func (t *Task) Remove() error {
	err := sess.DB("").C(taskC).RemoveId(t.ID)
	return err
}
