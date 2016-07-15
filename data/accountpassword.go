// Copyright 2016 The Solvist Author(s). All rights reserved.

package data

import (
	"crypto/rand"
	"crypto/sha1"

	"golang.org/x/crypto/pbkdf2"
)

// AccountPassword stores the password informations.
type AccountPassword struct {
	Salt       []byte `bson:"salt"`
	Iteration  int    `bson:"iteration"`
	KeyLength  int    `bson:"key_length"`
	DerivedKey []byte `bson:"derived_key"`
	Algorithm  string `bson:"algorithm"`
}

// NewAccountPassword generates an encrypted password.
func NewAccountPassword(clear string) (AccountPassword, error) {
	accPass := AccountPassword{}
	accPass.Salt = make([]byte, 32)
	_, err := rand.Read(accPass.Salt)
	if err != nil {
		return AccountPassword{}, err
	}

	accPass.Algorithm = "SHA1"
	accPass.Iteration = 4096
	accPass.KeyLength = 32
	accPass.DerivedKey = pbkdf2.Key(clear, accPass.Salt, accPass.Iteration, accPass.KeyLength, sha1.New)
	return accPass, nil
}

// IsValid checks whether a key's validity.
func (p AccountPassword) IsValid() bool {
	return len(p.DerivedKey) != 0
}
