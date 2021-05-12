package user

import (
	"github.com/oklog/ulid/v2"

	"github.com/CA22-game-creators/cookingbomb-apiserver/errors"
)

type ID ulid.ULID

func NewID(v ulid.ULID) (ID, error) {
	if (v == ulid.ULID{}) {
		return ID{}, errors.Internal("user id is nil")
	}

	return ID(v), nil
}
