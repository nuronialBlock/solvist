package data

import "labix.org/v2/mgo"

var sess *mgo.Session

const (
	accountC = "accounts"
	taskC    = "tasks"
)

func OpenDBSession(url string) (err error) {
	sess, err = mgo.Dial(url)
	return err
}
