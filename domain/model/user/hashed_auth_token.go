package user

import "github.com/CA22-game-creators/cookingbomb-apiserver/errors"

type HashedAuthToken []byte

func NewHashedAuthToken(v []byte) (HashedAuthToken, error) {
	if v == nil {
		return nil, errors.Internal("user hashed_auth_token is nil")
	}

	return HashedAuthToken(v), nil
}
