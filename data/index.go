package data

import "labix.org/v2/mgo"

var indices = map[string][]mgo.Index{
	accountC: {
		{
			Key:    []string{"Handle"},
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
