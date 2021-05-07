package user

import (
	"errors"
)

type HashedAuthToken []byte

func NewHashedAuthToken(v []byte) (HashedAuthToken, error) {
	if v == nil {
		return nil, errors.New("user hashed_auth_token is nil")
	}

	return HashedAuthToken(v), nil
}
