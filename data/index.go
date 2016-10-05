package data

import "gopkg.in/mgo.v2"

var indices = map[string][]mgo.Index{
	accountC: {
		{
			Key:    []string{"handle", "emails.address_norm"},
			Unique: true,
		},
	},
}

// MakeIndex ensures indexing of an unique key.
func MakeIndex() error {
	for col, idxs := range indices {
		for _, idx := range idxs {
			err := sess.DB("").C(col).EnsureIndex(idx)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
