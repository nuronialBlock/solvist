// Copyright @nuronial_block
package data

import (
	"time"

	"labix.org/v2/mgo/bson"
)

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
