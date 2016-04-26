// Copyright 2016 The Solvist Author(s). All rights reserved.

package data

// AccountPassword stores the password informations.
type AccountPassword struct {
	Salt       []byte `bson:"salt"`
	Iteration  int    `bson:"iteration"`
	KeyLength  int    `bson:"key_length"`
	DerivedKey []byte `bson:"derived_key"`
	Algorithm  string `bson:"algorithm"`
}
