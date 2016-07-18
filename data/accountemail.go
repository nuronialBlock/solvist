// Copyright 2016 The Solvist Author(s). All rights reserved.

package data

import (
	"time"

	"github.com/asaskevich/govalidator"
)

// AccountEmail stores the informations of user Email.
type AccountEmail struct {
	Address     string `bson:"address"`
	AddressNorm string `bson:"address_norm"`

	Primary bool `bson:"primary"`

	Verified   bool      `bson:"verified"`
	VerifiedAt time.Time `bson:"verified_at"`
	Token      string    `bson:"token"`
}

// NewAccountEmail returns a new AccountEmail,
// if address doesn't contain any error.
func NewAccountEmail(addr string) (AccountEmail, error) {
	norm, err := govalidator.NormalizeEmail(addr)
	if err != nil {
		return AccountEmail{}, err
	}

	return AccountEmail{
		Address:     addr,
		AddressNorm: norm,
	}, nil
}
