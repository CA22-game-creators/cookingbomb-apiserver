package user

import (
	"errors"

	"github.com/oklog/ulid/v2"
)

type ID ulid.ULID

func NewID(v ulid.ULID) (ID, error) {
	if (v == ulid.ULID{}) {
		return ID{}, errors.New("user id is nil")
	}

	return ID(v), nil
}
